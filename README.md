# Banking System API

A backend banking system built with Go, Gin, PostgreSQL, and JWT authentication. This project models banks, branches, customers, accounts, transactions, and loans in a clean REST API.

## Overview

This Banking System API is a simple backend for bank operations. Customers can sign up, log in, open accounts, deposit and withdraw money, transfer funds, view transaction history, and manage loans. Administrators can create banks and branches, open accounts on behalf of customers, view all records, and approve or reject loans.

## Features

- Customer registration and login with bcrypt password hashing
- JWT authentication with role-based access control (`customer` / `admin`)
- Create and manage banks and branches
- Create and retrieve customer profiles
- Open accounts for customers at specific branches
- Check account balances and account details
- Deposit money into accounts
- Withdraw money from accounts, with balance checks
- Transfer funds between accounts
- Automatic transaction records for deposits, withdrawals, and transfers
- View transaction history for each account
- Apply for loans and track loan status
- Admin approval/rejection of loans and customer repayment of pending balances

## Tech Stack

- Go
- Gin web framework
- PostgreSQL
- GORM ORM
- JWT authentication
- bcrypt password hashing

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

## Database Schema

All tables include an auto-managed `createdAt`, `updatedAt`, and `deletedAt` column supplied by GORM's `gorm.Model`. `updatedAt` and `deletedAt` (used for soft deletes) are omitted below for brevity.

### Bank

|    Field    |   Type         |
| ------------| ---------------|
| `id`        | `BIGSERIAL`    |
| `name`      | `VARCHAR(255)` |
| `headoffice`| `VARCHAR(255)` |
| `createdAt` | `TIMESTAMPTZ`  |


Column definitions:
- `id`: Unique bank record ID. In code, this comes from GORM's `gorm.Model`.
- `name`: Name of the bank.
- `headoffice`: Head office location of the bank.
- `createdAt`: Date and time when the bank record was created.

### Branch

|    Field    |      Type      |
| ------------| ---------------|
| `id`        | `BIGSERIAL`    |
| `bankID`    | `BIGINT`       |
| `name`      | `VARCHAR(255)` |
| `iFSCCode`  | `VARCHAR(20)`  |
| `address`   | `TEXT`         |
| `createdAt` | `TIMESTAMPTZ`  |

Column definitions:
- `id`: Unique branch record ID. In code, this comes from GORM's `gorm.Model`.
- `bankID`: ID of the bank to which the branch belongs, implemented as a foreign key.
- `name`: Name of the branch.
- `iFSCCode`: Unique IFSC code of the branch.
- `address`: Physical address of the branch.
- `createdAt`: Date and time when the branch record was created.

### Customer

|     Field      |      Type      |
| ---------------| ---------------|
| `id`           | `BIGSERIAL`    |
| `name`         | `VARCHAR(255)` |
| `email`        | `VARCHAR(255)` |
| `phone`        | `VARCHAR(20)`  |
| `passwordHash` | `VARCHAR(255)` |
| `role`         | `VARCHAR(20)`  |
| `createdAt`    | `TIMESTAMPTZ`  |

Column definitions:
- `id`: Unique customer record ID. In code, this comes from GORM's `gorm.Model`.
- `name`: Full name of the customer.
- `email`: Unique email address used for login.
- `phone`: Customer phone number.
- `passwordHash`: bcrypt hash of the customer's password, stored instead of plain text.
- `role`: Authorization role, defaults to `customer`; an `admin` can access administrative routes.
- `createdAt`: Date and time when the customer record was created.

### Account

|      Field      |      Type       |
| ----------------| ----------------|
| `id`            | `BIGSERIAL`     |
| `customerID`    | `BIGINT`        |
| `branchID`      | `BIGINT`        |
| `accountNumber` | `VARCHAR(50)`   |
| `balance`       | `NUMERIC(15,2)` |
| `accountType`   | `VARCHAR(20)`   |
| `status`        | `VARCHAR(20)`   |
| `createdAt`     | `TIMESTAMPTZ`   |

Column definitions:
- `id`: Unique account record ID. In code, this comes from GORM's `gorm.Model`.
- `customerID`: ID of the customer who owns the account, implemented as a foreign key.
- `branchID`: ID of the branch where the account is opened, implemented as a foreign key.
- `accountNumber`: Unique account number.
- `balance`: Current available balance in the account.
- `accountType`: Type of account, such as `Savings` or `Current`.
- `status`: Current account state, such as `Active` or `Inactive`.
- `createdAt`: Date and time when the account record was created.

### Transaction

|    Field    |      Type       |
| ------------| ----------------|
| `id`        | `BIGSERIAL`     |
| `accountID` | `BIGINT`        |
| `type`      | `VARCHAR(20)`   |
| `amount`    | `NUMERIC(15,2)` |
| `createdAt` | `TIMESTAMPTZ`   |

