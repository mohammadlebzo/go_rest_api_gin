# Golang REST API with Gin

## How to run

Run `./script.sh` or:

- Run `docker compose up -d` to create a MongoDB database.
- Run `cp .env.sample .env` to copy the content of `.env.sample` to a new `.env` file.
- Run `go run src/main.go` to run the project.

## Endpoits + Samples:

- **POST** `/api/login`. (_Authenticate a user and provide a token_)

  - Sample:
    ```
      {
          "username": "walter",
          "password": "testRaven621"
      }
    ```
  - Expected:
    ```
      {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.< Rest of The Token >"
      }
    ```

- **POST** `/api/admin/users`. (_Create a new user_)

  - Sample:
    ```
      {
          "id": 1,
          "username": "walter",
          "password": "testRaven621"
      }
    ```
  - Expected (_auth token needed_):
    ```
      {
        "message": "Registration Success"
      }
    ```

- **POST** `/api/admin/logout`. (_Invalidate the user token_)

  - Expected (_auth token needed_):

    ```
      {
        "message": "Logout Success"
      }
    ```

- **GET** `/api/admin/users/me`. (_Get the current authenticated user_)

  - Expected (_auth token needed_):
    ```
      {
        "data": {
          "id": 0,
          "username": "raven",
          "password": ""
        },
        "message": "success"
      }
    ```

- **GET** `/api/admin/users/:id`. (_Get a user if current user is authenticated user_)

  - Expected (_auth token needed_) / `:id = 1` :

    ```
      {
        "data": {
          "id": 1,
          "username": "walter",
          "password": ""
        },
        "message": "success"
      }
    ```

- **GET** `/api/admin/users`. (_Get all users if current user is authenticated user_)

  - Expected (_auth token needed_):

    ```
      [
        {
          "id": 0,
          "username": "raven",
          "password": "$2a$10$eWnbNsFKXxCto8xf7OjgX.P5OrAUhREvqQjIXBNXkUZC5bU9y4rNy"
        },
        {
          "id": 1,
          "username": "walter",
          "password": "$2a$10$yBMUpQush1Vv4rygEhCmYe.OWDXoL2HLAvla1GUHlV1HJQAbRldZG"
        }
      ]
    ```
