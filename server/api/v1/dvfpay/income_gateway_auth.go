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

type IncomeGatewayAuthApi struct {
}

var incomeGatewayAuthService = service.ServiceGroupApp.DvfpayServiceGroup.IncomeGatewayAuthService

// CreateIncomeGatewayAuth 创建IncomeGatewayAuth
// @Tags IncomeGatewayAuth
// @Summary 创建IncomeGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.IncomeGatewayAuth true "创建IncomeGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeGatewayAuth/createIncomeGatewayAuth [post]
func (incomeGatewayAuthApi *IncomeGatewayAuthApi) CreateIncomeGatewayAuth(c *gin.Context) {
	type IncomeGatewayAuthRequest struct {
		IncomeGatewayId uint   `json:"incomeGatewayId"`
		Merchants       []uint `json:"merchants"`
		Fee             int    `json:"fee"`
		LimitMax        int    `json:"limitMax"`
		LimitMin        int    `json:"limitMin"`
		LimitDay        int    `json:"limitDay"`
		LimitTotal      int    `json:"limitTotal"`
		Explain         string `json:"explain"`
	}
	var incomeGatewayAuthRequest IncomeGatewayAuthRequest
	_ = c.ShouldBindJSON(&incomeGatewayAuthRequest)
	var merchants []*system.SysUser
	for _, v := range incomeGatewayAuthRequest.Merchants {
		merchants = append(merchants, &system.SysUser{
			GVA_MODEL: global.GVA_MODEL{
				ID: v,
			},
		})
	}
	incomeGatewayAuth := dvfpay.IncomeGatewayAuth{
		IncomeGateway: dvfpay.IncomeGateway{
			GVA_MODEL: global.GVA_MODEL{
				ID: incomeGatewayAuthRequest.IncomeGatewayId,
			},
		},
		Merchants:  merchants,
		Fee:        &incomeGatewayAuthRequest.Fee,
		LimitMax:   &incomeGatewayAuthRequest.LimitMax,
		LimitMin:   &incomeGatewayAuthRequest.LimitMin,
		LimitDay:   &incomeGatewayAuthRequest.LimitDay,
		LimitTotal: &incomeGatewayAuthRequest.LimitTotal,
		Explain:    incomeGatewayAuthRequest.Explain,
	}
	if err := incomeGatewayAuthService.CreateIncomeGatewayAuth(incomeGatewayAuth); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIncomeGatewayAuth 删除IncomeGatewayAuth
// @Tags IncomeGatewayAuth
// @Summary 删除IncomeGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.IncomeGatewayAuth true "删除IncomeGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /incomeGatewayAuth/deleteIncomeGatewayAuth [delete]
func (incomeGatewayAuthApi *IncomeGatewayAuthApi) DeleteIncomeGatewayAuth(c *gin.Context) {
	var incomeGatewayAuth dvfpay.IncomeGatewayAuth
	_ = c.ShouldBindJSON(&incomeGatewayAuth)

	global.GVA_DB.Exec("delete "+
		"from dvfpay_income_gateway_auth_merchants "+
		"where income_gateway_auth_id = ?", incomeGatewayAuth.ID)

	if err := incomeGatewayAuthService.DeleteIncomeGatewayAuth(incomeGatewayAuth); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIncomeGatewayAuthByIds 批量删除IncomeGatewayAuth
// @Tags IncomeGatewayAuth
// @Summary 批量删除IncomeGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IncomeGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /incomeGatewayAuth/deleteIncomeGatewayAuthByIds [delete]
func (incomeGatewayAuthApi *IncomeGatewayAuthApi) DeleteIncomeGatewayAuthByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := incomeGatewayAuthService.DeleteIncomeGatewayAuthByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIncomeGatewayAuth 更新IncomeGatewayAuth
// @Tags IncomeGatewayAuth
// @Summary 更新IncomeGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.IncomeGatewayAuth true "更新IncomeGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /incomeGatewayAuth/updateIncomeGatewayAuth [put]
func (incomeGatewayAuthApi *IncomeGatewayAuthApi) UpdateIncomeGatewayAuth(c *gin.Context) {
	type IncomeGatewayAuthRequest struct {
		ID              uint   `json:"ID"`
		IncomeGatewayId uint   `json:"incomeGatewayId"`
		Merchants       []uint `json:"merchants"`
		Fee             int    `json:"fee"`
		LimitMax        int    `json:"limitMax"`
		LimitMin        int    `json:"limitMin"`
		LimitDay        int    `json:"limitDay"`
		LimitTotal      int    `json:"limitTotal"`
		Explain         string `json:"explain"`
	}
	var incomeGatewayAuthRequest IncomeGatewayAuthRequest
	_ = c.ShouldBindJSON(&incomeGatewayAuthRequest)
	var merchants []*system.SysUser
	for _, v := range incomeGatewayAuthRequest.Merchants {
		merchants = append(merchants, &system.SysUser{
			GVA_MODEL: global.GVA_MODEL{
				ID: v,
			},
		})
	}

	global.GVA_DB.Exec("delete "+
		"from dvfpay_income_gateway_auth_merchants "+
		"where income_gateway_auth_id = ?", incomeGatewayAuthRequest.ID)

	var incomeGatewayAuth dvfpay.IncomeGatewayAuth
	global.GVA_DB.First(&incomeGatewayAuth, incomeGatewayAuthRequest.ID)
	var incomeGateWay dvfpay.IncomeGateway
	global.GVA_DB.First(&incomeGateWay, incomeGatewayAuthRequest.IncomeGatewayId)
	incomeGatewayAuth.IncomeGateway = incomeGateWay
	incomeGatewayAuth.Merchants = merchants
	incomeGatewayAuth.Fee = &incomeGatewayAuthRequest.Fee
	incomeGatewayAuth.LimitMax = &incomeGatewayAuthRequest.LimitMax
	incomeGatewayAuth.LimitMin = &incomeGatewayAuthRequest.LimitMin
	incomeGatewayAuth.LimitDay = &incomeGatewayAuthRequest.LimitDay
	incomeGatewayAuth.LimitTotal = &incomeGatewayAuthRequest.LimitTotal
	incomeGatewayAuth.Explain = incomeGatewayAuthRequest.Explain
	if err := global.GVA_DB.Save(&incomeGatewayAuth).Error; err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIncomeGatewayAuth 用id查询IncomeGatewayAuth
// @Tags IncomeGatewayAuth
// @Summary 用id查询IncomeGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.IncomeGatewayAuth true "用id查询IncomeGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /incomeGatewayAuth/findIncomeGatewayAuth [get]
func (incomeGatewayAuthApi *IncomeGatewayAuthApi) FindIncomeGatewayAuth(c *gin.Context) {
	var incomeGatewayAuth dvfpay.IncomeGatewayAuth
	_ = c.ShouldBindQuery(&incomeGatewayAuth)
	if err, reincomeGatewayAuth := incomeGatewayAuthService.GetIncomeGatewayAuth(incomeGatewayAuth.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reincomeGatewayAuth": reincomeGatewayAuth}, c)
	}
}

// GetIncomeGatewayAuthList 分页获取IncomeGatewayAuth列表
// @Tags IncomeGatewayAuth
// @Summary 分页获取IncomeGatewayAuth列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.IncomeGatewayAuthSearch true "分页获取IncomeGatewayAuth列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeGatewayAuth/getIncomeGatewayAuthList [get]
func (incomeGatewayAuthApi *IncomeGatewayAuthApi) GetIncomeGatewayAuthList(c *gin.Context) {
	var pageInfo dvfpayReq.IncomeGatewayAuthSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := incomeGatewayAuthService.GetIncomeGatewayAuthInfoList(pageInfo); err != nil {
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

func (incomeGatewayAuthApi *IncomeGatewayAuthApi) GetMerchantIncomeGatewayAuthList(c *gin.Context) {
	var pageInfo dvfpayReq.IncomeGatewayAuthSearch
	_ = c.ShouldBindQuery(&pageInfo)
	merchantID := utils.GetUserID(c)
	if err, list, total := incomeGatewayAuthService.GetMerchantIncomeGatewayAuthInfoList(pageInfo, merchantID); err != nil {
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
