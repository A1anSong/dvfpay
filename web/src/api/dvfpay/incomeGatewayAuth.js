import service from '@/utils/request'

// @Tags IncomeGatewayAuth
// @Summary 创建IncomeGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeGatewayAuth true "创建IncomeGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeGatewayAuth/createIncomeGatewayAuth [post]
export const createIncomeGatewayAuth = (data) => {
  return service({
    url: '/incomeGatewayAuth/createIncomeGatewayAuth',
    method: 'post',
    data
  })
}

// @Tags IncomeGatewayAuth
// @Summary 删除IncomeGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeGatewayAuth true "删除IncomeGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /incomeGatewayAuth/deleteIncomeGatewayAuth [delete]
export const deleteIncomeGatewayAuth = (data) => {
  return service({
    url: '/incomeGatewayAuth/deleteIncomeGatewayAuth',
    method: 'delete',
    data
  })
}

// @Tags IncomeGatewayAuth
// @Summary 删除IncomeGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IncomeGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /incomeGatewayAuth/deleteIncomeGatewayAuth [delete]
export const deleteIncomeGatewayAuthByIds = (data) => {
  return service({
    url: '/incomeGatewayAuth/deleteIncomeGatewayAuthByIds',
    method: 'delete',
    data
  })
}

// @Tags IncomeGatewayAuth
// @Summary 更新IncomeGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeGatewayAuth true "更新IncomeGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /incomeGatewayAuth/updateIncomeGatewayAuth [put]
export const updateIncomeGatewayAuth = (data) => {
  return service({
    url: '/incomeGatewayAuth/updateIncomeGatewayAuth',
    method: 'put',
    data
  })
}

// @Tags IncomeGatewayAuth
// @Summary 用id查询IncomeGatewayAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IncomeGatewayAuth true "用id查询IncomeGatewayAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /incomeGatewayAuth/findIncomeGatewayAuth [get]
export const findIncomeGatewayAuth = (params) => {
  return service({
    url: '/incomeGatewayAuth/findIncomeGatewayAuth',
    method: 'get',
    params
  })
}

// @Tags IncomeGatewayAuth
// @Summary 分页获取IncomeGatewayAuth列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IncomeGatewayAuth列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeGatewayAuth/getIncomeGatewayAuthList [get]
export const getIncomeGatewayAuthList = (params) => {
  return service({
    url: '/incomeGatewayAuth/getIncomeGatewayAuthList',
    method: 'get',
    params
  })
}
