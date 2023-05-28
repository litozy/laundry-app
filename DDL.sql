CREATE DATABASE eniglaundry;

CREATE TABLE Customers (
  customer_id SERIAL PRIMARY KEY,
  customer_name VARCHAR(255) NOT NULL,
  phone_number VARCHAR(20) NOT NULL
);

CREATE TABLE Services (
  service_id SERIAL PRIMARY KEY,
  service_name VARCHAR(255) NOT NULL
  service_price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE Transactions (
  transaction_id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL,
  service_id INT NOT NULL,
  transaction_date_in DATE NOT NULL,
  transaction_date_done DATE NOT NULL,
  received_by VARCHAR(255) NOT NULL,
  quantity INT NOT NULL,
  unit VARCHAR(10) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  total_price DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY (customer_id) REFERENCES Customers (customer_id),
  FOREIGN KEY (service_id) REFERENCES Services (service_id)
);