# OTP-Based Authentication System

## Description

This project is an OTP-based authentication system built with **Go** and **MongoDB Cloud**. It includes functionalities such as user registration, OTP generation, validation, and user details update. The system also implements various validations and ensures the phone number is 10 digits long, preventing duplicate registrations.

## Features

1. **User Registration**
   - Validates phone numbers (10 digits).
   - Checks if the user is already registered.
   - Registers the user if the phone number is valid and not already in use.

2. **Login**
   - Generates an OTP for registered users.
   - Verifies if the user is already registered before sending an OTP.

3. **OTP Validation**
   - Verifies the OTP for login and account recovery.

4. **User Details Update**
   - Allows users to update their details.
   - Requires the phone number to be provided and valid.

5. **Rate Limiting**
   - Limits the number of registration requests (3 per minute) to prevent abuse.

## Tech Stack

- **Go (Golang)**: The backend language used for API development.
- **MongoDB Cloud**: Cloud-based database to store user data and OTP information.
- **Gin**: Web framework used for handling HTTP requests.
- **Golang Rate Limiting**: Prevents excessive requests to the registration API.

## Installation

### Prerequisites

- **Go 1.20+**
- **MongoDB Cloud Account**: Set up a MongoDB Atlas account and create a cluster.
- **Postman** or similar API testing tool.

### Steps to Run Locally

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
    
2. Install Go dependencies::
   ```bash
   go mod tidy

3. Set up your MongoDB Cloud connection string in a .env file:
    ```bash
    MONGO_URI=mongodb+srv://<username>:<password>@cluster.mongodb.net/authDB?retryWrites=true&w=majority
    PORT=8080

4. Run the Go application:
    ```bash
    go run main.go 

5. The Default server should be running on http://localhost:8080.


### Endpoints
-   POST /register: Registers a new user.
-   POST /login: Logs in a user by sending an OTP.
-   POST /validate-otp: Validates the OTP entered by the user.
-   GET /user/{mobileNumber}: Fetch user details by phone number.
-   PUT /user/{mobileNumber}: Updates user details.


# Testing with Postman

You can use **Postman** to test the following API endpoints for the OTP-Based Authentication System.

# API Documentation

## 1. Register User

- **URL**: `POST http://localhost:8080/register`
- **Method**: POST  
- **Headers**:  
  Content-Type: `application/json`
- **Body (JSON)**:
  ```json
  {
    "mobileNumber": "9058364966",
    "name": "Rohit Varshney",
    "deviceFingerprint": "abc123xyz"
  }
  ```
- **Description**: Registers a new user with a valid phone number, name, and device fingerprint.

---

## 2. Login

- **URL**: `POST http://localhost:8080/login`
- **Method**: POST  
- **Headers**:  
  Content-Type: `application/x-www-form-urlencoded`
- **Body (Form-Data)**:  
  ```
  mobileNumber=9058364966
  ```

---

## 3. Validate OTP

- **URL**: `POST http://localhost:8080/validate-otp`
- **Method**: POST  
- **Headers**:  
  Content-Type: `application/x-www-form-urlencoded`
- **Body (Form-Data)**:  
  ```
  mobileNumber=9058364966
  otp=123456
  ```

---

## 4. Get User Details

- **URL**: `GET http://localhost:8080/user/{mobileNumber}`  
  (Replace `{mobileNumber}` with the actual mobile number, e.g., `9058364966`)
- **Method**: GET  

---

## 5. Update User Details

- **URL**: `PUT http://localhost:8080/user/{mobileNumber}`  
  (Replace `{mobileNumber}` with the actual mobile number, e.g., `9058364966`)  
- **Method**: PUT  
- **Headers**:  
  Content-Type: `application/json`
- **Body (JSON)**:
  ```json
  {
    "name": "Rohit Varshney",
    "deviceFingerprint": "newFingerprint123"
  }
  ```
