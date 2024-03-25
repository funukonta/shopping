CREATE TABLE IF NOT EXISTS categories (
    id_category VARCHAR(15) PRIMARY KEY,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO categories (id_category,name) values ('CLTH','Clothes'),('FOOD','Foods'),('ETNC','Electronics');

CREATE TABLE IF NOT EXISTS products (
    id_product SERIAL PRIMARY KEY,
    name VARCHAR(255),
    price NUMERIC(10, 2),
    description TEXT,
    id_category VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category_id FOREIGN KEY (id_category) REFERENCES categories(id_category)
);

INSERT INTO products (name, price, description, id_category)
VALUES  ('T-Shirt', 59900.00, 'Cotton T-Shirt', 'CLTH'),
        ('Apples', 11900.00, 'Fresh red apples', 'FOOD'),
        ('Smartphone', 5999000.00, 'Latest model smartphone', 'ETNC');
