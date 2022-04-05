CREATE TABLE `user_card` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `card_name` varchar(25) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '持卡人姓名',
  `card_number` varchar(25) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '银行卡号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_card_number` (`card_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='银行卡表';
