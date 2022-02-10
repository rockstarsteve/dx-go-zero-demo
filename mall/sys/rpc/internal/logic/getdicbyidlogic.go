package logic

import (
	"context"
	"fmt"

	dic "go-zero-demo/mall/sys/rpc/dic"
	"go-zero-demo/mall/sys/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDicByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDicByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDicByIdLogic {
	return &GetDicByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDicByIdLogic) GetDicById(in *dic.IdRequest) (*dic.DicResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("从数据库中查询数据。。。。。。")

	return &dic.DicResponse{
		Id:  2,
		Key: "1",
		Val: "男",
	}, nil
}
