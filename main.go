package main

import (
	"net/http"
	_ "wechat_demo/config"
	_ "wechat_demo/log"
	"wechat_demo/wechat"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Phone string `json:"phone" binding:"required,min=11,max=11"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	mp := router.Group("/api/v1/wechat")
	wechat.RegisterRouter(mp)

	router.Run(":8080")
}
