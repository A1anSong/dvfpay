import service from '@/utils/request'

// @Tags IncomeOrder
// @Summary 创建IncomeOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeOrder true "创建IncomeOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeOrder/createIncomeOrder [post]
export const createIncomeOrder = (data) => {
  return service({
    url: '/incomeOrder/createIncomeOrder',
    method: 'post',
    data
  })
}

// @Tags IncomeOrder
// @Summary 删除IncomeOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeOrder true "删除IncomeOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /incomeOrder/deleteIncomeOrder [delete]
export const deleteIncomeOrder = (data) => {
  return service({
    url: '/incomeOrder/deleteIncomeOrder',
    method: 'delete',
    data
  })
}

// @Tags IncomeOrder
// @Summary 删除IncomeOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IncomeOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /incomeOrder/deleteIncomeOrder [delete]
export const deleteIncomeOrderByIds = (data) => {
  return service({
    url: '/incomeOrder/deleteIncomeOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags IncomeOrder
// @Summary 更新IncomeOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeOrder true "更新IncomeOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /incomeOrder/updateIncomeOrder [put]
export const updateIncomeOrder = (data) => {
  return service({
    url: '/incomeOrder/updateIncomeOrder',
    method: 'put',
    data
  })
}

// @Tags IncomeOrder
// @Summary 用id查询IncomeOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IncomeOrder true "用id查询IncomeOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /incomeOrder/findIncomeOrder [get]
export const findIncomeOrder = (params) => {
  return service({
    url: '/incomeOrder/findIncomeOrder',
    method: 'get',
    params
  })
}

// @Tags IncomeOrder
// @Summary 分页获取IncomeOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IncomeOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeOrder/getIncomeOrderList [get]
export const getIncomeOrderList = (params) => {
  return service({
    url: '/incomeOrder/getIncomeOrderList',
    method: 'get',
    params
  })
}

export const getMerchantIncomeOrderList = (params) => {
  return service({
    url: '/incomeOrder/getMerchantIncomeOrderList',
    method: 'get',
    params
  })
}

export const confirmIncomeOrder = (params) => {
  return service({
    url: '/incomeOrder/confirmIncomeOrder',
    method: 'put',
    params
  })
}

export const getMerchantStatisticsIncomeOrder = (params) => {
  return service({
    url: '/incomeOrder/getMerchantStatisticsIncomeOrder',
    method: 'get',
    params
  })
}

export const getMerchantTrendsCountIncomeOrder = (params) => {
  return service({
    url: '/incomeOrder/getMerchantTrendsCountIncomeOrder',
    method: 'get',
    params
  })
}

export const getMerchantTrendsSumIncomeOrder = (params) => {
  return service({
    url: '/incomeOrder/getMerchantTrendsSumIncomeOrder',
    method: 'get',
    params
  })
}
