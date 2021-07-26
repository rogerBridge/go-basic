DROP TABLE product;
CREATE TABLE product(
product_id CHAR(4) NOT NULL,
product_name VARCHAR(100) NOT NULL,
product_type VARCHAR(32) NOT NULL,
sale_price INTEGER ,
purchase_price INTEGER ,
regist_date DATE,
PRIMARY KEY(product_id)
);

BEGIN TRANSACTION;
INSERT INTO product VALUES ('0001', 'Tshirt', 'clothes', 1000, 500, '2009-09-20');
INSERT INTO product VALUES ('0002', 'Ticker', 'tool', 500, 320, '2009-09-11');
INSERT INTO product VALUES ('0003', 'SportTshirt', 'clothes', 4000, 2800, NULL);
INSERT INTO product VALUES ('0004', 'knife', 'kitchen', 3000, 2800, '2009-09-20');
INSERT INTO product VALUES ('0005', '高压锅', 'kitchen', 6800, 5000, '2009-01-15');
INSERT INTO product VALUES ('0006', '叉子', 'kitchen', 500, NULL, '2009-09-20');
INSERT INTO product VALUES ('0007', '菜板子', 'kitchen', 880, 790, '2008-04-28');
INSERT INTO product VALUES ('0008', '圆珠笔', 'tool', 100, NULL, '2009-11-11');
COMMIT;
