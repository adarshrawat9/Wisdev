# Wisdev

Wisdev is a backend-first developer identity platform that aggregates a developer's presence across multiple platforms into a single shareable profile.

The goal is to help developers showcase their skills, projects, coding activity, and professional presence through one unified profile.


# Vision

Developers often have their profiles spread across multiple platforms:

GitHub
LeetCode
HackerRank
Codeforces
CodeChef
LinkedIn
Personal Portfolio

Wisdev aims to provide:

One Profile. One Link. All Developer Insights.

## Tech Stack

#  Backend
   - Go
   - Gin
   - PostgreSQL
   - pgxpool
   - JWT
   - bcrypt

# Infrastructure
   - Docker
   - Docker Compose
   - golang-migrate

# Architecture

The project follows a layered architecture:

    HTTP Request
        ↓
    Handler
        ↓
    Service
        ↓
    Repository
        ↓
    Database


## Layers

# Handler

Responsible for:

   - HTTP Requests
   - JSON Binding
   - Status Codes
   - Response Formatting

# Service

Responsible for:

   - Business Logic
   - Validation
   - Authentication Logic

# Repository

Responsible for:

   - SQL Queries
   - Data Access

# Database

Responsible for:

   - PostgreSQL Connection Management
   - Connection Pooling


# Setup
   1. Clone Repository
       git clone <repository-url>
       cd wisdev

    2. Configure Environment
        Create a .env file and add your database configuration.

    3. Start PostgreSQL
        docker compose up -d

    4. Run Migrations
        make migrate-up

               or

        migrate -path migrations -database <database-url> up
        
    5. Start API
        go run ./cmd/api
