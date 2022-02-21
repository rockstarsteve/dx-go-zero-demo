package logic

import (
	"context"

	"go-zero-demo/mall/user/rpc/internal/svc"
	user "go-zero-demo/mall/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserLogic {
	return &SaveUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserLogic) SaveUser(in *user.SaveReq) (*user.SaveResp, error) {
	err := l.svcCtx.Model.SaveTras()
	if err != nil {
		return &user.SaveResp{}, err
	}

	return &user.SaveResp{
		Code: 200,
		Msg: "保存多个表成功",
	}, nil
}
