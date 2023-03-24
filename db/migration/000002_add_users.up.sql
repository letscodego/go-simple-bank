CREATE TABLE `users` (
  `username` varchar(255) PRIMARY KEY,
  `hashed_password` varchar(255) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `password_changed_at` timestamp NOT NULL DEFAULT "1970-01-01 00:00:01",
  `created_at` timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE `accounts` ADD FOREIGN KEY (`owner`) REFERENCES `users` (`username`);

ALTER TABLE `accounts` ADD CONSTRAINT `accounts_owner_currency` UNIQUE KEY (`owner`,`currency`);
