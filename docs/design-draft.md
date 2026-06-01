# hridoc 证件管理系统 - 设计草案

## 1. 系统架构

```
┌─────────────────────────────────────────────────────────────┐
│                         前端层                                │
│              Vue3 + Element Plus + Vite (SPA)               │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼ HTTP/JSON
┌─────────────────────────────────────────────────────────────┐
│                      后端 API 层                             │
│                   go-zero API Service                        │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │  User API   │  │  Cert API   │  │   Export API        │  │
│  │  用户管理    │  │  证件管理    │  │   导入/导出/水印     │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                              │
            ┌─────────────────┼─────────────────┐
            ▼                 ▼                 ▼
    ┌──────────────┐  ┌──────────────┐  ┌──────────────┐
    │   MySQL      │  │    Redis     │  │    MinIO     │
    │  业务数据     │  │  缓存/会话    │  │  文件存储     │
    └──────────────┘  └──────────────┘  └──────────────┘
```

### 部署拓扑
- **前端**：Nginx 托管静态资源（`dist/`）
- **后端**：直接运行 go-zero 编译出的 binary
- **依赖**：MySQL 5.7 + Redis + MinIO，全部通过 `docker-compose.yaml` 拉起

---

## 2. 数据库设计

### 2.1 用户表 (user)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT PK | 自增主键 |
| name | VARCHAR(64) | 姓名 |
| phone | VARCHAR(20) | 手机号 |
| email | VARCHAR(128) | 邮箱 |
| id_card | VARCHAR(18) | 身份证号（唯一） |
| education | VARCHAR(32) | 学历 |
| role | TINYINT | 角色：1-管理员，2-普通用户 |
| status | TINYINT | 状态：1-正常，2-禁用 |
| created_at / updated_at | DATETIME | 时间戳 |

### 2.2 证件类型表 (cert_category)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | INT PK | 自增主键 |
| name | VARCHAR(64) | 类型名称（如"软考证书"） |
| code | VARCHAR(32) | 类型编码 |
| description | VARCHAR(255) | 描述 |
| status | TINYINT | 状态：1-启用，2-禁用 |

### 2.3 证件表 (certificate)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT PK | 自增主键 |
| user_id | BIGINT FK | 关联用户 |
| category_id | INT FK | 关联证件类型 |
| name | VARCHAR(128) | 证书名称 |
| cert_no | VARCHAR(64) | 证书编号 |
| issuer | VARCHAR(128) | 发证机构 |
| issue_date | DATE | 发证日期 |
| expire_date | DATE | 有效期至（可为空） |
| level | VARCHAR(32) | 证书等级（初中高级） |
| file_url | VARCHAR(512) | MinIO 文件地址 |
| file_type | VARCHAR(16) | 文件类型：image/pdf |
| status | TINYINT | 状态：1-正常，2-过期 |
| created_at / updated_at | DATETIME | 时间戳 |

### 2.4 导出任务表 (export_task) - 用于记录导出历史
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT PK | 自增主键 |
| task_name | VARCHAR(128) | 任务名称（如"2024年软考证书导出"） |
| user_count | INT | 涉及人数 |
| cert_count | INT | 匹配到的证件数 |
| miss_count | INT | 缺证人数 |
| watermark_config | JSON | 水印配置JSON |
| file_url | VARCHAR(512) | 生成的压缩包地址 |
| status | TINYINT | 1-处理中，2-完成，3-失败 |
| created_at | DATETIME | 创建时间 |

---

## 3. 后端模块与 API 设计

### 3.1 模块划分

| 模块 | 职责 | 主要 API |
|------|------|----------|
| `sys` | 系统通用 | 登录、登出、上传接口 |
| `user` | 用户管理 | 增删改查用户 |
| `cert_category` | 证件类型 | 增删改查类型 |
| `certificate` | 证件管理 | 增删改查证件、预览 |
| `export` | 导入导出 | Excel 导入、预览匹配、导出打包 |

### 3.2 核心 API 列表

```
POST   /api/sys/login              登录（管理员密码登录）
POST   /api/sys/upload             通用文件上传（直传 MinIO）

GET    /api/user/list              用户列表（支持分页、按姓名/身份证搜索）
POST   /api/user/add               新增用户
POST   /api/user/update            更新用户
POST   /api/user/delete            删除用户

GET    /api/cert-category/list     证件类型列表
POST   /api/cert-category/add      新增类型
POST   /api/cert-category/update   更新类型
POST   /api/cert-category/delete   删除类型

GET    /api/certificate/list       证件列表（支持按用户、类型、等级筛选）
POST   /api/certificate/add        新增证件（含文件上传）
POST   /api/certificate/update     更新证件
POST   /api/certificate/delete     删除证件
GET    /api/certificate/preview/:id 证件预览（返回文件URL或缩略图）

POST   /api/export/import          上传 Excel 名单
POST   /api/export/preview         预览匹配结果（返回匹配/缺证/多证列表）
POST   /api/export/download        确认导出（传入水印配置 + 选中人员，返回压缩包下载链接）
GET    /api/export/task-list       导出历史列表
```

---

