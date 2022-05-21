import service from '@/utils/request'

// @Tags MerchantFunds
// @Summary 创建MerchantFunds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MerchantFunds true "创建MerchantFunds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /merchantFunds/createMerchantFunds [post]
export const createMerchantFunds = (data) => {
  return service({
    url: '/merchantFunds/createMerchantFunds',
    method: 'post',
    data
  })
}

// @Tags MerchantFunds
// @Summary 删除MerchantFunds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MerchantFunds true "删除MerchantFunds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /merchantFunds/deleteMerchantFunds [delete]
export const deleteMerchantFunds = (data) => {
  return service({
    url: '/merchantFunds/deleteMerchantFunds',
    method: 'delete',
    data
  })
}

// @Tags MerchantFunds
// @Summary 删除MerchantFunds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MerchantFunds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /merchantFunds/deleteMerchantFunds [delete]
export const deleteMerchantFundsByIds = (data) => {
  return service({
    url: '/merchantFunds/deleteMerchantFundsByIds',
    method: 'delete',
    data
  })
}

// @Tags MerchantFunds
// @Summary 更新MerchantFunds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MerchantFunds true "更新MerchantFunds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /merchantFunds/updateMerchantFunds [put]
export const updateMerchantFunds = (data) => {
  return service({
    url: '/merchantFunds/updateMerchantFunds',
    method: 'put',
    data
  })
}

// @Tags MerchantFunds
// @Summary 用id查询MerchantFunds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MerchantFunds true "用id查询MerchantFunds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /merchantFunds/findMerchantFunds [get]
export const findMerchantFunds = (params) => {
  return service({
    url: '/merchantFunds/findMerchantFunds',
    method: 'get',
    params
  })
}

// @Tags MerchantFunds
// @Summary 分页获取MerchantFunds列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MerchantFunds列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /merchantFunds/getMerchantFundsList [get]
export const getMerchantFundsList = (params) => {
  return service({
    url: '/merchantFunds/getMerchantFundsList',
    method: 'get',
    params
  })
}

export const getSelfMerchantFundsList = (params) => {
  return service({
    url: '/merchantFunds/getSelfMerchantFundsList',
    method: 'get',
    params
  })
}
