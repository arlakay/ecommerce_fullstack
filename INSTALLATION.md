# React Frontend and Go Backend Project Setup Guide

This guide will walk you through setting up a project with a React frontend and Go backend on your local machine.

## Prerequisites

Ensure you have the following installed:
- PostgreSQL (for database)
- Node.js and npm (for React)
- Go (for backend)
- Git (for version control)

## Database Setup (Postgre)
1. 	Database credential:
    ```
    	Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "ecommerce",
		Port:         5432,
    ```

## Backend Setup (Go)

1. Go dependencies used:
   ```
   	github.com/gin-gonic/gin v1.10.0
	github.com/golang-jwt/jwt/v4 v4.5.0
	golang.org/x/crypto v0.27.0
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.12
   ```

2. Run the Go server:
   ```
   go run main.go
   ```

## Frontend Setup (React)

1. Get All dependencies
   ```
   npm install
   ```

4. Start the React development server:
   ```
   npm start
   ```

## Running the Project

1. In one terminal, run the Go backend:
   ```
   cd backend
   go run main.go
   ```

2. In another terminal, run the React frontend:
   ```
   cd ecommerce-frontend
   npm start
   ```

3. Open a web browser and navigate to `http://localhost:3000`. You should see the React app with the message from the Go backend.

