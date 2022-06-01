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

export const statusType = (status) => {
  switch (status) {
    case 'success':
      return 'success'
    case 'PENDING':
      return 'warning'
    case 'SETTLED':
      return ''
    default:
      return status
  }
}
