import service from '@/utils/request'

// @Tags ExchangeRecord
// @Summary 创建ExchangeRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeRecord true "创建ExchangeRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeRecord/createExchangeRecord [post]
export const createExchangeRecord = (data) => {
  return service({
    url: '/exchangeRecord/createExchangeRecord',
    method: 'post',
    data
  })
}

// @Tags ExchangeRecord
// @Summary 删除ExchangeRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeRecord true "删除ExchangeRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeRecord/deleteExchangeRecord [delete]
export const deleteExchangeRecord = (data) => {
  return service({
    url: '/exchangeRecord/deleteExchangeRecord',
    method: 'delete',
    data
  })
}

// @Tags ExchangeRecord
// @Summary 删除ExchangeRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExchangeRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeRecord/deleteExchangeRecord [delete]
export const deleteExchangeRecordByIds = (data) => {
  return service({
    url: '/exchangeRecord/deleteExchangeRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags ExchangeRecord
// @Summary 更新ExchangeRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeRecord true "更新ExchangeRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exchangeRecord/updateExchangeRecord [put]
export const updateExchangeRecord = (data) => {
  return service({
    url: '/exchangeRecord/updateExchangeRecord',
    method: 'put',
    data
  })
}

// @Tags ExchangeRecord
// @Summary 用id查询ExchangeRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExchangeRecord true "用id查询ExchangeRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exchangeRecord/findExchangeRecord [get]
export const findExchangeRecord = (params) => {
  return service({
    url: '/exchangeRecord/findExchangeRecord',
    method: 'get',
    params
  })
}

// @Tags ExchangeRecord
// @Summary 分页获取ExchangeRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ExchangeRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeRecord/getExchangeRecordList [get]
export const getExchangeRecordList = (params) => {
  return service({
    url: '/exchangeRecord/getExchangeRecordList',
    method: 'get',
    params
  })
}

export const getMerchantExchangeRecordList = (params) => {
  return service({
    url: '/exchangeRecord/getMerchantExchangeRecordList',
    method: 'get',
    params
  })
}
