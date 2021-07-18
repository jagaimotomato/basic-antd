/*
 Navicat Premium Data Transfer

 Source Server         : 115.29.226.247
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : 115.29.226.247:3306
 Source Schema         : basic_antd

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 07/06/2021 16:31:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique_index`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 242 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (215, 'p', 'admin', '/department', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (217, 'p', 'admin', '/department', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (216, 'p', 'admin', '/department/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (218, 'p', 'admin', '/department/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (225, 'p', 'admin', '/dict/data', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (227, 'p', 'admin', '/dict/data', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (226, 'p', 'admin', '/dict/data/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (229, 'p', 'admin', '/dict/data/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (220, 'p', 'admin', '/dict/type', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (222, 'p', 'admin', '/dict/type', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (221, 'p', 'admin', '/dict/type/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (224, 'p', 'admin', '/dict/type/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (235, 'p', 'admin', '/info/nav', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (234, 'p', 'admin', '/info/user', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (239, 'p', 'admin', '/list/dictDataList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (238, 'p', 'admin', '/list/dictTypeList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (230, 'p', 'admin', '/list/permissionList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (236, 'p', 'admin', '/list/roleList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (219, 'p', 'admin', '/page/department', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (228, 'p', 'admin', '/page/dictData', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (223, 'p', 'admin', '/page/dictType', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (240, 'p', 'admin', '/page/loginLog', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (241, 'p', 'admin', '/page/operationLog', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (213, 'p', 'admin', '/page/permission', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (208, 'p', 'admin', '/page/role', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (203, 'p', 'admin', '/page/user', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (210, 'p', 'admin', '/permission', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (212, 'p', 'admin', '/permission', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (211, 'p', 'admin', '/permission/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (214, 'p', 'admin', '/permission/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (205, 'p', 'admin', '/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (207, 'p', 'admin', '/role', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (206, 'p', 'admin', '/role/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (209, 'p', 'admin', '/role/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (237, 'p', 'admin', '/tree/department', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (231, 'p', 'admin', '/tree/permission', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (232, 'p', 'admin', '/tree/permission/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (200, 'p', 'admin', '/user', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (202, 'p', 'admin', '/user', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (201, 'p', 'admin', '/user/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (204, 'p', 'admin', '/user/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (233, 'p', 'admin', '/user/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (93, 'p', 'test', '/department', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (95, 'p', 'test', '/department', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (94, 'p', 'test', '/department/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (96, 'p', 'test', '/department/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (103, 'p', 'test', '/dict/data', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (105, 'p', 'test', '/dict/data', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (104, 'p', 'test', '/dict/data/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (107, 'p', 'test', '/dict/data/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (98, 'p', 'test', '/dict/type', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (100, 'p', 'test', '/dict/type', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (99, 'p', 'test', '/dict/type/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (102, 'p', 'test', '/dict/type/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (113, 'p', 'test', '/info/nav', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (112, 'p', 'test', '/info/user', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (117, 'p', 'test', '/list/dictDataList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (116, 'p', 'test', '/list/dictTypeList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (108, 'p', 'test', '/list/permissionList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (114, 'p', 'test', '/list/roleList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (97, 'p', 'test', '/page/department', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (106, 'p', 'test', '/page/dictData', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (101, 'p', 'test', '/page/dictType', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (91, 'p', 'test', '/page/permission', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (86, 'p', 'test', '/page/role', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (81, 'p', 'test', '/page/user', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (88, 'p', 'test', '/permission', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (90, 'p', 'test', '/permission', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (89, 'p', 'test', '/permission/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (92, 'p', 'test', '/permission/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (83, 'p', 'test', '/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (85, 'p', 'test', '/role', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (84, 'p', 'test', '/role/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (87, 'p', 'test', '/role/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (115, 'p', 'test', '/tree/department', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (109, 'p', 'test', '/tree/permission', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (110, 'p', 'test', '/tree/permission/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (78, 'p', 'test', '/user', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (80, 'p', 'test', '/user', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (79, 'p', 'test', '/user/:id', 'DEL', '', '', '');
INSERT INTO `casbin_rule` VALUES (82, 'p', 'test', '/user/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (111, 'p', 'test', '/user/status', 'PUT', '', '', '');

-- ----------------------------
-- Table structure for department
-- ----------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department`  (
  `department_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `parent_id` bigint(20) NULL DEFAULT NULL,
  `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `describe` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`department_id`) USING BTREE,
  INDEX `idx_department_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of department
-- ----------------------------
INSERT INTO `department` VALUES (2, '淘宝', 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `department` VALUES (3, '天猫', 2, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `department` VALUES (4, 'bilibili', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `department` VALUES (5, '腾讯', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `department` VALUES (6, '支付宝', 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `department` VALUES (10, 'test2', 5, '', '', '', '', '2021-02-02 00:36:10.095', '2021-02-02 00:36:10.095', NULL);
INSERT INTO `department` VALUES (11, 'test', 6, '', '', '', '', '2021-02-02 00:36:55.444', '2021-02-02 00:36:55.444', NULL);
INSERT INTO `department` VALUES (14, 'test6', 13, '', '', '', '', '2021-02-02 12:01:41.570', '2021-02-02 12:01:41.570', NULL);
INSERT INTO `department` VALUES (15, '测试', 0, '', '', '', '', '2021-02-02 14:22:12.336', '2021-02-07 16:49:07.834', NULL);
INSERT INTO `department` VALUES (16, 'ceshi', 4, '', '', '1', '', '2021-02-07 16:43:14.610', '2021-02-07 16:43:14.610', NULL);

-- ----------------------------
-- Table structure for department_user
-- ----------------------------
DROP TABLE IF EXISTS `department_user`;
CREATE TABLE `department_user`  (
  `user_id` bigint(20) NULL DEFAULT NULL,
  `department_id` bigint(20) NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of department_user
-- ----------------------------
INSERT INTO `department_user` VALUES (1, 2);
INSERT INTO `department_user` VALUES (1, 3);
INSERT INTO `department_user` VALUES (1, 4);
INSERT INTO `department_user` VALUES (1, 4);

-- ----------------------------
-- Table structure for dict_data
-- ----------------------------
DROP TABLE IF EXISTS `dict_data`;
CREATE TABLE `dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT,
  `label` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `value` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`dict_code`) USING BTREE,
  INDEX `idx_dict_data_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dict_data
-- ----------------------------
INSERT INTO `dict_data` VALUES (1, '新增', 'add', 'sys_action', '1', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `dict_data` VALUES (2, '删除', 'del', 'sys_action', '1', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `dict_data` VALUES (3, '修改', 'edit', 'sys_action', '1', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `dict_data` VALUES (4, '查询', 'query', 'sys_action', '1', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `dict_data` VALUES (5, '详情', 'get', 'sys_action', '1', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `dict_data` VALUES (6, '启用', 'enable', 'sys_action', '1', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `dict_data` VALUES (7, '禁用', 'disable', 'sys_action', '1', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `dict_data` VALUES (8, '导入', 'import', 'sys_action', '1', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `dict_data` VALUES (9, '导出', 'export', 'sys_action', '1', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `dict_data` VALUES (10, 'test', 'deladsfsa', 'sys_action', '0', '1', '1', '2021-02-07 19:18:06.060', '2021-02-07 19:18:11.303', '2021-02-07 19:18:21.502');
INSERT INTO `dict_data` VALUES (11, '男', '1', 'sys_sex', '1', '1', '', '2021-05-19 18:07:11.494', '2021-05-19 18:07:11.494', NULL);
INSERT INTO `dict_data` VALUES (12, '女', '0', 'sys_sex', '1', '1', '', '2021-05-19 18:07:23.969', '2021-05-19 18:07:23.969', NULL);

-- ----------------------------
-- Table structure for dict_type
-- ----------------------------
DROP TABLE IF EXISTS `dict_type`;
CREATE TABLE `dict_type`  (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`dict_id`) USING BTREE,
  INDEX `idx_dict_type_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dict_type
-- ----------------------------
INSERT INTO `dict_type` VALUES (1, '系统操作', 'sys_action', '1', NULL, '1', NULL, '2021-02-07 18:38:36.256', NULL);
INSERT INTO `dict_type` VALUES (2, '测试', '222', '1', '1', '', '2021-02-07 18:38:46.696', '2021-02-07 18:38:46.696', '2021-02-07 18:38:51.083');
INSERT INTO `dict_type` VALUES (3, '性别', 'sys_sex', '1', '1', '', '2021-05-19 18:06:41.111', '2021-05-19 18:06:41.111', NULL);

-- ----------------------------
-- Table structure for login_log
-- ----------------------------
DROP TABLE IF EXISTS `login_log`;
CREATE TABLE `login_log`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `username` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `mobile` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `ip` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `browser` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `os` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `platform` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `login_time` timestamp(0) NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `msg` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_login_log_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 97 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for operation_log
-- ----------------------------
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作模块',
  `business_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作类型',
  `business_types` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'BusinessTypes',
  `method` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '函数',
  `request_method` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求方式',
  `operator_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作类型',
  `oper_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作者',
  `oper_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '访问地址',
  `oper_ip` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '客户端ip',
  `oper_location` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '访问位置',
  `oper_param` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求参数',
  `code` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '状态码',
  `status` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作状态',
  `oper_time` timestamp(0) NULL DEFAULT NULL COMMENT '操作时间',
  `json_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '返回数据',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `duration_time` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '耗时',
  `user_agent` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ua',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_operation_log_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 157 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`  (
  `permission_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `actions` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `target` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `hide_children` tinyint(4) NULL DEFAULT NULL,
  `hidden_header_content` tinyint(4) NULL DEFAULT NULL,
  `show` tinyint(4) NULL DEFAULT NULL,
  `redirect` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `describe` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `parent_id` bigint(20) NULL DEFAULT NULL,
  `component` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `icon` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `path` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `sort` bigint(20) NULL DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`permission_id`) USING BTREE,
  INDEX `idx_permission_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 135 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (1, '系统管理', 'system', NULL, NULL, 0, 0, 1, '/system/permission', NULL, 0, 'PageView', 'desktop', '/system', 1, NULL, '1', NULL, '2021-05-19 17:44:18.537', NULL);
INSERT INTO `permission` VALUES (2, '权限管理', 'permission', NULL, NULL, 0, 0, 1, NULL, NULL, 1, 'Permission', 'lock', '/system/permission', 0, NULL, '1', NULL, '2021-02-07 15:53:34.182', NULL);
INSERT INTO `permission` VALUES (3, '用户管理', 'user', NULL, NULL, 0, 0, 1, '', NULL, 1, 'User', 'user', '/system/user', 0, NULL, '1', NULL, '2021-02-06 14:43:02.965', NULL);
INSERT INTO `permission` VALUES (4, '角色管理', 'role', NULL, NULL, 0, 0, 1, '', NULL, 1, 'Role', 'deployment-unit', '/system/role', 0, NULL, '1', NULL, '2021-02-06 14:43:33.003', NULL);
INSERT INTO `permission` VALUES (5, '新增', '', 'add', '', 0, 0, 0, '', '新增', 2, '', '', '', 0, '', '', '2021-02-05 23:41:21.109', '2021-02-05 23:41:21.109', '2021-02-07 15:53:34.107');
INSERT INTO `permission` VALUES (6, '删除', '', 'del', '', 0, 0, 0, '', '删除', 2, '', '', '', 0, '', '', '2021-02-05 23:41:21.109', '2021-02-05 23:41:21.109', '2021-02-07 15:53:34.107');
INSERT INTO `permission` VALUES (7, '修改', '', 'edit', '', 0, 0, 0, '', '修改', 2, '', '', '', 0, '', '', '2021-02-05 23:41:21.109', '2021-02-05 23:41:21.109', '2021-02-07 15:53:34.107');
INSERT INTO `permission` VALUES (8, '查询', '', 'query', '', 0, 0, 0, '', '查询', 2, '', '', '', 0, '', '', '2021-02-05 23:41:21.109', '2021-02-05 23:41:21.109', '2021-02-07 15:53:34.107');
INSERT INTO `permission` VALUES (9, '详情', '', 'get', '', 0, 0, 0, '', '详情', 2, '', '', '', 0, '', '', '2021-02-05 23:41:21.109', '2021-02-05 23:41:21.109', '2021-02-07 15:53:34.107');
INSERT INTO `permission` VALUES (10, '首页', 'dashboard', NULL, NULL, 0, 0, 1, NULL, NULL, 0, 'Workplace', 'slack', '/dashboard', NULL, NULL, '1', NULL, '2021-02-06 14:44:09.901', NULL);
INSERT INTO `permission` VALUES (11, '导入', NULL, 'import', NULL, 0, 0, 0, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2021-02-07 15:53:34.107');
INSERT INTO `permission` VALUES (67, '新增', 'add_user', 'POST', '', 0, 0, 0, '', '新增', 3, '', '', '/user', 0, '', '1', '2021-02-06 12:22:00.311', '2021-04-25 16:35:19.689', NULL);
INSERT INTO `permission` VALUES (68, '删除', 'del_user', 'DEL', '', 0, 0, 0, '', '删除', 3, '', '', '/user/:id', 0, '', '1', '2021-02-06 12:22:00.311', '2021-04-25 16:35:26.334', NULL);
INSERT INTO `permission` VALUES (69, '修改', 'edit_user', 'PUT', '', 0, 0, 0, '', '修改', 3, '', '', '/user', 0, '', '1', '2021-02-06 12:22:00.311', '2021-04-25 16:35:37.173', NULL);
INSERT INTO `permission` VALUES (70, '查询', 'query_user', 'GET', '', 0, 0, 0, '', '查询用户列表', 3, '', '', '/page/user', 0, '', '1', '2021-02-06 12:22:00.311', '2021-04-25 16:35:46.156', NULL);
INSERT INTO `permission` VALUES (71, '详情', 'detail_user', 'GET', '', 0, 0, 0, '', '详情', 3, '', '', '/user/:id', 0, '', '1', '2021-02-06 12:22:00.311', '2021-04-25 16:35:55.176', NULL);
INSERT INTO `permission` VALUES (72, '新增', 'add_role', 'POST', '', 0, 0, 0, '', '新增', 4, '', '', '/role', 0, '', '1', '2021-02-06 12:22:09.425', '2021-04-25 16:36:31.442', NULL);
INSERT INTO `permission` VALUES (73, '删除', 'del_role', 'DEL', '', 0, 0, 0, '', '删除', 4, '', '', '/role/:id', 0, '', '1', '2021-02-06 12:22:09.425', '2021-04-25 16:36:51.882', NULL);
INSERT INTO `permission` VALUES (74, '修改', 'edit_role', 'PUT', '', 0, 0, 0, '', '修改', 4, '', '', '/role', 0, '', '2', '2021-02-06 12:22:09.425', '2021-04-25 17:21:29.748', NULL);
INSERT INTO `permission` VALUES (75, '查询', 'query_role', 'GET', '', 0, 0, 0, '', '查询', 4, '', '', '/page/role', 0, '', '1', '2021-02-06 12:22:09.425', '2021-04-25 16:38:58.572', NULL);
INSERT INTO `permission` VALUES (76, '详情', 'detail_role', 'GET', '', 0, 0, 0, '', '详情', 4, '', '', '/role/:id', 0, '', '2', '2021-02-06 12:22:09.425', '2021-04-25 17:21:46.913', NULL);
INSERT INTO `permission` VALUES (77, '新增', 'add_permission', 'POST', '', 0, 0, 0, '', '新增', 2, '', '', '/permission', 0, '', '1', '2021-02-07 15:53:34.176', '2021-04-25 16:34:07.503', NULL);
INSERT INTO `permission` VALUES (78, '删除', 'del_permission', 'DEL', '', 0, 0, 0, '', '删除', 2, '', '', '/permission/:id', 0, '', '1', '2021-02-07 15:53:34.176', '2021-04-25 16:34:19.938', NULL);
INSERT INTO `permission` VALUES (79, '修改', 'edit_permission', 'PUT', '', 0, 0, 0, '', '修改', 2, '', '', '/permission', 0, '', '1', '2021-02-07 15:53:34.176', '2021-04-25 16:34:28.676', NULL);
INSERT INTO `permission` VALUES (80, '查询', 'query_permission', 'GET', '', 0, 0, 0, '', '查询', 2, '', '', '/page/permission', 0, '', '1', '2021-02-07 15:53:34.176', '2021-04-25 16:34:37.940', NULL);
INSERT INTO `permission` VALUES (81, '详情', 'detail_permission', 'GET', '', 0, 0, 0, '', '详情', 2, '', '', '/permission/:id', 0, '', '1', '2021-02-07 15:53:34.176', '2021-04-25 16:34:55.701', NULL);
INSERT INTO `permission` VALUES (93, '部门管理', 'department', '', '', 0, 0, 1, '', '', 1, 'Department', 'cluster', '/system/department', 0, '1', '', '2021-02-07 16:20:15.135', '2021-02-07 16:20:15.135', NULL);
INSERT INTO `permission` VALUES (94, '新增', 'add_department', 'POST', '', 0, 0, 0, '', '新增', 93, '', '', '/department', 0, '', '1', '2021-02-07 16:20:15.177', '2021-04-25 16:51:18.471', NULL);
INSERT INTO `permission` VALUES (95, '删除', 'del_department', 'DEL', '', 0, 0, 0, '', '删除', 93, '', '', '/department/:id', 0, '', '1', '2021-02-07 16:20:15.177', '2021-04-25 16:51:38.624', NULL);
INSERT INTO `permission` VALUES (96, '修改', 'edit_department', 'PUT', '', 0, 0, 0, '', '修改', 93, '', '', '/department', 0, '', '1', '2021-02-07 16:20:15.177', '2021-04-25 16:51:59.299', NULL);
INSERT INTO `permission` VALUES (97, '详情', 'detail_department', 'GET', '', 0, 0, 0, '', '详情', 93, '', '', '/department/:id', 0, '', '1', '2021-02-07 16:20:15.177', '2021-04-25 16:52:19.700', NULL);
INSERT INTO `permission` VALUES (98, '查询', 'query_department', 'GET', '', 0, 0, 0, '', '查询', 93, '', '', '/page/department', 0, '', '1', '2021-02-07 16:20:15.177', '2021-04-25 16:52:46.847', NULL);
INSERT INTO `permission` VALUES (99, '字典管理', 'dict', '', '', 0, 0, 1, '/system/dicttype', '', 1, 'RouteView', 'build', '', 0, '1', '', '2021-02-07 18:33:17.511', '2021-02-07 18:33:17.511', NULL);
INSERT INTO `permission` VALUES (100, '字典类型', 'dicttype', '', '', 0, 0, 1, '', '', 99, 'DictType', '', '/system/dicttype', 0, '1', '', '2021-02-07 18:34:21.053', '2021-02-07 18:34:21.053', NULL);
INSERT INTO `permission` VALUES (101, '新增', 'add_dicttype', 'POST', '', 0, 0, 0, '', '新增', 100, '', '', '/dict/type', 0, '', '1', '2021-02-07 18:34:21.150', '2021-04-25 16:57:59.697', NULL);
INSERT INTO `permission` VALUES (102, '删除', 'del_dicttype', 'DEL', '', 0, 0, 0, '', '删除', 100, '', '', '/dict/type/:id', 0, '', '1', '2021-02-07 18:34:21.150', '2021-04-25 16:58:19.763', NULL);
INSERT INTO `permission` VALUES (103, '修改', 'edit_dicttype', 'PUT', '', 0, 0, 0, '', '修改', 100, '', '', '/dict/type', 0, '', '1', '2021-02-07 18:34:21.150', '2021-04-25 16:58:40.326', NULL);
INSERT INTO `permission` VALUES (104, '查询', 'query_dicttype', 'GET', '', 0, 0, 0, '', '查询', 100, '', '', '/page/dictType', 0, '', '1', '2021-02-07 18:34:21.150', '2021-04-25 16:59:16.437', NULL);
INSERT INTO `permission` VALUES (105, '详情', 'detail_dicttype', 'GET', '', 0, 0, 0, '', '详情', 100, '', '', '/dict/type/:id', 0, '', '1', '2021-02-07 18:34:21.150', '2021-04-25 16:59:37.742', NULL);
INSERT INTO `permission` VALUES (106, '字典键值', 'dictdata', '', '', 0, 0, 1, '', '', 99, 'DictData', '', '/system/dictdata', 0, '1', '', '2021-02-07 19:15:22.429', '2021-02-07 19:15:22.429', NULL);
INSERT INTO `permission` VALUES (107, '新增', 'add_dictdata', 'POST', '', 0, 0, 0, '', '新增', 106, '', '', '/dict/data', 0, '', '1', '2021-02-07 19:15:22.529', '2021-04-25 17:01:02.711', NULL);
INSERT INTO `permission` VALUES (108, '删除', 'del_dictdata', 'DEL', '', 0, 0, 0, '', '删除', 106, '', '', '/dict/data/:id', 0, '', '1', '2021-02-07 19:15:22.529', '2021-04-25 17:01:26.632', NULL);
INSERT INTO `permission` VALUES (109, '修改', 'edit_dictdata', 'PUT', '', 0, 0, 0, '', '修改', 106, '', '', '/dict/data', 0, '', '1', '2021-02-07 19:15:22.529', '2021-04-25 17:01:53.234', NULL);
INSERT INTO `permission` VALUES (110, '查询', 'query_dictdata', 'GET', '', 0, 0, 0, '', '查询', 106, '', '', '/page/dictData', 0, '', '1', '2021-02-07 19:15:22.529', '2021-04-25 17:03:12.048', NULL);
INSERT INTO `permission` VALUES (111, '详情', 'detail_dictdata', 'GET', '', 0, 0, 0, '', '详情', 106, '', '', '/dict/data/:id', 0, '', '1', '2021-02-07 19:15:22.529', '2021-04-25 17:02:21.496', NULL);
INSERT INTO `permission` VALUES (112, '日志管理', 'log', '', '', 0, 0, 1, '/system/loginlog', '', 1, 'RouteView', 'small-dash', '/system/log', 0, '1', '1', '2021-02-07 19:47:32.118', '2021-02-07 19:54:59.666', NULL);
INSERT INTO `permission` VALUES (113, '登录日志', 'loginlog', '', '', 0, 0, 1, '', '', 112, 'LoginLog', '', '/system/loginlog', 0, '1', '', '2021-02-07 19:48:34.767', '2021-02-07 19:48:34.767', NULL);
INSERT INTO `permission` VALUES (114, '操作日志', 'operationlog', '', '', 0, 0, 1, '', '', 112, 'OperationLog', '', '/system/operationlog', 0, '1', '', '2021-02-07 19:49:09.967', '2021-02-07 19:49:09.967', NULL);
INSERT INTO `permission` VALUES (115, '权限列表', '', 'GET', '', 0, 0, 1, '', '权限列表', 2, '', '', '/list/permissionList', 0, '1', '', '2021-04-25 15:13:42.587', '2021-04-25 15:13:42.587', NULL);
INSERT INTO `permission` VALUES (116, '权限树', '', 'GET', '', 0, 0, 1, '', '权限树', 2, '', '', '/tree/permission', 0, '1', '', '2021-04-25 15:14:22.448', '2021-04-25 15:14:22.448', NULL);
INSERT INTO `permission` VALUES (117, '角色权限树', '', 'GET', '', 0, 0, 1, '', '', 2, '', '', '/tree/permission/:id', 0, '1', '', '2021-04-25 15:14:51.447', '2021-04-25 15:14:51.447', NULL);
INSERT INTO `permission` VALUES (118, '修改用户状态', '', 'PUT', '', 0, 0, 1, '', '', 3, '', '', '/user/status', 0, '1', '', '2021-04-25 15:18:25.557', '2021-04-25 15:18:25.557', NULL);
INSERT INTO `permission` VALUES (119, '用户信息', '', 'GET', '', 0, 0, 1, '', '登录时获取用户信息 userinfo', 3, '', '', '/info/user', 0, '1', '', '2021-04-25 15:20:13.894', '2021-04-25 15:20:13.894', NULL);
INSERT INTO `permission` VALUES (120, '用户菜单', '', 'GET', '', 0, 0, 1, '', '登录时获取用户菜单', 3, '', '', '/info/nav', 0, '1', '', '2021-04-25 15:21:01.948', '2021-04-25 15:21:01.948', NULL);
INSERT INTO `permission` VALUES (121, '角色列表', '', 'GET', '', 0, 0, 1, '', '', 4, '', '', '/list/roleList', 0, '1', '', '2021-04-25 15:32:29.779', '2021-04-25 15:32:29.779', NULL);
INSERT INTO `permission` VALUES (122, '部门树', '', 'GET', '', 0, 0, 1, '', '', 93, '', '', '/tree/department', 0, '1', '', '2021-04-25 16:53:26.864', '2021-04-25 16:53:26.864', NULL);
INSERT INTO `permission` VALUES (123, '字典类型列表', 'list_dicttype', 'GET', '', 0, 0, 1, '', '', 100, '', '', '/list/dictTypeList', 0, '1', '', '2021-04-25 17:00:27.223', '2021-04-25 17:00:27.223', NULL);
INSERT INTO `permission` VALUES (124, '字典键值列表', '', 'GET', '', 0, 0, 1, '', '', 106, '', '', '/list/dictDataList', 0, '1', '', '2021-04-25 17:03:45.167', '2021-04-25 17:03:45.167', NULL);
INSERT INTO `permission` VALUES (125, '', 'loginlog', 'GET', '', 0, 0, 1, '', '', 113, '', '', '/page/loginLog', 0, '1', '1', '2021-05-19 14:35:51.409', '2021-05-19 14:36:23.612', NULL);
INSERT INTO `permission` VALUES (126, '', 'operationLog', 'GET', '', 0, 0, 1, '', '', 114, '', '', '/page/operationLog', 0, '1', '', '2021-05-19 14:37:16.482', '2021-05-19 14:37:16.482', NULL);
INSERT INTO `permission` VALUES (127, '账号设置', 'account', '', '', 0, 0, 1, '/account/center', '', 0, 'RouteView', 'setting', '/account', 2, '1', '1', '2021-05-19 14:52:08.249', '2021-05-19 17:44:23.380', NULL);
INSERT INTO `permission` VALUES (128, '个人中心', 'center', '', '', 0, 0, 1, '', '', 127, 'AccountCenter', '', '/account/center', 0, '1', '1', '2021-05-19 15:03:18.878', '2021-05-19 16:18:25.761', NULL);
INSERT INTO `permission` VALUES (129, '个人设置', 'settings', '', '', 1, 0, 1, '/account/settings/basic', '', 127, 'AccountSettings', '', '/account/settings', 0, '1', '1', '2021-05-19 15:05:25.631', '2021-05-19 16:19:04.461', NULL);
INSERT INTO `permission` VALUES (130, '基本设置', 'BasicSettings', '', '', 0, 0, 0, '', '', 129, 'BasicSettings', '', '/account/settings/basic', 0, '1', '1', '2021-05-19 15:07:47.833', '2021-05-19 16:19:28.889', NULL);
INSERT INTO `permission` VALUES (131, '安全设置', 'SecuritySettings', '', '', 0, 0, 0, '', '', 129, 'SecuritySettings', '', '/account/settings/security', 0, '1', '1', '2021-05-19 15:09:18.563', '2021-05-19 16:19:48.482', NULL);
INSERT INTO `permission` VALUES (132, '自定义设置', 'CustomSettings', '', '', 0, 0, 0, '', '', 129, 'CustomSettings', '', '/account/settings/custom', 0, '1', '', '2021-05-19 16:20:21.522', '2021-05-19 16:20:21.522', NULL);
INSERT INTO `permission` VALUES (133, '绑定设置', 'BindingSettings', '', '', 0, 0, 0, '', '', 129, 'BindingSettings', '', '/account/settings/binding', 0, '1', '', '2021-05-19 16:20:46.444', '2021-05-19 16:20:46.444', NULL);
INSERT INTO `permission` VALUES (134, '通知设置', 'NotificationSettings', '', '', 1, 0, 0, '', '', 129, 'NotificationSettings', '', '/account/settings/notification', 0, '1', '', '2021-05-19 16:21:08.316', '2021-05-19 16:21:08.316', NULL);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `describe` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` bigint(20) NULL DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`) USING BTREE,
  INDEX `idx_role_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 'admin', '管理员', '拥有所有权限', '1', 1, NULL, '1', NULL, '2021-05-19 15:10:03.741', NULL);

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission`  (
  `role_id` bigint(20) NULL DEFAULT NULL,
  `permission_id` bigint(20) NULL DEFAULT NULL,
  `role_key` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_permission
-- ----------------------------
INSERT INTO `role_permission` VALUES (1, 1, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 2, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 3, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 4, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 10, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 67, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 68, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 69, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 70, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 71, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 72, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 73, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 74, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 75, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 76, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 77, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 78, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 79, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 80, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 81, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 93, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 94, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 95, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 96, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 97, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 98, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 99, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 100, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 101, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 102, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 103, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 104, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 105, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 106, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 107, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 108, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 109, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 110, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 111, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 112, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 113, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 114, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 115, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 116, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 117, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 118, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 119, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 120, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 121, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 122, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 123, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 124, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 125, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 126, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 127, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 128, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 129, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 130, 'admin', '', '');
INSERT INTO `role_permission` VALUES (1, 131, 'admin', '', '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `real_name` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `avatar` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `email` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `last_login_ip` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `last_login_time` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `sex` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `bio` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `salt` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `login_times` int(11) NULL DEFAULT NULL,
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'test', 'test', 'assets/uploads/avatar/zhuwei/aefd3430-0783-432e-903b-02db2ca8fb00.jpg', '艾伦耶格尔', '878382660@qq.com', 'jLQ/xUKtNEnkjepl5JYMd8jWdMUnWt4HwfXCBUyAybE=', NULL, NULL, NULL, '1', '测试呀', '7249260', NULL, '1', NULL, 'test', NULL, '2021-05-19 18:07:40.343', NULL);

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `user_id` bigint(20) NULL DEFAULT NULL,
  `role_id` bigint(20) NULL DEFAULT NULL,
  `role_key` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (1, 1, 'admin', '', '');

SET FOREIGN_KEY_CHECKS = 1;
