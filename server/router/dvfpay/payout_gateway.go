package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayoutGatewayRouter struct {
}

// InitPayoutGatewayRouter 初始化 PayoutGateway 路由信息
func (s *PayoutGatewayRouter) InitPayoutGatewayRouter(Router *gin.RouterGroup) {
	payoutGatewayRouter := Router.Group("payoutGateway").Use(middleware.OperationRecord())
	payoutGatewayRouterWithoutRecord := Router.Group("payoutGateway")
	var payoutGatewayApi = v1.ApiGroupApp.DvfpayApiGroup.PayoutGatewayApi
	{
		payoutGatewayRouter.POST("createPayoutGateway", payoutGatewayApi.CreatePayoutGateway)   // 新建PayoutGateway
		payoutGatewayRouter.DELETE("deletePayoutGateway", payoutGatewayApi.DeletePayoutGateway) // 删除PayoutGateway
		payoutGatewayRouter.DELETE("deletePayoutGatewayByIds", payoutGatewayApi.DeletePayoutGatewayByIds) // 批量删除PayoutGateway
		payoutGatewayRouter.PUT("updatePayoutGateway", payoutGatewayApi.UpdatePayoutGateway)    // 更新PayoutGateway
	}
	{
		payoutGatewayRouterWithoutRecord.GET("findPayoutGateway", payoutGatewayApi.FindPayoutGateway)        // 根据ID获取PayoutGateway
		payoutGatewayRouterWithoutRecord.GET("getPayoutGatewayList", payoutGatewayApi.GetPayoutGatewayList)  // 获取PayoutGateway列表
	}
}
