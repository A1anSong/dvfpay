<template>
  <div>
    <div class="gva-table-box">
      <el-row :gutter="20">
        <el-col v-for="item in tableData" :key="item.merchantId" :xs="24" :sm="24">
          <el-card shadow="always" class="box-card">
            <div class="text item">MerchantId：{{ item.merchantId }}</div>
            <div class="text item">SecretKey：{{ item.secretKey }}</div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ApiKey',
}
</script>

<script setup>
// 全量引入格式化工具 请按需保留
import { ref } from 'vue'
import { getAPIKey } from '@/api/user'

// =========== 表格控制部分 ===========
const tableData = ref([])

// 查询
const getTableData = async() => {
  const table = await getAPIKey()
  if (table.code === 0) {
    tableData.value = table.data.list
  }
}

getTableData()
</script>

<style scoped>
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
