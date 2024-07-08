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

