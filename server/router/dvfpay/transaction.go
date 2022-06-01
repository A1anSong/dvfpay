package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TransactionRouter struct {
}

// InitTransactionRouter 初始化 Transaction 路由信息
func (s *TransactionRouter) InitTransactionRouter(Router *gin.RouterGroup) {
	transactionRouter := Router.Group("transaction").Use(middleware.OperationRecord())
	transactionRouterWithoutRecord := Router.Group("transaction")
	var transactionApi = v1.ApiGroupApp.DvfpayApiGroup.TransactionApi
	{
		transactionRouter.POST("createTransaction", transactionApi.CreateTransaction)             // 新建Transaction
		transactionRouter.DELETE("deleteTransaction", transactionApi.DeleteTransaction)           // 删除Transaction
		transactionRouter.DELETE("deleteTransactionByIds", transactionApi.DeleteTransactionByIds) // 批量删除Transaction
		transactionRouter.PUT("updateTransaction", transactionApi.UpdateTransaction)              // 更新Transaction
	}
	{
		transactionRouterWithoutRecord.GET("findTransaction", transactionApi.FindTransaction)       // 根据ID获取Transaction
		transactionRouterWithoutRecord.GET("getTransactionList", transactionApi.GetTransactionList) // 获取Transaction列表
		transactionRouterWithoutRecord.GET("getMerchantTransactionList", transactionApi.GetMerchantTransactionList)
	}
}
