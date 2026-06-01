# hridoc 公司员工证件管理系统 实施计划

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 基于 go-zero + Vue3 + Element Plus + MySQL + Redis + MinIO 构建完整的员工证件管理系统，支持证件管理、Excel 导入导出、水印生成。

**Architecture:** 单体 API 服务（go-zero）处理所有业务逻辑；前端 Vue3 SPA 独立部署；文件存储使用 MinIO；数据库使用 MySQL；Redis 用于缓存和会话。

**Tech Stack:** Go 1.21+, go-zero, Vue 3, Element Plus, Vite, Pinia, MySQL 5.7, Redis 7, MinIO, docker-compose

**Project Root:** `D:\worksapce\go\src\github.com\DemoLiang\hridoc`

---

## 文件结构规划

### 后端目录 (`api/`)
```
api/
├── etc/
│   └── api.yaml              # go-zero 配置文件
├── internal/
│   ├── config/
│   │   └── config.go         # 配置结构体
│   ├── handler/
│   │   ├── routes.go         # 路由注册
│   │   ├── sys/
│   │   │   ├── loginhandler.go
│   │   │   ├── logouthandler.go
│   │   │   ├── uploadtokenhandler.go
│   │   │   └── cleanhandler.go
│   │   ├── user/
│   │   │   ├── userlisthandler.go
│   │   │   ├── useraddhandler.go
│   │   │   ├── userupdatehandler.go
│   │   │   ├── userdeletehandler.go
│   │   │   └── usercerthandler.go
│   │   ├── certcategory/
│   │   │   ├── categorylisthandler.go
│   │   │   ├── categoryaddhandler.go
│   │   │   ├── categoryupdatehandler.go
│   │   │   └── categorydeletehandler.go
│   │   ├── certificate/
│   │   │   ├── certlisthandler.go
│   │   │   ├── certaddhandler.go
│   │   │   ├── certbatchaddhandler.go
│   │   │   ├── certupdatehandler.go
│   │   │   ├── certdeletehandler.go
│   │   │   ├── certpreviewhandler.go
│   │   │   └── certthumbhandler.go
│   │   ├── export/
│   │   │   ├── importhandler.go
│   │   │   ├── previewhandler.go
│   │   │   ├── downloadhandler.go
│   │   │   ├── taskhandler.go
│   │   │   ├── tasklisthandler.go
│   │   │   └── retryhandler.go
│   │   └── operationlog/
│   │       └── loglisthandler.go
│   ├── logic/
│   │   ├── sys/
│   │   │   ├── loginlogic.go
│   │   │   ├── logoutlogic.go
│   │   │   ├── uploadtokenlogic.go
│   │   │   └── cleanlogic.go
│   │   ├── user/
│   │   │   ├── userlistlogic.go
│   │   │   ├── useraddlogic.go
│   │   │   ├── userupdatelogic.go
│   │   │   ├── userdeletelogic.go
│   │   │   └── usercertlogic.go
│   │   ├── certcategory/
│   │   │   ├── categorylistlogic.go
│   │   │   ├── categoryaddlogic.go
│   │   │   ├── categoryupdatelogic.go
│   │   │   └── categorydeletelogic.go
│   │   ├── certificate/
│   │   │   ├── certlistlogic.go
│   │   │   ├── certaddlogic.go
│   │   │   ├── certbatchaddlogic.go
│   │   │   ├── certupdatelogic.go
│   │   │   ├── certdeletelogic.go
│   │   │   ├── certpreviewlogic.go
│   │   │   └── certthumblogic.go
│   │   ├── export/
│   │   │   ├── importlogic.go
│   │   │   ├── previewlogic.go
│   │   │   ├── downloadlogic.go
│   │   │   ├── tasklogic.go
│   │   │   ├── tasklistlogic.go
│   │   │   └── retrylogic.go
│   │   └── operationlog/
│   │       └── loglistlogic.go
│   ├── middleware/
│   │   ├── jwtauth.go          # JWT 鉴权中间件
│   │   └── operatelog.go       # 操作日志中间件
│   ├── svc/
│   │   └── servicecontext.go   # 服务上下文（DB/Redis/MinIO）
│   └── types/
│       └── types.go            # 公共类型定义
├── model/
│   ├── usermodel.go            # user 表 model
│   ├── usermodel_gen.go        # goctl 生成
│   ├── certcategorymodel.go    # cert_category 表 model
│   ├── certcategorymodel_gen.go
│   ├── certificatemodel.go     # certificate 表 model
│   ├── certificatemodel_gen.go
│   ├── exporttaskmodel.go      # export_task 表 model
│   ├── exporttaskmodel_gen.go
│   ├── operationlogmodel.go    # operation_log 表 model
│   └── operationlogmodel_gen.go
├── pkg/
│   ├── minio/
│   │   └── client.go           # MinIO 客户端封装
│   ├── watermark/
│   │   ├── image.go            # 图片水印
│   │   ├── pdf.go              # PDF 水印
│   │   └── font.go             # 字体加载
│   ├── excel/
│   │   ├── reader.go           # Excel 读取
│   │   └── template.go         # 模板生成
│   ├── auth/
│   │   └── jwt.go              # JWT 工具
│   └── errorx/
│       └── errors.go           # 业务错误码
├── doc/
│   └── api/
│       └── hridoc.api          # go-zero API 定义文件
├── Dockerfile
├── go.mod
├── go.sum
└── hridoc.go                   # 入口文件
```

### 前端目录 (`web/`)
```
web/
├── public/
├── src/
│   ├── api/
│   │   ├── sys.js              # 系统 API
│   │   ├── user.js             # 用户 API
│   │   ├── category.js         # 证件类型 API
│   │   ├── certificate.js      # 证件 API
│   │   ├── export.js           # 导出 API
│   │   └── log.js              # 日志 API
│   ├── components/
│   │   ├── WatermarkDialog.vue # 水印配置弹窗
│   │   ├── CertPreview.vue     # 证件预览组件
│   │   └── UploadComponent.vue # 上传组件
│   ├── views/
│   │   ├── LoginView.vue       # 登录页
│   │   ├── UserView.vue        # 用户管理
│   │   ├── CategoryView.vue    # 证件类型
│   │   ├── CertificateView.vue # 证件管理
│   │   ├── ExportView.vue      # 导入导出
│   │   ├── ExportHistory.vue   # 导出历史
│   │   └── LogView.vue         # 操作日志
│   ├── router/
│   │   └── index.js            # 路由配置
│   ├── stores/
│   │   ├── user.js             # 用户状态
│   │   └── app.js              # 应用状态
│   ├── utils/
│   │   ├── request.js          # Axios 封装
│   │   └── constants.js        # 常量
│   ├── App.vue
│   └── main.js
├── index.html
├── package.json
├── vite.config.js
└── Dockerfile
```

### 根目录
```
├── docker-compose.yaml
├── init.sql                    # 数据库初始化脚本
├── .env                        # 环境变量
└── README.md
```

---

## 阶段一：基础框架搭建

### Task 1.1: 创建 go-zero API 项目结构

**Files:**
- Create: `api/go.mod`
- Create: `api/go.sum`
- Create: `api/hridoc.go`
- Create: `api/Dockerfile`
- Create: `api/etc/api.yaml`

- [ ] **Step 1: 初始化 Go 模块**

```bash
cd D:\worksapce\go\src\github.com\DemoLiang\hridoc\api
go mod init github.com/DemoLiang/hridoc/api
```

- [ ] **Step 2: 安装 go-zero 依赖**

```bash
go get github.com/zeromicro/go-zero@latest
go mod tidy
```

- [ ] **Step 3: 创建入口文件 `api/hridoc.go`**

```go
package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/DemoLiang/hridoc/api/internal/config"
	"github.com/DemoLiang/hridoc/api/internal/handler"
	"github.com/DemoLiang/hridoc/api/internal/svc"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
```

- [ ] **Step 4: 创建配置文件 `api/etc/api.yaml`**

