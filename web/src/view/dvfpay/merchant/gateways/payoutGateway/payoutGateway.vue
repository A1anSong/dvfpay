<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="代收通道id">
          <el-input v-model="searchInfo.payoutGatewayId" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item label="商户id">
          <el-input v-model="searchInfo.merchantId" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
<!--      <div class="gva-btn-list">-->
<!--        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>-->
<!--        <el-popover v-model:visible="deleteVisible" placement="top" width="160">-->
<!--          <p>确定要删除吗？</p>-->
<!--          <div style="text-align: right; margin-top: 8px;">-->
<!--            <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>-->
<!--            <el-button size="small" type="primary" @click="onDelete">确定</el-button>-->
<!--          </div>-->
<!--          <template #reference>-->
<!--            <el-button-->
<!--                icon="delete"-->
<!--                size="small"-->
<!--                style="margin-left: 10px;"-->
<!--                :disabled="!multipleSelection.length"-->
<!--                @click="deleteVisible = true"-->
<!--            >删除-->
<!--            </el-button>-->
<!--          </template>-->
<!--        </el-popover>-->
<!--      </div>-->
      <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
      >
<!--        <el-table-column type="selection" width="55"/>-->
<!--        <el-table-column align="left" label="日期" width="180">-->
<!--          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
<!--        </el-table-column>-->
        <el-table-column align="left" label="代付通道" prop="payoutGateway.name" width="120"/>
<!--        <el-table-column align="left" label="商户" prop="merchantId" min-width="200">-->
<!--          <template #default="scope">-->
<!--            <el-cascader-->
<!--                v-model="scope.row.merchantIds"-->
<!--                :options="merchantOptions"-->
<!--                collapse-tags-->
<!--                :props="{ multiple:true,label:'nickName',value:'ID' }"-->
<!--            />-->
<!--          </template>-->
<!--        </el-table-column>-->
        <el-table-column align="left" label="手续费" prop="fee" width="120"/>
        <el-table-column align="left" label="单笔最高" prop="limitMax" width="120"/>
        <el-table-column align="left" label="单笔最低" prop="limitMin" width="120"/>
        <el-table-column align="left" label="单日限制" prop="limitDay" width="120"/>
        <el-table-column align="left" label="总量限制" prop="limitTotal"/>
<!--        <el-table-column align="left" label="按钮组">-->
<!--          <template #default="scope">-->
<!--            <el-button-->
<!--                type="text"-->
<!--                icon="edit"-->
<!--                size="small"-->
<!--                class="table-button"-->
<!--                @click="updatePayoutGatewayAuthFunc(scope.row)"-->
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
<!--        <el-form-item label="代付通道:">-->
<!--          <el-cascader-->
<!--              v-model="formData.payoutGatewayId"-->
<!--              style="width:100%"-->
<!--              :options="payoutGatewayOptions"-->
<!--              :props="{ label:'name',value:'ID' }"-->
<!--              filterable-->
<!--          />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="授权商户:">-->
<!--          <el-cascader-->
<!--              v-model="formData.merchants"-->
<!--              style="width:100%"-->
<!--              :options="merchantOptions"-->
<!--              :props="{ multiple:true,label:'nickName',value:'ID' }"-->
<!--              filterable-->
<!--          />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="手续费:">-->
<!--          <el-input v-model.number="formData.fee" clearable placeholder="请输入"/>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="单笔最高:">-->
<!--          <el-input v-model.number="formData.limitMax" clearable placeholder="请输入"/>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="单笔最低:">-->
<!--          <el-input v-model.number="formData.limitMin" clearable placeholder="请输入"/>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="单日限制:">-->
<!--          <el-input v-model.number="formData.limitDay" clearable placeholder="请输入"/>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="总量限制:">-->
<!--          <el-input v-model.number="formData.limitTotal" clearable placeholder="请输入"/>-->
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
  name: 'PayoutGatewayAuth',
}
</script>

<script setup>
import {
  // createPayoutGatewayAuth,
  // deletePayoutGatewayAuth,
  // deletePayoutGatewayAuthByIds,
  // updatePayoutGatewayAuth,
  // findPayoutGatewayAuth,
  getMerchantPayoutGatewayAuthList,
} from '@/api/dvfpay/payoutGatewayAuth'

// 全量引入格式化工具 请按需保留
// import { formatDate } from '@/utils/format'
// import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
// const formData = ref({
//   payoutGatewayId: '',
//   merchantId: [],
//   fee: '',
//   limitMax: '',
//   limitMin: '',
//   limitDay: '',
//   limitTotal: '',
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
  const table = await getMerchantPayoutGatewayAuthList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
//     deletePayoutGatewayAuthFunc(row)
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
//   const res = await deletePayoutGatewayAuthByIds({ ids })
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
// const updatePayoutGatewayAuthFunc = async(row) => {
//   const res = await findPayoutGatewayAuth({ ID: row.ID })
//   type.value = 'update'
//   if (res.code === 0) {
//     formData.value = res.data.repayoutGatewayAuth
//     const merchantIds = formData.value.merchants && formData.value.merchants.map(i => {
//       return [i.ID]
//     })
//     formData.value.merchants = merchantIds
//     dialogFormVisible.value = true
//   }
// }

// 删除行
// const deletePayoutGatewayAuthFunc = async(row) => {
//   const res = await deletePayoutGatewayAuth({ ID: row.ID })
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
//     payoutGatewayId: '',
//     merchantId: [],
//     fee: '',
//     limitMax: '',
//     limitMin: '',
//     limitDay: '',
//     limitTotal: '',
//   }
// }
// 弹窗确定
// const enterDialog = async() => {
//   formData.value.payoutGatewayId = formData.value.payoutGatewayId[0]
//   formData.value.merchants.forEach((item, index, arr) => {
//     arr[index] = arr[index][0]
//   })
//   let res
//   switch (type.value) {
//     case 'create':
//       res = await createPayoutGatewayAuth(formData.value)
//       break
//     case 'update':
//       res = await updatePayoutGatewayAuth(formData.value)
//       break
//     default:
//       res = await createPayoutGatewayAuth(formData.value)
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
