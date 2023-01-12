CREATE TABLE `goods_info` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `title` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
                              `cover` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图',
                              `picture` text NOT NULL COMMENT '商品图片的json格式数据',
                              `profile` varchar(255) NOT NULL DEFAULT '' COMMENT '商品简介',
                              `prices` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品价格',
                              `detail` text NOT NULL COMMENT '商品详情',
                              `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品类别',
                              `business_id` int(11) NOT NULL DEFAULT '0' COMMENT '商家ID',
                              `goods_key` varchar(255) NOT NULL DEFAULT '' COMMENT '商品关键字',
                              `op_id` int(11) NOT NULL DEFAULT '0' COMMENT '操作人ID',
                              `state` tinyint(11) NOT NULL DEFAULT '0' COMMENT '商品状态 1:未上架 2：以上架 3：以下架',
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`),
                              KEY `type_id` (`type_id`) USING BTREE,
                              KEY `b_t_id` (`business_id`,`type_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品信息表';

