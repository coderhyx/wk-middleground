package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wk-middleground/health_api/internal/logic"
	"wk-middleground/health_api/internal/svc"
	"wk-middleground/health_api/internal/types"
)

func ChatGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatGetRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChatGetLogic(r.Context(), svcCtx)
		resp, err := l.ChatGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
