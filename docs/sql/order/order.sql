CREATE TABLE IF NOT EXISTS `address` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `consignee` VARCHAR(255) NOT NULL COMMENT '收货人姓名',
  `email` VARCHAR(255) NOT NULL COMMENT '收货人邮箱',
  `mobile` VARCHAR(255) NOT NULL COMMENT '收货人手机号',
  `country` json NOT NULL COMMENT '国家 id：国家id name：国家名称',
  `province` json NOT NULL COMMENT '省 id:省id name:省名称',
  `city` json NOT NULL COMMENT '市 id:市id name:市名称',
  `coountry` json NOT NULL COMMENT '区 id:区id name:区名称',
  `street` json NOT NULL COMMENT '街道 id:街道id name:街道名称',
  `detailed_address` VARCHAR(255) NOT NULL COMMENT '详细地址',
  `postal_code` VARCHAR(255) NOT NULL COMMENT '邮编',
  `address_id` int NOT NULL COMMENT '地址id',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认地址',
  `label` VARCHAR(255) NOT NULL COMMENT '地址类型标签，家、公司等',
  `longitude` VARCHAR(255) NOT NULL COMMENT '经度',
  `latitude` VARCHAR(255) NOT NULL COMMENT '纬度',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '地址表';


CREATE TABLE IF NOT EXISTS `pay_method` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL COMMENT '支付方式名称',
  `desc` VARCHAR(255) NOT NULL COMMENT '支付方式描述',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '支付方式表';

-- 在线支付
-- 货到付款

CREATE TABLE IF NOT EXISTS `invoice_type` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL COMMENT '发票类型名称',
  `desc` VARCHAR(255) NOT NULL COMMENT '发票类型描述',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '发票类型表';

-- 个人
-- 公司


CREATE TABLE IF NOT EXISTS `order` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `order_no` VARCHAR(255) NOT NULL COMMENT '订单编号',
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户id',
  `address_id` BIGINT UNSIGNED NOT NULL COMMENT '地址id',
  `pay_method_id` BIGINT UNSIGNED NOT NULL COMMENT '支付方式id',
  `invoice_type_id` BIGINT UNSIGNED NOT NULL COMMENT '发票类型id',

  -- `invoice_title` VARCHAR(255) NOT NULL COMMENT '发票抬头',
  -- `invoice_content` VARCHAR(255) NOT NULL COMMENT '发票内容',
  -- `invoice_amount` DECIMAL(10,2) NOT NULL COMMENT '发票金额',
  -- `invoice_tax_no` VARCHAR(255) NOT NULL COMMENT '税号',

  `total_amount` DECIMAL(10,2) NOT NULL COMMENT '订单总金额',
  `comments` VARCHAR(255) NOT NULL COMMENT '订单备注',
  `order_status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '订单状态：0草稿，1下单',
  `pay_status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '支付状态：0未支付，1已支付',
  `logistics_status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '物流状态：0未发货，1已发货',
  `receipt_status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '收货状态：0未收货，1已收货',
  `sales_return_status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '退货状态：0未退货，1已退货',
  `invoice_status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '开票状态：0未开票，1已开票',
  `is_delete` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0未删除，1已删除',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '订单表';


-- 设计规则
-- 1.日期：年月日时分秒 //14或8(只要日期)位
-- 2.随机数字 //6或2位
-- 3.商家id //哪个商家的数据
-- 4.用户id //哪个用户的数据。取后4位
-- 5.商品id
-- 6.业务id字段 //支付通道
-- 7.其他唯一id字段


-- 订单主表：也叫订单表，保存订单的基本信息。
-- 订单商品表：保存订单中的商品信息。
-- 订单支付表：保存订单的支付和退款信息。
-- 订单优惠表：保存订单使用的所有优惠信息。


-- https://github.com/CrazyCodes/E-Commerce-SQL
