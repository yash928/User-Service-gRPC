CREATE DATABASE IF NOT EXISTS userdb;

USE userdb

CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY,           
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255),         
    city VARCHAR(100),         
    state VARCHAR(100),         
    country VARCHAR(100),         
    pincode VARCHAR(20),          
    phone_number VARCHAR(20),          
    marital_status VARCHAR(20),        
    height FLOAT                  
);

INSERT INTO users (id, name, address, city, state, country, pincode, phone_number, marital_status, height)
VALUES
('11111111-1111-1111-1111-111111111111', 'John Doe', '123 Elm St', 'Springfield', 'IL', 'INDIA', '62701', '555-1234', 'single', 180.5),
('1b2d0a52-1a7f-4c1e-9c24-5d5cb7c8d1b0', 'Alice Smith', '456 Maple Ave', 'Metropolis', 'NY', 'USA', '10101', '555-123-4567', 'married', 5.5),
('22222222-2222-2222-2222-222222222222', 'Jane Smith', '456 Oak St', 'Lincoln', 'NE', 'USA', '68508', '555-5678', 'married', 165),
('33333333-3333-3333-3333-333333333333', 'Alice Johnson', '789 Pine St', 'Columbus', 'OH', 'USA', '43215', '555-8765', 'single', 170.2),
('44444444-4444-4444-4444-444444444444', 'Bob Brown', '321 Maple St', 'Madison', 'WI', 'INDIA', '53703', '555-4321', 'single', 175.3),
('55555555-5555-5555-5555-555555555555', 'Eve White', '654 Cedar St', 'Chicago', 'IL', 'INDIA', '60614', '555-6789', 'single', 160.4);