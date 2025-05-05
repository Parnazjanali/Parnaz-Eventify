# Project Manifest: Evenntify

## 1. Problem Statement & Goal

### Problem
Organizing online events requires complex coordination, including user registration, event reservation, payment processing, and real-time reporting for organizers.

### Pain Points
- Users struggle to find a seamless platform to log in, reserve events, and make payments securely.
- Failed payments can lead to inconsistent reservations, causing frustration for users.
- Organizers lack real-time insights into reservations and payments, making event management inefficient.

### Solution
This project aims to build a scalable microservices-based system that allows users to log in with their username and password, reserve events, process payments securely, and provides organizers with real-time transaction reports. The system ensures consistency with rollback mechanisms for failed operations.

## 2. Project Overview

The Event Management System is a microservices-based platform designed to streamline online event reservations. Users can log in using their username and password, browse and reserve events, and make payments securely. Organizers can access real-time reports of reservations and transactions. The system ensures reliability by handling errors (e.g., payment failures) with rollback mechanisms.

## 3. Scope

### In Scope
- User login with username and password-based authentication.
- Event reservation and payment processing.
- Real-time transaction reporting for organizers.
- Error handling with rollback mechanisms (e.g., for payment failures).
- API documentation with Swagger.

### Out of Scope
- Support for cryptocurrency payments.
- Live chat between users.
- Mobile application development.

## 4. Architecture & Design

The system follows a microservices architecture with the following components:
- **API Gateway**: Central entry point for all user requests, handling routing and authentication.
- **Profile Management**: Manages user authentication (using username and password) and token generation.
- **Event Manager**: Handles event creation and availability checks.
- **Reserve Manager**: Manages event reservations and rollbacks.
- **Wallet Manager**: Processes payments and provides transaction logs.
- **Notif Manager**: Sends asynchronous notifications to users.

The services communicate via REST APIs using HTTP/JSON. Redis is used for caching tokens and asynchronous notifications.

## 5. Technical Stack

- **Programming Language**: Golang
- **API Framework**: Fiber (for REST APIs)
- **Database**: PostgreSQL (primary database for structured data)
- **Caching/Message Queue**: Redis (for token caching and Pub/Sub notifications)
- **Logging**: Zap (for structured logging)
- **Testing**: Go Testing (for unit tests)
- **Containerization**: Docker & Docker Compose
- **API Documentation**: Swagger

## 6. Assumptions

- Users log in using their username and password as the unique identifier.
- The system must be highly scalable and reliable to handle multiple concurrent users.
- Services are independent and communicate via REST APIs.
- Notifications are sent asynchronously to improve performance.

## 7. Constraints & Challenges

- Handling payment failures requires rollback mechanisms to ensure data consistency, which adds complexity.
- Real-time reporting may require caching (e.g., with Redis) to optimize performance.
- Ensuring security through proper token validation (JWT) across all services.

## 8. Non-Functional Requirements

- **Scalability**: The system should handle up to 1000 concurrent users.
- **Performance**: API response time should be under 200 milliseconds.
- **Security**: User data and transactions must be secured with JWT authentication.
- **Reliability**: The system must ensure data consistency with rollback mechanisms for failed operations.

## 9. Next Steps & Roadmap

- Design the database schema for PostgreSQL (e.g., User, Event, Reservation, Transaction).
- Implement microservices using Fiber and Golang (starting with API Gateway).
- Set up Redis for token caching and asynchronous notifications.
- Write unit tests for critical components (e.g., authentication, reservation) using Go Testing.
- Add structured logging with Zap for debugging and monitoring.
- Document APIs using Swagger.
- Deploy services using Docker Compose for local testing.

## 10. Appendices

- Call Flow diagrams for event reservation and real-time reporting are available in the project documentation.

* **Event Reservation Diagram : [click me](docs/rawFiles/EventReservation.png)**
* **Wallet Transaction Log : [click me](docs/rawFiles/walletTransactionLog.png)**