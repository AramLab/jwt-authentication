# Go JWT Authentication with Gin

This project demonstrates a simple implementation of JWT (JSON Web Token) authentication in Go using the Gin Gonic framework. The application is designed to handle user authentication and token management, providing a secure way to protect your API endpoints.

## Features

- **JWT Authentication:** Secure user authentication using JSON Web Tokens.
- **Access and Refresh Tokens:** Generate and validate access and refresh tokens.
- **Token Expiry and Refresh:** Manage token expiration and refreshing mechanisms.
- **User Management:** Basic user registration and login functionality.
- **Middleware Protection:** Protect API routes using JWT authentication middleware.

## Prerequisites

- Go 1.18 or later
- Gin Gonic
- MongoDB

## Setup

1. **Clone the Repository:**

```bash
git clone https://github.com/AramLab/go-jwt-authentication-with-gin.git
```

2. **Navigate to the Project Directory:**

```bash
cd go-jwt-authentication-with-gin
```

3. **Install Dependencies:**

```bash
go mod tidy
```

4. **Configure Environment Variables:**

Create a `.env` file in the project root and set the following variables:

```bash
PORT=your_port
SECRET_KEY=your_jwt_secret_key
MONGODB_URL=your_database_url
```

5. **Run the server:**

```bash
go run main.go
```

## API Endpoints

- POST /users/signup - Register a new user.
- POST /users/login - Authenticate and obtain access and refresh tokens.

- GET /users - Retrieves a list of all users in the system only for `ADMIN`.
- GET /users/:user_id - Retrieves the details of a specific user based on their `user_id`.

## Usage Examples

Once the server is running, you can interact with the API using tools like curl or Postman. Make sure to include the JWT token in the Authorization header for protected routes.

To register a new user:

```bash
curl -X POST http://localhost:<port>/users/signup -d '{"First_name":"Bil","Last_name":"Kosinsky", "Password":"password123", "Email":"email@gmail.com", "Phone":"000000000", "User_type":"USER"}' -H "Content-Type: application/json"
```

To log in and obtain tokens:

```bash
curl -X POST http://localhost:<port>/users/login -d '{"Password":"password123", "Email":"email@gmail.com"}' -H "Content-Type: application/json"
```

To get list of all users(if you are `ADMIN`):

```bash
curl -H "Authorization: Bearer <JWT_TOKEN>" http://localhost:<port>/users
```

To get user by id:

```bash
curl -H "Authorization: Bearer <JWT_TOKEN>" http://localhost:<port>/users/<user_id>
```

# Contributing

Feel free to open issues or submit pull requests if you have suggestions or improvements.