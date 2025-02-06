package user

import (
	"net/http"

	"github.com/micjn89757/VulFixMiner/server/basicService/internal/logic/user"
	"github.com/micjn89757/VulFixMiner/server/basicService/internal/svc"
	"github.com/micjn89757/VulFixMiner/server/basicService/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateUserInfoLogic(r.Context(), svcCtx)
		err := l.UpdateUserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
