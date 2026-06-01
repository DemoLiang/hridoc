package middleware

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/DemoLiang/hridoc/api/internal/config"
	"github.com/DemoLiang/hridoc/api/pkg/auth"
)

func JWTAuth(cfg config.Config) func(http.HandlerFunc) http.HandlerFunc {
	jwtUtil := auth.NewJWT(cfg.JwtAuth.AccessSecret, cfg.JwtAuth.AccessExpire)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if len(token) < 7 || token[:7] != "Bearer " {
				httpx.WriteJson(w, http.StatusUnauthorized, map[string]any{
					"code":    401,
					"message": "missing or invalid authorization header",
				})
				return
			}

			claims, err := jwtUtil.ParseToken(token[7:])
			if err != nil {
				logx.Errorf("JWT parse failed: %v", err)
				httpx.WriteJson(w, http.StatusUnauthorized, map[string]any{
					"code":    401,
					"message": "invalid token",
				})
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, "userId", claims.UserId)
			ctx = context.WithValue(ctx, "username", claims.Username)
			ctx = context.WithValue(ctx, "role", claims.Role)
			next(w, r.WithContext(ctx))
		}
	}
}

func GetUserId(ctx context.Context) int64 {
	if v, ok := ctx.Value("userId").(int64); ok {
		return v
	}
	return 0
}

func GetUsername(ctx context.Context) string {
	if v, ok := ctx.Value("username").(string); ok {
		return v
	}
	return ""
}

func GetRole(ctx context.Context) int {
	if v, ok := ctx.Value("role").(int); ok {
		return v
	}
	return 0
}
