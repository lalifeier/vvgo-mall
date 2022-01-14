CREATE TABLE `dict` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '字典数据id',
  `dict_type_id` bigint unsigned NOT NULL COMMENT '字典类型id',
  `type` varchar(100) NOT NULL DEFAULT '' COMMENT '字典类型',
  `label` varchar(100) NOT NULL DEFAULT '' COMMENT '字典标签',
  `value` varchar(100) NOT NULL DEFAULT '' COMMENT '字典键值',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:禁用 1:启用',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
  `sort` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否默认值 0:否 1:是',
  `create_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `create_by` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `update_at` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `update_by` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典数据表';
