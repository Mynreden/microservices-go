# Microservices CRUD Application

This application is a simple microservices-based system for managing users and posts. It consists of two microservices for handling CRUD operations on users and posts, and a gateway controller for routing requests.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Endpoints](#endpoints)
    - [User Service](#user-service)
    - [Post Service](#post-service)
- [Configuration](#configuration)
- [Setup](#setup)
- [Running the Application](#running-the-application)
- [License](#license)

## Overview

The application is composed of two main microservices:
1. **User Service**: Handles CRUD operations for users.
2. **Post Service**: Handles CRUD operations for posts.

A gateway controller is used to route requests to the appropriate microservice.

## Architecture

- **User Service**: Provides endpoints to create, retrieve, and delete users.
- **Post Service**: Provides endpoints to create, retrieve, and delete posts.
- **Gateway Controller**: Routes incoming requests to the appropriate microservice based on the URL path.

## Endpoints

### User Service

- **Create User**
    - **Endpoint**: `/api/users/create`
    - **Method**: POST
    - **Body**: JSON `{ "username": "user", "email": "user@example.com", "password": "pass" }`

- **Delete User**
    - **Endpoint**: `/api/users/delete/{id}`
    - **Method**: DELETE
    - **Params**: `{ "id": "user_id" }`

- **Get User**
    - **Endpoint**: `/api/users/{id}`
    - **Method**: GET
    - **Params**: `{ "id": "user_id" }`

### Post Service

- **Create Post**
    - **Endpoint**: `/api/posts/create`
    - **Method**: POST
    - **Body**: JSON `{ "title": "Post Title", "content": "Post content", "userId": "user_id" }`

- **Delete Post**
    - **Endpoint**: `/api/posts/delete/{id}`
    - **Method**: DELETE
    - **Params**: `{ "id": "post_id" }`

- **Get Post**
    - **Endpoint**: `/api/posts/{id}`
    - **Method**: GET
    - **Params**: `{ "id": "post_id" }`

- **Get User Posts**
    - **Endpoint**: `/api/posts/user/{id}`
    - **Method**: GET
    - **Params**: `{ "id": "user_id" }`

## Configuration

The application configuration is managed through environment variables:

- `ADDR`: The address for the gateway controller (default: `:8080`).
- `USER_ADDR`: The address for the User Service (default: `localhost:50051`).
- `POST_ADDR`: The address for the Post Service (default: `localhost:50052`).
- `DSN`: The Data Source Name for the PostgreSQL database.
- `STATIC_DIR`: Directory for static files (default: `/static`).

## Setup

1. **Clone the repository**:
   ```sh
   git clone https://github.com/yourusername/microservices-go.git
   cd microservices-go
2. **Install dependencies**:

    ```sh
    go mod tidy
    ```
3. **Run database migrations**:

    ```sh
    go run scripts/migrate.go
    ```
## Running the Application
1. Start the User Service:
    ```sh
    go run cmd/userservice/main.go
    ```
   
2. Start the Post Service:

    ```sh
    go run cmd/postservice/main.go
    ```
3. Start the Gateway Controller:

```sh
go run cmd/gateway/main.go
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.

Feel free to adjust the paths and details according to your project's str