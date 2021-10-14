#安装依赖
```go
go mod tidy
go mod install
```
# 迁移数据库结构
model的定义文档：
https://beego.me/docs/mvc/model/models.md

```go
//一定要编译！！！反正我不编译直接执行go run main orm syncdb 一直报错
go build main.go 
./main orm -h 查看帮助
./main orm orm syncdb -v
```
>使用 -force=1 可以 drop table 后再建表 
> 使用 -v 可以查看执行的 sql 语句

为了防止在orm迁移这一各种耽误时间，直接给出数据结构把
```sql

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

 Date: 14/10/2021 17:42:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

SET FOREIGN_KEY_CHECKS = 1;


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

 Date: 14/10/2021 17:43:19
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
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

SET FOREIGN_KEY_CHECKS = 1;


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

 Date: 14/10/2021 17:43:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

SET FOREIGN_KEY_CHECKS = 1;

```