<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>数据总览</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="8">
              <funds-bar-chart />
            </el-col>
            <el-col :xs="24" :sm="8">
              <statistics-income-orders-pie-chart />
            </el-col>
            <el-col :xs="24" :sm="8">
              <statistics-payout-orders-pie-chart />
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>代收订单近30天趋势</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="12">
              <trends-count-income-orders-chart />
            </el-col>
            <el-col :xs="24" :sm="12">
              <trends-sum-income-orders-chart />
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>代付订单近30天趋势</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="12">
              <trends-count-payout-orders-chart />
            </el-col>
            <el-col :xs="24" :sm="12">
              <trends-sum-payout-orders-chart />
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dash',
}
</script>

<script setup>
import FundsBarChart from '@/view/dvfpay/merchant/dash/dashCharts/fundsBarChart.vue'
import StatisticsIncomeOrdersPieChart from '@/view/dvfpay/merchant/dash/dashCharts/statisticsIncomeOrdersPieChart.vue'
import StatisticsPayoutOrdersPieChart from '@/view/dvfpay/merchant/dash/dashCharts/statisticsPayoutOrdersPieChart.vue'
import TrendsCountIncomeOrdersChart from '@/view/dvfpay/merchant/dash/dashCharts/trendsCountIncomeOrdersChart.vue'
import TrendsSumIncomeOrdersChart from '@/view/dvfpay/merchant/dash/dashCharts/trendsSumIncomeOrdersChart.vue'
import TrendsCountPayoutOrdersChart from '@/view/dvfpay/merchant/dash/dashCharts/trendsCountPayoutOrdersChart.vue'
import TrendsSumPayoutOrdersChart from '@/view/dvfpay/merchant/dash/dashCharts/trendsSumPayoutOrdersChart.vue'

</script>

<style lang="scss" scoped>
@mixin flex-center {
  display: flex;
  align-items: center;
}

.page {
  background: #f0f2f5;
  padding: 0;

  .gva-card-box {
    padding: 12px 16px;

    & + .gva-card-box {
      padding-top: 0px;
    }
  }

  .gva-card {
    box-sizing: border-box;
    background-color: #fff;
    border-radius: 2px;
    height: auto;
    padding: 26px 30px;
    overflow: hidden;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
  }

  .gva-top-card {
    height: 260px;
    @include flex-center;
    justify-content: space-between;
    color: #777;

    &-left {
      height: 100%;
      display: flex;
      flex-direction: column;

      &-title {
        font-size: 22px;
        color: #343844;
      }

      &-dot {
        font-size: 14px;
        color: #6B7687;
        margin-top: 24px;
      }

      &-rows {
        // margin-top: 15px;
        margin-top: 18px;
        color: #6B7687;
        width: 600px;
        align-items: center;
      }

      &-item {
        + .gva-top-card-left-item {
          margin-top: 24px;
        }

        margin-top: 14px;
      }
    }

    &-right {
      height: 600px;
      width: 600px;
      margin-top: 28px;
    }
  }

  ::v-deep(.el-card__header) {
    padding: 0;
    border-bottom: none;
  }

  .card-header {
    padding-bottom: 20px;
    border-bottom: 1px solid #e8e8e8;
  }

  .quick-entrance-title {
    height: 30px;
    font-size: 22px;
    color: #333;
    width: 100%;
    border-bottom: 1px solid #eee;
  }

  .quick-entrance-items {
    @include flex-center;
    justify-content: center;
    text-align: center;
    color: #333;

    .quick-entrance-item {
      padding: 16px 28px;
      margin-top: -16px;
      margin-bottom: -16px;
      border-radius: 4px;
      transition: all 0.2s;

      &:hover {
        box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
      }

      cursor: pointer;
      height: auto;
      text-align: center;
      // align-items: center;
      &-icon {
        width: 50px;
        height: 50px !important;
        border-radius: 8px;
        @include flex-center;
        justify-content: center;
        margin: 0 auto;

        i {
          font-size: 24px;
        }
      }

      p {
        margin-top: 10px;
      }
    }
  }

  .echart-box {
    padding: 14px;
  }
}

.dasboard-icon {
  font-size: 20px;
  color: rgb(85, 160, 248);
  width: 30px;
  height: 30px;
  margin-right: 10px;
  @include flex-center;
}

.flex-center {
  @include flex-center;
}

//小屏幕不显示右侧，将登陆框居中
@media (max-width: 750px) {
  .gva-card {
    padding: 20px 10px !important;

    .gva-top-card {
      height: auto;

      &-left {
        &-title {
          font-size: 20px !important;
        }

        &-rows {
          margin-top: 15px;
          align-items: center;
        }
      }

      &-right {
        display: none;
      }
    }

    .gva-middle-card {
      &-item {
        line-height: 20px;
      }
    }

    .dasboard-icon {
      font-size: 18px;
    }
  }
}

</style>
