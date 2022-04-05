CREATE TABLE `product_album` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `sku_id` int(11) unsigned NOT NULL COMMENT 'SKU ID',
    `name` varchar(25)  NOT NULL COMMENT '商品名称',
    `url` varchar(45)  DEFAULT NULL COMMENT '图片地址',
    `size` int(11) DEFAULT NULL COMMENT '视频大小',
    `desc` varchar(255)  NOT NULL COMMENT '图片介绍',
    `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
    `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '图片状态',
    `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '资源类型 0:图片 1:视频',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='专辑表';
