ALTER TABLE IF EXISTS `accounts` DROP CONSTRAINT IF EXISTS `accounts_owner_currency`;

ALTER TABLE IF EXISTS `accounts` DROP FOREIGN KEY `accounts_ibfk_1`;

DROP TABLE IF EXISTS `users`;


