-- 活动信息
-- 秒杀商品信息
-- 秒杀进度



CREATE TABLE IF NOT EXISTS `reserve_info` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `sku_id` int(11) unsigned NOT NULL COMMENT '商品编号',
  `reserve_start_time` DATETIME NOT NULL COMMENT '预约开始时间',
  `reserve_end_time` DATETIME NOT NULL COMMENT '预约结束时间',
  -- `reserve_num` INT UNSIGNED NOT NULL COMMENT '预约数量',
  -- `reserve_price` DECIMAL(10,2) UNSIGNED NOT NULL COMMENT '预约价格',
  -- `reserve_status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '预约状态',
  `seckill_start_time` DATETIME NOT NULL COMMENT '秒杀开始时间',
  `seckill_end_time` DATETIME NOT NULL COMMENT '秒杀结束时间',
  `is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '预约活动表';


CREATE TABLE IF NOT EXISTS `reserve_user` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `reserve_info_id` BIGINT UNSIGNED NOT NULL COMMENT '预约活动ID',
  `sku_id` int(11) unsigned NOT NULL COMMENT '商品编号',
  `reserve_user_id` BIGINT UNSIGNED NOT NULL COMMENT '预约用户',
  `reserve_user_name` VARCHAR(32) NOT NULL COMMENT '预约用户名',
  `reserve_time` DATETIME NOT NULL COMMENT '预约时间',
  `is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`)
  KEY `idx_reserve_info_id` (`reserve_info_id`)
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4
  COMMENT = '预约用户表'
)

-- redis
-- key reserve_info_{skuid}
-- key reserve_user_{userName} value sku_id list


-- addReservation(skuId, userName)
-- cancelReservation(skuId, userName)
-- getReservationInfoList(skuIds)
-- validateReservation(skuId, userName)
-- getMyReservationList(userName)

-- 给预约过的用户发消息提醒 深度分页
select * from reserve_user where id > (select id from reserve_user where sku_id = #{skuId} limit 100000,1) limit 20;

set seckill:28:commodity:189:stock 100

SELECT stock
FROM stock_info
WHERE commodity_id = 189 AND seckill_id = 28;

UPDATE stock_info
SET stock = stock - 1
WHERE commodity_id = 189 AND seckill_id = 28 AND stock > 0;


-- EVALSHA
-- 逻辑：首先判断活动库存是否存在，以及库存余量是否够本次购买数量，如果不够，则返回0，如果够则完成扣减并返回1
-- KEYS[1] : 活动库存的key
-- KEYS[2] : 活动库存的扣减数量
local stock = redis.call('get', KEYS[1])
if not stock or tonumber(stock) < tonumber(KEYS[2]) then
  return 0
end
redis.call('decrby',KEYS[1], KEYS[2])
return 1
