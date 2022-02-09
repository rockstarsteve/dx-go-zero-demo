package logic

import (
	"context"

	"go-zero-demo/jwttoken/internal/svc"
	"go-zero-demo/jwttoken/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) TestLogic {
	return TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test() (resp *types.InfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
