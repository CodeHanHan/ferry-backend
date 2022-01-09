CREATE TABLE `dept` (
    `dept_id` varchar(255) PRIMARY KEY NOT NULL,
    `parent_id` varchar(255)DEFAULT NULL,
    `dept_path` varchar(255) DEFAULT NULL,
    `dept_name` varchar(128) DEFAULT NULL,
    `dept_sort` int DEFAULT NULL,
    `leader` varchar(255) DEFAULT NULL,
    `phone` varchar(11) DEFAULT NULL,
    `email` varchar(64) DEFAULT NULL,
    `status` int DEFAULT NULL,
    `create_by` varchar(64) DEFAULT NULL,
    `update_by` varchar(64) DEFAULT NULL,
    `create_time` timestamp NULL DEFAULT  CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT  '2000-01-01 00:00:00',
    `delete_time` timestamp NULL DEFAULT  '2000-01-01 00:00:00'
);