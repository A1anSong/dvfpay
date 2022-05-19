import service from '@/utils/request'

// @Tags PayoutOrder
// @Summary 创建PayoutOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayoutOrder true "创建PayoutOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutOrder/createPayoutOrder [post]
export const createPayoutOrder = (data) => {
  return service({
    url: '/payoutOrder/createPayoutOrder',
    method: 'post',
    data
  })
}

// @Tags PayoutOrder
// @Summary 删除PayoutOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayoutOrder true "删除PayoutOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payoutOrder/deletePayoutOrder [delete]
export const deletePayoutOrder = (data) => {
  return service({
    url: '/payoutOrder/deletePayoutOrder',
    method: 'delete',
    data
  })
}

// @Tags PayoutOrder
// @Summary 删除PayoutOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayoutOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payoutOrder/deletePayoutOrder [delete]
export const deletePayoutOrderByIds = (data) => {
  return service({
    url: '/payoutOrder/deletePayoutOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags PayoutOrder
// @Summary 更新PayoutOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayoutOrder true "更新PayoutOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payoutOrder/updatePayoutOrder [put]
export const updatePayoutOrder = (data) => {
  return service({
    url: '/payoutOrder/updatePayoutOrder',
    method: 'put',
    data
  })
}

// @Tags PayoutOrder
// @Summary 用id查询PayoutOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayoutOrder true "用id查询PayoutOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payoutOrder/findPayoutOrder [get]
export const findPayoutOrder = (params) => {
  return service({
    url: '/payoutOrder/findPayoutOrder',
    method: 'get',
    params
  })
}

// @Tags PayoutOrder
// @Summary 分页获取PayoutOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PayoutOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutOrder/getPayoutOrderList [get]
export const getPayoutOrderList = (params) => {
  return service({
    url: '/payoutOrder/getPayoutOrderList',
    method: 'get',
    params
  })
}

export const getMerchantPayoutOrderList = (params) => {
  return service({
    url: '/payoutOrder/getMerchantPayoutOrderList',
    method: 'get',
    params
  })
}
