package wework

import (
	"net/http"
	wxcpt "wechat_demo/wechat-work/wxbizmsgcrypt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type WeWorkController struct {
	Token          string
	EncodingAESKey string
	ReceiverID     string
	WXBizMsgCrypt  *wxcpt.WXBizMsgCrypt
}

func NewWeWorkController() *WeWorkController {
	ctrl := &WeWorkController{
		Token:          viper.GetString("WEWORK_TOKEN"),
		EncodingAESKey: viper.GetString("WEWORK_ENCODE_AES_KEY"),
		ReceiverID:     viper.GetString("WEWORK_RECV_ID"),
	}
	ctrl.WXBizMsgCrypt = wxcpt.NewWXBizMsgCrypt(ctrl.Token, ctrl.EncodingAESKey, ctrl.ReceiverID, wxcpt.XmlType)
	return ctrl
}

func (ctrl *WeWorkController) VerifyURL(ctx *gin.Context) {
	verifyMsgSign := ctx.Query("msg_signature")
	verifyTimestamp := ctx.Query("timestamp")
	verifyNonce := ctx.Query("nonce")
	verifyEchoStr := ctx.Query("echostr")
	echoStr, cryptErr := ctrl.WXBizMsgCrypt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	if cryptErr != nil {
		ctx.JSON(http.StatusBadRequest, cryptErr)
		return
	}
	ctx.String(http.StatusOK, string(echoStr))
}

func (ctrl *WeWorkController) HandleMsg(ctx *gin.Context) {

}
