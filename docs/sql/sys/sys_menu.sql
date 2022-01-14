CREATE TABLE `sys_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父级ID',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `url` varchar(30) NOT NULL DEFAULT '' COMMENT '菜单URL',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:目录 1:菜单 2:按钮',
  `perms` varchar(500) NOT NULL DEFAULT '' comment '授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)',
  `icon` varchar(30) NOT NULL DEFAULT '' COMMENT '菜单图标',
  `sort` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';
