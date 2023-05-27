CREATE TABLE `score` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT,
     `created_at` datetime(3) DEFAULT NULL,
     `updated_at` datetime(3) DEFAULT NULL,
     `deleted_at` datetime(3) DEFAULT NULL,
     `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
     `score` float NOT NULL DEFAULT '0' COMMENT '积分',
     `created_by` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;