package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rosenlo/alertmanager-webhook-wechat/api/v1"
)

type Service interface {
	WebService() *gin.Engine
}

func New() Service {
	return new(v1.Service)
}