## 4. 核心流程：导入 → 预览 → 导出

### 4.1 Excel 导入

HR 下载模板 → 填写名单（姓名 + 身份证号，可选证件类型） → 上传 Excel

**模板格式**：
```
| 姓名 | 身份证号 | 证件类型（可选） | 证件等级（可选） |
```

后端解析 Excel，按身份证号匹配 user 表，返回预览数据。

### 4.2 预览匹配

返回三种状态：
- **匹配成功**：找到用户，且按条件（类型/等级）筛选后有证件
- **缺证**：找到用户，但没有符合条件的证件
- **未匹配**：身份证号在系统中不存在

对于"匹配成功"的用户，展示其符合条件的所有证件（一个人可能有多证）。

### 4.3 导出打包

HR 在预览页：
1. 勾选要导出的人员（或全选）
2. 配置水印（文字、位置、字号、透明度、颜色、旋转角度）
3. 点击"导出"

后端处理：
1. 读取勾选人员的证件文件（从 MinIO 下载到临时目录）
2. 对每份证件图片应用水印（使用 Go 图像库如 `fogleman/gg` 或 `golang.org/x/image`）
3. PDF 证件是否需要转图片打水印后再转回？或者直接用 PDF 库加水印？
   → 建议：PDF 用 `pdfcpu` 或 `fpdf` 加水印；图片用 Go 图像库
4. 将所有处理后的文件打包为 ZIP
5. 上传 ZIP 到 MinIO，返回下载链接
6. 记录导出任务到 export_task 表

---

## 5. 水印设计

水印配置结构（JSON）：
```json
{
  "text": "仅供XX公司内部使用",
  "position": "center",      // center / diagonal / bottom / top-left / top-right / bottom-left / bottom-right
  "fontSize": 48,
  "opacity": 0.3,            // 0.0 ~ 1.0
  "color": "#808080",        // 十六进制颜色
  "rotation": -45,           // 旋转角度（对角线用 -45）
  "fontFamily": "simhei"     // 字体（需内置支持中文）
}
```

- **图片水印**：使用 `fogleman/gg` 或 `github.com/nfnt/resize` + 标准库绘制
- **PDF 水印**：使用 `github.com/pdfcpu/pdfcpu` 添加文字水印层
- **字体问题**：需内置中文字体（如思源黑体或系统 simhei），避免服务器无中文字体导致乱码

---

## 6. 文件存储设计

### MinIO Bucket 结构

```
hridoc/
├── certificates/           # 原始证件文件
│   ├── 2024/06/user_001_cert_001.jpg
│   ├── 2024/06/user_001_cert_002.pdf
│   └── ...
└── exports/                # 导出压缩包
    ├── export_20240601_143022.zip
    └── ...
```

### 文件上传流程
1. 前端直传 MinIO（前端先向后端申请预签名 URL，然后直传 MinIO）
   → 优点：不占用后端带宽
   → 或者后端代理上传（简单但占带宽）
2. 上传成功后前端将文件 URL 提交给后端，保存到 certificate 表

---

## 7. 前端页面设计

### 页面清单

| 页面 | 说明 |
|------|------|
| 登录页 | 管理员密码登录 |
| 用户管理 | 用户增删改查、分页、搜索 |
| 证件类型 | 类型增删改查 |
| 证件管理 | 证件增删改查、按用户/类型/等级筛选、预览 |
| 导入导出 | 上传 Excel、预览匹配结果、配置水印、下载 |
| 导出历史 | 查看历史导出任务和下载链接 |

### 关键交互
- **证件预览**：点击证件行→弹窗/抽屉展示证件图片或 PDF（图片直接展示，PDF 用 iframe 或 pdf.js）
- **导出预览页**：表格展示匹配结果，三种状态用不同颜色标识（绿/红/灰），可勾选要导出的行
- **水印配置**：弹窗表单，实时预览水印效果（用 Canvas 在前端模拟）

---

## 8. 技术栈汇总

| 层级 | 选型 |
|------|------|
| 前端框架 | Vue 3 + Composition API |
| UI 库 | Element Plus |
| 构建工具 | Vite |
| 状态管理 | Pinia |
| HTTP 客户端 | Axios |
| 后端框架 | go-zero (api) |
| ORM | go-zero 内置 sqlx + gen |
| 缓存 | Redis（go-zero 内置 cache） |
| 文件存储 | MinIO |
| Excel 处理 | 前端 `xlsx.js` / 后端 `github.com/xuri/excelize/v2` |
| 图像处理 | `github.com/fogleman/gg` |
| PDF 处理 | `github.com/pdfcpu/pdfcpu` |
| 字体 | 内置思源黑体或系统 simhei |

---

## 9. docker-compose 扩展

在现有 mysql/redis/etcd 基础上，增加 MinIO：

```yaml
  minio:
    image: minio/minio
    ports:
      - 19000:9000
      - 19001:9001
    environment:
      MINIO_ROOT_USER: minio
      MINIO_ROOT_PASSWORD: minio123
    command: server /data --console-address ":9001"
```

---

以上为设计草案，请审阅各部分内容是否有需要调整的地方。
