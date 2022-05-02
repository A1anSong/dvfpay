package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TransactionRecordRouter struct {
}

// InitTransactionRecordRouter 初始化 TransactionRecord 路由信息
func (s *TransactionRecordRouter) InitTransactionRecordRouter(Router *gin.RouterGroup) {
	transactionRecordRouter := Router.Group("transactionRecord").Use(middleware.OperationRecord())
	transactionRecordRouterWithoutRecord := Router.Group("transactionRecord")
	var transactionRecordApi = v1.ApiGroupApp.DvfpayApiGroup.TransactionRecordApi
	{
		transactionRecordRouter.POST("createTransactionRecord", transactionRecordApi.CreateTransactionRecord)   // 新建TransactionRecord
		transactionRecordRouter.DELETE("deleteTransactionRecord", transactionRecordApi.DeleteTransactionRecord) // 删除TransactionRecord
		transactionRecordRouter.DELETE("deleteTransactionRecordByIds", transactionRecordApi.DeleteTransactionRecordByIds) // 批量删除TransactionRecord
		transactionRecordRouter.PUT("updateTransactionRecord", transactionRecordApi.UpdateTransactionRecord)    // 更新TransactionRecord
	}
	{
		transactionRecordRouterWithoutRecord.GET("findTransactionRecord", transactionRecordApi.FindTransactionRecord)        // 根据ID获取TransactionRecord
		transactionRecordRouterWithoutRecord.GET("getTransactionRecordList", transactionRecordApi.GetTransactionRecordList)  // 获取TransactionRecord列表
	}
}
