# Scheduler Billing Engine

## Description

The Scheduler Billing Engine is a loan management system built using Golang, Gin, and PostgreSQL with GORM. It provides functionalities for managing loans, tracking payments, and monitoring borrower statuses. The system allows for the creation of loans, making payments, and checking outstanding amounts and delinquency status.

## Features

- Create loans with specified amounts and interest rates.
- Schedule weekly payments for loans.
- Track outstanding amounts and payment statuses.
- Check if a borrower is delinquent based on missed payments.
- Soft delete functionality for loans and payments.
- Billing records associated with each payment.

## Table of Contents

- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Running the Application](#running-the-application)
- [Testing](#testing)
- [Contributing](#contributing)

## Tech Stack

- **Backend**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **CI/CD**: Docker

## Getting Started

### Prerequisites

- Go (version 1.23 or higher)
- PostgreSQL

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/balqisgautama/billing-engine.git
   cd your-folder-name
   ```

2. Install the necessary Go packages:

   ```bash
   go mod tidy
   ```

## Running the Application
To deploy the FinTech API using Docker and Nodemon, you can use the provided `nodemon.json`, `Dockerfile` and `docker-compose.yml`.

1. **Installing Nodemon**:
    In the root directory of your project, run:
    ```bash
    npm install -g nodemon
    ```

2. **Build and Run the Application**:
    In the root directory of your project, run:
    ```bash
    nodemon
    ```
    or
    ```bash
    make run
    ```
3. **Access the API**:
    The API will be accessible at `http://localhost:8080`.

## Testing

To run tests, use:

```bash
go test ./...
```
or
```bash
make test
```

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.