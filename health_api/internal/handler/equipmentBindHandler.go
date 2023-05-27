package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wk-middleground/health_api/internal/logic"
	"wk-middleground/health_api/internal/svc"
	"wk-middleground/health_api/internal/types"
)

func EquipmentBindHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EquipmentBindRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewEquipmentBindLogic(r.Context(), svcCtx)
		resp, err := l.EquipmentBind(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
