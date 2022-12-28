CREATE TABLE `product_brand` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '品牌ID',
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '品牌名称',
    `desc` varchar(255)  NOT NULL DEFAULT '' COMMENT '品牌描述',
    `logo_url` varchar(255)  NOT NULL DEFAULT '' COMMENT '品牌logo图片',
    `sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
    `create_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
    `update_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
    `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 1:enable, 0:disable, -1:deleted',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unq_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='品牌表';
