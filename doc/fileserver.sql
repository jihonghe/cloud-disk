CREATE TABLE `file_meta` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `file_hash` char(40) NOT NULL COMMENT '文件hash',
    `file_name` varchar(256) NOT NULL COMMENT '文件名',
    `file_size` bigint(20) DEFAULT 0 COMMENT '文件大小',
    `file_address` varchar(1024) NOT NULL COMMENT '文件存储位置',
    `create_time` datetime default NOW() COMMENT '创建日期',
    `update_time` datetime default NOW() COMMENT '更新日期',
    `status` int(11) NOT NULL DEFAULT 1 COMMENT '文件状态(禁用0/可用1/已删除-1)',
    `ext1` int(11) COMMENT '备用字段1',
    `ext2` text COMMENT '备用字段2',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_file_hash` (`file_hash`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `nick_name` varchar(128) DEFAULT 'new_user' COMMENT '用户昵称',
    `user_name` varchar(64) NOT NULL COMMENT '用户名，可用于账号登录',
    `password` varchar(256) NOT NULL COMMENT '用户加密密码',
    `email` varchar(64) NOT NULL COMMENT '邮箱，可用于账号验证及找回',
    `phone` varchar(128) COMMENT '手机号',
    `email_validated` bool DEFAULT FALSE COMMENT '邮箱是否已验证，FALSE表示未验证',
    `phone_validated` bool DEFAULT FALSE COMMENT '手机号是否已验证，FALSE表示未验证',
    `signup_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
    `last_active` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最后活跃时间',
    `profile` text COMMENT '用户属性',
    `status` int(11) NOT NULL DEFAULT 1 COMMENT '账户状态（禁用0/启用1/锁定2/标记删除3）',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_name` (`user_name`),
    UNIQUE KEY `idx_email` (`email`),
    KEY `idx_status` (`status`)
)ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
CREATE TABLE `user_token` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL COMMENT '用户名',
    `user_token` char(40) NOT NULL COMMENT '用户登录token',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
