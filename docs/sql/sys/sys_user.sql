DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账号id',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '姓名',
  `phone` varchar(15) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(30) NOT NULL DEFAULT '' COMMENT '邮箱',
  `nickname` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` enum('male','female','unknow') NOT NULL DEFAULT 'unknow' COMMENT '性别',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
  `dept_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '部门id',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`),
  KEY `idx_phone` (`phone`),
  KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