Column definitions:
- `id`: Unique transaction record ID. In code, this comes from GORM's `gorm.Model`.
- `accountID`: ID of the account linked with the transaction, implemented as a foreign key.
- `type`: Transaction type, such as `Deposit`, `Withdraw`, `Transfer In`, or `Transfer Out`.
- `amount`: Amount involved in the transaction.
- `createdAt`: Date and time when the transaction record was created.

Transactions are generated automatically by deposits, withdrawals, and transfers — there is no endpoint for creating one directly.

### Loan

|       Field       |      Type       |
| ------------------| ----------------|
| `id`              | `BIGSERIAL`     |
| `customerID`      | `BIGINT`        |
| `principalAmount` | `NUMERIC(15,2)` |
| `interestRate`    | `NUMERIC(5,2)`  |
| `totalAmount`     | `NUMERIC(15,2)` |
| `pendingAmount`   | `NUMERIC(15,2)` |
| `status`          | `VARCHAR(20)`   |
| `createdAt`       | `TIMESTAMPTZ`   |

Column definitions:
- `id`: Unique loan record ID. In code, this comes from GORM's `gorm.Model`.
- `customerID`: ID of the customer who applied for the loan, implemented as a foreign key.
- `principalAmount`: Original loan amount requested by the customer.
- `interestRate`: Interest rate applied to the loan.
- `totalAmount`: Total amount to be repaid.
- `pendingAmount`: Remaining unpaid loan amount, set from `totalAmount` when the loan is created.
- `status`: Current loan status — `Pending`, `Approved`, `Rejected`, or `Closed`.
- `createdAt`: Date and time when the loan record was created.

## Architecture

The project follows a layered architecture with a clear separation of concerns.

```text
Client
  ↓
Routes + Middleware
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

- **Routes**: define API endpoints, attach middleware, and map requests to handlers
- **Middleware**: verifies JWTs and enforces the `admin` role on protected routes
- **Handlers**: process HTTP requests and responses
- **Services**: implement business logic and validation
- **Repositories**: perform database queries and persistence
- **Models**: represent database entities and data structures

## API Endpoints

Endpoints are listed in the order they're registered in `routes/routes.go`. **Protected** routes require `Authorization: Bearer <token>`. **Admin** routes require that token's `role` claim to be `admin`.

### Root

| Method | Path  | Access | Description               |
| ------ | ----- | ------ | ------------------------- |
| GET    | `/`   | Public | Check API running message |

### Banks

| Method | Path     | Access | Description         |
| ------ | -------- | ------ | ------------------- |
| POST   | `/banks` | Admin  | Create a new bank   |
| GET    | `/banks` | Public | List all banks      |

### Branches

| Method | Path        | Access | Description       |
| ------ | ----------- | ------ | ----------------- |
| POST   | `/branches` | Admin  | Create a branch   |
| GET    | `/branches` | Public | List all branches |

### Customers

| Method | Path | Access | Description |
| ------ | ------------------------- | --------- | ---------------------------- |
| POST   | `/customers`              | Public    | Create a customer            |
| GET    | `/customers`              | Admin     | List all customers           |
| GET    | `/customers/:id/accounts` | Protected | List accounts for a customer |

### Accounts

| Method | Path | Access | Description |
| ------ | ---------------------------- | --------- | ---------------------------------------- |
| POST   | `/accounts`                  | Admin     | Open an account                          |
| GET    | `/accounts`                  | Admin     | List all accounts                        |
| POST   | `/accounts/:id/deposit`      | Protected | Deposit money into an account            |
| POST   | `/accounts/:id/withdraw`     | Protected | Withdraw money from an account           |
| GET    | `/accounts/:id`              | Protected | Get account details                      |
| POST   | `/accounts/transfer`         | Protected | Transfer money between accounts          |
| GET    | `/accounts/:id/transactions` | Protected | Get transaction history for an account   |

### Transactions

| Method | Path |  Access   | Description |                       |
| ------ | ---------------- | ------------| --------------------- |
| GET    | `/transactions`  | Admin       | List all transactions |

> There is no `POST /transactions` endpoint — records are created automatically by deposits, withdrawals, and transfers.

### Loans

| Method | Path | Access | Description |
| ------ | --------------------- | --------- | ---------------- |
| POST   | `/loans`              | Protected | Apply for a loan |
| GET    | `/loans`              | Protected | List loans       |
| POST   | `/loans/:id/approve`  | Admin     | Approve a loan   |
| POST   | `/loans/:id/reject`   | Admin     | Reject a loan    |
| POST   | `/loans/:id/repay`    | Protected | Repay a loan     |

### Authentication

| Method | Path | Access | Description |
| ------ | -------- | ------ | -------------------------------- |
| POST   | `/login` | Public | Log in and receive a JWT token   |

## API Request and Response Examples

### Check Server Status

`GET /`

Response:

```json
{
  "message": "Banking API Running",
  "status": "Server is running successfully",
  "usage": "Use Postman or any API client to access the available endpoints"
}
```

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

`POST /banks` — **Admin**

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
  "createdAt": "2026-07-14T10:00:00Z"
}
```

### Create Branch

