import service from '@/utils/request'

// @Tags PayoutGateway
// @Summary 创建PayoutGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayoutGateway true "创建PayoutGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutGateway/createPayoutGateway [post]
export const createPayoutGateway = (data) => {
  return service({
    url: '/payoutGateway/createPayoutGateway',
    method: 'post',
    data
  })
}

// @Tags PayoutGateway
// @Summary 删除PayoutGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayoutGateway true "删除PayoutGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payoutGateway/deletePayoutGateway [delete]
export const deletePayoutGateway = (data) => {
  return service({
    url: '/payoutGateway/deletePayoutGateway',
    method: 'delete',
    data
  })
}

// @Tags PayoutGateway
// @Summary 删除PayoutGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayoutGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payoutGateway/deletePayoutGateway [delete]
export const deletePayoutGatewayByIds = (data) => {
  return service({
    url: '/payoutGateway/deletePayoutGatewayByIds',
    method: 'delete',
    data
  })
}

// @Tags PayoutGateway
// @Summary 更新PayoutGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayoutGateway true "更新PayoutGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payoutGateway/updatePayoutGateway [put]
export const updatePayoutGateway = (data) => {
  return service({
    url: '/payoutGateway/updatePayoutGateway',
    method: 'put',
    data
  })
}

// @Tags PayoutGateway
// @Summary 用id查询PayoutGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayoutGateway true "用id查询PayoutGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payoutGateway/findPayoutGateway [get]
export const findPayoutGateway = (params) => {
  return service({
    url: '/payoutGateway/findPayoutGateway',
    method: 'get',
    params
  })
}

// @Tags PayoutGateway
// @Summary 分页获取PayoutGateway列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PayoutGateway列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payoutGateway/getPayoutGatewayList [get]
export const getPayoutGatewayList = (params) => {
  return service({
    url: '/payoutGateway/getPayoutGatewayList',
    method: 'get',
    params
  })
}
