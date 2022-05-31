export const statusDisplay = (status) => {
  switch (status) {
    case 'success':
      return '成功'
    case 'PENDING':
      return '待定'
    case 'SETTLED':
      return '结算'
    default:
      return status
  }
}
