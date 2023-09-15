package controller

import (
	"xj/xapi-backend/enums"
	gerror "xj/xapi-backend/g_error"
	ghandle "xj/xapi-backend/g_handle"
	glog "xj/xapi-backend/g_log"
	"xj/xapi-backend/middleware"

	"github.com/gin-gonic/gin"
)

type RateLimitConfig struct {
	IP                string  `json:"ip"`
	RequestsPerSecond float64 `json:"requests_per_second"`
	BucketSize        int     `json:"bucket_size"`
}

//	@Summary		获得具体IP的限流配置
//	@Description	获得具体IP的限流配置
//	@Tags			管理配置
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			ip	query		string	true	"ip地址"
//	@Success		200	{object}	object	"具体IP的限流配置"
//	@Router			/manage/config/ratelimit [get]
func GetIPRateLimitConfig(c *gin.Context) {
	ip := c.Query("ip")
	if ip == "" {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	requestsPerSecond, bucketSize := middleware.GetIPRateLimitConfig(ip)

	resData := RateLimitConfig{
		IP:                ip,
		RequestsPerSecond: requestsPerSecond,
		BucketSize:        bucketSize,
	}

	ghandle.HandlerSuccess(c, "IP的限流配置获取成功", resData)
}

//	@Summary		更新具体IP限流配置失败
//	@Description	更新具体IP限流配置失败
//	@Tags			管理配置
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			request	body		RateLimitConfig	true	"限流配置"
//	@Success		200		{object}	object
//	@Router			/manage/config/ratelimit [put]
func UpdateIPRateLimitConfig(c *gin.Context) {
	var params *RateLimitConfig
	if err := c.ShouldBindJSON(&params); err != nil {
		glog.Log.Errorf("RateLimitConfig err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	if params.RequestsPerSecond <= 0 || params.BucketSize <= 0 {
		glog.Log.Errorf("参数IP限流配置无效，RequestsPerSecond=%v, BucketSize=%v", params.RequestsPerSecond, params.BucketSize)
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "限流配置无效"))
		return
	}
	if err := middleware.UpdateIPRateLimitConfig(params.IP, params.RequestsPerSecond, params.BucketSize); err != nil {
		c.Error(gerror.NewAbortErr(int(enums.UpdateIPRateLimitConfigFailed), "更新IP限流配置失败"))
		return
	}

	ghandle.HandlerSuccess(c, "更新IP限流配置成功", nil)
}
