import service from '@/utils/request'

// @Tags PayoutGatewayAuth
// @Summary 创建PayoutGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayoutGatewayAuth true "创建PayoutGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutGatewayAuth/createPayoutGatewayAuth [post]
export const createPayoutGatewayAuth = (data) => {
  return service({
    url: '/payoutGatewayAuth/createPayoutGatewayAuth',
    method: 'post',
    data
  })
}

// @Tags PayoutGatewayAuth
// @Summary 删除PayoutGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayoutGatewayAuth true "删除PayoutGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payoutGatewayAuth/deletePayoutGatewayAuth [delete]
export const deletePayoutGatewayAuth = (data) => {
  return service({
    url: '/payoutGatewayAuth/deletePayoutGatewayAuth',
    method: 'delete',
    data
  })
}

// @Tags PayoutGatewayAuth
// @Summary 删除PayoutGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayoutGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payoutGatewayAuth/deletePayoutGatewayAuth [delete]
export const deletePayoutGatewayAuthByIds = (data) => {
  return service({
    url: '/payoutGatewayAuth/deletePayoutGatewayAuthByIds',
    method: 'delete',
    data
  })
}

// @Tags PayoutGatewayAuth
// @Summary 更新PayoutGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayoutGatewayAuth true "更新PayoutGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payoutGatewayAuth/updatePayoutGatewayAuth [put]
export const updatePayoutGatewayAuth = (data) => {
  return service({
    url: '/payoutGatewayAuth/updatePayoutGatewayAuth',
    method: 'put',
    data
  })
}

// @Tags PayoutGatewayAuth
// @Summary 用id查询PayoutGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayoutGatewayAuth true "用id查询PayoutGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payoutGatewayAuth/findPayoutGatewayAuth [get]
export const findPayoutGatewayAuth = (params) => {
  return service({
    url: '/payoutGatewayAuth/findPayoutGatewayAuth',
    method: 'get',
    params
  })
}

// @Tags PayoutGatewayAuth
// @Summary 分页获取PayoutGatewayAuth列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PayoutGatewayAuth列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutGatewayAuth/getPayoutGatewayAuthList [get]
export const getPayoutGatewayAuthList = (params) => {
  return service({
    url: '/payoutGatewayAuth/getPayoutGatewayAuthList',
    method: 'get',
    params
  })
}

export const getMerchantPayoutGatewayAuthList = (params) => {
  return service({
    url: '/payoutGatewayAuth/getMerchantPayoutGatewayAuthList',
    method: 'get',
    params
  })
}
