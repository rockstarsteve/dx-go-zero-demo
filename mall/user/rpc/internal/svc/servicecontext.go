package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/mall/sys/rpc/dicclient"
	"go-zero-demo/mall/user/rpc/internal/config"
	"go-zero-demo/mall/user/rpc/internal/server/model"
)

type ServiceContext struct {
	Config config.Config
	DicRpc dicclient.Dic

	Model  model.SysUserModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DicRpc: dicclient.NewDic(zrpc.MustNewClient(c.DicRpc)),
		Model:  model.NewSysUserModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}
