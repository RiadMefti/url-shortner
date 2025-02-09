# URL Shortener

A fast and lightweight URL shortening service built with Go and SQLite. This service allows you to create shortened versions of long URLs, making them easier to share and manage.

# Deployed version 
[url-shortner](http://d440csc080g08ck4ock40o8k.46.202.176.17.sslip.io/)

## Features

- Create shortened URLs from long URLs
- Persistent storage using SQLite database
- Docker support for easy deployment
- Fast redirect response times
- Simple and clean REST API

## Prerequisites

Before you begin, ensure you have met the following requirements:
* Go 1.23.2 or higher
* Docker (if running containerized)
* SQLite3
* Git for version control

## Installation

### Local Development
```bash
# Clone the repository
git clone https://github.com/RiadMefti/url-shortner.git

# Change into the project directory
cd url-shortner

# Install dependencies
go mod download
```

### Docker Deployment
```bash
# Build the Docker image
docker build -t url-shortener .

# Run the container
docker run -p 8080:8080 url-shortener
```

## Usage

### API Endpoints

```bash
# Access it 
http://localhost:8080
```

### Running Locally
```bash
# Run the application
go run main.go

# Build the application
go build -o url-shortener main.go
```
