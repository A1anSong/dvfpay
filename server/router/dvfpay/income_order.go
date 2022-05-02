package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IncomeOrderRouter struct {
}

// InitIncomeOrderRouter 初始化 IncomeOrder 路由信息
func (s *IncomeOrderRouter) InitIncomeOrderRouter(Router *gin.RouterGroup) {
	incomeOrderRouter := Router.Group("incomeOrder").Use(middleware.OperationRecord())
	incomeOrderRouterWithoutRecord := Router.Group("incomeOrder")
	var incomeOrderApi = v1.ApiGroupApp.DvfpayApiGroup.IncomeOrderApi
	{
		incomeOrderRouter.POST("createIncomeOrder", incomeOrderApi.CreateIncomeOrder)   // 新建IncomeOrder
		incomeOrderRouter.DELETE("deleteIncomeOrder", incomeOrderApi.DeleteIncomeOrder) // 删除IncomeOrder
		incomeOrderRouter.DELETE("deleteIncomeOrderByIds", incomeOrderApi.DeleteIncomeOrderByIds) // 批量删除IncomeOrder
		incomeOrderRouter.PUT("updateIncomeOrder", incomeOrderApi.UpdateIncomeOrder)    // 更新IncomeOrder
	}
	{
		incomeOrderRouterWithoutRecord.GET("findIncomeOrder", incomeOrderApi.FindIncomeOrder)        // 根据ID获取IncomeOrder
		incomeOrderRouterWithoutRecord.GET("getIncomeOrderList", incomeOrderApi.GetIncomeOrderList)  // 获取IncomeOrder列表
	}
}
