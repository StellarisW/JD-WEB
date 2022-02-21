/*
《Go Web编程实战派从入门到精通》源码
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uid` int(10) DEFAULT '0' COMMENT '用户编号',
  `phone` varchar(30) DEFAULT '' COMMENT '用户手机',
  `name` varchar(30) DEFAULT '' COMMENT '用户姓名',
  `zipcode` varchar(20) DEFAULT '' COMMENT '邮政编码',
  `address` varchar(250) DEFAULT '' COMMENT '地址',
  `default_address` int(11) DEFAULT '0' COMMENT '默认地址',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='地址信息';

-- ----------------------------
-- Table structure for administrator
-- ----------------------------
DROP TABLE IF EXISTS `administrator`;
CREATE TABLE `administrator` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT '' COMMENT '用户名',
  `password` varchar(100) DEFAULT '' COMMENT '密码',
  `mobile` varchar(100) DEFAULT NULL COMMENT '手机号',
  `email` varchar(50) DEFAULT '' COMMENT '邮箱',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `role_id` int(10) DEFAULT '0' COMMENT '角色编号',
  `is_super` tinyint(4) DEFAULT '0' COMMENT '是否超级管理员',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='管理员表';

-- ----------------------------
-- Table structure for auth
-- ----------------------------
DROP TABLE IF EXISTS `auth`;
CREATE TABLE `auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `module_name` varchar(80) NOT NULL DEFAULT '',
  `action_name` varchar(80) DEFAULT '' COMMENT '操作名称',
  `type` tinyint(1) DEFAULT '0' COMMENT '节点类型',
  `url` varchar(250) DEFAULT '' COMMENT '跳转地址',
  `module_id` int(11) DEFAULT '0' COMMENT '模块编号',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `description` varchar(250) DEFAULT '' COMMENT '描述',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态',
  `checked` tinyint(1) DEFAULT '0' COMMENT '是否检验',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8 COMMENT='权限控制';

-- ----------------------------
-- Table structure for jwt
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklist`;
CREATE TABLE `jwt_blacklist` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `jwt` varchar(255) NOT NULL DEFAULT '' COMMENT 'JWT令牌',
  `expire_time` timestamp NOT NULL COMMENT '过期时间',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='jwt黑名单';

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(50) DEFAULT '' COMMENT '标题',
  `banner_type` tinyint(4) DEFAULT '0' COMMENT '类型',
  `banner_img` varchar(100) DEFAULT '' COMMENT '图片地址',
  `link` varchar(200) DEFAULT '' COMMENT '连接',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `status` int(10) DEFAULT '0' COMMENT '状态',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='焦点图表';

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(250) DEFAULT '' COMMENT '标题',
  `price` decimal(10,2) DEFAULT '0.00',
  `product_version` varchar(50) DEFAULT '' COMMENT '版本',
  `num` int(11) DEFAULT '0' COMMENT '数量',
  `product_gift` varchar(100) DEFAULT '' COMMENT '商品礼物',
  `product_fitting` varchar(100) DEFAULT '' COMMENT '商品搭配',
  `product_color` varchar(50) DEFAULT '' COMMENT '商品颜色',
  `product_img` varchar(150) DEFAULT '' COMMENT '商品图片',
  `product_attr` varchar(100) DEFAULT '' COMMENT '商品属性',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='购物车';

-- ----------------------------
-- Records of cart
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `link` varchar(250) DEFAULT '' COMMENT '连接',
  `position` int(10) DEFAULT '0' COMMENT '位置',
  `is_opennew` int(10) DEFAULT '0' COMMENT '是否新打开',
  `relation` varchar(100) DEFAULT '' COMMENT '关系',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `status` int(10) DEFAULT '0' COMMENT '状态',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `order_id` varchar(100) DEFAULT '' COMMENT '订单编号',
  `uid` int(11) DEFAULT '0' COMMENT '用户编号',
  `all_price` decimal(10,2) DEFAULT '0.00' COMMENT '价格',
  `phone` varchar(30) DEFAULT '' COMMENT '电话',
  `name` varchar(100) DEFAULT '' COMMENT '名字',
  `address` varchar(250) DEFAULT '' COMMENT '地址',
  `zipcode` varchar(30) DEFAULT '' COMMENT '邮编',
  `pay_status` tinyint(4) DEFAULT '0' COMMENT '支付状态',
  `pay_type` tinyint(4) DEFAULT '0' COMMENT '支付类型',
  `order_status` tinyint(4) DEFAULT '0' COMMENT '订单状态',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for order_item
-- ----------------------------
DROP TABLE IF EXISTS `order_item`;
CREATE TABLE `order_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '订单编号',
  `order_id` varchar(100) DEFAULT '0' COMMENT '订单编号',
  `uid` int(11) DEFAULT '0' COMMENT '用户编号',
  `product_title` varchar(100) DEFAULT '' COMMENT '商品标题',
  `product_id` int(11) DEFAULT '0' COMMENT '商品编号',
  `product_img` varchar(200) DEFAULT '' COMMENT '商品图片',
  `product_price` decimal(10,2) DEFAULT '0.00' COMMENT '商品价格',
  `product_num` int(10) DEFAULT '0' COMMENT '商品数量',
  `product_version` varchar(100) DEFAULT '' COMMENT '商品版本',
  `product_color` varchar(100) DEFAULT '' COMMENT '商品颜色',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `sub_title` varchar(100) DEFAULT '' COMMENT '子标题',
  `product_sn` varchar(50) DEFAULT '',
  `cate_id` int(10) DEFAULT '0' COMMENT '分类id',
  `click_count` int(10) DEFAULT '0' COMMENT '点击数',
  `product_number` int(10) DEFAULT '0' COMMENT '商品编号',
  `price` decimal(10,2) DEFAULT '0.00' COMMENT '价格',
  `market_price` decimal(10,2) DEFAULT '0.00' COMMENT '市场价格',
  `relation_product` varchar(100) DEFAULT '' COMMENT '关联商品',
  `product_attr` varchar(100) DEFAULT '' COMMENT '商品属性',
  `product_version` varchar(100) DEFAULT '' COMMENT '商品版本',
  `product_img` varchar(100) DEFAULT '' COMMENT '商品图片',
  `product_gift` varchar(100) DEFAULT '',
  `product_fitting` varchar(100) DEFAULT '',
  `product_color` varchar(100) DEFAULT '' COMMENT '商品颜色',
  `product_keywords` varchar(100) DEFAULT '' COMMENT '关键词',
  `product_desc` varchar(50) DEFAULT '' COMMENT '描述',
  `product_content` varchar(100) DEFAULT '' COMMENT '内容',
  `is_delete` tinyint(4) DEFAULT '0' COMMENT '是否删除',
  `is_hot` tinyint(4) DEFAULT '0' COMMENT '是否热门',
  `is_best` tinyint(4) DEFAULT '0' COMMENT '是否畅销',
  `is_new` tinyint(4) DEFAULT '0' COMMENT '是否新品',
  `product_type_id` tinyint(4) DEFAULT '0' COMMENT '商品类型编号',
  `sort` int(10) DEFAULT '0' COMMENT '商品分类',
  `status` tinyint(4) DEFAULT '0' COMMENT '商品状态',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='商品';

DROP TABLE IF EXISTS `product_collect`;
CREATE TABLE `product_collect`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_id` int(11) DEFAULT '0',
    `product_id` int(11) DEFAULT '0',
    `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
    `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='商品收藏';

-- ----------------------------
-- Table structure for product_attr
-- ----------------------------
DROP TABLE IF EXISTS `product_attr`;
CREATE TABLE `product_attr` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT '0' COMMENT '商品编号',
  `attribute_cate_id` int(10) DEFAULT '0' COMMENT '属性分类编号',
  `attribute_id` int(10) DEFAULT '0' COMMENT '属性编号',
  `attribute_title` varchar(100) DEFAULT '' COMMENT '属性标题',
  `attribute_type` int(10) DEFAULT '0' COMMENT '属性类型',
  `attribute_value` varchar(100) DEFAULT '' COMMENT '属性值',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COMMENT='商品属性';

-- ----------------------------
-- Table structure for product_cate
-- ----------------------------
DROP TABLE IF EXISTS `product_cate`;
CREATE TABLE `product_cate` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(200) DEFAULT '' COMMENT '分类名称',
  `cate_img` varchar(200) DEFAULT '' COMMENT '分类图片',
  `link` varchar(250) DEFAULT '' COMMENT '链接',
  `template` text COMMENT '模版',
  `pid` int(10) DEFAULT '0' COMMENT '父编号',
  `sub_title` varchar(100) DEFAULT '' COMMENT '子标题',
  `keywords` varchar(250) DEFAULT '' COMMENT '关键字',
  `description` text COMMENT '描述',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='商品分类';

-- ----------------------------
-- Table structure for product_color
-- ----------------------------
DROP TABLE IF EXISTS `product_color`;
CREATE TABLE `product_color` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `color_name` varchar(100) DEFAULT '' COMMENT '颜色名字',
  `color_value` varchar(100) DEFAULT '' COMMENT '颜色值',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `checked` tinyint(4) DEFAULT '0' COMMENT '是否检验',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for product_image
-- ----------------------------
DROP TABLE IF EXISTS `product_image`;
CREATE TABLE `product_image` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT '0' COMMENT '商品编号',
  `img_url` varchar(250) DEFAULT '' COMMENT '图片地址',
  `color_id` int(11) DEFAULT '0' COMMENT '颜色编号',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for product_type
-- ----------------------------
DROP TABLE IF EXISTS `product_type`;
CREATE TABLE `product_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `description` varchar(500) DEFAULT '' COMMENT '描述',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for product_type_attribute
-- ----------------------------
DROP TABLE IF EXISTS `product_type_attribute`;
CREATE TABLE `product_type_attribute` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cate_id` int(11) DEFAULT '0' COMMENT '分类编号',
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `attr_type` tinyint(4) DEFAULT '0' COMMENT '属性类型',
  `attr_value` varchar(100) DEFAULT '' COMMENT '属性值',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '标题名称',
  `description` varchar(500) DEFAULT '' COMMENT '描述',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for role_auth
-- ----------------------------
DROP TABLE IF EXISTS `role_auth`;
CREATE TABLE `role_auth` (
  `auth_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '权限编号',
  `role_id` int(11) DEFAULT '0' COMMENT '角色编号',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`auth_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for setting
-- ----------------------------
DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `site_title` varchar(100) DEFAULT '' COMMENT '商城名称',
  `site_logo` varchar(250) DEFAULT '' COMMENT '商城图标',
  `site_keywords` varchar(100) DEFAULT '' COMMENT '商城关键字',
  `site_description` varchar(500) DEFAULT '' COMMENT '商城描述',
  `no_picture` varchar(100) DEFAULT '' COMMENT '没有图片显示',
  `site_icp` varchar(50) DEFAULT '' COMMENT '商城ICP',
  `site_tel` varchar(50) DEFAULT '' COMMENT '商城手机号',
  `search_keywords` varchar(250) DEFAULT '' COMMENT '搜索关键字',
  `tongji_code` varchar(500) DEFAULT '' COMMENT '统计编码',
  `appid` varchar(50) DEFAULT '' COMMENT 'oss appid',
  `app_secret` varchar(80) DEFAULT '' COMMENT 'oss app_secret',
  `end_point` varchar(200) DEFAULT '' COMMENT 'oss 终端点',
  `bucket_name` varchar(200) DEFAULT '' COMMENT 'oss 桶名称',
  `oss_status` tinyint(4) DEFAULT '0' COMMENT 'oss 状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of setting
-- ----------------------------
BEGIN;
INSERT INTO `setting` VALUES (1, 'LeastMall商城', 'ab', 'LeastMall商城', 'LeastMall商城', 'a', 'LeastMall商城', '88888888', 'LeastMall商城', 'dd', 'b', 'c', 'd', 'a', 1);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `phone` varchar(30) DEFAULT '' COMMENT '手机号',
  `password` varchar(80) DEFAULT '' COMMENT '密码',
  `last_ip` varchar(50) DEFAULT '' COMMENT '最近ip',
  `email` varchar(80) DEFAULT '' COMMENT '邮编',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_sms
-- ----------------------------
DROP TABLE IF EXISTS `user_sms`;
CREATE TABLE `user_sms` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(50) DEFAULT '' COMMENT 'ip地址',
  `phone` varchar(50) DEFAULT '' COMMENT '手机号',
  `send_count` int(10) DEFAULT '0' COMMENT '发送统计',
  `add_day` varchar(200) DEFAULT '' COMMENT '添加日期',
  `sign` varchar(80) DEFAULT '' COMMENT '签名',
  `update_time` datetime NOT NULL DEFAULT NOW() COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
