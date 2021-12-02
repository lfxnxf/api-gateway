CREATE TABLE `vehicle_info` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `name` varchar(128) NOT NULL DEFAULT '' COMMENT '名字',
    `load_num` tinyint NOT NULL DEFAULT '0' COMMENT '荷载人数',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='车辆信息字典表';

CREATE TABLE `users` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
     `name` varchar(128) NOT NULL DEFAULT '' COMMENT '名字',
     `address` varchar(256) NOT NULL DEFAULT '' COMMENT '地址',
     `boss_id` bigint NOT NULL COMMENT '老板uid',
     `identity` tinyint NOT NULL DEFAULT '0' COMMENT '1:老板，2：管理员，3：乘务员，4：家长',
     `phone` bigint NOT NULL DEFAULT '0' COMMENT '手机号',
     `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
     `token` char(32) NOT NULL DEFAULT '' COMMENT 'token',
     `status` tinyint not null default 1 comment '状态，1：正常，2：删除',
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`),
     KEY `idx_phone_password` (`phone`,`password`),
     KEY `idx_token` (`token`),
     KEY `idx_boss_id` (`boss_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='用户表';

insert into users (name, address, boss_id, identity, phone)values ('徐帆', '北京', 0, 1, 13888888888);
insert into users (name, address, boss_id, identity, phone)values ('陈国珍', '三河', 1, 3, 13666666666);

CREATE TABLE `vehicle` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `vehicle_info_id` int NOT NULL DEFAULT 0 COMMENT '车辆类型id，对应vehicle_info',
    `boss_id` bigint NOT NULL DEFAULT 0 COMMENT '老板id',
    `license_plate` char(8) NOT NULL DEFAULT '' COMMENT '车牌号',
    `driver_id` bigint NOT NULL DEFAULT 0 COMMENT '当前司机',
    `status` tinyint not null default 1 comment '状态，1：正常，2：删除',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_boss_id` (`boss_id`),
    KEY `idx_driver_id` (`driver_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='车辆信息';


CREATE TABLE `sites` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `name` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '名字',
    `longitude` VARCHAR(128) not null DEFAULT '' comment '经度',
    `latitude` VARCHAR(128) not null DEFAULT '' comment '纬度',
    `status` tinyint not null default 1 comment '状态，1：正常，2：删除',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='站点';

CREATE TABLE `vehicle_sites` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `vehicle_id` bigint(20) not null default 0 comment '车辆id',
    `site_id` bigint(20) not null default 0 comment '站点id',
    `sort` tinyint not null default 0 comment '站点排序',
    `status` tinyint not null default 1 comment '状态，1：正常，2：删除',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_vehicle_id_site_id` (`vehicle_id`, `site_id`),
    KEY `idx_vehicle_id_sort` (`vehicle_id`, `sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='车辆站点';


CREATE TABLE `shifts` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `name` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '名字',
    `status` tinyint not null default 1 comment '状态，1：正常，2：删除',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='班次';

CREATE TABLE `shifts_sites` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `shift_id` bigint(20) not null default 0 comment '班次id',
    `site_id` bigint(20) not null default 0 comment '站点id',
    `sort` tinyint not null default 0 comment '站点排序',
    `status` tinyint not null default 1 comment '状态，1：正常，2：删除',
    `arrive_time` varchar(30) not null default '' comment '到站时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_shift_id_site_id` (`shift_id`, `site_id`),
    KEY `idx_shift_id_sort` (`shift_id`, `sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='班次站点';

INSERT INTO `school`.`vehicle_info`(`id`, `name`, `load_num`, `create_time`, `update_time`) VALUES (1, '小型车辆', 9, '2021-12-01 08:20:59', '2021-12-01 08:21:17');
INSERT INTO `school`.`vehicle_info`(`id`, `name`, `load_num`, `create_time`, `update_time`) VALUES (2, '中型车辆', 20, '2021-12-01 08:21:15', '2021-12-01 08:21:45');
INSERT INTO `school`.`vehicle_info`(`id`, `name`, `load_num`, `create_time`, `update_time`) VALUES (3, '大型车辆', 36, '2021-12-01 08:21:40', '2021-12-01 08:21:47');