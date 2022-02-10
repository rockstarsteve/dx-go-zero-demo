package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/mall/sys/rpc/dicclient"
	"go-zero-demo/mall/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DicRpc dicclient.Dic
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DicRpc: dicclient.NewDic(zrpc.MustNewClient(c.DicRpc)),
	}
}
