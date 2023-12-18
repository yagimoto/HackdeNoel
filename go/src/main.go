package main

import (
    "log"

    "github.com/gin-gonic/gin"
)

type Scores struct {
    ScoreID int `json:"score_id"`
    Score int `json:"score"`
    UserID int `json:"user_id"`
}

type Ranking struct{
    Scores []Scores
}

type User struct {
    UserID int `json:"user_id"`
    UserName string `json:"user_name"`
    Password string `json:"password"`
}

func main() {
    router := gin.Default()

    router.GET("/login", login)
    router.GET("/logout", logout)
    router.GET("/sign-up", signUp)
    router.GET("/ranking", showRanking)
    router.GET("/score", getScore)
    router.GET("/result", getResult)

    if err := router.Run(); err != nil {
        log.Fatal("Server Run Failed.: ", err)
    }

}

func showRanking(c *gin.Context){
    // model.GetRanking()をJSONにして返す
}

func getResult(c *gin.Context){
    user_id := getUserIDforHeader(c)
    score := c.Query("score")

    HighScore := model.GetHighScore(user_id)

    if(score > HighScore) {
        model.updateHighScore(userid)
    } else {
        c.JSON(200, gin.H{
			"score": score,
		})
    }
    // model.GetHighScore()
    // 比較
    // ハイスコアならmodel.updateHighScore()
    // じゃないならそのままの値を返す
}

func getScore(c *gin.Context){
    // model.getScore()
    // 返ってきた値を足し算する
}
func login(c *gin.Context) {
    username := c.Query("username")
    password := c.Query("password")

    // model.searchUser()
    // 存在してたらmodel.createToken()
}

func logout(c *gin.Context) {
    // model.DeleteToken();
}

func signUp(c *gin.Context) {
    username := c.Query("username")
    password := c.Query("password")

    // model.searchUserで同じ名前のアカウントがないか確認
    if(model.searchUser(username)) {

        // なければ、model.createUserで新規アカウントをDBに登録
        // トークンを生成
        user := User{userid, username, password}
        model.createUser(user)
    }
    
}

func getUserIDforHeader(c *gin.Context) int {
    claims := c.MustGet("claims").(jwt.MapClaims)
    userID := int(claims["user_id"].(float64)) // ユーザーIDの取得
    return userID
}
