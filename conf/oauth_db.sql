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

 Date: 08/02/2021 17:38:57
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
  `app_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT 'key',
  `userinfo` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '用户名',
  `expires` datetime(0) NOT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `access_token`(`access_token`) USING BTREE COMMENT 'token 唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 254 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci COMMENT = '访问令牌' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of access_token_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for oauth2_tbl
-- ----------------------------
DROP TABLE IF EXISTS `oauth2_tbl`;
CREATE TABLE `oauth2_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '应用的唯一标识',
  `app_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '公匙',
  `app_secret` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '私匙',
  `expire_time` datetime(0) NOT NULL COMMENT 'appid超时时间',
  `token_expire_time` int(11) NOT NULL COMMENT 'token过期时间',
  `strict_sign` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否强制验签:0：用户自定义，1：强制',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `app_key`(`app_key`) USING BTREE COMMENT 'app_key 唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci COMMENT = 'oauth2 配置' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of oauth2_tbl
-- ----------------------------
INSERT INTO `oauth2_tbl` VALUES (1, '', 'normal', '4EE0A9A43B9B911C067BEE5CC50A9972', '2025-01-01 00:00:00', 1000000, 1);
INSERT INTO `oauth2_tbl` VALUES (2, '', 'hainlp', '4EE0A9A43B9B911C067BEE5CC50A9972', '2025-01-01 00:00:00', 100000, 1);
INSERT INTO `oauth2_tbl` VALUES (3, '', 'wwwthings', '4EE0A9A43B9B911C067BEE5CC50A9972', '2025-01-01 00:00:00', 100000, 1);
INSERT INTO `oauth2_tbl` VALUES (4, '', 'apiserver', '98D93FEB1370D9F35DA3FFFE6083F906\r\n', '2025-01-01 00:00:00', 100000, 1);

-- ----------------------------
-- Table structure for refresh_token_tbl
-- ----------------------------
DROP TABLE IF EXISTS `refresh_token_tbl`;
CREATE TABLE `refresh_token_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `refresh_token` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '刷新令牌',
  `token_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '令牌类型',
  `app_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '访问令牌',
  `userinfo` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '用户名',
  `expires` datetime(0) NOT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `refresh_token`(`refresh_token`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 254 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci COMMENT = '刷新令牌' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of refresh_token_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for user_account_tbl
-- ----------------------------
DROP TABLE IF EXISTS `user_account_tbl`;
CREATE TABLE `user_account_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户账号',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
  `account_type` int(11) NOT NULL DEFAULT 0 COMMENT '帐号类型:0手机号，1邮件',
  `app_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT 'normal' COMMENT '[@fk oauth2_client_tbl.app_key]oauth2_client_tbl表的id(验签id)',
  `user_info_id` int(11) NOT NULL COMMENT '用户附加信息id',
  `reg_time` datetime(0) NULL DEFAULT NULL COMMENT '注册时间',
  `reg_ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '注册ip',
  `describ` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL COMMENT '描述',
  `vaild` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否有效',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UNIQ_5696AD037D3656A4`(`username`) USING BTREE,
  INDEX `user_info_id`(`user_info_id`) USING BTREE,
  INDEX `app_key`(`app_key`) USING BTREE,
  CONSTRAINT `user_account_tbl_ibfk_1` FOREIGN KEY (`app_key`) REFERENCES `oauth2_tbl` (`app_key`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `user_account_tbl_ibfk_2` FOREIGN KEY (`user_info_id`) REFERENCES `user_info_tbl` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_account_tbl
-- ----------------------------
INSERT INTO `user_account_tbl` VALUES (1, 'admin', 'CF08A5402C5CAE463619A9060482B2E2', 0, 'hainlp', 1, '2017-08-16 15:50:59', NULL, 'admin123456!', 1, '', NULL, '', NULL);

-- ----------------------------
-- Table structure for user_info_tbl
-- ----------------------------
DROP TABLE IF EXISTS `user_info_tbl`;
CREATE TABLE `user_info_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `headurl` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_info_tbl
-- ----------------------------
INSERT INTO `user_info_tbl` VALUES (1, '管理员', '');

SET FOREIGN_KEY_CHECKS = 1;
