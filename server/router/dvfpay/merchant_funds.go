package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MerchantFundsRouter struct {
}

// InitMerchantFundsRouter 初始化 MerchantFunds 路由信息
func (s *MerchantFundsRouter) InitMerchantFundsRouter(Router *gin.RouterGroup) {
	merchantFundsRouter := Router.Group("merchantFunds").Use(middleware.OperationRecord())
	merchantFundsRouterWithoutRecord := Router.Group("merchantFunds")
	var merchantFundsApi = v1.ApiGroupApp.DvfpayApiGroup.MerchantFundsApi
	{
		merchantFundsRouter.POST("createMerchantFunds", merchantFundsApi.CreateMerchantFunds)   // 新建MerchantFunds
		merchantFundsRouter.DELETE("deleteMerchantFunds", merchantFundsApi.DeleteMerchantFunds) // 删除MerchantFunds
		merchantFundsRouter.DELETE("deleteMerchantFundsByIds", merchantFundsApi.DeleteMerchantFundsByIds) // 批量删除MerchantFunds
		merchantFundsRouter.PUT("updateMerchantFunds", merchantFundsApi.UpdateMerchantFunds)    // 更新MerchantFunds
	}
	{
		merchantFundsRouterWithoutRecord.GET("findMerchantFunds", merchantFundsApi.FindMerchantFunds)        // 根据ID获取MerchantFunds
		merchantFundsRouterWithoutRecord.GET("getMerchantFundsList", merchantFundsApi.GetMerchantFundsList)  // 获取MerchantFunds列表
	}
}
