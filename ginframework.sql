/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50720
Source Host           : localhost:3306
Source Database       : ginframework

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2020-01-09 19:44:43
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'professor', '/ginFrameWork/resource1', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'student', '/ginFrameWork/resource1', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/ginFrameWork/*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'student', '/ginFrameWork/resource2', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'professor', '/ginFrameWork/resource2', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'student', '/ginFrameWork/index', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/ginFrameWork/*', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'professor', '/ginFrameWork/index', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'professor', '/ginFrameWork/resource2', 'POST', '', '', '');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` enum('admin','professor','student') DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', 'admin', 'admin', 'admin');
INSERT INTO `users` VALUES ('2', 'professor', 'professor', 'professor');
INSERT INTO `users` VALUES ('3', 'student', 'student', 'student');
