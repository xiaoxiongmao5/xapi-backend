package ghandle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 禁止状态
func HandlerNoAuth(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "禁止状态"})
	c.Abort()
}

// 状态未授权
func HandlerUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "状态未授权"})
	c.Abort()
}

// 从上下文信息获取信息失败
func HandlerContextError(c *gin.Context, key string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "从上下文信息获取[" + key + "]失败"})
	c.Abort()
}

// 业务成功
func HandlerSuccess(c *gin.Context, msg string, data any) {
	if data == nil {
		c.JSON(http.StatusOK, gin.H{"result": 0, "msg": msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": 0, "msg": msg, "data": data})
}
