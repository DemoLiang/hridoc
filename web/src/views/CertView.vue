<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import request from '@/utils/request'

interface Cert {
  id: number
  userId: number
  userName: string
  categoryId: number
  categoryName: string
  name: string
  certNo: string
  issuer: string
  issueDate: string
  expireDate: string
  level: string
  fileUrl: string
  fileType: string
  thumbUrl: string
  status: number
}

interface User {
  id: number
  name: string
}

interface Category {
  id: number
  name: string
}

const list = ref<Cert[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const keyword = ref('')
const loading = ref(false)

const users = ref<User[]>([])
const categories = ref<Category[]>([])

const dialogVisible = ref(false)
const isEdit = ref(false)
const form = ref<Partial<Cert>>({
  userId: undefined,
  categoryId: undefined,
  name: '',
  certNo: '',
  issuer: '',
  issueDate: '',
  expireDate: '',
  level: '',
  fileUrl: '',
  fileType: '',
  status: 1,
})

const selectedCerts = ref<Cert[]>([])
const previewVisible = ref(false)
const previewUrl = ref('')
const isPdf = ref(false)

async function fetchList() {
  loading.value = true
  try {
    const res: any = await request.get('/api/cert/list', {
      params: { page: page.value, pageSize: pageSize.value, keyword: keyword.value },
    })
    list.value = res.data.list || []
    total.value = res.data.total || 0
  } finally {
    loading.value = false
  }
}

async function fetchUsersAndCategories() {
  const [uRes, cRes]: any = await Promise.all([
    request.get('/api/user/list', { params: { page: 1, pageSize: 1000 } }),
    request.get('/api/category/list', { params: { page: 1, pageSize: 1000 } }),
  ])
  users.value = uRes.data.list || []
  categories.value = cRes.data.list || []
}

function onSearch() {
  page.value = 1
  fetchList()
}

function onAdd() {
  isEdit.value = false
  form.value = { name: '', certNo: '', issuer: '', issueDate: '', expireDate: '', level: '', fileUrl: '', fileType: '', status: 1 }
  fetchUsersAndCategories()
  dialogVisible.value = true
}

function onEdit(row: Cert) {
  isEdit.value = true
  form.value = { ...row }
  fetchUsersAndCategories()
  dialogVisible.value = true
}

async function onDelete(row: Cert) {
  try {
    await ElMessageBox.confirm('确认删除该证件？', '提示', { type: 'warning' })
    await request.post('/api/cert/delete', { id: row.id })
    ElMessage.success('删除成功')
    fetchList()
  } catch {
    // cancelled
  }
}

async function onDeleteBatch() {
  if (selectedCerts.value.length === 0) {
    ElMessage.warning('请先选择要删除的证件')
    return
  }
  try {
    await ElMessageBox.confirm(`确认删除选中的 ${selectedCerts.value.length} 个证件？`, '提示', { type: 'warning' })
    await request.post('/api/cert/delete/batch', {
      ids: selectedCerts.value.map(c => c.id),
    })
    ElMessage.success('批量删除成功')
    selectedCerts.value = []
    fetchList()
  } catch {
    // cancelled
  }
}

function handleSelectionChange(val: Cert[]) {
  selectedCerts.value = val
}

function openInNewTab() {
  window.open(previewUrl.value, '_blank')
}

function authProxyUrl(fileUrl: string): string {
  const token = localStorage.getItem('token') || ''
  return `/api/file/proxy?url=${encodeURIComponent(fileUrl)}&token=${token}`
}

async function onPreview(row: Cert) {
  if (!row.fileUrl || row.fileUrl === '') {
    ElMessage.warning('该证件没有上传附件')
    return
  }
  previewUrl.value = row.fileUrl ? authProxyUrl(row.fileUrl) : ''
  isPdf.value = row.fileType === 'pdf'
  previewVisible.value = true
}

async function onSubmit() {
  if (!form.value.userId || !form.value.categoryId || !form.value.name) {
    ElMessage.warning('请填写完整信息')
    return
  }
  try {
    const url = isEdit.value ? '/api/cert/update' : '/api/cert/add'
    await request.post(url, form.value)
    ElMessage.success(isEdit.value ? '更新成功' : '添加成功')
    dialogVisible.value = false
    fetchList()
  } catch {
    // handled by interceptor
  }
}

async function handleUpload(options: any) {
  const fd = new FormData()
  fd.append('file', options.file)
  try {
    const res: any = await request.post('/api/upload', fd, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
    form.value.fileUrl = res.data.url
    form.value.fileType = res.data.fileType
    form.value.thumbUrl = res.data.thumbUrl || ''
    ElMessage.success('上传成功')
  } catch {
    options.onError(new Error('上传失败'))
  }
}

onMounted(fetchList)
</script>

<template>
  <div>
    <div class="toolbar">
      <el-input v-model="keyword" placeholder="搜索姓名/证件名称" style="width: 240px; margin-right: 12px" @keyup.enter="onSearch" />
      <el-button type="primary" @click="onSearch">搜索</el-button>
      <el-button type="success" @click="onAdd">新增证件</el-button>
      <el-button type="danger" :disabled="selectedCerts.length === 0" @click="onDeleteBatch">批量删除</el-button>
    </div>
    <el-table :data="list" v-loading="loading" stripe @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" />
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="userName" label="员工" />
      <el-table-column prop="categoryName" label="证件类型" />
      <el-table-column prop="name" label="证件名称" />
      <el-table-column prop="certNo" label="证件编号" />
      <el-table-column prop="issuer" label="发证机构" />
      <el-table-column prop="level" label="等级" width="100" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="row.status === 1 ? 'success' : 'danger'">{{ row.status === 1 ? '正常' : '禁用' }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250">
        <template #default="{ row }">
          <el-popover :width="240" trigger="hover" placement="top" :disabled="!row.fileUrl">
            <template #default>
              <div style="text-align: center">
                <el-image
                  v-if="row.fileType !== 'pdf' && row.fileType !== ''"
                  :src="authProxyUrl(row.fileUrl)"
                  style="max-width: 200px; max-height: 200px"
                  fit="contain"
                />
                <div v-else style="padding: 10px; text-align: center; color: #999">
                  {{ row.fileType === 'pdf' ? 'PDF 文件，点击查看' : '暂无可预览附件' }}
                </div>
              </div>
            </template>
            <template #reference>
              <el-button size="small" type="primary" @click="onPreview(row)">预览</el-button>
            </template>
          </el-popover>
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

    <el-dialog v-model="previewVisible" title="证件预览" width="800px" destroy-on-close>
      <div style="text-align: center; min-height: 300px">
        <iframe v-if="isPdf" :src="previewUrl" style="width: 100%; height: 600px; border: none" />
        <el-image v-else :src="previewUrl" style="max-width: 100%; max-height: 600px" fit="contain" />
      </div>
      <template #footer>
        <el-button @click="previewVisible = false">关闭</el-button>
        <el-button type="primary" @click="openInNewTab">新窗口打开</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="dialogVisible" :title="isEdit ? '编辑证件' : '新增证件'" width="560px">
      <el-form :model="form" label-width="90px">
        <el-form-item label="员工" required>
          <el-select v-model="form.userId" style="width: 100%">
            <el-option v-for="u in users" :key="u.id" :label="u.name" :value="u.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="证件类型" required>
          <el-select v-model="form.categoryId" style="width: 100%">
            <el-option v-for="c in categories" :key="c.id" :label="c.name" :value="c.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="证件名称" required>
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item label="证件编号">
          <el-input v-model="form.certNo" />
        </el-form-item>
        <el-form-item label="发证机构">
          <el-input v-model="form.issuer" />
        </el-form-item>
        <el-form-item label="发证日期">
          <el-date-picker v-model="form.issueDate" type="date" value-format="YYYY-MM-DD" style="width: 100%" />
        </el-form-item>
        <el-form-item label="到期日期">
          <el-date-picker v-model="form.expireDate" type="date" value-format="YYYY-MM-DD" style="width: 100%" />
        </el-form-item>
        <el-form-item label="等级">
          <el-input v-model="form.level" />
        </el-form-item>
        <el-form-item label="附件">
          <el-upload
            action="#"
            :http-request="handleUpload"
            :show-file-list="false"
            accept="image/*,application/pdf"
          >
            <el-button type="primary">上传附件</el-button>
          </el-upload>
          <el-link v-if="form.fileUrl" :href="form.fileUrl" target="_blank">查看附件</el-link>
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
