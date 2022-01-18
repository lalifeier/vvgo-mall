CREATE TABLE `account_login_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账号id',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '登录类型 0：用户名登录 1：手机号登录 2：邮箱登录',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '登录状态 0：登录失败 1：登录成功',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `login_at` int(11) NOT NULL DEFAULT '0' COMMENT '登录时间',
  `login_ip` varchar(12) NOT NULL DEFAULT '' COMMENT '登录ip',
  `login_ua` varchar(255) NOT NULL DEFAULT '' COMMENT '登录浏览器user agent',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账号登录日志表';
