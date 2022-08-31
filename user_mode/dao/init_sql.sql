DROP database  if EXISTS users;

create database users;

use users;


drop table if EXISTS users;

create table users(
            id varchar(20) primary key,
			username varchar(50),
			password varchar(30) not NULL,
			create_time TIMESTAMP not null DEFAULT(now()),
			delete_time TIMESTAMP
)

