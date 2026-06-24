<script setup lang="ts">
import { ref, onMounted } from 'vue'
import {
  User,
  Document,
  Tickets,
  Download,
  ArrowRight,
  TrendCharts,
  Calendar,
  Timer,
} from '@element-plus/icons-vue'
import request from '@/utils/request'

const stats = ref({
  userCount: 0,
  certCount: 0,
  categoryCount: 0,
  exportCount: 0,
})
const loading = ref(false)

const shortcuts = [
  {
    title: '用户管理',
    path: '/users',
    icon: User,
    color: '#6366f1',
    bg: 'linear-gradient(135deg, #6366f1 0%, #818cf8 100%)',
    desc: '管理公司员工的账号信息',
  },
  {
    title: '证件类型',
    path: '/categories',
    icon: Tickets,
    color: '#10b981',
    bg: 'linear-gradient(135deg, #10b981 0%, #34d399 100%)',
    desc: '维护各类证件的分类定义',
  },
  {
    title: '证件管理',
    path: '/certificates',
    icon: Document,
    color: '#f59e0b',
    bg: 'linear-gradient(135deg, #f59e0b 0%, #fbbf24 100%)',
    desc: '查看与维护员工证件信息',
  },
  {
    title: '导出数据',
    path: '/export',
    icon: Download,
    color: '#ef4444',
    bg: 'linear-gradient(135deg, #ef4444 0%, #f87171 100%)',
    desc: '按条件导出证件汇总报表',
  },
]

async function fetchStats() {
  loading.value = true
  try {
    const [uRes, cRes, catRes, eRes]: any = await Promise.all([
      request.get('/api/user/list', { params: { page: 1, pageSize: 1 } }),
      request.get('/api/cert/list', { params: { page: 1, pageSize: 1 } }),
      request.get('/api/category/list', { params: { page: 1, pageSize: 1 } }),
      request.get('/api/task/list', { params: { page: 1, pageSize: 1 } }),
    ])
    stats.value = {
      userCount: uRes.data.total || 0,
      certCount: cRes.data.total || 0,
      categoryCount: catRes.data.total || 0,
      exportCount: eRes.data.total || 0,
    }
  } finally {
    loading.value = false
  }
}

onMounted(fetchStats)
</script>

<template>
  <div class="dashboard">
    <div class="page-header">
      <div class="page-title">
        <h1>欢迎回到 HRIDoc</h1>
        <p>实时掌握企业证件数据动态</p>
      </div>
      <el-tag type="info" effect="light" round>
        <el-icon><Calendar /></el-icon>
        {{ new Date().toLocaleDateString() }}
      </el-tag>
    </div>

    <!-- 统计卡片 -->
    <el-row :gutter="16" v-loading="loading">
      <el-col :xs="24" :sm="12" :lg="6">
        <div class="stat-card">
          <div class="stat-icon-wrapper" style="background: linear-gradient(135deg, #6366f1, #818cf8)">
            <el-icon size="24" color="#fff"><User /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.userCount }}</div>
            <div class="stat-label">员工总数</div>
          </div>
          <div class="stat-trend">
            <span class="trend-up"><TrendCharts /> +12%</span>
          </div>
        </div>
      </el-col>

      <el-col :xs="24" :sm="12" :lg="6">
        <div class="stat-card">
          <div class="stat-icon-wrapper" style="background: linear-gradient(135deg, #10b981, #34d399)">
            <el-icon size="24" color="#fff"><Document /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.certCount }}</div>
            <div class="stat-label">证件总数</div>
          </div>
          <div class="stat-trend">
            <span class="trend-up"><TrendCharts /> +8%</span>
          </div>
        </div>
      </el-col>

      <el-col :xs="24" :sm="12" :lg="6">
        <div class="stat-card">
          <div class="stat-icon-wrapper" style="background: linear-gradient(135deg, #f59e0b, #fbbf24)">
            <el-icon size="24" color="#fff"><Tickets /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.categoryCount }}</div>
            <div class="stat-label">证件类型</div>
          </div>
          <div class="stat-trend">
            <span class="trend-up"><TrendCharts /> +5%</span>
          </div>
        </div>
      </el-col>

      <el-col :xs="24" :sm="12" :lg="6">
        <div class="stat-card">
          <div class="stat-icon-wrapper" style="background: linear-gradient(135deg, #ef4444, #f87171)">
            <el-icon size="24" color="#fff"><Download /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-value">{{ stats.exportCount }}</div>
            <div class="stat-label">导出任务</div>
          </div>
          <div class="stat-trend">
            <span class="trend-up"><TrendCharts /> +23%</span>
          </div>
        </div>
      </el-col>
    </el-row>

    <!-- 快捷入口 -->
    <el-row :gutter="16" style="margin-top: 20px">
      <el-col :span="24">
        <div class="section-header">
          <span class="section-title"><el-icon size="15"><Timer /></el-icon> 快捷入口</span>
          <span class="section-desc">快速进入常用功能模块</span>
        </div>
      </el-col>

      <el-col
        v-for="item in shortcuts"
        :key="item.path"
        :xs="24"
        :sm="12"
        :lg="6"
        style="margin-bottom: 16px"
      >
        <div class="shortcut-card" @click="$router.push(item.path)">
          <div class="shortcut-left">
            <div class="shortcut-icon" :style="{ background: item.bg }">
              <el-icon size="22" color="#fff">
                <component :is="item.icon" />
              </el-icon>
            </div>
            <div class="shortcut-info">
              <div class="shortcut-title">{{ item.title }}</div>
              <div class="shortcut-desc">{{ item.desc }}</div>
            </div>
          </div>
          <div class="shortcut-arrow">
            <el-icon><ArrowRight /></el-icon>
          </div>
        </div>
      </el-col>
    </el-row>

    <!-- 底部占位 -->
    <el-row :gutter="16" style="margin-top: 8px">
      <el-col :span="24">
        <div class="tip-card">
          <el-icon size="16" color="#6366f1"><InfoFilled /></el-icon>
          <span>提示：点击上方快捷入口可快速进入对应模块，系统支持 Excel 批量导入与带水印的导出报表。</span>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
