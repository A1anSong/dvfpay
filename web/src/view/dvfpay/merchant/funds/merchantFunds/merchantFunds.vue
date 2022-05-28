<template>
  <div>
    <div class="gva-table-box">
      <el-row :gutter="20">
        <el-col v-for="item in tableData" :key="item.currency" :xs="24" :sm="6">
          <el-card shadow="hover" class="box-card">
            <template #header>
              <div class="card-header">
                <span>{{ item.currency }}</span>
                <div v-if="item.currency === 'USDT'" style="text-align: right">
                  <el-button class="button" type="primary">充值</el-button>
                  <el-button v-if="item.available > 0" class="button" type="primary">提现</el-button>
                </div>
                <div v-else-if="item.available > 0">
                  <el-button class="button" type="primary" @click="openExchangeDialog(item.currency)">兑换USDT</el-button>
                </div>
              </div>
            </template>
            <div class="text item">可用：{{
              currencySymbols(item.currency) + (item.available / 100).toLocaleString()
            }}
            </div>
            <div class="text item">不可用：{{
              currencySymbols(item.currency) + (item.unavailable / 100).toLocaleString()
            }}
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <el-dialog v-model="exchangeDialogFormVisible" :before-close="closeExchangeDialog" title="兑换USDT">
      <el-form :model="exchangeFormData" label-position="right" label-width="80px">
        <el-form-item label="金额:">
          <el-input-number v-model="exchangeFormData.fromAmount" :precision="2" :step="0.01" style="width: 100%" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeExchangeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterExchangeDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'MerchantFunds',
}
</script>

<script setup>
import {
  getSelfMerchantFundsList,
} from '@/api/dvfpay/merchantFunds'

// 全量引入格式化工具 请按需保留
import { currencySymbols } from '@/utils/dvfpay/currencySymbols'
import { ref } from 'vue'
import { createExchangeRecord } from '@/api/dvfpay/exchangeRecord'
import { ElMessage } from 'element-plus'

const exchangeFormData = ref({
  from: '',
  fromAmount: 0,
})

// 自动化生成的字典（可能为空）以及字段

// =========== 表格控制部分 ===========
const tableData = ref([])

// 查询
const getTableData = async() => {
  const table = await getSelfMerchantFundsList({ page: 1, pageSize: 999 })
  if (table.code === 0) {
    tableData.value = table.data.list
  }
}

getTableData()

const exchangeDialogFormVisible = ref(false)

const openExchangeDialog = (from) => {
  exchangeFormData.value.from = from
  exchangeDialogFormVisible.value = true
}

const closeExchangeDialog = () => {
  exchangeDialogFormVisible.value = false
  exchangeFormData.value = {
    from: '',
    fromAmount: 0,
  }
}

const enterExchangeDialog = async() => {
  const exchangeData = {
    from: exchangeFormData.value.from,
    to: 'USDT',
    fromAmount: exchangeFormData.value.fromAmount,
  }
  exchangeData.fromAmount = exchangeData.fromAmount * 100
  const res = await createExchangeRecord(exchangeData)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功',
    })
    closeExchangeDialog()
    getTableData()
  }
}
</script>

<style>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 20px;
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
