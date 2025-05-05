# API Contract: Profile Manager

## Overview

The **Profile Manager** service is responsible for managing user identities in **Parnaz-Eventify**, including registration, user validation for login, profile updates, and deletions. This service supports core user lifecycle actions and integrates with the API Gateway to validate credentials.

## Protocol

* **Type**: REST
* **Data Format**: HTTP/JSON
* **Base URL**: `http://127.0.0.1:8083`

## Endpoints

### 1. Register User

* **Method**: POST
* **Path**: `/users/register`
* **Description**: Registers a new user in the system.
* **Inputs**:

  ```json
  {
    "username": "string", // Required, unique username
    "password": "string", // Required, encrypted or hashed
    "email": "string",    // Optional
    "fullName": "string"   // Optional
  }
  ```
* **Outputs**:

  * **201 Created**:

    ```json
    {
      "userId": "user_123",
      "username": "parnaz123"
    }
    ```
* **Errors**:

  * **400 Bad Request**: Missing or invalid fields.
  * **409 Conflict**: Username already exists.

---

### 2. Validate Credentials (for Login)

* **Method**: POST
* **Path**: `/users/validate`
* **Description**: Validates username and password during login (called by API Gateway).
* **Inputs**:

  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
* **Outputs**:

  * **200 OK**:

    ```json
    {
      "valid": true,
      "userId": "user_123"
    }
    ```
  * **401 Unauthorized**:

    ```json
    {
      "valid": false,
      "error": "Invalid credentials"
    }
    ```

---

### 3. Get User by ID

* **Method**: GET
* **Path**: `/users/{userId}`
* **Description**: Retrieves basic profile info using user ID.
* **Outputs**:

  * **200 OK**:

    ```json
    {
      "userId": "user_123",
      "username": "parnaz123",
      "email": "parnaz@example.com",
      "fullName": "Parnaz P."
    }
    ```
  * **404 Not Found**: User not found.

---

### 4. Update User Info

* **Method**: PUT
* **Path**: `/users/{userId}`
* **Description**: Updates a user's profile information.
* **Body**:

  ```json
  {
    "email": "string",
    "fullName": "string"
  }
  ```
* **Outputs**:

  * **200 OK**:

    ```json
    {
      "userId": "user_123",
      "email": "new@example.com",
      "fullName": "New Name"
    }
    ```
  * **404 Not Found**: User not found.

---

### 5. Delete User

* **Method**: DELETE
* **Path**: `/users/{userId}`
* **Description**: Deletes a user account.
* **Outputs**:

  * **200 OK**:

    ```json
    {
      "deleted": true,
      "userId": "user_123"
    }
    ```
  * **404 Not Found**: User not found.

---
