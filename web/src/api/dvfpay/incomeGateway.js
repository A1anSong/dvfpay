import service from '@/utils/request'

// @Tags IncomeGateway
// @Summary 创建IncomeGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeGateway true "创建IncomeGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeGateway/createIncomeGateway [post]
export const createIncomeGateway = (data) => {
  return service({
    url: '/incomeGateway/createIncomeGateway',
    method: 'post',
    data
  })
}

// @Tags IncomeGateway
// @Summary 删除IncomeGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeGateway true "删除IncomeGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /incomeGateway/deleteIncomeGateway [delete]
export const deleteIncomeGateway = (data) => {
  return service({
    url: '/incomeGateway/deleteIncomeGateway',
    method: 'delete',
    data
  })
}

// @Tags IncomeGateway
// @Summary 删除IncomeGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IncomeGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /incomeGateway/deleteIncomeGateway [delete]
export const deleteIncomeGatewayByIds = (data) => {
  return service({
    url: '/incomeGateway/deleteIncomeGatewayByIds',
    method: 'delete',
    data
  })
}

// @Tags IncomeGateway
// @Summary 更新IncomeGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeGateway true "更新IncomeGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /incomeGateway/updateIncomeGateway [put]
export const updateIncomeGateway = (data) => {
  return service({
    url: '/incomeGateway/updateIncomeGateway',
    method: 'put',
    data
  })
}

// @Tags IncomeGateway
// @Summary 用id查询IncomeGateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IncomeGateway true "用id查询IncomeGateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /incomeGateway/findIncomeGateway [get]
export const findIncomeGateway = (params) => {
  return service({
    url: '/incomeGateway/findIncomeGateway',
    method: 'get',
    params
  })
}

// @Tags IncomeGateway
// @Summary 分页获取IncomeGateway列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IncomeGateway列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /incomeGateway/getIncomeGatewayList [get]
export const getIncomeGatewayList = (params) => {
  return service({
    url: '/incomeGateway/getIncomeGatewayList',
    method: 'get',
    params
  })
}
