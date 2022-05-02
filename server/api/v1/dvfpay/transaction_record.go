package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TransactionRecordApi struct {
}

var transactionRecordService = service.ServiceGroupApp.DvfpayServiceGroup.TransactionRecordService


// CreateTransactionRecord 创建TransactionRecord
// @Tags TransactionRecord
// @Summary 创建TransactionRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.TransactionRecord true "创建TransactionRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /transactionRecord/createTransactionRecord [post]
func (transactionRecordApi *TransactionRecordApi) CreateTransactionRecord(c *gin.Context) {
	var transactionRecord dvfpay.TransactionRecord
	_ = c.ShouldBindJSON(&transactionRecord)
	if err := transactionRecordService.CreateTransactionRecord(transactionRecord); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTransactionRecord 删除TransactionRecord
// @Tags TransactionRecord
// @Summary 删除TransactionRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.TransactionRecord true "删除TransactionRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /transactionRecord/deleteTransactionRecord [delete]
func (transactionRecordApi *TransactionRecordApi) DeleteTransactionRecord(c *gin.Context) {
	var transactionRecord dvfpay.TransactionRecord
	_ = c.ShouldBindJSON(&transactionRecord)
	if err := transactionRecordService.DeleteTransactionRecord(transactionRecord); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTransactionRecordByIds 批量删除TransactionRecord
// @Tags TransactionRecord
// @Summary 批量删除TransactionRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TransactionRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /transactionRecord/deleteTransactionRecordByIds [delete]
func (transactionRecordApi *TransactionRecordApi) DeleteTransactionRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := transactionRecordService.DeleteTransactionRecordByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTransactionRecord 更新TransactionRecord
// @Tags TransactionRecord
// @Summary 更新TransactionRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.TransactionRecord true "更新TransactionRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /transactionRecord/updateTransactionRecord [put]
func (transactionRecordApi *TransactionRecordApi) UpdateTransactionRecord(c *gin.Context) {
	var transactionRecord dvfpay.TransactionRecord
	_ = c.ShouldBindJSON(&transactionRecord)
	if err := transactionRecordService.UpdateTransactionRecord(transactionRecord); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTransactionRecord 用id查询TransactionRecord
// @Tags TransactionRecord
// @Summary 用id查询TransactionRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.TransactionRecord true "用id查询TransactionRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /transactionRecord/findTransactionRecord [get]
func (transactionRecordApi *TransactionRecordApi) FindTransactionRecord(c *gin.Context) {
	var transactionRecord dvfpay.TransactionRecord
	_ = c.ShouldBindQuery(&transactionRecord)
	if err, retransactionRecord := transactionRecordService.GetTransactionRecord(transactionRecord.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retransactionRecord": retransactionRecord}, c)
	}
}

// GetTransactionRecordList 分页获取TransactionRecord列表
// @Tags TransactionRecord
// @Summary 分页获取TransactionRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.TransactionRecordSearch true "分页获取TransactionRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /transactionRecord/getTransactionRecordList [get]
func (transactionRecordApi *TransactionRecordApi) GetTransactionRecordList(c *gin.Context) {
	var pageInfo dvfpayReq.TransactionRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := transactionRecordService.GetTransactionRecordInfoList(pageInfo); err != nil {
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
