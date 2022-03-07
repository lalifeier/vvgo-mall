CREATE TABLE `coupon_batch` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '批次ID',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '批次名称',
  `coupon_name` varchar(128) NOT NULL DEFAULT '' COMMENT '券名称',
  `rule_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '规则id',
  `total_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '总数量',
  `assign_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '已发放券数量',
  `used_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '已使用数量',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优惠券批次表';

CREATE TABLE `coupon_rule` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '规则id',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '规则名称',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '优惠券类型，0:满减，1:折扣',
  `content` BLOB NOT NULL DEFAULT '' COMMENT '规则内容',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优惠券规则表';


// 使用门槛
threshold: 5.01
// 优惠金额
amount: 5
// 使用范围，0—全场，1—商家，2—类别，3—商品
use_range: 3
// 商品 id
commodity_id: 10
// 每个用户可以领取的数量
receive_count: 1
// 是否互斥，true 表示互斥，false 表示不互斥
is_mutex: true
// 领取开始时间
receive_started_at: 2020-11-1 00:08:00
// 领取结束时间
receive_ended_at: 2020-11-6 00:08:00
// 使用开始时间
use_started_at: 2020-11-1 00:00:00
// 使用结束时间
use_ended_at: 2020-11-11 11:59:59


CREATE TABLE `coupon` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '优惠券id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `batch_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '批次ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0：未使用，1：已使用，2：已过期,3：冻结',
  `order_id` varchar(128) NOT NULL DEFAULT '' COMMENT '订单ID',
  `received_time` DATETIME NOT NULL DEFAULT '' COMMENT '领取时间',
  `validate_time` DATETIME NOT NULL DEFAULT '' COMMENT '有效日期',
  `used_time` DATETIME NOT NULL DEFAULT '' COMMENT '使用时间',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优惠券表';

CREATE TABLE `coupon_opt_record` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `coupon_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '优惠券id',
  `operating` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '操作，0—锁定，1—核销，2—解锁',
  `operated_at` DATETIME NOT NULL DEFAULT '' COMMENT '操作时间',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优惠券表';

CREATE TABLE `notify_msg` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '信息内容id',
  `coupon_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '优惠券id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `notify_day` varchar(128) NOT NULL DEFAULT '' COMMENT '需要进行通知的日期'.
  `notify_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '通知类型，1过期提醒',
  `notify_time` TIMESTAMP NOT NULL DEFAULT '' COMMENT '通知的时间，在该时间戳所在天内通知'.
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '通知状态，0初始状态 1成功 2失败',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='通知信息表';

CREATE TABLE `message` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '信息id',
  `rec_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '接受者id',
  `message_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '信息内容id',
  `is_read` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否已读'
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='信息表';

CREATE TABLE `message_content` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '信息内容id',
  `send_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '发送者id',
  `content` varchar(128)  NOT NULL DEFAULT '' COMMENT '站内信内容',
  `send_time` DATETIME NOT NULL DEFAULT '' COMMENT '发送时间'
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='信息内容表';
