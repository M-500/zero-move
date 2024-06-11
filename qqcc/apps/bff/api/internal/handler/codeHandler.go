package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qqcc/apps/bff/api/internal/logic"
	"qqcc/apps/bff/api/internal/svc"
)

func CodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCodeLogic(r.Context(), svcCtx)
		resp, err := l.Code()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
