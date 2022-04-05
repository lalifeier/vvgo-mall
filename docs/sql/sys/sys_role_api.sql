CREATE TABLE `sys_role_api` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `api_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '接口ID',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色接口关联表';
