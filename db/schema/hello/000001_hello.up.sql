CREATE TABLE `message` (
  `id` varchar(50) PRIMARY KEY,
  `sender` varchar(50) NOT NULL,
  `info` varchar(256),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);