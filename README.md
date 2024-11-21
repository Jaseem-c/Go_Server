# GO Api Server

## overview
This project is a simple GO Api server,designed to understand basic of building and running a go server .it demonstrates how to run a server,setup routes,handle requests,and return json responses.additionaly add a feature which show how to serve static front files like html,css etc.


## features

### 1. RESTful API Design
- Implements a simple API endpoint (`/api`) that represents:
  - How to create an Api server.
  - Handle HTTP requests using function "handleFunc" in `net/http` package
  - Return JSON responses using  `encoding/json` package.
  - Use structs to define and encode JSON response.

### 2.Struct Handling
  - Define custom structs.
  - Encode structs to JSON.

### 3.HTTP Header
  - configure HTTP headers (e.g., Content-Type: application/json) to ensure proper response format.
  - Encode structs to JSON. 

### 4.Static File Serving:
  - server can serve static files for front-end application(html,css,js) from `public` directory


## How it Works

 - API Endpoint (`/api`):
    - The /api route responds with a JSON object.
      {
         "text": "Front end connected to Go Api Server"
      }
    - Response is generated by struct

 - Static File Serving:
    - Files in the ./public directory (index.html) are served when the root (/) is accessed.


## How to Run the Project

### 1. clone Repository
  - git clone <repository-url>
  - cd <respository-directory>
### 2. Run the server
  - go run main.go
### 3. Access the Application
  - visit `http://localhost:8080/api` for the API response.

## Key Concept Demonstrated
 - Struct
 - HTTP server creation
 - Route handlers
 - JSON convertion
 - Serving Static Files  






 