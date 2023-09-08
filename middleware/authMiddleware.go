package middleware

import (
	"xj/xapi-backend/enums"
	gconfig "xj/xapi-backend/g_config"
	gerror "xj/xapi-backend/g_error"
	ghandle "xj/xapi-backend/g_handle"
	gstore "xj/xapi-backend/g_store"
	"xj/xapi-backend/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 判断已登录状态的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取当前的 Token
		tokenCookie, err := c.Cookie("token")
		if err != nil || tokenCookie == "" {
			ghandle.HandlerUnauthorized(c)
			return
		}

		// 验证当前 Token
		token, err := jwt.Parse(tokenCookie, func(token *jwt.Token) (interface{}, error) {
			return []byte(gconfig.SecretKey), nil
		})
		if err != nil || !token.Valid {
			c.Error(gerror.NewAbortErr(int(enums.Unauthorized), "Unauthorized"))
			c.Abort()
			return
		}

		// 从 Token 中获取用户信息
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Error(gerror.NewAbortErr(int(enums.Unauthorized), "Unauthorized"))
			c.Abort()
			return
		}

		// 重新生成 Token，并更新有效期
		userAccount := claims["user_account"].(string)
		userRole := claims["user_role"].(string)
		newToken, err := utils.GenerateToken(userAccount, userRole)
		if err != nil {
			c.Error(gerror.NewAbortErr(int(enums.GenerateTokenFailed), err.Error()))
			c.Abort()
			return
		}

		// 删除旧的 token
		delete(gstore.TokenMemoryStore, tokenCookie)

		// 更新内存中的 token 数据
		gstore.TokenMemoryStore[newToken] = true

		// 将新的 token 返回给前端
		domain, _ := utils.GetDomainFromReferer(c.Request.Referer())
		c.SetCookie("token", newToken, 3600, "/", domain, false, true)

		// 在此可以将 claims 中的用户信息保存到上下文中，供后续处理使用
		c.Set("user_account", claims["user_account"])
		c.Set("user_role", claims["user_role"])

		c.Next()
	}
}
