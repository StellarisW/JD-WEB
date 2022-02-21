-- ----------------------------
-- Records of address
-- ----------------------------
BEGIN;
INSERT INTO `address` VALUES (1, 1, '13882348889', '王五', '100000', '北京市xxxxx区xxxxx街道xxxx号',0, default, default);
INSERT INTO `address` VALUES (2, 2, '13812388888', '张三', '600000', '成都市xxxx区xxxx街道xxxx号', 0, default, default);
INSERT INTO `address` VALUES (3, 3, '13887456889', '李四', '100000', '重庆市xxxxx区xxxxx街道xxxx号', 0, default, default);
INSERT INTO `address` VALUES (4, 4, '13834522889', '王五', '100000', '烟台市xxxxx区xxxxx街道xxxx号', 1, default, default);
INSERT INTO `address` VALUES (5, 5, '13123123389', '王五', '100000', '北京市xxxxx区xxxxx街道xxxx号', 1, default, default);
INSERT INTO `address` VALUES (6, 6, '13564568889', '王五', '100000', '北京市xxxxx区xxxxx街道xxxx号', 1, default, default);
INSERT INTO `address` VALUES (7, 7, '13845668889', '王五', '100000', '北京市xxxxx区xxxxx街道xxxx号', 1, default, default);
INSERT INTO `address` VALUES (8, 8, '13845648889', '王五', '100000', '北京市xxxxx区xxxxx街道xxxx号', 1, default, default);
INSERT INTO `address` VALUES (9, 9, '13854666889', '王五', '100000', '北京市xxxxx区xxxxx街道xxxx号', 1, default, default);
INSERT INTO `address` VALUES (10, 10, '13888888889', '王五', '100000', '北京市xxxxx区xxxxx街道xxxx号', 1, default, default);
COMMIT;

