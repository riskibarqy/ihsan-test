# Riski Ramdan - Ihsan Solusi Assesment Test

## Overview

This is a backend service for user and balance history management. It follows a clean architecture structure, making it modular and scalable.

## What is Clean Architecture?

Clean Architecture is a software design philosophy that separates the concerns of an application into distinct layers, making the system more maintainable, testable, and scalable. The main layers include:

-   **Entities**: Represent the core business logic and rules.
-   **Use Cases**: Contain application-specific business rules and coordinate between repositories and entities.
-   **Interface Adapters**: Handle user interfaces, controllers, gateways, and data conversions.
-   **Infrastructure**: Includes frameworks, external services, and databases.

By structuring the application this way, dependencies point inward, reducing the risk of tight coupling and making the codebase more adaptable to changes.

## Project Structure

```
cmd/
  ├── main.go  # Entry point of the application
datatransfers/
  ├── base.go  # Base data transfer objects
  ├── user.go  # User-related DTOs
  ├── user_balance_history.go  # DTOs for user balance history
internal/
  ├── config/
  │   ├── config.go  # Application configuration setup
  ├── constants/
  │   ├── base.go  # Base constants used across the application
  │   ├── error.go  # Error definitions and handling
  ├── delivery/
  │   ├── http/
  │   │   ├── router.go  # HTTP router setup and initialization
  │   │   ├── user_handler.go  # HTTP handler for user operations
  │   │   ├── user_balance_history_handler.go  # HTTP handler for user balance history
  ├── domain/
  │   ├── user.go  # User entity definition
  │   ├── user_balance_history.go  # User balance history entity
  │   ├── user_repository.go  # User repository interface
  │   ├── user_balance_history_repository.go  # User balance history repository interface
  ├── repository/
  │   ├── user_repository.go  # Implementation of user repository
  │   ├── user_balance_history_repository.go  # Implementation of user balance history repository
  ├── usecase/
  │   ├── user_usecase.go  # Business logic for user operations
  │   ├── user_balance_history_usecase.go  # Business logic for user balance history
migrations/
  ├── SQL migration files for database schema changes
pkg/
  ├── database/
  │   ├── database.go  # Database connection setup and management
  ├── utils.go  # Utility functions used across the application
.env  # Environment variables configuration
.env_example  # Example environment variables template
.gitignore  # Ignored files and directories
Dockerfile  # Docker container setup configuration
docker-compose.yaml  # Docker Compose configuration for local development
go.mod / go.sum  # Go module dependencies
LICENSE  # License file
Makefile  # Automation scripts for building and running the application
README.md  # Project documentation
```

## Installation & Setup

### Prerequisites

-   Go 1.21+
-   Docker & Docker Compose
-   PostgreSQL

### Steps

1.  Clone the repository:
    ```sh
    git clone https://github.com/riskibarqy/ihsan-test
    cd ihsan-test
    ```
2.  Copy the example environment file and configure it:
    ```sh
    cp .env_example .env
    ```
    
3.  Run the application using Docker:
    ```sh
    docker-compose up --build -d
    ```
    
## API Endpoints

Refer to `Ihsan Solusi Test.postman_collection.json` for available API endpoints.

## Explanation Video

https://drive.google.com/file/d/1NpL-mUvkS87DmLQsnzfThVHMx6fyGk04/view?usp=drive_link

## License

This project is licensed under the MIT License. See `LICENSE` for details.
