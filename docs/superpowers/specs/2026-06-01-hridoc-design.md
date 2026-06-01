---
name: hridoc-system-design
description: hridoc 公司员工证件管理系统完整设计文档 - 单体 API + Vue3 SPA，支持证件管理、Excel 导入导出、水印生成
type: design
date: 2026-06-01
---

# hridoc 公司员工证件管理系统设计文档

## 1. 概述

### 1.1 项目背景
公司需要一套员工证件管理系统，用于集中管理员工的各类证书，支持 HR 通过 Excel 名单批量筛选人员、导出带水印的证件图片/PDF。

### 1.2 技术选型
| 层级 | 技术栈 |
|------|--------|
| 前端 | Vue 3 + Element Plus + Vite + Pinia + Axios |
| 后端 | go-zero (单体 api 服务) |
| 数据库 | MySQL 5.7 |
| 缓存 | Redis |
| 文件存储 | MinIO (docker-compose 拉起) |
| Excel 处理 | 后端: `github.com/xuri/excelize/v2` |
| 图像处理 | `github.com/fogleman/gg` |
| PDF 处理 | `github.com/pdfcpu/pdfcpu` |
| 字体加载 | 配置文件指定中文字体路径 |

---

## 2. 系统架构

```
+-------------------------------------------------------------+
|                         前端层                               |
|              Vue3 + Element Plus + Vite (SPA)              |
+-------------------------------------------------------------+
                              |
                              v HTTP/JSON
+-------------------------------------------------------------+
|                      后端 API 层                            |
|                   go-zero API Service                       |
|  +-------------+  +-------------+  +---------------------+ |
|  |  User API   |  |  Cert API   |  |   Export API        | |
|  |  用户管理    |  |  证件管理    |  |   导入/导出/水印     | |
|  +-------------+  +-------------+  +---------------------+ |
|  +-------------+  +-------------+                          |
|  |  Sys API    |  |  System     |                          |
|  |  系统通用    |  |  操作日志    |                          |
|  +-------------+  +-------------+                          |
+-------------------------------------------------------------+
                              |
            +-----------------+-----------------+
            v                 v                 v
    +--------------+  +--------------+  +--------------+
    |   MySQL      |  |    Redis     |  |    MinIO     |
    |  业务数据     |  |  缓存/会话    |  |  文件存储     |
    +--------------+  +--------------+  +--------------+
```

### 2.1 部署拓扑
- **前端**：Nginx 托管静态资源
- **后端**：go-zero 单 binary 运行
- **依赖**：MySQL 5.7 + Redis + MinIO，全部通过 `docker-compose.yaml` 拉起
- **字体**：Docker 构建时打包进镜像，路径由配置文件指定

---

## 3. 数据库设计

### 3.1 用户表 (user)
```sql
CREATE TABLE `user` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` VARCHAR(64) NOT NULL COMMENT '姓名',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号',
  `email` VARCHAR(128) DEFAULT NULL COMMENT '邮箱',
  `id_card` VARCHAR(18) NOT NULL COMMENT '身份证号',
  `education` VARCHAR(32) DEFAULT NULL COMMENT '学历',
  `role` TINYINT NOT NULL DEFAULT 2 COMMENT '角色：1-管理员，2-普通用户',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-正常，2-禁用',
  `password` VARCHAR(128) NOT NULL COMMENT '密码（bcrypt加密）',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_id_card` (`id_card`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
```

### 3.2 证件类型表 (cert_category)
```sql
CREATE TABLE `cert_category` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` VARCHAR(64) NOT NULL COMMENT '类型名称',
  `code` VARCHAR(32) NOT NULL COMMENT '类型编码',
  `description` VARCHAR(255) DEFAULT NULL COMMENT '描述',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-启用，2-禁用',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='证件类型表';
