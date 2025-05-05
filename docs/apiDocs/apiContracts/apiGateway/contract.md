# API Contract: API Gateway

## Overview
The API Gateway acts as the central entry point for all client requests, routing them to the appropriate microservices (e.g., Profile Management, Event Manager, Reserve Manager). It handles authentication, request routing, and response aggregation.

## Protocol
- **Type**: REST
- **Data Format**: HTTP/JSON
- **Base URL**: `http://127.0.0.1:8080`

## Endpoints

### 1. User Login
- **Method**: POST
- **Path**: `auth/login`
- **Description**: Authenticates a user with their username and password, returning a JWT token.
- **Inputs**:
  - **Request Body**:
    ```
    {
      "username": "string", // User's username (required)
      "password": "string"  // User's password (required)
    }
    ```
    Example:
    ```
    {
      "username": "parnaz123",
      "password": "mypassword"
    }
    ```
- **Outputs**:
  - **200 OK**:
    ```json
    {
      "token": "string" // JWT token for authentication
    }
    ```
    Example:
    ```json
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
    ```
- **Errors**:
  - **400 Bad Request**: Invalid request (e.g., missing username or password).
    ```json
    {
      "error": "Missing username or password"
    }
    ```
  - **401 Unauthorized**: Invalid credentials.
    ```json
    {
      "error": "Invalid username or password"
    }
    ```
- **Security**: None (public endpoint).


### 2. Reserve an Event
- **Method**: POST
- **Path**: `/events/reserve`
- **Description**: Reserves an event for the authenticated user and processes payment.
- **Inputs**:
  - **Request Body**:
    ```json
    {
      "eventId": "string" // Event ID to reserve (required)
    }
    ```
    Example:
    ```json
    {
      "eventId": "event_123"
    }
    ```
- **Outputs**:
  - **200 OK**:
    ```json
    {
      "reservationId": "string", // Reservation ID
      "eventId": "string"        // Event ID
    }
    ```
    Example:
    ```json
    {
      "reservationId": "res_456",
      "eventId": "event_123"
    }
    ```
- **Errors**:
  - **400 Bad Request**: Invalid request (e.g., missing event ID).
    ```json
    {
      "error": "Missing event ID"
    }
    ```
  - **401 Unauthorized**: Missing or invalid JWT token.
    ```json
    {
      "error": "Unauthorized: Invalid or missing token"
    }
    ```
  - **404 Not Found**: Event not found or not available.
    ```json
    {
      "error": "Event not found or not available"
    }
    ```
  - **500 Internal Server Error**: Payment failure or other internal error.
    ```json
    {
      "error": "Payment failed, reservation rolled back"
    }
    ```
- **Security**:
  - Requires JWT token in the `Authorization` header (`Bearer <token>`).