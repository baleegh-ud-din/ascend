# Ascend

Ascend is a full-stack web application designed for efficient management and tracking of tickets and tasks. It consists of a backend server written in Go and a frontend built using React JS.

## 🚀 Features
- RESTful API backend built with Go
- Frontend built using React with Vite
- PostgreSQL database integration
- Kubernetes deployment support
- Background jobs scheduling
- Secure authentication system

## 📁 Project Structure
```
.
├── LICENSE
├── Makefile
├── README.md
├── app
│   ├── server
│   │   └── ascend
│   └── static
│       ├── assets
│       │   ├── index-D8b4DHJx.css
│       │   ├── index-DO1VeTRM.js
│       │   └── react-CHdo91hT.svg
│       ├── index.html
│       └── vite.svg
├── bin
│   └── v0.0.1
│       └── ascend
├── cmd
│   └── main.go
├── config
│   └── config.go
├── database
│   ├── database.go
│   ├── dbutils.go
│   ├── migrations.go
│   └── schema.go
├── emails
│   └── email.go
├── go.mod
├── go.sum
├── handlers
│   └── handlers.go
├── jobs
│   └── jobs.go
├── jwt
│   └── jwt.go
├── k8s
│   ├── Dockerfile
│   └── deployment.yaml
├── middleware
│   └── middleware.go
├── models
│   └── models.go
├── routes
│   └── routes.go
├── ui
│   ├── README.md
│   ├── eslint.config.js
│   ├── index.html
│   ├── node_modules/
│   ├── package.json
│   ├── pnpm-lock.yaml
│   ├── public
│   │   └── vite.svg
│   ├── src
│   │   ├── App.css
│   │   ├── App.tsx
│   │   ├── assets
│   │   │   └── react.svg
│   │   ├── components
│   │   │   ├── common
│   │   │   └── layout
│   │   ├── context
│   │   ├── hooks
│   │   ├── index.css
│   │   ├── main.tsx
│   │   ├── pages
│   │   ├── services
│   │   ├── types
│   │   ├── utils
│   │   └── vite-env.d.ts
│   ├── tsconfig.app.json
│   ├── tsconfig.json
│   ├── tsconfig.node.json
│   └── vite.config.ts
└── utils
    └── logger.go
```

## 🛠️ Setup & Installation
### **1️⃣ Install Dependencies**
Ensure you have the following installed:
- Go 1.24.0 or later
- Node.js & pnpm (for UI)
- PostgreSQL
- Docker & Kubernetes (for deployment)

### **2️⃣ Configure Environment**
Copy the `.env` file and update the necessary variables:
```sh
cp .env.example .env
```

### **3️⃣ Start Development Servers**
To run both the backend and frontend servers in development mode:
```sh
make dev
```
This will start:
- **Frontend** at `http://localhost:5173`
- **Backend** at `http://localhost:8443`

### **4️⃣ Build & Run the Application**
```sh
make build
make start
```

## 🚢 Deployment
### **Docker Build & Run**
```sh
docker build -t ascend-app -f k8s/Dockerfile .
docker run -p 8443:8443 ascend-app
```

### **Kubernetes Deployment**
```sh
kubectl apply -f k8s/deployment.yaml
```

## 📜 License
This project is licensed under the [Apache License](LICENSE).

## ✨ Contributors
- **Shaik** – Lead Developer
- Open for contributions! Feel free to submit issues & PRs.

---
For any queries, reach out at [hello@sbkashif.com]. 🚀

