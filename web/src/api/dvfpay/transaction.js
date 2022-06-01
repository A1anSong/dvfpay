import service from '@/utils/request'

// @Tags Transaction
// @Summary 创建Transaction
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Transaction true "创建Transaction"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /transaction/createTransaction [post]
export const createTransaction = (data) => {
  return service({
    url: '/transaction/createTransaction',
    method: 'post',
    data
  })
}

// @Tags Transaction
// @Summary 删除Transaction
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Transaction true "删除Transaction"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /transaction/deleteTransaction [delete]
export const deleteTransaction = (data) => {
  return service({
    url: '/transaction/deleteTransaction',
    method: 'delete',
    data
  })
}

// @Tags Transaction
// @Summary 删除Transaction
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Transaction"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /transaction/deleteTransaction [delete]
export const deleteTransactionByIds = (data) => {
  return service({
    url: '/transaction/deleteTransactionByIds',
    method: 'delete',
    data
  })
}

// @Tags Transaction
// @Summary 更新Transaction
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Transaction true "更新Transaction"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /transaction/updateTransaction [put]
export const updateTransaction = (data) => {
  return service({
    url: '/transaction/updateTransaction',
    method: 'put',
    data
  })
}

// @Tags Transaction
// @Summary 用id查询Transaction
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Transaction true "用id查询Transaction"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /transaction/findTransaction [get]
export const findTransaction = (params) => {
  return service({
    url: '/transaction/findTransaction',
    method: 'get',
    params
  })
}

// @Tags Transaction
// @Summary 分页获取Transaction列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Transaction列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /transaction/getTransactionList [get]
export const getTransactionList = (params) => {
  return service({
    url: '/transaction/getTransactionList',
    method: 'get',
    params
  })
}

export const getMerchantTransactionList = (params) => {
  return service({
    url: '/transaction/getMerchantTransactionList',
    method: 'get',
    params
  })
}
