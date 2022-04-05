-- 暂存购物车
{
  "cart": [
    {
      "sku_id": 1,
      "count": 1,
      "timestamp": 1527058983,
      "selected": true
    },
  ]
}


-- 用户购物车
CREATE TABLE `user_cart` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `sku_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品sku_id',
  `count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
  `timestamp` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '加购时间',
  `selected` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否选中',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户购物车表';


-- hash数据结构
-- 增加购物车商品
hset cart:user_id sku_id {count: 1, timestamp: 1527058983, selected: true}

-- 获取购物车所有商品（全选）
hgetall cart:user_id

-- 购物车商品数量
hlen cart:user_id

-- 删除商品
hdel cart:user_id sku_id

-- 增加商品
hincrby cart:user_id sku_id:num 1

hmset cart:user_id sku_id:num 1 sku_id:info {timestamp: 1527058983, selected: true}