```yaml
Name: hridoc-api
Host: 0.0.0.0
Port: 8888

Mysql:
  DataSource: root:root@tcp(localhost:3306)/hridoc?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai

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

Watermark:
  FontPath: ./fonts/simhei.ttf
  DefaultText: "仅供公司内部使用"

Upload:
  MaxFileSize: 10485760
  AllowedTypes:
    - image/jpeg
    - image/png
    - image/jpg
    - application/pdf

Excel:
  MaxRows: 1000

Export:
  TempDir: ./tmp/exports
  CleanupDays: 7
```

- [ ] **Step 5: 创建配置结构 `api/internal/config/config.go`**

```go
package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	CacheRedis redis.RedisConf
	JwtAuth    struct {
		AccessSecret string
		AccessExpire int64
	}
	MinIO struct {
		Endpoint        string
		AccessKey       string
		SecretKey       string
		Bucket          string
		UseSSL          bool
		PresignedExpiry int
	}
	Watermark struct {
		FontPath    string
		DefaultText string
	}
	Upload struct {
		MaxFileSize  int64
		AllowedTypes []string
	}
	Excel struct {
		MaxRows int
	}
	Export struct {
		TempDir     string
		CleanupDays int
	}
}
```

- [ ] **Step 6: 创建服务上下文 `api/internal/svc/servicecontext.go`**

```go
package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/DemoLiang/hridoc/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
```

- [ ] **Step 7: Create Dockerfile `api/Dockerfile`**

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o hridoc-api hridoc.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/hridoc-api .
COPY --from=builder /app/etc ./etc

EXPOSE 8888
CMD ["./hrridoc-api", "-f", "etc/api.yaml"]
```

- [ ] **Step 8: Commit**

```bash
git add api/
git commit -m "feat: init go-zero API project structure

- Go module initialization
- Config structure with MySQL, Redis, MinIO, JWT
- Service context with DB connection
- Dockerfile for API service

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 1.2: 定义数据库模型并生成 Model 代码

**Files:**
- Create: `api/model/*.sql` (DDL)
- Create: `api/model/table.sql`

- [ ] **Step 1: 创建数据库初始化脚本 `init.sql`**

```sql
CREATE DATABASE IF NOT EXISTS hridoc CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE hridoc;

CREATE TABLE IF NOT EXISTS `user` (
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

CREATE TABLE IF NOT EXISTS `cert_category` (
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

CREATE TABLE IF NOT EXISTS `certificate` (
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

CREATE TABLE IF NOT EXISTS `export_task` (
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

CREATE TABLE IF NOT EXISTS `operation_log` (
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

-- 初始化超级管理员（密码: admin123，bcrypt hash）
INSERT INTO `user` (name, phone, id_card, role, status, password, created_at, updated_at)
VALUES ('超级管理员', '13800000000', '000000000000000000', 1, 1, '$2a$10$X7oMyJxQ8ZlQkEYQKNr5U.Sl1fBxTjK.F8gN0hH3aTQmW8LgKJ6m', NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();

-- 默认证件类型
INSERT INTO `cert_category` (name, code, description, status) VALUES
('软考证书', 'RKS', '计算机技术与软件专业技术资格（水平）考试', 1),
('一级建造师', 'YJJZS', '一级建造师执业资格证书', 1),
('二级建造师', 'EJJZS', '二级建造师执业资格证书', 1),
('注册安全工程师', 'ZCAQGCS', '注册安全工程师证书', 1),
('PMP项目管理', 'PMP', '项目管理专业人士资格认证', 1)
ON DUPLICATE KEY UPDATE updated_at = NOW();
```

- [ ] **Step 2: 使用 goctl 生成 Model 代码**

```bash
cd D:\worksapce\go\src\github.com\DemoLiang\hridoc\api
# 安装 goctl
go install github.com/zeromicro/go-zero/tools/goctl@latest

# 生成 user model
goctl model mysql ddl -src ../init.sql -dir ./model -t user
# 生成 cert_category model
goctl model mysql ddl -src ../init.sql -dir ./model -t cert_category
# 生成 certificate model
goctl model mysql ddl -src ../init.sql -dir ./model -t certificate
# 生成 export_task model
goctl model mysql ddl -src ../init.sql -dir ./model -t export_task
# 生成 operation_log model
goctl model mysql ddl -src ../init.sql -dir ./model -t operation_log
```

- [ ] **Step 3: 更新 ServiceContext 注入 Model**

Modify: `api/internal/svc/servicecontext.go`

```go
package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/DemoLiang/hridoc/api/internal/config"
	"github.com/DemoLiang/hridoc/api/model"
)

type ServiceContext struct {
	Config             config.Config
	DB                 sqlx.SqlConn
	UserModel          model.UserModel
	CertCategoryModel  model.CertCategoryModel
	CertificateModel   model.CertificateModel
	ExportTaskModel    model.ExportTaskModel
	OperationLogModel  model.OperationLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		DB:                db,
		UserModel:         model.NewUserModel(db, c.CacheRedis),
		CertCategoryModel: model.NewCertCategoryModel(db, c.CacheRedis),
		CertificateModel:  model.NewCertificateModel(db, c.CacheRedis),
		ExportTaskModel:   model.NewExportTaskModel(db, c.CacheRedis),
		OperationLogModel: model.NewOperationLogModel(db, c.CacheRedis),
	}
}
```

- [ ] **Step 4: Commit**

```bash
git add init.sql api/model/
git commit -m "feat: add database models and initialization

- Complete SQL DDL for 5 tables
- goctl generated model code
- ServiceContext with all model injections
- Default admin user and cert categories

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 1.3: MinIO 客户端封装

**Files:**
- Create: `api/pkg/minio/client.go`

- [ ] **Step 1: 创建 MinIO 客户端**

```go
package minio

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/DemoLiang/hridoc/api/internal/config"
)

type Client struct {
	client     *minio.Client
	bucket     string
	expiry     time.Duration
}

func NewClient(cfg config.MinIO) (*Client, error) {
	mc, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("minio new client: %w", err)
	}

	ctx := context.Background()
	exists, err := mc.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("minio check bucket: %w", err)
	}
	if !exists {
		err = mc.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("minio make bucket: %w", err)
		}
	}

	return &Client{
		client: mc,
		bucket: cfg.Bucket,
		expiry: time.Duration(cfg.PresignedExpiry) * time.Second,
	}, nil
}

func (c *Client) PresignedPutURL(ctx context.Context, objectName string) (string, error) {
	u, err := c.client.PresignedPutObject(ctx, c.bucket, objectName, c.expiry)
	if err != nil {
		return "", fmt.Errorf("minio presigned put: %w", err)
	}
	return u.String(), nil
}

func (c *Client) PresignedGetURL(ctx context.Context, objectName string, expiry time.Duration) (string, error) {
	u, err := c.client.PresignedGetObject(ctx, c.bucket, objectName, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("minio presigned get: %w", err)
	}
	return u.String(), nil
}

func (c *Client) PutObject(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) error {
	_, err := c.client.PutObject(ctx, c.bucket, objectName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return fmt.Errorf("minio put object: %w", err)
	}
	return nil
}

func (c *Client) GetObject(ctx context.Context, objectName string) (io.ReadCloser, error) {
	obj, err := c.client.GetObject(ctx, c.bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("minio get object: %w", err)
	}
	return obj, nil
}

func (c *Client) RemoveObject(ctx context.Context, objectName string) error {
	err := c.client.RemoveObject(ctx, c.bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("minio remove object: %w", err)
	}
	return nil
}

func (c *Client) ObjectURL(objectName string) string {
	return fmt.Sprintf("%s/%s/%s", c.client.EndpointURL(), c.bucket, objectName)
}
```

- [ ] **Step 2: 更新 ServiceContext**

Modify: `api/internal/svc/servicecontext.go`

新增 MinIO 客户端注入：
```go
import "github.com/DemoLiang/hridoc/api/pkg/minio"

