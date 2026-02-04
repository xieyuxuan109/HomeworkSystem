/*
SQLyog Community v13.2.1 (64 bit)
MySQL - 8.0.43 : Database - homework
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`homework` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `homework`;

/*Table structure for table `homeworks` */

DROP TABLE IF EXISTS `homeworks`;

CREATE TABLE `homeworks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '作业id',
  `title` varchar(200) DEFAULT NULL COMMENT '作业标题',
  `description` text COMMENT '作业描述/作业要求',
  `department` enum('backend','frontend','sre','product','design','android','ios') DEFAULT NULL COMMENT '所属部门',
  `creator_id` bigint unsigned DEFAULT NULL COMMENT '发布者id',
  `deadline` timestamp NULL DEFAULT NULL COMMENT '截止时间(Asia/Shanghai)',
  `allow_late` tinyint(1) DEFAULT NULL COMMENT '是否允许补交',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `creator_id` (`creator_id`),
  KEY `id_title` (`title`),
  CONSTRAINT `homeworks_ibfk_1` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



DROP TABLE IF EXISTS `submissions`;

CREATE TABLE `submissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `homework_id` bigint unsigned DEFAULT NULL COMMENT '作业id',
  `student_id` bigint unsigned DEFAULT NULL COMMENT '学员id',
  `content` text COMMENT '提交内容/链接',
  `file_url` varchar(500) DEFAULT NULL COMMENT '附件地址',
  `is_late` tinyint(1) DEFAULT NULL COMMENT '是否迟交',
  `score` int DEFAULT NULL COMMENT '分数',
  `comment` text COMMENT '老师评语',
  `is_excellent` tinyint(1) DEFAULT NULL COMMENT '是否是优秀作业',
  `reviewer_id` bigint DEFAULT NULL,
  `reviewed_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `submission_count` int DEFAULT '0' COMMENT '提交人数',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `student_id` (`student_id`),
  KEY `reviewer_id` (`reviewer_id`),
  KEY `fk_submissions_homework` (`homework_id`),
  CONSTRAINT `fk_submissions_homework` FOREIGN KEY (`homework_id`) REFERENCES `homeworks` (`id`),
  CONSTRAINT `submissions_ibfk_1` FOREIGN KEY (`homework_id`) REFERENCES `homeworks` (`id`) ON DELETE CASCADE,
  CONSTRAINT `submissions_ibfk_2` FOREIGN KEY (`student_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `username` varchar(50) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '昵称',
  `role` enum('student','admin') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'student' COMMENT '角色',
  `department` enum('backend','frontend','sre','product','design','android','ios') DEFAULT NULL COMMENT '所属部门',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '""' COMMENT '邮箱',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `id_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;