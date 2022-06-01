export const transactionOperation = (operation) => {
  switch (operation) {
    case 'DEPOSIT':
      return '充值'
    case 'WITHDRAW':
      return '提现'
    default:
      return operation
  }
}

export const transactionStatus = (status) => {
  switch (status) {
    case 'RELEASED':
      return '已放行'
    case 'REVIEWING':
      return '审核中'
    case 'REJECTED':
      return '已拒绝'
    default:
      return status
  }
}

export const statusType = (status) => {
  switch (status) {
    case 'RELEASED':
      return 'success'
    case 'REVIEWING':
      return 'warning'
    case 'REJECTED':
      return 'danger'
    default:
      return status
  }
}