.dashboard {
  animation: fadeIn 0.5s ease;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.page-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24px;
}

.page-title h1 {
  font-size: 20px;
  font-weight: 600;
  color: #1e293b;
  margin: 0 0 4px;
}

.page-title p {
  font-size: 13px;
  color: #94a3b8;
  margin: 0;
}

.stat-card {
  background: #fff;
  border-radius: 14px;
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
  border: 1px solid #f1f5f9;
  transition: all 0.3s;
  position: relative;
  overflow: hidden;
}

.stat-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
}

.stat-card::after {
  content: '';
  position: absolute;
  top: 0;
  right: 0;
  width: 80px;
  height: 80px;
  background: radial-gradient(circle, rgba(99, 102, 241, 0.04) 0%, transparent 70%);
  border-radius: 50%;
  transform: translate(30%, -30%);
}

.stat-icon-wrapper {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 26px;
  font-weight: 700;
  color: #1e293b;
  line-height: 1.2;
}

.stat-label {
  font-size: 13px;
  color: #94a3b8;
  margin-top: 4px;
}

.stat-trend {
  font-size: 12px;
  font-weight: 500;
}

.trend-up {
  display: flex;
  align-items: center;
  gap: 2px;
  color: #10b981;
}

.section-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 4px;
}

.section-title {
  font-size: 15px;
  font-weight: 600;
  color: #334155;
  display: flex;
  align-items: center;
  gap: 6px;
  line-height: 1;
}

.section-desc {
  font-size: 12px;
  color: #94a3b8;
  margin-left: auto;
}

.shortcut-card {
  background: #fff;
  border-radius: 14px;
  padding: 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
  border: 1px solid #f1f5f9;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
  transition: all 0.3s;
}

.shortcut-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
  border-color: #e2e8f0;
}

.shortcut-left {
  display: flex;
  align-items: center;
  gap: 14px;
}

.shortcut-icon {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.shortcut-title {
  font-size: 15px;
  font-weight: 600;
  color: #1e293b;
}

.shortcut-desc {
  font-size: 12px;
  color: #94a3b8;
  margin-top: 2px;
}

.shortcut-arrow {
  color: #cbd5e1;
  transition: all 0.3s;
}

.shortcut-card:hover .shortcut-arrow {
  color: #6366f1;
  transform: translateX(3px);
}

.tip-card {
  background: #fff;
  border-radius: 12px;
  padding: 16px 20px;
  display: flex;
  align-items: center;
  gap: 10px;
  border: 1px solid #f1f5f9;
  font-size: 13px;
  color: #64748b;
}

.tip-card span {
  line-height: 1.5;
}
</style>
