CREATE TABLE `sys_dict` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '字典数据id',
  `dict_type_id` int unsigned NOT NULL DEFAULT 0 COMMENT '字典类型id',
  `type` varchar(100) NOT NULL DEFAULT '' COMMENT '字典类型',
  `label` varchar(100) NOT NULL DEFAULT '' COMMENT '字典标签',
  `value` varchar(100) NOT NULL DEFAULT '' COMMENT '字典键值',
  `status` tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '状态 0:禁用 1:启用',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
  `sort` tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '排序',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '是否默认值 0:否 1:是',
  `is_deleted` tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '是否删除 0:否 1:是',
  `create_by` bigint unsigned NOT NULL DEFAULT 0 COMMENT '更新人',
  `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` bigint unsigned NOT NULL DEFAULT 0 COMMENT '更新人',
  `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci  COMMENT='字典数据表';
