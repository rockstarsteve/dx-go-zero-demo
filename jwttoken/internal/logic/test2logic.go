package logic

import (
	"context"

	"go-zero-demo/jwttoken/internal/svc"
	"go-zero-demo/jwttoken/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Test2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTest2Logic(ctx context.Context, svcCtx *svc.ServiceContext) Test2Logic {
	return Test2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Test2Logic) Test2() (resp *types.InfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
