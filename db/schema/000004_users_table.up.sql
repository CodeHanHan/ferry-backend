CREATE TABLE `users_table` (
  `id` varchar(100) PRIMARY KEY,
  `user_name` varchar(255),
  `password` varchar(255),
  `email` varchar(255),
  `role` varchar(255),
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);   