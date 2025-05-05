# Wallet Manager Contract

## Overview

Wallet Manager handles wallets and transactions in Parnaz-Eventify, including lookup, payments, rollbacks, balance, history, and deposits.

## Protocol

- Type: REST
- Data Format: HTTP/JSON
- Base URL: http://localhost:8082

## Route Group

- /wallet: Wallet management endpoints.
## Endpoints

### 1. Get Wallet ID by User ID
- **Method**: GET  
- **Path**: `/wallet/user/{userId}`  
- **Description**: Retrieves or creates a wallet associated with the specified user ID.
- **Inputs**:  
  - **Path Parameter**: `userId` – User’s unique identifier (required)
- **Outputs**:
  - **200 OK**:
    ```json
    {
      "userId": "user_123",
      "walletId": "wallet_abc"
    }
    ```
- **Errors**:
  - **404 Not Found**:
    ```json
    {
      "error": "User not found"
    }
    ```

---

### 2. Add Balance (Deposit)
- **Method**: POST  
- **Path**: `/wallet/{walletId}/deposit`  
- **Description**: Adds funds to the specified wallet.
- **Inputs**:
  - **Path Parameter**: `walletId` – Wallet’s unique identifier (required)
  - **Request Body**:
    ```json
    {
      "amount": 100000,
      "referenceId": "txn_001",
      "description": "Charge from gateway"
    }
    ```
- **Outputs**:
  - **200 OK**:
    ```json
    {
      "walletId": "wallet_abc",
      "newBalance": 150000
    }
    ```
- **Errors**:
  - **400 Bad Request**:
    ```json
    {
      "error": "Invalid request body"
    }
    ```

---

### 3. Pay 
- **Method**: POST  
- **Path**: `/wallet/{walletId}/pay`  
- **Description**: Deducts amount from wallet for a purchase.
- **Inputs**:
  - **Path Parameter**: `walletId` – Wallet’s unique identifier (required)
  - **Request Body**:
    ```json
    {
      "amount": 50000,
      "referenceId": "txn_002",
      "description": "Purchase gift card"
    }
    ```
- **Outputs**:
  - **200 OK**:
    ```json
    {
      "walletId": "wallet_abc",
      "newBalance": 100000
    }
    ```
- **Errors**:
  - **400 Bad Request**:
    ```json
    {
      "error": "Insufficient balance"
    }
    ```

---

### 4. Rollback Transaction
- **Method**: POST  
- **Path**: `/wallet/{walletId}/rollback`  
- **Description**: Reverses a transaction based on its reference ID.
- **Inputs**:
  - **Path Parameter**: `walletId` – Wallet’s unique identifier (required)
  - **Request Body**:
    ```json
    {
      "referenceId": "txn_002",
      "reason": "Gift card service failed"
    }
    ```
- **Outputs**:
  - **200 OK**:
    ```json
    {
      "walletId": "wallet_abc",
      "rolledBack": true,
      "newBalance": 150000
    }
    ```
- **Errors**:
  - **400 Bad Request**:
    ```json
    {
      "error": "Invalid reference ID"
    }
    ```

---

### 5. Get Wallet Balance
- **Method**: GET  
- **Path**: `/wallet/{walletId}/get-balance`  
- **Description**: Returns the current balance of the wallet.
- **Inputs**:
  - **Path Parameter**: `walletId` – Wallet’s unique identifier (required)
- **Outputs**:
  - **200 OK**:
    ```json
    {
      "walletId": "wallet_abc",
      "balance": 150000
    }
    ```

---

### 6. Get Last Transactions
- **Method**: GET  
- **Path**: `/wallet/{walletId}/get-transactions`  
- **Query**: `limit` (default: 10)  
- **Description**: Fetches the last N transactions of a wallet.
- **Inputs**:
  - **Path Parameter**: `walletId` – Wallet’s unique identifier (required)
  - **Query Parameter**: `limit` – Number of transactions to return (optional)
- **Outputs**:
  - **200 OK**:
    ```json
    [
      {
        "txnId": "txn_001",
        "type": "deposit",
        "amount": 100000,
        "timestamp": "2025-05-04T09:00:00Z"
      },
      {
        "txnId": "txn_002",
        "type": "pay",
        "amount": 50000,
        "timestamp": "2025-05-04T10:00:00Z"
      }
    ]
    ```