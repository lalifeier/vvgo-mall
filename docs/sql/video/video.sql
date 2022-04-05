CREATE TABLE `video` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `hash_code` varchar(32) NOT NULL COMMENT '唯一hash码',
  `resolution` varchar(32) NOT NULL COMMENT '分辨率 1080P,720P,480P,360P,240P',
  `size` int(11) unsigned NOT NULL COMMENT '文件大小 单位KB',
  `duration` int(11) unsigned NOT NULL COMMENT '时长 单位秒',
  `url` varchar(255) NOT NULL COMMENT '视频地址',
  `meta_data` json NOT NULL COMMENT '元数据 {language:xxx, tag:xxx, category:xxx}'
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

CREATE TABLE `video_chunk` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `video_id` int(11) unsigned NOT NULL COMMENT '视频id',
  `start_time` int(11) unsigned NOT NULL COMMENT '开始时间 单位秒',
  `end_time` int(11) unsigned NOT NULL COMMENT '结束时间 单位秒',
  `folder` varchar(255) NOT NULL COMMENT '视频分片目录',
  `resolution` varchar(32) NOT NULL COMMENT '分辨率 1080P,720P,480P,360P,240P'
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频分片表';

CREATE TABLE `thumbnail` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `video_id` int(11) unsigned NOT NULL COMMENT '视频id',
  `folder` varchar(255) NOT NULL COMMENT '缩略图目录',
  `size` int(11) unsigned NOT NULL COMMENT '缩略图大小 单位KB',
  `type` varchar(32) NOT NULL COMMENT '缩略图类型',
  `moment` int(11) unsigned NOT NULL COMMENT '缩略图时间 单位秒'
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='缩略图表';

CREATE TABLE `user_video` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `video_id` int(11) unsigned NOT NULL COMMENT '视频id',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `meta_data` json NOT NULL COMMENT '元数据 {language:xxx, tag:xxx, category:xxx}'
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户视频关系表';
