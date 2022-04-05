DROP TABLE IF EXISTS `casbin_rules`;
CREATE TABLE `casbin_rules` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ptype` varchar(255)  NOT NULL DEFAULT '',
  `v0` varchar(255)  NOT NULL DEFAULT '',
  `v1` varchar(255)  NOT NULL DEFAULT '',
  `v2` varchar(255)  NOT NULL DEFAULT '',
  `v3` varchar(255)  NOT NULL DEFAULT '',
  `v4` varchar(255)  NOT NULL DEFAULT '',
  `v5` varchar(255)  NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
