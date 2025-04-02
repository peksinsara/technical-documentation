# TechDocs - Technical Documentation Platform

A modern platform for creating and managing technical documentation, service specifications, use cases, and diagrams.

## Features

- 📝 Markdown-based document creation and editing
- 🔍 Advanced search and filtering
- 📊 Service documentation
- 📋 Use case management

## Tech Stack

### Backend
- Go
- Gin web framework
- GORM ORM
- MySQL database
- JWT authentication

### Frontend
- Vue 3
- Vue Router
- Pinia state management
- Tailwind CSS
- Headless UI
- Heroicons

## Getting Started

### Prerequisites

- Go 1.16 or later
- Node.js 16 or later
- MySQL 8.0 or later

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file with your configuration:
   ```
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=techdocs
   JWT_SECRET=your_secret_key
   SERVER_PORT=8080
   ```

4. Run the backend server:
   ```bash
   go run cmd/api/main.go
   ```

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

4. Build for production:
   ```bash
   npm run build
   ```

## Project Structure

```
.
├── backend/
│   ├── cmd/
│   │   └── api/
│   │   ├── internal/
│   │   │   ├── config/
│   │   │   ├── handler/
│   │   │   ├── middleware/
│   │   │   ├── model/
│   │   │   ├── repository/
│   │   │   └── service/
│   │   └── pkg/
│   │   │   ├── auth/
│   │   │   ├── database/
│   │   │   └── logger/
│   └── frontend/
│       ├── src/
│       │   ├── components/
│       │   ├── views/
│       │   ├── router/
│       │   ├── stores/
│       │   └── assets/
│       └── public/
```


## License

This project is licensed under the MIT License - see the LICENSE file for details. 