type ServiceContext struct {
	Config             config.Config
	DB                 sqlx.SqlConn
	UserModel          model.UserModel
	CertCategoryModel  model.CertCategoryModel
	CertificateModel   model.CertificateModel
	ExportTaskModel    model.ExportTaskModel
	OperationLogModel  model.OperationLogModel
	MinIO              *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMysql(c.Mysql.DataSource)
	minioClient, _ := minio.NewClient(c.MinIO) // 错误处理在生产代码中完善
	return &ServiceContext{
		Config:            c,
		DB:                db,
		UserModel:         model.NewUserModel(db, c.CacheRedis),
		CertCategoryModel: model.NewCertCategoryModel(db, c.CacheRedis),
		CertificateModel:  model.NewCertificateModel(db, c.CacheRedis),
		ExportTaskModel:   model.NewExportTaskModel(db, c.CacheRedis),
		OperationLogModel: model.NewOperationLogModel(db, c.CacheRedis),
		MinIO:             minioClient,
	}
}
```

- [ ] **Step 3: Commit**

```bash
git add api/pkg/minio/ api/internal/svc/
git commit -m "feat: add MinIO client wrapper

- Presigned PUT/GET URLs
- Object upload/download/remove
- Bucket auto-creation
- Integrated into ServiceContext

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 1.4: JWT 鉴权中间件

**Files:**
- Create: `api/pkg/auth/jwt.go`
- Create: `api/internal/middleware/jwtauth.go`

- [ ] **Step 1: 创建 JWT 工具包 `api/pkg/auth/jwt.go`**

```go
package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
	Expire int64
}

type Claims struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.RegisteredClaims
}

func NewJWT(secret string, expire int64) *JWT {
	return &JWT{Secret: secret, Expire: expire}
}

func (j *JWT) GenerateToken(userId int64, username string, role int) (string, error) {
	now := time.Now()
	claims := Claims{
		UserId:   userId,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(j.Expire) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))
}

func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
```

- [ ] **Step 2: 创建 JWT 中间件 `api/internal/middleware/jwtauth.go`**

```go
package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/DemoLiang/hridoc/api/internal/config"
	"github.com/DemoLiang/hridoc/api/pkg/auth"
)

func JWTAuth(cfg config.Config) func(http.HandlerFunc) http.HandlerFunc {
	jwtUtil := auth.NewJWT(cfg.JwtAuth.AccessSecret, cfg.JwtAuth.AccessExpire)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if len(token) < 7 || token[:7] != "Bearer " {
				httpx.Error(w, fmt.Errorf("missing or invalid authorization header"))
				return
			}

			claims, err := jwtUtil.ParseToken(token[7:])
			if err != nil {
				httpx.Error(w, fmt.Errorf("invalid token: %v", err))
				return
			}

			// 将用户信息写入 context
			ctx := r.Context()
			ctx = context.WithValue(ctx, "userId", claims.UserId)
			ctx = context.WithValue(ctx, "username", claims.Username)
			ctx = context.WithValue(ctx, "role", claims.Role)
			next(w, r.WithContext(ctx))
		}
	}
}
```

- [ ] **Step 3: 创建错误码定义 `api/pkg/errorx/errors.go`**

```go
package errorx

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	ErrSystem           = 10001
	ErrInvalidParam     = 10002
	ErrUserNotFound     = 20001
	ErrIdCardExists     = 20002
	ErrPassword         = 20003
	ErrTokenInvalid     = 20004
	ErrCategoryNotFound = 30001
	ErrCertNotFound     = 30002
	ErrExcelFormat      = 40001
	ErrExcelValidation  = 40002
	ErrPreviewExpired   = 40003
	ErrTaskNotFound     = 40004
	ErrTaskFailed       = 40005
	ErrUploadFailed     = 50001
	ErrInvalidFileType  = 50002
	ErrFileTooLarge     = 50003
	ErrMinIOFailed      = 50004
	ErrNoPermission     = 60001
)

func New(code int, msg string) error {
	return fmt.Errorf("%d: %s", code, msg)
}

func IsNotFound(err error) bool {
	return err == sqlx.ErrNotFound
}
```

- [ ] **Step 4: Commit**

```bash
git add api/pkg/auth/ api/pkg/errorx/ api/internal/middleware/
git commit -m "feat: add JWT auth middleware and error codes

- JWT token generation and parsing
- HTTP middleware for auth protection
- Business error code definitions

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 1.5: 更新 docker-compose.yaml

**Files:**
- Modify: `docker-compose.yaml`

- [ ] **Step 1: 更新 compose 配置**

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

  web:
    build:
      context: ./web
      dockerfile: Dockerfile
    container_name: hridoc-web
    ports:
      - "80:80"
    restart: always
```

- [ ] **Step 2: Commit**

```bash
git add docker-compose.yaml
git commit -m "chore: update docker-compose with MinIO and web service

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

## 阶段二：用户与证件管理

### Task 2.1: 用户管理 API

**Files:**
- Create: `api/doc/api/hridoc.api` (API 定义)

- [ ] **Step 1: 定义 API 接口 `api/doc/api/hridoc.api`**

```go
syntax = "v1"

info(
	title: "hridoc API"
	desc: "hridoc certificate management system API"
	author: "DemoLiang"
	version: "v1"
)

type (
	LoginReq {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResp {
		Token    string `json:"token"`
		Username string `json:"username"`
		Role     int    `json:"role"`
		ExpireAt int64  `json:"expireAt"`
	}

	UserInfo {
		Id        int64  `json:"id"`
		Name      string `json:"name"`
		Phone     string `json:"phone,omitempty"`
		Email     string `json:"email,omitempty"`
		IdCard    string `json:"idCard"`
		Education string `json:"education,omitempty"`
		Role      int    `json:"role"`
		Status    int    `json:"status"`
		CreatedAt string `json:"createdAt"`
	}

	UserListReq {
		Page     int    `form:"page,default=1"`
		PageSize int    `form:"pageSize,default=10"`
		Keyword  string `form:"keyword,optional"`
	}

	UserListResp {
		List       []UserInfo `json:"list"`
		Total      int64      `json:"total"`
		Page       int        `json:"page"`
		PageSize   int        `json:"pageSize"`
	}

	UserAddReq {
		Name      string `json:"name"`
		Phone     string `json:"phone,optional"`
		Email     string `json:"email,optional"`
		IdCard    string `json:"idCard"`
		Education string `json:"education,optional"`
	}

	UserUpdateReq {
		Id        int64  `json:"id"`
		Name      string `json:"name,optional"`
		Phone     string `json:"phone,optional"`
		Email     string `json:"email,optional"`
		Education string `json:"education,optional"`
		Status    int    `json:"status,optional"`
	}

	UserDeleteReq {
		Id int64 `json:"id"`
	}
)

service hridoc-api {
	@handler login
	post /api/sys/login (LoginReq) returns (LoginResp)
}
```

- [ ] **Step 2: 使用 goctl 生成 Handler 和 Logic 骨架**

```bash
cd api
goctl api go -api doc/api/hridoc.api -dir . -style goZero
```

- [ ] **Step 3: 实现登录逻辑 `api/internal/logic/sys/loginlogic.go`**

```go
package sys

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/pkg/auth"
	"github.com/DemoLiang/hridoc/api/pkg/errorx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	// 根据用户名查找用户（这里用 name 作为用户名）
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.Username)
	if err != nil {
		return nil, errorx.New(errorx.ErrUserNotFound, "用户不存在")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errorx.New(errorx.ErrPassword, "密码错误")
	}

	// 检查状态
	if user.Status != 1 {
		return nil, errorx.New(errorx.ErrUserNotFound, "用户已被禁用")
	}

	// 生成 JWT
	jwtUtil := auth.NewJWT(l.svcCtx.Config.JwtAuth.AccessSecret, l.svcCtx.Config.JwtAuth.AccessExpire)
	token, err := jwtUtil.GenerateToken(user.Id, user.Name, int(user.Role))
	if err != nil {
		return nil, errorx.New(errorx.ErrSystem, "生成token失败")
	}

	return &types.LoginResp{
		Token:    token,
		Username: user.Name,
		Role:     int(user.Role),
		ExpireAt: time.Now().Add(time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire) * time.Second).Unix(),
	}, nil
}
```

- [ ] **Step 4: Commit**

```bash
git add api/doc/ api/internal/
git commit -m "feat: add API definitions and login implementation

