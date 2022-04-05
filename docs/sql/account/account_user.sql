CREATE TABLE `account_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '用户名',
  `phone` varchar(15) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(30) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `create_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `create_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT '创建ip',
  `last_login_at` int(11) NOT NULL DEFAULT '0' COMMENT '最后一次登录时间',
  `last_login_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT '最后一次登录ip',
  `login_count` int(11) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0:禁用 1:启用 -1:删除',
  PRIMARY KEY (`id`),
  KEY `idx_username` (`username`),
  KEY `idx_phone` (`phone`),
  KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账户表';
