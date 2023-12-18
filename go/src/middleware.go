package middleware

import (
    "net/http"
    "os"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

var secretKey = os.Getenv("SECRET_KEY")

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // 署名検証用の鍵
            return []byte(secretKey), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        claims := token.Claims.(jwt.MapClaims)
        c.Set("claims", claims)
        c.Next()
    }
}