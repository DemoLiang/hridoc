<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user'
import request from '@/utils/request'
import { User, Lock, DocumentChecked } from '@element-plus/icons-vue'

const router = useRouter()
const userStore = useUserStore()

const form = ref({
  username: '',
  password: '',
})
const loading = ref(false)

async function handleLogin() {
  if (!form.value.username || !form.value.password) {
    ElMessage.warning('请输入用户名和密码')
    return
  }
  loading.value = true
  try {
    const res: any = await request.post('/api/login', form.value)
    userStore.setToken(res.data.token)
    ElMessage.success('登录成功')
    router.push('/')
  } catch {
    // error handled by interceptor
  } finally {
    loading.value = false
  }
}

const floatingItems = ref<{ left: string; top: string; size: number; delay: number; duration: number }[]>([])
onMounted(() => {
  for (let i = 0; i < 20; i++) {
    floatingItems.value.push({
      left: Math.random() * 100 + '%',
      top: Math.random() * 100 + '%',
      size: Math.random() * 60 + 20,
      delay: Math.random() * 5,
      duration: Math.random() * 10 + 10,
    })
  }
})
</script>

<template>
  <div class="login-page">
    <div class="floating-bg">
      <div
        v-for="(item, idx) in floatingItems"
        :key="idx"
        class="floating-item"
        :style="{
          left: item.left,
          top: item.top,
          width: item.size + 'px',
          height: item.size + 'px',
          animationDelay: item.delay + 's',
          animationDuration: item.duration + 's',
        }"
      />
    </div>

    <div class="login-content">
      <div class="login-left">
        <div class="brand">
          <el-icon size="56" color="#fff"><DocumentChecked /></el-icon>
          <h1 class="brand-title">HRIDoc</h1>
          <p class="brand-subtitle">企业员工证件数字化管理平台</p>
          <div class="brand-features">
            <div class="feature-item">
              <div class="feature-icon">&#10003;</div>
              <span>员工证件集中管理</span>
            </div>
            <div class="feature-item">
              <div class="feature-icon">&#10003;</div>
              <span>Excel 批量导入导出</span>
            </div>
            <div class="feature-item">
              <div class="feature-icon">&#10003;</div>
              <span>水印保护与安全审计</span>
            </div>
          </div>
        </div>
      </div>

      <div class="login-right">
        <div class="glass-card">
          <div class="card-header">
            <h2 class="welcome-text">欢迎回来</h2>
            <p class="hint-text">请登录您的账号</p>
          </div>

          <el-form :model="form" @submit.prevent="handleLogin" class="login-form">
            <el-form-item>
              <el-input
                v-model="form.username"
                placeholder="用户名"
                size="large"
                class="fancy-input"
              >
                <template #prefix>
                  <el-icon><User /></el-icon>
                </template>
              </el-input>
            </el-form-item>

            <el-form-item>
              <el-input
                v-model="form.password"
                type="password"
                placeholder="密码"
                size="large"
                show-password
                class="fancy-input"
                @keyup.enter="handleLogin"
              >
                <template #prefix>
                  <el-icon><Lock /></el-icon>
                </template>
              </el-input>
            </el-form-item>

            <el-form-item style="margin-top: 8px">
              <el-button
                type="primary"
                size="large"
                class="login-btn"
                :loading="loading"
                @click="handleLogin"
              >
                登 录
              </el-button>
            </el-form-item>
          </el-form>

          <div class="divider">
            <span>默认账号 admin / admin123</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%);
  position: relative;
  overflow: hidden;
}

.floating-bg {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 0;
}

.floating-item {
  position: absolute;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(99, 102, 241, 0.15) 0%, transparent 70%);
  animation: floatUp linear infinite;
}

@keyframes floatUp {
  0% {
    transform: translateY(0) scale(1);
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  90% {
    opacity: 1;
  }
  100% {
    transform: translateY(-100vh) scale(0.5);
    opacity: 0;
  }
}

.login-content {
  display: flex;
  width: 900px;
  max-width: 95vw;
  z-index: 1;
  animation: fadeInUp 0.8s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.login-left {
  flex: 1;
  display: flex;
  align-items: center;
  padding-right: 60px;
}

.brand {
  color: #fff;
}

.brand-title {
  font-size: 42px;
  font-weight: 700;
  margin: 16px 0 8px;
  letter-spacing: 2px;
  background: linear-gradient(90deg, #fff, #a5b4fc);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.brand-subtitle {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 40px;
}

.brand-features {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.feature-item {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 14px;
  color: rgba(255, 255, 255, 0.85);
}

.feature-icon {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  background: linear-gradient(135deg, #6366f1, #8b5cf6);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  color: #fff;
}

.login-right {
  width: 380px;
  flex-shrink: 0;
}

.glass-card {
  background: rgba(255, 255, 255, 0.08);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 20px;
  padding: 40px 36px;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.4);
}

.card-header {
  text-align: center;
  margin-bottom: 32px;
}

.welcome-text {
  font-size: 24px;
  font-weight: 600;
  color: #fff;
  margin: 0 0 6px;
}

.hint-text {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.55);
  margin: 0;
}

.login-form :deep(.el-form-item) {
  margin-bottom: 20px;
}

.fancy-input :deep(.el-input__wrapper) {
  background: rgba(255, 255, 255, 0.06);
  border: 1px solid rgba(255, 255, 255, 0.1);
  box-shadow: none !important;
  border-radius: 10px;
  padding: 4px 12px;
  transition: all 0.3s;
}

.fancy-input :deep(.el-input__wrapper:hover),
.fancy-input :deep(.el-input__wrapper.is-focus) {
  background: rgba(255, 255, 255, 0.1);
  border-color: rgba(99, 102, 241, 0.5);
}

.fancy-input :deep(.el-input__inner) {
  color: #fff;
  font-size: 14px;
}

.fancy-input :deep(.el-input__inner::placeholder) {
  color: rgba(255, 255, 255, 0.4);
}

.fancy-input :deep(.el-input__prefix) {
  color: rgba(255, 255, 255, 0.5);
  margin-right: 8px;
}

.login-btn {
  width: 100%;
  height: 44px;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 500;
  letter-spacing: 2px;
  background: linear-gradient(135deg, #6366f1, #8b5cf6);
  border: none;
  transition: all 0.3s;
}

.login-btn:hover {
  background: linear-gradient(135deg, #818cf8, #a78bfa);
  transform: translateY(-1px);
  box-shadow: 0 8px 20px rgba(99, 102, 241, 0.35);
}

.divider {
  margin-top: 24px;
  text-align: center;
  position: relative;
}

.divider::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(255,255,255,0.15), transparent);
}

.divider span {
  position: relative;
  background: transparent;
  padding: 0 12px;
  font-size: 12px;
  color: rgba(255, 255, 255, 0.35);
}

@media (max-width: 768px) {
  .login-content {
    flex-direction: column;
    align-items: center;
  }
  .login-left {
    display: none;
  }
  .login-right {
    width: 100%;
  }
}
</style>
