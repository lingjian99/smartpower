package handler

import (
	"net/http"
	"smartpower/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"

	"smartpower/diagram/internal/logic"
	"smartpower/diagram/internal/svc"
	"smartpower/diagram/internal/types"
)

func StructDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			response.WithParameterError(w, err)
			return
		}

		l := logic.NewStructDataLogic(r.Context(), svcCtx)
		resp, err := l.StructData(&req)
		response.Response(w, resp, err)
	}
}
