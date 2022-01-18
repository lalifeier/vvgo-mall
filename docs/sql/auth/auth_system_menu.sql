CREATE TABLE `auth_system_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `system_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '系统ID',
  `menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统菜单关联表';
