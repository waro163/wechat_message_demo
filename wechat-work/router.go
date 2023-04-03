package wework

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	ctrl := NewWeWorkController()
	r.GET("/msg", ctrl.VerifyURL)
	r.POST("/msg", ctrl.HandleMsg)
}
