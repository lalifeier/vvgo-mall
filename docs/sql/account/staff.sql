CREATE TABLE `staff` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账号id',
    `email` varchar(30) NOT NULL DEFAULT '' COMMENT '员工邮箱',
    `phone` varchar(15) NOT NULL DEFAULT '' COMMENT '员工手机号',
    `name` varchar(30) NOT NULL DEFAULT '' COMMENT '员工姓名',
    `nickname` varchar(30) NOT NULL DEFAULT '' COMMENT '员工昵称',
    `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '员工头像(相对路径)',
    `gender` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '员工性别 1:男性 2:女性',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`uid`),
    KEY `idx_email` (`email`),
    KEY `idx_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='员工表';
