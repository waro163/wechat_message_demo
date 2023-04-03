package main

import (
	"net/http"
	_ "wechat_demo/config"
	_ "wechat_demo/log"
	"wechat_demo/wechat"
	wechatmsg "wechat_demo/wechat-msg"
	wework "wechat_demo/wechat-work"

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

	event := router.Group("/api/v1/wechatmsg")
	wechatmsg.RegisterRouter(event)

	work := router.Group("/api/v1/wework")
	wework.RegisterRouter(work)

	router.Run(":8080")
}
