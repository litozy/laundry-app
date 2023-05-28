-- buat template
INSERT INTO Customers (customer_name, phone_number)
VALUES ('Jessica', '0812654987');

INSERT INTO Transactions (customer_id, service_id, transaction_date, received_by, quantity, unit, price, total_price)
VALUES (1, 1, '2022-08-18', 'Mirna', 5, 'KG', 7000.00, 35000.00),
       (1, 2, '2022-08-18', 'Mirna', 1, 'Buah', 50000.00, 50000.00),
       (1, 3, '2022-08-18', 'Mirna', 2, 'Buah', 25000.00, 50000.00);

INSERT INTO Services (service_name, service_price)
VALUES ('Cuci + Setrika', 35000), ('Laundry Bedcover', 50000), ('Laundry Boneka', 25000);

--kode yang ada di programnya
INSERT INTO Transactions (customer_id, service_id, transaction_date_in, transaction_date_done, received_by, quantity, unit, price, total_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

INSERT INTO Customers (customer_name, phone_number) VALUES ($1, $2);

DELETE FROM Customers WHERE customer_id = $1;

UPDATE Customers SET customer_name=$1, phone_number=$2 WHERE customer_id=$3;