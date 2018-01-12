CREATE DATABASE `myboard` CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `myboard`;

CREATE TABLE `board` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `text` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `event_stream` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `event` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
