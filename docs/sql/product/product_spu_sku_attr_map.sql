CREATE TABLE `product_spu_sku_attr_map` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `spu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'SPU ID',
    `sku_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'SKU ID',
    `attr_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '销售属性ID',
    `attr_name` varchar(255) NOT NULL DEFAULT '0' COMMENT '销售属性名称',
    `attr_value_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '销售属性值ID',
    `attr_value_name` varchar(255) NOT NULL DEFAULT '0' COMMENT '销售属性值',
    `create_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人staff_id',
    `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 1:enable, 0:disable, -1:deleted',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关联关系冗余表';
