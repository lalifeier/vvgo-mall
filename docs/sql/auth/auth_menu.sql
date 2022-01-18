CREATE TABLE `auth_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `system_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '系统ID',
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父菜单ID',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 1:菜单 2:按钮',
  `url` varchar(30) NOT NULL DEFAULT '' COMMENT '菜单URL',
  `perms` varchar(500) NOT NULL DEFAULT '' comment '授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
  `sort` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:禁用 1:启用 -1:删除',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';
