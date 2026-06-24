<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import request from '@/utils/request'

interface Category {
  id: number
  name: string
  code: string
  description: string
  status: number
}

const list = ref<Category[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const keyword = ref('')
const loading = ref(false)

const dialogVisible = ref(false)
const isEdit = ref(false)
const form = ref<Partial<Category>>({
  name: '',
  code: '',
  description: '',
  status: 1,
})

async function fetchList() {
  loading.value = true
  try {
    const res: any = await request.get('/api/category/list', {
      params: { page: page.value, pageSize: pageSize.value, keyword: keyword.value },
    })
    list.value = res.data.list || []
    total.value = res.data.total || 0
  } finally {
    loading.value = false
  }
}

function onSearch() {
  page.value = 1
  fetchList()
}

function onAdd() {
  isEdit.value = false
  form.value = { name: '', code: '', description: '', status: 1 }
  dialogVisible.value = true
}

function onEdit(row: Category) {
  isEdit.value = true
  form.value = { ...row }
  dialogVisible.value = true
}

async function onDelete(row: Category) {
  try {
    await ElMessageBox.confirm('确认删除该证件类型？', '提示', { type: 'warning' })
    await request.post('/api/category/delete', { id: row.id })
    ElMessage.success('删除成功')
    fetchList()
  } catch {
    // cancelled
  }
}

async function onSubmit() {
  if (!form.value.name || !form.value.code) {
    ElMessage.warning('请填写名称和编码')
    return
  }
  try {
    const url = isEdit.value ? '/api/category/update' : '/api/category/add'
    await request.post(url, form.value)
    ElMessage.success(isEdit.value ? '更新成功' : '添加成功')
    dialogVisible.value = false
    fetchList()
  } catch {
    // handled by interceptor
  }
}

onMounted(fetchList)
</script>

<template>
  <div>
    <div class="toolbar">
      <el-input v-model="keyword" placeholder="搜索名称/编码" style="width: 240px; margin-right: 12px" @keyup.enter="onSearch" />
      <el-button type="primary" @click="onSearch">搜索</el-button>
      <el-button type="success" @click="onAdd">新增类型</el-button>
    </div>
    <el-table :data="list" v-loading="loading" stripe>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="name" label="名称" />
      <el-table-column prop="code" label="编码" />
      <el-table-column prop="description" label="描述" show-overflow-tooltip />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="row.status === 1 ? 'success' : 'danger'">{{ row.status === 1 ? '正常' : '禁用' }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <el-button size="small" @click="onEdit(row)">编辑</el-button>
          <el-button size="small" type="danger" @click="onDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      v-model:current-page="page"
      v-model:page-size="pageSize"
      :total="total"
      layout="total, prev, pager, next"
      @change="fetchList"
      style="margin-top: 16px"
    />

    <el-dialog v-model="dialogVisible" :title="isEdit ? '编辑证件类型' : '新增证件类型'" width="500px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="名称" required>
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item label="编码" required>
          <el-input v-model="form.code" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="form.description" type="textarea" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="form.status">
            <el-option label="正常" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="onSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
.toolbar {
  margin-bottom: 16px;
}
</style>
