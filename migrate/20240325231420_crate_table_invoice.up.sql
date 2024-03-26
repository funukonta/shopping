CREATE TABLE IF NOT EXISTS invoice (
    id_invoice SERIAL PRIMARY KEY,
    id_customer INT NOT NULL,
    status VARCHAR(50) DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     CONSTRAINT fk_customer_id FOREIGN KEY (id_customer) REFERENCES customers(id)
);

CREATE TABLE IF NOT EXISTS invoice_detail (
    id_invoice_detail SERIAL PRIMARY KEY,
    id_invoice INT NOT NULL,
    id_product INT NOT NULL,
    qty INT,
    sub_total DECIMAL(10,2),
    CONSTRAINT fk_invoice_id FOREIGN KEY (id_invoice) REFERENCES invoice(id_invoice),
    CONSTRAINT fk_product_id FOREIGN KEY (id_product) REFERENCES products(id_product)
);
