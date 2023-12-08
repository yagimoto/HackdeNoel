package main

import (
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", func(ctx *gin.Context) {
        ctx.String(200, "Hello World")
    })

    if err := router.Run(); err != nil {
        log.Fatal("Server Run Failed.: ", err)
    }
}