`POST /branches` — **Admin**

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

> The `passwordHash` request field is treated as a plaintext password and is bcrypt-hashed before storage.

Response:

```json
{
  "id": 1,
  "name": "Tushar Gupta",
  "email": "tushar@example.com",
  "phone": "9999999999",
  "role": "customer"
}
```

### Get Customer Accounts

`GET /customers/:id/accounts` — **Protected**

Response:

```json
[
  {
    "id": 1,
    "customerID": 1,
    "branchID": 1,
    "accountNumber": "ACC1001",
    "balance": 5000,
    "accountType": "Savings",
    "status": "Active"
  }
]
```

### Create Account

`POST /accounts` — **Admin**

Request:

```json
{
  "customerID": 1,
  "branchID": 1,
  "accountNumber": "ACC1001",
  "balance": 0,
  "accountType": "Savings",
  "status": "Active"
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
  "accountType": "Savings",
  "status": "Active"
}
```

### Deposit Money

`POST /accounts/:id/deposit` — **Protected**

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

`POST /accounts/:id/withdraw` — **Protected**

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

### Get Account Details

`GET /accounts/:id` — **Protected**

Response:

```json
{
  "id": 1,
  "customerID": 1,
  "branchID": 1,
  "accountNumber": "ACC1001",
  "balance": 4000,
  "accountType": "Savings",
  "status": "Active"
}
```

### Transfer Money

`POST /accounts/transfer` — **Protected**

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

### Get Transaction History for an Account

`GET /accounts/:id/transactions` — **Protected**

Response:

```json
[
  {
    "id": 1,
    "accountID": 1,
    "type": "Deposit",
    "amount": 5000,
    "createdAt": "2026-07-14T10:20:00Z"
  },
  {
    "id": 2,
    "accountID": 1,
    "type": "Withdraw",
    "amount": 1000,
    "createdAt": "2026-07-14T10:22:00Z"
  }
]
```

### List All Transactions

`GET /transactions` — **Admin**

Response:

```json
[
  {
    "id": 1,
    "accountID": 1,
    "type": "Deposit",
    "amount": 5000,
    "createdAt": "2026-07-14T10:20:00Z"
  }
]
```

### Create Loan

`POST /loans` — **Protected**

Request:

```json
{
  "principalAmount": 100000,
  "interestRate": 10,
  "totalAmount": 110000
}
```

> The customer is taken from the JWT, so `customerID` does not need to be provided in the request body. New loans start with `Pending` status and `pendingAmount` equal to `totalAmount`.

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

### List Loans

`GET /loans` — **Protected**

Response:

```json
[
  {
    "id": 1,
    "customerID": 1,
    "principalAmount": 100000,
    "interestRate": 10,
    "totalAmount": 110000,
    "pendingAmount": 110000,
    "status": "Pending"
  }
]
```

### Approve Loan

`POST /loans/:id/approve` — **Admin**

Response:

```json
{
  "message": "loan approved successfully"
}
```

### Reject Loan

`POST /loans/:id/reject` — **Admin**

Response:

```json
{
  "message": "loan rejected successfully"
}
```

### Repay Loan

`POST /loans/:id/repay` — **Protected**

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

## Project Structure

```text
banking-system/
├── cmd/
│   └── main.go
├── config/
│   └── database.go
├── handlers/
│   ├── account_handler.go
│   ├── auth_handler.go
│   ├── bank_handler.go
│   ├── branch_handler.go
│   ├── customer_handler.go
│   ├── loan_handler.go
│   └── transaction_handler.go
├── middleware/
│   ├── admin_middleware.go
│   └── jwt_middleware.go
├── models/
│   ├── account.go
│   ├── bank.go
│   ├── branch.go
│   ├── customer.go
│   ├── loan.go
│   └── transaction.go
├── repository/
│   ├── account_repository.go
│   ├── bank_repository.go
│   ├── branch_repository.go
│   ├── customer_repository.go
│   ├── loan_repository.go
│   └── transaction_repository.go
├── routes/
│   └── routes.go
├── services/
│   ├── account_service.go
│   ├── auth_service.go
│   ├── bank_service.go
│   ├── branch_service.go
│   ├── customer_service.go
│   ├── loan_service.go
│   └── transaction_service.go
├── utils/
│   └── jwt.go
├── docs/
├── migrations/
├── scripts/
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── render.yaml
└── README.md
```

## Environment Variables

Example `.env` values:

```env
DATABASE_URL=postgres://user:password@localhost:5432/banking_db?sslmode=disable
JWT_SECRET=your_jwt_secret
PORT=8080
```

`DATABASE_URL` and `JWT_SECRET` are required. `PORT` defaults to `8080` when omitted.

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
6. Verify that it's running:
   ```bash
   curl http://localhost:8080/
   ```

The server automatically creates or updates the tables for Bank, Branch, Customer, Account, Transaction, and Loan on startup.

## Author

**Tushar Gupta**

B.Tech Computer Science & Engineering, Delhi Technological University (DTU)
