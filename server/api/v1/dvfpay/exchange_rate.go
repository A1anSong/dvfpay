package dvfpay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dvfpayReq "github.com/flipped-aurora/gin-vue-admin/server/model/dvfpay/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ExchangeRateApi struct {
}

var exchangeRateService = service.ServiceGroupApp.DvfpayServiceGroup.ExchangeRateService


// CreateExchangeRate 创建ExchangeRate
// @Tags ExchangeRate
// @Summary 创建ExchangeRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.ExchangeRate true "创建ExchangeRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeRate/createExchangeRate [post]
func (exchangeRateApi *ExchangeRateApi) CreateExchangeRate(c *gin.Context) {
	var exchangeRate dvfpay.ExchangeRate
	_ = c.ShouldBindJSON(&exchangeRate)
	if err := exchangeRateService.CreateExchangeRate(exchangeRate); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteExchangeRate 删除ExchangeRate
// @Tags ExchangeRate
// @Summary 删除ExchangeRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.ExchangeRate true "删除ExchangeRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeRate/deleteExchangeRate [delete]
func (exchangeRateApi *ExchangeRateApi) DeleteExchangeRate(c *gin.Context) {
	var exchangeRate dvfpay.ExchangeRate
	_ = c.ShouldBindJSON(&exchangeRate)
	if err := exchangeRateService.DeleteExchangeRate(exchangeRate); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExchangeRateByIds 批量删除ExchangeRate
// @Tags ExchangeRate
// @Summary 批量删除ExchangeRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExchangeRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exchangeRate/deleteExchangeRateByIds [delete]
func (exchangeRateApi *ExchangeRateApi) DeleteExchangeRateByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := exchangeRateService.DeleteExchangeRateByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExchangeRate 更新ExchangeRate
// @Tags ExchangeRate
// @Summary 更新ExchangeRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dvfpay.ExchangeRate true "更新ExchangeRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exchangeRate/updateExchangeRate [put]
func (exchangeRateApi *ExchangeRateApi) UpdateExchangeRate(c *gin.Context) {
	var exchangeRate dvfpay.ExchangeRate
	_ = c.ShouldBindJSON(&exchangeRate)
	if err := exchangeRateService.UpdateExchangeRate(exchangeRate); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExchangeRate 用id查询ExchangeRate
// @Tags ExchangeRate
// @Summary 用id查询ExchangeRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpay.ExchangeRate true "用id查询ExchangeRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exchangeRate/findExchangeRate [get]
func (exchangeRateApi *ExchangeRateApi) FindExchangeRate(c *gin.Context) {
	var exchangeRate dvfpay.ExchangeRate
	_ = c.ShouldBindQuery(&exchangeRate)
	if err, reexchangeRate := exchangeRateService.GetExchangeRate(exchangeRate.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexchangeRate": reexchangeRate}, c)
	}
}

// GetExchangeRateList 分页获取ExchangeRate列表
// @Tags ExchangeRate
// @Summary 分页获取ExchangeRate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dvfpayReq.ExchangeRateSearch true "分页获取ExchangeRate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeRate/getExchangeRateList [get]
func (exchangeRateApi *ExchangeRateApi) GetExchangeRateList(c *gin.Context) {
	var pageInfo dvfpayReq.ExchangeRateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := exchangeRateService.GetExchangeRateInfoList(pageInfo); err != nil {
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
