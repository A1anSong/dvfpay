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

type PayoutOrderApi struct {
}

var payoutOrderService = service.ServiceGroupApp.DvfpayServiceGroup.PayoutOrderService

// CreatePayoutOrder 创建PayoutOrder
// @Tags PayoutOrder
// @Summary 创建PayoutOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.PayoutOrder true "创建PayoutOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutOrder/createPayoutOrder [post]
func (payoutOrderApi *PayoutOrderApi) CreatePayoutOrder(c *gin.Context) {
	var payoutOrder dvfpay.PayoutOrder
	_ = c.ShouldBindJSON(&payoutOrder)
	if err := payoutOrderService.CreatePayoutOrder(payoutOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePayoutOrder 删除PayoutOrder
// @Tags PayoutOrder
// @Summary 删除PayoutOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.PayoutOrder true "删除PayoutOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payoutOrder/deletePayoutOrder [delete]
func (payoutOrderApi *PayoutOrderApi) DeletePayoutOrder(c *gin.Context) {
	var payoutOrder dvfpay.PayoutOrder
	_ = c.ShouldBindJSON(&payoutOrder)
	if err := payoutOrderService.DeletePayoutOrder(payoutOrder); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayoutOrderByIds 批量删除PayoutOrder
// @Tags PayoutOrder
// @Summary 批量删除PayoutOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayoutOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payoutOrder/deletePayoutOrderByIds [delete]
func (payoutOrderApi *PayoutOrderApi) DeletePayoutOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := payoutOrderService.DeletePayoutOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayoutOrder 更新PayoutOrder
// @Tags PayoutOrder
// @Summary 更新PayoutOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.PayoutOrder true "更新PayoutOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payoutOrder/updatePayoutOrder [put]
func (payoutOrderApi *PayoutOrderApi) UpdatePayoutOrder(c *gin.Context) {
	var payoutOrder dvfpay.PayoutOrder
	_ = c.ShouldBindJSON(&payoutOrder)
	if err := payoutOrderService.UpdatePayoutOrder(payoutOrder); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPayoutOrder 用id查询PayoutOrder
// @Tags PayoutOrder
// @Summary 用id查询PayoutOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.PayoutOrder true "用id查询PayoutOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payoutOrder/findPayoutOrder [get]
func (payoutOrderApi *PayoutOrderApi) FindPayoutOrder(c *gin.Context) {
	var payoutOrder dvfpay.PayoutOrder
	_ = c.ShouldBindQuery(&payoutOrder)
	if err, repayoutOrder := payoutOrderService.GetPayoutOrder(payoutOrder.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repayoutOrder": repayoutOrder}, c)
	}
}

// GetPayoutOrderList 分页获取PayoutOrder列表
// @Tags PayoutOrder
// @Summary 分页获取PayoutOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.PayoutOrderSearch true "分页获取PayoutOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutOrder/getPayoutOrderList [get]
func (payoutOrderApi *PayoutOrderApi) GetPayoutOrderList(c *gin.Context) {
	var pageInfo dvfpayReq.PayoutOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := payoutOrderService.GetPayoutOrderInfoList(pageInfo); err != nil {
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

func (payoutOrderApi *PayoutOrderApi) GetMerchantPayoutOrderList(c *gin.Context) {
	var pageInfo dvfpayReq.PayoutOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	merchantID := utils.GetUserID(c)
	if err, list, total := payoutOrderService.GetMerchantPayoutOrderInfoList(pageInfo, merchantID); err != nil {
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
