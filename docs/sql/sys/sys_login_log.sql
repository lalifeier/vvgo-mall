DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账号id',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '账号名',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '登录类型 0：用户名登录 1：手机号登录 2：邮箱登录',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '登录状态 0：登录失败 1：登录成功',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `login_at` int(11) NOT NULL DEFAULT '0' COMMENT '登录时间',
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT '登录ip',
  `ua` varchar(255) NOT NULL DEFAULT '' COMMENT '登录浏览器user agent',
  `browser` varchar(255) NOT NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(255) NOT NULL DEFAULT '' COMMENT '操作系统',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账号登录日志表';
