-- 通用抽奖工具(万能胶Glue) glue_session_prizes_timer 抽奖场次奖品定时投放器表
CREATE TABLE `glue_session_prizes_timer` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `session_prizes_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '抽奖场次奖品ID',
    `delivery_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '定时投放奖品数量的时间',
    `prize_quantity` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '奖品数量',
    `create_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
    `update_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
    `status` tinyint(1)  NOT NULL DEFAULT '0' COMMENT '状态 -1:deleted, 0:wait, 1:success',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='抽奖场次奖品定时投放器表';
