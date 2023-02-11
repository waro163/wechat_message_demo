module wechat_demo

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.5
	github.com/silenceper/wechat/v2 v2.1.3
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/viper v1.11.0
	github.com/waro163/wechat-message v0.0.0-20221102074714-f9423deccb1d
)

replace github.com/silenceper/wechat/v2 => /Users/waro/work/wechat

replace github.com/waro163/wechat-message => /Users/waro/work/wechat-message
