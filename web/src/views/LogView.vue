<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import request from '@/utils/request'

interface LogItem {
  id: number
  operatorId: number
  operatorName: string
  module: string
  action: string
  target: string
  detail: string
  ip: string
  createdAt: string
}

const list = ref<LogItem[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const loading = ref(false)

const moduleFilter = ref('')
const actionFilter = ref('')
const operatorIdFilter = ref('')

const moduleOptions = [
  { label: '全部模块', value: '' },
  { label: '用户管理', value: '用户管理' },
  { label: '证件管理', value: '证件管理' },
  { label: '证件类型', value: '证件类型' },
  { label: '导出任务', value: '导出任务' },
  { label: '导入任务', value: '导入任务' },
  { label: '文件上传', value: '文件上传' },
  { label: '系统', value: '系统' },
]

const actionOptions = [
  { label: '全部动作', value: '' },
  { label: '新增/执行', value: '新增/执行' },
  { label: '查询', value: '查询' },
  { label: '更新', value: '更新' },
  { label: '删除', value: '删除' },
]

async function fetchList() {
  loading.value = true
  try {
    const params: any = { page: page.value, pageSize: pageSize.value }
    if (moduleFilter.value) params.module = moduleFilter.value
    if (actionFilter.value) params.action = actionFilter.value
    if (operatorIdFilter.value) params.operatorId = operatorIdFilter.value

    const res: any = await request.get('/api/operation-log/list', { params })
    list.value = res.data.list || []
    total.value = res.data.total || 0
  } catch {
    ElMessage.error('获取操作日志失败')
  } finally {
    loading.value = false
  }
}

function onSearch() {
  page.value = 1
  fetchList()
}

function onReset() {
  moduleFilter.value = ''
  actionFilter.value = ''
  operatorIdFilter.value = ''
  page.value = 1
  fetchList()
}

onMounted(() => {
  fetchList()
})
</script>

<template>
  <div>
    <el-card class="card">
      <template #header>
        <span>操作日志</span>
      </template>

      <el-form inline class="search-form">
        <el-form-item label="模块">
          <el-select v-model="moduleFilter" placeholder="选择模块" clearable style="width: 160px">
            <el-option
              v-for="opt in moduleOptions"
              :key="opt.value"
              :label="opt.label"
              :value="opt.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="动作">
          <el-select v-model="actionFilter" placeholder="选择动作" clearable style="width: 160px">
            <el-option
              v-for="opt in actionOptions"
              :key="opt.value"
              :label="opt.label"
              :value="opt.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="操作人ID">
          <el-input v-model="operatorIdFilter" placeholder="操作人ID" clearable style="width: 140px" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSearch">查询</el-button>
          <el-button @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>

      <el-table :data="list" v-loading="loading" stripe size="small">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="operatorName" label="操作人" width="120" />
        <el-table-column prop="module" label="模块" width="120">
          <template #default="{ row }">
            {{ moduleOptions.find(o => o.value === row.module)?.label || row.module }}
          </template>
        </el-table-column>
        <el-table-column prop="action" label="动作" width="120">
          <template #default="{ row }">
            {{ actionOptions.find(o => o.value === row.action)?.label || row.action }}
          </template>
        </el-table-column>
        <el-table-column prop="target" label="操作对象" min-width="160" show-overflow-tooltip />
        <el-table-column prop="ip" label="IP" width="130" />
        <el-table-column prop="createdAt" label="操作时间" width="170" />
      </el-table>

      <el-pagination
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :total="total"
        layout="total, prev, pager, next"
        @change="fetchList"
        style="margin-top: 16px"
      />
    </el-card>
  </div>
</template>

<style scoped>
.card {
  margin-bottom: 16px;
}
.search-form {
  margin-bottom: 16px;
}
</style>
