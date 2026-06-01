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
