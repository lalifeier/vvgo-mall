-- 通用抽奖工具(万能胶Glue) glue_session 抽奖场次表
CREATE TABLE `glue_session` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '场次ID',
    `activity_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '活动ID',
    `times_limit` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '抽奖次数限制，0默认不限制',
    `start_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '场次开始时间',
    `end_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '场次结束时间',
    `create_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
    `update_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
    `status` tinyint(1)  NOT NULL DEFAULT '0' COMMENT '状态 -1:deleted, 0:disable, 1:enable',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='抽奖场次表';
