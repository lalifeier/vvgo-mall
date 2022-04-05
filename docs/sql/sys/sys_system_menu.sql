CREATE TABLE `sys_system_menu` (
  `system_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '系统ID',
  `menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  PRIMARY KEY (`system_id`, `menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统菜单关联表';
