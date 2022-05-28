<template>
  <div class="dashboard-line-box">
    <div class="dashboard-line-title">
      代付订单统计
    </div>
    <div
      ref="echart"
      class="dashboard-line"
    />
  </div>
</template>

<script>
export default {
  name: 'StatisticsPayoutOrdersPieChart',
}
</script>

<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
import { getMerchantStatisticsPayoutOrder } from '@/api/dvfpay/payoutOrder'

const chart = shallowRef(null)
const echart = ref(null)
const initChart = () => {
  chart.value = echarts.init(echart.value, '', { locale: 'ZH' })
  getData()
}

const statisticsIncomeOrderData = ref([])

const getData = async() => {
  chart.value.showLoading()
  const table = await getMerchantStatisticsPayoutOrder()
  if (table.code === 0) {
    table.data.list && table.data.list.forEach((count) => {
      if (count.name === '成功') {
        count.itemStyle = {
          color: '#91cc75',
        }
      }
      if (count.name === '待定') {
        count.itemStyle = {
          color: '#fac858',
        }
      }
      if (count.name === '结算') {
        count.itemStyle = {
          color: '#5470c6',
        }
      }
      statisticsIncomeOrderData.value.push(count)
    })
    chart.value.setOption({
      grid: {
        left: '40',
        right: '20',
        top: '40',
        bottom: '20',
      },
      legend: {},
      tooltip: {},
      series: [{
        type: 'pie',
        top: '10%',
        data: statisticsIncomeOrderData.value,
      }],
    })
    chart.value.hideLoading()
  }
}

onMounted(async() => {
  await nextTick()
  initChart()
  window.addEventListener('resize', () => {
    chart.value.resize()
  })
})

onUnmounted(() => {
  if (!chart.value) {
    return
  }
  chart.value.dispose()
  chart.value = null
})
</script>

<style lang="scss" scoped>
.dashboard-line-box {
  .dashboard-line {
    background-color: #fff;
    height: 240px;
    width: 100%;
  }

  .dashboard-line-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
}
</style>
