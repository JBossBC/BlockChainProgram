DROP database  if EXISTS users;

create database users;

use users;


drop table if EXISTS users;

create table users(
            id  bigint primary key AUTO_INCREMENT,
			username varchar(50),
            password varchar(256) not NULL,
			create_time TIMESTAMP not null DEFAULT(now()),
			delete_time TIMESTAMP,
            update_time TIMESTAMP not null DEFAULT(now())
)

