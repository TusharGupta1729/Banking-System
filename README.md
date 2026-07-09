# Banking System API

A backend banking system built with Go, Gin, PostgreSQL, and JWT authentication. This project models banks, branches, customers, accounts, transactions, and loans in a clean REST API.

---

## Overview

This Banking System API is a simple backend for bank operations. Users can sign up, log in, open accounts, deposit and withdraw money, view transactions, and manage loans.

---

## Features

- Customer registration and login with secure password hashing
- JWT authentication for deposit, withdrawal, and transfer routes
- Create and manage banks and branches
- Create and retrieve customer profiles
- Open accounts for customers at specific branches
- Check account balances and account details
- Deposit money into accounts
- Withdraw money from accounts, with balance checks
- View transaction history for each account
- Apply for loans and track loan status
- Repay pending loan balances

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

| Field      | Type      |
|------------|-----------|
| id         | SERIAL    |
| name       | VARCHAR   |
| headoffice | VARCHAR   |
| createdAt  | TIMESTAMP |

Column definitions:

- `id`: Unique bank record ID. In code, this comes from GORM's `gorm.Model`, which uses a `uint` primary key by default.
- `name`: Name of the bank.
- `headoffice`: Head office location of the bank.
- `createdAt`: Date and time when the bank record was created. In code, this comes from GORM's `gorm.Model`.

### Branch

| Field    | Type    |
|----------|---------|
| id       | SERIAL  |
| bankID   | INTEGER |
| name     | VARCHAR |
| iFSCCode | VARCHAR |
| address  | TEXT    |

Column definitions:

- `id`: Unique branch record ID. In code, this comes from GORM's `gorm.Model`, which uses a `uint` primary key by default.
- `bankID`: ID of the bank to which the branch belongs. In code, this is implemented as a `uint` foreign key.
- `name`: Name of the branch.
- `iFSCCode`: Unique IFSC code of the branch.
- `address`: Physical address of the branch.

### Customer

| Field        | Type      |
|--------------|-----------|
| id           | SERIAL    |
| name         | VARCHAR   |
| email        | VARCHAR   |
| phone        | VARCHAR   |
| passwordHash | VARCHAR   |
| createdAt    | TIMESTAMP |

Column definitions:

- `id`: Unique customer record ID. In code, this comes from GORM's `gorm.Model`, which uses a `uint` primary key by default.
- `name`: Full name of the customer.
- `email`: Unique email address used for login.
- `phone`: Customer phone number.
- `passwordHash`: Hashed password stored securely instead of plain text.
- `createdAt`: Date and time when the customer record was created. In code, this comes from GORM's `gorm.Model`.

### Account

| Field         | Type      |
|---------------|-----------|
| id            | SERIAL    |
| customerID    | INTEGER   |
| branchID      | INTEGER   |
| accountNumber | VARCHAR   |
| balance       | DECIMAL   |
| accountType   | VARCHAR   |
| createdAt     | TIMESTAMP |

Column definitions:

- `id`: Unique account record ID. In code, this comes from GORM's `gorm.Model`, which uses a `uint` primary key by default.
- `customerID`: ID of the customer who owns the account. In code, this is implemented as a `uint` foreign key.
- `branchID`: ID of the branch where the account is opened. In code, this is implemented as a `uint` foreign key.
- `accountNumber`: Unique account number.
- `balance`: Current available balance in the account.
- `accountType`: Type of account, such as savings or current.
- `createdAt`: Date and time when the account record was created. In code, this comes from GORM's `gorm.Model`.

### Transaction

| Field     | Type      |
|-----------|-----------|
| id        | SERIAL    |
| accountID | INTEGER   |
| type      | VARCHAR   |
| amount    | DECIMAL   |
| createdAt | TIMESTAMP |

Column definitions:

- `id`: Unique transaction record ID. In code, this comes from GORM's `gorm.Model`, which uses a `uint` primary key by default.
- `accountID`: ID of the account linked with the transaction. In code, this is implemented as a `uint` foreign key.
- `type`: Transaction type, such as deposit, withdraw, transfer in, or transfer out.
- `amount`: Amount involved in the transaction.
- `createdAt`: Date and time when the transaction record was created. In code, this comes from GORM's `gorm.Model`.

### Loan

| Field           | Type      |
|-----------------|-----------|
| id              | SERIAL    |
| customerID      | INTEGER   |
| principalAmount | DECIMAL   |
| interestRate    | DECIMAL   |
| totalAmount     | DECIMAL   |
| pendingAmount   | DECIMAL   |
| status          | VARCHAR   |
| createdAt       | TIMESTAMP |

Column definitions:

- `id`: Unique loan record ID. In code, this comes from GORM's `gorm.Model`, which uses a `uint` primary key by default.
- `customerID`: ID of the customer who applied for the loan. In code, this is implemented as a `uint` foreign key.
- `principalAmount`: Original loan amount requested by the customer.
- `interestRate`: Interest rate applied to the loan.
- `totalAmount`: Total amount to be repaid.
- `pendingAmount`: Remaining unpaid loan amount.
- `status`: Current loan status, such as pending, approved, rejected, or closed.
- `createdAt`: Date and time when the loan record was created. In code, this comes from GORM's `gorm.Model`.

