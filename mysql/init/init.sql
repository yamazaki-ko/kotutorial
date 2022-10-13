DROP DATABASE IF EXISTS dev_db;

CREATE DATABASE dev_db;
use dev_db;

CREATE TABLE `users` (
  `id` bigint unsigned not null auto_increment,
  `user_name` varchar(255) not null,
  `login_id` varchar(255) not null,
  `password` varchar(255) not null,
  `created_at` datetime not null default current_timestamp,
  `updated_at` datetime not null default current_timestamp on update current_timestamp,
  primary key (id)
);

INSERT INTO `users` (`user_name`,`login_id`,`password`) 
VALUES ('デフォルトユーザ','default','password');