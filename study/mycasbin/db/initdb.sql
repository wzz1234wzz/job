CREATE DATABASE IF NOT EXISTS wzz_db default charset utf8 COLLATE utf8_general_ci;

DROP TABLE IF EXISTS `user`;

CREATE TABLE IF NOT EXISTS `user` (
`ysid` bigint(20) NOT NULL AUTO_INCREMENT,
`name` varchar(128) NOT NULL UNIQUE,
`pwd_hash` varchar(128) ,
`group` VARCHAR(20) NOT NULL,
`rank` tinyint(1) DEFAULT '1',
`email` varchar(128),
`phone` varchar(16),
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`is_activated` tinyint(1) DEFAULT '1',
 PRIMARY KEY (`ysid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `user` (`name`, `pwd_hash`, `group`, `rank`,`email`, `phone`, `created_at`, `updated_at`, `is_activated`) VALUES ('root', '$2a$10$6fod9YfBJCv.AdunawtlBO4VcojeaqBEWz9wVGJsuYXcMWOEHjwSq', '', 1, '', '', '2020-11-19 15:53:34', '2020-11-19 15:53:34', '1');
