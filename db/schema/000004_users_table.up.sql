CREATE TABLE `user` (
  `id` varchar(100) PRIMARY KEY,
  `username` varchar(255) UNIQUE NOT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT 'common',
  `dept` varchar(255) DEFAULT NULL,
  `post` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `create_by` varchar(255) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp DEFAULT '2000-01-01 00:00:00'
);