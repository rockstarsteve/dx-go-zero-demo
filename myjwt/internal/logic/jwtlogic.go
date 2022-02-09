package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/myjwt/internal/svc"
)

type JwtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) JwtLogic {
	return JwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JwtLogic) Jwt() error {
	// todo: add your logic here and delete this line

	return nil
}
