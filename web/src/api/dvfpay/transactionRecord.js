import service from '@/utils/request'

// @Tags TransactionRecord
// @Summary 创建TransactionRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TransactionRecord true "创建TransactionRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /transactionRecord/createTransactionRecord [post]
export const createTransactionRecord = (data) => {
  return service({
    url: '/transactionRecord/createTransactionRecord',
    method: 'post',
    data
  })
}

// @Tags TransactionRecord
// @Summary 删除TransactionRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TransactionRecord true "删除TransactionRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /transactionRecord/deleteTransactionRecord [delete]
export const deleteTransactionRecord = (data) => {
  return service({
    url: '/transactionRecord/deleteTransactionRecord',
    method: 'delete',
    data
  })
}

// @Tags TransactionRecord
// @Summary 删除TransactionRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TransactionRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /transactionRecord/deleteTransactionRecord [delete]
export const deleteTransactionRecordByIds = (data) => {
  return service({
    url: '/transactionRecord/deleteTransactionRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags TransactionRecord
// @Summary 更新TransactionRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TransactionRecord true "更新TransactionRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /transactionRecord/updateTransactionRecord [put]
export const updateTransactionRecord = (data) => {
  return service({
    url: '/transactionRecord/updateTransactionRecord',
    method: 'put',
    data
  })
}

// @Tags TransactionRecord
// @Summary 用id查询TransactionRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TransactionRecord true "用id查询TransactionRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /transactionRecord/findTransactionRecord [get]
export const findTransactionRecord = (params) => {
  return service({
    url: '/transactionRecord/findTransactionRecord',
    method: 'get',
    params
  })
}

// @Tags TransactionRecord
// @Summary 分页获取TransactionRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TransactionRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /transactionRecord/getTransactionRecordList [get]
export const getTransactionRecordList = (params) => {
  return service({
    url: '/transactionRecord/getTransactionRecordList',
    method: 'get',
    params
  })
}
