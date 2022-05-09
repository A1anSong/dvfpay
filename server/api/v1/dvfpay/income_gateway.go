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

type IncomeGatewayApi struct {
}

var incomeGatewayService = service.ServiceGroupApp.DvfpayServiceGroup.IncomeGatewayService

// CreateIncomeGateway 创建IncomeGateway
// @Tags IncomeGateway
// @Summary 创建IncomeGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.IncomeGateway true "创建IncomeGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeGateway/createIncomeGateway [post]
func (incomeGatewayApi *IncomeGatewayApi) CreateIncomeGateway(c *gin.Context) {
	var incomeGateway dvfpay.IncomeGateway
	_ = c.ShouldBindJSON(&incomeGateway)
	if err := incomeGatewayService.CreateIncomeGateway(incomeGateway); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIncomeGateway 删除IncomeGateway
// @Tags IncomeGateway
// @Summary 删除IncomeGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.IncomeGateway true "删除IncomeGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /incomeGateway/deleteIncomeGateway [delete]
func (incomeGatewayApi *IncomeGatewayApi) DeleteIncomeGateway(c *gin.Context) {
	var incomeGateway dvfpay.IncomeGateway
	_ = c.ShouldBindJSON(&incomeGateway)
	if err := incomeGatewayService.DeleteIncomeGateway(incomeGateway); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIncomeGatewayByIds 批量删除IncomeGateway
// @Tags IncomeGateway
// @Summary 批量删除IncomeGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IncomeGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /incomeGateway/deleteIncomeGatewayByIds [delete]
func (incomeGatewayApi *IncomeGatewayApi) DeleteIncomeGatewayByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := incomeGatewayService.DeleteIncomeGatewayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIncomeGateway 更新IncomeGateway
// @Tags IncomeGateway
// @Summary 更新IncomeGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.IncomeGateway true "更新IncomeGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /incomeGateway/updateIncomeGateway [put]
func (incomeGatewayApi *IncomeGatewayApi) UpdateIncomeGateway(c *gin.Context) {
	var incomeGateway dvfpay.IncomeGateway
	_ = c.ShouldBindJSON(&incomeGateway)
	if err := incomeGatewayService.UpdateIncomeGateway(incomeGateway); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIncomeGateway 用id查询IncomeGateway
// @Tags IncomeGateway
// @Summary 用id查询IncomeGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.IncomeGateway true "用id查询IncomeGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /incomeGateway/findIncomeGateway [get]
func (incomeGatewayApi *IncomeGatewayApi) FindIncomeGateway(c *gin.Context) {
	var incomeGateway dvfpay.IncomeGateway
	_ = c.ShouldBindQuery(&incomeGateway)
	if err, reincomeGateway := incomeGatewayService.GetIncomeGateway(incomeGateway.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reincomeGateway": reincomeGateway}, c)
	}
}

// GetIncomeGatewayList 分页获取IncomeGateway列表
// @Tags IncomeGateway
// @Summary 分页获取IncomeGateway列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.IncomeGatewaySearch true "分页获取IncomeGateway列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeGateway/getIncomeGatewayList [get]
func (incomeGatewayApi *IncomeGatewayApi) GetIncomeGatewayList(c *gin.Context) {
	var pageInfo dvfpayReq.IncomeGatewaySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := incomeGatewayService.GetIncomeGatewayInfoList(pageInfo); err != nil {
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
