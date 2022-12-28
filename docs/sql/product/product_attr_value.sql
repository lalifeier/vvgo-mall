CREATE TABLE `product_attr_value` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '销售属性值ID',
    `attr_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '销售属性ID',
    `value` varchar(255)  NOT NULL DEFAULT '' COMMENT '销售属性值',
    `desc` varchar(255)  NOT NULL DEFAULT '' COMMENT '销售属性值描述',
    `create_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
    `update_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
    `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 1:enable, 0:disable, -1:deleted',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售属性值';
