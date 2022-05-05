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
	var incomeGatewayAuth dvfpay.IncomeGatewayAuth
	_ = c.ShouldBindJSON(&incomeGatewayAuth)
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
	var incomeGatewayAuth dvfpay.IncomeGatewayAuth
	_ = c.ShouldBindJSON(&incomeGatewayAuth)
	if err := incomeGatewayAuthService.UpdateIncomeGatewayAuth(incomeGatewayAuth); err != nil {
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
