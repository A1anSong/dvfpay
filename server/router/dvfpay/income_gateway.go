package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IncomeGatewayRouter struct {
}

// InitIncomeGatewayRouter 初始化 IncomeGateway 路由信息
func (s *IncomeGatewayRouter) InitIncomeGatewayRouter(Router *gin.RouterGroup) {
	incomeGatewayRouter := Router.Group("incomeGateway").Use(middleware.OperationRecord())
	incomeGatewayRouterWithoutRecord := Router.Group("incomeGateway")
	var incomeGatewayApi = v1.ApiGroupApp.DvfpayApiGroup.IncomeGatewayApi
	{
		incomeGatewayRouter.POST("createIncomeGateway", incomeGatewayApi.CreateIncomeGateway)   // 新建IncomeGateway
		incomeGatewayRouter.DELETE("deleteIncomeGateway", incomeGatewayApi.DeleteIncomeGateway) // 删除IncomeGateway
		incomeGatewayRouter.DELETE("deleteIncomeGatewayByIds", incomeGatewayApi.DeleteIncomeGatewayByIds) // 批量删除IncomeGateway
		incomeGatewayRouter.PUT("updateIncomeGateway", incomeGatewayApi.UpdateIncomeGateway)    // 更新IncomeGateway
	}
	{
		incomeGatewayRouterWithoutRecord.GET("findIncomeGateway", incomeGatewayApi.FindIncomeGateway)        // 根据ID获取IncomeGateway
		incomeGatewayRouterWithoutRecord.GET("getIncomeGatewayList", incomeGatewayApi.GetIncomeGatewayList)  // 获取IncomeGateway列表
	}
}
