CREATE TABLE `authorities` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(255) NOT NULL COMMENT '权限名称',
  `authory` varchar(255) NOT NULL COMMENT '权限代码',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci  COMMENT='oauth授权表';

CREATE TABLE `oauth_approvals` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `client_id` varchar(255) NOT NULL COMMENT '客户端id',
  `scope` varchar(255) NOT NULL COMMENT '权限',
  `expires` datetime NOT NULL COMMENT '过期时间',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci  COMMENT='oauth授权表';

CREATE TABLE `oauth_client_details` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `client_id` varchar(255) NOT NULL COMMENT '客户端id',
  `resource_ids` varchar(255) NOT NULL COMMENT '资源id',
  `client_secret` varchar(255) NOT NULL COMMENT '客户端密钥',
  `scope` varchar(255) NOT NULL COMMENT '权限',
  `authorized_grant_types` varchar(255) NOT NULL COMMENT '授权类型',
  `web_server_redirect_uri` varchar(255) NOT NULL COMMENT '回调地址',
  `authorities` varchar(255) NOT NULL COMMENT '权限',
  `access_token_validity` int(11) unsigned NOT NULL COMMENT 'token有效期',
  `refresh_token_validity` int(11) unsigned NOT NULL COMMENT '刷新token有效期',
  `additional_information` varchar(255) NOT NULL COMMENT '附加信息',
  `autoapprove` varchar(255) NOT NULL COMMENT '自动授权',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci  COMMENT='oauth授权表';

CREATE TABLE `oauth_code` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `code` varchar(255) NOT NULL COMMENT '授权码',
  `authentication` varchar(255) NOT NULL COMMENT '认证',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci  COMMENT='oauth授权表';
