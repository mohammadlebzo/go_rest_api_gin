# Golang REST API with Gin

## How to run

- Run `docker compose up -d` to create a MongoDB database.
- Run `cp .env.sample .env` to copy the content of `.env.sample` to a new `.env` file.
- Run `go run src/main.go` to run the project.

## Endpoits + Samples:

- **POST** `/api/login`. (_Authenticate a user and provide a token_)

  ```
    {
        "username": "raven",
        "password": "testRaven621"
    }
  ```

- **POST** `/api/admin/users`. Create a new user

  ```
    {
        "id": 0,
        "username": "raven",
        "password": "testRaven621"
    }
  ```

- **POST** `/api/admin/logout`. (_Invalidate the user token_)
- **GET** `/api/admin/users/me`. (_Get the current authenticated user_)
- **GET** `/api/admin/users/:id`. (_Get a user if current user is authenticated user_)
- **GET** `/api/admin/users`. (_Get all users if current user is authenticated user_)
