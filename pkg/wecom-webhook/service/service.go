package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rosenlo/wecom-webhook/pkg/wecom-webhook/service/v1"
)

type Service interface {
	WebService() *gin.Engine
}

func NewService() Service {
	return new(v1.Service)
}
