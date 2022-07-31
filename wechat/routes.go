package wechat

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.Any("", HandleWechatEvent)
	r.GET("/login", WechatLogin)
	r.GET("/user_info", WechatUserInfo)
}
