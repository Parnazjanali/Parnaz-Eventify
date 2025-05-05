# API Contract: Reserve Manager

## Overview

The Reserve Manager in **Parnaz-Eventify** is responsible for managing event reservations and cancellations. While total capacity is maintained by the Event Manager service, Reserve Manager tracks actual reservations and determines available slots. It also handles logic for canceling reservations.

## Protocol

* **Type**: REST
* **Data Format**: HTTP/JSON
* **Base URL**: `http://127.0.0.1:8084`

## Endpoints

### 1. Reserve an Event

- **Method**: POST  
- **Path**: `/reservations`  
- **Description**: Reserves an event for the user if capacity is available.

- **Inputs**:
  - **Request Body**:
    ```json
    {
      "eventId": "string",  // ID of the event to reserve
      "userId": "string"    // ID of the user making the reservation
    }
    ```
    Example:
    ```json
    {
      "eventId": "event_123",
      "userId": "user_456"
    }
    ```

- **Outputs**:
  - **200 OK**:
    ```json
    {
      "reservationId": "res_789",
      "status": "reserved"
    }
    ```
  - **409 Conflict**: Not enough capacity
    ```json
    {
      "error": "Event is fully booked"
    }
    ```

- **Errors**:
  - **400 Bad Request**: Missing required fields
  - **404 Not Found**: Event not found
  - **500 Internal Server Error**: Unexpected failure during reservation

- **Security**:
  - Requires JWT token in the `Authorization` header (`Bearer <token>`)

---

### 2. Cancel Reservation

- **Method**: DELETE  
- **Path**: `/reservations/{reservationId}`  
- **Description**: Cancels a reservation and releases capacity.

- **Inputs**:
  - **Path Parameter**: `reservationId` — ID of the reservation to cancel

- **Outputs**:
  - **200 OK**:
    ```json
    {
      "reservationId": "res_789",
      "status": "cancelled"
    }
    ```

  - **404 Not Found**:
    ```json
    {
      "error": "Reservation not found"
    }
    ```

- **Errors**:
  - **400 Bad Request**: Invalid or missing reservation ID
  - **500 Internal Server Error**: Error during cancellation

- **Security**:
  - Requires JWT token in the `Authorization` header (`Bearer <token>`)

---
### Notes

Coordination with Event Manager is done to fetch the total capacity.

