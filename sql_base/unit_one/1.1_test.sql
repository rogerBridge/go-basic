CREATE TABLE addressbook (
                             regist_no int not null ,
                             primary key (regist_no) ,
                             name varchar(128) not null ,
                             address varchar(256) not null ,
                             tel_no char(10),
                             mail_address char(20)
);

-- Add more columns after create table
ALTER TABLE addressbook ADD COLUMN postal_code char(8) not null;

-- DROP TABLE addressbook;