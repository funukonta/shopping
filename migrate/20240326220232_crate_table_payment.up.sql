CREATE TABLE IF NOT EXISTS payment (
    id_payment SERIAL PRIMARY KEY,
    id_invoice int not null,
    payment_amount DECIMAL(10,2),
    payment_method INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_invoice_id FOREIGN KEY (id_invoice) REFERENCES invoice(id_invoice)
);

CREATE TABLE IF NOT EXISTS payment_method (
    id_payment_method SERIAL PRIMARY KEY,
    name VARCHAR(255)
);
