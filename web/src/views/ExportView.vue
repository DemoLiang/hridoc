<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import request from '@/utils/request'

interface Category {
  id: number
  name: string
  code: string
}

interface User {
  id: number
  name: string
}

interface PreviewCategory {
  categoryCode: string
  categoryName: string
  hasCert: boolean
  certId: number
  certName: string
}

interface ExcelPreviewUser {
  rowIndex: number
  name: string
  idCard: string
  matchStatus: number
  userId: number
  userName: string
  categories: PreviewCategory[]
  missReason: string
}

interface PreviewData {
  totalCount: number
  matchedCount: number
  missCount: number
  unmatchedCount: number
  users: ExcelPreviewUser[]
}

interface Task {
  id: number
  taskName: string
  userCount: number
  certCount: number
  missCount: number
  status: number
  failReason: string
  fileUrl: string
  createdAt: string
  completedAt: string
}

const categories = ref<Category[]>([])
const selectedCategories = ref<string[]>([])
const watermarkText = ref('')
const watermarkMode = ref<'diagonal' | 'horizontal'>('diagonal')
const watermarkColor = ref('#D0E0FF')
const watermarkOpacity = ref(0.05)
const watermarkFontSize = ref(44)

const watermarkPreviewVisible = ref(false)
const watermarkPreviewUrl = ref('')
const watermarkPreviewLoading = ref(false)

const uploadedFile = ref<File | null>(null)
const previewLoading = ref(false)
const previewData = ref<PreviewData | null>(null)

const taskList = ref<Task[]>([])
const taskTotal = ref(0)
const taskPage = ref(1)
const taskPageSize = ref(10)
const taskLoading = ref(false)

const users = ref<User[]>([])
const selectedUsers = ref<number[]>([])
const userLoading = ref(false)

async function fetchCategories() {
  const res: any = await request.get('/api/category/list', { params: { page: 1, pageSize: 1000 } })
  categories.value = res.data.list || []
}

async function fetchUsers() {
  userLoading.value = true
  try {
    const res: any = await request.get('/api/user/list', { params: { page: 1, pageSize: 1000 } })
    users.value = res.data.list || []
  } finally {
    userLoading.value = false
  }
}

function handleFileChange(file: any) {
  uploadedFile.value = file.raw
  previewData.value = null
}

function buildFormData(): FormData | null {
  if (selectedCategories.value.length === 0) {
    ElMessage.warning('请选择证件类型')
    return null
  }
  const formData = new FormData()
  if (uploadedFile.value) {
    formData.append('file', uploadedFile.value)
  }
  if (selectedUsers.value.length > 0) {
    formData.append('userIds', selectedUsers.value.join(','))
  }
  formData.append('categoryCodes', selectedCategories.value.join(','))
  formData.append('watermarkMode', watermarkMode.value)
  if (watermarkText.value) {
    formData.append('watermarkText', watermarkText.value)
    formData.append('watermarkColor', watermarkColor.value)
    formData.append('watermarkOpacity', String(watermarkOpacity.value))
    formData.append('watermarkFontSize', String(watermarkFontSize.value))
  }
  return formData
}

function downloadFile(url: string, filename: string) {
  const a = document.createElement('a')
  a.href = url
  a.download = filename
  a.style.display = 'none'
  document.body.appendChild(a)
  a.click()
  setTimeout(() => {
    document.body.removeChild(a)
  }, 200)
}

async function onDownloadTemplate() {
  try {
    const token = localStorage.getItem('token') || ''
    const res = await fetch('/api/excel/template/download', {
      headers: { Authorization: `Bearer ${token}` },
    })
    if (!res.ok) {
      ElMessage.error('下载模板失败')
      return
    }
    const blob = await res.blob()
    const url = URL.createObjectURL(blob)
    downloadFile(url, 'import_template.xlsx')
    URL.revokeObjectURL(url)
  } catch {
    ElMessage.error('下载模板请求失败')
  }
}

