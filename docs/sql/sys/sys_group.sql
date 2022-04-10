DROP TABLE IF EXISTS `sys_group`;
CREATE TABLE `sys_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父级ID',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '用户组名称',
  `sort` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:禁用 1:启用',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户组表';
