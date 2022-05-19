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

type IncomeOrderApi struct {
}

var incomeOrderService = service.ServiceGroupApp.DvfpayServiceGroup.IncomeOrderService

// CreateIncomeOrder 创建IncomeOrder
// @Tags IncomeOrder
// @Summary 创建IncomeOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.IncomeOrder true "创建IncomeOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeOrder/createIncomeOrder [post]
func (incomeOrderApi *IncomeOrderApi) CreateIncomeOrder(c *gin.Context) {
	var incomeOrder dvfpay.IncomeOrder
	_ = c.ShouldBindJSON(&incomeOrder)
	if err := incomeOrderService.CreateIncomeOrder(incomeOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIncomeOrder 删除IncomeOrder
// @Tags IncomeOrder
// @Summary 删除IncomeOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.IncomeOrder true "删除IncomeOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /incomeOrder/deleteIncomeOrder [delete]
func (incomeOrderApi *IncomeOrderApi) DeleteIncomeOrder(c *gin.Context) {
	var incomeOrder dvfpay.IncomeOrder
	_ = c.ShouldBindJSON(&incomeOrder)
	if err := incomeOrderService.DeleteIncomeOrder(incomeOrder); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIncomeOrderByIds 批量删除IncomeOrder
// @Tags IncomeOrder
// @Summary 批量删除IncomeOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IncomeOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /incomeOrder/deleteIncomeOrderByIds [delete]
func (incomeOrderApi *IncomeOrderApi) DeleteIncomeOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := incomeOrderService.DeleteIncomeOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIncomeOrder 更新IncomeOrder
// @Tags IncomeOrder
// @Summary 更新IncomeOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.IncomeOrder true "更新IncomeOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /incomeOrder/updateIncomeOrder [put]
func (incomeOrderApi *IncomeOrderApi) UpdateIncomeOrder(c *gin.Context) {
	var incomeOrder dvfpay.IncomeOrder
	_ = c.ShouldBindJSON(&incomeOrder)
	if err := incomeOrderService.UpdateIncomeOrder(incomeOrder); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIncomeOrder 用id查询IncomeOrder
// @Tags IncomeOrder
// @Summary 用id查询IncomeOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.IncomeOrder true "用id查询IncomeOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /incomeOrder/findIncomeOrder [get]
func (incomeOrderApi *IncomeOrderApi) FindIncomeOrder(c *gin.Context) {
	var incomeOrder dvfpay.IncomeOrder
	_ = c.ShouldBindQuery(&incomeOrder)
	if err, reincomeOrder := incomeOrderService.GetIncomeOrder(incomeOrder.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reincomeOrder": reincomeOrder}, c)
	}
}

// GetIncomeOrderList 分页获取IncomeOrder列表
// @Tags IncomeOrder
// @Summary 分页获取IncomeOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.IncomeOrderSearch true "分页获取IncomeOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeOrder/getIncomeOrderList [get]
func (incomeOrderApi *IncomeOrderApi) GetIncomeOrderList(c *gin.Context) {
	var pageInfo dvfpayReq.IncomeOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := incomeOrderService.GetIncomeOrderInfoList(pageInfo); err != nil {
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

func (incomeOrderApi *IncomeOrderApi) GetMerchantIncomeOrderList(c *gin.Context) {
	var pageInfo dvfpayReq.IncomeOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	merchantID := utils.GetUserID(c)
	if err, list, total := incomeOrderService.GetMerchantIncomeOrderInfoList(pageInfo, merchantID); err != nil {
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

func (incomeOrderApi *IncomeOrderApi) ConfirmIncomeOrder(c *gin.Context) {
	var incomeOrder dvfpay.IncomeOrder
	_ = c.ShouldBindJSON(&incomeOrder)
	if err := incomeOrderService.ConfirmIncomeOrder(incomeOrder); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
