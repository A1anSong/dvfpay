<template>
  <div class="dashboard-line-box">
    <div class="dashboard-line-title">
      订单数量
    </div>
    <div
      ref="echart"
      class="dashboard-line"
    />
  </div>
</template>

<script>
export default {
  name: 'TrendsCountIncomeOrdersChart',
}
</script>

<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
import { getTrendsCountIncomeOrder } from '@/api/dvfpay/incomeOrder'

const chart = shallowRef(null)
const echart = ref(null)
const initChart = () => {
  chart.value = echarts.init(echart.value, '', { locale: 'ZH' })
  getData()
}

const countData = ref([])
const countSeries = ref([])

const getData = async() => {
  chart.value.showLoading()
  const table = await getTrendsCountIncomeOrder()
  if (table.code === 0) {
    if (table.data.list !== undefined) {
      let index = 0
      for (const key in table.data.list) {
        const data = []
        table.data.list[key].forEach((count) => {
          data.push({
            date: count.date,
            count: count.count,
          })
        })
        countData.value.push({
          source: data,
        })
        countSeries.value.push({
          name: key,
          datasetIndex: index,
          type: 'line',
          smooth: true,
        })
        index++
      }
    }
    chart.value.setOption({
      grid: {
        left: '40',
        right: '20',
        top: '40',
        bottom: '20',
      },
      legend: {},
      tooltip: {
        trigger: 'axis',
        axisPointer: { type: 'cross' },
      },
      dataset: countData.value,
      xAxis: {
        type: 'time',
        axisTick: {
          show: false,
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
      series: countSeries.value,
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
