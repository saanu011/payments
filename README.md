# Pismo Payments

- [Pismo Payments](#pismo-payments)
  - [Requirements](#Requirements)
  - [How to run](#how-to-run)
  - [How to test](#how-to-test)
  - [How to use](#how-to-use)
  - [Endpoints](#endpoints)
    - [Create Account](#create-account)
    - [Get Account](#get-account)
    - [Create Transaction](#create-transaction)
  - [Architecture](#architecture)
  - [Database](#database)
  - [Tests](#tests)

## Requirements

Have you ever thought about how a payment system works? This project is a simple implementation of a payment system. It allows you to create accounts and transactions.

## How to run

To run the project, you need to have Docker installed. Then, you can run the following command:

```bash
./bin/app_start
```

What this command does?
1. copies configs from sample config file to a new file
2. install dependencies (if there is some missing dependencies please run `go mod tidy`)
3. starts necessary services (database, etc)
4. runs the migration (we have a simple migration to create tables)
5. starts the application

## How to test

To test the project, you need to have Docker installed. It will help with DB related unit tests.
You can run the following command:

```bash
./bin/test_unit
```

## Endpoints

### Create Account

```azure
POST /accounts

request: {"email": "john.doe@mail.com", "password": "secret password"}
```

### Get Account

```azure
GET /accounts/:account_id

response: {"id": "some_id", "email": "john.doe@mail.com", "password": "secret password"}
```

### Create Transaction

```azure
POST /transaction

request: {"account_id": "account id", "operation_type_id": 1, "amount": 105.2}
```