# JWT-based Authentication TODO Backend using Go Fiber

This project is a simple TODO backend built using Go Fiber, with JWT-based authentication. It provides endpoints for user authentication (signup and signin) and TODO CRUD operations.

## Installation

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   ```

2. **Navigate to the project directory:**

   ```bash
   cd <project-directory>
   ```

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Build the project:**

   ```bash
   go build
   ```

5. **Run the executable:**

   ```bash
   ./<executable-name>
   ```

## API Endpoints

### User Authentication

#### Create User

- **Endpoint:** `POST /api/user/signup`
- **Request Body:**
  ```json
  {
      "name": "name",
      "email": "email",
      "password": "password"
  }
  ```

#### SignIn User

- **Endpoint:** `POST /api/user/signin`
- **Request Body:**
  ```json
  {
      "email": "email",
      "password": "password"
  }
  ```

### TODO CRUD Operations

#### Create TODO

- **Endpoint:** `POST /api/todos`
- **Request Body:**
  ```json
  {
      "title": "--title--",
      "sub_title": "send",
      "estimate_sprint": "1h",
      "is_done": true,
      "priority": 2
  }
  ```

#### List TODOs

- **Endpoint:** `GET /api/todos/`
- **Response Body:**
  ```json
  {
      "data": [
          {
              "id": "fdf1aa98-de31-409a-ae6c-6a5c78e3d32a",
              "title": "work on major project",
              "sub_title": "send mails",
              "estimate_sprint": "5h",
              "is_done": false,
              "priority": 2
          },
          // Additional TODOs...
      ]
  }
  ```
#### Edit/Delete TODOs

- **Endpoint:** `GET /api/todos/--uuid--`
- **Response Body:**
  ```json
  {
      "data": [
          {
              "id": "fdf1aa98-de31-409a-ae6c-6a5c78e3d32a",
              "title": "work on major project",
              "sub_title": "send mails",
              "estimate_sprint": "5h",
              "is_done": false,
              "priority": 2
          },
          // Additional TODOs...
      ]
  }
  ```

## Authentication and Authorization

- Authentication is handled using JWT tokens. Users need to sign in to obtain a token, which should be included in the Authorization header for accessing TODO endpoints.

## Contributing

Feel free to contribute to this project by creating issues or submitting pull requests. Please follow the [code of conduct](CODE_OF_CONDUCT.md).

## License

This project is licensed under the [MIT License](LICENSE).
