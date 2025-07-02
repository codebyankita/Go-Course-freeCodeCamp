# 📚 Go Bookstore API

A simple RESTful API for managing a bookstore, built with Go using the Gorilla Mux router and GORM ORM for MySQL.

## 🛠️ Features

- CRUD operations for books and authors
- RESTful design
- Uses GORM ORM with MySQL support
- Built with `mux` router

## 📦 Project Structure

```

go-bookstore/
│
├── main.go                 # Entry point
├── go.mod                 # Go module file
├── go.sum                 # Go dependency checksums
├── controllers/           # Handlers for HTTP endpoints
├── models/                # DB models
├── routes/                # Router setup
└── config/                # DB connection setup

````

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/codebyankita/go-bookstore.git
cd go-bookstore
````

### 2. Initialize Go module (already done)

```bash
go mod init github.com/codebyankita/go-bookstore
```

### 3. Install dependencies

```bash
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql
```

### 4. Setup MySQL database

Make sure MySQL is running and create a database:

```sql
CREATE DATABASE go_bookstore;
```

Update your DB config in `config/config.go`:

```go
const (
  DBUser     = "root"
  DBPassword = "your_password"
  DBName     = "go_bookstore"
)
```

### 5. Run the server

```bash
go run main.go
```

Server will start at:

```
http://localhost:8000
```

---

## 📌 Sample API Endpoints

| Method | Endpoint     | Description        |
| ------ | ------------ | ------------------ |
| GET    | /movies      | Get all movies     |
| GET    | /movies/{id} | Get movie by ID    |
| POST   | /movies      | Create new movie   |
| PUT    | /movies/{id} | Update movie by ID |
| DELETE | /movies/{id} | Delete movie by ID |

### Example Payload

```json
{
  "isbn": "123456789",
  "title": "The Go Workshop",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

---

## 🧪 Test API with Postman

You can test your endpoints using Postman or `curl`.

```bash
curl http://localhost:8000/movies
```

---

## 🧰 Tech Stack

* **Go** – Programming language
* **Gorilla Mux** – Routing
* **GORM** – ORM for database
* **MySQL** – Database

---

## 📄 License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.

```

---

Let me know if you want me to generate:
- Example `config.go`, `routes.go`, and `controller.go` files
- Dockerfile and docker-compose setup
- Swagger documentation

I can help scaffold the full bookstore API if you're building it step by step.
```

