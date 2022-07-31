package wechat

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

func WechatLogin(c *gin.Context) {
	req := c.Request
	// host2 := req.URL
	// fmt.Printf("%#v\n", host2)
	uri := fmt.Sprintf("https://%s/api/v1/wechat/user_info", req.Host)
	log.Debugf("redirect uri %s\n", uri)
	auth.Redirect(c.Writer, c.Request, uri, "snsapi_userinfo", "")
	// c.Redirect(http.StatusTemporaryRedirect, "/ping")
}

func WechatUserInfo(c *gin.Context) {
	code, exist := c.GetQuery("code")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "lost query parameters 'code'",
		})
		return
	}
	res, err := auth.GetUserAccessToken(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "server error" + err.Error()})
		return
	}
	user_info, err := auth.GetUserInfo(res.AccessToken, res.OpenID, "zh_CN")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "server error" + err.Error()})
		return
	}
	// fmt.Printf("%#v\n", user_info)
	log.Debug(user_info)
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
