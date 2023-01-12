CREATE TABLE `goods_type` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `title` varchar(255) NOT NULL DEFAULT '' COMMENT '类别名称',
                              `explain` varchar(255) NOT NULL DEFAULT '' COMMENT '类别说明',
                              `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父类ID',
                              `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0:冻结 1:正常',
                              `op_id` int(11) NOT NULL DEFAULT '0' COMMENT '操作人ID',
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`),
                              KEY `parent_id` (`parent_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品类别表';

