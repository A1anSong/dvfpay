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

type MerchantFundsApi struct {
}

var merchantFundsService = service.ServiceGroupApp.DvfpayServiceGroup.MerchantFundsService

// CreateMerchantFunds 创建MerchantFunds
// @Tags MerchantFunds
// @Summary 创建MerchantFunds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.MerchantFunds true "创建MerchantFunds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /merchantFunds/createMerchantFunds [post]
func (merchantFundsApi *MerchantFundsApi) CreateMerchantFunds(c *gin.Context) {
	var merchantFunds dvfpay.MerchantFunds
	_ = c.ShouldBindJSON(&merchantFunds)
	if err := merchantFundsService.CreateMerchantFunds(merchantFunds); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMerchantFunds 删除MerchantFunds
// @Tags MerchantFunds
// @Summary 删除MerchantFunds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.MerchantFunds true "删除MerchantFunds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /merchantFunds/deleteMerchantFunds [delete]
func (merchantFundsApi *MerchantFundsApi) DeleteMerchantFunds(c *gin.Context) {
	var merchantFunds dvfpay.MerchantFunds
	_ = c.ShouldBindJSON(&merchantFunds)
	if err := merchantFundsService.DeleteMerchantFunds(merchantFunds); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMerchantFundsByIds 批量删除MerchantFunds
// @Tags MerchantFunds
// @Summary 批量删除MerchantFunds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MerchantFunds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /merchantFunds/deleteMerchantFundsByIds [delete]
func (merchantFundsApi *MerchantFundsApi) DeleteMerchantFundsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := merchantFundsService.DeleteMerchantFundsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMerchantFunds 更新MerchantFunds
// @Tags MerchantFunds
// @Summary 更新MerchantFunds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.MerchantFunds true "更新MerchantFunds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /merchantFunds/updateMerchantFunds [put]
func (merchantFundsApi *MerchantFundsApi) UpdateMerchantFunds(c *gin.Context) {
	var merchantFunds dvfpay.MerchantFunds
	_ = c.ShouldBindJSON(&merchantFunds)
	if err := merchantFundsService.UpdateMerchantFunds(merchantFunds); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMerchantFunds 用id查询MerchantFunds
// @Tags MerchantFunds
// @Summary 用id查询MerchantFunds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.MerchantFunds true "用id查询MerchantFunds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /merchantFunds/findMerchantFunds [get]
func (merchantFundsApi *MerchantFundsApi) FindMerchantFunds(c *gin.Context) {
	var merchantFunds dvfpay.MerchantFunds
	_ = c.ShouldBindQuery(&merchantFunds)
	if err, remerchantFunds := merchantFundsService.GetMerchantFunds(merchantFunds.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remerchantFunds": remerchantFunds}, c)
	}
}

// GetMerchantFundsList 分页获取MerchantFunds列表
// @Tags MerchantFunds
// @Summary 分页获取MerchantFunds列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.MerchantFundsSearch true "分页获取MerchantFunds列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /merchantFunds/getMerchantFundsList [get]
func (merchantFundsApi *MerchantFundsApi) GetMerchantFundsList(c *gin.Context) {
	var pageInfo dvfpayReq.MerchantFundsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := merchantFundsService.GetMerchantFundsInfoList(pageInfo); err != nil {
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

func (merchantFundsApi *MerchantFundsApi) GetSelfMerchantFundsList(c *gin.Context) {
	var pageInfo dvfpayReq.MerchantFundsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	merchantID := utils.GetUserID(c)
	if err, list, total := merchantFundsService.GetSelfMerchantFundsInfoList(pageInfo, merchantID); err != nil {
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