```

### 3.3 证件表 (certificate)
```sql
CREATE TABLE `certificate` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` BIGINT NOT NULL COMMENT '关联用户ID',
  `category_id` INT NOT NULL COMMENT '关联证件类型ID',
  `name` VARCHAR(128) NOT NULL COMMENT '证书名称',
  `cert_no` VARCHAR(64) DEFAULT NULL COMMENT '证书编号',
  `issuer` VARCHAR(128) DEFAULT NULL COMMENT '发证机构',
  `issue_date` DATE DEFAULT NULL COMMENT '发证日期',
  `expire_date` DATE DEFAULT NULL COMMENT '有效期至',
  `level` VARCHAR(32) DEFAULT NULL COMMENT '证书等级（初中高级）',
  `file_url` VARCHAR(512) NOT NULL COMMENT 'MinIO 文件URL',
  `file_type` VARCHAR(16) NOT NULL COMMENT '文件类型：image/pdf',
  `thumb_url` VARCHAR(512) DEFAULT NULL COMMENT '缩略图URL',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-正常，2-过期',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_level` (`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='证件表';
```

### 3.4 导出任务表 (export_task)
```sql
CREATE TABLE `export_task` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `task_name` VARCHAR(128) DEFAULT NULL COMMENT '任务名称',
  `user_count` INT NOT NULL DEFAULT 0 COMMENT '涉及人数',
  `cert_count` INT NOT NULL DEFAULT 0 COMMENT '匹配证件数',
  `miss_count` INT NOT NULL DEFAULT 0 COMMENT '缺证人数',
  `watermark_config` JSON COMMENT '水印配置',
  `file_url` VARCHAR(512) DEFAULT NULL COMMENT '生成的压缩包URL',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-处理中，2-完成，3-失败',
  `fail_reason` VARCHAR(512) DEFAULT NULL COMMENT '失败原因',
  `created_by` BIGINT NOT NULL COMMENT '创建人ID',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `completed_at` DATETIME DEFAULT NULL COMMENT '完成时间',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_created_by` (`created_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='导出任务表';
```

### 3.5 操作日志表 (operation_log)
```sql
CREATE TABLE `operation_log` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `operator_id` BIGINT NOT NULL COMMENT '操作人ID',
  `operator_name` VARCHAR(64) NOT NULL COMMENT '操作人姓名',
  `module` VARCHAR(32) NOT NULL COMMENT '操作模块',
  `action` VARCHAR(64) NOT NULL COMMENT '操作动作',
  `target` VARCHAR(255) DEFAULT NULL COMMENT '操作对象描述',
  `detail` JSON DEFAULT NULL COMMENT '详细数据',
  `ip` VARCHAR(64) DEFAULT NULL COMMENT '操作IP',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_operator` (`operator_id`),
  KEY `idx_module_action` (`module`, `action`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作日志表';
```

---

## 4. 后端 API 设计

### 4.1 模块与路由

```
POST   /api/sys/login                    登录
POST   /api/sys/logout                   登出
GET    /api/sys/upload-token             获取 MinIO 预签名上传URL
POST   /api/sys/clean-old-exports        手动触发清理旧导出文件

GET    /api/user/list                    用户列表
POST   /api/user/add                     新增用户
POST   /api/user/update                  更新用户
POST   /api/user/delete                  删除用户
GET    /api/user/:id/certs               查询某用户的所有证件

GET    /api/cert-category/list           证件类型列表
POST   /api/cert-category/add            新增类型
POST   /api/cert-category/update         更新类型
POST   /api/cert-category/delete         删除类型

GET    /api/certificate/list             证件列表（支持按用户/类型/等级筛选）
POST   /api/certificate/add              新增证件
POST   /api/certificate/batch-add        批量新增证件
POST   /api/certificate/update           更新证件
POST   /api/certificate/delete           删除证件
GET    /api/certificate/preview/:id      证件预览（返回文件URL）
GET    /api/certificate/thumb/:id        证件缩略图（返回缩略图URL）

POST   /api/export/import                上传 Excel 名单
POST   /api/export/preview               预览匹配结果
POST   /api/export/download              确认导出（异步模式）
GET    /api/export/task/:id              查询导出任务状态
GET    /api/export/task-list             导出历史列表
POST   /api/export/retry/:id             重试失败的导出任务

GET    /api/operation-log/list           操作日志列表
```

### 4.2 核心数据结构

#### 登录请求/响应
```go
type LoginReq struct {
    Username string `json:"username"`  // 管理员账号
    Password string `json:"password"`
}

type LoginResp struct {
    Token     string `json:"token"`
    Username  string `json:"username"`
    Role      int    `json:"role"`
    ExpireAt  int64  `json:"expireAt"`
}
```

#### 水印配置
```go
type WatermarkConfig struct {
    Text       string  `json:"text"`        // 水印文字
    Position   string  `json:"position"`    // center/diagonal/bottom/topleft/topright/bottomleft/bottomright
    FontSize   int     `json:"fontSize"`    // 字号
    Opacity    float64 `json:"opacity"`     // 透明度 0.0~1.0
    Color      string  `json:"color"`       // 十六进制颜色
    Rotation   int     `json:"rotation"`    // 旋转角度
    FontFamily string  `json:"fontFamily"`  // 字体名称
}
```

#### 预览匹配结果
```go
type ExportPreviewItem struct {
    RowIndex     int           `json:"rowIndex"`     // Excel 行号
    Name         string        `json:"name"`         // 姓名
    IdCard       string        `json:"idCard"`       // 身份证号
    Category     string        `json:"category"`     // 目标证件类型（Excel填的）
    Level        string        `json:"level"`        // 目标等级（Excel填的）
    Status       int           `json:"status"`       // 1-匹配成功, 2-缺证, 3-未匹配（无此人）
    UserId       int64         `json:"userId"`       // 系统用户ID
    Certificates []CertBrief   `json:"certificates"` // 匹配到的证件列表（Status=1时有值）
    MissReason   string        `json:"missReason"`   // 缺证原因或异常提示
}

type ExportPreviewResp struct {
    TotalCount     int                  `json:"totalCount"`
    MatchedCount   int                  `json:"matchedCount"`
    MissCount      int                  `json:"missCount"`
    UnmatchedCount int                  `json:"unmatchedCount"`
    Items          []ExportPreviewItem  `json:"items"`
    Warnings       []string             `json:"warnings"`   // 格式异常提示
}
```

#### 导出请求
```go
type ExportDownloadReq struct {
    PreviewToken string           `json:"previewToken"` // 预览令牌（避免重复解析Excel）
    SelectedIds  []int64          `json:"selectedIds"`  // 勾选的用户ID列表
    Watermark    WatermarkConfig  `json:"watermark"`    // 水印配置
    TaskName     string           `json:"taskName"`     // 任务名称（可选）
}

type ExportDownloadResp struct {
    TaskId  int64  `json:"taskId"`
    Status  int    `json:"status"`   // 1-处理中
}
```

#### 导出任务查询
```go
type ExportTaskResp struct {
    Id           int64           `json:"id"`
    TaskName     string          `json:"taskName"`
    UserCount    int             `json:"userCount"`
    CertCount    int             `json:"certCount"`
    MissCount    int             `json:"missCount"`
    Status       int             `json:"status"`       // 1-处理中, 2-完成, 3-失败
    FileUrl      string          `json:"fileUrl"`      // 完成后返回的下载链接
    FailReason   string          `json:"failReason"`
    CreatedAt    string          `json:"createdAt"`
    CompletedAt  string          `json:"completedAt"`
}
```

---

## 5. 核心业务流程

### 5.1 登录鉴权流程

```
1. 管理员输入账号密码
2. 后端校验密码（bcrypt）
3. 生成 JWT Token（存入 Redis，key: token:{user_id}）
4. 返回 Token 和用户信息
5. 前端后续请求 Header 携带 Authorization: Bearer {token}
6. 后端 JWT 中间件校验 Token + Redis 检查是否有效
```

### 5.2 Excel 导入 - 预览匹配 - 确认导出

```
+------------------+------------------+------------------+
|   Step 1: 上传   |  Step 2: 预览    | Step 3: 导出     |
+------------------+------------------+------------------+
|                  |                  |                  |
| 下载模板         | 展示匹配结果     | 配置水印         |
| 填写名单         | 勾选人员         | 提交导出任务     |
| 上传 Excel       | 查看证件明细     | 异步处理         |
|                  |                  | 轮询状态         |
| 解析 + 匹配用户   | 三种状态标识     | 下载 ZIP         |
| 自动创建缺失用户  | 绿/红/灰         |                  |
| 缓存到 Redis     |                  |                  |
| 返回预览数据     | 全选/反选        |                  |
| + warnings       |                  |                  |
+------------------+------------------+------------------+
```

**导出异步处理：**
1. 从 Redis 读取缓存的匹配结果
2. 筛选勾选的人员
3. 从 MySQL 查询这些人的证件信息
4. 从 MinIO 下载证件文件到本地临时目录
5. 遍历每份文件打水印（图片用 gg/PDF 用 pdfcpu）
6. 重命名文件：`{姓名}_{证件类型}_{等级}.{ext}`
7. 打包 ZIP -> 上传 MinIO
8. 更新任务状态
9. 清理临时文件

### 5.3 缩略图生成

```
用户上传证件图片：
1. 前端获取 MinIO 预签名 URL
2. 直传 MinIO
3. 后端收到上传完成通知
4. 读取原图 -> gg 库缩放（最大边 400px）
5. 上传缩略图到 MinIO /thumbs/
6. 更新 certificate.thumb_url
```

### 5.4 操作日志

| 模块 | 动作 | 触发时机 |
|-----|------|---------|
| user | add/update/delete | 用户增删改 |
| certificate | add/batch-add/update/delete | 证件增删改 |
| export | download | 提交导出任务时 |
| export | clean | 手动触发清理时 |

---

## 6. MinIO 存储结构

```
hridoc-bucket/
+-- certificates/             # 原始证件文件
|   +-- {user_id}/
|       +-- {cert_id}_original.jpg
|       +-- {cert_id}_original.pdf
+-- thumbs/                   # 缩略图
|   +-- {user_id}/
|       +-- {cert_id}_thumb.jpg
+-- exports/                  # 导出压缩包
    +-- export_{task_id}_{timestamp}.zip
```

### 预签名 URL 过期时间
| 场景 | 过期时间 |
|------|---------|
| 上传 URL | 5 分钟 |
| 预览 URL | 5 分钟 |
| 下载 URL | 1 小时 |
| 缩略图 URL | 5 分钟 |

---

## 7. 前端页面设计

### 7.2 页面清单

| 页面 | 路径 | 说明 |
|------|------|------|
| 登录页 | /login | 管理员账号密码登录 |
| 用户管理 | /user | 用户列表、增删改查、分页搜索 |
| 证件类型 | /category | 证件类型增删改查 |
| 证件管理 | /certificate | 证件列表、筛选、预览、上传 |
| 导入导出 | /export | Excel 上传、预览、水印配置、导出 |
| 导出历史 | /export/history | 历史任务列表、下载、重试 |
| 操作日志 | /log | 操作日志查询 |

### 7.3 关键交互

**水印配置弹窗**：
- 表单字段：文字、位置（7 个选项）、字号滑块、透明度滑块、颜色选择器、角度滑块
- 实时 Canvas 预览（前端用 Canvas 模拟，帮助 HR 调整）

**导出预览页**：
- 统计栏：总人数 / 匹配成功 / 缺证 / 未匹配
- 表格：勾选列 + 行号 + 姓名 + 身份证 + 类型 + 状态 + 证件展开
- 状态颜色：匹配=绿色 / 缺证=红色 / 未匹配=灰色
- 操作栏：全选/反选、配置水印、确认导出

**证件预览**：
- 图片：弹窗展示，支持缩放
- PDF：iframe 或新标签页打开

---

## 8. 配置文件

### api.yaml
```yaml
Name: hridoc-api
Host: 0.0.0.0
Port: 8888

Mysql:
  DataSource: root:${PASSWORD}@tcp(mysql:3306)/hridoc?charset=utf8mb4

CacheRedis:
  - Host: redis:6379
    Pass: ${PASSWORD}
    Type: node

MinIO:
  Endpoint: minio:9000
  AccessKey: minio
  SecretKey: minio123
  Bucket: hridoc-bucket
  UseSSL: false

JwtAuth:
  AccessSecret: ${JWT_SECRET}
  AccessExpire: 86400

Watermark:
  FontPath: /app/fonts/simhei.ttf
  DefaultText: "仅供公司内部使用"

Upload:
  MaxFileSize: 10485760
  AllowedTypes: ["image/jpeg", "image/png", "image/jpg", "application/pdf"]

Excel:
  MaxRows: 1000

Export:
  TempDir: /tmp/hridoc/exports
  CleanupDays: 7
```

---

## 9. 错误码

| 错误码 | 说明 |
|--------|------|
| 10001 | 系统错误 |
| 10002 | 参数错误 |
| 20001 | 用户不存在 |
| 20002 | 身份证号已存在 |
| 20003 | 密码错误 |
| 20004 | Token 无效或过期 |
| 30001 | 证件类型不存在 |
| 30002 | 证件不存在 |
| 40001 | Excel 格式错误 |
| 40002 | Excel 内容验证失败 |
| 40003 | 预览令牌过期 |
| 40004 | 导出任务不存在 |
| 40005 | 导出任务失败 |
| 50001 | 文件上传失败 |
| 50002 | 不支持的文件类型 |
| 50003 | 文件大小超限 |
| 50004 | MinIO 操作失败 |
| 60001 | 无权限 |

---

## 10. Docker Compose 完整配置

```yaml
version: "3.8"

services:
  mysql:
    image: mysql:5.7
    container_name: hridoc-mysql
    ports:
      - "13306:3306"
    volumes:
      - ${DATA_DIR}/mysql/data:/var/lib/mysql
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql:ro
    environment:
      MYSQL_ROOT_PASSWORD: ${PASSWORD}
      MYSQL_DATABASE: hridoc
    restart: always

  redis:
    image: redis:7-alpine
    container_name: hridoc-redis
    ports:
      - "16379:6379"
    environment:
      TZ: Asia/Shanghai
    restart: always
    command: redis-server --requirepass ${PASSWORD} --appendonly yes

  minio:
    image: minio/minio:latest
    container_name: hridoc-minio
    ports:
      - "19000:9000"
      - "19001:9001"
    volumes:
      - ${DATA_DIR}/minio:/data
    environment:
      MINIO_ROOT_USER: minio
      MINIO_ROOT_PASSWORD: minio123
    restart: always
    command: server /data --console-address ":9001"

  api:
    build:
      context: ./api
      dockerfile: Dockerfile
    container_name: hridoc-api
    ports:
      - "8888:8888"
    volumes:
      - ./api/etc/api.yaml:/app/etc/api.yaml:ro
      - ${DATA_DIR}/fonts:/app/fonts:ro
      - ${DATA_DIR}/tmp:/tmp/hridoc
    environment:
      TZ: Asia/Shanghai
    depends_on:
      - mysql
      - redis
      - minio
    restart: always

  frontend:
    build:
      context: ./web
      dockerfile: Dockerfile
    container_name: hridoc-web
    ports:
      - "80:80"
    restart: always
```

---

## 11. 数据初始化

### 超级管理员
```sql
INSERT INTO `user` (name, phone, id_card, role, status, password)
VALUES ('超级管理员', '13800000000', '000000000000000000', 1, 1, '$2a$10$xxxx');
```

### 默认证件类型
```sql
INSERT INTO `cert_category` (name, code, description, status) VALUES
('软考证书', 'RKS', '计算机技术与软件专业技术资格考试', 1),
('一级建造师', 'YJJZS', '一级建造师执业资格证书', 1),
('二级建造师', 'EJJZS', '二级建造师执业资格证书', 1),
('注册安全工程师', 'ZCAQGCS', '注册安全工程师证书', 1),
('PMP项目管理', 'PMP', '项目管理专业人士资格认证', 1);
```

---

## 12. 实现阶段

### 阶段一：基础框架
- go-zero API 项目初始化
- Docker Compose 配置（含 MinIO）
- 数据库表 + 初始数据
- JWT 鉴权 + 操作日志中间件
- MinIO 上传封装

### 阶段二：用户与证件管理
- 用户管理 CRUD
- 证件类型管理 CRUD
- 证件管理 CRUD + 上传 + 缩略图
- 文件预览接口

### 阶段三：导入导出核心
- Excel 模板下载 + 解析
- 预览匹配逻辑（含自动创建用户）
- 水印生成（图片 + PDF）
- 异步导出任务（ZIP 打包 + MinIO 上传）
- 导出任务状态轮询
- 导出历史列表

### 阶段四：前端
- Vue3 项目初始化
- 登录页 + 路由守卫
- 用户管理 / 证件类型 / 证件管理
- 导入导出页（水印配置 + 预览表格）
- 导出历史页 + 操作日志页

### 阶段五：收尾
- 操作日志完整接入
- 手动清理功能
- 测试与联调
- 部署文档

---

*文档版本: v1.0*  
*最后更新: 2026-06-01*
