package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TransactionApi struct {
}

var transactionService = service.ServiceGroupApp.DvfpayServiceGroup.TransactionService

// CreateTransaction 创建Transaction
// @Tags Transaction
// @Summary 创建Transaction
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.Transaction true "创建Transaction"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /transaction/createTransaction [post]
func (transactionApi *TransactionApi) CreateTransaction(c *gin.Context) {
	var transaction dvfpay.Transaction
	_ = c.ShouldBindJSON(&transaction)
	merchantID := utils.GetUserID(c)
	if err := transactionService.CreateTransaction(transaction, merchantID); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		if err.Error() == "余额不足，提现失败！" {
			response.FailWithMessage("余额不足，提现失败！", c)
		} else {
			response.FailWithMessage("提现失败", c)
		}
	} else {
		response.OkWithMessage("提交成功", c)
	}
}

// DeleteTransaction 删除Transaction
// @Tags Transaction
// @Summary 删除Transaction
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.Transaction true "删除Transaction"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /transaction/deleteTransaction [delete]
func (transactionApi *TransactionApi) DeleteTransaction(c *gin.Context) {
	var transaction dvfpay.Transaction
	_ = c.ShouldBindJSON(&transaction)
	if err := transactionService.DeleteTransaction(transaction); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTransactionByIds 批量删除Transaction
// @Tags Transaction
// @Summary 批量删除Transaction
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Transaction"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /transaction/deleteTransactionByIds [delete]
func (transactionApi *TransactionApi) DeleteTransactionByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := transactionService.DeleteTransactionByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTransaction 更新Transaction
// @Tags Transaction
// @Summary 更新Transaction
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.Transaction true "更新Transaction"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /transaction/updateTransaction [put]
func (transactionApi *TransactionApi) UpdateTransaction(c *gin.Context) {
	var transaction dvfpay.Transaction
	_ = c.ShouldBindJSON(&transaction)
	if err := transactionService.UpdateTransaction(transaction); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTransaction 用id查询Transaction
// @Tags Transaction
// @Summary 用id查询Transaction
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.Transaction true "用id查询Transaction"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /transaction/findTransaction [get]
func (transactionApi *TransactionApi) FindTransaction(c *gin.Context) {
	var transaction dvfpay.Transaction
	_ = c.ShouldBindQuery(&transaction)
	if err, retransaction := transactionService.GetTransaction(transaction.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retransaction": retransaction}, c)
	}
}

// GetTransactionList 分页获取Transaction列表
// @Tags Transaction
// @Summary 分页获取Transaction列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.TransactionSearch true "分页获取Transaction列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /transaction/getTransactionList [get]
func (transactionApi *TransactionApi) GetTransactionList(c *gin.Context) {
	var pageInfo dvfpayReq.TransactionSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := transactionService.GetTransactionInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
