package logic

import (
	"context"

	"go-zero-demo/jwttoken/internal/svc"
	"go-zero-demo/jwttoken/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	//// todo: add your logic here and delete this line
	//
	//return
	return &types.GetUserResponse{Name: "kim"}, nil
}
