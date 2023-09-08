package middleware

import (
	"net/http"
	gerror "xj/xapi-backend/g_error"

	"github.com/gin-gonic/gin"
)

// 捕获中断业务异常 中间件
func G_ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 判断上层业务抛出的错误类型
		if err := c.Errors.Last(); err != nil {
			if abortError, ok := err.Err.(*gerror.AbortError); ok {
				// 生成错误响应并终止请求处理
				c.JSON(http.StatusOK, gin.H{"result": abortError.Code, "msg": abortError.Message})
				c.Abort()
				return
			}
		}
	}
}
