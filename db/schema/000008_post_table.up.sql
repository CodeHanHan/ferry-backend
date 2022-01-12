CREATE TABLE `sys_post` (
  `post_id` varchar(255) PRIMARY KEY NOT NULL,
  `post_name` varchar(128) DEFAULT NULL,
  `post_code` varchar(128) DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `status` int DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT '2000-01-01 00:00:00',
  `delete_time` timestamp NULL DEFAULT '2000-01-01 00:00:00'
);