CREATE TABLE `certificate` (
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `id` bigint auto_increment COMMENT '证书ID',
    `name` varchar(64) NOT NULL COMMENT '证书名称',
    `path` varchar(64) DEFAULT NULL COMMENT '存储位置',
    `user_id` bigint COMMENT '证书持有人',
    PRIMARY KEY (`id`),
    KEY `idx_name` (`name`),
    KEY `idx_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `certificate_type` (
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `id` bigint auto_increment COMMENT '分类ID',
    `name` varchar(64) NOT NULL COMMENT '分类名称',
    PRIMARY KEY (`id`),
    KEY `idx_name` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;