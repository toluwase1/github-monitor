# GitHub Monitor Project

The **GitHub Monitor** project consists of three services: `commit-fetcher`, `commit-monitor`, and `github-service`. The purpose of this project is to fetch commits from a GitHub repository, monitor the repository for new commits, and provide API services for querying commit data. The services use RabbitMQ for messaging and PostgreSQL for data storage.

## Table of Contents
1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Environment Setup](#environment-setup)
4. [Services](#services)
   - [Commit Fetcher](#commit-fetcher)
   - [Commit Monitor](#commit-monitor)
   - [GitHub Service](#github-service)
5. [Running the Project](#running-the-project)
6. [API Documentation](#api-documentation)
7. [Contributing](#contributing)
8. [License](#license)

## Overview

This project fetches commits from a specified GitHub repository, monitors it for new commits, and provides a REST API for accessing commit data. The architecture leverages RabbitMQ for signaling between services and PostgreSQL as the database.

## Architecture

- **commit-fetcher**: Fetches commits from the specified repository for a given date range and sends a signal to RabbitMQ upon completion.
- **commit-monitor**: Listens for the completion signal and starts monitoring the repository for new commits.
- **github-service**: Provides REST API endpoints to access commit and repository data.

## Environment Setup

### Prerequisites

- Docker and Docker Compose
- Make

### Environment Variables

Create a `.env` file in the root of the project with the following content:

```ini
# Postgres Config
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=github_monitor

# GitHub Config
GITHUB_TOKEN=your_token

# Date Config
START_DATE=2024-08-02
SINCE_DATE=2024-08-01
UNTIL_DATE=2024-08-04

# RabbitMQ Config
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
REPO_NAME=chromium
