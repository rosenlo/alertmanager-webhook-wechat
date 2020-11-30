package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rosenlo/toolkits/log"
	"github.com/spf13/viper"
)

type Service struct{}

func (s *Service) WebService() *gin.Engine {
	routes := gin.Default()
	routes.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, I'm WeCom WebHook for Alertmanager")
	})
	routes.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	routes.POST("/webhook", s.Webhook)

	log.Debugf("will start with address: %v", viper.GetString("address"))
	return routes
}
