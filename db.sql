CREATE DATABASE shorten_link;

USE shorten_link;

CREATE TABLE shortlinks(
   ID int auto_increment PRIMARY KEY ,
   shortlink varchar(7) not null,
   link varchar(255) not null,
   created_at timestamp default current_timestamp                             null,
   updated_at timestamp default current_timestamp on update current_timestamp null,
   expired_at timestamp default current_timestamp null,

   index(shortlink)
);
