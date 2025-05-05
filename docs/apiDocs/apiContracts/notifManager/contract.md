# API Contract: Notification Manager

## Overview
The **Notification Manager** handles the creation, delivery, and retrieval of notifications in **Parnaz-Eventify**. It stores and processes notifications using **Redis** for real-time processing.

## Protocol
- **Type**: REST
- **Data Format**: HTTP/JSON
- **Base URL**: `http://127.0.0.1:8085`

## Endpoints

### 1. Send Notification

- **Method**: POST  
- **Path**: `/notifications/send`  
- **Description**: Sends a notification to a user or a list of users.

- **Inputs**:
  - **Request Body**:
    ```json
    {
      "userId": ["user_123", "user_456"],   // List of user IDs
      "message": "string",                    // Message content of the notification
      "type": "string"                        // Type of notification (e.g., event, reminder)
    }
    ```
    Example:
    ```json
    {
      "userIds": ["user_123", "user_456"],
      "message": "Event 'Party' is starting soon!",
      "type": "event"
    }
    ```

- **Outputs**:
  - **200 OK**:
    ```json
    {
      "status": "notifications sent",
      "userIds": ["user_123", "user_456"]
    }
    ```
  - **400 Bad Request**: Missing required fields
    ```json
    {
      "error": "Invalid input"
    }
    ```

- **Errors**:
  - **500 Internal Server Error**: Issue with notification delivery

- **Security**:
  - Requires JWT token in the `Authorization` header (`Bearer <token>`)

---

### 2. Get Notifications for User

- **Method**: GET  
- **Path**: `/notifications/{userId}`  
- **Description**: Retrieves unread notifications for a specific user.

- **Inputs**:
  - **Path Parameter**: `userId` — ID of the user

- **Outputs**:
  - **200 OK**:
    ```json
    [
      {
        "notificationId": "notif_001",
        "message": "Event 'Party' is starting soon!",
        "timestamp": "2025-05-04T12:00:00Z",
        "status": "unread"
      },
      {
        "notificationId": "notif_002",
        "message": "Reminder: Don't forget your appointment.",
        "timestamp": "2025-05-04T13:00:00Z",
        "status": "unread"
      }
    ]
    ```

  - **404 Not Found**: No notifications found for the user
    ```json
    {
      "error": "No notifications found"
    }
    ```

- **Errors**:
  - **500 Internal Server Error**: Issue retrieving notifications

- **Security**:
  - Requires JWT token in the `Authorization` header (`Bearer <token>`)

---
