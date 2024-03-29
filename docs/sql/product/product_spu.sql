CREATE TABLE `product_spu` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT 'spu名称',
    `desc` varchar(255)  NOT NULL DEFAULT '' COMMENT 'spu描述',
    `brand_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '品牌ID',
    `category_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分类ID',
    `unit` varchar(255)  NOT NULL DEFAULT '' COMMENT 'spu单位',
    `main_url` text COMMENT '商品介绍主图 多个图片逗号分隔',
    `banner_url` text COMMENT 'banner图片 多个图片逗号分隔',
    `price_fee` int unsigned NOT NULL DEFAULT 0 COMMENT '售价，整数方式保存',
    `price_scale` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '售价，金额对应的小数位数',
    `market_price_fee` int unsigned NOT NULL DEFAULT 0 COMMENT '市场价，整数方式保存',
    `market_price_scale` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '市场价，金额对应的小数位数',
    `create_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
    `update_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
    `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 1:enable, 0:disable, -1:deleted',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT AUTO_INCREMENT=666666 CHARSET=utf8mb4 COMMENT='spu表';
