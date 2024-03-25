CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    address TEXT,
    phone VARCHAR(22),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO customers (name,email,password,address,phone) values ('Evan Roy','evanroy36@gmail.com','funukonta','Sidoarjo','089661108834');