
# MiniTodo App

A simple backend API for a Todo application built with Go (Gin, GORM).

## 📌 Overview

This project is part of a learning phase to build a scalable and maintainable Go application.  
The final goal is to apply this experience to the development of the DriveShareSNS app.

## 🛠️ Tech Stack (Planned)

- Go
- Gin (Web Framework)
- GORM (ORM)
- PostgreSQL
- JWT Authentication
- Docker (later)
- Table-driven tests (later)

## 🚀 Getting Started

### 1. Clone this project

```bash
git clone https://github.com/akihiro-coder/driveshare-backend-go.git
cd driveshare-backend-go/learning/go_todo_app
```

### 2. Initialize Go module

```bash
go mod init github.com/akihiro-coder/driveshare-backend-go/learning/go_todo_app
```

### 3. Set up environment variables

Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=todo_app
JWT_SECRET=your_jwt_secret
```

### 4. Create .gitignore

Include the following:

```gitignore
/bin/
/build/
/pkg/mod/
*.exe
*.out
*.test
.vscode/
.idea/
.DS_Store
.env
```

### 5. Install Go dependencies

```bash
go mod tidy
```

### 6. Directory Structure (initial setup)

```
go_todo_app/
├── cmd/             # Entry point
├── internal/        # Handler, service, repository, model, etc.
├── pkg/             # Shared utils (e.g., JWT)
├── .env             # Local environment variables
├── .gitignore
├── go.mod
├── README.md
```

## 📝 Status

- [x] Step 1: Initialize go mod
- [x] Step 2: Setup .env, .gitignore, README
- [ ] Step 3: Setup directory structure
- [ ] Step 4: Start writing `main.go`