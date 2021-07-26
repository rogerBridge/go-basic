SELECT product_name, regist_date FROM product WHERE regist_date>'2009-04-28';

SELECT product_name, product_type, (sale_price*0.9-purchase_price) AS profit FROM product WHERE sale_price*0.9-purchase_price>100 AND (product_type = 'workTools' OR product_type = 'clothes');