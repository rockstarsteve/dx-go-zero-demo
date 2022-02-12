package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/jwttoken/internal/config"
	"go-zero-demo/jwttoken/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
