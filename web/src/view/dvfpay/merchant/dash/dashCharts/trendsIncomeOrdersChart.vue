<template>
  <div class="dashboard-line-box">
    <div class="dashboard-line-title">
      近30天代收订单趋势
    </div>
    <div
      ref="echart"
      class="dashboard-line"
    />
  </div>
</template>

<script>
export default {
  name: 'TrendsIncomeOrdersChart',
}
</script>

<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
import { getTrendsMerchantIncomeOrder } from '@/api/dvfpay/incomeOrder'

const chart = shallowRef(null)
const echart = ref(null)
const initChart = () => {
  chart.value = echarts.init(echart.value, '', { locale: 'ZH' })
  getData()
}

const sumData = ref([])
const countData = ref([])

const getData = async() => {
  chart.value.showLoading()
  const table = await getTrendsMerchantIncomeOrder()
  if (table.code === 0) {
    table.data.sumList && table.data.sumList.forEach((sum) => {
      sumData.value.push({
        date: sum.date,
        sum: sum.sum / 100,
      })
    })
    table.data.countList && table.data.countList.forEach((count) => {
      countData.value.push({
        date: count.date,
        count: count.count,
      })
    })
    chart.value.setOption({
      grid: {
        left: '50',
        right: '30',
        top: '40',
        bottom: '20',
      },
      legend: {},
      tooltip: {
        trigger: 'axis',
        axisPointer: { type: 'cross' },
      },
      dataset: [{
        source: sumData.value,
      }, {
        source: countData.value,
      }],
      xAxis: {
        type: 'time',
        maxInterval: 3600 * 24 * 1000,
        axisTick: {
          show: false,
          alignWithLabel: true,
        },
        axisLabel: {
          formatter: {
            month: '{MMMM}',
            day: '{d}日',
          },
        },
        axisLine: {
          show: false,
        },
      },
      yAxis: [{
        position: 'left',
      }, {
        position: 'right',
      }],
      series: [{
        name: '收单金额',
        datasetIndex: 0,
        yAxisIndex: 0,
        type: 'bar',
        itemStyle: {
          borderRadius: [5, 5, 0, 0],
        },
      }, {
        name: '订单数量',
        datasetIndex: 1,
        yAxisIndex: 1,
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
    height: 360px;
    width: 100%;
  }

  .dashboard-line-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
}
</style>
