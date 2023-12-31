package ghandle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 状态未授权 401
func HandlerUnauthorized(c *gin.Context) {
	//  "Error 1044 (42000): Access denied for user 'xapiuser'@'%' to database 'xapi'"
	// glog.Log.Infof("Error : Access denied for user '%s'@'%s' to 业务 '%s'", "用户名", "用户补充", "业务名")
	c.JSON(http.StatusOK, gin.H{"result": http.StatusUnauthorized, "msg": "状态未授权"})
	c.Abort()
}

// 禁止状态 403
func HandlerForbidden(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": http.StatusForbidden, "msg": "禁止状态"})
	c.Abort()
}

// 从上下文信息获取信息失败 204
func HandlerContextError(c *gin.Context, key string) {
	c.JSON(http.StatusOK, gin.H{"result": http.StatusNoContent, "msg": "从上下文信息获取[" + key + "]失败"})
	c.Abort()
}

// 未知错误 500
func HandlerServerError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": http.StatusInternalServerError, "msg": "未知错误"})
	c.Abort()
}

// Dubbo加载失败 424
func HandlerDobboLoadFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"result": http.StatusFailedDependency, "msg": "Dubbo加载失败: " + msg})
	c.Abort()
}

// ——————————————————————以下是业务——————————————————————

// 业务成功 200
func HandlerSuccess(c *gin.Context, msg string, data any) {
	if data == nil {
		c.JSON(http.StatusOK, gin.H{"result": 0, "msg": msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": 0, "msg": msg, "data": data})
}

// 参数错误 200
func HandlerParamError(c *gin.Context, paramName string) {
	if paramName == "" {
		c.JSON(http.StatusOK, gin.H{"result": 1, "msg": "参数错误"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": 1, "msg": "参数[" + paramName + "]错误"})
	}
	c.Abort()
}
