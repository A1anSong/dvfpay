package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IncomeGatewayAuthRouter struct {
}

// InitIncomeGatewayAuthRouter 初始化 IncomeGatewayAuth 路由信息
func (s *IncomeGatewayAuthRouter) InitIncomeGatewayAuthRouter(Router *gin.RouterGroup) {
	incomeGatewayAuthRouter := Router.Group("incomeGatewayAuth").Use(middleware.OperationRecord())
	incomeGatewayAuthRouterWithoutRecord := Router.Group("incomeGatewayAuth")
	var incomeGatewayAuthApi = v1.ApiGroupApp.DvfpayApiGroup.IncomeGatewayAuthApi
	{
		incomeGatewayAuthRouter.POST("createIncomeGatewayAuth", incomeGatewayAuthApi.CreateIncomeGatewayAuth)             // 新建IncomeGatewayAuth
		incomeGatewayAuthRouter.DELETE("deleteIncomeGatewayAuth", incomeGatewayAuthApi.DeleteIncomeGatewayAuth)           // 删除IncomeGatewayAuth
		incomeGatewayAuthRouter.DELETE("deleteIncomeGatewayAuthByIds", incomeGatewayAuthApi.DeleteIncomeGatewayAuthByIds) // 批量删除IncomeGatewayAuth
		incomeGatewayAuthRouter.PUT("updateIncomeGatewayAuth", incomeGatewayAuthApi.UpdateIncomeGatewayAuth)              // 更新IncomeGatewayAuth
	}
	{
		incomeGatewayAuthRouterWithoutRecord.GET("findIncomeGatewayAuth", incomeGatewayAuthApi.FindIncomeGatewayAuth)       // 根据ID获取IncomeGatewayAuth
		incomeGatewayAuthRouterWithoutRecord.GET("getIncomeGatewayAuthList", incomeGatewayAuthApi.GetIncomeGatewayAuthList) // 获取IncomeGatewayAuth列表
	}
}
