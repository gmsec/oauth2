/*
 Navicat Premium Data Transfer

 Source Server         : 192.155.1.150
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : 192.155.1.150:3306
 Source Schema         : oauth_db

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 02/07/2021 18:15:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for access_token_tbl
-- ----------------------------
DROP TABLE IF EXISTS `access_token_tbl`;
CREATE TABLE `access_token_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `access_token` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '访问令牌',
  `token_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '令牌类型',
  `app_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '应用的唯一标识',
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '用户名',
  `expires` datetime(0) NOT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `access_token`(`access_token`) USING BTREE COMMENT 'token 唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 1111 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci COMMENT = '访问令牌' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for oauth2_tbl
-- ----------------------------
DROP TABLE IF EXISTS `oauth2_tbl`;
CREATE TABLE `oauth2_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '应用的唯一标识',
  `app_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '公匙',
  `app_secret` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '私匙',
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户账号',
  `expire_time` datetime(0) NOT NULL COMMENT 'appid超时时间',
  `token_expire_time` int(11) NOT NULL COMMENT 'token过期时间',
  `strict_sign` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否强制验签:0：用户自定义，1：强制',
  `oauth_info` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL COMMENT 'oauth信息(一般base64 json串)',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `app_key`(`app_key`) USING BTREE COMMENT 'app_key 唯一',
  UNIQUE INDEX `app_id`(`app_id`) USING BTREE COMMENT 'app_id 唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 219 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci COMMENT = 'oauth2 配置' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for refresh_token_tbl
-- ----------------------------
DROP TABLE IF EXISTS `refresh_token_tbl`;
CREATE TABLE `refresh_token_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `refresh_token` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '刷新令牌',
  `token_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '令牌类型',
  `app_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '应用的唯一标识',
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '用户名',
  `expires` datetime(0) NOT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `refresh_token`(`refresh_token`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1126 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci COMMENT = '刷新令牌' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_account_tbl
-- ----------------------------
DROP TABLE IF EXISTS `user_account_tbl`;
CREATE TABLE `user_account_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户账号',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
  `account_type` int(11) NOT NULL DEFAULT 0 COMMENT '帐号类型:0手机号，1邮件',
  `app_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT 'normal' COMMENT '[@fk oauth2_tbl.app_id]oauth2_tbl表的id(验签id)',
  `user_info_id` int(11) NOT NULL COMMENT '[@fk user_info_tbl.id]用户附加信息id',
  `reg_time` datetime(0) NULL DEFAULT NULL COMMENT '注册时间',
  `reg_ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '注册ip',
  `describ` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL COMMENT '描述',
  `vaild` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否有效',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  INDEX `user_info_id`(`user_info_id`) USING BTREE,
  INDEX `app_key`(`app_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 107 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_info_tbl
-- ----------------------------
DROP TABLE IF EXISTS `user_info_tbl`;
CREATE TABLE `user_info_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户账号',
  `user_info` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL COMMENT '用户信息(一般base64 json串)',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 107 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
