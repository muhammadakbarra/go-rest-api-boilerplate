# 🚀 Go REST API Boilerplate

[![Go Report Card](https://goreportcard.com/badge/github.com/muhammadakbarra/go-rest-api-boilerplate)](https://goreportcard.com/report/github.com/muhammadakbarra/go-rest-api-boilerplate)
[![Go Doc](https://img.shields.io/badge/godoc-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/muhammadakbarra/go-rest-api-boilerplate)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A modern, scalable, and high-performance REST API boilerplate built with **Go**, **Chi Router**, and **PostgreSQL**. Designed to be the perfect starting point for your next microservice or web backend.

---

## ✨ Features

- **Standard Layout**: Follows [golang-standards/project-layout](https://github.com/golang-standards/project-layout).
- **Fast Routing**: Powered by [go-chi/chi](https://github.com/go-chi/chi) for lightweight and idiomatic routing.
- **Auto Swagger**: Interactive API documentation with [swaggo/swag](https://github.com/swaggo/swag).
- **Environment Ready**: Configuration management via `.env` files.
- **Clean Architecture**: Clear separation between Handlers, Repositories, and Models.
- **PostgreSQL**: Robust data persistence with `pgx` connection pooling.

---

## 🛠 Tech Stack

- **Language:** Go 1.22+
- **Framework:** [go-chi/chi v5](https://github.com/go-chi/chi)
- **Database:** PostgreSQL
- **Driver:** [jackc/pgx v5](https://github.com/jackc/pgx)
- **Documentation:** [swaggo/http-swagger](https://github.com/swaggo/http-swagger)

---

## 📂 Project Structure

```text
.
├── cmd/
│   └── api/            # Main entry point (main.go)
├── internal/
│   ├── config/         # App configuration & env loading
│   ├── database/       # DB connection & pooling logic
│   └── posts/          # Domain logic (Handler, Repository, Model)
├── docs/               # Auto-generated Swagger files
├── .env.example        # Template for env variables
└── go.mod              # Dependency management
```

---

## 🚀 Getting Started

### 1. Prerequisites
- Go installed (version 1.22 or higher)
- PostgreSQL running locally or via Docker

### 2. Installation
Clone the repository and install dependencies:
```bash
git clone https://github.com/muhammadakbarra/go-rest-api-boilerplate.git
cd go-rest-api-boilerplate
go mod tidy
```

### 3. Configuration
Copy the example environment file and update your credentials:
```bash
cp .env.example .env
```
Edit `.env`:
```env
APP_PORT=8080
DATABASE_URL=postgresql://user:password@localhost:5432/your_db_name
```

### 4. Database Setup
Once `.env` is configured, run the migration to create the necessary tables:
```bash
make migrate
```

### 5. Running the App
```bash
make run
```
The server will start at `http://localhost:8080`.

---

## 📖 API Documentation

The documentation is automatically generated. Once the server is running, visit:
👉 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

### Updating Documentation
If you modify any API handlers, regenerate the Swagger files:
```bash
make swagger
```

---

## 🛠 Development Commands

| Command | Description |
| :--- | :--- |
| `make run` | Start the development server |
| `make migrate` | Run database migrations |
| `make swagger` | Regenerate Swagger documentation |
| `make build` | Build the binary |
| `make test` | Run all tests |
| `make help` | Show all available commands |

---

## 🤝 Contributing

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📝 License

Distributed under the MIT License. See `LICENSE` for more information.

---
<p align="center">Made with ❤️ for the Go Community</p>
