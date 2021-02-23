/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : basic_antd

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 23/02/2021 14:41:49
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
  UNIQUE INDEX `unique_index`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (1, 'p', 'admin', '/info/user', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (24, 'p', 'test', '/department/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (25, 'p', 'test', '/dict/data/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (26, 'p', 'test', '/dict/type/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (3, 'p', 'test', '/info/nav', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (2, 'p', 'test', '/info/user', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (15, 'p', 'test', '/list/dictDataList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (16, 'p', 'test', '/list/dictTypeList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (13, 'p', 'test', '/list/permissionList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (14, 'p', 'test', '/list/roleList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (20, 'p', 'test', '/other/dictData/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (7, 'p', 'test', '/page/department', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (12, 'p', 'test', '/page/dictData', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (10, 'p', 'test', '/page/dictType', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (11, 'p', 'test', '/page/dictType', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (8, 'p', 'test', '/page/loginLog', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (9, 'p', 'test', '/page/operationLog', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (5, 'p', 'test', '/page/permission', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (6, 'p', 'test', '/page/role', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (4, 'p', 'test', '/page/user', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (22, 'p', 'test', '/permission/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (23, 'p', 'test', '/role/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (19, 'p', 'test', '/tree/department', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (17, 'p', 'test', '/tree/permission', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (18, 'p', 'test', '/tree/permission/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (21, 'p', 'test', '/user/:id', 'GET', NULL, NULL, NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

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
INSERT INTO `department_user` VALUES (3, 4);
INSERT INTO `department_user` VALUES (3, 10);
INSERT INTO `department_user` VALUES (3, 15);

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
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dict_type
-- ----------------------------
INSERT INTO `dict_type` VALUES (1, '系统操作', 'sys_action', '1', NULL, '1', NULL, '2021-02-07 18:38:36.256', NULL);
INSERT INTO `dict_type` VALUES (2, '测试', '222', '1', '1', '', '2021-02-07 18:38:46.696', '2021-02-07 18:38:46.696', '2021-02-07 18:38:51.083');

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
) ENGINE = InnoDB AUTO_INCREMENT = 76 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of login_log
-- ----------------------------
INSERT INTO `login_log` VALUES (1, '2021-02-05 17:42:07.720', '2021-02-05 17:42:07.720', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:42:08', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (2, '2021-02-05 17:42:22.538', '2021-02-05 17:42:22.538', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:42:23', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (3, '2021-02-05 17:43:06.964', '2021-02-05 17:43:06.964', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:43:07', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (4, '2021-02-05 17:45:45.388', '2021-02-05 17:45:45.388', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:45:45', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (5, '2021-02-05 17:48:50.143', '2021-02-05 17:48:50.143', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:48:50', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (6, '2021-02-05 17:48:56.590', '2021-02-05 17:48:56.590', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:48:57', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (7, '2021-02-05 17:49:33.506', '2021-02-05 17:49:33.506', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:49:34', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (8, '2021-02-05 17:51:46.508', '2021-02-05 17:51:46.508', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:51:47', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (9, '2021-02-05 17:53:20.187', '2021-02-05 17:53:20.187', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:53:20', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (10, '2021-02-05 17:58:33.485', '2021-02-05 17:58:33.485', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:58:33', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (11, '2021-02-05 17:59:05.877', '2021-02-05 17:59:05.877', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 17:59:06', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (12, '2021-02-05 18:02:24.099', '2021-02-05 18:02:24.099', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 18:02:24', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (13, '2021-02-05 18:04:50.393', '2021-02-05 18:04:50.393', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 18:04:50', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (14, '2021-02-05 18:05:18.913', '2021-02-05 18:05:18.913', NULL, '', '', 'testt\'s', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 18:05:19', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '用户不存在');
INSERT INTO `login_log` VALUES (15, '2021-02-05 18:05:33.764', '2021-02-05 18:05:33.764', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 18:05:34', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (16, '2021-02-05 18:05:46.653', '2021-02-05 18:05:46.653', NULL, '', '', 'test', '0', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 18:05:47', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '密码错误');
INSERT INTO `login_log` VALUES (17, '2021-02-05 18:08:56.483', '2021-02-05 18:08:56.483', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-05 18:08:56', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (18, '2021-02-20 17:14:49.202', '2021-02-20 17:14:49.202', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-20 17:14:49', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (19, '2021-02-20 17:15:21.775', '2021-02-20 17:15:21.775', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-20 17:15:22', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (20, '2021-02-20 17:15:21.924', '2021-02-20 17:15:21.924', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-20 17:15:22', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (21, '2021-02-20 17:17:16.375', '2021-02-20 17:17:16.375', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-20 17:17:16', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (22, '2021-02-20 17:17:16.500', '2021-02-20 17:17:16.500', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-20 17:17:16', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (23, '2021-02-20 17:35:15.169', '2021-02-20 17:35:15.169', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-20 17:35:15', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (24, '2021-02-20 17:35:15.349', '2021-02-20 17:35:15.349', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-20 17:35:15', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (25, '2021-02-20 17:36:52.781', '2021-02-20 17:36:52.781', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-20 17:36:53', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (26, '2021-02-20 17:36:52.989', '2021-02-20 17:36:52.989', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-20 17:36:53', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (27, '2021-02-22 10:28:47.133', '2021-02-22 10:28:47.133', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 10:28:47', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (28, '2021-02-22 10:28:47.264', '2021-02-22 10:28:47.264', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 10:28:47', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (29, '2021-02-22 10:28:59.722', '2021-02-22 10:28:59.722', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 10:29:00', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (30, '2021-02-22 10:28:59.837', '2021-02-22 10:28:59.837', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 10:29:00', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (31, '2021-02-22 10:50:40.386', '2021-02-22 10:50:40.386', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 10:50:40', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (32, '2021-02-22 10:50:40.529', '2021-02-22 10:50:40.529', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 10:50:41', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (33, '2021-02-22 10:58:16.398', '2021-02-22 10:58:16.398', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 10:58:16', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (34, '2021-02-22 10:58:16.570', '2021-02-22 10:58:16.570', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 10:58:17', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (35, '2021-02-22 11:00:08.770', '2021-02-22 11:00:08.770', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:00:09', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (36, '2021-02-22 11:00:08.912', '2021-02-22 11:00:08.912', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:00:09', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (37, '2021-02-22 11:02:19.496', '2021-02-22 11:02:19.496', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:02:19', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (38, '2021-02-22 11:02:19.649', '2021-02-22 11:02:19.649', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:02:20', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (39, '2021-02-22 11:03:37.746', '2021-02-22 11:03:37.746', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:03:38', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (40, '2021-02-22 11:04:24.195', '2021-02-22 11:04:24.195', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:04:24', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (41, '2021-02-22 11:04:29.608', '2021-02-22 11:04:29.608', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:04:30', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (42, '2021-02-22 11:04:29.751', '2021-02-22 11:04:29.751', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:04:30', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (43, '2021-02-22 11:14:35.067', '2021-02-22 11:14:35.067', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:14:35', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (44, '2021-02-22 11:14:35.190', '2021-02-22 11:14:35.190', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:14:35', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (45, '2021-02-22 11:19:38.877', '2021-02-22 11:19:38.877', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:19:39', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (46, '2021-02-22 11:19:39.057', '2021-02-22 11:19:39.057', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:19:39', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (47, '2021-02-22 11:24:45.860', '2021-02-22 11:24:45.860', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:24:46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (48, '2021-02-22 11:24:46.057', '2021-02-22 11:24:46.057', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:24:46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (49, '2021-02-22 11:26:15.844', '2021-02-22 11:26:15.844', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:26:16', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (50, '2021-02-22 11:26:16.060', '2021-02-22 11:26:16.060', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:26:16', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (51, '2021-02-22 11:28:04.995', '2021-02-22 11:28:04.995', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:28:05', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (52, '2021-02-22 11:28:05.149', '2021-02-22 11:28:05.149', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:28:05', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (53, '2021-02-22 11:28:16.891', '2021-02-22 11:28:16.891', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:28:17', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (54, '2021-02-22 11:28:17.039', '2021-02-22 11:28:17.039', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:28:17', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (55, '2021-02-22 11:28:29.554', '2021-02-22 11:28:29.554', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:28:30', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (56, '2021-02-22 11:28:29.769', '2021-02-22 11:28:29.769', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:28:30', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (57, '2021-02-22 11:29:04.594', '2021-02-22 11:29:04.594', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:29:05', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (58, '2021-02-22 11:29:04.794', '2021-02-22 11:29:04.794', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:29:05', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (59, '2021-02-22 11:29:23.723', '2021-02-22 11:29:23.723', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:29:24', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (60, '2021-02-22 11:30:29.058', '2021-02-22 11:30:29.058', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:30:29', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (61, '2021-02-22 11:30:37.980', '2021-02-22 11:30:37.980', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:30:38', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (62, '2021-02-22 11:30:38.172', '2021-02-22 11:30:38.172', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:30:38', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (63, '2021-02-22 11:30:53.693', '2021-02-22 11:30:53.693', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:30:54', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (64, '2021-02-22 11:30:53.877', '2021-02-22 11:30:53.877', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:30:54', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (65, '2021-02-22 11:34:41.701', '2021-02-22 11:34:41.701', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:34:42', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (66, '2021-02-22 11:34:41.910', '2021-02-22 11:34:41.910', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:34:42', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (67, '2021-02-22 11:46:25.358', '2021-02-22 11:46:25.358', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:46:25', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (68, '2021-02-22 11:46:25.518', '2021-02-22 11:46:25.518', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:46:26', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (69, '2021-02-22 11:48:17.974', '2021-02-22 11:48:17.974', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 11:48:18', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (70, '2021-02-22 17:26:28.490', '2021-02-22 17:26:28.490', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 17:26:28', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (71, '2021-02-22 17:26:40.099', '2021-02-22 17:26:40.099', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 17:26:40', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (72, '2021-02-22 17:26:54.572', '2021-02-22 17:26:54.572', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 17:26:55', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (73, '2021-02-22 17:27:27.146', '2021-02-22 17:27:27.146', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 17:27:27', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (74, '2021-02-22 18:04:18.243', '2021-02-22 18:04:18.243', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 18:04:18', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (75, '2021-02-22 18:04:33.138', '2021-02-22 18:04:33.138', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-22 18:04:33', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');
INSERT INTO `login_log` VALUES (76, '2021-02-23 11:26:24.863', '2021-02-23 11:26:24.863', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-23 11:26:25', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '退出登录');
INSERT INTO `login_log` VALUES (77, '2021-02-23 11:26:31.557', '2021-02-23 11:26:31.557', NULL, '', '', 'test', '1', '0', '::1', '未知位置', 'Chrome87.0.4280.88', 'Windows 10', 'Windows', '2021-02-23 11:26:32', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36', '登录成功');

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
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of operation_log
-- ----------------------------
INSERT INTO `operation_log` VALUES (1, '2021-02-05 17:42:07.879', '2021-02-05 17:42:07.879', NULL, '1', '', '', '11', '', 'POST', 'POST', '', 'test', '/auth/logout', '::1', '未知位置', '', '200', '1', '2021-02-05 17:42:08', '', '', '222.4006ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (2, '2021-02-05 17:42:22.703', '2021-02-05 17:42:22.703', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 17:42:23', '', '', '171.5892ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (3, '2021-02-05 17:43:07.085', '2021-02-05 17:43:07.085', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 17:43:07', '', '', '117.6581ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (4, '2021-02-05 17:45:45.576', '2021-02-05 17:45:45.576', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 17:45:46', '', '', '161.0871ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (5, '2021-02-05 17:48:50.303', '2021-02-05 17:48:50.303', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 17:48:50', '', '', '171.8165ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (6, '2021-02-05 17:48:56.718', '2021-02-05 17:48:56.718', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 17:48:57', '', '', '106.2095ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (7, '2021-02-05 17:49:33.596', '2021-02-05 17:49:33.596', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 17:49:34', '', '', '77.7915ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (8, '2021-02-05 17:51:46.663', '2021-02-05 17:51:46.663', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 17:51:47', '', '', '171.527ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (9, '2021-02-05 17:53:20.331', '2021-02-05 17:53:20.331', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 17:53:20', '', '', '140.6426ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (10, '2021-02-05 17:58:33.629', '2021-02-05 17:58:33.629', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 17:58:34', '', '', '168.8183ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (11, '2021-02-05 17:59:05.994', '2021-02-05 17:59:05.994', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 17:59:06', '', '', '145.6081ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (12, '2021-02-05 18:02:24.262', '2021-02-05 18:02:24.262', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 18:02:24', '', '', '173.5357ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (13, '2021-02-05 18:04:50.587', '2021-02-05 18:04:50.587', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 18:04:51', '', '', '193.8906ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (14, '2021-02-05 18:05:19.075', '2021-02-05 18:05:19.075', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 18:05:19', '', '', '155.3793ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (15, '2021-02-05 18:05:33.926', '2021-02-05 18:05:33.926', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 18:05:34', '', '', '167.0711ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (16, '2021-02-05 18:05:46.833', '2021-02-05 18:05:46.833', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 18:05:47', '', '', '184.5379ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (17, '2021-02-05 18:08:56.612', '2021-02-05 18:08:56.612', NULL, '', '', '用户登录', '10', '', 'POST', 'POST', '', '-', '/auth/login', '::1', '未知位置', '', '200', '1', '2021-02-05 18:08:57', '', '', '181.6648ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (18, '2021-02-05 22:20:43.319', '2021-02-05 22:20:43.319', NULL, '1', '', '', '2', '', 'PUT', 'PUT', '', 'test', '/permission', '::1', '未知位置', '', '200', '1', '2021-02-05 22:20:43', '', '', '110.8047ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (19, '2021-02-05 22:41:05.970', '2021-02-05 22:41:05.970', NULL, '1', '', '', '1', '', 'POST', 'POST', '', 'test', '/permission', '::1', '未知位置', '', '200', '1', '2021-02-05 22:41:06', '', '', '14.9604ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (20, '2021-02-05 22:43:38.381', '2021-02-05 22:43:38.381', NULL, '1', '', '', '1', '', 'POST', 'POST', '', 'test', '/permission', '::1', '未知位置', '', '200', '1', '2021-02-05 22:43:38', '', '', '56.1352ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (21, '2021-02-05 22:46:52.961', '2021-02-05 22:46:52.961', NULL, '1', '', '', '1', '', 'POST', 'POST', '', 'test', '/permission', '::1', '未知位置', '', '200', '1', '2021-02-05 22:46:53', '', '', '103.3124ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (22, '2021-02-05 23:01:37.214', '2021-02-05 23:01:37.214', NULL, '1', '', '', '1', '', 'POST', 'POST', '', 'test', '/permission', '::1', '未知位置', '', '200', '1', '2021-02-05 23:01:37', '', '', '28.554ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (23, '2021-02-05 23:17:42.828', '2021-02-05 23:17:42.828', NULL, '1', '', '', '2', '', 'PUT', 'PUT', '', 'test', '/permission', '::1', '未知位置', '', '200', '1', '2021-02-05 23:17:43', '', '', '101.2383ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');
INSERT INTO `operation_log` VALUES (24, '2021-02-05 23:25:56.248', '2021-02-05 23:25:56.248', NULL, '1', '', '', '2', '', 'PUT', 'PUT', '', 'test', '/permission', '::1', '未知位置', '', '200', '1', '2021-02-05 23:25:56', '', '', '69.8138ms', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36');

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
) ENGINE = InnoDB AUTO_INCREMENT = 115 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (1, '系统管理', 'system', NULL, NULL, 0, 0, 1, '/system/permission', NULL, 0, 'PageView', 'desktop', '/system', 0, NULL, '1', NULL, '2021-01-26 11:36:54.892', NULL);
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
INSERT INTO `permission` VALUES (67, '新增', '', 'add', '', 0, 0, 0, '', '新增', 3, '', '', '', 0, '', '', '2021-02-06 12:22:00.311', '2021-02-06 12:22:00.311', NULL);
INSERT INTO `permission` VALUES (68, '删除', '', 'del', '', 0, 0, 0, '', '删除', 3, '', '', '', 0, '', '', '2021-02-06 12:22:00.311', '2021-02-06 12:22:00.311', NULL);
INSERT INTO `permission` VALUES (69, '修改', '', 'edit', '', 0, 0, 0, '', '修改', 3, '', '', '', 0, '', '', '2021-02-06 12:22:00.311', '2021-02-06 12:22:00.311', NULL);
INSERT INTO `permission` VALUES (70, '查询', '', 'query', '', 0, 0, 0, '', '查询', 3, '', '', '', 0, '', '', '2021-02-06 12:22:00.311', '2021-02-06 12:22:00.311', NULL);
INSERT INTO `permission` VALUES (71, '详情', '', 'get', '', 0, 0, 0, '', '详情', 3, '', '', '', 0, '', '', '2021-02-06 12:22:00.311', '2021-02-06 12:22:00.311', NULL);
INSERT INTO `permission` VALUES (72, '新增', '', 'add', '', 0, 0, 0, '', '新增', 4, '', '', '', 0, '', '', '2021-02-06 12:22:09.425', '2021-02-06 12:22:09.425', NULL);
INSERT INTO `permission` VALUES (73, '删除', '', 'del', '', 0, 0, 0, '', '删除', 4, '', '', '', 0, '', '', '2021-02-06 12:22:09.425', '2021-02-06 12:22:09.425', NULL);
INSERT INTO `permission` VALUES (74, '修改', '', 'edit', '', 0, 0, 0, '', '修改', 4, '', '', '', 0, '', '', '2021-02-06 12:22:09.425', '2021-02-06 12:22:09.425', NULL);
INSERT INTO `permission` VALUES (75, '查询', '', 'query', '', 0, 0, 0, '', '查询', 4, '', '', '', 0, '', '', '2021-02-06 12:22:09.425', '2021-02-06 12:22:09.425', NULL);
INSERT INTO `permission` VALUES (76, '详情', '', 'get', '', 0, 0, 0, '', '详情', 4, '', '', '', 0, '', '', '2021-02-06 12:22:09.425', '2021-02-06 12:22:09.425', NULL);
INSERT INTO `permission` VALUES (77, '新增', '', 'add', '', 0, 0, 0, '', '新增', 2, '', '', '', 0, '', '', '2021-02-07 15:53:34.176', '2021-02-07 15:53:34.176', NULL);
INSERT INTO `permission` VALUES (78, '删除', '', 'del', '', 0, 0, 0, '', '删除', 2, '', '', '', 0, '', '', '2021-02-07 15:53:34.176', '2021-02-07 15:53:34.176', NULL);
INSERT INTO `permission` VALUES (79, '修改', '', 'edit', '', 0, 0, 0, '', '修改', 2, '', '', '', 0, '', '', '2021-02-07 15:53:34.176', '2021-02-07 15:53:34.176', NULL);
INSERT INTO `permission` VALUES (80, '查询', '', 'query', '', 0, 0, 0, '', '查询', 2, '', '', '', 0, '', '', '2021-02-07 15:53:34.176', '2021-02-07 15:53:34.176', NULL);
INSERT INTO `permission` VALUES (81, '详情', '', 'get', '', 0, 0, 0, '', '详情', 2, '', '', '', 0, '', '', '2021-02-07 15:53:34.176', '2021-02-07 15:53:34.176', NULL);
INSERT INTO `permission` VALUES (93, '部门管理', 'department', '', '', 0, 0, 1, '', '', 1, 'Department', 'cluster', '/system/department', 0, '1', '', '2021-02-07 16:20:15.135', '2021-02-07 16:20:15.135', NULL);
INSERT INTO `permission` VALUES (94, '新增', '', 'add', '', 0, 0, 0, '', '新增', 93, '', '', '', 0, '', '', '2021-02-07 16:20:15.177', '2021-02-07 16:20:15.177', NULL);
INSERT INTO `permission` VALUES (95, '删除', '', 'del', '', 0, 0, 0, '', '删除', 93, '', '', '', 0, '', '', '2021-02-07 16:20:15.177', '2021-02-07 16:20:15.177', NULL);
INSERT INTO `permission` VALUES (96, '修改', '', 'edit', '', 0, 0, 0, '', '修改', 93, '', '', '', 0, '', '', '2021-02-07 16:20:15.177', '2021-02-07 16:20:15.177', NULL);
INSERT INTO `permission` VALUES (97, '详情', '', 'get', '', 0, 0, 0, '', '详情', 93, '', '', '', 0, '', '', '2021-02-07 16:20:15.177', '2021-02-07 16:20:15.177', NULL);
INSERT INTO `permission` VALUES (98, '查询', '', 'query', '', 0, 0, 0, '', '查询', 93, '', '', '', 0, '', '', '2021-02-07 16:20:15.177', '2021-02-07 16:20:15.177', NULL);
INSERT INTO `permission` VALUES (99, '字典管理', 'dict', '', '', 0, 0, 1, '/system/dicttype', '', 1, 'RouteView', 'build', '', 0, '1', '', '2021-02-07 18:33:17.511', '2021-02-07 18:33:17.511', NULL);
INSERT INTO `permission` VALUES (100, '字典类型', 'dicttype', '', '', 0, 0, 1, '', '', 99, 'DictType', '', '/system/dicttype', 0, '1', '', '2021-02-07 18:34:21.053', '2021-02-07 18:34:21.053', NULL);
INSERT INTO `permission` VALUES (101, '新增', '', 'add', '', 0, 0, 0, '', '新增', 100, '', '', '', 0, '', '', '2021-02-07 18:34:21.150', '2021-02-07 18:34:21.150', NULL);
INSERT INTO `permission` VALUES (102, '删除', '', 'del', '', 0, 0, 0, '', '删除', 100, '', '', '', 0, '', '', '2021-02-07 18:34:21.150', '2021-02-07 18:34:21.150', NULL);
INSERT INTO `permission` VALUES (103, '修改', '', 'edit', '', 0, 0, 0, '', '修改', 100, '', '', '', 0, '', '', '2021-02-07 18:34:21.150', '2021-02-07 18:34:21.150', NULL);
INSERT INTO `permission` VALUES (104, '查询', '', 'query', '', 0, 0, 0, '', '查询', 100, '', '', '', 0, '', '', '2021-02-07 18:34:21.150', '2021-02-07 18:34:21.150', NULL);
INSERT INTO `permission` VALUES (105, '详情', '', 'get', '', 0, 0, 0, '', '详情', 100, '', '', '', 0, '', '', '2021-02-07 18:34:21.150', '2021-02-07 18:34:21.150', NULL);
INSERT INTO `permission` VALUES (106, '字典键值', 'dictdata', '', '', 0, 0, 1, '', '', 99, 'DictData', '', '/system/dictdata', 0, '1', '', '2021-02-07 19:15:22.429', '2021-02-07 19:15:22.429', NULL);
INSERT INTO `permission` VALUES (107, '新增', '', 'add', '', 0, 0, 0, '', '新增', 106, '', '', '', 0, '', '', '2021-02-07 19:15:22.529', '2021-02-07 19:15:22.529', NULL);
INSERT INTO `permission` VALUES (108, '删除', '', 'del', '', 0, 0, 0, '', '删除', 106, '', '', '', 0, '', '', '2021-02-07 19:15:22.529', '2021-02-07 19:15:22.529', NULL);
INSERT INTO `permission` VALUES (109, '修改', '', 'edit', '', 0, 0, 0, '', '修改', 106, '', '', '', 0, '', '', '2021-02-07 19:15:22.529', '2021-02-07 19:15:22.529', NULL);
INSERT INTO `permission` VALUES (110, '查询', '', 'query', '', 0, 0, 0, '', '查询', 106, '', '', '', 0, '', '', '2021-02-07 19:15:22.529', '2021-02-07 19:15:22.529', NULL);
INSERT INTO `permission` VALUES (111, '详情', '', 'get', '', 0, 0, 0, '', '详情', 106, '', '', '', 0, '', '', '2021-02-07 19:15:22.529', '2021-02-07 19:15:22.529', NULL);
INSERT INTO `permission` VALUES (112, '日志管理', 'log', '', '', 0, 0, 1, '/system/loginlog', '', 1, 'RouteView', 'small-dash', '/system/log', 0, '1', '1', '2021-02-07 19:47:32.118', '2021-02-07 19:54:59.666', NULL);
INSERT INTO `permission` VALUES (113, '登录日志', 'loginlog', '', '', 0, 0, 1, '', '', 112, 'LoginLog', '', '/system/loginlog', 0, '1', '', '2021-02-07 19:48:34.767', '2021-02-07 19:48:34.767', NULL);
INSERT INTO `permission` VALUES (114, '操作日志', 'operationlog', '', '', 0, 0, 1, '', '', 112, 'OperationLog', '', '/system/operationlog', 0, '1', '', '2021-02-07 19:49:09.967', '2021-02-07 19:49:09.967', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 'admin', '管理员', '拥有所有权限', '1', 1, NULL, '1', NULL, '2021-02-07 19:49:18.445', NULL);
INSERT INTO `role` VALUES (2, 'user', '用户', '部分权限', '1', 1, NULL, NULL, NULL, '2021-01-28 14:09:57.446', NULL);
INSERT INTO `role` VALUES (10, 'test', 'test', '', '1', 0, '', '1', '2021-01-29 13:37:32.145', '2021-02-22 17:26:49.256', NULL);

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
INSERT INTO `role_permission` VALUES (2, 1, NULL, NULL, NULL);
INSERT INTO `role_permission` VALUES (2, 2, NULL, NULL, NULL);
INSERT INTO `role_permission` VALUES (2, 3, NULL, NULL, NULL);
INSERT INTO `role_permission` VALUES (2, 5, NULL, NULL, NULL);
INSERT INTO `role_permission` VALUES (2, 6, NULL, NULL, NULL);
INSERT INTO `role_permission` VALUES (2, 7, NULL, NULL, NULL);
INSERT INTO `role_permission` VALUES (2, 11, NULL, NULL, NULL);
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
INSERT INTO `role_permission` VALUES (10, 1, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 2, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 3, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 4, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 10, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 67, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 68, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 69, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 70, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 71, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 72, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 73, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 74, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 75, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 76, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 77, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 78, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 79, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 80, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 81, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 93, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 94, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 95, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 96, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 97, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 98, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 99, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 100, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 101, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 102, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 103, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 104, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 105, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 106, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 107, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 108, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 109, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 110, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 111, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 112, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 113, 'test', '', '');
INSERT INTO `role_permission` VALUES (10, 114, 'test', '', '');

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
  `login_times` int(11) NULL DEFAULT NULL,
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updated_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'test', 'test', 'test', 'test', NULL, 'e10adc3949ba59abbe56e057f20f883e', NULL, NULL, NULL, '1', NULL, '1', NULL, 'test', NULL, '2021-02-02 18:24:18.776', NULL);
INSERT INTO `user` VALUES (2, 'user', NULL, NULL, NULL, NULL, 'e10adc3949ba59abbe56e057f20f883e', NULL, NULL, NULL, '0', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `user` VALUES (3, '测试', '测试啊', '', '测试的昵称', '', 'e10adc3949ba59abbe56e057f20f883e', '122222', '', '', '1', 0, '1', 'test', '', '2021-02-03 13:24:02.105', '2021-02-03 13:24:02.105', NULL);
INSERT INTO `user` VALUES (4, '', '', '', '', '', 'd41d8cd98f00b204e9800998ecf8427e', '', '', '', '1', 0, '0', 'test', '1', '2021-02-03 14:21:56.046', '2021-02-07 16:54:25.095', NULL);
INSERT INTO `user` VALUES (5, 'fasdf', '', '', '', '', '6a204bd89f3c8348afd5c77c717a097a', '', '', '', '1', 0, '1', 'test', '', '2021-02-03 14:33:09.394', '2021-02-03 14:33:09.394', '2021-02-03 16:12:56.660');

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
INSERT INTO `user_role` VALUES (1, 2, 'user', NULL, NULL);
INSERT INTO `user_role` VALUES (2, 1, 'admin', NULL, NULL);
INSERT INTO `user_role` VALUES (3, 1, 'admin', '', '');
INSERT INTO `user_role` VALUES (3, 2, 'user', '', '');
INSERT INTO `user_role` VALUES (3, 10, 'test', '', '');
INSERT INTO `user_role` VALUES (1, 10, 'test', NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
