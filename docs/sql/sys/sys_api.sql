DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `group` varchar(32) NOT NULL DEFAULT '' COMMENT '接口分组',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '接口名称',
  `path` varchar(255) NOT NULL DEFAULT '' COMMENT '接口路径',
  `method` varchar(16) NOT NULL DEFAULT '' COMMENT '接口请求方式',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '接口描述',
  `permission` varchar(255) NOT NULL DEFAULT '' COMMENT '接口权限',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:禁用 1:启用 -1:删除',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci  COMMENT='接口表';
