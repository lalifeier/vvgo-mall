CREATE TABLE `product_member_price` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `sku_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'SKU ID',
    `member_level_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '会员等级ID',
    `member_price` decimal(10,2) unsigned NOT NULL DEFAULT '0' COMMENT '会员价',
    `member_level_name` varchar(255) NOT NULL DEFAULT '' COMMENT '会员等级名称',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT AUTO_INCREMENT=666666 CHARSET=utf8mb4 COMMENT='商品会员价格表';
