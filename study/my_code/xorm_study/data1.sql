CREATE TABLE `lt_user` (
	  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
	  `blacktime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '黑名单限制到期时间',
	  `realname` varchar(50) NOT NULL DEFAULT '' COMMENT '联系人',
	  `mobile` varchar(50) NOT NULL DEFAULT '' COMMENT '手机号',
	  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '联系地址',
	  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
	  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
	  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'IP地址',
	  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40652 DEFAULT CHARSET=utf8;
