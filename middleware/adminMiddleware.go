package middleware

import (
	"xj/xapi-backend/enums"
	gerror "xj/xapi-backend/g_error"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户信息
		userrole, exists := c.Get("user_role")
		if !exists || userrole.(string) != "admin" {
			c.Error(gerror.NewAbortErr(int(enums.NotAdminRole), "无权限"))
			c.Abort()
			return
		}
		c.Next()
	}
}
