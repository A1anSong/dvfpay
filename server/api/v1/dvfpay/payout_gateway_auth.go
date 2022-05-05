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

type PayoutGatewayAuthApi struct {
}

var payoutGatewayAuthService = service.ServiceGroupApp.DvfpayServiceGroup.PayoutGatewayAuthService

// CreatePayoutGatewayAuth 创建PayoutGatewayAuth
// @Tags PayoutGatewayAuth
// @Summary 创建PayoutGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.PayoutGatewayAuth true "创建PayoutGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutGatewayAuth/createPayoutGatewayAuth [post]
func (payoutGatewayAuthApi *PayoutGatewayAuthApi) CreatePayoutGatewayAuth(c *gin.Context) {
	var payoutGatewayAuth dvfpay.PayoutGatewayAuth
	_ = c.ShouldBindJSON(&payoutGatewayAuth)
	if err := payoutGatewayAuthService.CreatePayoutGatewayAuth(payoutGatewayAuth); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePayoutGatewayAuth 删除PayoutGatewayAuth
// @Tags PayoutGatewayAuth
// @Summary 删除PayoutGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.PayoutGatewayAuth true "删除PayoutGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payoutGatewayAuth/deletePayoutGatewayAuth [delete]
func (payoutGatewayAuthApi *PayoutGatewayAuthApi) DeletePayoutGatewayAuth(c *gin.Context) {
	var payoutGatewayAuth dvfpay.PayoutGatewayAuth
	_ = c.ShouldBindJSON(&payoutGatewayAuth)
	if err := payoutGatewayAuthService.DeletePayoutGatewayAuth(payoutGatewayAuth); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayoutGatewayAuthByIds 批量删除PayoutGatewayAuth
// @Tags PayoutGatewayAuth
// @Summary 批量删除PayoutGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayoutGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payoutGatewayAuth/deletePayoutGatewayAuthByIds [delete]
func (payoutGatewayAuthApi *PayoutGatewayAuthApi) DeletePayoutGatewayAuthByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := payoutGatewayAuthService.DeletePayoutGatewayAuthByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayoutGatewayAuth 更新PayoutGatewayAuth
// @Tags PayoutGatewayAuth
// @Summary 更新PayoutGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.PayoutGatewayAuth true "更新PayoutGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payoutGatewayAuth/updatePayoutGatewayAuth [put]
func (payoutGatewayAuthApi *PayoutGatewayAuthApi) UpdatePayoutGatewayAuth(c *gin.Context) {
	var payoutGatewayAuth dvfpay.PayoutGatewayAuth
	_ = c.ShouldBindJSON(&payoutGatewayAuth)
	if err := payoutGatewayAuthService.UpdatePayoutGatewayAuth(payoutGatewayAuth); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPayoutGatewayAuth 用id查询PayoutGatewayAuth
// @Tags PayoutGatewayAuth
// @Summary 用id查询PayoutGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.PayoutGatewayAuth true "用id查询PayoutGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payoutGatewayAuth/findPayoutGatewayAuth [get]
func (payoutGatewayAuthApi *PayoutGatewayAuthApi) FindPayoutGatewayAuth(c *gin.Context) {
	var payoutGatewayAuth dvfpay.PayoutGatewayAuth
	_ = c.ShouldBindQuery(&payoutGatewayAuth)
	if err, repayoutGatewayAuth := payoutGatewayAuthService.GetPayoutGatewayAuth(payoutGatewayAuth.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repayoutGatewayAuth": repayoutGatewayAuth}, c)
	}
}

// GetPayoutGatewayAuthList 分页获取PayoutGatewayAuth列表
// @Tags PayoutGatewayAuth
// @Summary 分页获取PayoutGatewayAuth列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.PayoutGatewayAuthSearch true "分页获取PayoutGatewayAuth列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutGatewayAuth/getPayoutGatewayAuthList [get]
func (payoutGatewayAuthApi *PayoutGatewayAuthApi) GetPayoutGatewayAuthList(c *gin.Context) {
	var pageInfo dvfpayReq.PayoutGatewayAuthSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := payoutGatewayAuthService.GetPayoutGatewayAuthInfoList(pageInfo); err != nil {
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
