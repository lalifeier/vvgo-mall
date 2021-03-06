CREATE TABLE `sys_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `system_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '系统ID',
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父级ID',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `component` varchar(255) NOT NULL DEFAULT '' COMMENT '组件',
  `url` varchar(30) NOT NULL DEFAULT '' COMMENT '菜单URL',
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单图标',
  `is_cache` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否缓存 0:否 1:是',
  `is_affix` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否固定到标签栏 0:否 1:是',
  `is_breadcrumb` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否显示面包屑 0:否 1:是',
  `is_hidden` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否隐藏 0:否 1:是',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
  `sort` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:禁用 1:启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';
