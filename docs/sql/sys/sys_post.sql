CREATE TABLE `sys_post` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '岗位名称',
  `remark` varchar(30) NOT NULL DEFAULT '' COMMENT '岗位备注',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='岗位表';
