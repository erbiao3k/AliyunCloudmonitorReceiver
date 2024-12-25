package router

import (
	"AliyunCloudmonitorReceiver/config"
	"AliyunCloudmonitorReceiver/controller"
	"AliyunCloudmonitorReceiver/public"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 运维告警中心接口/msg/webhook
	r.POST(config.AlertWebhook+":id", controller.WebHook)

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"errmsg":   "MethodNotAllowed",
			"time":     public.NowTime(),
			"yourip":   c.ClientIP(),
			"resource": c.Request.Host + c.Request.RequestURI,
		})
	})

	return r
}
