package customer

import (
	"PowerX/internal/logic/mp/customer/auth"
	"net/http"

	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthByProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewAuthByProfileLogic(r.Context(), svcCtx)
		resp, err := l.AuthByProfile()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
