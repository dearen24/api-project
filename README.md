# 🚀 API Project – Backend Framework Performance Comparison

This project is an **exploration and benchmark** of different backend frameworks.  
The goal is to compare **performance, scalability, and developer experience** using the same API specification and database setup.
This is an **ONGOING** project. Feel free to contact me on LinkedIn (https://www.linkedin.com/in/dearen-hippy-2994ab1b7)
I will upload the testing results soon!

---

## 📌 Frameworks Included

Currently implemented:

- 🐍 [Django](https://www.djangoproject.com/) (`djangoapi`)
- ⚡ [FastAPI](https://fastapi.tiangolo.com/) (`fast_api`)
- 🌐 [Express.js](https://expressjs.com/) (`express-api`)
- 🥷 [NestJS](https://nestjs.com/) (`nest-api`)
- 🐹 [Go Gin](https://gin-gonic.com/) (`go-gin-api`)
- 🌱 [Go Fiber](https://gofiber.io/) (`go-fiber-api`)
- ☕ [Spring Boot](https://spring.io/projects/spring-boot) (`spring-boot-api`)

More may be added over time!

---

## 🗄️ Database Setup

- **MySQL** (containerized with Docker)  
- Schema and seed data are in:
  - `create-table.txt`
  - `insert-table.txt`
 
Currently I only used 1 table for the API and also the testing. The other will be added soon!

---

## 📊 Benchmarking Tools

Performance tests are powered by:

- [Grafana k6](https://k6.io/) → load testing  
- Docker Compose → orchestrating DB and services  

Each API is tested against the same conditions to ensure **fair performance comparisons**.

---

## 🏗️ Getting Started

### 1️⃣ Clone the repo
```bash
git clone https://github.com/dearen24/api-project.git
cd api-project
