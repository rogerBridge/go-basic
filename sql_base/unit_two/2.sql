-- create table
CREATE TABLE chars
(chr CHAR(3) NOT NULL,
PRIMARY KEY (chr));

-- insert data
BEGIN TRANSACTION ;
INSERT INTO chars VALUES ('1') ;
INSERT INTO chars VALUES ('2') ;
INSERT INTO chars VALUES ('3') ;
INSERT INTO chars VALUES ('10') ;
INSERT INTO chars VALUES ('11') ;
INSERT INTO chars VALUES  ('222');
COMMIT;
