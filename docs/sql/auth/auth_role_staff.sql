CREATE TABLE `auth_role_staff` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `staff_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '员工ID',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色员工关联表';
