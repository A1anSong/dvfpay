package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
	dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
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
	type PayoutGatewayAuthRequest struct {
		PayoutGatewayId uint   `json:"payoutGatewayId"`
		Merchants       []uint `json:"merchants"`
		Fee             int    `json:"fee"`
		LimitMax        int    `json:"limitMax"`
		LimitMin        int    `json:"limitMin"`
		LimitDay        int    `json:"limitDay"`
		LimitTotal      int    `json:"limitTotal"`
	}
	var payoutGatewayAuthRequest PayoutGatewayAuthRequest
	_ = c.ShouldBindJSON(&payoutGatewayAuthRequest)
	var merchants []*system.SysUser
	for _, v := range payoutGatewayAuthRequest.Merchants {
		merchants = append(merchants, &system.SysUser{
			GVA_MODEL: global.GVA_MODEL{
				ID: v,
			},
		})
	}
	payoutGatewayAuth := dvfpay.PayoutGatewayAuth{
		PayoutGateway: dvfpay.PayoutGateway{
			GVA_MODEL: global.GVA_MODEL{
				ID: payoutGatewayAuthRequest.PayoutGatewayId,
			},
		},
		Merchants:  merchants,
		Fee:        &payoutGatewayAuthRequest.Fee,
		LimitMax:   &payoutGatewayAuthRequest.LimitMax,
		LimitMin:   &payoutGatewayAuthRequest.LimitMin,
		LimitDay:   &payoutGatewayAuthRequest.LimitDay,
		LimitTotal: &payoutGatewayAuthRequest.LimitTotal,
	}
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

	global.GVA_DB.Exec("delete "+
		"from dvfpay_payout_gateway_auth_merchants "+
		"where payout_gateway_auth_id = ?", payoutGatewayAuth.ID)

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
	type PayoutGatewayAuthRequest struct {
		ID              uint   `json:"ID"`
		PayoutGatewayId uint   `json:"payoutGatewayId"`
		Merchants       []uint `json:"merchants"`
		Fee             int    `json:"fee"`
		LimitMax        int    `json:"limitMax"`
		LimitMin        int    `json:"limitMin"`
		LimitDay        int    `json:"limitDay"`
		LimitTotal      int    `json:"limitTotal"`
	}
	var payoutGatewayAuthRequest PayoutGatewayAuthRequest
	_ = c.ShouldBindJSON(&payoutGatewayAuthRequest)
	var merchants []*system.SysUser
	for _, v := range payoutGatewayAuthRequest.Merchants {
		merchants = append(merchants, &system.SysUser{
			GVA_MODEL: global.GVA_MODEL{
				ID: v,
			},
		})
	}

	global.GVA_DB.Exec("delete "+
		"from dvfpay_payout_gateway_auth_merchants "+
		"where payout_gateway_auth_id = ?", payoutGatewayAuthRequest.ID)

	var payoutGatewayAuth dvfpay.PayoutGatewayAuth
	global.GVA_DB.First(&payoutGatewayAuth, payoutGatewayAuthRequest.ID)
	var payoutGateWay dvfpay.PayoutGateway
	global.GVA_DB.First(&payoutGateWay, payoutGatewayAuthRequest.PayoutGatewayId)
	payoutGatewayAuth.PayoutGateway = payoutGateWay
	payoutGatewayAuth.Merchants = merchants
	payoutGatewayAuth.Fee = &payoutGatewayAuthRequest.Fee
	payoutGatewayAuth.LimitMax = &payoutGatewayAuthRequest.LimitMax
	payoutGatewayAuth.LimitMin = &payoutGatewayAuthRequest.LimitMin
	payoutGatewayAuth.LimitDay = &payoutGatewayAuthRequest.LimitDay
	payoutGatewayAuth.LimitTotal = &payoutGatewayAuthRequest.LimitTotal
	if err := global.GVA_DB.Save(&payoutGatewayAuth).Error; err != nil {
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

func (payoutGatewayAuthApi *PayoutGatewayAuthApi) GetMerchantPayoutGatewayAuthList(c *gin.Context) {
	var pageInfo dvfpayReq.PayoutGatewayAuthSearch
	_ = c.ShouldBindQuery(&pageInfo)
	merchantID := utils.GetUserID(c)
	if err, list, total := payoutGatewayAuthService.GetMerchantPayoutGatewayAuthInfoList(pageInfo, merchantID); err != nil {
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
