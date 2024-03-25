CREATE TABLE IF NOT EXISTS shopping_cart (
    id_cart SERIAL PRIMARY KEY,
    id_customer INT NOT NULL,
    CONSTRAINT fk_customer_id FOREIGN KEY (id_customer) REFERENCES customers(id)
);

-- evan
INSERT INTO shopping_cart (id_customer) values (1);

CREATE TABLE IF NOT EXISTS shopping_cart_detail (
    id_cart INT,
    id_product INT,
    qty INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_cart_id FOREIGN KEY (id_cart) REFERENCES shopping_cart(id_cart),
    CONSTRAINT fk_product_id FOREIGN KEY (id_product) REFERENCES products(id_product)
);

-- t-shirt and phone
INSERT INTO shopping_cart_detail (id_product,qty) values (1,3),(3,1);