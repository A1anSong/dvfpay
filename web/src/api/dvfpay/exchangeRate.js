import service from '@/utils/request'

// @Tags ExchangeRate
// @Summary 创建ExchangeRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeRate true "创建ExchangeRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeRate/createExchangeRate [post]
export const createExchangeRate = (data) => {
  return service({
    url: '/exchangeRate/createExchangeRate',
    method: 'post',
    data
  })
}

// @Tags ExchangeRate
// @Summary 删除ExchangeRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeRate true "删除ExchangeRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeRate/deleteExchangeRate [delete]
export const deleteExchangeRate = (data) => {
  return service({
    url: '/exchangeRate/deleteExchangeRate',
    method: 'delete',
    data
  })
}

// @Tags ExchangeRate
// @Summary 删除ExchangeRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExchangeRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeRate/deleteExchangeRate [delete]
export const deleteExchangeRateByIds = (data) => {
  return service({
    url: '/exchangeRate/deleteExchangeRateByIds',
    method: 'delete',
    data
  })
}

// @Tags ExchangeRate
// @Summary 更新ExchangeRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeRate true "更新ExchangeRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exchangeRate/updateExchangeRate [put]
export const updateExchangeRate = (data) => {
  return service({
    url: '/exchangeRate/updateExchangeRate',
    method: 'put',
    data
  })
}

// @Tags ExchangeRate
// @Summary 用id查询ExchangeRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExchangeRate true "用id查询ExchangeRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exchangeRate/findExchangeRate [get]
export const findExchangeRate = (params) => {
  return service({
    url: '/exchangeRate/findExchangeRate',
    method: 'get',
    params
  })
}

// @Tags ExchangeRate
// @Summary 分页获取ExchangeRate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ExchangeRate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeRate/getExchangeRateList [get]
export const getExchangeRateList = (params) => {
  return service({
    url: '/exchangeRate/getExchangeRateList',
    method: 'get',
    params
  })
}
