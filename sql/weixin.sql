/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50620
Source Host           : localhost:3306
Source Database       : weixin

Target Server Type    : MYSQL
Target Server Version : 50620
File Encoding         : 65001

Date: 2016-04-01 19:18:53
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for menu_type
-- ----------------------------
DROP TABLE IF EXISTS `menu_type`;
CREATE TABLE `menu_type` (
  `id` varchar(32) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `value` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu_type
-- ----------------------------
INSERT INTO `menu_type` VALUES ('510B3EA4F7F811E5BEC23010B3A0F15C', '点击推事件', 'click');
INSERT INTO `menu_type` VALUES ('57B35D33F7F811E5BEC23010B3A0F15C', '跳转URL', 'view');
INSERT INTO `menu_type` VALUES ('5F23D9F3F7F811E5BEC23010B3A0F15C', '扫码推事件', 'scancode_push');

-- ----------------------------
-- Table structure for pubnum
-- ----------------------------
DROP TABLE IF EXISTS `pubnum`;
CREATE TABLE `pubnum` (
  `id` varchar(32) NOT NULL COMMENT 'id',
  `name` varchar(50) DEFAULT NULL COMMENT '订阅/公众号名称',
  `appid` varchar(18) DEFAULT NULL COMMENT '第三方用户唯一凭证',
  `appsecret` varchar(32) DEFAULT NULL COMMENT '第三方用户唯一凭证密钥',
  `userid` varchar(32) DEFAULT NULL,
  `created` int(11) DEFAULT NULL,
  `token` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pubnum
-- ----------------------------
INSERT INTO `pubnum` VALUES ('1F2F0533F6ED11E58F303010B3A0F15C', 'gh_4c5d3e510995', 'wx6c8a3beb5709e3d1', 'd4624c36b6795d1d99dcf0547af5443d', '423F7D111E5A', '1446542802', 'a07fbd22a3bd11e5ada43010b3a0f');
INSERT INTO `pubnum` VALUES ('62754E6DF7FA11E5BEC23010B3A0F15C', '内涵段子', '', '', '423F7D111E5B', '1459509047', '');

-- ----------------------------
-- Table structure for token
-- ----------------------------
DROP TABLE IF EXISTS `token`;
CREATE TABLE `token` (
  `id` varchar(32) NOT NULL,
  `token` varchar(512) DEFAULT NULL,
  `pubnumid` varchar(32) DEFAULT NULL,
  `created` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of token
-- ----------------------------
INSERT INTO `token` VALUES ('0C59F1ECF7B011E59B003010B3A0F15C', 'Tg7hhyE4Wuijy521qTAgzEZ6MadauHizwCj8cvbu6h2sxgL0u2KM7j6OieoXAXdyw_doPueT-fCHV4R5h98Iq8Hvsl3s6Om2ECZcQesjY0g7f5hSZB7jmcctJ344kOxYXTUeADATPQ', '1F2F0533F6ED11E58F303010B3A0F15C', '1459477119');
INSERT INTO `token` VALUES ('204F5549F7E411E598703010B3A0F15C', 'oOh0vxjY5ZaaafEtY1P4WoPQZ_yQLR6uv1jt4EVuF7g3QUhWEu8Iu-hrurxP6nFS9astGOuEUJNw8yKem3CWbie7EvSvgl96q1xnWcFQVzwiaNT82uAE___T2txbMVBiIOOjADANTT', '1F2F0533F6ED11E58F303010B3A0F15C', '1459499487');
INSERT INTO `token` VALUES ('52F01223F7D111E5BF573010B3A0F15C', 'MuKuJpajmommH2r8B2Xs1cTfDNCmNQMgp7pKX01TSlH0--GhB5Txdix4V2pkqk0lzfDaTBt5vSBqbCE4qonnfwKd2mJAMR9PsJlS0_opxeVJYKZgpFhmgdVZW2AXP6eOYXObAJARQH', '1F2F0533F6ED11E58F303010B3A0F15C', '1459491411');
INSERT INTO `token` VALUES ('8C14251EF7F811E5BEC23010B3A0F15C', 'L5BjaStYEnmY0MiJCnfc8U69zNr1Z8Ga5QBx4b3oE1EJ1CawlQKkQQzz4Pc-5k_7KUf3FxsqVehtX6MjQmgwTu9V5kjqjkx44M5mF6AsnPT7Bs8c76McSJwc6Z3_Mhu5HSRbAIAUFE', '1F2F0533F6ED11E58F303010B3A0F15C', '1459508257');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(32) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `nickname` varchar(50) DEFAULT NULL,
  `account` varchar(50) DEFAULT NULL,
  `passwd` varchar(50) DEFAULT NULL,
  `created` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('423F7D111E5A', '张三', '无天', 'admin', '111111', '1446542802');
INSERT INTO `user` VALUES ('423F7D111E5B', '张二', '无地', 'zhanger', '111111', '1446542812');
