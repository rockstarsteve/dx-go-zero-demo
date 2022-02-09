package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/jwttoken/internal/logic"
	"go-zero-demo/jwttoken/internal/svc"
	"go-zero-demo/jwttoken/internal/types"
)

func JwtHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JwtTokenRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewJwtLogic(r.Context(), svcCtx)
		resp, err := l.Jwt(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