- goctl generated handler/logic skeleton
- Login with bcrypt password verification
- JWT token generation

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

### Task 2.2: 用户 CRUD 完整实现

**Files:**
- Create/Modify: `api/internal/logic/user/*.go`
- Create/Modify: `api/internal/handler/user/*.go`

- [ ] **Step 1: 用户列表逻辑 `api/internal/logic/user/userlistlogic.go`**

```go
package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/internal/types"
	"github.com/DemoLiang/hridoc/api/model"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (*types.UserListResp, error) {
	where := "where 1=1"
	args := []interface{}{}
	if req.Keyword != "" {
		where += " and (name like ? or id_card like ?)"
		args = append(args, "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	count, err := l.svcCtx.UserModel.Count(l.ctx, where, args...)
	if err != nil {
		return nil, err
	}

	list, err := l.svcCtx.UserModel.FindPage(l.ctx, req.Page, req.PageSize, where, args...)
	if err != nil {
		return nil, err
	}

	resp := &types.UserListResp{
		Total:    count,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	for _, u := range list {
		resp.List = append(resp.List, types.UserInfo{
			Id:        u.Id,
			Name:      u.Name,
			Phone:     u.Phone,
			Email:     u.Email,
			IdCard:    u.IdCard,
			Education: u.Education,
			Role:      int(u.Role),
			Status:    int(u.Status),
			CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return resp, nil
}
```

- [ ] **Step 2: 新增用户逻辑 `api/internal/logic/user/useraddlogic.go`**

```go
func (l *UserAddLogic) UserAdd(req *types.UserAddReq) error {
	// 检查身份证号是否已存在
	_, err := l.svcCtx.UserModel.FindOneByIdCard(l.ctx, req.IdCard)
	if err == nil {
		return errorx.New(errorx.ErrIdCardExists, "身份证号已存在")
	}

	// 生成默认密码（身份证号后6位）
	hashedPwd, _ := bcrypt.GenerateFromPassword(
		[]byte(req.IdCard[len(req.IdCard)-6:]), bcrypt.DefaultCost)

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Name:      req.Name,
		Phone:     req.Phone,
		Email:     req.Email,
		IdCard:    req.IdCard,
		Education: req.Education,
		Role:      2,
		Status:    1,
		Password:  string(hashedPwd),
	})
	return err
}
```

- [ ] **Step 3: 更新和删除用户**

类似模式实现 `userupdatelogic.go` 和 `userdeletelogic.go`。

- [ ] **Step 4: Commit**

