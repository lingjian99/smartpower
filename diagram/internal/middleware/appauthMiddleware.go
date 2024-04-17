package middleware

import (
	"fmt"
	"net/http"
	"smartpower/pkg/def/defuser"
	"smartpower/pkg/utils"
	"smartpower/pkg/xerr"

	"smartpower/pkg/response"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

var (
	ErrNoTokenInRequest = errors.New("no token present in request")
)

type AppAuthMiddleware struct {
	cc cache.Cache
}
func NewAppAuthMiddleware(cc cache.Cache) *AppAuthMiddleware {
	return &AppAuthMiddleware{
		cc: cc,
	}
}

func (m *AppAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := ExtractToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		userId, _ := utils.CtxInt64(r.Context(), "userId")
		var cToken string
		err = m.cc.GetCtx(r.Context(), fmt.Sprintf("%s%d", defuser.CliUserTokenPrefix, userId), &cToken)

		if err != nil && !m.cc.IsNotFound(err) {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			response.Response(w, nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "err: %+v", err))
			return
		}
		//封禁,会将redis中的token删除
		if m.cc.IsNotFound(err) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if cToken != tokenString {
			http.Error(w, "您的帐号已在另一台设备登录", http.StatusUnauthorized)
			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}

func ExtractToken(req *http.Request)(string, error) {
	tokenHeader := req.Header.Get("Authorization")
	if tokenHeader == "" {
		return "", ErrNoTokenInRequest
	}
	return tokenHeader, nil}
