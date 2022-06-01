<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="订单id">
          <el-input v-model="searchInfo.orderId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="搜索条件">
            <el-option
              key="success"
              label="成功"
              value="success"
            />
            <el-option
              key="PENDING"
              label="待定"
              value="PENDING"
            />
            <el-option
              key="SETTLED"
              label="结算"
              value="SETTLED"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="订单信息">
          <el-input v-model="searchInfo.metaData" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="日期">
          <el-date-picker v-model="searchInfo.arrivalTime" type="date" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <!--            <div class="gva-btn-list">-->
      <!--              <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>-->
      <!--              <el-popover v-model:visible="deleteVisible" placement="top" width="160">-->
      <!--                <p>确定要删除吗？</p>-->
      <!--                <div style="text-align: right; margin-top: 8px;">-->
      <!--                  <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>-->
      <!--                  <el-button size="small" type="primary" @click="onDelete">确定</el-button>-->
      <!--                </div>-->
      <!--                <template #reference>-->
      <!--                  <el-button-->
      <!--                    icon="delete"-->
      <!--                    size="small"-->
      <!--                    style="margin-left: 10px;"-->
      <!--                    :disabled="!multipleSelection.length"-->
      <!--                    @click="deleteVisible = true"-->
      <!--                  >删除-->
      <!--                  </el-button>-->
      <!--                </template>-->
      <!--              </el-popover>-->
      <!--            </div>-->
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <!--        <el-table-column type="selection" width="55" />-->
        <!--        <el-table-column align="left" label="日期" width="180">-->
        <!--          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
        <!--        </el-table-column>-->
        <el-table-column align="left" label="到账时间" min-width="180">
          <template #default="scope">{{ formatDate(scope.row.arrivalTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="商户订单号" prop="merchantOrderId" min-width="120" />
        <el-table-column align="left" label="订单id" prop="orderId" min-width="120" />
        <el-table-column align="left" label="金额" min-width="120">
          <template #default="scope">{{
            currencySymbols(scope.row.currency) + (scope.row.amount / 100).toFixed(2)
          }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="币种" prop="currency" min-width="120" />
        <el-table-column align="left" label="状态" min-width="120">
          <template #default="scope">
            <el-tag
              :type="statusType(scope.row.status)"
              effect="dark"
            >
              {{ statusDisplay(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="付款人" prop="payer" min-width="120" />
        <el-table-column align="left" label="备注" prop="remark" min-width="120" />
        <el-table-column align="left" label="商户名称" prop="merchant.nickName" min-width="120" />
        <el-table-column align="left" label="通道名称" prop="incomeGatewayAuth.incomeGateway.name" min-width="120" />
        <el-table-column align="left" label="确认上分" min-width="120">
          <template #default="scope">{{ scope.row.confirmed ? '是' : '否' }}</template>
        </el-table-column>
        <!--        <el-table-column align="left" label="按钮组">-->
        <!--          <template #default="scope">-->
        <!--            <el-button-->
        <!--                type="text"-->
        <!--                icon="edit"-->
        <!--                size="small"-->
        <!--                class="table-button"-->
        <!--                @click="updateIncomeOrderFunc(scope.row)"-->
        <!--            >变更-->
        <!--            </el-button>-->
        <!--            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>-->
        <!--          </template>-->
        <!--        </el-table-column>-->
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <!--    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">-->
    <!--      <el-form :model="formData" label-position="right" label-width="80px">-->
    <!--        <el-form-item label="到账时间:">-->
    <!--          <el-date-picker v-model="formData.arrivalTime" type="date" style="width:100%" placeholder="选择日期" clearable/>-->
    <!--        </el-form-item>-->
    <!--        <el-form-item label="订单id:">-->
    <!--          <el-input v-model="formData.orderId" clearable placeholder="请输入"/>-->
    <!--        </el-form-item>-->
    <!--        <el-form-item label="金额:">-->
    <!--          <el-input v-model.number="formData.amount" clearable placeholder="请输入"/>-->
    <!--        </el-form-item>-->
    <!--        <el-form-item label="币种:">-->
    <!--          <el-input v-model="formData.currency" clearable placeholder="请输入"/>-->
    <!--        </el-form-item>-->
    <!--        <el-form-item label="状态:">-->
    <!--          <el-input v-model="formData.status" clearable placeholder="请输入"/>-->
    <!--        </el-form-item>-->
    <!--        <el-form-item label="付款人:">-->
    <!--          <el-input v-model="formData.payer" clearable placeholder="请输入"/>-->
    <!--        </el-form-item>-->
    <!--        <el-form-item label="备注:">-->
    <!--          <el-input v-model="formData.remark" clearable placeholder="请输入"/>-->
    <!--        </el-form-item>-->
    <!--        <el-form-item label="确认上分:">-->
    <!--          <el-select v-model="formData.confirmed" placeholder="请输入" style="width: 100%">-->
    <!--            <el-option-->
    <!--                :key="true"-->
    <!--                label="已确认"-->
    <!--                :value="true"-->
    <!--            />-->
    <!--            <el-option-->
    <!--                :key="false"-->
    <!--                label="未确认"-->
    <!--                :value="false"-->
    <!--            />-->
    <!--          </el-select>-->
    <!--        </el-form-item>-->
    <!--      </el-form>-->
    <!--      <template #footer>-->
    <!--        <div class="dialog-footer">-->
    <!--          <el-button size="small" @click="closeDialog">取 消</el-button>-->
    <!--          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>-->
    <!--        </div>-->
    <!--      </template>-->
    <!--    </el-dialog>-->
  </div>
</template>

<script>
export default {
  name: 'IncomeOrder',
}
</script>

<script setup>
import {
  // createIncomeOrder,
  // deleteIncomeOrder,
  // deleteIncomeOrderByIds,
  // updateIncomeOrder,
  // findIncomeOrder,
  getIncomeOrderList,
} from '@/api/dvfpay/incomeOrder'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { currencySymbols } from '@/utils/dvfpay/currency'
import { statusDisplay, statusType } from '@/utils/dvfpay/order'
// import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
// const formData = ref({
//   arrivalTime: new Date(),
//   orderId: '',
//   amount: '',
//   currency: '',
//   status: '',
//   payer: '',
//   remark: '',
//   confirmed: '',
// })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getIncomeOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
// const deleteRow = (row) => {
//   ElMessageBox.confirm('确定要删除吗?', '提示', {
//     confirmButtonText: '确定',
//     cancelButtonText: '取消',
//     type: 'warning',
//   }).then(() => {
//     deleteIncomeOrderFunc(row)
//   })
// }

// 批量删除控制标记
// const deleteVisible = ref(false)

// 多选删除
// const onDelete = async() => {
//   const ids = []
//   if (multipleSelection.value.length === 0) {
//     ElMessage({
//       type: 'warning',
//       message: '请选择要删除的数据',
//     })
//     return
//   }
//   multipleSelection.value &&
//   multipleSelection.value.map(item => {
//     ids.push(item.ID)
//   })
//   const res = await deleteIncomeOrderByIds({ ids })
//   if (res.code === 0) {
//     ElMessage({
//       type: 'success',
//       message: '删除成功',
//     })
//     if (tableData.value.length === ids.length && page.value > 1) {
//       page.value--
//     }
//     deleteVisible.value = false
//     getTableData()
//   }
// }

// 行为控制标记（弹窗内部需要增还是改）
// const type = ref('')

// 更新行
// const updateIncomeOrderFunc = async(row) => {
//   const res = await findIncomeOrder({ ID: row.ID })
//   type.value = 'update'
//   if (res.code === 0) {
//     formData.value = res.data.reincomeOrder
//     dialogFormVisible.value = true
//   }
// }

// 删除行
// const deleteIncomeOrderFunc = async(row) => {
//   const res = await deleteIncomeOrder({ ID: row.ID })
//   if (res.code === 0) {
//     ElMessage({
//       type: 'success',
//       message: '删除成功',
//     })
//     if (tableData.value.length === 1 && page.value > 1) {
//       page.value--
//     }
//     getTableData()
//   }
// }

// 弹窗控制标记
// const dialogFormVisible = ref(false)

// 打开弹窗
// const openDialog = () => {
//   type.value = 'create'
//   dialogFormVisible.value = true
// }

// 关闭弹窗
// const closeDialog = () => {
//   dialogFormVisible.value = false
//   formData.value = {
//     arrivalTime: new Date(),
//     orderId: '',
//     amount: '',
//     currency: '',
//     status: '',
//     payer: '',
//     remark: '',
//     confirmed: '',
//   }
// }
// 弹窗确定
// const enterDialog = async() => {
//   let res
//   switch (type.value) {
//     case 'create':
//       res = await createIncomeOrder(formData.value)
//       break
//     case 'update':
//       res = await updateIncomeOrder(formData.value)
//       break
//     default:
//       res = await createIncomeOrder(formData.value)
//       break
//   }
//   if (res.code === 0) {
//     ElMessage({
//       type: 'success',
//       message: '创建/更改成功',
//     })
//     closeDialog()
//     getTableData()
//   }
// }
</script>

<style>
</style>
