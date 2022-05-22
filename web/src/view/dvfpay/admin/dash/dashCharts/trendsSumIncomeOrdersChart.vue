<template>
  <div class="dashboard-line-box">
    <div class="dashboard-line-title">
      订单金额
    </div>
    <div
      ref="echart"
      class="dashboard-line"
    />
  </div>
</template>

<script>
export default {
  name: 'TrendsSumIncomeOrdersChart',
}
</script>

<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
import { getTrendsSumIncomeOrder } from '@/api/dvfpay/incomeOrder'

const chart = shallowRef(null)
const echart = ref(null)
const initChart = () => {
  chart.value = echarts.init(echart.value, '', { locale: 'ZH' })
  getData()
}

const usdData = ref([])
const eurData = ref([])
const gbpData = ref([])

const getData = async() => {
  chart.value.showLoading()
  const table = await getTrendsSumIncomeOrder()
  if (table.code === 0) {
    table.data.usdList && table.data.usdList.forEach((sum) => {
      usdData.value.push({
        date: sum.date,
        sum: sum.sum / 100,
      })
    })
    table.data.eurList && table.data.eurList.forEach((sum) => {
      eurData.value.push({
        date: sum.date,
        sum: sum.sum / 100,
      })
    })
    table.data.gbpList && table.data.gbpList.forEach((sum) => {
      gbpData.value.push({
        date: sum.date,
        sum: sum.sum / 100,
      })
    })
    chart.value.setOption({
      grid: {
        left: '50',
        right: '20',
        top: '40',
        bottom: '20',
      },
      legend: {},
      tooltip: {
        trigger: 'axis',
        axisPointer: { type: 'cross' },
      },
      dataset: [{
        source: usdData.value,
      }, {
        source: eurData.value,
      }, {
        source: gbpData.value,
      }],
      xAxis: {
        type: 'time',
        axisTick: {
          show: false,
          alignWithLabel: true,
        },
        axisLabel: {
          formatter: {
            month: '{MMM}',
            day: '{d}日',
          },
        },
        axisLine: {
          show: false,
        },
      },
      yAxis: {},
      series: [{
        name: 'USD',
        datasetIndex: 0,
        type: 'line',
        smooth: true,
      }, {
        name: 'EUR',
        datasetIndex: 1,
        type: 'line',
        smooth: true,
      }, {
        name: 'GBP',
        datasetIndex: 2,
        type: 'line',
        smooth: true,
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
