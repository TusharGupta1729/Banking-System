# Banking System API

A backend banking system built with Go, Gin, PostgreSQL, and JWT authentication. This project models banks, branches, customers, accounts, transactions, and loans in a clean REST API.

---

## Overview

This Banking System API is a simple backend for bank operations. Users can sign up, log in, open accounts, deposit and withdraw money, view transactions, and manage loans.

---

## Features

- Customer registration and login with secure password hashing
- JWT authentication for protected banking routes
- Create and manage banks and branches
- Create and retrieve customer profiles
- Open savings accounts for customers at specific branches
- Check account balances and account details
- Deposit money into accounts
- Withdraw money from accounts, with balance checks
- View transaction history for each account
- Apply for loans and track loan status
- Repay pending loan balances

---

## Tech Stack

- Go
- Gin web framework
- PostgreSQL
- GORM ORM
- JWT authentication

---

## System Design

### Core Entities

- Bank
- Branch
- Customer
- Account
- Transaction
- Loan

### Relationships

```text
Bank
 └── Branch
      └── Account
             └── Transaction

Customer
 ├── Account
 └── Loan
```

---

## Database Schema

### Bank

| Field       | Type      |
|-------------|-----------|
| id          | UUID      |
| name        | VARCHAR   |
| head_office | VARCHAR   |
| created_at  | TIMESTAMP |

### Branch

| Field      | Type      |
|------------|-----------|
| id         | UUID      |
| bank_id    | UUID      |
| name       | VARCHAR   |
| ifsc_code  | VARCHAR   |
| address    | TEXT      |

### Customer

| Field       | Type      |
|-------------|-----------|
| id          | UUID      |
| name        | VARCHAR   |
| email       | VARCHAR   |
| phone       | VARCHAR   |
| password    | VARCHAR   |
| created_at  | TIMESTAMP |

### Account

| Field          | Type      |
|----------------|-----------|
| id             | UUID      |
| customer_id    | UUID      |
| branch_id      | UUID      |
| account_number | VARCHAR   |
| balance        | DECIMAL   |
| account_type   | VARCHAR   |
| status         | VARCHAR   |
| created_at     | TIMESTAMP |

### Transaction

| Field      | Type      |
|------------|-----------|
| id         | UUID      |
| account_id  | UUID     |
| type       | VARCHAR   |
| amount     | DECIMAL   |
| created_at | TIMESTAMP |

### Loan

| Field            | Type      |
|------------------|-----------|
| id               | UUID      |
| customer_id      | UUID      |
| principal_amount | DECIMAL   |
| interest_rate    | DECIMAL   |
| total_amount     | DECIMAL   |
| pending_amount   | DECIMAL   |
| status           | VARCHAR   |
| created_at       | TIMESTAMP |

---

## API Endpoints

### Authentication

- POST /auth/register — Register a new user
- POST /auth/login — Log in and receive a JWT token

### Banks

- POST /banks — Create a new bank
- GET /banks — List all banks

### Branches

- POST /branches — Create a branch
- GET /banks/:id/branches — List branches for a specific bank

### Accounts

- POST /accounts — Open a savings account
- GET /accounts/:id — Get account details
- POST /accounts/:id/deposit — Deposit money into an account
- POST /accounts/:id/withdraw — Withdraw money from an account
- GET /accounts/:id/transactions — Get transaction history for an account

### Loans

- POST /loans — Apply for a loan
- GET /loans/:id — Get loan details
- POST /loans/:id/repay — Repay a loan
- GET /customers/:id/loans — List loans for a customer

---

## Architecture

The project follows a layered architecture with a clear separation of concerns.

```
Client
  ↓
Routes
  ↓
Handlers
  ↓
Services
  ↓
Repositories
  ↓
PostgreSQL
```

Responsibilities:
- Routes: define API endpoints and map requests to handlers
- Handlers: process HTTP requests and responses
- Services: implement business logic and validation
- Repositories: perform database queries and persistence
- Models: represent database entities and data structures

## Project Structure

```text
banking-system/
├── cmd/
│   └── main.go
├── config/
│   └── database.go
├── handlers/
│   ├── auth_handler.go
│   ├── account_handler.go
│   ├── bank_handler.go
│   ├── branch_handler.go
│   └── loan_handler.go
├── middleware/
│   └── auth_middleware.go
├── models/
│   ├── bank.go
│   ├── branch.go
│   ├── customer.go
│   ├── account.go
│   ├── transaction.go
│   └── loan.go
├── repository/
│   ├── bank_repository.go
│   ├── account_repository.go
│   └── loan_repository.go
├── services/
│   ├── auth_service.go
│   ├── account_service.go
│   └── loan_service.go
├── routes/
│   └── routes.go
├── migrations/
├── docs/
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

## Environment Variables

Example `.env` values:

```env
DATABASE_URL=postgres://user:password@localhost:5432/banking_db?sslmode=disable
JWT_SECRET=your_jwt_secret
PORT=8080
```

---

## Setup & Run

1. Clone the repository.
2. Create a PostgreSQL database.
3. Copy `.env.example` to `.env` and set your database credentials.
4. Install dependencies:

```bash
go mod download
```

5. Run the application:

```bash
go run ./cmd/main.go
```

---

## Author

Tushar Gupta

B.Tech Computer Science & Engineering
Delhi Technological University (DTU)
