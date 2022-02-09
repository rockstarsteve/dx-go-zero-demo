package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/myjwt/internal/logic"
	"go-zero-demo/myjwt/internal/svc"
)

func JwtHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewJwtLogic(r.Context(), svcCtx)
		err := l.Jwt()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
