/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50620
Source Host           : localhost:3306
Source Database       : weixin

Target Server Type    : MYSQL
Target Server Version : 50620
File Encoding         : 65001

Date: 2016-04-18 17:54:47
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for media
-- ----------------------------
DROP TABLE IF EXISTS `media`;
CREATE TABLE `media` (
  `id` varchar(32) NOT NULL,
  `created` int(11) DEFAULT NULL,
  `mediaId` varchar(64) DEFAULT NULL COMMENT '媒体素材id',
  `ctype` varchar(50) DEFAULT NULL COMMENT '类型有图片（image）、语音（voice）、视频（video）和缩略图（thumb)',
  `localName` varchar(255) DEFAULT NULL COMMENT '本地名称',
  `url` varchar(255) DEFAULT NULL COMMENT '素材url',
  `saveType` int(1) DEFAULT '0' COMMENT '素材类型：0临时素材，1永久素材',
  `title` varchar(255) DEFAULT NULL,
  `introduction` varchar(2000) DEFAULT NULL,
  `pubnumId` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of media
-- ----------------------------
INSERT INTO `media` VALUES ('00868592014E11E6B7FE3010B3A0F15C', '1460534520', '09ERr6uwQo5BLdZVOMFM5u8v1TYXQ05x_FvI-XY4wLbu_bphr20U_Wd3GhEyNhTc', 'video', 'upload/media/FE5454EE014D11E6B7FE3010B3A0F15C.mp4', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('0B621923020A11E6ADFA3010B3A0F15C', '1460615284', 'aAgFthFu1xA3kWObYcL8RM6qZog_lA5Ul3JddZFShKk', 'video', 'upload/media/066DF74D020A11E6ADFA3010B3A0F15C.mp4', '', '1', '测试标题', '测试描述', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('0BB80723016A11E6BA5D3010B3A0F15C', '1460546565', 'Ng_Ki8Xp2ea3iFBQsJtLkGR0aJI5Uo1TwBfUHWU2KlT9YLJaU50Yx8wHs2tI96dz', 'video', 'upload/media/09AEA293016A11E6BA5D3010B3A0F15C.mp4', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('0F5AF773016A11E6BA5D3010B3A0F15C', '1460546571', 'HmdTtCvdXukvH2xCRsUTElyklD2D_38Q_HaqxF-aCz0R7NoXrR0mFKv5qIi2YXGY', 'image', 'upload/media/0EC212D3016A11E6BA5D3010B3A0F15C.jpg', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('19BFA443016A11E6BA5D3010B3A0F15C', '1460546589', '', 'thumb', 'upload/media/193A6EB3016A11E6BA5D3010B3A0F15C.jpg', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('1F342043016A11E6BA5D3010B3A0F15C', '1460546598', '-p4hamMX3fjWfmd_Z9mUuht9Y7E5G4_qeJjR9o72nso', 'voice', 'upload/media/1E6FBED3016A11E6BA5D3010B3A0F15C.wav', '', '1', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('302CDD05014511E68B623010B3A0F15C', '1460453806', 'PvJrti-v1Q3sYoH19vuadaHKMFTlrVWcwQkT_heQMd4', 'image', 'upload/media/0D0C22B5009211E699413010B3A0F15C.png', 'https://mmbiz.qlogo.cn/mmbiz/rLRxm6HPRCJASMmwiaZC58moULN0Hu8Jd0MYz2XD7kyQlbacPxdumy4aqSLkHkibn60uO83JoVJYfUEUPwGZpQXA/0?wx_fmt=png', '1', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('3068D61E014511E68B623010B3A0F15C', '1460431900', 'ma6WQLXpXgMMRhKwDRloAQLe0_1f59lbn3vmfVFl_1k', 'video', 'upload/media/09AEA293016A11E6BA5D3010B3A0F15C.mp4', '', '1', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('3077074B014511E68B623010B3A0F15C', '1460430867', 'ma6WQLXpXgMMRhKwDRloAdb17ByvUJHbdI_LZBSzoqM', 'video', 'upload/media/09AEA293016A11E6BA5D3010B3A0F15C.mp4', '', '1', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('3087F7AA014511E68B623010B3A0F15C', '1460356306', 'CLNWW7Ft6wXNu9-Qf3Fgatp-oKz4BQK4jt3IyCsxqMU', 'video', 'upload/media/09AEA293016A11E6BA5D3010B3A0F15C.mp4', '', '1', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('309C437F014511E68B623010B3A0F15C', '1460356270', 'CLNWW7Ft6wXNu9-Qf3Fgaia69-9TyxvwfnI3YNiq7Hk', 'video', 'upload/media/09AEA293016A11E6BA5D3010B3A0F15C.mp4', '', '1', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('30A9B157014511E68B623010B3A0F15C', '1460356233', 'CLNWW7Ft6wXNu9-Qf3FgasuN_DU2v6yVSctSzlssXt0', 'video', 'upload/media/09AEA293016A11E6BA5D3010B3A0F15C.mp4', '', '1', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('733C72A701E911E6858C3010B3A0F15C', '1460601285', 'QsGfPXdTO1fTV2I2VTudABBudpxRn4XspYQP-ViXO2k', 'image', 'upload/media/6FB4E27601E911E6858C3010B3A0F15C.jpg', 'https://mmbiz.qlogo.cn/mmbiz/rLRxm6HPRCJh8owxLDAsALFX8705POunZMDBib7ujyZuJvziaib6NnQp2N3ibxv7yXtjX7hh9H5BmjkkGu0TCSjRxA/0?wx_fmt=jpeg', '1', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('770274CD020F11E6B0533010B3A0F15C', '1460617612', 'jGslaa7Jxon8LbKHVlLDcOQTPX0foUPaFjA-R-F_zEhOG4jQddYRIC5MmR1W27RB', 'video', 'upload/media/72ADA44B020F11E6B0533010B3A0F15C.mp4', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('8047CF83016A11E6AEF13010B3A0F15C', '1460546761', '', 'thumb', 'upload/media/7FB44213016A11E6AEF13010B3A0F15C.jpg', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('84C7F5E3016911E6A3393010B3A0F15C', '1460546339', 'elK1JgRypdxQDZJw8kCz7x00u_Dm47h10VOF6YF08w-PjqG7usLSdG5Q8bRz5PEj', 'video', 'upload/media/82297D43016911E6A3393010B3A0F15C.mp4', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('8AB57D03016A11E6AEF13010B3A0F15C', '1460546778', '-p4hamMX3fjWfmd_Z9mUugMnUYXb9Axt6mEfMYq9DVo', 'thumb', 'upload/media/8A196413016A11E6AEF13010B3A0F15C.jpg', 'https://mmbiz.qlogo.cn/mmbiz/rLRxm6HPRCISuryxh5Tav7jZeOZCpOkfQ5hqD1XYkU5eArRK5iamLmNLFCbRxEPPCck2eZhBHMo8Sia4RgFGUiaHg/0?wx_fmt=jpeg', '1', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('9C7CAC73016511E6A1263010B3A0F15C', '1460544661', 'jbzNRuRcEPw0tEw6ufq5kXxodmwDrQBR3oi2HOhRYcc', 'video', 'upload/media/970C0103016511E6A1263010B3A0F15C.mp4', '', '1', '标题', '描述描述描述描述', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('C6E35CA101FF11E692DE3010B3A0F15C', '1460610874', 'hDgk8aDHF82-8RW2t9mJcarkDhGqY-kvWGIeK7fEgqH-BS9X0CxE7moQ2rRf28aT', 'image', 'upload/media/C60B94BB01FF11E692DE3010B3A0F15C.jpg', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('C88EF4B301FF11E692DE3010B3A0F15C', '1460610877', 'jDfzglCwlSvWu1vH4XF9gE3wuIFhgPv7aKLUCj7sjJ1bEkSpx4aJKzGmDlopEtvW', 'image', 'upload/media/C60B94BB01FF11E692DE3010B3A0F15C.jpg', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('C95E09D001FF11E692DE3010B3A0F15C', '1460610878', 'pkJZ7hbVzGUmdAT1pS5lOwTnzfqoRPRFDA6H3B-OoN4km1fvcT7YNJI1FVskkX7q', 'image', 'upload/media/C60B94BB01FF11E692DE3010B3A0F15C.jpg', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('CA27526701FF11E692DE3010B3A0F151', '1460610879', '9vC_UOIn3zrHiKWUA3gJVrl6FqmX06RKwShgRchITdRhbepdeypxLqMglsiK6UZ1', 'image', 'upload/media/C60B94BB01FF11E692DE3010B3A0F15C.jpg', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');
INSERT INTO `media` VALUES ('CA27526701FF11E692DE3010B3A0F15C', '1460610879', '9vC_UOIn3zrHiKWUA3gJVrl6FqmX06RKwShgRchITdRhbepdeypxLqMglsiK6UZN', 'image', 'upload/media/C60B94BB01FF11E692DE3010B3A0F15C.jpg', '', '0', '', '', '1F2F0533F6ED11E58F303010B3A0F15C');

-- ----------------------------
-- Table structure for media_news
-- ----------------------------
DROP TABLE IF EXISTS `media_news`;
CREATE TABLE `media_news` (
  `id` varchar(32) NOT NULL,
  `title` varchar(50) DEFAULT NULL COMMENT '标题',
  `thumb_media_id` varchar(64) DEFAULT NULL COMMENT '图文消息的封面图片素材id',
  `thumb_url` varchar(255) DEFAULT NULL COMMENT '图文消息的封面图片的地址',
  `show_cover_pic` int(1) DEFAULT NULL COMMENT '	是否显示封面，0为false，即不显示，1为true，即显示',
  `author` varchar(50) DEFAULT NULL,
  `digest` varchar(500) DEFAULT NULL,
  `content` text COMMENT '图文消息的具体内容，支持HTML标签，必须少于2万字符',
  `url` varchar(255) DEFAULT NULL COMMENT '图文页的URL',
  `content_source_url` varchar(255) DEFAULT NULL COMMENT '图文消息的原文地址，即点击“阅读原文”后的URL',
  `mediaId` varchar(64) DEFAULT NULL COMMENT '所属的图文素材ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of media_news
-- ----------------------------

-- ----------------------------
-- Table structure for media_news_temp
-- ----------------------------
DROP TABLE IF EXISTS `media_news_temp`;
CREATE TABLE `media_news_temp` (
  `id` varchar(32) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `pid` varchar(32) DEFAULT NULL,
  `content` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of media_news_temp
-- ----------------------------
INSERT INTO `media_news_temp` VALUES ('1', '标题', 'root', null);
INSERT INTO `media_news_temp` VALUES ('10', '正文5', '3', '<section class=\"wx96Diy\" data-source=\"bj.96weixin.com\">\r\n			<section><section class=\"96wx-bdc\" style=\"margin: 0px; border: 1px solid rgb(0, 187, 236); text-align: center;\"><section class=\"96wx-bdc\" style=\"margin: 1em auto; border-radius: 6em; border: 0.1em solid rgb(0, 187, 236); width: 12em; height: 12em;\"><section class=\"96wx-bgc\" style=\"margin: 0.4em auto; border-radius: 5.5em; width: 11em; height: 11em; text-align: center; display: table; max-height: 11em; background-color: rgb(0, 187, 236);\"><section style=\"margin: 1em; padding: 1em; color: white; line-height: 1.2em; font-size: 1.5em; vertical-align: middle; display: table-cell; max-height: 11em;\">请输入标题</section></section></section><section class=\"96wx-bgc\" style=\"margin: 1em 1em 2em; padding: 0.5em 1em; border-radius: 1em; color: white; line-height: 1.5em; font-size: 1em; display: inline-block; white-space: nowrap; max-width: 100%; background-color: rgb(0, 187, 236);\">副标题</section></section><section class=\"96wx-bdc\" style=\"border-width: 0px 1px 1px; border-style: none solid solid; border-color: currentColor rgb(0, 187, 236) rgb(0, 187, 236); margin: 0px; padding: 1em; text-align: center; color: rgb(0, 0, 0); line-height: 1.4em; font-size: 1em;\"><p>请输入内容请输入内容<br>请输入内容请输入内容</p></section></section>\r\n		</section>');
INSERT INTO `media_news_temp` VALUES ('11', '正文6', '3', '<section class=\"wx96Diy\" data-source=\"bj.96weixin.com\">\r\n			<section class=\"96wx-bdtc\"><section data-id=\"46\" class=\"96wx-bdtc\" style=\"border: 0px none; padding: 0px; box-sizing: border-box; margin: 0px; font-size: 14px; font-family: 微软雅黑;\"><section style=\"box-sizing: border-box; padding: 0px; margin: 0px;\"><section style=\"display: inline-block; line-height: 20px; vertical-align: top; margin-top: 0.5em; text-align: center; color: inherit; box-sizing: border-box; padding: 0px;\"><img class=\"96wx-img\" data-wxsrc=\"https://mmbiz.qlogo.cn/mmbiz/iaGswicCbWm6ic2rtZ8OSsoVQKbk32I4libujyBHy1RMPWtwXktiao3QPVSzEF4HaYxawwANnbQYrUtLXSo79PgqRjw/0?wx_fmt=png\" style=\"box-sizing:border-box; color:inherit; margin:0px; padding:0px; vertical-align:top; width:40px\" data-width=\"40px\" src=\"http://cdn.96weixin.com/sucai/zhengwen/8.png\"><p class=\"96wx-bdtc\" data-brushtype=\"text\" placeholder=\"96编辑器\" style=\"box-sizing: border-box; color: inherit; font-size: 12px; line-height: 30px; margin-top: 0px; margin-bottom: 0px; padding: 0px; white-space: normal;\">96编辑器</p></section><section style=\"display: inline-block; margin-top: 1.2em; vertical-align: top; width: 20px; padding-top: 10px; padding-bottom: 10px; color: inherit; box-sizing: border-box;\" data-width=\"20px\"><section class=\"96wx-bdrc\" style=\"border-right-width: 20px; border-left-width: 0px; border-right-style: solid; border-right-color: rgb(213, 72, 11); border-left-color: rgb(213, 72, 11); display: inline-block; height: 15px; width: 15px; vertical-align: top; float: left; color: inherit; box-sizing: border-box; padding: 0px; margin: 0px; border-bottom-width: 12px !important; border-top-style: solid !important; border-bottom-style: solid !important; border-top-color: transparent !important; border-bottom-color: transparent !important;\" data-width=\"15px\"></section></section><section class=\"96wx-bgc\" style=\"display: inline-block; width: 75%; margin-top: 0.7em;padding: 1em; border-top-left-radius: 1em; border-top-right-radius: 1em; border-bottom-right-radius: 1em; border-bottom-left-radius: 1em; font-size: 1em; background-color: rgb(213, 72, 11); border-color: rgb(213, 72, 11); box-sizing: border-box;\" data-width=\"65%\"><p style=\"color: inherit; border-color: rgb(254, 228, 198); box-sizing: border-box; padding: 0px; margin-top: 0px; margin-bottom: 0px; white-space: normal;\">请输入对话的内容</p></section></section><section style=\"width: 0px; height: 0px; clear: both; box-sizing: border-box; padding: 0px; margin: 0px;\"></section></section></section>\r\n		</section>');
INSERT INTO `media_news_temp` VALUES ('2', '关注', 'root', null);
INSERT INTO `media_news_temp` VALUES ('3', '正文', 'root', null);
INSERT INTO `media_news_temp` VALUES ('4', '标题1', '1', '<section class=\"wx96Diy\" data-source=\"bj.96weixin.com\">\r\n			<section class=\"96wxDiy\" style=\"clear: both; position: relative; width: 100%; margin: 0px auto; overflow: hidden;\"><section style=\"border:none;margin:5px 0px;padding:0px 5px;\"><section style=\"box-sizing: border-box; color: inherit; border-color: rgb(27, 156, 226);\"><span class=\"96wx-bgc\" style=\"border-top-left-radius: 0px; border-top-right-radius: 0em; border-bottom-right-radius: 0em; border-bottom-left-radius: 0px; box-sizing: border-box; display: block; font-size: 16px; text-align: center; padding: 0.5em; border-color: rgb(42, 102, 134); background-color: rgb(217, 10, 98);\"><p style=\"color:rgb(255, 255, 255)\">输入内容</p></span></section><section class=\"96wx-bdtc 96wx-bdrc\" style=\"width: 0px; border-right-width: 4px; border-right-style: solid; border-right-color: rgb(217, 10, 98); border-top-width: 4px; border-top-style: solid; border-top-color: rgb(217, 10, 98); box-sizing: border-box; color: inherit; height: 4px; border-left-width: 4px !important; border-left-style: solid !important; border-left-color: transparent !important; border-bottom-width: 4px !important; border-bottom-style: solid !important; border-bottom-color: transparent !important;\"></section><section class=\"96wx-bdtc 96wx-bdlc\" style=\"float: right; height: 4px; border-right-width: 4px; border-right-style: solid; border-top-width: 4px; border-top-style: solid; border-top-color: rgb(217, 10, 98); box-sizing: border-box; color: inherit; width: 0px; margin-top: -8px; border-left-color: rgb(217, 10, 98); border-left-width: 4px !important; border-left-style: solid !important; border-right-color: transparent !important; border-bottom-width: 4px !important; border-bottom-style: solid !important; border-bottom-color: transparent !important;\"></section></section></section>\r\n		</section>');
INSERT INTO `media_news_temp` VALUES ('5', '正文1', '3', '<section class=\"wx96Diy\" data-source=\"bj.96weixin.com\">\r\n			<section class=\"96wxDiy\" style=\"clear: both; position: relative; width: 100%; margin: 0px auto; overflow: hidden;\"><section style=\"position:static;box-sizing: border-box;margin: 0.5em 0;\"><section style=\"color: inherit;/* text-align: center; */\"><section style=\"border: 0px none;\"><section style=\"margin-top: 10px; margin-left: 5px;\"><section style=\"width: 48%; display: inline-block; float:left;box-shadow: rgb(128, 130, 110) 0px 2px 5px;\"><section style=\"padding: 10px; min-height: 140px; text-align: center; background: rgb(146, 208, 80);\"><h3 class=\"96wxDiy\" style=\"margin-bottom: 10px;\">春夏</h3><section style=\"color: inherit;\"><p class=\"96wxDiy\" style=\"line-height: 24px; zoom: 1; margin-bottom: 0px; white-space: pre-wrap; text-align: left;\">春天随着落花走了,夏天披着一身的绿叶儿在暖风里蹦跳着走来了。</p></section></section></section><section style=\" width: 48%; display: inline-block; text-align: center; box-shadow: rgb(128, 130, 110) 0px 2px 5px;\"><section style=\"padding: 10px; min-height: 140px; background: rgb(146, 208, 80);\"><h3 class=\"96wxDiy\" style=\"margin-bottom: 10px;\">盛夏</h3><section style=\"color: inherit;\"><p class=\"96wxDiy\" style=\"line-height: 24px; zoom: 1; margin-bottom: 0px; white-space: pre-wrap; text-align: left;\">盛夏,天热得连蜻蜓都只敢贴着树荫处飞,好像怕阳光伤了自己的翅膀。</p></section></section></section></section></section></section></section></section>\r\n		</section>');
INSERT INTO `media_news_temp` VALUES ('6', '关注1', '2', '<section class=\"wx96Diy\" data-source=\"bj.96weixin.com\">\r\n										<p><img width=\"100%\" class=\"96wx-img\" data-wxsrc=\"https://mmbiz.qlogo.cn/mmbiz/iaGswicCbWm6icjxgD1EaN9UjrBrGnGvVpXcDajlUzcjcuAsJJxVvEAdGsTctwQcw8Fd46JQIzhZQUakgMcnCdM7Q/0\" src=\"http://cdn.96weixin.com/sucai/guanzhu/4.jpg\"></p></section>');
INSERT INTO `media_news_temp` VALUES ('7', '正文2', '3', '<section class=\"wx96Diy\" data-source=\"bj.96weixin.com\">\r\n			<section id=\"96weixinzw081639\" class=\"96wx-bdtc\" style=\"border: 0px rgb(0, 187, 236); margin: 1.8em 0px 0.5em;\"><section class=\"96wx-bdc\" style=\"border: 2px solid rgb(0, 187, 236); padding: 10px; margin-left: 1.3em; font-size: 1em; font-family: inherit; font-weight: inherit; text-align: inherit; text-decoration: inherit;\"><section class=\"96wx-bdc 96wx-bgc\" style=\"border: 2px solid rgb(0, 187, 236); padding: 0px 5px; display: inline-block; float: left; vertical-align: top; margin: -1.3em 0px 0px -1.3em; font-size: 1.5em; font-family: inherit; font-weight: inherit; text-align: inherit; text-decoration: inherit; background-color: rgb(0, 187, 236); color: rgb(255, 255, 255);\"><section style=\"color: inherit;\">01</section></section><section style=\"clear: both; color: inherit;\"></section><section style=\"color: inherit;\"><section class=\"96wx-bdtc\" style=\"color: inherit;\">输入文字输入文字</section></section></section></section>\r\n		</section>');
INSERT INTO `media_news_temp` VALUES ('8', '正文3', '3', '<section class=\"wx96Diy\" data-source=\"bj.96weixin.com\">\r\n			<section style=\"padding: 8px 5px; text-align: center; border: 1px solid rgb(0, 187, 236);font-size: 1em; font-family: inherit; font-weight: inherit; text-decoration: inherit; color: rgb(51, 51, 51); box-sizing: border-box;\" class=\"96wx-bdc\"><section style=\"margin-top: -1.2em; padding: 0px; border: none; box-sizing: border-box;\" class=\"96wx-bdtc\"><span style=\"display: inline-block; padding: 4px 8px; color: rgb(255, 255, 255); font-size: 1em; line-height: 1.4;  border-top-left-radius: 32px 8px; border-top-right-radius: 16px 48px; border-bottom-left-radius: 16px 48px; border-bottom-right-radius: 64px 8px; font-family: inherit; font-weight: inherit; text-decoration: inherit; background-color: rgb(0, 187, 236); border-color: rgb(0, 187, 236); box-sizing: border-box;\" class=\"96wx-bgc\"><section class=\"96wx-bdtc\" style=\"box-sizing: border-box;\">请输入标题9</section></span></section><section style=\"padding: 8px; font-size: 1em; line-height: 1.4; font-family: inherit; box-sizing: border-box;\" class=\"96wx-bdtc\"><section class=\"96wx-bdtc\" style=\"box-sizing: border-box; text-align: left;\">可输入内容，标题背景可以更换颜色，素材边框可以更换颜色，背景为白色，无阴影</section></section></section>\r\n		</section>');
INSERT INTO `media_news_temp` VALUES ('9', '正文4', '3', '<section class=\"wx96Diy\" data-source=\"bj.96weixin.com\">\r\n			<section class=\"wx96Diy 96wx-bgpic\" style=\"width: 100%; padding: 15px 15px 15px 15px; border-width: 1px; border-top-style: solid; border-right-style: solid; border-bottom-style: solid; border-color: rgb(204, 204, 204); max-width: 100%; line-height: 1.5; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-left-radius: 5px; border-bottom-right-radius: 5px; background-image: url(http://cdn.96weixin.com/sucai/zhengwen/3.jpg); box-sizing: border-box !important; word-wrap: break-word !important; background-position: 10px 11px;\" data-wxsrc=\"https://mmbiz.qlogo.cn/mmbiz/iaGswicCbWm6ib4sQwRuoty4m9wFibZ7KDaXxcibe4wUadQ255QHm6Xff3H6WicibXckwpazjNUJ5YDUI7erfdv0c0CnQ/0?wx_fmt=jpeg\"><fieldset class=\"wx96Diy\" style=\"margin: 0px; padding: 5px; max-width: 100%; border: 1px solid rgb(204, 204, 204); font-family: 宋体; box-sizing: border-box !important; word-wrap: break-word !important;\">可在这输入内容， &nbsp;- 96微信图文排版,微信图文编辑器,微信公众号编辑器，微信编辑首选。</fieldset></section>\r\n		</section>');

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
INSERT INTO `menu_type` VALUES ('1ACB1374FA5711E58AED6057182B3A68', 'scancode_waitmsg_扫码推事件且弹出“消息接收中”提示框', 'scancode_waitmsg');
INSERT INTO `menu_type` VALUES ('2E485C4AFA5711E58AED6057182B3A68', 'pic_sysphoto_弹出系统拍照发图', 'pic_sysphoto');
INSERT INTO `menu_type` VALUES ('510B3EA4F7F811E5BEC23010B3A0F15C', 'click_点击推事件', 'click');
INSERT INTO `menu_type` VALUES ('57B35D33F7F811E5BEC23010B3A0F15C', 'view_跳转URL', 'view');
INSERT INTO `menu_type` VALUES ('5F23D9F3F7F811E5BEC23010B3A0F15C', 'scancode_push_扫码推事件', 'scancode_push');

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
INSERT INTO `pubnum` VALUES ('2FF57F87FB0011E5AA223010B3A0F15C', '自己的', 'wx7387b8a980dcaaee', '15e338d726360b0cd479247cf4473349', '423F7D111E5A', '1459841392', 'a07fbd22a3bd11e5ada43010b3a0f');
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
INSERT INTO `token` VALUES ('03138A7804C211E6B5103010B3A0F15C', 'GtpPCpMw7Te4Xt_IPse0BrjSuOf68dKL9JpTbw4jqsVUXB_nMX9O4-L-JwkNdl6PidoUJVIrIIgyB4chaJ8hQY_BrW5_wmOGx3q2PlLyQJSYBqwxPAas9HTvurT8kjUrBQVcAJAVKP', '2FF57F87FB0011E5AA223010B3A0F15C', '1460914200');
INSERT INTO `token` VALUES ('05DF8DFE051A11E686CB3010B3A0F15C', 'MT1zj0xXRIRTu-5-AQoJRF0szoB7HP0m2lXPcpNGeauU3vFgw6s4FiMdfYPPULUNO5YdAr6Cj-_kJLblcJffZB2J34HYI759hMQU8c9uqVMxsAUmi-9tWuL6gCV3f_IdRYFaAJAWDK', '1F2F0533F6ED11E58F303010B3A0F15C', '1460952000');
INSERT INTO `token` VALUES ('0630460E051A11E686CB3010B3A0F15C', 'OX_DfPtPih3bvC346uRZ4JXKOX58L1RLQvDjEeyYFF5cszf1VqP5KDkFn7A4PfSp4EArMDZh1ATqkbJy9PQrouEHgb9U7_opbx7aXYrYrS82O8BagC1sHq1V3J0gpgs1ZPVfAGAFVB', '2FF57F87FB0011E5AA223010B3A0F15C', '1460952001');
INSERT INTO `token` VALUES ('071C5261006311E6AC1F3010B3A0F15C', '_m9EAwa_cR53DjFQu7mLgVAEkJd-W07SJsfDOF9arGvxOk-dsNeWOPPiJWW755Qsec6mytx1iKIOFGWhKl5OrszZHeH_R2he-wv6odKFAUYTHUgACAGUH', '1F2F0533F6ED11E58F303010B3A0F15C', '1460433600');
INSERT INTO `token` VALUES ('07564E21006311E6AC1F3010B3A0F15C', '3Dp-ur5QiOGZsfg_UjldJpMp6isyvOBUk7UyNcHWOPDo_rAtckRMttFr68z_hHN_phk24o9Ra0YG5RsdkD7H8_jpT3vNzXphrZfTEaN4mNXsc_i_qcFno7ah2jbCIJG3VGZcABADOH', '2FF57F87FB0011E5AA223010B3A0F15C', '1460433600');
INSERT INTO `token` VALUES ('0C5901C4FAFC11E5BE763010B3A0F15C', 'A56860XyDLrrhZPDnYTnQe9gtb4EcZ8s_lIQZk471nnUXG_QC5Zm2rxiRTnUmLXXQ8WzDn2fSqiUFpibfzEvwHg7CvzFqfttmfU_n7aGBW5thGtpawFEIZ2iPViz_m9GJNHbAGAOSM', '1F2F0533F6ED11E58F303010B3A0F15C', '1459839615');
INSERT INTO `token` VALUES ('0C59F1ECF7B011E59B003010B3A0F15C', 'Tg7hhyE4Wuijy521qTAgzEZ6MadauHizwCj8cvbu6h2sxgL0u2KM7j6OieoXAXdyw_doPueT-fCHV4R5h98Iq8Hvsl3s6Om2ECZcQesjY0g7f5hSZB7jmcctJ344kOxYXTUeADATPQ', '1F2F0533F6ED11E58F303010B3A0F15C', '1459477119');
INSERT INTO `token` VALUES ('0D10C51AFC6B11E5B73B3010B3A0F15C', 'qyoqQcG2ckVEsOQnl4QrcQdbQFMhgl4kX-bLC_EOesgQoj1BQdtNqtKB8yQAIaRyNistdJuEpT9UeB9CCT8z3d63FyR_1fOryia29nzTCsDpOAcO3XIq74fze31TWEF1IJJiAEAPQQ', '1F2F0533F6ED11E58F303010B3A0F15C', '1459997241');
INSERT INTO `token` VALUES ('1024B745FFF611E5934D3010B3A0F15C', 'N-GP0CMo9xLR3FKR8JVy-zvgxskHCGF2SLtdqdgpo7dkTtJMevgZXJ3Z12zVksOhbCokiE5o5MiaZ26Rdogy2VYov90xFNAA8eINeW5H-A2dgYOumRG-SqKAlBHJ3ir2RIRiAGAEVE', '1F2F0533F6ED11E58F303010B3A0F15C', '1460386800');
INSERT INTO `token` VALUES ('108920F7FFF611E5934D3010B3A0F15C', '0w3pNtBNFtqKBPgWceszseMkDzQOxtMxE-Zn9f02gBqAlOd4O18qZhBXsr394O3iDzF75CO5vmyNStWqDSCifO0-tnN532tVYwO7wOt1LGoUJOdADAFUZ', '2FF57F87FB0011E5AA223010B3A0F15C', '1460386800');
INSERT INTO `token` VALUES ('11163D5E050511E6B5103010B3A0F15C', 'KmP899SVnNCwleZI5OqLLGTO1y0yR-_wNm9KUR78rvsI6JDi8gbmVGuiTG4A-jBpO8AIMw6U48R6VJSqkoOxPzeZmYJb_BTF20GVE8GNP3aUn4EuYyjFH7tHwoNJzWVNYYTaAGAMMH', '1F2F0533F6ED11E58F303010B3A0F15C', '1460943000');
INSERT INTO `token` VALUES ('12B3398C004E11E6934D3010B3A0F15C', 'X9WpJQ3FyUHXl-ahOIjJPGFTDV_ylqqEK1FxBtpKb47ff_xVSI-SNwq1dh-4TncEzmcHpnQW5EQd_wUZPCT5Au8zzPk88n2Blod5vEMaZYfIsXWu-QF_NohfaZlUQlaTZATfABAOPI', '1F2F0533F6ED11E58F303010B3A0F15C', '1460424600');
INSERT INTO `token` VALUES ('12DA98AE004E11E6934D3010B3A0F15C', 'mCeQmCpf0FjCVstu8_EX4EyiZ7jmRhWB0RkUC8TNSz6D-m6XQmeCoeG0qXC1gGGmiHHtDm_uDf2t8nNqTLc8cy33UfPoqQim989DUk7OowkZ9f2gJZgZkDjPOR86MPjeBWXjAJAWRS', '2FF57F87FB0011E5AA223010B3A0F15C', '1460424600');
INSERT INTO `token` VALUES ('13F0AA70021B11E6B0533010B3A0F15C', '1SM1B_P8yXqtLtHUB13deXIdCe-iqO06_A0w1TZtyy8x8UrrVMSU1tOuyQhOhHacF_CkE3neDVtpNKSvJMaq_2bemQG-b7WqcSGMvvVzSJYOPHfAJAHIA', '2FF57F87FB0011E5AA223010B3A0F15C', '1460622600');
INSERT INTO `token` VALUES ('1508057D00A611E692FC3010B3A0F15C', 'Q9-uULekRf3Q6u1d6ExoxjuQmOwmPonHmZ6Ubt9MFyEAwpDndkwBKEWCSiF_1YEkz4OrxMPb2limNA7qGNIO-bf3lIkLoKZPYsPhrs8CiHLNa657BPThPWI4C_oZOyVDBVMdAHAXRI', '2FF57F87FB0011E5AA223010B3A0F15C', '1460462399');
INSERT INTO `token` VALUES ('16A84C8001E411E6AF103010B3A0F15C', '82QyMHT-Bj3GP-SHvriv0xjPbS8RFsqd6G1Zxsw84DImzAeXp77csIKGiE0bkZAwGgoMBSlZMt_cjDTXFu6cpwLW05mQmcCzucHTSfLCUa02ay3VEwksBx8VHWJcXj_rSHJhAIAYNP', '1F2F0533F6ED11E58F303010B3A0F15C', '1460598982');
INSERT INTO `token` VALUES ('17BFB090044011E6B5103010B3A0F15C', 'GSpvQCUYZfyTAVevVCIL__EQblJFAGoSapsXBi7j8k33e7QZl_r5O9j8kjODhpN4F-retgS8JezTmI6gFvzwOqM9G4-WqIMVajlq4vgGv7YW617JadvpKxxvv_rD-f_wDJDdAHAYEY', '1F2F0533F6ED11E58F303010B3A0F15C', '1460858400');
INSERT INTO `token` VALUES ('190929D402CB11E6AB2E3010B3A0F15C', 'cAFpdIV5PadomOdxoBMrsCF21Uw9u_fjk3FjRg9JEEPVFBUAFn5tVuFRBkslP_3QsJ3eUOhQsVyeiwq954xAtVH1Teq_hqkbT1-fjX5RjKg26d8Q1oxqR9lYIiOm1dXPMLYeAHADCP', '2FF57F87FB0011E5AA223010B3A0F15C', '1460698200');
INSERT INTO `token` VALUES ('191DACB0FF8911E5BC003010B3A0F15C', 'iddy4gz3aRkLEG1XcvARQfk8W8gYy-8c-p0PPZzMv1iTfJAag9Mq4oHtadN62huRpsrwa7ExXJKl3bCS3EysHZ2RiYNy4H7r8Ue-VY24gjCiToZou1DqELY1X-GlslsBYNSjAIAKZE', '2FF57F87FB0011E5AA223010B3A0F15C', '1460340000');
INSERT INTO `token` VALUES ('1A2CE743049811E6B5103010B3A0F15C', '5VVOT4PpTxtvKQmu9PQunQzmHpJ4FNwcdH0aXdpHPbqoH-gzbVGsX7D91Od31AOtcIxGR2xYMlUgGfxyj1gXIIG80qRktltg9xmyczzsgjJSKRCFQBW9Qm4WFE2wQ2GaXSFaABAEZM', '1F2F0533F6ED11E58F303010B3A0F15C', '1460896200');
INSERT INTO `token` VALUES ('1A5D8574015611E681D43010B3A0F15C', '-iwx5Y671IqOifTEOTQ6yn3K2iS_WHOoI1onsGbgbOChkxWohp9Ly-IK4qbN7DSV3AZM6AHGyl2Rgw2OW0Lg-ehCF13WF74cqSrBdTNb4Pa3sR5kYVU3H892xJ7RpYlWXURaAHAHEB', '2FF57F87FB0011E5AA223010B3A0F15C', '1460538000');
INSERT INTO `token` VALUES ('1CADD16E04F011E6B5103010B3A0F15C', '-h-tvaUNJ266PMFJLgmTG7ZB8RW4x2PE3SSnga-LA_NPWf_8FznDq26a2ZQdsew3wlvjiLhCkEOHaEsaOcb-s6N2QgK12POi76MX8Xuzp0oxSTmh6jH9fJE1K5zWilXhULZeAHACTH', '1F2F0533F6ED11E58F303010B3A0F15C', '1460934000');
INSERT INTO `token` VALUES ('1FA7B1C3020611E6ADFA3010B3A0F15C', 'C87ohOR37y2XLzz0ZJNnvFHA5K-rI3wluOKQ0X82mhe1OqSXyBEQCnVQBlzqovJcKRkT9DQwSZEqeqjtwfCOOFWoIQx32NorLp5v3uyb0TgKq7g-AW1sUke5dih6hxArSSPfACAHHE', '1F2F0533F6ED11E58F303010B3A0F15C', '1460613600');
INSERT INTO `token` VALUES ('2002F379020611E6ADFA3010B3A0F15C', '33d3KDMMoRWl5VxH8NCmXkuLQx0nQG7IYiDkGedvueWa12H-VT-hgfu54tXyRxYqbgHs1QJ8sWYrP2DtjNqdIHbL210Ub7JCn4-lVSwG7pitxjOB1EHBbSDVEV9uneHhEGMdAJAZBH', '2FF57F87FB0011E5AA223010B3A0F15C', '1460613601');
INSERT INTO `token` VALUES ('204F5549F7E411E598703010B3A0F15C', 'oOh0vxjY5ZaaafEtY1P4WoPQZ_yQLR6uv1jt4EVuF7g3QUhWEu8Iu-hrurxP6nFS9astGOuEUJNw8yKem3CWbie7EvSvgl96q1xnWcFQVzwiaNT82uAE___T2txbMVBiIOOjADANTT', '1F2F0533F6ED11E58F303010B3A0F15C', '1459499487');
INSERT INTO `token` VALUES ('20B1840903D311E6B5103010B3A0F15C', 'yTqRdMGh1cMaTN6yYhjnhB9GFEkhW-hgExNcftkQphOL77WtWh-m6woIZ6np66gCI5KjYlUHURcjFNwt1qUwsYa-erftKup0B1_iLmAEXeUkxppL44Z5TE6L788Tt6vDJDDdADAUOO', '2FF57F87FB0011E5AA223010B3A0F15C', '1460811600');
INSERT INTO `token` VALUES ('20F4FC63016511E6A1263010B3A0F15C', 'R21H08bIpy_BC5f8LuhjJxo77b7GSvbRChMVVKYzztm2h9M2xQGSUJ1PCgp14v9-4OmbMz2D23EPjsF-KbfJl1V9taOF0CiffNvKEwkW5ZZRWk3EZ0hm14oPYKo2dv7DQVLjAEAGLB', '1F2F0533F6ED11E58F303010B3A0F15C', '1460544453');
INSERT INTO `token` VALUES ('211DFB1CFD4F11E58D1A3010B3A0F15C', 'JoggzQZKxHFWkXf2C6hSYiTSAZOR5QWQhD6Z9_TFKvwyw64qku4Lk-u1xGv4l_KZQbFyWRu5QVC8vpWazAtWRCBCflk8RsgogbILcs9ZDovqdU0I9MdDJqvgZ6zOqN-3MTBaAFASZH', '2FF57F87FB0011E5AA223010B3A0F15C', '1460095200');
INSERT INTO `token` VALUES ('220F04F5025E11E6AB2E3010B3A0F15C', '5C1OUZGSXLQX1HApIQ-h_4LmD4cY3AtKRUEcRoiAUf8ESsOkQ-VEOUWPbk3ExpKSWa-T66HBbks6HcwF9yMND-yuHU0PKa-U9owhrQddle5ZwYG_ZjaWc2Jw44MLoJG3YRLcABARKE', '2FF57F87FB0011E5AA223010B3A0F15C', '1460651400');
INSERT INTO `token` VALUES ('2312F10E009111E6A8553010B3A0F15C', '3Dp-ur5QiOGZsfg_UjldJvm_M_-UzLIgzuQOWZXwDyQq5uaVVEbNOf_OWUxIYTFV4JprdfguDqCH3wjurKADkWIUpjB0ss70YG2N64D4MLhJ4aQ8imAYoRVKinfydeJfWJUeAAANAL', '2FF57F87FB0011E5AA223010B3A0F15C', '1460453404');
INSERT INTO `token` VALUES ('2322C720042B11E6B5103010B3A0F15C', '-h-tvaUNJ266PMFJLgmTG_rv4DRFS38rT_l3XJ3OjhZZGBgLNtTsd223wdc-mM0mCUozEEdhgHOkOHj73CKn-kP75SV8TmPDGo_RN4j6VopgeqTJhA2SI_f5V-ehJjSRFUNiADAUJC', '1F2F0533F6ED11E58F303010B3A0F15C', '1460849400');
INSERT INTO `token` VALUES ('25D232FD048311E6B5103010B3A0F15C', 'ucT6l_i6PrZda5nSjj1ctKbksX4zOaoWKnqT5aHt2eXnGValL-Cb4deGa_HGDRMVK7Nvi4ysg0gKIwBAHHQEyRoB6ebOXWByws-geYpsaeARbNEkkv9RdP1GJ-dWvP8bZULbABAFAR', '1F2F0533F6ED11E58F303010B3A0F15C', '1460887200');
INSERT INTO `token` VALUES ('285A075404DB11E6B5103010B3A0F15C', 'A56860XyDLrrhZPDnYTnQaoS9BSaimj00BJHxI-3oov-aQlhiAWH3aYYClYTzJClLxRrSWD5zOQoV34hqp8LJh_Ow0-6t-eShRcxZ5J5PUp3nhAdMEuuJvEkXfgI2DzcFMHbAAAOVQ', '1F2F0533F6ED11E58F303010B3A0F15C', '1460925000');
INSERT INTO `token` VALUES ('2DDD174F024911E6AB2E3010B3A0F15C', 'qbbizEJlb7YoD1JSo4SiGWyYY0I87MFSblFEECXl6YZPqvZM4w1P9ajCGApCQ2wM9ytKDJiRfoF2Y3iXWh0TgjMSsf7QjfkdVLcWJC_5hSmIKI721SWxRdRwEOpap-3cOKOdADAVRV', '1F2F0533F6ED11E58F303010B3A0F15C', '1460642400');
INSERT INTO `token` VALUES ('2EB63C80041611E6B5103010B3A0F15C', 'N-GP0CMo9xLR3FKR8JVy-x-JwIGNNIIU-Ct0aSlQVRzDRx-wxfYkAVTuZCKEEoH7TcP2XQv5Mik262QmiUMPtwZZE9r7UQC6sKtugQpkfAwZrsX9XBKf1b5xNvX2ayHbUTLcADAZYQ', '1F2F0533F6ED11E58F303010B3A0F15C', '1460840400');
INSERT INTO `token` VALUES ('3026A9D002A111E6AB2E3010B3A0F15C', 'HEiFk8sWpZxbHw6hbNIfvZwX88Qu5l6WiMsxUU7tzgGgXh4GbT4l_z95obpVWiQXfLci_ezxP4K8lSg49S11HozW6F_QQDNTG73nF9e07QnJo_QRbKjzw4PvLdmQs-YAIHZdADALKC', '1F2F0533F6ED11E58F303010B3A0F15C', '1460680200');
INSERT INTO `token` VALUES ('3162C594046E11E6B5103010B3A0F15C', 'FHBAvU27YkiWwJFEu8og-y8bmXeI13-xqn4L6Q9fCFXX1ACIu6LOzqFHZ6bOfTOfxQfSKYtamuHDyKCnWjG0QdDeQC_yK14-3wVcTuPlqd4M94-_dF9PvBllX-i6nEfRZAAgADAUVV', '2FF57F87FB0011E5AA223010B3A0F15C', '1460878200');
INSERT INTO `token` VALUES ('3532FA67035111E6B5103010B3A0F15C', 'nSHDKA-Ra6hxfuIq95Z-FiHG0eskjQRiJ1A24BTCyPtHupeMJzHbIOjmzxF-1Jusx77U3dUtR2ARm530vlycD0gvygIwp-hTPWmabOHZvGzIjtMeQHH0oYKiiGzMGZehLFBeAHAOIE', '1F2F0533F6ED11E58F303010B3A0F15C', '1460755800');
INSERT INTO `token` VALUES ('35661857035111E6B5103010B3A0F15C', 'D5R0EENPUW54co9Jwjvan9pLuO0bVihOs3OwdQURdZW_rGQg47nnmrgHPKD4CwuHOknT9ktNG_0aDNivdOdNiBtekyvXWaFM-upQ4-V7gydoVfl665MRmdExeommis4uUXMbAGACFU', '2FF57F87FB0011E5AA223010B3A0F15C', '1460755800');
INSERT INTO `token` VALUES ('356CC779FA1E11E590636057182B3A68', 'VaBRE-rERCds50L0Nbzv9cajaXh753sZ4RCIOuzlHXMqlbywQ1LD46hgUrifZvVgrIfvVcJTE9uXDDT3q7qKVehTuna7sg6fgzqQ3n7r_USWZVsaykjmJeu3KMVRJzN-PGBcAIAULG', '1F2F0533F6ED11E58F303010B3A0F15C', '1459744335');
INSERT INTO `token` VALUES ('37DC8D0103A911E6B5103010B3A0F15C', 'wA75FqzPEUCd343j_Yr1Y6v1NZyO98KqOttRt9BfSTeaFj1K6wMryKxNkxZimSysRQPbZlXsFQIR0lQghEq5ABF7PzyRiZBNOv0ntP4_iYARVMfABADHY', '1F2F0533F6ED11E58F303010B3A0F15C', '1460793600');
INSERT INTO `token` VALUES ('3A4B86A0040111E6B5103010B3A0F15C', '7-g-PNyptCTxyIOWh9LmtaMcJGfGOtZ5zbgPYUBzpLTNCLTI-IXzaa5zI99sNSCqYH3m2PYTIs48LtciYU6Stb8KC-tTCGteoJlvqOHDI3ihRwpAdKTYIgoWweMWTxJbTHVeAAAZNB', '1F2F0533F6ED11E58F303010B3A0F15C', '1460831400');
INSERT INTO `token` VALUES ('3D0D1D1A011711E692FC3010B3A0F15C', 'QI8g7QdUymXVhNKh4qvdrH6CXujDm_N1Duf5i2HsSGGJZLpeersITTJ4U8eTITeTp5cU8XvU729d_ngiH9By5XgnmdqMuXqxSxKROiKhxipWsLY4pcMOWTDid4MOq58qRAEdADABKO', '1F2F0533F6ED11E58F303010B3A0F15C', '1460511000');
INSERT INTO `token` VALUES ('3D39AB5A011711E692FC3010B3A0F15C', 'FHBAvU27YkiWwJFEu8og-7VFWwx1b1qwPhbNri6xujiG5POf95qu85CLxmSO47wX10Hs7rcyKyMrWSeh9vvbwWtF7feMw7wzdDK--PomgztRa5gm3LV1ootX6KwD16fXMDXiAAAFWQ', '2FF57F87FB0011E5AA223010B3A0F15C', '1460511000');
INSERT INTO `token` VALUES ('3DD9BC43015311E681D43010B3A0F15C', 'Z-Ugkg6Yn1ynZoND3Hq2pzc-rDoC-Xi-I95Sc490p_SHHlZ-yEu4RNMb2EWoTodOlPr7JlQQhgffdK8Wt1bkTqkfzQYe5is-oDxQ9qq8gogMPUcAIADGL', '1F2F0533F6ED11E58F303010B3A0F15C', '1460536771');
INSERT INTO `token` VALUES ('3F889A5A04B111E6B5103010B3A0F15C', 'lkRoAwC8yo5mFoqkmf-1QdDFABPH1vhlRl5LW9DtMquwdxuJPFedn25umm0e1yk7Hs4oM4pCjRliRaFwkJ2e3kX9CFG-4uEbMK0-88iTnIMaP9tE6Lj4avW1eQzqI2tWMVRbACASBX', '2FF57F87FB0011E5AA223010B3A0F15C', '1460907000');
INSERT INTO `token` VALUES ('406E30F8FA4811E5BF416057182B3A68', 'dvRp4dxrf98kGWdfLxz_sI0Nx6NYvnKu_twKEHUmLDoykIyJjq7_JqDfwQLYEYVlbhUiX-BsABit9nBveIqgcb-wFczOAxIB9jBTIFreAz-dMjXL-BhxtI8s74IXJ_crTRPhABAXVT', '1F2F0533F6ED11E58F303010B3A0F15C', '1459762393');
INSERT INTO `token` VALUES ('40C5AC77033C11E6B5103010B3A0F15C', 'HcjibEcPSYQ2-CqtT-YQJAnnfP5bOQaxDaYmVpswvGNEPOlyHD2e8kwEEcpXjPh3GfmKHDxg71mdLDvtnUjkInrteK1LX9PHsjL-jlhvb9tTyA0ORE9unN6ScVKYtPHWZGPgAHAVCL', '1F2F0533F6ED11E58F303010B3A0F15C', '1460746800');
INSERT INTO `token` VALUES ('40F87C47033C11E6B5103010B3A0F15C', '33d3KDMMoRWl5VxH8NCmXoHb1f7SyIiPiVhimSQxyoUBSD7JIxRNaCj-jtRGRbM9werUe3l5X9Zdzif8lQ792AwUvHW6I9ZYFks4Erx6c1WM82XFpcKXCq6G_PkEBrjGZAYjACANOM', '2FF57F87FB0011E5AA223010B3A0F15C', '1460746800');
INSERT INTO `token` VALUES ('45FEB7D103EC11E6B5103010B3A0F15C', 'L5BjaStYEnmY0MiJCnfc8cjcwGKfC1jon9oUL6ZD16WDsvGfWXMO4r6rlxI0zOhbSrdxns9wpC2xNXzp29yK6Y4fTojIcP9w4JSsygRhTRBYu6BqkIo2Jzsm-wX4lqwUCNOjAIAKVW', '1F2F0533F6ED11E58F303010B3A0F15C', '1460822400');
INSERT INTO `token` VALUES ('48C00896010211E692FC3010B3A0F15C', '-h-tvaUNJ266PMFJLgmTG012tdqKTl8zGl4WQlo61VbTFOQDSFVv9jF0SAEA_JtXY7I9qZPAkyqm8DG5j6FQukkm4Ax-lU5dlV0EfX8Gb_AErGRBnHGgdwOvi8d02dcAQBQgAGAYJY', '1F2F0533F6ED11E58F303010B3A0F15C', '1460502000');
INSERT INTO `token` VALUES ('48FFD258010211E692FC3010B3A0F15C', 'Q9-uULekRf3Q6u1d6ExoxhgNg85XT-OQ8xJ3bIHv9s-50pjIfNRVJMQd5eZQkOiY9ZcNsCQ3ZYN45RUx5O-8qiPuis-rXBNntiXqu3E9F6wHTLjAAABNF', '2FF57F87FB0011E5AA223010B3A0F15C', '1460502000');
INSERT INTO `token` VALUES ('4C59E527032711E6B5103010B3A0F15C', 'PaM44uvGzQ4zP4VoZq4x-Rio3BBlHbkueWsnRsnwmVuQm7gV4kx4yIoK0AmZWAAxhmBOBadtO0m1eO9ALaPiwU7vGePW5ttk0OoGcycUFluFRnqRnYS34NiXOfcdw1s3IYCgAHALBF', '1F2F0533F6ED11E58F303010B3A0F15C', '1460737800');
INSERT INTO `token` VALUES ('4C87D2F7032711E6B5103010B3A0F15C', '1SM1B_P8yXqtLtHUB13deYCx4dfSdY5gj4ChfYwx79d7SfJauprd47TQzyeAlfXZXYo6_uDl1DLlbT4p3DuHReGndxV1MO8F5ojKLS74upefrbe7LPk5s1Tq-MNmjALuAGMfADAKQW', '2FF57F87FB0011E5AA223010B3A0F15C', '1460737800');
INSERT INTO `token` VALUES ('4CB7614DFFE511E5934D3010B3A0F15C', 'w_rnO3MVuKWnl7EDVWwYxXlkj1o622gw89y511EtAG4-cgXSPc68s8rJJNn-Wz3Y7ykQP5YKsqruPnJchP89kSUfDzAW6xwDeEsE7sVdQEwM_8cEV2ePyqzgv9ka_BNAKCEgADACUW', '1F2F0533F6ED11E58F303010B3A0F15C', '1460379600');
INSERT INTO `token` VALUES ('4D03394EFFE511E5934D3010B3A0F15C', '0_wrwphw54xsuDWUWwrJu6rDbNh3CAaA0hJNB0j7weaA2SHbLfb3T0AaDklHrQkVE3y5i9jjkjGNWEIPtQ9i3E763XvadOMcOdPvmb3w0x4KHJeABASPR', '2FF57F87FB0011E5AA223010B3A0F15C', '1460379600');
INSERT INTO `token` VALUES ('4F2E8980003D11E6934D3010B3A0F15C', 'Z-Ugkg6Yn1ynZoND3Hq2p31PLGRndsHyzJwuaOYUQTViIvjPzFPVtKTeSv5HipzJlCt9I49czP76JmLwS1A7A9h8oSNvZBmfkNSX_mwvQl_qKzfHD3PxrlZ0VPkXRAuhOFDaAFAKBZ', '1F2F0533F6ED11E58F303010B3A0F15C', '1460417400');
INSERT INTO `token` VALUES ('4F5FFB04003D11E6934D3010B3A0F15C', 'XB_LrfRHHip3TT2DzCRYYQcwOkxFQuWp01hHzdovjiHgXkB91ymrAzaKpSJdBzWDs-971G3G5g91ijgHmcd8wWbqI0Eq5hYt3MSzGMbrUka-5izKcwPW2b_jRZwDXcNjPNZjAHADBN', '2FF57F87FB0011E5AA223010B3A0F15C', '1460417400');
INSERT INTO `token` VALUES ('4FF11D5CFCAC11E5AB523010B3A0F15C', 'Z-Ugkg6Yn1ynZoND3Hq2p_zgT8FWtYG2e87_ek7dIBIGGiC8sInWX_0niskPNCrnUe5mCrUwJqdIYKBCD0GYSYYMlykUtrjyvz-vJbWNz43l5cPKVRHTZ5usW6yHbqn-DFKeAHAXZE', '1F2F0533F6ED11E58F303010B3A0F15C', '1460025271');
INSERT INTO `token` VALUES ('52F01223F7D111E5BF573010B3A0F15C', 'MuKuJpajmommH2r8B2Xs1cTfDNCmNQMgp7pKX01TSlH0--GhB5Txdix4V2pkqk0lzfDaTBt5vSBqbCE4qonnfwKd2mJAMR9PsJlS0_opxeVJYKZgpFhmgdVZW2AXP6eOYXObAJARQH', '1F2F0533F6ED11E58F303010B3A0F15C', '1459491411');
INSERT INTO `token` VALUES ('542D71E500ED11E692FC3010B3A0F15C', 'OVciNBoiNLrDtzx07qmGSSALBtnqICtdecAo7kkGE5mNNLR42Rj6PeTsbqFRCcODLxHiS4sx1kPlwftJEP_uYTNyJ58Tr0qBtoJ_bIRQfG7FKyZ0M_udINoXe8Qb9kapPSXiAAAEHM', '1F2F0533F6ED11E58F303010B3A0F15C', '1460493000');
INSERT INTO `token` VALUES ('545C230500ED11E692FC3010B3A0F15C', 'o3VR8vqnAI7x85Js2-dNeWxX0BPDFKnJejp7lkHZXIpeEE62jMZ8fBq-6oOkgg5GaTdcWyk5TVavO_KqiD_L-4nbKDNMXBmYf_K4r11X5oTSSWxjakG_1qpWPmir99EiRESiAGAGTH', '2FF57F87FB0011E5AA223010B3A0F15C', '1460493000');
INSERT INTO `token` VALUES ('5583DF2B02BA11E6AB2E3010B3A0F15C', 'JoggzQZKxHFWkXf2C6hSYvj0cm81ghNfzF0LUKccdpBFU0cVFpV1oB-7USBbPmUuPy-RRaMHILBRSkRqkqYDohfRzgu3sA1dQFv21zoC098WTTU-BPyLUdJpG-f4UfUoKADgAIATJM', '2FF57F87FB0011E5AA223010B3A0F15C', '1460691000');
INSERT INTO `token` VALUES ('56CC2C8C014511E68B623010B3A0F15C', 'TZMbJEpOLOEaWIky499eVQP7-l-HcTgY8AVu4G6J1Jh9psolqAaJVa5TkuzOWF7A5zBBEDWl2mEGSb7NwmjKDmF_MKSMfg3IiBzwCmltrfqtKmIJbZt65XJBF1eP2s57CBFgABATSW', '2FF57F87FB0011E5AA223010B3A0F15C', '1460530800');
INSERT INTO `token` VALUES ('57FF845C031211E6B5103010B3A0F15C', '6T_s5yROyvyAf1yAd7r9Xtrd40cqZF-uOioV35tGJ2bYYTryxPLHz8901VwF37ckTDylft9BiVTSsTYSMaHRZI88w5MMKJRhXtHoyqDVBKVZLgW2qp7JM-GKRndLHJhWQSYeAGAMOU', '1F2F0533F6ED11E58F303010B3A0F15C', '1460728800');
INSERT INTO `token` VALUES ('58313558FFD011E5934D3010B3A0F15C', 'FOoMZ0e5JTtIiW2-yahgOWXOjL-4O0DHNRbdySoIkqo9ew_Nd2J5mFLReKhWo8C9ZmR9RpV31E3eevBk0w25E0mOUJzOLrwurgsjWqNwN248e_-rtINFdqEB39Ej2KHcBEUeAAAPSI', '1F2F0533F6ED11E58F303010B3A0F15C', '1460370600');
INSERT INTO `token` VALUES ('5837AB5C031211E6B5103010B3A0F15C', 'beNmefdA2_86liFAWg1dAcgrfj9vwcs9keaALzaE2xDUoY2uK3_iidu6SMr9sPbGH0MCdWAOf9OepUoXquHlv6JWZlcarixyNIejpzGs6ZcSCYhAAAFPL', '2FF57F87FB0011E5AA223010B3A0F15C', '1460728800');
INSERT INTO `token` VALUES ('586083EEFFD011E5934D3010B3A0F15C', 'YeHoFwaxsiXCWVLBmLcdAIDZMkuYIkrdK3bUs9MYi785Yw2jQcO0cwBaEQpoWnqQQ4mQRcXGnnKGaPkG4V-kpTuvqoFGXcPdyAiFmwNKYvAZ6J7H4zBNzfk83GCUMBRLSHQaAIAZOX', '2FF57F87FB0011E5AA223010B3A0F15C', '1460370600');
INSERT INTO `token` VALUES ('5C1B620C01F511E689F53010B3A0F15C', '4fMr244kiW2TRTnvPeUDvAfi2WwqQ4_ZJDcEdN6CBgKbJx-iRJPSoBWhU1N0ggocKaNFA6otj3-PtVw1rZVzD4EGeJOF7kpBr_16xSbETYkQPIeAFAFEX', '1F2F0533F6ED11E58F303010B3A0F15C', '1460606400');
INSERT INTO `token` VALUES ('5C5E392201F511E689F53010B3A0F15C', 'aDRUJMVSYCJcdqUOf7wcd6UmbyjBr33G75TphD4yiayXIn4ThOH5d3kYGzLOjETOdfwktYUf_eeLX25rp-KWdZEWCK61f3SdEWOBADq2rZcLMChAIASTJ', '2FF57F87FB0011E5AA223010B3A0F15C', '1460606400');
INSERT INTO `token` VALUES ('5D514E8303C211E6B5103010B3A0F15C', 'UrGOuf3s7uVi8oI4-tsL_2QyOQexNOcrl9QXnqTt3wqlk0WylZE9qgbymy9N_UolhifSh-enzPOcffd-WoS96kfSnLsVDiDTAJHjVePTh3MQZEfABAENS', '2FF57F87FB0011E5AA223010B3A0F15C', '1460804400');
INSERT INTO `token` VALUES ('5E84FF3D024D11E6AB2E3010B3A0F15C', 'unnPvkdq_bEN6pPfA44S8rLjV9yiLsd8DzDPztTiWrIhhM1_103nsbIvk0i1ao2LsKVYsk-5aCle471ncSeFfySRE5EGRQleKuBGJW-9VaxPCqBZTqzVhAuV2TxDpCnnEWRgAJAXPK', '2FF57F87FB0011E5AA223010B3A0F15C', '1460644200');
INSERT INTO `token` VALUES ('6046920500D811E692FC3010B3A0F15C', 'vePyAjvSCkyBrBiLAVQnjaw4YTS6Uuu03MuLz9l1GpOOUY05fHLAn0iu0JBxvcwgmmDZwkRSulRHGKnIZVBUv-BUwCHCquEeJsmI3uudDA5xGy0CBm97SDA5jqdnRwPWPBZaACAWIG', '1F2F0533F6ED11E58F303010B3A0F15C', '1460484000');
INSERT INTO `token` VALUES ('608A78D500D811E692FC3010B3A0F15C', 'uaF5k4bnxPq25Prq3hZfJOWUXoeQArUmQMRdppBJrNdPTnKvpmj-ka-EX78l33d1qS9F89CZkUy1CoYqwcMbdQ4xsirL0mUIgva5hbdrsdQK8RM-hc_HetOtnTMSQZCUUZZgADAHMM', '2FF57F87FB0011E5AA223010B3A0F15C', '1460484001');
INSERT INTO `token` VALUES ('62459524047211E6B5103010B3A0F15C', '_m9EAwa_cR53DjFQu7mLgRRGSv5UXIg-4paCMiW_G23SYQNuWWKUlRzQVai_hqADQrl6zmnarJmEiNc1TtljUvWAD9vOE7Dt4QiGsOtHEfWaPgz1-Y0V2GOcEybjt2TZCXOgAGAMHY', '1F2F0533F6ED11E58F303010B3A0F15C', '1460880000');
INSERT INTO `token` VALUES ('63B8867EFFBB11E5AFEF3010B3A0F15C', 'w3iwyV07W4JBbZOzc6A4GANlbmMTK6t2_blePeQWsDbLkhxMJvjjyjv6vG1rr8HZnE87yYRLgPbWHVFFMmVB3mm5N8PN3BxBMW70k8Smz-OwBrn8m6NwAYyBQPAYsBHRAHQfAIAQXG', '2FF57F87FB0011E5AA223010B3A0F15C', '1460361600');
INSERT INTO `token` VALUES ('6431358C02FD11E680C63010B3A0F15C', 'i161uQ7RVOVyjzTYwWeJmyTkLHHBoNCQJbAszG23J1uAL_dHPrueTH8eXeway461koFTvnMUMj-WcCiHWcegfynqxTRtz8upPFzI30-Y3tltjNTMeveXZAQgWLvy7dd3IDUjAAAMHY', '1F2F0533F6ED11E58F303010B3A0F15C', '1460719801');
INSERT INTO `token` VALUES ('64D2C0CE04CA11E6B5103010B3A0F15C', 'i161uQ7RVOVyjzTYwWeJmyT94gsrbY529zt6ByGFJeftaXzKw3POCMFD5nccrtV3LlQLpePm0YvnuS-Mld7mA8OIB5GsayVEGrkH9-negKERERdADAGTG', '1F2F0533F6ED11E58F303010B3A0F15C', '1460917800');
INSERT INTO `token` VALUES ('6A36AF3D023811E6AB2E3010B3A0F15C', 'Z8zGXVtyQU2uMEqrrE-KZ-SrLgFJ2aaxEaaaY7xnKa0tdO3h8k9jgUiQJElfJlnTbFUDmKhegcKoZF8e3iwlVEeCG6LtfEdtlwsqoRYZhBlXKhXOMCjcEzIJfX25ObVYQNVjABAIFN', '1F2F0533F6ED11E58F303010B3A0F15C', '1460635200');
INSERT INTO `token` VALUES ('6B72790C00C311E692FC3010B3A0F15C', 'C87ohOR37y2XLzz0ZJNnvNQGoB-dhnjKIsnqf05yarL_9ipxFOZc6gdeUuhCk1wIp2oaZ-9MDSapAw_Kt43gxT3FTh-hx8qmpAFxUfDZ9s64GpzZmV9YO1oXxHc3suFFGRHhABAMXN', '1F2F0533F6ED11E58F303010B3A0F15C', '1460475000');
INSERT INTO `token` VALUES ('6C9E76F5029011E6AB2E3010B3A0F15C', 'MT1zj0xXRIRTu-5-AQoJRGaMMQRvjIJOVVeQR90oCrPvN-epQR6F2Y-lTDa9FXp0tSm4Gq5_M0IHnnbBeVx8FtLT902v2OKWgVWW-SSlHY_aQ68R509sxIoBFKQ7_EgxAZJgAGAUTD', '1F2F0533F6ED11E58F303010B3A0F15C', '1460673000');
INSERT INTO `token` VALUES ('6DD95C20045D11E6B5103010B3A0F15C', 'FHBAvU27YkiWwJFEu8og-yRbnmH_aYj6csgQgUq4xExilaoVO8xBp80gKDXOtB-HsBU78VGQVEmYkkGL6Gp4f57NH-MFGvXOhAx6GFD1WHF4-MmKWgdw2HCBjfiCRYa_FWXgABAVMN', '2FF57F87FB0011E5AA223010B3A0F15C', '1460871000');
INSERT INTO `token` VALUES ('6F38B7F902E811E680403010B3A0F15C', 'Ig8LSE2BfvAb5ZG2L8VT-0K_7KLZfIbwS062CgFKyt4URd8uQ5UEuqw8777SEHpQxZ7SCo8a9z7tSs6QiGMXZ71XQTSi1FHO4rmB-8MEglIRYMgAEAJQE', '1F2F0533F6ED11E58F303010B3A0F15C', '1460710800');
INSERT INTO `token` VALUES ('744F2BD3039811E6B5103010B3A0F15C', '-BGJPTGTRwRNFXS_cgjN2WJ7mbI40A--0zK3uqEEzh8zd-koQ08WHew1zEaeGVLyOCYjj1onv9wjkEv2q0eVgBM65uBn2SNOTKwV0O2TlZhhp6Km_zPX-jNsKiT1HVK7HTIdAGAOAP', '1F2F0533F6ED11E58F303010B3A0F15C', '1460786400');
INSERT INTO `token` VALUES ('7A9A3921044811E6B5103010B3A0F15C', 'aDouAnXGXhhn_Y3qwy947U7z8ahS-Z_d38UNIKjp6Zl8HFldXr9Qj3xDHRwAhVZo9P27HfZnAPnRYGKM24h-nBq0COt0yA8I1tH8uZvO54W2Y-eTBjEbsIikz3jsTP5lUASbAJACIA', '2FF57F87FB0011E5AA223010B3A0F15C', '1460862002');
INSERT INTO `token` VALUES ('7B6D08D102D311E680403010B3A0F15C', 'z1gn_XvmfJK_lMIye4_jWxLXKVY3_1gQ9cd_wmAsfSZcqANgEuQdwFnu0tuGeAqb5rmlsS4k0bUdHCUSCk8LlQIq2knAtb7QUMFZ7HnC9cFlJe9wzBWqnuaHpYXZVB7lHUPjAGAOBW', '1F2F0533F6ED11E58F303010B3A0F15C', '1460701801');
INSERT INTO `token` VALUES ('7BFF57F704A011E6B5103010B3A0F15C', '2fa5hizhOucUjEar9NTNCOf_NL8GnwX_qcgBm7ja22oYhMWyDfnBnZ_eldU9nAwLKQTWi1BWzaWwnlo1aGr5Hx1yXWjZTfBr5r7XINsZPPDOuTLD16vGxSQzqumiHrURWGAfAAAEOI', '2FF57F87FB0011E5AA223010B3A0F15C', '1460899800');
INSERT INTO `token` VALUES ('7D72F7D9FF8711E5BC003010B3A0F15C', 'aHFMnw66KoR4F7lzIJJrBp5t7sRugnrQICDsnjtL_k6ITa-wqaD5xmhTV81ToxyFe7kvCarbCLK97N0yVuckJavypZ55KwSPy6eVeXeUX3sDPRiAGAMYA', '1F2F0533F6ED11E58F303010B3A0F15C', '1460339309');
INSERT INTO `token` VALUES ('8290C65103DB11E6B5103010B3A0F15C', 'Yefumfh9-Spq2Mu4tST3vWv8DToCnMt2UZq09XeyMNf6YgiyLiRjzXyWJ6U4qf5rkT9dm76rSLqRyQngZE4GGvg_oz-KyAMuJ4iMUIHdH1BlnAd_VDUySVohWoFicxhWWMTiABARHO', '1F2F0533F6ED11E58F303010B3A0F15C', '1460815200');
INSERT INTO `token` VALUES ('84EB9BF0043311E6B5103010B3A0F15C', '33d3KDMMoRWl5VxH8NCmXj__tmmxPevhLve7AHO4wzo5n2SQerk1edx7cLzFL4QWiNhscvYYh3XrfuBQlk6-QO38q2CNms5NAo2oiL0sBTW4kKj045zTHFnGdda4kw_aMWYfAGAVFH', '2FF57F87FB0011E5AA223010B3A0F15C', '1460853000');
INSERT INTO `token` VALUES ('8B60F25EFC8311E5AD3A3010B3A0F15C', '6DVzO0JzFOA88fo144kH5enediqzG-Il-DeVF0ajEqM6wrDAXDx3eDUDdkNeEav3EFxvL47BG2v6_vSw2KtRLHl38ZgEhxiVWKI-tI8j_WAkaZF49sdFSkDAE1MMRWfBKTUbACASBC', '1F2F0533F6ED11E58F303010B3A0F15C', '1460007761');
INSERT INTO `token` VALUES ('8BA7DD39002C11E6934D3010B3A0F15C', 'vajsk5To8GolRpVLG5i4jsnx6iFgV7bMxFcgIpJTJVSRftPt-YxVYko8LlZREhsT4mI1AbIzS9AqQyJbLJX_kdbPBHb8hhL5Ywe3Ps2PCs2q9w2e9l6maExMw6GKWPN-CGBjABAWAV', '1F2F0533F6ED11E58F303010B3A0F15C', '1460410200');
INSERT INTO `token` VALUES ('8BDE09CC002C11E6934D3010B3A0F15C', 'isDVhEWSd__VJJqNVnE1YUAdR2rmXuHb4rN8GE0EI8GEXEobCe0RV_8MTrotIebA25BDuWefBa3_fMkBG4DdKrAvjF2p6FZTJlggO83vfLuYGd0V9okjQAA1GzmN613BUNBbAHABUI', '2FF57F87FB0011E5AA223010B3A0F15C', '1460410200');
INSERT INTO `token` VALUES ('8C14251EF7F811E5BEC23010B3A0F15C', 'L5BjaStYEnmY0MiJCnfc8U69zNr1Z8Ga5QBx4b3oE1EJ1CawlQKkQQzz4Pc-5k_7KUf3FxsqVehtX6MjQmgwTu9V5kjqjkx44M5mF6AsnPT7Bs8c76McSJwc6Z3_Mhu5HSRbAIAUFE', '1F2F0533F6ED11E58F303010B3A0F15C', '1459508257');
INSERT INTO `token` VALUES ('9087EAF0041E11E6B5103010B3A0F15C', 'DUrEwh9yma1t0UDkVH9f-DL4NXYQc4LYSub8Zp2PxYq_w48hO71kIxeM1nae7SJzg7RBySn6_4NcvSs1iSe-aIsI_A1KsbrzRnQZv8XWxo4WQSbAIAFZB', '2FF57F87FB0011E5AA223010B3A0F15C', '1460844000');
INSERT INTO `token` VALUES ('91F372EC02A911E6AB2E3010B3A0F15C', 'lp8MhmtSB76Mgkj05rMqKHLOBql-an8W-I4Pyb6pqovaMo4GTbk8VGgiQwzCg_mWCF55PWFyHSsh9sp1THF_ZvVE1vGlPQZa1sa6UQL9J_17xVpFuPnNkR8Buxo2AQbFFRIeADAMYP', '2FF57F87FB0011E5AA223010B3A0F15C', '1460683800');
INSERT INTO `token` VALUES ('947FDC1C030111E680C63010B3A0F15C', 'isDVhEWSd__VJJqNVnE1YWtU9F0A8z0bY-_6iYUqcCJ4r2nmxdeCdj235HpX12Zt-5DcsFfgYi19rGs_BvNFcEGt6UfEko9Fj4IN9u0AKrAwPm4mdSrIOGM1yefS7RrYEBWfAHATQL', '2FF57F87FB0011E5AA223010B3A0F15C', '1460721600');
INSERT INTO `token` VALUES ('959E0262FA3711E5905D6057182B3A68', '5VVOT4PpTxtvKQmu9PQunbrxU96fzgGE5R_NxJiJOidlW65mj99LQip2ya9o9LWtjDMwbj8MRcHkmqVAFMvmbeIT4viF5DI9vlziCABgxNw265pl3eJCIOYjGzOvB-PnKEWcACAOIA', '1F2F0533F6ED11E58F303010B3A0F15C', '1459755234');
INSERT INTO `token` VALUES ('9649F1DDFA5811E58AED6057182B3A68', 'Tg7hhyE4Wuijy521qTAgzLQFdVQPkGDX0-gZESEZ6DV_dIPSvByMFUFIGJAhrKPivrejntDNzM7Gqaql7FxpuQpQd722L1bQN4MohkcZwVzYXpNvbN1XcAOY2lMU5bnKSSCgAGANNX', '1F2F0533F6ED11E58F303010B3A0F15C', '1459769409');
INSERT INTO `token` VALUES ('9846FD6201E411E6AF103010B3A0F15C', 'aDRUJMVSYCJcdqUOf7wcd-gdWDb-cEIrcjUJh7NIiTT0rEi6jZ02yJss0xT-9v1382kGx34Yk2i2ts7A_OSZxXU_kruZSWj8oa90f6LKjRWGlkAH8RtsOV6d0fNpLuihUKWjAFAKAM', '2FF57F87FB0011E5AA223010B3A0F15C', '1460599200');
INSERT INTO `token` VALUES ('998730E703B111E6B5103010B3A0F15C', 'QyMLQBHNbWezVr_3BGj9U8PHg6FQbJ_yKFcLM4H1tL4Bu36MHDBkkopixDvSSsovW_DgnqJMSRukILv-rOceo0twiZd2JsPuqUCs3vau_5YUBvErjgEX01JEj7a8AI4WREZjAEAWJE', '2FF57F87FB0011E5AA223010B3A0F15C', '1460797200');
INSERT INTO `token` VALUES ('9AFC591E023C11E6AB2E3010B3A0F15C', 'SCZlJh8wNx36XM0Y5E76tD3hrko_aVgtPjz_O7tSekvrPfJpJA4tNyhBfyqZENKBtBhjofeZBknS0f6oIeBt90MAAXtJXnjntJt1xh5fIezX9OIlqfVCvph26-wG-d7lLJOeAIAHFK', '2FF57F87FB0011E5AA223010B3A0F15C', '1460637000');
INSERT INTO `token` VALUES ('9B718DDCF82811E583636057182B3A68', 'lnPfi8KcicOcH8Jy-HBoULxyf3lHTuImgpmlfnDAz6X9LpAgwDw4YusxwziNiDf27zdlOrVYqR8woGJUuFyqril8wRungRMgg5kcwBKabENaYerbbsucYgZ4PBsgsI0qGVGjAJAKZK', '1F2F0533F6ED11E58F303010B3A0F15C', '1459528899');
INSERT INTO `token` VALUES ('9C1322F0040911E6B5103010B3A0F15C', 'JBVFYPmgL0OkRhK-wZ_iC8YORkg505q3ADe4v8BlyfG9BGSAXoxWo3_CCXeTGHWf7QVjHAViCk5wxslxnPZnPNGt8kH36Qy4sDHLCQmHG8KExjWQOZeZGnjB1XNZHcnSQPGbAJAWVM', '2FF57F87FB0011E5AA223010B3A0F15C', '1460835000');
INSERT INTO `token` VALUES ('9C545E3600C711E692FC3010B3A0F15C', 'g3PG9xSsk5YGNGJeGhYk6ndkFwLejsKFRtsVV8PHppLBLEO_LfLBluxEzhrIuVDHqXbJgOsoldmHQRAwRnqbORwPYFxmtMr0Azh-dVnfQ6-uFC2Hyz8jcrtKbLRrKeTWFFMeAEAIIP', '2FF57F87FB0011E5AA223010B3A0F15C', '1460476800');
INSERT INTO `token` VALUES ('9D7993DD008D11E6AB663010B3A0F15C', 'djHLmFu36ph33g-SBNlrfQgQApldwYXP_70Q7UR2AcB2CScxSSKlyK0hQzbNkMQQIKfUewQrd_TesUVmltUN0v00wQh_F-K1IyUSfsPtbaOLKZEk8h_tmkoRotaYRRDQWFWiAEAIRI', '1F2F0533F6ED11E58F303010B3A0F15C', '1460451891');
INSERT INTO `token` VALUES ('9D7F4570029411E6AB2E3010B3A0F15C', 'F6rxXtw0N8o-U8zYu8OS8LDhqz-aDWU1j7DuPQL-7dhvzlsYqzndlCEeJL2OvwBgLusp4pNtFXuHLnVSQeQPjB3edrmvmW5nvppfFKja6OdtAc3A0yXQ6Nuryk5FWTYNCONeACAXJI', '2FF57F87FB0011E5AA223010B3A0F15C', '1460674800');
INSERT INTO `token` VALUES ('9EE81ED0046111E6B5103010B3A0F15C', 'gPJ8NQ6UraUncnaX1lWa9bXTTtiiKZV27T3vCOkZfOexj-fEnopCIjN05DUNuB8Wkni4OZOACOUF1vXdzd_tRZ6jT2VUro05svPDLASh0eET7VASc5T1l0AvkkOWSADkKBOfAFAQHU', '1F2F0533F6ED11E58F303010B3A0F15C', '1460872800');
INSERT INTO `token` VALUES ('A14BEF7B04B911E6B5103010B3A0F15C', 'C87ohOR37y2XLzz0ZJNnvG1c167ErauqI5DXfe0Fw8RJTD5fYDZF_kr4AwUcAUZ78HeybzylzgGF7EGUEKS2dhFXOU7B6Ir2sVHRaMlgSpe0ST7bM64FIUzMn5j_M9KkXZUfAGAEGB', '1F2F0533F6ED11E58F303010B3A0F15C', '1460910600');
INSERT INTO `token` VALUES ('A19B6223017711E68B013010B3A0F15C', 'dvRp4dxrf98kGWdfLxz_sALWoPVz63KkxOber7MH5cqBF4gB-n70I-_1F9Kd1i5fwAjSmUQr9C8osK5RqjzILQID6duEFr3xpaDxdV729RmmYgXNgZdoo0xbvBaJTz20SGXdAFAWKB', '1F2F0533F6ED11E58F303010B3A0F15C', '1460552400');
INSERT INTO `token` VALUES ('A22D7FCFFD4B11E58D1A3010B3A0F15C', 'U2j2pWH1IxHp1Z9Mg2fHvVfE4Mj3UPEP4GAvNXkJBxOYI5noZNsSxegBDk1njHL0wfjfAsrPczMecjTFke5n-dc3zJopMQfdJwBCHFIYI5W4iDoXpWlF_GIzeJm1Us8wCWPgAHAMGZ', '1F2F0533F6ED11E58F303010B3A0F15C', '1460093698');
INSERT INTO `token` VALUES ('A538B24F039C11E6B5103010B3A0F15C', 'GtpPCpMw7Te4Xt_IPse0BjOpOLEoz93HEtX7wiy5ZgSXeqL2ZkMeGZrBieedVjpP2u4CI0k1VmOkL9T-OF6pZbNsBWaoyzEvYYVT-vXlLgpUinY1P1Yuz20W4WnrqSrUOMYaAIATRD', '2FF57F87FB0011E5AA223010B3A0F15C', '1460788200');
INSERT INTO `token` VALUES ('A692FE78022711E6B50F3010B3A0F15C', 'AWTYXG3dBPKPNJqJsM1ogL262Fehyvl3pDd-5twPRy5uiszt3tAKQ82oDCFESynL-WzxH9zelKzIyrOKOn6y9WHUb1m9_XwMQDLI31PkEkgi6-Wz6QjP6NJLaDAxLV-iORJcACAILD', '1F2F0533F6ED11E58F303010B3A0F15C', '1460628000');
INSERT INTO `token` VALUES ('A7C3BAAD03F411E6B5103010B3A0F15C', 'CqM92ND9inrNW8XBPB8o0m3y3Ci7Rhg-ILEvR420diMGmS6mGrwE9B_48tLp4o3lMqjds6ApLZnQMCt8pdyX8OdKy-104F81JAZR6KKnJocV6MGM5xT2KDZz7t1rWNprHRZhAIABLT', '2FF57F87FB0011E5AA223010B3A0F15C', '1460826000');
INSERT INTO `token` VALUES ('A7EC6B0E00B211E692FC3010B3A0F15C', 'N_xAmdojvqDeSFbJQmtoGzpgr4PrHclE2mkTMBD5livbVZZvZqIi5zb-D6bcCjniKLfIt4QRvctymytvuqefs2wA-qNE-6rRnQ_ZPnEkCMz2lsSVb_gu08Iu3TqxnJjjVXHfAAAYAK', '1F2F0533F6ED11E58F303010B3A0F15C', '1460467800');
INSERT INTO `token` VALUES ('A8D47E8E027F11E6AB2E3010B3A0F15C', 'aHFMnw66KoR4F7lzIJJrBrj3gAfTZ0gjjBa9YYcO-5QcSu7-IEmQqDjrM0vy0lws9okJczekX4sS7YLgetmkOayeF9ZIxGBIj__qnEHH2OH8NV2PZVEjws6UF4nTHAgvSLXaAIAIBN', '1F2F0533F6ED11E58F303010B3A0F15C', '1460665799');
INSERT INTO `token` VALUES ('ADD133E1014111E6813D3010B3A0F15C', 'OAnPjRY6NOfyyT_0AnrXeAHXnQs2oaR8abrQW5ivUa9YH_Of0WPTd0knj9DQq49gMWySsP21ywWp_PvC8WD3ntcY3WabIW5caH6ioNJUB6i9fadd1pvfcY1VnBhMnXPMGDVeAJAJGY', '1F2F0533F6ED11E58F303010B3A0F15C', '1460529228');
INSERT INTO `token` VALUES ('AF5E7F8E04FC11E6B5103010B3A0F15C', 'beNmefdA2_86liFAWg1dAal0umr7aMqn2IgYffS76A8dd81fUhdFor1hSVnyC_ItN8w2oXcv1gfPdlm982-IcizuWGZZKxyTIdSoAY-8D9q_v8zkfnE4Ro-EOlsZSIDmUOBhAGAGAN', '2FF57F87FB0011E5AA223010B3A0F15C', '1460939400');
INSERT INTO `token` VALUES ('B02755AFFBDE11E59D1F3010B3A0F15C', 'vePyAjvSCkyBrBiLAVQnjfP49AlWUEciPlda4Gr_aI4VTXmlOs0W2UqQvNdpKAFhPNyGGdq90ViOs2J7k4YTS3ooWyw_7z3WXeCjVzVhx9-E2Uf1KLCFOcQGl0lTtXK6MTQcADARYV', '1F2F0533F6ED11E58F303010B3A0F15C', '1459936956');
INSERT INTO `token` VALUES ('B0C8A8A2038711E6B5103010B3A0F15C', 'pZXvqWIDdoIQvIPIIwzwuhQPZMJ0vnpR3ysq9V0M7mBum2tn_Qgt3j3vlhzg0q3zIY4ZRLhj48vaYt8zSPyp_VR96yj7eSkrNLCRz3y8kATggeiZW57VsRP77yyuFnRNASEeAEAIAM', '1F2F0533F6ED11E58F303010B3A0F15C', '1460779200');
INSERT INTO `token` VALUES ('B12E781CFD2911E594C03010B3A0F15C', '7-g-PNyptCTxyIOWh9LmtVvjSkp8IkkDVDfPAkv6wJzlHf--lSR2pQ_f8fZudlxkYDZ9diUNtkiqbNxbsMJfzFlaMM3DW83HNbfVxDmR5MmQ_M02vWJi6bAURmbxS66MPIVaAGACVI', '1F2F0533F6ED11E58F303010B3A0F15C', '1460079121');
INSERT INTO `token` VALUES ('B14507C1038711E6B5103010B3A0F15C', 'qnjqZ2GqXPa53Od_90D07SAlpYg9CWnOVQCYI8vqPgzE0VLAUyH2FTzDQ6jO-R3KF6XQV8kqE9MZe_cv6gv-uw0-Kh4v0YA2YntH2hOig2fqwGNIzQpY5MKwK2R0F681MEVbABAPHY', '2FF57F87FB0011E5AA223010B3A0F15C', '1460779201');
INSERT INTO `token` VALUES ('B2734F87FD3A11E597D73010B3A0F15C', '-BGJPTGTRwRNFXS_cgjN2boWmpHec0OYtQdfW9ygbdoCLfTptTrKBadPWs11IQqktZmHWS7qSmDtbhxSOsiYJMHBK3oRewe4ZvqPW0awIHxjmj2K8APkBrKChC9vsf7dYACdAEANPR', '1F2F0533F6ED11E58F303010B3A0F15C', '1460086424');
INSERT INTO `token` VALUES ('B29CAA86FD3A11E597D73010B3A0F15C', 'hRrAaBOw_pyXPLwG2RFszGoTkEFjdEcjAUX61zfTa7szKmC9yJU9-KCT-ZD9Rh3lp4ZXkRBjnoBc3F5foT4u3_pcuf8YVQtnEAbPGXMIhCYuV5-6Dl3ssNVCs24_1vNlOISaAFAXRB', '2FF57F87FB0011E5AA223010B3A0F15C', '1460086425');
INSERT INTO `token` VALUES ('B4B48D62026A11E6AB2E3010B3A0F15C', 'eOlSEa9vxOKfm7pWkvpSOdIyWZswlT4k-V-NCwiEvJjkJ9_qfpx81mRNlh5yoPMXX37JMppkSLKkpsqt9KQUzMbklSnDNUKULaoSmMtQw4g-O21OwGrw2ZTsZbwYok2XBXKeAAAVRE', '1F2F0533F6ED11E58F303010B3A0F15C', '1460656800');
INSERT INTO `token` VALUES ('B747A97F02C211E6AB2E3010B3A0F15C', 'TSpCt4OZJVezYYa5euvtCMqJFh6z1rFf7HdBLWCa7WUQwGi0SvTd5YjpsfwTgPxsr32EWP38spDI1hnF0LrHJyCFxB8Q7E-R6vfwswebMLNoTXdxfWj9yGnaoN5c0iLLKTDhACAGPZ', '1F2F0533F6ED11E58F303010B3A0F15C', '1460694600');
INSERT INTO `token` VALUES ('B87AC8BE048F11E6B5103010B3A0F15C', 'eVW7sgZm5cl2XB65bSsv2tMxV2TXmaQ_6dfI8EM0SdpU2VwOui8YvzB9lFyzplZ2Z2F9Emd9aGpuHo9_j9ibS3tXya_ZaDKq4j7V03tLU33poI2c0fSPwhKNOgk4Q78aAPWfABATIP', '2FF57F87FB0011E5AA223010B3A0F15C', '1460892600');
INSERT INTO `token` VALUES ('BA6EB5CCFC9811E599263010B3A0F15C', 'Yefumfh9-Spq2Mu4tST3vftUCyrNJFSKKEPPrR_WQG-p_D-HyN9rGw4e_622b4SLwUZq4Nb_VgQvlIiaiEOoyQhURpb2gL5oSKDbKX90V7lyDpYerqT0dVcnhayT4JAvPTAaAFALPW', '1F2F0533F6ED11E58F303010B3A0F15C', '1460016859');
INSERT INTO `token` VALUES ('BAF85D8E04E711E6B5103010B3A0F15C', 'qDFd7qcIHqfC3uSSMcZkZShtq-E_ZvFCZEFwn1iWNojl2w1NyqtTJpCFMFB_soqK3Cci8Bm0ZK_j-Gi5hBvjI_MaEtMrdBitn7Zym_PU_-B_RCJ1RGBe1_iQhOvzEgW0YOWbAHANDL', '2FF57F87FB0011E5AA223010B3A0F15C', '1460930400');
INSERT INTO `token` VALUES ('BD162C92FB0111E5AAF33010B3A0F15C', 'cYGk8qQVOuD_8fBttfnvoUNw-zpF-uDHVkOSFrQD87oodyluZsOZVdvTu3CvW3Jcm6nmXiabEZRmCKnmvW6uuBU3GyGN51i3R0pDbV76GmXrOwsJovftzMfyGDwr1BBEVFYdAIAGQJ', '2FF57F87FB0011E5AA223010B3A0F15C', '1459842059');
INSERT INTO `token` VALUES ('BDADDBAD053F11E6AFF53010B3A0F15C', 'gj6FT0nQM6rKyAI-j3jQkB6kUlgZCm-rMGB7nmT-wYkvF5KxAunYGY3jo5Xn1xzjhC9SZO_gb-9cp0CcGSENlSUOHeonMzHvClFcfg8H6o_1KyJNJ7q8RKTQs76A7ymCDVJhACAYQE', '1F2F0533F6ED11E58F303010B3A0F15C', '1460968200');
INSERT INTO `token` VALUES ('BDE0ACCA053F11E6AFF53010B3A0F15C', 'DUrEwh9yma1t0UDkVH9f-K3zLHIqrRJPG1keH5QljFVb7At05saJ5dnW7bWHpX_SXMsi2jJ-oVdlKY-gWDTFxigLlwd03Ss-0LWDBQKdKgJvPNGgojhtBEozWJ_8LHkMWWCjAAAMMS', '2FF57F87FB0011E5AA223010B3A0F15C', '1460968200');
INSERT INTO `token` VALUES ('BEF6BAA003CA11E6B5103010B3A0F15C', 'x7XAtg6CBs1ZNgFnuVi8ncfnt9UEFTrEfmH6np2_XDacrQfV3JPsCltwEASVZwbrkVTBeT9-bZVykHFNOiS85r0T2-gEHPj309xPL2kSjHEXPCdAHALDQ', '1F2F0533F6ED11E58F303010B3A0F15C', '1460808000');
INSERT INTO `token` VALUES ('C31F057EFFBC11E5AFEF3010B3A0F15C', 'E_mTU45r0Qax6V4qPrV_IWXc_6Zm3YhAWvaqtkPl__4TzGhBrv1a1oyh-yCqUFPKBRHKseBAlJCxZ4O2r92VsR2Te7AB2DmHTBgV2z3Y0UB_mtbMF626govwfzHse1NYBUFaAGAOHS', '1F2F0533F6ED11E58F303010B3A0F15C', '1460362189');
INSERT INTO `token` VALUES ('C699716504D211E6B5103010B3A0F15C', '5TBA5f1YBm-CNI4NlubqTSLErhSg20SFIrKgu8ehQjiiaYdUyOxyZgZo1ayNtjR6PvNseKXWmbt-rJct30otVlLfpkK77y3yqvja1BAm6geeDU3Ld1KTfx0uIwdKSmd2ZBJjAHACEM', '2FF57F87FB0011E5AA223010B3A0F15C', '1460921400');
INSERT INTO `token` VALUES ('C7FCA695001B11E6934D3010B3A0F15C', 'FOoMZ0e5JTtIiW2-yahgObfX_VWUaTYSs1SGbFQn-gan70XmpAzUU6C1iWZzt90Y09NTHo728E6kiq1F_cLCWv-o597zgutP_JW2mpMtPABgMkVvkUTT7qh3hb_80bneGMYiAAAJPB', '1F2F0533F6ED11E58F303010B3A0F15C', '1460403000');
INSERT INTO `token` VALUES ('C8260085001B11E6934D3010B3A0F15C', 'rtrNtPSI-DJ1ZcqmB-Tn8mehyRRMIeisf8DwCGd9VDECa2ofIrQVODhY85bhzWD3kTCyqnT0TbcGrQUKUWVJQrP-kedtm5UVWM-QYGbL6BZc9QBYQ4Kk1tVMt_hV2p8_ADKhAJAPSZ', '2FF57F87FB0011E5AA223010B3A0F15C', '1460403000');
INSERT INTO `token` VALUES ('CA44A9AC052A11E688B93010B3A0F15C', '6T_s5yROyvyAf1yAd7r9XhSz0wTN4-lg3BL-ftz_QMMHq0ET1HU3VXMJ4sRTaGEWvPMr67dgDJwj1RiCRSM2qHt0cFa2Xnpmi9ZAv0MHnJoTOHcAJAIDW', '1F2F0533F6ED11E58F303010B3A0F15C', '1460959202');
INSERT INTO `token` VALUES ('CB69CEAC052A11E688B93010B3A0F15C', 'juGyEsMF85_Mfu0pBBUyTxlfQfN08mIHyhf94vjk0QmPcWNWVdKDvusU6JVjQzSijhf3fnQC8SOgnIe1Gf5Yq4ePR2r6fQC02aXKyA-z37sHPAhAEALTE', '2FF57F87FB0011E5AA223010B3A0F15C', '1460959204');
INSERT INTO `token` VALUES ('D0B743FC02F011E680403010B3A0F15C', 'o3VR8vqnAI7x85Js2-dNeQaMFawAUHSe5-YNvdZ35c4Xbg4bx4MqmLfZ5v9h8B8bD8ey4L_EjzwnP8D_8Z7pWExirpPF72p1sYd1LUAU3z4DXUiAAAVPM', '2FF57F87FB0011E5AA223010B3A0F15C', '1460714399');
INSERT INTO `token` VALUES ('D261F7D3017B11E68B013010B3A0F15C', 'DrzzqJGMvn7kTO7sm_aLJUTDu9QxkMHVP2Ax4FNSu6Qm94JTqQlsUKoFhAFYo8M0s_Oj9P5jIIOY32C3J0wFgSD2aWlYMIwaTOON2DzimLAPP3SVObtn99qjaQRWJtwxCVShAHAQJG', '2FF57F87FB0011E5AA223010B3A0F15C', '1460554200');
INSERT INTO `token` VALUES ('D3AF0D24000611E6934D3010B3A0F15C', 'qbbizEJlb7YoD1JSo4SiGVpZ4qBECGm-NIM8UgtkNpVgyOqbNNsXOPKzKWJaZz77tvJ_DyDlQjZ-kkSA2XHIXDFqvHjuDGkOQnV9fk-q5o7CUEj_F9KUZp53TgthKB6aVQHdABAGBS', '1F2F0533F6ED11E58F303010B3A0F15C', '1460394000');
INSERT INTO `token` VALUES ('D3E90A60000611E6934D3010B3A0F15C', 'vP52v6EWQxEd7aRjPo2UDCMI-CWvOIYYY7y-BZQocCSH_nmbEnAdoDGVJYnj1zTTO5Kcsby28obx1O0n6Zoy6G9j7KatQeoIG-xzOab8etBTTihzCvZ0Y2URP995QmbrTXUaAJAMQN', '2FF57F87FB0011E5AA223010B3A0F15C', '1460394000');
INSERT INTO `token` VALUES ('D776DF7F022B11E6B50F3010B3A0F15C', 'RaV4_F-tSm9CYa96SAp1SxmKOqX2GSrBbiZ_9zlh4kIckduQ4hJdTfQTfaoB3hOvkDvmdx0jlqRYp4JKee78HEO5Wy9I7Yj-ydBNYen1bkngwAA-M7o1Pygtvvu2i7PoDOLeAEATEB', '2FF57F87FB0011E5AA223010B3A0F15C', '1460629800');
INSERT INTO `token` VALUES ('D8E1FABAFD7411E5B40A3010B3A0F15C', 'Aey09kyjbGT-QFW5rAxQ7yIOTi5j34kZ6Ek4G2sA0rfB7OmVa8HjNAYUbG2Tbrhlkav3B0PvGR29eHl3MGSnL3thUMPslSRjkoi9UqIiSn1Xr7gp6C236WpmFIp8u66xYUKiAEAUQG', '1F2F0533F6ED11E58F303010B3A0F15C', '1460111400');
INSERT INTO `token` VALUES ('D90280EE00B611E692FC3010B3A0F15C', 'eVW7sgZm5cl2XB65bSsv2toeRD_jCLOTfCvGObk31b_SZc6sdOcpWGs6jlJ9CVTyXpJ8frq6L2jJ1II1d3j0294FrXAoA4p2b02f-4U5nv9p2w75rtDdJvoEbGXJJoxLIXCbAHANCH', '2FF57F87FB0011E5AA223010B3A0F15C', '1460469600');
INSERT INTO `token` VALUES ('D914A37AFD7411E5B40A3010B3A0F15C', 'JBVFYPmgL0OkRhK-wZ_iCxtHK8O8iGlMLJeV7HKUxfs5iIjtovueUiqgX1vDjCuKUVzo5Xay3IUOBFWI8vFD4p5nWkSpOPjHzm1fFjurPBRepi7CI0_6AnYNijt7ieAlSDThAEAYJZ', '2FF57F87FB0011E5AA223010B3A0F15C', '1460111400');
INSERT INTO `token` VALUES ('D9F78C1E028311E6AB2E3010B3A0F15C', 'PS9vQJfCZvw2eEOaTbSw2rC4ZnO_HyeNOY7AXjxWX8hCJ-34pX9JHovsia0iTIZW_ne5zbfwH0JhXhvo0VU2LIYWyIaxJPXNtXeJ4Y0UVRfVw8RWRjgQ_G9H7VcZAX99MYPdACAZLC', '2FF57F87FB0011E5AA223010B3A0F15C', '1460667600');
INSERT INTO `token` VALUES ('DB4F4BA7045011E6B5103010B3A0F15C', '-BGJPTGTRwRNFXS_cgjN2X6dBwm_lkiecS22Qe_QFs9hPssJBBN_F3CxDAWdl66FtDr1rOtTSFfXCgsyxgg4u8tBTTPoYJRmVd1XeLOuBmc07PyHFaBqRw9kBAWHHzUEMIRbADAUQZ', '1F2F0533F6ED11E58F303010B3A0F15C', '1460865600');
INSERT INTO `token` VALUES ('DC8DB12802DB11E680403010B3A0F15C', 'H4EIWLij93PhxVvFw7iKIuPg1unfUyHSO8ov_zFJbD0bAbFWXi0EbWSrwwxcWsIHBPuBTl_UcnTBfnhv47BdDQo8Ns7OK6Vv6Xt05ek54xph2hvOkjAU9uce64KcoSo3GIWaAJAUWK', '2FF57F87FB0011E5AA223010B3A0F15C', '1460705400');
INSERT INTO `token` VALUES ('DDC0631904A811E6B5103010B3A0F15C', 'nSHDKA-Ra6hxfuIq95Z-Fo5urB9hywdfcD_tWI4Swl2D2o_zxf1iE-UPYG0sa49SonF5bhzVp0dkXR0OZzyFVsTvNwZ9pBwZMDCbB9357J42waQUMUb8CNnBS4PqR3i9FZKgACABVG', '1F2F0533F6ED11E58F303010B3A0F15C', '1460903400');
INSERT INTO `token` VALUES ('DFCA4916F81111E5B9926057182B3A68', 'R21H08bIpy_BC5f8LuhjJw-5pNhrN-pkpqhfT3Am7h0Gr0106VyHtiMXMV-gCXo8jvhM42mgeRcQGXQkVQE2ASCZFRBOqkhm6vdkdNecMUs7NZL1w961qInm23Z1WBemTTWjAEAEEH', '1F2F0533F6ED11E58F303010B3A0F15C', '1459519135');
INSERT INTO `token` VALUES ('E191D863FFAB11E5B0483010B3A0F15C', 'vePyAjvSCkyBrBiLAVQnjdI8n0CLMeYcZVHaRakSfMiFl-62MoKo57WPRG95Thmny3v9l7osvjs171jUBrihB9406CoQFAr8cDC36aql0GoITAeAIAYJQ', '1F2F0533F6ED11E58F303010B3A0F15C', '1460354939');
INSERT INTO `token` VALUES ('E3085C7C021611E6B0533010B3A0F15C', 'QWbvZK44slJRit5QyRngEJMsqF4RD9vnshuY7UTAjb44En-qOThc43PXjKEsMIxv8Xf6WhICqJu6Zym1oYvqJalyS7UOUqcH1Sk0XETY7Fql7IdVGTuWj1fVtKRQGKGIUXDaADAYBH', '1F2F0533F6ED11E58F303010B3A0F15C', '1460620800');
INSERT INTO `token` VALUES ('E45071DA03E311E6B5103010B3A0F15C', 'nYmQ9iprafleq8LEKbcPJkGkEWf3MzWRfcpxF4Tm3_wWdd-pZfERXfBVG0TzKJkTinAr2KFxzycGGizOkYP54oYxzjS0sutDHz2UwxCo_773bdFKkH1qor6T7YTxkfKhKNJaADADGC', '2FF57F87FB0011E5AA223010B3A0F15C', '1460818800');
INSERT INTO `token` VALUES ('E46BD79900A111E692FC3010B3A0F15C', 'N_xAmdojvqDeSFbJQmtoG6XUa4x6Jctvh6lYB9eT50w06KTAJtXoqfVTTj2gQsRKiLJMRGafrEIlU8LPnAb9ewUaaHM65O7_nDOTiZi2SSFZTvYami4Cizd0u6RHKNDrSBXhADAGPL', '1F2F0533F6ED11E58F303010B3A0F15C', '1460460600');
INSERT INTO `token` VALUES ('E48FBD45FD5F11E5B0023010B3A0F15C', 'pZXvqWIDdoIQvIPIIwzwuvytm8vUuUFTGiZXtJRi98AhuGnedA2tgLFDo4z0fX-zTF3M2CaEaaF5WBSRu75SgquUqPrYvrLLEysNtqcxR8STx0c6SKTka4gdAVO62LqaSWOiAEAFJB', '1F2F0533F6ED11E58F303010B3A0F15C', '1460102400');
INSERT INTO `token` VALUES ('E4BAC5FFFD5F11E5B0023010B3A0F15C', 'yIt2PFR4yNflDlE572ZwCo5nJOnL0CFI38tIgMsSUwZfNnDhtQ1MfF_vw3aMhdqBwPOESTGylzB5azuAsS3DHTRfu6aQJIrdxnZDHPjQwgLUZ-mp-Pnh2-VQ_Afta05sDJOjAHAOUZ', '2FF57F87FB0011E5AA223010B3A0F15C', '1460102400');
INSERT INTO `token` VALUES ('E597AB14026E11E6AB2E3010B3A0F15C', 'F0-QdpQ9lPWUKXTtpyIhXgo75v3FPRf21NjL8e5yBU93f-i5YqLHRseLNonPAXkP7qqTGl3F3L2u5FoWDGR-NEHNWJtiMRrsGpQ1Px8ZKLiDvIkIWdPWRsrSC_xe2zlVBLBdABAUEU', '2FF57F87FB0011E5AA223010B3A0F15C', '1460658600');
INSERT INTO `token` VALUES ('ECBC3443016911E6BA5D3010B3A0F15C', 'zxmhatMwgfg1yBtVz9BVVTF0D9y3Zud7GuL4jHxymYGDsigvMuEn1h5CvYr11N4BfT4o0Ex3_nOCAGjaN-wIWNq9I4uqo_svbp1jjIr2M6Fh6VCZVzVhfLzbprB3i8QaKAJgAAAVLU', '2FF57F87FB0011E5AA223010B3A0F15C', '1460546513');
INSERT INTO `token` VALUES ('ED544E58037611E6B5103010B3A0F15C', 'OAnPjRY6NOfyyT_0AnrXeJdpHtns7JOROVSOMjoZq5t9Co4lVKTPJYq0tgnwLNaLMF0qv3RRG6gVg9DEto9xkI5zjVZ_fvQM1qy9aG5NliZGKT2oolAdy7goRUuWqQqnJRQcAIABJI', '1F2F0533F6ED11E58F303010B3A0F15C', '1460772000');
INSERT INTO `token` VALUES ('ED8A05B8037611E6B5103010B3A0F15C', 'qnjqZ2GqXPa53Od_90D07eDieogjwhRzMwSUcL9gCsjcM7fjoTjqEOTFME15WnbsAr5hK7qF6ULPPcRsaHahEQMQMFix6R1dLDbilXqbJY-_xfYfTcS8MPD2BnyKVgVtPSWjADAWZD', '2FF57F87FB0011E5AA223010B3A0F15C', '1460772000');
INSERT INTO `token` VALUES ('F130A252025911E6AB2E3010B3A0F15C', '4IDBs7LB-ij5cwAlNrR1qbVtDcru2PzfbUQ5Nu70bXsaR8-jOnqILIZ2YJqvGG80f9XS8AO00R4byETzsceH1_3fCLa36zxXpl9TBDbt8Wr5GUjPgQXJilqGiObMZWmuIIHcAEAZZY', '1F2F0533F6ED11E58F303010B3A0F15C', '1460649600');
INSERT INTO `token` VALUES ('F3C5BA4C02B111E6AB2E3010B3A0F15C', 'fSx5F3uKQfI3wAFB79W8k-iCKgvSjMGhbgLmIqk5o1RjncL62q3ZzRzaCh0uZkouivyB2E2abKr_33KVQHqgNxWq1EFwtZjKfsBpTt6bAnAzMyNH8r8TAndX0OIiEnW1LDFeABAOGI', '1F2F0533F6ED11E58F303010B3A0F15C', '1460687400');
INSERT INTO `token` VALUES ('F4EFD8A0047E11E6B5103010B3A0F15C', 'UTy4klGxtt0yu90fM8DoMSv_wg3oPui8bomSdrX1aq5ocdxFYTft48_IWR2lgOlIV3O75TuMfdeMqWbSfAg2u3Krd-cIqlYzoOFKhz6atzhNPHSSEVl621r9AwOLkZ60PRBcAEAJLV', '2FF57F87FB0011E5AA223010B3A0F15C', '1460885400');
INSERT INTO `token` VALUES ('F9022B9FFD3511E58C223010B3A0F15C', 'JoggzQZKxHFWkXf2C6hSYtMDoWfxX7_o9cgHMm-b3jE_BL6_ZzrpWyEna20k0blNmFpEfqFKAmPa2F7et-ja38se5nZhaF3SB5jPz97mAQpNVez2r3vM6qfV1AelVfxKGMRgAGAAGL', '2FF57F87FB0011E5AA223010B3A0F15C', '1460084395');
INSERT INTO `token` VALUES ('F9EEE33F036111E6B5103010B3A0F15C', 'O5Q5mWlP9I99N55uW9bG6Mn4OK8QMtS16x_MFR9BvQbGjjQCUyLpGuBJitTnSri9y6h6ma3iabppTizpIZMBGvZ6241CNqalFv6H4eEEynEHhsCfe4BIjxBLj40EqfnvBRIbAAAIYP', '1F2F0533F6ED11E58F303010B3A0F15C', '1460763002');
INSERT INTO `token` VALUES ('FA267243036111E6B5103010B3A0F15C', '2fa5hizhOucUjEar9NTNCAWGWkBykjVOuNVB070z4BIIYkBiFCeAlGVBAqpI4j5pDAAzXV1QAezRQkW-mPGZDE9o_DLL887W1a-t8D9rFHkTCSdAHAMVL', '2FF57F87FB0011E5AA223010B3A0F15C', '1460763002');
INSERT INTO `token` VALUES ('FA67EE06FAD011E5AA6F3010B3A0F15C', 'U2j2pWH1IxHp1Z9Mg2fHvcmYZLZgUX-CfhkoWb5mT4sz_c7mAroQHDMbLmZYk2vGfbEwO3fYNwl841uUsrXOMLAx0HpHS_A7WyhOusKnSjgUcbcq6k7gTesUARD2bmp9SFQaAEAJWH', '1F2F0533F6ED11E58F303010B3A0F15C', '1459821116');
INSERT INTO `token` VALUES ('FB851F8803B911E6B5103010B3A0F15C', 'pk8PpKM3XIGtJd-eOCdsDcyE7emngHNOSkJ4Hkhv8FPV4rcvrHYdXSwU_m5N0_XPn7uL0fYtjkhV0sHoDEvDqgs3bUmgX64Hh6d9T_6F2EW_nvMxRYnO48w-6ij9ZP7cABAbADASTS', '1F2F0533F6ED11E58F303010B3A0F15C', '1460800800');
INSERT INTO `token` VALUES ('FC78BB41007711E6A47E3010B3A0F15C', 'G1O3VDkKPptLXKvT6bDFfEthtu9O-4SC73nKUK2OpUrz6BCQACy_0Bc1QvvnnEgICJtfGYaxLC9q_RK4KE07wU4hhwMhKTltRvoNGWc0zIkUDReABAMVJ', '1F2F0533F6ED11E58F303010B3A0F15C', '1460442601');
INSERT INTO `token` VALUES ('FCBDB381007711E6A47E3010B3A0F15C', 'yIt2PFR4yNflDlE572ZwCrLL50_XSyN3xb1V-U4Oyn5-mJnifTgX5W00wzws0j0vAaF5T6iEd45b0eDiuQMYG1ftGwdS0xRGuf0Y6e-MH_N8Tj_jtXc0WbEFapxDAPjkVTBfAFAECK', '2FF57F87FB0011E5AA223010B3A0F15C', '1460442602');
INSERT INTO `token` VALUES ('FEB09FB7FB0F11E588B13010B3A0F15C', '2J2sCf8ToNYasqXVPb6qdLDtho4r1dJxzukBbUN2u_-6eJR08FNZplb8b6XGol7xAry3u55p4PeOxePdO2TeoRtSThxYhgWiEC1VhPs27MQ8YObvfqfU5RkJ26k8wRQRNDHbADAQGK', '1F2F0533F6ED11E58F303010B3A0F15C', '1459848182');

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
