# User-Service-gRPC
# This is the small gRPC Service which handles user data 

# Steps to build and Run GRPC Server:
1. cd ./User-Service-gRPC
2. Create .env file, insert below data in the file:
    DB_USER=<user>
    DB_PASSWORD=<password>
    DB_NAME=<db_name>
    DB_HOST=db
    DB_PORT=3306

    ENVIRONMENT=<env>
    APP_PORT=8000
    GRPC_PORT=9000

3. run command to build and run grpc server and mysql container: docker compose up --build -d

# Steps to build Client:
1. cd ./User-Service-Client
2. Create .env file, insert below data in the file:
    ENVIRONMENT=<env>
    APP_PORT=8000
    USER_SERVICE_URL= grpc-server:9000 
    # here grpc-server is the container name for the grpc server 
3. run command to build and run grpc client container: docker compose up --build -d

# Endpoints for grpc server:
  GET Request: localhost:8000/api
  1. /user/{{user_id}} : to fetch user details by id
  2. /user?id={{user_id}}&id={{user_id}} : to fetch multiple user details by given a list of ids.
  3. /user/filter?country={{country}}&marital_status={{status}} : to fetch list of users based on filter of country and marital status

# Sample Data can be used 
INSERT INTO users (id, name, address, city, state, country, pincode, phone_number, marital_status, height)
VALUES
('11111111-1111-1111-1111-111111111111', 'John Doe', '123 Elm St', 'Springfield', 'IL', 'INDIA', '62701', '555-1234', 'single', 180.5),
('1b2d0a52-1a7f-4c1e-9c24-5d5cb7c8d1b0', 'Alice Smith', '456 Maple Ave', 'Metropolis', 'NY', 'USA', '10101', '555-123-4567', 'married', 5.5),
('22222222-2222-2222-2222-222222222222', 'Jane Smith', '456 Oak St', 'Lincoln', 'NE', 'USA', '68508', '555-5678', 'married', 165),
('33333333-3333-3333-3333-333333333333', 'Alice Johnson', '789 Pine St', 'Columbus', 'OH', 'USA', '43215', '555-8765', 'single', 170.2),
('44444444-4444-4444-4444-444444444444', 'Bob Brown', '321 Maple St', 'Madison', 'WI', 'INDIA', '53703', '555-4321', 'single', 175.3),
('55555555-5555-5555-5555-555555555555', 'Eve White', '654 Cedar St', 'Chicago', 'IL', 'INDIA', '60614', '555-6789', 'single', 160.4);
