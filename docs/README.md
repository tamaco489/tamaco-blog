# Documentation

This directory contains project documentation.

## Table of Contents

- [Technology Stack](./tech-stack.md) - Detailed technology stack information
- [Development Rules](../.cursor/rules/dev-rules/globals.mdc) - Development rules and guidelines

## Architecture Overview

![Architecture Diagram](./architecture.png)

## Overview

This project is a blog system built with Next.js. It provides hosting through AWS S3 and CloudFront, data management through Supabase, and CMS functionality through Strapi.

For detailed technical architecture, please refer to [Technology Stack](./tech-stack.md).

## Development Environment Setup

### Infrastructure Development

Execute infrastructure commands from the `infra/ecr` directory.

```bash
cd infra/ecr
make init AWS_PROFILE=xxx ENV=xxx   # Initialize Terraform workspace
make plan AWS_PROFILE=xxx ENV=xxx   # Plan infrastructure changes
make apply AWS_PROFILE=xxx ENV=xxx  # Apply infrastructure changes
make ci-check       # Run quality checks (format + validate + lint + security)
```

### Frontend Development

Execute frontend commands from the `frontend` directory.

```bash
cd frontend
npm ci      # Install dependencies
npm run dev # Start development server
```

### Backend Development

Execute backend commands from the `backend/api/article` directory.

#### Initial Setup and Launch

```bash
cd backend/api/article
make install-tools # Create environment files and install development tools
make up            # Start Docker environment (including PostgreSQL)
make migrate-up    # Run database migrations
make load-masters  # Load master data
make logs          # Output API server logs (startup confirmation)
```

#### OpenAPI Development

```bash
cd backend/api/article
make bundle-openapi # Bundle OpenAPI specification
make gen-api        # Generate API interfaces and type definitions
```

#### Database Development Workflow

##### 1. Adding Migration Files

```bash
# Create new migration file
make migrate-create NAME=create_users_table

# Apply and verify migrations
make migrate-up     # Apply migrations
make migrate-status # Check status
make migrate-down   # Rollback (when necessary)
```

##### 2. Adding sqlc Query Files

```bash
# Create query file (internal/repository/queries/table_name.sql)
# Generate sqlc code
make gen-sqlc
```

##### 3. Loading Master Data

```bash
# Load master data
make load-masters   # Load master data
make reset-masters  # Delete all and reload master data

# Create new master file (scripts/masters/number_table_name.sql)
```
