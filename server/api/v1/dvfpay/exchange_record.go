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

type ExchangeRecordApi struct {
}

var exchangeRecordService = service.ServiceGroupApp.DvfpayServiceGroup.ExchangeRecordService

// CreateExchangeRecord 创建ExchangeRecord
// @Tags ExchangeRecord
// @Summary 创建ExchangeRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.ExchangeRecord true "创建ExchangeRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeRecord/createExchangeRecord [post]
func (exchangeRecordApi *ExchangeRecordApi) CreateExchangeRecord(c *gin.Context) {
	var exchangeRecord dvfpay.ExchangeRecord
	_ = c.ShouldBindJSON(&exchangeRecord)
	merchantID := utils.GetUserID(c)
	if err := exchangeRecordService.CreateExchangeRecord(exchangeRecord, merchantID); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		if err.Error() == "余额不足，兑换失败！" {
			response.FailWithMessage("余额不足，兑换失败！", c)
		} else {
			response.FailWithMessage("兑换失败", c)
		}
	} else {
		response.OkWithMessage("兑换成功", c)
	}
}

// DeleteExchangeRecord 删除ExchangeRecord
// @Tags ExchangeRecord
// @Summary 删除ExchangeRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.ExchangeRecord true "删除ExchangeRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeRecord/deleteExchangeRecord [delete]
func (exchangeRecordApi *ExchangeRecordApi) DeleteExchangeRecord(c *gin.Context) {
	var exchangeRecord dvfpay.ExchangeRecord
	_ = c.ShouldBindJSON(&exchangeRecord)
	if err := exchangeRecordService.DeleteExchangeRecord(exchangeRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExchangeRecordByIds 批量删除ExchangeRecord
// @Tags ExchangeRecord
// @Summary 批量删除ExchangeRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExchangeRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exchangeRecord/deleteExchangeRecordByIds [delete]
func (exchangeRecordApi *ExchangeRecordApi) DeleteExchangeRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := exchangeRecordService.DeleteExchangeRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExchangeRecord 更新ExchangeRecord
// @Tags ExchangeRecord
// @Summary 更新ExchangeRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.ExchangeRecord true "更新ExchangeRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exchangeRecord/updateExchangeRecord [put]
func (exchangeRecordApi *ExchangeRecordApi) UpdateExchangeRecord(c *gin.Context) {
	var exchangeRecord dvfpay.ExchangeRecord
	_ = c.ShouldBindJSON(&exchangeRecord)
	if err := exchangeRecordService.UpdateExchangeRecord(exchangeRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExchangeRecord 用id查询ExchangeRecord
// @Tags ExchangeRecord
// @Summary 用id查询ExchangeRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.ExchangeRecord true "用id查询ExchangeRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exchangeRecord/findExchangeRecord [get]
func (exchangeRecordApi *ExchangeRecordApi) FindExchangeRecord(c *gin.Context) {
	var exchangeRecord dvfpay.ExchangeRecord
	_ = c.ShouldBindQuery(&exchangeRecord)
	if err, reexchangeRecord := exchangeRecordService.GetExchangeRecord(exchangeRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexchangeRecord": reexchangeRecord}, c)
	}
}

// GetExchangeRecordList 分页获取ExchangeRecord列表
// @Tags ExchangeRecord
// @Summary 分页获取ExchangeRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.ExchangeRecordSearch true "分页获取ExchangeRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeRecord/getExchangeRecordList [get]
func (exchangeRecordApi *ExchangeRecordApi) GetExchangeRecordList(c *gin.Context) {
	var pageInfo dvfpayReq.ExchangeRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := exchangeRecordService.GetExchangeRecordInfoList(pageInfo); err != nil {
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

func (exchangeRecordApi *ExchangeRecordApi) GetMerchantExchangeRecordList(c *gin.Context) {
	var pageInfo dvfpayReq.ExchangeRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	merchantID := utils.GetUserID(c)
	if err, list, total := exchangeRecordService.GetMerchantExchangeRecordInfoList(pageInfo, merchantID); err != nil {
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
