# Tournament Leaderboard API

This project is a Go-based backend API for managing Coding tournament's . It allows for the creation of tournaments, user registration, and admin management. The API includes features for handling users, admins, and tournament-related data.

---

### Features

-   **Admin Management**: APIs for admin signup and login.
-   **User Management**: APIs for user registration and authentication.
-   **Tournament Management**:
    -   Create and manage tournaments.
    -   Add problems and test cases to tournaments.
    -   Add participants to ongoing tournaments.
    -   Retrieve tournament details, ongoing tournaments, and participant lists.
-   **Authentication**: Secure endpoints using **JWT (JSON Web Tokens)** for both user and admin roles.
-   **Metrics**: Integrated **Prometheus** for monitoring application metrics.
-   **API Documentation**: Auto-generated **Swagger** documentation for all endpoints.

---

### Project Structure

The project follows a clean architecture pattern, with distinct layers for handlers, services, and repositories.

-   **`cmd/api`**: The entry point for the application.
-   **`config`**: Handles application configuration loading.
-   **`internal`**: Contains the core logic of the application.
    -   **`admin`**: Admin-related handlers, services, and repository.
    -   **`user`**: User-related handlers, services, and repository.
    -   **`tournament`**: Tournament-related handlers, services, and repository.
    -   **`common`**: Shared components like JWT functionality and middleware.
-   **`docs`**: Auto-generated Swagger documentation files.

---

### Getting Started

#### Prerequisites

-   Go 1.22 or higher
-   Docker (for database setup)

#### Setup

1.  **Clone the repository**:
    ```bash
    git clone [https://github.com/SLANGERES/Tournament-Lederboard.git](https://github.com/SLANGERES/Tournament-Lederboard.git)
    cd Tournament-Lederboard
    ```

2.  **Configure the environment**:
    Create a `.env` file in the root directory with your database connection strings and JWT key. You can use a file like `.env.example` as a template.

3.  **Run with Docker Compose**:
    This will set up the necessary databases (e.g., PostgreSQL).
    ```bash
    docker-compose up -d
    ```

4.  **Install dependencies**:
    ```bash
    go mod tidy
    ```

5.  **Generate Swagger documentation**:
    You need to install the `swag` tool first: `go install github.com/swaggo/swag/cmd/swag@latest`.
    ```bash
    swag init -g cmd/api/main.go
    ```

6.  **Run the application**:
    ```bash
    go run cmd/api/main.go
    ```

The API will be running on the address specified in your configuration (e.g., `http://localhost:8080`).

---

### API Documentation

The API endpoints are documented using **Swagger**. Once the server is running, you can access the interactive documentation at:

-   `http://localhost:8080/swagger/index.html`

This will provide a complete list of all available endpoints, their parameters, and expected responses.

---

### Endpoints

Below is a brief overview of the key endpoints:

| Method | Path                                          | Description                                       |
| :----- | :-------------------------------------------- | :------------------------------------------------ |
| `POST` | `/v1/admins/signup`                           | Register a new admin.                             |
| `POST` | `/v1/admins/login`                            | Authenticate an admin and receive a JWT.          |
| `POST` | `/v1/users/signup`                            | Register a new user.                              |
| `POST` | `/v1/users/login`                             | Authenticate a user and receive a JWT.            |
| `POST` | `/v1/tournaments`                             | Create a new tournament (Admin only).             |
| `GET`  | `/v1/tournaments/ongoing`                     | Get a list of all ongoing tournaments.            |
| `GET`  | `/v1/tournaments/{id}`                        | Get details of a specific tournament.             |
| `POST` | `/v1/tournaments/addproblem/{id}`             | Add a new problem to a tournament (Admin only).   |
| `POST` | `/v1/tournaments/addtestcase/{id}`            | Add a new test case to a problem (Admin only).    |
| `POST` | `/v1/tournaments/ongoing/addparticipant/{id}` | Add a participant to an ongoing tournament.       |
| `GET`  | `/v1/tournaments/ongoing/{id}/problem`        | Get problems for an ongoing tournament.           |
| `GET`  | `/v1/tournaments/{id}/participants`           | Get all participants of a tournament.             |

---

### Metrics

Prometheus metrics are exposed at the `/metrics` endpoint.

-   **Endpoint**: `http://localhost:7070/metrics`

---

### Contributing

Feel free to submit issues or pull requests to improve the project.