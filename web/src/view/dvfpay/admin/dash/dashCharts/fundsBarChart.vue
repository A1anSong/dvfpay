<template>
  <div class="dashboard-line-box">
    <div class="dashboard-line-title">
      钱包余额
    </div>
    <div
      ref="echart"
      class="dashboard-line"
    />
  </div>
</template>

<script>
export default {
  name: 'FundsBarChart',
}
</script>

<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
import { getStatisticsMerchantFundsList } from '@/api/dvfpay/merchantFunds'

const chart = shallowRef(null)
const echart = ref(null)
const initChart = () => {
  chart.value = echarts.init(echart.value, '', { locale: 'ZH' })
  getData()
}

const currencyData = ref([])

const getData = async() => {
  chart.value.showLoading()
  const table = await getStatisticsMerchantFundsList({})
  if (table.code === 0) {
    table.data.list && table.data.list.forEach((funds) => {
      if (funds.available !== 0 || funds.unavailable !== 0) {
        currencyData.value.push({
          currency: funds.currency,
          '可用': funds.available / 100,
          '不可用': funds.unavailable / 100,
        })
      }
    })
    chart.value.setOption({
      grid: {
        left: '50',
        right: '20',
        top: '40',
        bottom: '20',
      },
      legend: {},
      tooltip: {},
      dataset: {
        source: currencyData.value,
      },
      xAxis: {
        type: 'category',
        axisTick: {
          show: false,
        },
        axisLine: {
          show: false,
        },
      },
      yAxis: {},
      series: [{
        type: 'bar',
        name: '可用',
        itemStyle: {
          borderRadius: [5, 5, 0, 0],
          color: '#91cc75',
        },
      }, {
        type: 'bar',
        name: '不可用',
        itemStyle: {
          borderRadius: [5, 5, 0, 0],
          color: '#ee6666',
        },
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
