package wechat

import (
	_ "wechat_demo/config"

	"github.com/silenceper/wechat/v2/cache"
	oa "github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/spf13/viper"
)

var memory = cache.NewMemory()

var _ cache.Cache = (*cache.Memory)(nil)

var conf = &config.Config{
	AppID:          viper.GetString("APP_ID"),
	AppSecret:      viper.GetString("APP_SECRET"),
	Token:          viper.GetString("TOKEN"),
	EncodingAESKey: viper.GetString("AES_KEY"),
	Cache:          memory,
}

var (
	mp   = oa.NewOfficialAccount(conf)
	auth = mp.GetOauth()
)

func handleMsg(mm *message.MixMessage) *message.Reply {
	if mm.MsgType == "text" {
		if mm.Content == "人工" {
			data := message.NewText("")
			return &message.Reply{MsgType: message.MsgTypeTransfer, MsgData: data}
		} else {
			data := message.NewText(mm.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: data}
		}
	}
	return nil
}
