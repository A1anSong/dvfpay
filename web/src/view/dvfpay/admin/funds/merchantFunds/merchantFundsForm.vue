<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="商户id:">
          <el-input v-model.number="formData.merchantId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="币种:">
          <el-input v-model="formData.currency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="可用:">
          <el-input v-model.number="formData.available" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="不可用:">
          <el-input v-model.number="formData.unavailable" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MerchantFunds',
}
</script>

<script setup>
import {
  createMerchantFunds,
  updateMerchantFunds,
  findMerchantFunds,
} from '@/api/dvfpay/merchantFunds'

// 自动获取字典
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'

const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
  merchantId: 0,
  currency: '',
  available: 0,
  unavailable: 0,
})

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findMerchantFunds({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.remerchantFunds
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createMerchantFunds(formData.value)
      break
    case 'update':
      res = await updateMerchantFunds(formData.value)
      break
    default:
      res = await createMerchantFunds(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功',
    })
  }
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