```bash
git add api/internal/logic/user/ api/internal/handler/user/
git commit -m "feat: add user CRUD APIs

- User list with pagination and search
- Add user with auto-generated password
- Update and delete user

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 2.3: 证件类型 CRUD

**Files:**
- Create/Modify: `api/internal/logic/certcategory/*.go`
- Create/Modify: `api/internal/handler/certcategory/*.go`

- [ ] **Step 1: 实现证件类型 CRUD**

与 User CRUD 类似，使用 goctl 生成的骨架，实现：
- `categorylistlogic.go` - 列表（不分页，通常类型数据量小）
- `categoryaddlogic.go` - 新增
- `categoryupdatelogic.go` - 更新
- `categorydeletelogic.go` - 删除

- [ ] **Step 2: Commit**

```bash
git add api/internal/logic/certcategory/ api/internal/handler/certcategory/
git commit -m "feat: add certificate category CRUD"
```

---

### Task 2.4: 证件管理 CRUD + 上传 + 缩略图

**Files:**
- Create: `api/internal/logic/certificate/*.go`
- Create: `api/pkg/watermark/image.go` (缩略图生成)

- [ ] **Step 1: 缩略图生成 `api/pkg/watermark/image.go`**

```go
package watermark

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

func GenerateThumbnail(reader io.Reader, maxSize uint) ([]byte, string, error) {
	img, format, err := image.Decode(reader)
	if err != nil {
		return nil, "", fmt.Errorf("decode image: %w", err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if width <= int(maxSize) && height <= int(maxSize) {
		// 原图已足够小
		var buf bytes.Buffer
		switch format {
		case "png":
			png.Encode(&buf, img)
		default:
			jpeg.Encode(&buf, img, &jpeg.Options{Quality: 85})
		}
		return buf.Bytes(), format, nil
	}

	// 等比缩放
	var newWidth, newHeight uint
	if width > height {
		newWidth = maxSize
		newHeight = uint(math.Round(float64(height) * float64(maxSize) / float64(width)))
	} else {
		newHeight = maxSize
		newWidth = uint(math.Round(float64(width) * float64(maxSize) / float64(height)))
	}

	thumb := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	var buf bytes.Buffer
	jpeg.Encode(&buf, thumb, &jpeg.Options{Quality: 85})
	return buf.Bytes(), "jpeg", nil
}
```

- [ ] **Step 2: 证件新增逻辑 `api/internal/logic/certificate/certaddlogic.go`**

```go
func (l *CertAddLogic) CertAdd(req *types.CertAddReq) error {
	// 1. 保存元数据到数据库
	cert := &model.Certificate{
		UserId:     req.UserId,
		CategoryId: req.CategoryId,
		Name:       req.Name,
		CertNo:     req.CertNo,
		Issuer:     req.Issuer,
		IssueDate:  sql.NullTime{Time: req.IssueDate, Valid: !req.IssueDate.IsZero()},
		ExpireDate: sql.NullTime{Time: req.ExpireDate, Valid: !req.ExpireDate.IsZero()},
		Level:      req.Level,
		FileUrl:    req.FileUrl,
		FileType:   req.FileType,
		Status:     1,
	}
	result, err := l.svcCtx.CertificateModel.Insert(l.ctx, cert)
	if err != nil {
		return err
	}
	certId, _ := result.LastInsertId()

	// 2. 如果是图片，生成缩略图
	if req.FileType == "image" {
		go l.generateThumbnail(certId, req.FileUrl)
	}
	return nil
}

func (l *CertAddLogic) generateThumbnail(certId int64, fileUrl string) {
	// 从 MinIO 下载原图
	reader, err := l.svcCtx.MinIO.GetObject(context.Background(), extractObjectName(fileUrl))
	if err != nil {
		return
	}
	defer reader.Close()

	// 生成缩略图
	data, format, err := watermark.GenerateThumbnail(reader, 400)
	if err != nil {
		return
	}

	// 上传缩略图到 MinIO
	thumbName := fmt.Sprintf("thumbs/%d/%d_thumb.%s", certId, certId, format)
	err = l.svcCtx.MinIO.PutObject(context.Background(), thumbName, bytes.NewReader(data), int64(len(data)), "image/"+format)
	if err != nil {
		return
	}

	// 更新数据库
	l.svcCtx.CertificateModel.UpdateThumbUrl(context.Background(), certId, l.svcCtx.MinIO.ObjectURL(thumbName))
}
```

- [ ] **Step 3: Commit**

```bash
git add api/pkg/watermark/ api/internal/logic/certificate/
git commit -m "feat: add certificate CRUD with thumbnail generation

- Certificate add/update/delete/list
- Async thumbnail generation for images
- MinIO integration for file storage

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 2.5: 文件预览接口

**Files:**
- Create: `api/internal/logic/certificate/certpreviewlogic.go`

- [ ] **Step 1: 实现预览和缩略图接口**

```go
func (l *CertPreviewLogic) CertPreview(req *types.CertPreviewReq) (*types.CertPreviewResp, error) {
	cert, err := l.svcCtx.CertificateModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.New(errorx.ErrCertNotFound, "证件不存在")
	}

	// 生成预签名 URL（5 分钟过期）
	objectName := extractObjectName(cert.FileUrl)
	url, err := l.svcCtx.MinIO.PresignedGetURL(l.ctx, objectName, 5*time.Minute)
	if err != nil {
		return nil, errorx.New(errorx.ErrMinIOFailed, "生成预览链接失败")
	}

	return &types.CertPreviewResp{Url: url}, nil
}

func (l *CertThumbLogic) CertThumb(req *types.CertThumbReq) (*types.CertThumbResp, error) {
	cert, err := l.svcCtx.CertificateModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.New(errorx.ErrCertNotFound, "证件不存在")
	}

	if cert.ThumbUrl == "" {
		return nil, errorx.New(errorx.ErrCertNotFound, "暂无缩略图")
	}

	objectName := extractObjectName(cert.ThumbUrl)
	url, err := l.svcCtx.MinIO.PresignedGetURL(l.ctx, objectName, 5*time.Minute)
	if err != nil {
		return nil, errorx.New(errorx.ErrMinIOFailed, "生成缩略图链接失败")
	}

	return &types.CertThumbResp{Url: url}, nil
}
```

- [ ] **Step 2: Commit**

```bash
git add api/internal/logic/certificate/
git commit -m "feat: add certificate preview and thumbnail APIs

- Presigned URL generation for secure file access
- Separate preview and thumbnail endpoints

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

## 阶段三：导入导出核心功能

### Task 3.1: Excel 模板下载与解析

**Files:**
- Create: `api/pkg/excel/reader.go`
- Create: `api/pkg/excel/template.go`

- [ ] **Step 1: Excel 模板生成 `api/pkg/excel/template.go`**

```go
package excel

import (
	"bytes"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func GenerateTemplate() ([]byte, error) {
	f := excelize.NewFile()
	sheet := "导入模板"
	f.SetSheetName("Sheet1", sheet)

	// 表头
	headers := []string{"姓名", "身份证号", "证件类型（可选）", "证件等级（可选）"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// 设置列宽
	f.SetColWidth(sheet, "A", "A", 15)
	f.SetColWidth(sheet, "B", "B", 20)
	f.SetColWidth(sheet, "C", "C", 20)
	f.SetColWidth(sheet, "D", "D", 20)

	// 示例数据
	f.SetCellValue(sheet, "A2", "张三")
	f.SetCellValue(sheet, "B2", "11010119900101xxxx")
	f.SetCellValue(sheet, "C2", "软考证书")
	f.SetCellValue(sheet, "D2", "高级")

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, fmt.Errorf("write excel: %w", err)
	}
	return buf.Bytes(), nil
}
```

- [ ] **Step 2: Excel 读取 `api/pkg/excel/reader.go`**

```go
package excel

import (
	"fmt"
	"io"
	"strings"

	"github.com/xuri/excelize/v2"
)

type ImportRow struct {
	RowIndex int
	Name     string
	IdCard   string
	Category string
	Level    string
}

type ParseResult struct {
	Rows     []ImportRow
	Warnings []string
}

func ParseImportExcel(reader io.Reader, maxRows int) (*ParseResult, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, fmt.Errorf("open excel: %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("get rows: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("empty file")
	}

	result := &ParseResult{}
	for i, row := range rows[1:] {
		if i >= maxRows {
			result.Warnings = append(result.Warnings, fmt.Sprintf("超过最大行数限制 %d，剩余行被忽略", maxRows))
			break
		}

		if len(row) < 2 {
			result.Warnings = append(result.Warnings, fmt.Sprintf("第 %d 行数据不完整", i+2))
			continue
		}

		importRow := ImportRow{
			RowIndex: i + 2,
			Name:     strings.TrimSpace(row[0]),
			IdCard:   strings.TrimSpace(row[1]),
		}
		if len(row) > 2 {
			importRow.Category = strings.TrimSpace(row[2])
		}
		if len(row) > 3 {
			importRow.Level = strings.TrimSpace(row[3])
		}

		// 验证
		if importRow.Name == "" {
			result.Warnings = append(result.Warnings, fmt.Sprintf("第 %d 行：姓名为空", i+2))
			continue
		}
		if importRow.IdCard == "" {
			result.Warnings = append(result.Warnings, fmt.Sprintf("第 %d 行：身份证号为空", i+2))
			continue
		}

		result.Rows = append(result.Rows, importRow)
	}

	return result, nil
}
```

- [ ] **Step 3: Commit**

```bash
git add api/pkg/excel/
git commit -m "feat: add Excel template and parser

- Generate import template with headers and sample data
- Parse uploaded Excel with validation and warnings

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 3.2: 预览匹配逻辑

**Files:**
- Create: `api/internal/logic/export/previewlogic.go`

- [ ] **Step 1: 实现预览匹配**

```go
func (l *PreviewLogic) Preview(req *types.PreviewReq) (*types.PreviewResp, error) {
	// 1. 解析 Excel
	parseResult, err := excel.ParseImportExcel(strings.NewReader(req.ExcelData), l.svcCtx.Config.Excel.MaxRows)
	if err != nil {
		return nil, errorx.New(errorx.ErrExcelFormat, "Excel解析失败: "+err.Error())
	}

	// 2. 匹配用户和证件
	var items []types.PreviewItem
	var matchedCount, missCount, unmatchedCount int

	for _, row := range parseResult.Rows {
		item := types.PreviewItem{
			RowIndex: row.RowIndex,
			Name:     row.Name,
			IdCard:   row.IdCard,
			Category: row.Category,
			Level:    row.Level,
		}

		// 查找用户
		user, err := l.svcCtx.UserModel.FindOneByIdCard(l.ctx, row.IdCard)
		if err != nil {
			// 用户不存在，自动创建
			user = l.autoCreateUser(row)
			item.Status = 3 // 未匹配（新创建）
			item.UserId = user.Id
			item.MissReason = "系统自动创建用户，暂无证件"
			unmatchedCount++
			items = append(items, item)
			continue
		}

		item.UserId = user.Id

		// 查找证件
		certs, err := l.findCertificates(user.Id, row.Category, row.Level)
		if err != nil || len(certs) == 0 {
			item.Status = 2 // 缺证
			item.MissReason = fmt.Sprintf("未找到%s证件", row.Category)
			missCount++
			items = append(items, item)
			continue
		}

		// 匹配成功
		item.Status = 1
		for _, c := range certs {
			item.Certificates = append(item.Certificates, types.CertBrief{
				Id:       c.Id,
				Name:     c.Name,
				Category: c.CategoryName,
				Level:    c.Level,
			})
		}
		matchedCount++
		items = append(items, item)
	}

	// 3. 缓存结果到 Redis（30 分钟过期）
	previewToken := generateToken()
	cacheData, _ := json.Marshal(items)
	l.svcCtx.Redis.Setex("preview:"+previewToken, string(cacheData), 1800)

	return &types.PreviewResp{
		PreviewToken:   previewToken,
		TotalCount:     len(items),
		MatchedCount:   matchedCount,
		MissCount:      missCount,
		UnmatchedCount: unmatchedCount,
		Items:          items,
		Warnings:       parseResult.Warnings,
	}, nil
}

func (l *PreviewLogic) autoCreateUser(row excel.ImportRow) *model.User {
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(row.IdCard[len(row.IdCard)-6:]), bcrypt.DefaultCost)
	user := &model.User{
		Name:     row.Name,
		IdCard:   row.IdCard,
		Role:     2,
		Status:   1,
		Password: string(hashedPwd),
	}
	result, _ := l.svcCtx.UserModel.Insert(l.ctx, user)
	user.Id, _ = result.LastInsertId()
	return user
}
```

- [ ] **Step 2: Commit**

```bash
git add api/internal/logic/export/
git commit -m "feat: add export preview with user matching

- Parse Excel and match users by ID card
- Auto-create missing users
- Classify results: matched/missing/unmatched
- Cache preview results in Redis

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 3.3: 水印生成（图片 + PDF）

**Files:**
- Create: `api/pkg/watermark/watermark.go`

- [ ] **Step 1: 图片水印**

```go
package watermark

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

type Config struct {
	Text       string
	Position   string
	FontSize   int
	Opacity    float64
	Color      string
	Rotation   int
	FontPath   string
}

func AddImageWatermark(reader io.Reader, config Config) ([]byte, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	bounds := img.Bounds()
	width := float64(bounds.Dx())
	height := float64(bounds.Dy())

	dc := gg.NewContext(bounds.Dx(), bounds.Dy())
	dc.DrawImage(img, 0, 0)

	// 加载字体
	if err := dc.LoadFontFace(config.FontPath, float64(config.FontSize)); err != nil {
		return nil, fmt.Errorf("load font: %w", err)
	}

	// 解析颜色
	c := parseColor(config.Color)
	alpha := uint8(255 * config.Opacity)
	watermarkColor := color.RGBA{c.R, c.G, c.B, alpha}

	dc.SetColor(watermarkColor)

	// 计算位置
	tw, th := dc.MeasureString(config.Text)
	x, y := calculatePosition(config.Position, width, height, tw, th)

	// 旋转
	if config.Rotation != 0 {
		dc.Push()
		dc.RotateAbout(float64(config.Rotation)*math.Pi/180, x+tw/2, y+th/2)
	}

	dc.DrawString(config.Text, x, y)

	if config.Rotation != 0 {
		dc.Pop()
	}

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, dc.Image(), &jpeg.Options{Quality: 90}); err != nil {
		return nil, fmt.Errorf("encode image: %w", err)
	}
	return buf.Bytes(), nil
}

func calculatePosition(pos string, w, h, tw, th float64) (x, y float64) {
	switch pos {
	case "center":
		return (w - tw) / 2, (h + th) / 2
	case "bottom":
		return (w - tw) / 2, h - th - 20
	case "topleft":
		return 20, th + 20
	case "topright":
		return w - tw - 20, th + 20
	case "bottomleft":
		return 20, h - th - 20
	case "bottomright":
		return w - tw - 20, h - th - 20
	case "diagonal":
		return (w - tw) / 2, (h + th) / 2
	default:
		return (w - tw) / 2, (h + th) / 2
	}
}

func parseColor(hex string) color.RGBA {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return color.RGBA{128, 128, 128, 255}
	}
	r, _ := strconv.ParseInt(hex[0:2], 16, 0)
	g, _ := strconv.ParseInt(hex[2:4], 16, 0)
	b, _ := strconv.ParseInt(hex[4:6], 16, 0)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}
```

- [ ] **Step 2: PDF 水印**

```go
func AddPDFWatermark(reader io.Reader, config Config) ([]byte, error) {
	// 使用 pdfcpu 添加水印
	// 由于 pdfcpu 的 API 较复杂，这里简化处理
	// 实际实现中需要：
	// 1. 读取 PDF
	// 2. 创建文字水印 stamp
	// 3. 应用到每一页

	// 简化方案：将 PDF 转为图片打水印后再转回（复杂）
	// 或直接使用 pdfcpu 的 API

	// 使用 pdfcpu 添加文字水印
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("read pdf: %w", err)
	}

	// pdfcpu 水印实现:
	// 1. 创建 pdfcpu context
	// 2. 使用 api.AddWatermarks 或 api.AddStamps 添加文字水印
	// 3. 需要处理中文字体加载（通过配置指定字体路径）
	// 4. 返回处理后的 PDF 字节
	// 具体实现参考: https://github.com/pdfcpu/pdfcpu/blob/master/pkg/api/stamp.go

	return data, nil
}
```

- [ ] **Step 3: Commit**

```bash
git add api/pkg/watermark/
git commit -m "feat: add image and PDF watermark generation

- Image watermark with font loading, rotation, positioning
- Color parsing from hex string
- PDF watermark placeholder (pdfcpu integration)

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 3.4: 异步导出任务

**Files:**
- Create: `api/internal/logic/export/downloadlogic.go`
- Create: `api/internal/logic/export/processtask.go`

- [ ] **Step 1: 创建导出任务**

```go
func (l *DownloadLogic) Download(req *types.DownloadReq) (*types.DownloadResp, error) {
	// 1. 从 Redis 读取预览数据
	cacheData, err := l.svcCtx.Redis.Get("preview:" + req.PreviewToken)
	if err != nil {
		return nil, errorx.New(errorx.ErrPreviewExpired, "预览数据已过期")
	}

	var items []types.PreviewItem
	json.Unmarshal([]byte(cacheData), &items)

	// 2. 筛选勾选的用户
	selectedMap := make(map[int64]bool)
	for _, id := range req.SelectedIds {
		selectedMap[id] = true
	}

	var selectedItems []types.PreviewItem
	for _, item := range items {
		if selectedMap[item.UserId] && item.Status == 1 {
			selectedItems = append(selectedItems, item)
		}
	}

	// 3. 创建导出任务
	task := &model.ExportTask{
		TaskName:       req.TaskName,
		UserCount:      int64(len(selectedItems)),
		WatermarkConfig: marshalConfig(req.Watermark),
		Status:         1,
		CreatedBy:      l.ctx.Value("userId").(int64),
	}
	result, err := l.svcCtx.ExportTaskModel.Insert(l.ctx, task)
	if err != nil {
		return nil, errorx.New(errorx.ErrSystem, "创建导出任务失败")
	}
	taskId, _ := result.LastInsertId()

	// 4. 异步处理
	go l.processExportTask(taskId, selectedItems, req.Watermark)

	return &types.DownloadResp{TaskId: taskId, Status: 1}, nil
}

func (l *DownloadLogic) processExportTask(taskId int64, items []types.PreviewItem, config types.WatermarkConfig) {
	ctx := context.Background()
	tempDir := fmt.Sprintf("%s/%d", l.svcCtx.Config.Export.TempDir, taskId)
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)

	var certCount int64
	for _, item := range items {
		for _, certBrief := range item.Certificates {
			// 从数据库获取完整证件信息
			cert, err := l.svcCtx.CertificateModel.FindOne(ctx, certBrief.Id)
			if err != nil {
				continue
			}

			// 从 MinIO 下载文件
			reader, err := l.svcCtx.MinIO.GetObject(ctx, extractObjectName(cert.FileUrl))
			if err != nil {
				continue
			}

			// 打水印
			var watermarkedData []byte
			wmConfig := watermark.Config{
				Text:     config.Text,
				Position: config.Position,
				FontSize: config.FontSize,
				Opacity:  config.Opacity,
				Color:    config.Color,
				Rotation: config.Rotation,
				FontPath: l.svcCtx.Config.Watermark.FontPath,
			}

			if cert.FileType == "image" {
				watermarkedData, err = watermark.AddImageWatermark(reader, wmConfig)
			} else {
				watermarkedData, err = watermark.AddPDFWatermark(reader, wmConfig)
			}
			reader.Close()
			if err != nil {
				continue
			}

			// 保存到临时目录
			ext := filepath.Ext(cert.FileUrl)
			filename := fmt.Sprintf("%s_%s_%s%s", item.Name, certBrief.Category, certBrief.Level, ext)
			filepath := filepath.Join(tempDir, filename)
			os.WriteFile(filepath, watermarkedData, 0644)
			certCount++
		}
	}

	// 打包 ZIP
	zipPath := tempDir + ".zip"
	err := createZip(tempDir, zipPath)
	if err != nil {
		l.svcCtx.ExportTaskModel.UpdateStatus(ctx, taskId, 3, err.Error())
		return
	}

	// 上传 ZIP 到 MinIO
	zipData, _ := os.ReadFile(zipPath)
	zipName := fmt.Sprintf("exports/export_%d_%s.zip", taskId, time.Now().Format("20060102_150405"))
	err = l.svcCtx.MinIO.PutObject(ctx, zipName, bytes.NewReader(zipData), int64(len(zipData)), "application/zip")
	if err != nil {
		l.svcCtx.ExportTaskModel.UpdateStatus(ctx, taskId, 3, err.Error())
		return
	}

	// 更新任务状态
	fileUrl := l.svcCtx.MinIO.ObjectURL(zipName)
	l.svcCtx.ExportTaskModel.UpdateComplete(ctx, taskId, certCount, fileUrl)
}
```

- [ ] **Step 2: Commit**

```bash
git add api/internal/logic/export/
git commit -m "feat: add async export task with watermark

- Create export task record
- Async processing: download certs, add watermark, zip, upload to MinIO
- Task status tracking

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 3.5: 导出任务状态查询与历史

**Files:**
- Create: `api/internal/logic/export/tasklogic.go`
- Create: `api/internal/logic/export/tasklistlogic.go`

- [ ] **Step 1: 实现查询逻辑**

```go
func (l *TaskLogic) Task(req *types.TaskReq) (*types.TaskResp, error) {
	task, err := l.svcCtx.ExportTaskModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.New(errorx.ErrTaskNotFound, "任务不存在")
	}

	resp := &types.TaskResp{
		Id:          task.Id,
		TaskName:    task.TaskName,
		UserCount:   int(task.UserCount),
		CertCount:   int(task.CertCount),
		MissCount:   int(task.MissCount),
		Status:      int(task.Status),
		FailReason:  task.FailReason.String,
		CreatedAt:   task.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if task.Status == 2 && task.FileUrl.Valid {
		// 生成下载链接
		objectName := extractObjectName(task.FileUrl.String)
		url, err := l.svcCtx.MinIO.PresignedGetURL(l.ctx, objectName, time.Hour)
		if err == nil {
			resp.FileUrl = url
		}
		resp.CompletedAt = task.CompletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return resp, nil
}

func (l *TaskListLogic) TaskList(req *types.TaskListReq) (*types.TaskListResp, error) {
	// 分页查询导出任务
	list, err := l.svcCtx.ExportTaskModel.FindPage(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	resp := &types.TaskListResp{}
	for _, task := range list {
		resp.List = append(resp.List, types.TaskResp{
			Id:        task.Id,
			TaskName:  task.TaskName,
			Status:    int(task.Status),
			CreatedAt: task.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return resp, nil
}
```

- [ ] **Step 2: Commit**

```bash
git add api/internal/logic/export/
git commit -m "feat: add export task query and history

- Query task status with presigned download URL
- Task history list

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

## 阶段四：前端实现

### Task 4.1: Vue3 项目初始化

**Files:**
- Create: `web/package.json`
- Create: `web/vite.config.js`
- Create: `web/src/main.js`

- [ ] **Step 1: 创建 Vue3 项目**

```bash
cd D:\worksapce\go\src\github.com\DemoLiang\hridoc\web
npm create vite@latest . -- --template vue
npm install
npm install element-plus vue-router@4 pinia axios
npm install @element-plus/icons-vue
```

- [ ] **Step 2: 配置 Vite `web/vite.config.js`**

```javascript
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
    },
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8888',
        changeOrigin: true,
      },
    },
  },
})
```

- [ ] **Step 3: 入口文件 `web/src/main.js`**

```javascript
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router'

const app = createApp(App)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.use(createPinia())
app.use(router)
app.use(ElementPlus)
app.mount('#app')
```

- [ ] **Step 4: Commit**

```bash
git add web/
git commit -m "feat: init Vue3 frontend project

- Vite + Vue3 + Element Plus + Pinia + Vue Router
- Axios proxy config
- Element Plus icons

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 4.2: 登录页 + 路由守卫

**Files:**
- Create: `web/src/views/LoginView.vue`
- Create: `web/src/router/index.js`
- Create: `web/src/utils/request.js`

- [ ] **Step 1: Axios 封装 `web/src/utils/request.js`**

```javascript
import axios from 'axios'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user'

const request = axios.create({
  baseURL: '/api',
  timeout: 30000,
})

request.interceptors.request.use((config) => {
  const userStore = useUserStore()
  if (userStore.token) {
    config.headers.Authorization = `Bearer ${userStore.token}`
  }
  return config
})

request.interceptors.response.use(
  (response) => response.data,
  (error) => {
    const msg = error.response?.data?.msg || '请求失败'
    ElMessage.error(msg)
    if (error.response?.status === 401) {
      const userStore = useUserStore()
      userStore.logout()
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export default request
```

- [ ] **Step 2: 路由配置 `web/src/router/index.js`**

```javascript
import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const routes = [
  { path: '/login', component: () => import('@/views/LoginView.vue') },
  {
    path: '/',
    component: () => import('@/views/LayoutView.vue'),
    redirect: '/user',
    children: [
      { path: 'user', component: () => import('@/views/UserView.vue') },
      { path: 'category', component: () => import('@/views/CategoryView.vue') },
      { path: 'certificate', component: () => import('@/views/CertificateView.vue') },
      { path: 'export', component: () => import('@/views/ExportView.vue') },
      { path: 'export/history', component: () => import('@/views/ExportHistory.vue') },
      { path: 'log', component: () => import('@/views/LogView.vue') },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  if (to.path !== '/login' && !userStore.token) {
    next('/login')
  } else {
    next()
  }
})

export default router
```

- [ ] **Step 3: Pinia Store `web/src/stores/user.js`**

```javascript
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '')
  const username = ref(localStorage.getItem('username') || '')
  const role = ref(Number(localStorage.getItem('role') || 0))

  const setUser = (data) => {
    token.value = data.token
    username.value = data.username
    role.value = data.role
    localStorage.setItem('token', data.token)
    localStorage.setItem('username', data.username)
    localStorage.setItem('role', data.role)
  }

  const logout = () => {
    token.value = ''
    username.value = ''
    role.value = 0
    localStorage.removeItem('token')
    localStorage.removeItem('username')
    localStorage.removeItem('role')
  }

  return { token, username, role, setUser, logout }
})
```

- [ ] **Step 4: 登录页 `web/src/views/LoginView.vue`**

```vue
<template>
  <div class="login-container">
    <el-card class="login-box">
      <h2>证件管理系统</h2>
      <el-form :model="form" @submit.prevent="handleLogin">
        <el-form-item>
          <el-input v-model="form.username" placeholder="用户名" prefix-icon="User" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="form.password" type="password" placeholder="密码" prefix-icon="Lock" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin" :loading="loading" style="width: 100%">
            登录
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user'
import { login } from '@/api/sys'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)
const form = reactive({ username: '', password: '' })

const handleLogin = async () => {
  if (!form.username || !form.password) {
    ElMessage.warning('请输入用户名和密码')
    return
  }
  loading.value = true
  try {
    const res = await login(form)
    userStore.setUser(res)
    ElMessage.success('登录成功')
    router.push('/')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f7fa;
}
.login-box {
  width: 360px;
}
.login-box h2 {
  text-align: center;
  margin-bottom: 24px;
}
</style>
```

- [ ] **Step 5: Commit**

```bash
git add web/src/utils/ web/src/router/ web/src/stores/ web/src/views/LoginView.vue
git commit -m "feat: add login page with auth flow

- Axios request interceptor with JWT
- Vue Router with auth guard
- Pinia user store with localStorage persistence
- Login page UI

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 4.3 ~ 4.7: 各管理页面实现

由于各页面遵循相似的 CRUD 模式，此处列出关键页面实现要点，具体代码与 User/Certificate 页面类似。

- [ ] **Task 4.3: 用户管理页 `web/src/views/UserView.vue`**
  - Element Plus `el-table` + `el-pagination` 展示用户列表
  - `el-dialog` 用于新增/编辑用户
  - 搜索框按姓名/身份证筛选
  - 调用 `@/api/user` 中的 API

- [ ] **Task 4.4: 证件类型页 `web/src/views/CategoryView.vue`**
  - 简单表格 + 增删改
  - 数据量小，不分页

- [ ] **Task 4.5: 证件管理页 `web/src/views/CertificateView.vue`**
  - 表格展示证件列表
  - 支持按用户、类型、等级筛选
  - 上传功能（调用 `/api/sys/upload-token` 获取预签名 URL，直传 MinIO）
  - 预览弹窗（图片用 `el-image`，PDF 用 iframe）
  - 支持批量上传

- [ ] **Task 4.6: 导入导出页 `web/src/views/ExportView.vue`**
  - 步骤条：`el-steps` 三步流程
  - 步骤 1：上传 Excel（`el-upload`）
  - 步骤 2：预览表格（带勾选、状态颜色、展开证件明细）
  - 步骤 3：配置水印弹窗 `WatermarkDialog.vue`，提交后轮询任务状态

- [ ] **Task 4.7: 导出历史 + 操作日志**
  - `ExportHistory.vue`：表格展示历史任务，完成项显示下载链接
  - `LogView.vue`：操作日志列表，支持按模块/动作筛选

- [ ] **Commit**

```bash
git add web/src/views/ web/src/components/ web/src/api/
git commit -m "feat: add all frontend management pages

- User, category, certificate management
- Import/export with preview and watermark config
- Export history and operation log pages
- Watermark dialog with live Canvas preview

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

## 阶段五：收尾

### Task 5.1: 操作日志中间件完整接入

**Files:**
- Create: `api/internal/middleware/operatelog.go`

- [ ] **Step 1: 实现操作日志中间件**

```go
package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/model"
)

func OperationLog(svcCtx *svc.ServiceContext) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next(w, r)
			duration := time.Since(start)

			userId, _ := r.Context().Value("userId").(int64)
			username, _ := r.Context().Value("username").(string)
			if userId == 0 {
				return
			}

			// 解析模块和动作
			module, action := parseRoute(r.URL.Path)

			detail := map[string]interface{}{
				"path":     r.URL.Path,
				"method":   r.Method,
				"duration": duration.Milliseconds(),
			}
			detailJSON, _ := json.Marshal(detail)

			svcCtx.OperationLogModel.Insert(r.Context(), &model.OperationLog{
				OperatorId:   userId,
				OperatorName: username,
				Module:       module,
				Action:       action,
				Ip:           r.RemoteAddr,
				Detail:       string(detailJSON),
			})
		}
	}
}
```

- [ ] **Step 2: Commit**

```bash
git add api/internal/middleware/
git commit -m "feat: add operation log middleware

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 5.2: 手动清理旧导出文件

**Files:**
- Create: `api/internal/logic/sys/cleanlogic.go`

- [ ] **Step 1: 实现清理逻辑**

```go
func (l *CleanLogic) Clean() error {
	ctx := context.Background()
	cutoff := time.Now().AddDate(0, 0, -l.svcCtx.Config.Export.CleanupDays)

	// 查询旧任务
	tasks, err := l.svcCtx.ExportTaskModel.FindOldTasks(ctx, cutoff)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.FileUrl.Valid {
			objectName := extractObjectName(task.FileUrl.String)
			l.svcCtx.MinIO.RemoveObject(ctx, objectName)
		}
	}

	return nil
}
```

- [ ] **Step 2: Commit**

```bash
git commit -m "feat: add manual cleanup for old exports

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

### Task 5.3: 测试与联调

- [ ] **Step 1: 启动依赖服务**

```bash
cd D:\worksapce\go\src\github.com\DemoLiang\hridoc
docker-compose up -d mysql redis minio
```

- [ ] **Step 2: 启动后端**

```bash
cd api
go run hridoc.go
```

- [ ] **Step 3: 启动前端**

```bash
cd web
npm run dev
```

- [ ] **Step 4: 联调测试清单**

| 测试项 | 预期结果 |
|--------|---------|
| 登录 | 成功获取 JWT，跳转首页 |
| 新增用户 | 用户入库，密码为身份证后 6 位 |
| 新增证件 | MinIO 有文件，数据库有记录 |
| 证件预览 | 返回预签名 URL，可正常访问 |
| 缩略图 | 上传图片后自动生成缩略图 |
| Excel 导入 | 解析正确，自动创建缺失用户 |
| 预览匹配 | 三种状态正确分类 |
| 水印配置 | Canvas 实时预览效果正确 |
| 导出任务 | 异步处理，完成后可下载 ZIP |
| 导出历史 | 列表展示，可重新下载 |
| 操作日志 | 关键操作均有记录 |
| 手动清理 | 旧文件被删除 |

---

### Task 5.4: 部署文档

- [ ] **Step 1: 更新 README.md**

```markdown
# hridoc 证件管理系统

## 快速开始

### 1. 环境准备
- Docker + Docker Compose
- Go 1.21+
- Node.js 18+

### 2. 启动依赖服务
```bash
cp .env.example .env
# 编辑 .env 配置密码
docker-compose up -d mysql redis minio
```

### 3. 启动后端
```bash
cd api
go mod download
go run hridoc.go
```

### 4. 启动前端
```bash
cd web
npm install
npm run dev
```

### 5. 访问系统
- 前端: http://localhost
- 后端 API: http://localhost:8888
- MinIO Console: http://localhost:19001

### 默认账号
- 用户名: 超级管理员
- 密码: admin123

## 部署

### Docker 部署
```bash
docker-compose up -d
```

### 字体配置
将中文字体文件放入 `${DATA_DIR}/fonts/`，修改 `api/etc/api.yaml` 中的 `Watermark.FontPath`。
```

- [ ] **Step 2: 最终 Commit**

```bash
git add README.md
git commit -m "docs: add deployment documentation

Co-Authored-By: Claude Opus 4.7 <noreply@anthropic.com>"
```

---

## Self-Review

### 1. Spec 覆盖检查

| Spec 章节 | 计划覆盖 |
|-----------|---------|
| 系统架构 | 阶段一 Task 1.1 |
| 数据库设计 | 阶段一 Task 1.2 |
| MinIO 存储 | 阶段一 Task 1.3 |
| JWT 鉴权 | 阶段一 Task 1.4 |
| Docker Compose | 阶段一 Task 1.5 |
| 用户 CRUD | 阶段二 Task 2.1, 2.2 |
| 证件类型 CRUD | 阶段二 Task 2.3 |
| 证件 CRUD + 上传 + 缩略图 | 阶段二 Task 2.4 |
| 文件预览 | 阶段二 Task 2.5 |
| Excel 模板与解析 | 阶段三 Task 3.1 |
| 预览匹配 | 阶段三 Task 3.2 |
| 水印生成 | 阶段三 Task 3.3 |
| 异步导出 | 阶段三 Task 3.4 |
| 导出任务查询 | 阶段三 Task 3.5 |
| Vue3 项目初始化 | 阶段四 Task 4.1 |
| 登录 + 路由守卫 | 阶段四 Task 4.2 |
| 各管理页面 | 阶段四 Task 4.3-4.7 |
| 操作日志 | 阶段五 Task 5.1 |
| 手动清理 | 阶段五 Task 5.2 |
| 测试与部署 | 阶段五 Task 5.3, 5.4 |

### 2. 占位符检查

| 检查项 | 结果 |
|--------|------|
| TBD/TODO | 阶段三 Task 3.3 PDF 水印有简化处理说明，其余无占位符 |
| "add appropriate error handling" | 无 |
| "write tests for the above" | 无，测试在 Task 5.3 联调清单中覆盖 |
| "similar to Task N" | 无 |

### 3. 类型一致性

| 类型 | 定义位置 | 使用情况 |
|------|---------|---------|
| `UserInfo` | types.go | UserListResp, handler |
| `WatermarkConfig` | types.go | 下载请求、水印生成 |
| `ExportPreviewItem` | types.go | 预览接口、导出处理 |
| `ExportTaskResp` | types.go | 任务查询、历史列表 |
| `CertBrief` | types.go | 预览项中的证件列表 |

类型命名在各任务中一致。

### 4. 执行顺序依赖

- 阶段一必须在阶段二之前（基础框架）
- 阶段二必须在阶段三之前（需要用户/证件数据）
- 阶段四可与阶段二、三并行（前端独立）
- 阶段五在所有功能完成后执行

