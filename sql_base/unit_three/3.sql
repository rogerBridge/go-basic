DROP TABLE product_ins;
CREATE TABLE product_ins
(
product_id char(4) not null,
product_name varchar(100) not null,
product_type varchar(32) not null,
sale_price integer default 0,
purchase_price integer ,
regist_date date ,
primary key (product_id)
);

INSERT INTO product_ins (product_id, product_name, product_type, sale_price, purchase_price, regist_date) VALUES ('0001', 'Tshirt', 'clothes', 1000, 500, '2020-06-18');
INSERT INTO product_ins (product_id, product_name, product_type, sale_price, purchase_price, regist_date) VALUES ('0002', 'Pen', 'tool', 100, 50, '2020-06-18');
INSERT INTO product_ins (product_id, product_name, product_type, sale_price, purchase_price, regist_date) VALUES ('0003', 'Shoes', 'tool', 200, 100, '2020-06-18');

DROP TABLE product_ins_copy;
/* 对一个表格做备份操作 */
CREATE TABLE product_ins_copy (
                                  product_id char(4) not null,
                                  product_name varchar(100) not null,
                                  product_type varchar(32) not null,
                                  sale_price integer default 0,
                                  purchase_price integer ,
                                  regist_date date ,
                                  primary key (product_id)
);
INSERT INTO product_ins_copy (product_id, product_name, product_type, sale_price, purchase_price, regist_date) SELECT product_id, product_name, product_type, sale_price, purchase_price, regist_date FROM product_ins;
-- UPDATE product_ins_copy SET regist_date = '2009-10-10';

-- DELETE FROM product_ins_copy;

-- DELETE FROM product_ins_copy WHERE regist_date>'2020-06-18';

-- TRUNCATE product_ins_copy; -- higher speed compare to DELETE

-- UPDATE product_ins_copy SET sale_price=sale_price*1.2 WHERE product_type='clothes';

-- UPDATE product_ins_copy SET sale_price=sale_price*10, purchase_price=purchase_price/2 WHERE product_type='tool';

-- 事务, 什么是事务呢? 一个处理单元里面的内容, 要不全部处理完, 要不全部没有处理
-- BEGIN TRANSACTION;
-- UPDATE product_ins_copy SET sale_price=sale_price+10 WHERE product_type='tool';
-- UPDATE product_ins_copy SET sale_price=sale_price-10 WHERE product_type='clothes';
-- COMMIT;
-- ROLLBACK

-- BEGIN TRANSACTION;
-- INSERT INTO product VALUES ('0005', 'T恤', '衣服', 1000, 500, '2009-09-20');
-- INSERT INTO product VALUES ('0006', '打孔器', '办公用品', 500, 320, '2009-09-11');
-- INSERT INTO product VALUES ('0007', '运动T恤', '衣服', 4000, 2800, NULL);
-- COMMIT;