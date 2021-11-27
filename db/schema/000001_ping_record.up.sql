CREATE TABLE `ping_record` (
  `ping_id` varchar(100) PRIMARY KEY,
  `message` varchar(255),
  `reply` varchar(255),
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
