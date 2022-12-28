-- 通用抽奖工具(万能胶Glue) glue_session_prizes 抽奖场次奖品表
CREATE TABLE `glue_session_prizes` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `session_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '场次ID',
    `node` varchar(255)  NOT NULL DEFAULT '' COMMENT '节点标识 按时间抽奖: 空值, 按抽奖次数抽奖: 第几次参与值, 按数额范围区间抽奖: 数额区间上限值',
    `prize_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '奖品类型 1:优惠券, 2:积分, 3:实物, 4:空奖 ...',
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '奖品名称',
    `pic_url` varchar(255)  NOT NULL DEFAULT '' COMMENT '奖品图片',
    `value` varchar(255)  NOT NULL DEFAULT '' COMMENT '奖品抽象值 优惠券:优惠券ID, 积分:积分值, 实物: sku ID',
    `probability` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '中奖概率1~100',
    `create_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
    `update_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
    `status` tinyint(1)  NOT NULL DEFAULT '0' COMMENT '状态 -1:deleted, 0:disable, 1:enable',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='抽奖场次奖品表';
