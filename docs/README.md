# GitHub API Data Fetching and Service

This project implements a service in Golang to interact with GitHub's public APIs. It fetches data from the Chromium repository, stores it in a PostgreSQL database, and provides endpoints for accessing repository details, commits, and top authors.

## Table of Contents
1. [Features](#features)
2. [Setup](#setup)
3. [Usage](#usage)
    - [Fetch Commits](#fetch-commits)
    - [Fetch Repository Details](#fetch-repository-details)
    - [Reset Collection](#reset-collection)
    - [Get Top N Authors](#get-top-n-authors)
4. [Environment Variables](#environment-variables)
5. [Endpoints](#endpoints)
6. [Development](#development)
7. [License](#license)

## Features
- Fetches commit and repository data from the GitHub API.
- Stores data in a PostgreSQL database.
- Provides endpoints for retrieving commit history, repository details, and top authors by commit count.
- Supports resetting the collection of commits from a specific date.

## Setup

### Prerequisites
- Golang (v1.16+)
- PostgreSQL
- GitHub Personal Access Token

### Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/toluwase1/github-monitor.git
   cd github-monitor
2. **Create a .env file in the root directory with the following content:**

DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_username
DB_PASSWORD=your_password
DB_NAME=github_monitor
GITHUB_TOKEN=your_token
START_DATE=2021-01-01
3. Run the following
a. go mod tidy
go run cmd/github-monitor/main.go or make run

