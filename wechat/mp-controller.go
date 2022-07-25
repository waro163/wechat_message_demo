package wechat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleWechatEvent(c *gin.Context) {
	server := mp.GetServer(c.Request, c.Writer)
	server.SetMessageHandler(handleMsg)
	err := server.Serve()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	server.Send()
}
