<template>
  <div>
    <div class="gva-table-box">
      <el-row :gutter="20">
        <el-col v-for="item in tableData" :key="item.ID" :xs="24" :sm="6">
          <el-card shadow="hover" class="box-card">
            <template #header>
              <div class="card-header">
                <span>{{ item.incomeGateway.name }}</span>
              </div>
            </template>
            <div class="text item">币种：{{ item.incomeGateway.currency }}</div>
            <div class="text item">状态：{{ item.incomeGateway.status }}</div>
            <div class="text item">手续费：{{ (item.fee / 100).toFixed(2) + '%' }}</div>
            <div class="text item">单笔最高：{{ item.limitMax.toLocaleString() }}</div>
            <div class="text item">单笔最低：{{ item.limitMin.toLocaleString() }}</div>
            <div class="text item">单日限制：{{ item.limitDay.toLocaleString() }}</div>
            <div class="text item">总量限制：{{ item.limitTotal.toLocaleString() }}</div>
            <div v-if="item.incomeGateway.type === 'gateway'" class="text item">
              接口地址：{{ item.incomeGateway.gateWayURL +'?id='+ item.ID }}
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
  name: 'IncomeGatewayAuth',
}
</script>

<script setup>
import {
  getMerchantIncomeGatewayAuthList,
} from '@/api/dvfpay/incomeGatewayAuth'

// 全量引入格式化工具 请按需保留
import { ref } from 'vue'

// =========== 表格控制部分 ===========
const tableData = ref([])

// 查询
const getTableData = async() => {
  const table = await getMerchantIncomeGatewayAuthList({ page: 1, pageSize: 999 })
  if (table.code === 0) {
    tableData.value = table.data.list
  }
}

getTableData()
</script>

<style>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  margin-bottom: 20px;
}
</style>
