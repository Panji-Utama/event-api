# Event Management API

Event Management API is a Go application built with Gin and Gorm to manage events, including creation, retrieval, updating, and deletion.

## Description

Event Management API provides endpoints to manage events, allowing users to perform CRUD (Create, Read, Update, Delete) operations on event data. It provides a RESTful API interface that can be consumed by various clients.

## Technologies Used

- [Gin](https://github.com/gin-gonic/gin): Web framework for Go.
- [Gorm](https://gorm.io/): ORM library for Go.
- [ElephantSQL](https://www.elephantsql.com/): PostgreSQL as a Service.

## Installation

1. Clone the repository:
    ```
    git clone https://github.com/Panji-Utama/event-api.git
    ```
2. Install dependencies:
    ```
    cd event-api
    go mod tidy
    ```

## Usage

To use the Event Management API, follow these steps:

1. Set up an ElephantSQL PostgreSQL database.
2. Configure the database connection details in your environment or `.env` file. Example:
    ```
    DATABASE_URL=your_database_url
    ```
3. Run the application:
    ```
    go run main.go
    ```
4. Access the API endpoints using a REST client or integrate them into your application.

Alternatively, you can use CompileDaemon to automatically compile and run your Go application whenever changes are detected. Here's how to use CompileDaemon:

1. Run your application with CompileDaemon to automatically restart the server when changes are made to your code. Use the following command:
    ```
    CompileDaemon -command="./your_application_binary"
    ```

Replace `your_application_binary` with the binary name of your Go application.

## License

This project is licensed under the [MIT License](LICENSE).
