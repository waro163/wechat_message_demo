package wechatmsg

import (
	"github.com/gin-gonic/gin"
	"github.com/waro163/wechat-message/mini/serve"
	"github.com/waro163/wechat-message/mp"
	msg "github.com/waro163/wechat-message/mp/message"
)

func HandleMsg(mixMsg mp.MixMessage) (interface{}, error) {
	text := msg.NewText("hello wechat message")
	return text, nil
}

var server = &serve.Server{
	SkipValidata: true,
	LogMsg:       true,
	Handler:      HandleMsg,
}

func HandleWechatEvent() gin.HandlerFunc {
	return gin.WrapH(server)
}
