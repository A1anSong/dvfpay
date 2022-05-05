package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayoutGatewayAuthRouter struct {
}

// InitPayoutGatewayAuthRouter 初始化 PayoutGatewayAuth 路由信息
func (s *PayoutGatewayAuthRouter) InitPayoutGatewayAuthRouter(Router *gin.RouterGroup) {
	payoutGatewayAuthRouter := Router.Group("payoutGatewayAuth").Use(middleware.OperationRecord())
	payoutGatewayAuthRouterWithoutRecord := Router.Group("payoutGatewayAuth")
	var payoutGatewayAuthApi = v1.ApiGroupApp.DvfpayApiGroup.PayoutGatewayAuthApi
	{
		payoutGatewayAuthRouter.POST("createPayoutGatewayAuth", payoutGatewayAuthApi.CreatePayoutGatewayAuth)             // 新建PayoutGatewayAuth
		payoutGatewayAuthRouter.DELETE("deletePayoutGatewayAuth", payoutGatewayAuthApi.DeletePayoutGatewayAuth)           // 删除PayoutGatewayAuth
		payoutGatewayAuthRouter.DELETE("deletePayoutGatewayAuthByIds", payoutGatewayAuthApi.DeletePayoutGatewayAuthByIds) // 批量删除PayoutGatewayAuth
		payoutGatewayAuthRouter.PUT("updatePayoutGatewayAuth", payoutGatewayAuthApi.UpdatePayoutGatewayAuth)              // 更新PayoutGatewayAuth
	}
	{
		payoutGatewayAuthRouterWithoutRecord.GET("findPayoutGatewayAuth", payoutGatewayAuthApi.FindPayoutGatewayAuth)       // 根据ID获取PayoutGatewayAuth
		payoutGatewayAuthRouterWithoutRecord.GET("getPayoutGatewayAuthList", payoutGatewayAuthApi.GetPayoutGatewayAuthList) // 获取PayoutGatewayAuth列表
	}
}
