CREATE SCHEMA `kumparan_test_db` ;

CREATE TABLE `kumparan_test_db`.`articles` (
       `id` INT NOT NULL AUTO_INCREMENT,
       `author` VARCHAR(100) NOT NULL,
       `title` VARCHAR(225) NOT NULL,
       `body` TEXT NOT NULL,
       `created_at` TIMESTAMP NOT NULL,
       PRIMARY KEY (`id`)
);
