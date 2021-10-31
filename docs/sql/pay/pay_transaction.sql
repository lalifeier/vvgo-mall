-- -----------------------------------------------------
-- Table 创建支付流水表
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `pay_transaction` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_id` VARCHAR(32) NOT NULL COMMENT '应用id',
  `pay_method_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付方式id，可以用来识别支付，如：支付宝、微信、Paypal等',
  `app_order_id` VARCHAR(64) NOT NULL COMMENT '应用方订单号',
  `transaction_id` VARCHAR(64) NOT NULL COMMENT '本次交易唯一id，整个支付系统唯一，生成他的原因主要是 order_id对于其它应用来说可能重复',
  `total_fee` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付金额，整数方式保存',
  `scale` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '金额对应的小数位数',
  `currency_code` CHAR(3) NOT NULL DEFAULT 'CNY' COMMENT '交易的币种',
  `pay_channel` VARCHAR(64) NOT NULL COMMENT '选择的支付渠道，比如：支付宝中的花呗、信用卡等',
  `expire_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单过期时间',
  `return_url` VARCHAR(255) NOT NULL COMMENT '支付后跳转url',
  `notify_url` VARCHAR(255) NOT NULL COMMENT '支付后，异步通知url',
  `email` VARCHAR(64) NOT NULL COMMENT '用户的邮箱',
  `sign_type` VARCHAR(10) NOT NULL DEFAULT 'RSA' COMMENT '采用的签方式：MD5 RSA RSA2 HASH-MAC等',
  `intput_charset` CHAR(5) NOT NULL DEFAULT 'UTF-8' COMMENT '字符集编码方式',
  `payment_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '第三方支付成功的时间',
  `notify_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '收到异步通知的时间',
  `finish_time` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '通知上游系统的时间',
  `trade_no` VARCHAR(64) NOT NULL COMMENT '第三方的流水号',
  `transaction_code` VARCHAR(64) NOT NULL COMMENT '真实给第三方的交易code，异步通知的时候更新',
  `order_status` TINYINT NOT NULL DEFAULT 0 COMMENT '0:等待支付，1:待付款完成， 2:完成支付，3:该笔交易已关闭，-1:支付失败',
  `create_at` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_at` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `create_ip` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建的ip，这可能是自己服务的ip',
  `update_ip` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新的ip',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_tradid` (`transaction_id`),
  INDEX `idx_trade_no` (`trade_no`),
  INDEX `idx_ctime` (`create_at`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '发起支付的数据';
