/*
 Navicat Premium Data Transfer

 Source Server         : localhost_root
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : zset_api

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 20/10/2021 16:19:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id_id` int NOT NULL COMMENT '用户外键',
  `topic_id_id` int NOT NULL COMMENT '话题外键',
  `content` longtext NOT NULL,
  `is_active` int NOT NULL DEFAULT '1' COMMENT '1草稿箱，0非草稿箱',
  `is_delete` int NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `title` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=220 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of post
-- ----------------------------
BEGIN;
INSERT INTO `post` VALUES (1, 1, 1, 'python 续写', 1, 0, '2021-10-20 15:31:40', '2021-10-20 15:31:42', 'python+linux');
INSERT INTO `post` VALUES (2, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (3, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (4, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (5, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (6, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (7, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (8, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (9, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (10, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (11, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (12, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (13, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (14, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (15, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (16, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (17, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (18, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (19, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (20, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (21, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (22, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (23, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (24, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (25, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (26, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (27, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (28, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (29, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (30, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (31, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (32, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (33, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (34, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (35, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (36, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (37, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (38, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (39, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (40, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (41, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (42, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (43, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (44, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (45, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (46, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (47, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (48, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (49, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (50, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (51, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (52, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (53, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (54, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (55, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (56, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (57, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (58, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (59, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (60, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (61, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (62, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (63, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (64, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (65, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (66, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (67, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (68, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (69, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (70, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (71, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (72, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (73, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (74, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (75, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (76, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (77, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (78, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (79, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (80, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (81, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (82, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (83, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (84, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (85, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (86, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (87, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (88, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (89, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (90, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (91, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (92, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (93, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (94, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (95, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (96, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (97, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (98, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (99, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (100, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (101, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (102, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (103, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (104, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (105, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (106, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (107, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (108, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (109, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (110, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (111, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (112, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (113, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (114, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (115, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (116, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (117, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (118, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (119, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (120, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (121, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (122, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (123, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (124, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (125, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (126, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (127, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (128, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (129, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (130, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (131, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (132, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (133, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (134, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (135, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (136, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (137, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (138, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (139, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (140, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (141, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (142, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (143, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (144, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (145, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (146, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (147, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (148, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (149, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (150, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (151, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (152, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (153, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (154, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (155, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (156, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (157, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (158, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (159, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (160, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (161, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (162, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (163, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (164, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (165, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (166, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (167, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (168, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (169, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (170, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (171, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (172, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (173, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (174, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (175, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (176, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (177, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (178, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (179, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (180, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (181, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (182, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (183, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (184, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (185, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (186, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (187, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (188, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (189, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (190, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (191, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (192, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (193, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (194, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (195, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (196, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (197, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (198, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (199, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (200, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (201, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (202, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (203, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (204, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (205, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (206, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (207, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (208, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (209, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (210, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (211, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (212, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (213, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (214, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (215, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (216, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (217, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (218, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
INSERT INTO `post` VALUES (219, 1, 2, 'xue LINES', 1, 0, NULL, NULL, '我爱学习');
COMMIT;

-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic` (
  `id` int NOT NULL AUTO_INCREMENT,
  `index` int NOT NULL DEFAULT '1' COMMENT '排列顺序',
  `is_active` int NOT NULL DEFAULT '1' COMMENT '1上架，0下架',
  `is_delete` int NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of topic
-- ----------------------------
BEGIN;
INSERT INTO `topic` VALUES (1, 1, 1, 0, NULL, NULL, 'asdasd');
INSERT INTO `topic` VALUES (2, 2, 1, 0, NULL, NULL, '11sadasdsa');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `age` int DEFAULT NULL COMMENT '年龄',
  `gender` int DEFAULT NULL COMMENT '1:男,2:女,3:未知',
  `phone` bigint DEFAULT NULL COMMENT '电话号码',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `is_active` int NOT NULL DEFAULT '1' COMMENT '1启用，0停用',
  `is_delete` int NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  `profile` varchar(255) DEFAULT NULL COMMENT '头像oss',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'panda', 'd41d8cd98f00b204e9800998ecf8427e', 0, 0, 0, '', 0, 0, '', '2021-10-19 04:07:51', '2021-10-19 04:07:51');
INSERT INTO `user` VALUES (2, 'python', 'd41d8cd98f00b204e9800998ecf8427e', 111, 11, 0, '18778331181@sina.cn22', 0, 0, '12211212', '2021-10-19 04:08:06', '2021-10-19 04:44:51');
INSERT INTO `user` VALUES (3, 'python2', 'd41d8cd98f00b204e9800998ecf8427e', 0, 0, 0, '', 0, 0, '', '2021-10-19 04:41:27', '2021-10-19 04:41:27');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
