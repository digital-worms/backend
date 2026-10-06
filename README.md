<div align="center">

<img src="https://raw.githubusercontent.com/digital-worms/.github/main/assets/digital-worms-logo.png" width="220" alt="Digital Worms Logo">

# Digital Worms Backend

### Backend service for the Digital Worms platform

A Go backend for organizing events, managing shared expenses, storing memories and building collaborative features for groups of friends.

</div>

<p align="center">
  <img src="https://img.shields.io/badge/Go-Backend-00ADD8?logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/PostgreSQL-Database-4169E1?logo=postgresql&logoColor=white" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/Docker-Containerized-2496ED?logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/API-REST-6A5ACD" alt="REST API">
  <img src="https://img.shields.io/badge/Status-Active_Development-orange" alt="Status">
</p>

---

## About

Digital Worms is a social platform designed for groups of friends.

The goal is to bring events, participants, shared expenses, media, notifications and other group activities into one application.

This repository contains the backend part of the platform.

---

## Tech Stack

- Go
- PostgreSQL
- Docker
- REST API
- Git
- Linux

---

## Architecture

The backend is being designed with a layered architecture and clear separation of responsibilities.

The project will contain layers for:

- HTTP handlers
- business logic
- repositories
- database models
- DTOs
- configuration
- middleware
- infrastructure

The architecture is being developed with maintainability and future scalability in mind.

---

## Planned Modules

### Authentication

- User registration
- Login
- JWT authentication
- Refresh tokens
- User profiles

### Events

- Event creation
- Participants
- Roles
- Event information
- Activity planning

### Shared Expenses

- Expense creation
- Participants
- Expense splitting
- Balance calculation

### Media

- Event photos
- Event videos
- Shared memories

### Notifications

- Event updates
- Invitations
- Expense notifications
- Activity notifications

---

## Future Ideas

Digital Worms is planned as a long-term project.

Future functionality may include:

- Real-time messenger
- AI-powered features
- Redis caching
- gRPC
- Kafka
- Music and DJ-set management
- Games and group activities
- Travel planning
- CI/CD
- Kubernetes deployment

---

## Project Status

🚧 **Active development**

The backend architecture and core functionality are currently being designed and implemented.

---

## Development

### Database migrations

The Goose CLI is managed by the Go module, so no separate global installation is needed.

1. Copy `.env.example` to `.env` and adjust the local database credentials if needed.
2. Start PostgreSQL with Docker Compose:

   ```sh
   docker compose up -d postgres
   ```

3. Validate the migration files and inspect the database migration status:

   ```sh
   go tool goose -dir ./migrations validate
   go tool goose -dir ./migrations status
   ```

4. Apply pending migrations:

   ```sh
   go tool goose -dir ./migrations up
   ```

Goose reads the connection settings from `.env` (`GOOSE_DRIVER` and `GOOSE_DBSTRING`).
`validate` checks the migration files; `status` and `up` connect to PostgreSQL.

To roll back the latest migration in a local development database, use
`go tool goose -dir ./migrations down`. This runs the migration's `Down` section and
can remove database objects and their data, so do not use it against a database whose
data you need to keep.

---

<div align="center">

### Digital Worms

**Friends. Events. Expenses. Memories.**

</div>
