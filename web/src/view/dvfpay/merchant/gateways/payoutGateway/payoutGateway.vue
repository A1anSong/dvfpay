<template>
  <div>
    <div class="gva-table-box">
      <el-row :gutter="20">
        <el-col v-for="item in tableData" :key="item.ID" :xs="24" :sm="6">
          <el-card shadow="hover" class="box-card">
            <template #header>
              <div class="card-header">
                <span>{{ item.payoutGateway.name }}</span>
              </div>
            </template>
            <div class="text item">币种：{{ item.payoutGateway.currency }}</div>
            <div class="text item">状态：{{ item.payoutGateway.status }}</div>
            <div class="text item">手续费：{{ (item.fee / 100).toFixed(2) + '%' }}</div>
            <div class="text item">单笔最高：{{ item.limitMax.toLocaleString() }}</div>
            <div class="text item">单笔最低：{{ item.limitMin.toLocaleString() }}</div>
            <div class="text item">单日限制：{{ item.limitDay.toLocaleString() }}</div>
            <div class="text item">总量限制：{{ item.limitTotal.toLocaleString() }}</div>
            <div v-if="item.payoutGateway.type === 'gateway'" class="text item">
              接口地址：{{ item.payoutGateway.gateWayURL +'?id='+ item.ID }}
            </div>
            <div class="text item">说明：<br><span v-html="item.explain" /></div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PayoutGatewayAuth',
}
</script>

<script setup>
import {
  getMerchantPayoutGatewayAuthList,
} from '@/api/dvfpay/payoutGatewayAuth'

// 全量引入格式化工具 请按需保留
import { ref } from 'vue'

// =========== 表格控制部分 ===========
const tableData = ref([])

// 查询
const getTableData = async() => {
  const table = await getMerchantPayoutGatewayAuthList({ page: 1, pageSize: 999 })
  if (table.code === 0) {
    tableData.value = table.data.list
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()
</script>

<style>
</style>
