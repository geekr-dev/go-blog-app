CREATE DATABASE IF NOT EXISTS blog DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;

CREATE TABLE `articles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '文章摘要',
  `feature_image` varchar(200) DEFAULT '' COMMENT '封面图片',
  `content` text COMMENT '文章内容',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '作者',
  `updated_by` varchar(100) DEFAULT '' COMMENT '编辑',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态: 0、未发布，1、已发布',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章表'

CREATE TABLE `tags` (
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(100) DEFAULT '' COMMENT '标签名称',
	`state` tinyint(3) unsigned DEFAULT 1 COMMENT '状态: 0、已禁用，1、已启用',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
	`created_by` VARCHAR(100) DEFAULT '' COMMENT '作者',
	`updated_by` VARCHAR(100) DEFAULT '' COMMENT '编辑',
	`deleted_at` TIMESTAMP DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签表';

CREATE TABLE `article_tags` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int unsigned NOT NULL COMMENT '文章ID',
  `tag_id` int unsigned NOT NULL COMMENT '标签ID',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` TIMESTAMP DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章标签关联表'

CREATE TABLE `auth`(
  `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` VARCHAR(20) DEFAULT '' COMMENT 'APP KEY',
  `app_secret` VARCHAR(50) DEFAULT '' COMMENT 'APP SECRET',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` TIMESTAMP DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户认证表'