# API Contract: Event Manager

## Overview
The Event Manager is responsible for managing events in the Parnaz-Eventify system. It handles retrieving the list of events, fetching details of a specific event, and checking event availability for reservations.

## Protocol
- **Type**: REST
- **Data Format**: HTTP/JSON
- **Base URL**: `http://127.0.0.1:8081`

## Route Groups
- **/events**: Endpoints related to event management.

## Endpoints

### 1. List Events
- **Method**: GET
- **Path**: `/events`
- **Description**: Retrieves a list of available events with basic information.
- **Inputs**: None
- **Outputs**:
  - **200 OK**:
    ```json
    [
      {
        "id": "string",    // Event ID
        "title": "string", // Event title
        "date": "string"   // Event date (ISO 8601 format)
      }
    ]
    ```
    Example:
    ```json
    [
      {
        "id": "event_123",
        "title": "Tech Conference 2025",
        "date": "2025-05-01T10:00:00Z"
      },
      {
        "id": "event_124",
        "title": "Music Festival 2025",
        "date": "2025-06-01T18:00:00Z"
      }
    ]
    ```
- **Errors**:
  - **401 Unauthorized**: Missing or invalid JWT token.
    ```json
    {
      "error": "Unauthorized: Invalid or missing token"
    }
    ```
  - **500 Internal Server Error**: Failed to fetch events.
    ```json
    {
      "error": "Internal server error"
    }
    ```
- **Security**: None(Public Endpoint)

### 2. Get Event Details
- **Method**: GET
- **Path**: `/events/{eventId}`
- **Description**: Retrieves detailed information about a specific event by its ID.
- **Inputs**:
  - **Path Parameter**:
    - `eventId`: The ID of the event (required).
      Example: `event_123`
- **Outputs**:
  - **200 OK**:
    ```json
    {
      "id": "string",           // Event ID
      "title": "string",        // Event title
      "date": "string",         // Event date (ISO 8601 format)
      "capacity": "integer",    // Total capacity
      "availableSeats": "integer" // Available seats
    }
    ```
    Example:
    ```json
    {
      "id": "event_123",
      "title": "Tech Conference 2025",
      "date": "2025-05-01T10:00:00Z",
      "capacity": 100,
      "availableSeats": 50
    }
    ```
- **Errors**:
  - **400 Bad Request**: Invalid event ID.
    ```json
    {
      "error": "Invalid event ID"
    }
    ```
  - **401 Unauthorized**: Missing or invalid JWT token.
    ```json
    {
      "error": "Unauthorized: Invalid or missing token"
    }
    ```
  - **404 Not Found**: Event not found.
    ```json
    {
      "error": "Event not found"
    }
    ```
  - **500 Internal Server Error**: Internal server error.
    ```json
    {
      "error": "Internal server error"
    }
    ```
- **Security**: None(Public Endpoint)

### 3. Check Event Availability
- **Method**: GET
- **Path**: `/events/{eventId}/availability`
- **Description**: Checks if an event is available for reservation (e.g., has available seats).
- **Inputs**:
  - **Path Parameter**:
    - `eventId`: The ID of the event (required).
      Example: `event_123`
- **Outputs**:
  - **200 OK**:
  ```json
  {
  "eventId": "string",      // Event ID
  "available": "boolean",   // Whether the event is available
  "status": "string",       // Availability status (e.g., "AVAILABLE", "NOT_AVAILABLE")
  "availableSeats": "integer" // Number of available seats
  }
  ```
  Example:
  ```json
  {
  
  "eventId": "event_123",
  "available": true,
  "status": "AVAILABLE",
  "availableSeats": 50

  }
  ````

- **Errors**:
  - **400 Bad Request**: Invalid event ID.
    ```json
    {
      "error": "Invalid event ID"
    }
    ```
  - **401 Unauthorized**: Missing or invalid JWT token.
    ```json
    {
      "error": "Unauthorized: Invalid or missing token"
    }
    ```
  - **404 Not Found**: Event not found.
    ```json
    {
      "error": "Event not found"
    }
    ```
  - **500 Internal Server Error**: Internal server error.
    ```json
    {
      "error": "Internal server error"
    }
    ```
- **Security**:
  - Requires JWT token in the `Authorization` header (`Bearer <token>`).