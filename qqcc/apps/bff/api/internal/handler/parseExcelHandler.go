package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qqcc/apps/bff/api/internal/logic"
	"qqcc/apps/bff/api/internal/svc"
	"qqcc/apps/bff/api/internal/types"
)

func ParseExcelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ParseExcelRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewParseExcelLogic(r.Context(), svcCtx)
		resp, err := l.ParseExcel(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
