package order

import (
	"context"
	"fmt"
	"go-zero-demo/mall/user/rpc/userclient"

	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) SaveOrderLogic {
	return SaveOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrderLogic) SaveOrder(req types.SaveReq) (resp *types.SaveResp, err error) {
	// todo: add your logic here and delete this line

	saveResp, err := l.svcCtx.UserRpc.SaveUser(l.ctx, &userclient.SaveReq{})

	if err != nil {
		fmt.Println("保存用户有误！！！")
	}

	return &types.SaveResp{
		Code:     saveResp.Code,
		Msg: saveResp.Msg,
	}, err
}