async function onPreview() {
  if (selectedCategories.value.length === 0) {
    ElMessage.warning('请选择证件类型')
    return
  }
  if (!uploadedFile.value) {
    ElMessage.warning('请先上传 Excel 名单文件')
    return
  }

  previewLoading.value = true
  try {
    const formData = new FormData()
    formData.append('file', uploadedFile.value)
    formData.append('categoryCodes', selectedCategories.value.join(','))
    const res: any = await request.post('/api/preview', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
    if (res.code === 0) {
      previewData.value = res.data
    } else {
      ElMessage.error(res.message || '预览失败')
    }
  } catch (e) {
    ElMessage.error('预览请求失败')
  } finally {
    previewLoading.value = false
  }
}

async function onExport() {
  const formData = buildFormData()
  if (!formData) return

  if (!uploadedFile.value && selectedUsers.value.length === 0) {
    ElMessage.warning('请上传 Excel 名单或选择人员')
    return
  }

  try {
    const res: any = await request.post('/api/export', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
    if (res.code === 0) {
      ElMessage.success('导出任务已创建，任务ID: ' + res.data.taskId)
      fetchTaskList()
    } else {
      ElMessage.error(res.message || '创建导出任务失败')
    }
  } catch {
    ElMessage.error('导出请求失败')
  }
}

async function onUserExport() {
  if (selectedCategories.value.length === 0) {
    ElMessage.warning('请选择证件类型')
    return
  }
  if (selectedUsers.value.length === 0) {
    ElMessage.warning('请选择要导出的用户')
    return
  }

  const formData = new FormData()
  formData.append('categoryCodes', selectedCategories.value.join(','))
  formData.append('userIds', selectedUsers.value.join(','))
  formData.append('watermarkMode', watermarkMode.value)
  if (watermarkText.value) {
    formData.append('watermarkText', watermarkText.value)
    formData.append('watermarkColor', watermarkColor.value)
    formData.append('watermarkOpacity', String(watermarkOpacity.value))
    formData.append('watermarkFontSize', String(watermarkFontSize.value))
  }

  try {
    const res: any = await request.post('/api/export', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
    if (res.code === 0) {
      ElMessage.success('导出任务已创建，任务ID: ' + res.data.taskId)
      fetchTaskList()
    } else {
      ElMessage.error(res.message || '创建导出任务失败')
    }
  } catch {
    ElMessage.error('导出请求失败')
  }
}

async function onWatermarkPreview() {
  if (!watermarkText.value) {
    ElMessage.warning('请输入水印文字')
    return
  }

  watermarkPreviewLoading.value = true
  try {
    const formData = new FormData()
    formData.append('watermarkText', watermarkText.value)
    formData.append('watermarkMode', watermarkMode.value)
    formData.append('watermarkColor', watermarkColor.value)
    formData.append('watermarkOpacity', String(watermarkOpacity.value))
    formData.append('watermarkFontSize', String(watermarkFontSize.value))

    const token = localStorage.getItem('token') || ''
    const res = await fetch('/api/watermark/preview', {
      method: 'POST',
      headers: { Authorization: `Bearer ${token}` },
      body: formData,
    })
    if (!res.ok) {
      ElMessage.error('预览生成失败')
      return
    }
    const blob = await res.blob()
    watermarkPreviewUrl.value = URL.createObjectURL(blob)
    watermarkPreviewVisible.value = true
  } catch {
    ElMessage.error('预览请求失败')
  } finally {
    watermarkPreviewLoading.value = false
  }
}

async function fetchTaskList() {
  taskLoading.value = true
  try {
    const res: any = await request.get('/api/task/list', {
      params: { page: taskPage.value, pageSize: taskPageSize.value },
    })
    taskList.value = res.data.list || []
    taskTotal.value = res.data.total || 0
  } finally {
    taskLoading.value = false
  }
}

async function onDownload(task: Task) {
  try {
    const token = localStorage.getItem('token') || ''
    const res = await fetch(`/api/task/download/${task.id}/file`, {
      headers: { Authorization: `Bearer ${token}` },
    })
    if (!res.ok) {
      ElMessage.error('下载文件失败')
      return
    }
    const blob = await res.blob()
    const zipBlob = new Blob([blob], { type: 'application/zip' })
    const url = URL.createObjectURL(zipBlob)
    const a = document.createElement('a')
    a.href = url
    a.download = `task_${task.id}.zip`
    a.style.display = 'none'
    document.body.appendChild(a)
    a.click()
    setTimeout(() => {
      document.body.removeChild(a)
      URL.revokeObjectURL(url)
    }, 200)
  } catch {
    ElMessage.error('下载请求失败')
  }
}

async function onClean() {
  try {
    await ElMessageBox.confirm('确认清理旧导出文件？', '提示', { type: 'warning' })
    const res: any = await request.post('/api/clean')
    ElMessage.success('已清理 ' + (res.data.deletedFiles || 0) + ' 个文件')
    fetchTaskList()
  } catch {
    // cancelled
  }
}

function statusText(status: number) {
  return { 1: '处理中', 2: '已完成', 3: '失败' }[status] || '未知'
}

function statusType(status: number) {
  return { 1: 'info', 2: 'success', 3: 'danger' }[status] || 'info'
}

function matchStatusText(status: number) {
  return { 0: '未匹配', 1: '身份证号匹配', 2: '姓名匹配', 3: '重名无法匹配' }[status] || '未知'
}

function matchStatusType(status: number) {
  return { 0: 'danger', 1: 'success', 2: 'success', 3: 'warning' }[status] || 'info'
}

onMounted(() => {
  fetchCategories()
  fetchUsers()
  fetchTaskList()
})
</script>

<template>
  <div>
    <el-card class="card">
      <template #header>
        <span>选择导出条件</span>
      </template>
      <el-form label-width="120px">
        <el-form-item label="上传 Excel 名单">
          <el-upload
            action="#"
            :auto-upload="false"
            :limit="1"
            :on-change="handleFileChange"
            accept=".xlsx,.xls"
          >
            <el-button type="primary">选择文件</el-button>
            <template #tip>
              <div class="upload-tip">
                请上传包含"姓名"和"身份证号"两列的 Excel 文件，第 1 行为表头
                <el-link type="primary" style="margin-left: 8px" @click="onDownloadTemplate">下载导入模板</el-link>
              </div>
            </template>
          </el-upload>
        </el-form-item>
        <el-form-item label="或选择人员">
          <el-select-v2
            v-model="selectedUsers"
            :options="users.map(u => ({ value: u.id, label: u.name }))"
            placeholder="直接选择系统用户导出，可多选"
            multiple
            collapse-tags
            :loading="userLoading"
            style="width: 400px"
          />
        </el-form-item>
        <el-form-item label="选择证件类型">
          <el-select-v2
            v-model="selectedCategories"
            :options="categories.map(c => ({ value: c.code, label: c.name }))"
            placeholder="请选择证件类型"
            multiple
            collapse-tags
            style="width: 400px"
          />
        </el-form-item>
        <el-form-item label="水印文字">
          <el-input v-model="watermarkText" placeholder="可选，导出文件将添加水印" style="width: 400px" />
        </el-form-item>
        <el-form-item label="水印方向">
          <el-radio-group v-model="watermarkMode">
            <el-radio label="diagonal">45度斜铺</el-radio>
            <el-radio label="horizontal">横向平铺</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="水印颜色">
          <el-color-picker v-model="watermarkColor" :predefine="['#D0D0D0', '#A0A0A0', '#808080', '#000000']" />
        </el-form-item>
        <el-form-item label="水印透明度">
          <el-slider v-model="watermarkOpacity" :min="0.01" :max="1" :step="0.01" show-input style="width: 400px" />
        </el-form-item>
        <el-form-item label="水印字号">
          <el-input-number v-model="watermarkFontSize" :min="12" :max="120" :step="1" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="previewLoading" @click="onPreview" :disabled="!uploadedFile">预览匹配结果</el-button>
          <el-button type="info" :loading="watermarkPreviewLoading" @click="onWatermarkPreview" :disabled="!watermarkText">预览水印效果</el-button>
          <el-button type="success" :disabled="(previewData && previewData.matchedCount === 0) || !uploadedFile" @click="onExport">创建导出任务</el-button>
          <el-button type="success" :disabled="selectedUsers.length === 0 || selectedCategories.length === 0" @click="onUserExport">选人导出</el-button>
          <el-button type="danger" @click="onClean">清理旧文件</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card v-if="previewData" class="card">
      <template #header>
        <span>
          预览结果 — 共 {{ previewData.totalCount }} 人，
          <el-tag type="success">匹配成功 {{ previewData.matchedCount }}</el-tag>
          <el-tag type="warning" style="margin-left: 8px">缺证 {{ previewData.missCount }}</el-tag>
          <el-tag type="danger" style="margin-left: 8px">未匹配 {{ previewData.unmatchedCount }}</el-tag>
        </span>
      </template>
      <el-table :data="previewData.users" stripe size="small">
        <el-table-column prop="rowIndex" label="Excel 行号" width="100" />
        <el-table-column prop="name" label="姓名" width="120" />
        <el-table-column prop="idCard" label="身份证号" width="180" />
        <el-table-column prop="matchStatus" label="匹配状态" width="130">
          <template #default="{ row }">
            <el-tag :type="matchStatusType(row.matchStatus)">{{ matchStatusText(row.matchStatus) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="missReason" label="提示" min-width="180" />
        <el-table-column
          v-for="cat in categories.filter(c => selectedCategories.includes(c.code))"
          :key="cat.code"
          :label="cat.name"
          width="120"
        >
          <template #default="{ row }">
            <template v-if="row.matchStatus === 1 || row.matchStatus === 2">
              <el-tag
                v-if="row.categories.find((c: any) => c.categoryCode === cat.code)?.hasCert"
                type="success"
                size="small"
              >
                {{ row.categories.find((c: any) => c.categoryCode === cat.code)?.certName }}
              </el-tag>
              <el-tag v-else type="danger" size="small">缺证</el-tag>
            </template>
            <span v-else>-</span>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-card class="card">
      <template #header>
        <div class="card-header">
          <span>导出任务历史</span>
          <el-button :loading="taskLoading" size="small" @click="fetchTaskList">刷新</el-button>
        </div>
      </template>
      <el-table :data="taskList" v-loading="taskLoading" stripe>
        <el-table-column prop="taskName" label="任务名称" />
        <el-table-column prop="userCount" label="人数" width="80" />
        <el-table-column prop="certCount" label="证件数" width="90" />
        <el-table-column prop="missCount" label="缺证数" width="90" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="statusType(row.status)">{{ statusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="failReason" label="失败原因" show-overflow-tooltip />
        <el-table-column prop="createdAt" label="创建时间" />
        <el-table-column prop="completedAt" label="完成时间" />
        <el-table-column label="操作" width="120">
          <template #default="{ row }">
            <el-button v-if="row.status === 2" size="small" type="primary" @click="onDownload(row)">下载</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        v-model:current-page="taskPage"
        v-model:page-size="taskPageSize"
        :total="taskTotal"
        layout="total, prev, pager, next"
        @change="fetchTaskList"
        style="margin-top: 16px"
      />
    </el-card>

    <el-dialog v-model="watermarkPreviewVisible" title="水印效果预览" width="800px" destroy-on-close>
      <div style="text-align: center; min-height: 300px">
        <el-image :src="watermarkPreviewUrl" style="max-width: 100%; max-height: 600px" fit="contain" />
      </div>
    </el-dialog>
  </div>
</template>

<style scoped>
.card {
  margin-bottom: 16px;
}
.upload-tip {
  color: #999;
  font-size: 12px;
  margin-top: 8px;
}
</style>
