# NexaForm

## Project Overview

This project implements a flexible, secure, and customizable online questionnaire system. Built using **Go**, the system focuses on:

- **User management** with secure authentication and role-based access control.
- **Dynamic questionnaire creation** with various question types and customization options.
- **Real-time updates** and **data analytics** to monitor questionnaire responses.
- **Robust security** measures, including OTP verification, JWT/OAuth, and encrypted data storage.

## Features

### 1. User Authentication and Security
- **Registration and Login:** Users register using their email and national ID.
- **Authentication:** Secure login using JWT/OAuth2 with support for two-factor authentication (2FA).
- **Role-Based Access Control (RBAC):** Differentiated permissions for admins, creators, and participants.

### 2. Questionnaire Management
- **Question Types:** Supports text, multiple-choice, conditional logic, and file uploads.
- **Customizability:** Allows control over anonymity, time limits, and randomization of questions.
- **User Roles:** Ownership-based access and temporary role delegation.

### 3. Response and Analytics
- **Real-Time Updates:** Uses Server-Sent Events (SSE) or WebSockets for live monitoring.
- **Analytics:** Generates insights like response times, average completion rates, and correctness percentages.

### 4. Notifications and Interactions
- **Event Notifications:** Alerts for changes in questionnaires, role updates, or canceled responses.
- **Chat Rooms:** Enables discussion among participants with private and group options.

### 5. Data and Security Management
- **Data Persistence:** Utilizes MySQL for structured storage of questionnaire data.
- **Logging and Monitoring:** Comprehensive structured logging for system transparency.
- **Encryption:** Secure handling of sensitive data like responses and user credentials.

## Architecture and Technical Details

### Tech Stack
- **Backend:** Go
- **Database:** MySQL
- **Authentication:** JWT/OAuth2 with 2FA support
- **Real-Time Communication:** SSE or WebSockets
- **Containerization:** Docker for seamless deployment

### Design Principles
- **Modular Architecture:** Ensures scalability and maintainability.
- **Secure Coding Practices:** Implements rate limiting, encrypted communication, and input validation.
- **Standard Logging Formats:** Logs are stored in JSON for easy indexing and monitoring.

### Patterns Used
- Dependency Injection
- Factory Design Pattern
- Repository Pattern

## Installation

### Prerequisites
- Go 1.20+
- MySQL 8.0+
- Docker (optional, for containerized deployment)

### Setup Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/online-questionnaire.git
   cd online-questionnaire
   ```
2. Configure the database in the `.env` file:
   ```env
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_NAME=questionnaire
   ```
3. Run the migrations:
   ```bash
   go run migrate.go
   ```
4. Start the server:
   ```bash
   go run main.go
   ```

## Testing
- **Unit Tests:** Execute `go test ./...` to run unit tests for all modules.
- **End-to-End Tests:** Use `Postman` or `cURL` to test API endpoints.

## Contribution Guidelines
1. Fork the repository and create a new branch for your feature.
2. Ensure your code adheres to the project's modular architecture.
3. Submit a pull request with detailed descriptions of your changes.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

---

Feel free to customize the README further to match the exact requirements or branding of your project.
