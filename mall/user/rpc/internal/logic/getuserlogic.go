package logic

import (
	"context"
	"go-zero-demo/mall/sys/rpc/dicclient"
	"go-zero-demo/mall/user/rpc/internal/svc"
	user "go-zero-demo/mall/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	//从dic服务获取性别
	genderId := int64(1)

	dic, err := l.svcCtx.DicRpc.GetDicById(l.ctx, &dicclient.IdRequest{
		Id: genderId,
	})
	sysUser, err2 := l.svcCtx.Model.FindOne(genderId)
	if err2 != nil || err != nil {
		return nil, err
	}

	return &user.UserResponse{
		Id: string(sysUser.Id),
		Name:   sysUser.Name,
		Gender: dic.Val + " ->这个是从字典rpc表中查询到的数据",
		Age:    12,
	}, nil
}
