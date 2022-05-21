package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExchangeRecordRouter struct {
}

// InitExchangeRecordRouter 初始化 ExchangeRecord 路由信息
func (s *ExchangeRecordRouter) InitExchangeRecordRouter(Router *gin.RouterGroup) {
	exchangeRecordRouter := Router.Group("exchangeRecord").Use(middleware.OperationRecord())
	exchangeRecordRouterWithoutRecord := Router.Group("exchangeRecord")
	var exchangeRecordApi = v1.ApiGroupApp.DvfpayApiGroup.ExchangeRecordApi
	{
		exchangeRecordRouter.POST("createExchangeRecord", exchangeRecordApi.CreateExchangeRecord)             // 新建ExchangeRecord
		exchangeRecordRouter.DELETE("deleteExchangeRecord", exchangeRecordApi.DeleteExchangeRecord)           // 删除ExchangeRecord
		exchangeRecordRouter.DELETE("deleteExchangeRecordByIds", exchangeRecordApi.DeleteExchangeRecordByIds) // 批量删除ExchangeRecord
		exchangeRecordRouter.PUT("updateExchangeRecord", exchangeRecordApi.UpdateExchangeRecord)              // 更新ExchangeRecord
	}
	{
		exchangeRecordRouterWithoutRecord.GET("findExchangeRecord", exchangeRecordApi.FindExchangeRecord)       // 根据ID获取ExchangeRecord
		exchangeRecordRouterWithoutRecord.GET("getExchangeRecordList", exchangeRecordApi.GetExchangeRecordList) // 获取ExchangeRecord列表
		exchangeRecordRouterWithoutRecord.GET("getMerchantExchangeRecordList", exchangeRecordApi.GetMerchantExchangeRecordList)
	}
}
