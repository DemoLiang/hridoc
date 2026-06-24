# hridoc - 公司员工证件管理系统

hridoc 是一个基于 go-zero + Vue3 的公司员工证件管理系统，支持员工信息管理、证件类型维护、证件上传与预览、Excel 批量导入、按需导出汇总报表（含水印）以及操作日志记录等功能。

## 技术栈

- **后端**: Go 1.22 + go-zero v1.7.4 + MySQL 5.7 + Redis 7 + MinIO
- **前端**: Vue 3 + TypeScript + Vite + Element Plus 2.14.1 + Pinia
- **部署**: Docker + Docker Compose

## 功能特性

- 用户管理：员工账号的增删改查，支持身份证、学历、联系方式等信息
- 证件类型管理：自定义证件分类（如软考证书、一级建造师等）
- 证件管理：证件上传（图片/PDF）、缩略图生成、文件预览
- Excel 导入：按模板批量导入用户与证件信息
- 导出报表：选择员工和证件类型，异步生成汇总 Excel，支持水印配置
- 操作日志：自动记录关键接口的操作人、模块、动作与 IP
- 文件存储：基于 MinIO 的对象存储，支持预签名下载链接

## 环境要求

- Docker >= 20.10
- Docker Compose >= 2.0
- Go >= 1.22（本地开发）
- Node.js >= 20（本地开发）

## 目录结构

```
hridoc/
├── api/                    # go-zero 后端
│   ├── etc/api.yaml        # 后端配置文件
│   ├── internal/           # 业务逻辑
│   ├── model/              # 数据库模型
│   ├── pkg/                # 公共包（水印、错误码、MinIO 客户端）
│   └── Dockerfile
├── web/                    # Vue3 前端
│   ├── src/views/          # 页面视图
│   ├── src/router/         # 路由配置
│   ├── src/stores/         # Pinia 状态管理
│   └── Dockerfile
├── init.sql                # 数据库初始化脚本
├── docker-compose.yaml     # 服务编排
└── README.md
```

## 快速启动（Docker Compose）

### 1. 创建数据目录与 .env 文件

```bash
mkdir -p /data/hridoc/mysql /data/hridoc/minio /data/hridoc/fonts /data/hridoc/tmp
cp init.sql /data/hridoc/
```

创建 `.env` 文件：

```bash
DATA_DIR=/data/hridoc
PASSWORD=your_db_password
```

### 2. 启动服务

```bash
docker compose up -d
```

服务启动后：

| 服务 | 地址 | 说明 |
|------|------|------|
| Web 前端 | http://localhost | Nginx 托管的 SPA |
| API 后端 | http://localhost:8888 | go-zero REST API |
| MinIO 控制台 | http://localhost:19001 | 对象存储管理界面 |
| MySQL | localhost:13306 | 数据库 |
| Redis | localhost:16379 | 缓存 |

### 3. 初始化数据库

首次启动时，MySQL 会自动执行 `init.sql` 创建表结构并插入默认管理员账号。

### 4. 访问系统

打开浏览器访问 http://localhost，使用默认管理员账号登录：

- **账号**: `admin`
- **密码**: `admin123`

## 本地开发

### 后端

```bash
cd api
go mod download
go run hridoc.go -f etc/api.yaml
```

### 前端

```bash
cd web
npm install
npm run dev
```

前端开发服务器默认运行在 http://localhost:5173，API 代理到 http://localhost:8888。

## 配置说明

### 后端配置（api/etc/api.yaml）

```yaml
Name: hridoc-api
Host: 0.0.0.0
Port: 8888

Mysql:
  DataSource: root:root@tcp(localhost:3306)/hridoc?charset=utf8mb4&parseTime=true

CacheRedis:
  - Host: localhost:6379
    Pass: ""
    Type: node

JwtAuth:
  AccessSecret: hridoc-secret-key-change-in-production
  AccessExpire: 86400

MinIO:
  Endpoint: localhost:9000
  AccessKey: minio
  SecretKey: minio123
  Bucket: hridoc-bucket
  UseSSL: false
  PresignedExpiry: 300

Upload:
  MaxFileSize: 10485760
  AllowedTypes:
    - image/jpeg
    - image/png
    - image/jpg
    - application/pdf

Export:
  TempDir: ./tmp/exports
  CleanupDays: 7
```

### 前端配置（web/.env）

```
VITE_API_BASE_URL=http://localhost:8888
```

## API 概览

| 接口 | 方法 | 说明 |
|------|------|------|
| /api/login | POST | 登录获取 JWT |
| /api/user/* | GET/POST | 用户 CRUD |
| /api/category/* | GET/POST | 证件类型 CRUD |
| /api/cert/* | GET/POST | 证件 CRUD |
| /api/upload | POST | 文件上传 |
| /api/preview | POST | 导出预览 |
| /api/export | POST | 创建导出任务 |
| /api/task/list | GET | 任务历史 |
| /api/task/download/:id | GET | 获取下载链接 |
| /api/excel/template | GET | 下载导入模板 |
| /api/excel/import | POST | Excel 批量导入 |
| /api/clean | POST | 清理旧导出文件 |

## 默认账号

系统初始化时会自动创建以下账号：

| 角色 | 用户名 | 密码 |
|------|--------|------|
| 管理员 | admin | admin123 |
| 普通用户 | user1 | 123456 |

## 注意事项

1. **生产环境部署前**，务必修改以下配置：
   - `JwtAuth.AccessSecret`：JWT 签名密钥
   - MySQL 与 Redis 密码
   - MinIO 的 `SecretKey`
   - 前端 `VITE_API_BASE_URL`

2. **MinIO Bucket**：首次启动后需要确保 `hridoc-bucket` 存在，可登录 MinIO 控制台（http://localhost:19001，账号 minio / minio123）手动创建，或配置自动创建策略。

3. **水印字体**：PDF 水印使用 pdfcpu 内置字体，图片水印使用 Go 基础字体。如需自定义中文字体，请将 TTF 字体文件放入 `./fonts/` 目录并在配置中指定 `Watermark.FontPath`。

4. **导出文件清理**：系统支持手动调用 `/api/clean` 清理过期导出文件，建议在服务器配置定时任务自动执行。
