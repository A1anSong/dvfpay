export const currencySymbols = (currency) => {
  switch (currency) {
    case 'USD':
      return '$'
    case 'USDT':
      return '₮'
    case 'EUR':
      return '€'
    case 'GBP':
      return '£'
    default:
      return currency
  }
}
