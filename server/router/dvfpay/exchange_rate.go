package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExchangeRateRouter struct {
}

// InitExchangeRateRouter 初始化 ExchangeRate 路由信息
func (s *ExchangeRateRouter) InitExchangeRateRouter(Router *gin.RouterGroup) {
	exchangeRateRouter := Router.Group("exchangeRate").Use(middleware.OperationRecord())
	exchangeRateRouterWithoutRecord := Router.Group("exchangeRate")
	var exchangeRateApi = v1.ApiGroupApp.DvfpayApiGroup.ExchangeRateApi
	{
		exchangeRateRouter.POST("createExchangeRate", exchangeRateApi.CreateExchangeRate)   // 新建ExchangeRate
		exchangeRateRouter.DELETE("deleteExchangeRate", exchangeRateApi.DeleteExchangeRate) // 删除ExchangeRate
		exchangeRateRouter.DELETE("deleteExchangeRateByIds", exchangeRateApi.DeleteExchangeRateByIds) // 批量删除ExchangeRate
		exchangeRateRouter.PUT("updateExchangeRate", exchangeRateApi.UpdateExchangeRate)    // 更新ExchangeRate
	}
	{
		exchangeRateRouterWithoutRecord.GET("findExchangeRate", exchangeRateApi.FindExchangeRate)        // 根据ID获取ExchangeRate
		exchangeRateRouterWithoutRecord.GET("getExchangeRateList", exchangeRateApi.GetExchangeRateList)  // 获取ExchangeRate列表
	}
}
