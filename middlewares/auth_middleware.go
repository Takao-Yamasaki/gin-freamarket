package middlewares

import (
	"gin-fleamarket/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 認証ミドルウェア
func AuthMiddleware(authService services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ヘッダーの検証
		header := ctx.GetHeader("Authorization")
		if header == "" {
			// 後続のハンドラー関数の処理を中止
			ctx.AbortWithStatus(http.StatusUnauthorized)
			// 処理を終了
			return
		}

		// JWTトークンはBearerで始まる
		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStrings := strings.TrimPrefix(header, "Bearer ")
		user, err := authService.GetUserFromToken(tokenStrings)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// リクエストの生存期間中に保持されるキーと値のペア
		// レスポンスが返されるまでの間、取得することができる
		ctx.Set("user", user)

		// 処理を次のミドルウェアまたは目的のメソッドに移行
		ctx.Next()
	}
}
