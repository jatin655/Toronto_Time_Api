# Toronto Time API

## Overview

This project provides an API that returns the current time in Toronto and logs each API request to a MySQL database.

## Features

- **Current Time API**: Returns Toronto time in JSON format via the `/current-time` endpoint.
- **Database Logging**: Logs each API request to a MySQL database.

## Prerequisites

Before setting up the project, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.19 or higher).
- [MySQL](https://dev.mysql.com/downloads/installer/).
- Required Go packages:
  - [Gorilla Mux](https://github.com/gorilla/mux) for routing.
  - [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) for MySQL integration.

## Setup Instructions

### 1. MySQL Database Setup

1. **Start MySQL**: Ensure the MySQL service is running.
2. **Log in to MySQL**:

   ```bash
   mysql -u root -p
    ```

3. **Create Database**:

```bash
CREATE DATABASE TimeLogger;
```

4. **Use the Database**:

``` bash
USE TimeLogger;
```

5. **Create Table**:

``` bash
CREATE TABLE time_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME NOT NULL
);
```

6. **Use the Database**:

```bash
USE TimeLogger;
Create Table:
sql
Copy code
CREATE TABLE time_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME NOT NULL
);
```

7. **Exit MySQL**:

```bash
exit
```

---


### 2. Project Setup

1. **Clone the Repository**:

```bash

git clone <repository-url>
cd toronto-time-api
```

2. **Initialize the Go Module**:

```bash

go mod init toronto-time-api
```

3. **Install Dependencies**:

```bash

go get -u github.com/gorilla/mux
go get -u github.com/go-sql-driver/mysql
```

4. **Configure Database Connection: Open db.go and update the dsn variable with your MySQL credentials**:

```bash
dsn := "username:password@tcp(127.0.0.1:3306)/TimeLogger"

```
Replace username and password with your MySQL credentials.

---



### 3. Run the Application

1. **Start the server**:

```bash

go run main.go db.go
```

The API will start at http://localhost:8080.

---


### 4. Test the API

1. **Use curl or a tool like Postman**:

```bash

curl http://localhost:8080/current-time
```

2. **Verify the response**:

```bash
{
    "current_time": "2024-12-01 15:30:00"
}
```


---


### Project Structure

```bash
toronto-time-api/
├── db.go           # Database connection setup
├── main.go         # Main application and API endpoint
├── go.mod          # Go module file
├── README.md       # Project documentation

```


### Notes

- Ensure the MySQL service is running when testing the API.
- Use appropriate error handling for production deployment.