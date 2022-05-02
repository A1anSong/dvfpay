<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="到账时间:">
          <el-date-picker v-model="formData.arrivalTime" type="date" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="订单id:">
          <el-input v-model="formData.orderId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="金额:">
          <el-input v-model.number="formData.amount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="币种:">
          <el-input v-model="formData.currency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-input v-model="formData.status" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付款人:">
          <el-input v-model="formData.payer" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
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
  name: 'PayoutOrder',
}
</script>

<script setup>
import {
  createPayoutOrder,
  updatePayoutOrder,
  findPayoutOrder,
} from '@/api/dvfpay/payoutOrder'

// 自动获取字典
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'

const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
  arrivalTime: new Date(),
  orderId: '',
  amount: 0,
  currency: '',
  status: '',
  payer: '',
  remark: '',
})

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findPayoutOrder({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.repayoutOrder
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
      res = await createPayoutOrder(formData.value)
      break
    case 'update':
      res = await updatePayoutOrder(formData.value)
      break
    default:
      res = await createPayoutOrder(formData.value)
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