---

## API Endpoints

### Root

- GET / — Check API running message

### Authentication

- POST /login — Log in and receive a JWT token

### Banks

- POST /banks — Create a new bank
- GET /banks — List all banks

### Branches

- POST /branches — Create a branch
- GET /branches — List all branches

### Customers

- POST /customers — Create a customer
- GET /customers — List all customers
- GET /customers/:id/accounts — List accounts for a customer

### Accounts

- POST /accounts — Open an account
- GET /accounts — List all accounts
- GET /accounts/:id — Get account details
- POST /accounts/:id/deposit — Deposit money into an account
- POST /accounts/:id/withdraw — Withdraw money from an account
- POST /accounts/transfer — Transfer money between accounts
- GET /accounts/:id/transactions — Get transaction history for an account

### Transactions

- POST /transactions — Create a transaction
- GET /transactions — List all transactions

### Loans

- POST /loans — Apply for a loan
- GET /loans — List all loans
- POST /loans/:id/approve — Approve a loan
- POST /loans/:id/reject — Reject a loan
- POST /loans/:id/repay — Repay a loan

---

## API Request and Response Examples

### Login

`POST /login`

Request:

```json
{
  "email": "tushar@example.com",
  "password": "123456"
}
```

Response:

```json
{
  "token": "jwt_token_here"
}
```

### Create Bank

`POST /banks`

Request:

```json
{
  "name": "State Bank",
  "headoffice": "Delhi"
}
```

Response:

```json
{
  "id": 1,
  "name": "State Bank",
  "headoffice": "Delhi",
  "createdAt": "2026-07-09T10:00:00Z"
}
```

### Create Branch

`POST /branches`

Request:

```json
{
  "bankID": 1,
  "name": "Delhi Main Branch",
  "iFSCCode": "SBIN0001234",
  "address": "Connaught Place, Delhi"
}
```

Response:

```json
{
  "id": 1,
  "bankID": 1,
  "name": "Delhi Main Branch",
  "iFSCCode": "SBIN0001234",
  "address": "Connaught Place, Delhi"
}
```

### Create Customer

`POST /customers`

Request:

```json
{
  "name": "Tushar Gupta",
  "email": "tushar@example.com",
  "phone": "9999999999",
  "passwordHash": "123456"
}
```

Response:

```json
{
  "id": 1,
  "name": "Tushar Gupta",
  "email": "tushar@example.com",
  "phone": "9999999999"
}
```

### Create Account

`POST /accounts`

Request:

```json
{
  "customerID": 1,
  "branchID": 1,
  "accountNumber": "ACC1001",
  "balance": 0,
  "accountType": "Savings"
}
```

Response:

```json
{
  "id": 1,
  "customerID": 1,
  "branchID": 1,
  "accountNumber": "ACC1001",
  "balance": 0,
  "accountType": "Savings"
}
```

### Deposit Money

`POST /accounts/:id/deposit`

Request:

```json
{
  "amount": 5000
}
```

Response:

```json
{
  "message": "deposit successful"
}
```

### Withdraw Money

`POST /accounts/:id/withdraw`

Request:

```json
{
  "amount": 1000
}
```

Response:

```json
{
  "message": "withdraw successful"
}
```

### Transfer Money

`POST /accounts/transfer`

Request:

```json
{
  "fromAccount": 1,
  "toAccount": 2,
  "amount": 1500
}
```

Response:

```json
{
  "message": "transfer successful"
}
```

### Create Transaction

`POST /transactions`

Request:

```json
{
  "accountID": 1,
  "type": "Deposit",
  "amount": 5000
}
```

Response:

```json
{
  "id": 1,
  "accountID": 1,
  "type": "Deposit",
  "amount": 5000
}
```

### Create Loan

`POST /loans`

Request:

```json
{
  "customerID": 1,
  "principalAmount": 100000,
  "interestRate": 10,
  "totalAmount": 110000
}
```

Response:

```json
{
  "id": 1,
  "customerID": 1,
  "principalAmount": 100000,
  "interestRate": 10,
  "totalAmount": 110000,
  "pendingAmount": 110000,
  "status": "Pending"
}
```

### Approve Loan

`POST /loans/:id/approve`

Response:

```json
{
  "message": "loan approved successfully"
}
```

### Reject Loan

`POST /loans/:id/reject`

Response:

```json
{
  "message": "loan rejected successfully"
}
```

### Repay Loan

`POST /loans/:id/repay`

Request:

```json
{
  "amount": 5000
}
```

Response:

```json
{
  "message": "loan repayment successful"
}
```

### Error Response

```json
{
  "error": "invalid request body"
}
```

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
