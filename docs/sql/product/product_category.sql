CREATE TABLE `product_category` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
    `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '分类名称',
    `desc` varchar(255)  NOT NULL DEFAULT '' COMMENT '分类描述',
    `pic_url` varchar(255)  NOT NULL DEFAULT '' COMMENT '分类图片',
    `path` varchar(255)  NOT NULL DEFAULT '' COMMENT '分类地址{pid}-{child_id}-...',
    `create_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人staff_id',
    `update_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改人staff_id',
    `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 1:enable, 0:disable, -1:deleted',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分类表';
