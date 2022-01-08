CREATE TABLE `role` (
  `role_id` varchar(255) PRIMARY KEY NOT NULL,
  `role_name` varchar(128) DEFAULT NULL unique,
  `status` int DEFAULT NULL,
  `role_key` varchar(128) DEFAULT NULL,
  `role_sort` int DEFAULT NULL,
  `flag` varchar(128) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `admin` int DEFAULT 0,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT '2000-01-01 00:00:00',
  `delete_time` timestamp NULL DEFAULT '2000-01-01 00:00:00'
);