# User-Service-gRPC
# This is the small gRPC Service which handles user data 

Steps to build and Run GRPC Server:
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

Steps to build Client:
1. cd ./User-Service-Client
2. Create .env file, insert below data in the file:
    ENVIRONMENT=<env>
    APP_PORT=8000
    USER_SERVICE_URL= grpc-server:9000 
    # here grpc-server is the container name for the grpc server 
3. run command to build and run grpc client container: docker compose up --build -d

Endpoints for grpc server:
- GET Request: localhost:8000/api
    - /user/{{user_id}} : to fetch user details by id
    - /user?id={{user_id}}&id={{user_id}} : to fetch multiple user details by given a list of ids.
    - /user/filter?country={{country}}&marital_status={{status}} : to fetch list of users based on filter of country and marital status