-- ----------------------------
-- Records of administrator
-- ----------------------------
BEGIN;
INSERT INTO `administrator` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '1888888888', 'admin@163.com', 1, 1, 1, DEFAULT, DEFAULT);
INSERT INTO `administrator` VALUES (2, '工程师', '', '13999999999', 'programmer1@163.com', 1, 2, 0, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of auth
-- ----------------------------
BEGIN;
INSERT INTO `auth` VALUES (1, '系统管理员', 'get', 3, '3', 0, 0, '', 0,  0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (2, '组织部门', '', 3, '', 0, 0, '', 0,  0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (3, '模块管理', '', 3, '', 0, 0, '', 0,  0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (4, 'Banner管理', '', 3, '', 0, 0, '', 0,  0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (5, '商品管理', '', 3, '', 0, 0, '', 0,  0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (6, '订单管理', '', 3, '', 0, 0, '', 0,  0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (7, '设置管理', '', 3, '', 0, 0, '', 0,  0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (8, '管理员列表', '管理员列表', 2, '/administrator', 1, 0, '', 0,  1, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (9, '', '新增管理员', 2, '/administrator/add', 1, 0, '', 0,  0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (10, '', '部门列表', 2, '/role', 2, 0, '', 0,  0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (11, '', '新增部门', 2, '/role/add', 2, 0, '', 0, 0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (12, '', '新增权限', 2, '/auth/add', 3, 0, '', 0, 0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (13, '', '权限列表', 2, '/auth', 3, 0, '', 0,  0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (14, '', 'Banner列表', 2, '/banner', 4, 0, '', 0, 0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (15, '', '新增Banner', 2, '/banner/add', 4, 0, '', 0, 0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (16, '', '商品列表', 2, '/product', 5, 0, '', 0, 0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (17, '', '商品分类', 2, '/productCate', 5, 0, '', 0, 0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (18, '', '商品类型', 2, '/productType', 5, 0, '', 0, 0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (19, '', '订单列表', 2, '/order', 6, 0, '', 0, 0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (20, '', '导航管理', 2, '/menu', 7, 0, '', 0, 0, DEFAULT, DEFAULT);
INSERT INTO `auth` VALUES (21, '', '商城设置', 2, '/setting', 7, 0, '', 0, 0, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of jwt_blacklist
-- ----------------------------
BEGIN;
INSERT INTO `jwt_blacklist` VALUES (1, DATE_ADD(NOW(),INTERVAL '1' WEEK), DEFAULT, DEFAULT);
INSERT INTO `jwt_blacklist` VALUES (2, DATE_ADD(NOW(),INTERVAL '1' WEEK), DEFAULT, DEFAULT);
INSERT INTO `jwt_blacklist` VALUES (3, DATE_ADD(NOW(),INTERVAL '1' WEEK), DEFAULT, DEFAULT);
INSERT INTO `jwt_blacklist` VALUES (4, DATE_ADD(NOW(),INTERVAL '1' WEEK), DEFAULT, DEFAULT);
INSERT INTO `jwt_blacklist` VALUES (5, DATE_ADD(NOW(),INTERVAL '1' WEEK), DEFAULT, DEFAULT);
INSERT INTO `jwt_blacklist` VALUES (6, DATE_ADD(NOW(),INTERVAL '1' WEEK), DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of banner
-- ----------------------------
BEGIN;
INSERT INTO `banner` VALUES (1, 'banner1', 1, 'resource/upload/20201121/1605944396119681000.jpg', '/', 1000, 1, DEFAULT, DEFAULT);
INSERT INTO `banner` VALUES (2, 'banner2', 1, 'resource/upload/20201024/1603504839411539000.jpg', '', 1000, 1, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (1, 'PC电脑', '/category_1.html', 1, 1, '1', 10, 1, DEFAULT, DEFAULT);
INSERT INTO `menu` VALUES (2, '手机', '', 1, 2, '1', 10, 1, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of order
-- ----------------------------
BEGIN;
INSERT INTO `order` VALUES (1, '2020102316351850', 1, 66.66, '18899999999', '张三', '成都市高新区', '600000', 0, 0, 0, DEFAULT, DEFAULT);
INSERT INTO `order` VALUES (2, '2020112016527196', 1, 6666.66, '13888888889', '王五', '北京xxxxx区xxxxx街道xxxx号', '100000', 0, 0, 0, DEFAULT, DEFAULT);
INSERT INTO `order` VALUES (3, '2021011810071850', 1, 1999.00, '13888888889', '王五', '北京xxxxx区xxxxx街道xxxx号', '100000', 0, 0, 0, DEFAULT, DEFAULT);
INSERT INTO `order` VALUES (4, '2021011810151576', 1, 1999.00, '13888888889', '王五', '北京xxxxx区xxxxx街道xxxx号', '100000', 0, 0, 0, DEFAULT, DEFAULT);
INSERT INTO `order` VALUES (5, '2021011811371850', 1, 10995.00, '13888888889', '王五', '北京xxxxx区xxxxx街道xxxx号', '100000', 0, 0, 0, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of order_item
-- ----------------------------
BEGIN;
INSERT INTO `order_item` VALUES (1, 1, 1, 'go web from ', 1, 'resource/upload/20201023/1603440321653795000.png', 66.66, 1, '2', 'red', DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` VALUES (1, '电脑', '', '', 1, 100, 0, 1999.00, 2599.00, '', '', '第一版', 'resource/upload/20201203/1607005318313835000.jpg', '', '', '1', '', '', '', 0, 1, 1, 0, 0, 0, 1, DEFAULT, DEFAULT);
INSERT INTO `product` VALUES (2, '笔记本', '', '', 1, 1, 5, 3999.00, 4699.00, '23,24,39', ' 格式: 颜色:红色,白色,黄色 | 尺寸:41,42,43', '', 'resource/upload/20210118/1610940762322803000.jpg', '', '', '1', '', '', '', 0, 1, 1, 1, 1, 0, 1, DEFAULT, DEFAULT);
INSERT INTO `product` VALUES (3, '手机', '', '', 1, 100, 6, 999.00, 1299.00, '', '', '', 'resource/upload/20210118/1610940776037983000.jpg', '', '', '1', '', '', '', 0, 1, 1, 1, 0, 1, 1, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of product_attr
-- ----------------------------
BEGIN;
INSERT INTO `product_attr` VALUES (12, 2, 1, 1, '平板电脑', 1, '', 10, 1, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of product_cate
-- ----------------------------
BEGIN;
INSERT INTO `product_cate` VALUES (1, '手机', '', '', '', 0, '手机', '手机', '手机', 0, 1, DEFAULT, DEFAULT);
INSERT INTO `product_cate` VALUES (2, 'PC电脑', '', '', '', 0, 'PC电脑', '', '', 0, 1, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of product_color
-- ----------------------------
BEGIN;
INSERT INTO `product_color` VALUES (1, '黑色', '#ffffff', 1, 1, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of product_image
-- ----------------------------
BEGIN;
INSERT INTO `product_image` VALUES (1, 1, '/resource/upload/20201024/1603519200684359000.jpg', 0, 10, 1, DEFAULT, DEFAULT);
INSERT INTO `product_image` VALUES (2, 1, '/resource/upload/20201024/1603519285204437000.jpg', 0, 10, 1, DEFAULT, DEFAULT);
INSERT INTO `product_image` VALUES (6, 2, '/resource/upload/20210118/1610940522542324000.jpg', 0, 10, 1, DEFAULT, DEFAULT);
INSERT INTO `product_image` VALUES (7, 2, '/resource/upload/20210118/1610940522573123000.jpg', 0, 10, 1, DEFAULT, DEFAULT);
INSERT INTO `product_image` VALUES (8, 3, '/resource/upload/20210118/1610940548355473000.jpg', 0, 10, 1, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of product_type
-- ----------------------------
BEGIN;
INSERT INTO `product_type` VALUES (1, '电脑', '', 1, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of product_type_attribute
-- ----------------------------
BEGIN;
INSERT INTO `product_type_attribute` VALUES (1, 1, '平板电脑', 1, '', 1, 10, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, '超级管理员', '超级管理员', 1, DEFAULT, DEFAULT);
INSERT INTO `role` VALUES (2, '技术部', '技术部', 1, DEFAULT, DEFAULT);
INSERT INTO `role` VALUES (3, '运营部', '运营部', 1, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of role_auth
-- ----------------------------
BEGIN;
INSERT INTO `role_auth` VALUES (1, 1, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, '13888888888', 'e10adc3949ba59abbe56e057f20f883e', '', 'admin@qq.com', 1, DEFAULT, DEFAULT);
INSERT INTO `user` VALUES (2, '18389999991', 'e10adc3949ba59abbe56e057f20f883e', '127.0.0.1', '', 0, DEFAULT, DEFAULT);
COMMIT;

-- ----------------------------
-- Records of user_sms
-- ----------------------------
BEGIN;
INSERT INTO `user_sms` VALUES (1, '127.0.0.1', '18389999991', 1, '20201102', 'e178c966721a75236355d935ac3dd9ff', DEFAULT, DEFAULT);
INSERT INTO `user_sms` VALUES (2, '127.0.0.1', '13889999992', 1, '20201119', '200f0a43fbc9c0ae40f26432269cae91', DEFAULT, DEFAULT);
INSERT INTO `user_sms` VALUES (3, '127.0.0.1', '13889999999', 1, '20201119', '5abc5dca4b31c1cfe222693ee1c5bd1c', DEFAULT, DEFAULT);
COMMIT;