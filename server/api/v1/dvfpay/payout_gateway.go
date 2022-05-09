package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PayoutGatewayApi struct {
}

var payoutGatewayService = service.ServiceGroupApp.DvfpayServiceGroup.PayoutGatewayService

// CreatePayoutGateway 创建PayoutGateway
// @Tags PayoutGateway
// @Summary 创建PayoutGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.PayoutGateway true "创建PayoutGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutGateway/createPayoutGateway [post]
func (payoutGatewayApi *PayoutGatewayApi) CreatePayoutGateway(c *gin.Context) {
	var payoutGateway dvfpay.PayoutGateway
	_ = c.ShouldBindJSON(&payoutGateway)
	if err := payoutGatewayService.CreatePayoutGateway(payoutGateway); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePayoutGateway 删除PayoutGateway
// @Tags PayoutGateway
// @Summary 删除PayoutGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.PayoutGateway true "删除PayoutGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payoutGateway/deletePayoutGateway [delete]
func (payoutGatewayApi *PayoutGatewayApi) DeletePayoutGateway(c *gin.Context) {
	var payoutGateway dvfpay.PayoutGateway
	_ = c.ShouldBindJSON(&payoutGateway)
	if err := payoutGatewayService.DeletePayoutGateway(payoutGateway); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayoutGatewayByIds 批量删除PayoutGateway
// @Tags PayoutGateway
// @Summary 批量删除PayoutGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayoutGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payoutGateway/deletePayoutGatewayByIds [delete]
func (payoutGatewayApi *PayoutGatewayApi) DeletePayoutGatewayByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := payoutGatewayService.DeletePayoutGatewayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayoutGateway 更新PayoutGateway
// @Tags PayoutGateway
// @Summary 更新PayoutGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.PayoutGateway true "更新PayoutGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payoutGateway/updatePayoutGateway [put]
func (payoutGatewayApi *PayoutGatewayApi) UpdatePayoutGateway(c *gin.Context) {
	var payoutGateway dvfpay.PayoutGateway
	_ = c.ShouldBindJSON(&payoutGateway)
	if err := payoutGatewayService.UpdatePayoutGateway(payoutGateway); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPayoutGateway 用id查询PayoutGateway
// @Tags PayoutGateway
// @Summary 用id查询PayoutGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.PayoutGateway true "用id查询PayoutGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payoutGateway/findPayoutGateway [get]
func (payoutGatewayApi *PayoutGatewayApi) FindPayoutGateway(c *gin.Context) {
	var payoutGateway dvfpay.PayoutGateway
	_ = c.ShouldBindQuery(&payoutGateway)
	if err, repayoutGateway := payoutGatewayService.GetPayoutGateway(payoutGateway.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repayoutGateway": repayoutGateway}, c)
	}
}

// GetPayoutGatewayList 分页获取PayoutGateway列表
// @Tags PayoutGateway
// @Summary 分页获取PayoutGateway列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.PayoutGatewaySearch true "分页获取PayoutGateway列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutGateway/getPayoutGatewayList [get]
func (payoutGatewayApi *PayoutGatewayApi) GetPayoutGatewayList(c *gin.Context) {
	var pageInfo dvfpayReq.PayoutGatewaySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := payoutGatewayService.GetPayoutGatewayInfoList(pageInfo); err != nil {
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
