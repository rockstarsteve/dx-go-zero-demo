package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	errorx "go-zero-demo/jwttoken/internal/common"
	"time"

	"go-zero-demo/jwttoken/internal/svc"
	"go-zero-demo/jwttoken/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *JwtLogic) Jwt(req types.JwtTokenRequest) (*types.JwtTokenResponse, error) {

	if req.UserId != "007" {
		return &types.JwtTokenResponse{}, errorx.NewCodeError(int(500),"用户名不存在")
	}

	var accessExpire = l.svcCtx.Config.JwtAuth.AccessExpire

	now := time.Now().Unix()
	accessToken, err := l.GenToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, nil, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.JwtTokenResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *JwtLogic) GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = "007"
	//for k, v := range payloads {
	//	claims[k] = v
	//}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
