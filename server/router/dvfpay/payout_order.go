package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayoutOrderRouter struct {
}

// InitPayoutOrderRouter 初始化 PayoutOrder 路由信息
func (s *PayoutOrderRouter) InitPayoutOrderRouter(Router *gin.RouterGroup) {
	payoutOrderRouter := Router.Group("payoutOrder").Use(middleware.OperationRecord())
	payoutOrderRouterWithoutRecord := Router.Group("payoutOrder")
	var payoutOrderApi = v1.ApiGroupApp.DvfpayApiGroup.PayoutOrderApi
	{
		payoutOrderRouter.POST("createPayoutOrder", payoutOrderApi.CreatePayoutOrder)   // 新建PayoutOrder
		payoutOrderRouter.DELETE("deletePayoutOrder", payoutOrderApi.DeletePayoutOrder) // 删除PayoutOrder
		payoutOrderRouter.DELETE("deletePayoutOrderByIds", payoutOrderApi.DeletePayoutOrderByIds) // 批量删除PayoutOrder
		payoutOrderRouter.PUT("updatePayoutOrder", payoutOrderApi.UpdatePayoutOrder)    // 更新PayoutOrder
	}
	{
		payoutOrderRouterWithoutRecord.GET("findPayoutOrder", payoutOrderApi.FindPayoutOrder)        // 根据ID获取PayoutOrder
		payoutOrderRouterWithoutRecord.GET("getPayoutOrderList", payoutOrderApi.GetPayoutOrderList)  // 获取PayoutOrder列表
	}
}
