
# Go Server Project

## Overview
This project is a simple Go web server that serves static HTML files and handles HTTP requests. It includes a form for submitting user data (name and address) via a POST request and a basic endpoint for a GET request. The server is built using Go's `net/http` package and serves files from a `static` directory.

## Project Structure
The project contains the following files:
- **main.go**: The main Go application file that defines the server, including:
  - A file server to serve static HTML files from the `static` directory.
  - A `/form` endpoint to handle POST requests from a form submission.
  - A `/hello` endpoint to handle GET requests and return a simple greeting.
- **static/index.html**: A basic static webpage with a heading, served as the default page at the root URL (`/`).
- **static/form.html**: An HTML form with fields for "Name" and "Address," which submits data to the `/form` endpoint via a POST request.

## Prerequisites
- **Go**: Ensure Go (version 1.22 or later) is installed. Verify with:
  ```bash
  go version
  ```
- A web browser to test the server and form submission.
- A terminal to run commands.

## Setup Instructions
1. **Navigate to the Project Directory**:
   Change to the project directory:
   ```bash
   cd /Users/ankitakapadiya/Documents/code/Golang/go-server
   ```

2. **Initialize a Go Module**:
   Since the project uses Go modules, initialize one if not already present:
   ```bash
   go mod init go-server
   ```
   This creates a `go.mod` file to manage dependencies. The module name `go-server` is used for simplicity.

3. **Create the Static Directory**:
   Ensure a `static` directory exists in the project root to store HTML files:
   ```bash
   mkdir static
   ```
   Place the `index.html` and `form.html` files in the `static` directory.

4. **Verify Dependencies**:
   The project uses only the standard `net/http` package, so no external dependencies are required. Run the following to ensure the `go.mod` file is up-to-date:
   ```bash
   go mod tidy
   ```

## Running the Server
1. **Run the Server**:
   Execute the following command to start the server:
   ```bash
   go run main.go
   ```
   The server will start on port 8080, and you should see the message:
   ```
   Starting server at port 8080
   ```

2. **Build the Executable (Optional)**:
   To create a binary executable:
   ```bash
   go build
   ```
   This generates an executable (e.g., `go-server` on macOS/Linux). Run it with:
   ```bash
   ./go-server
   ```

## Testing the Application
1. **Access the Static Page**:
   - Open a browser and navigate to `http://localhost:8080/`.
   - You should see a page with the heading "Static Website" from `index.html`.

2. **Access the Form**:
   - Navigate to `http://localhost:8080/form.html`.
   - The page displays a form with fields for "Name" and "Address" and a "Submit" button.
   - Enter values in the fields and click "Submit."
   - The server will respond with a message like:
     ```
     POST request successful
     Name = <entered_name>
     Address = <entered_address>
     ```

3. **Test the Hello Endpoint**:
   - Navigate to `http://localhost:8080/hello` in your browser.
   - You should see the response:
     ```
     hello!
     ```
   - If you try a different path (e.g., `/invalid`) or use a non-GET method, you’ll receive a "404 not found" or "method is not supported" error.

4. **Test with cURL (Optional)**:
   - Test the `/hello` endpoint:
     ```bash
     curl http://localhost:8080/hello
     ```
     Expected output: `hello!`
   - Test the `/form` endpoint with a POST request:
     ```bash
     curl -X POST -d "name=John&address=123 Main St" http://localhost:8080/form
     ```
     Expected output:
     ```
     POST request successfulName = John
     Address = 123 Main St
     ```

## Troubleshooting
- **Server Fails to Start**:
  - Ensure port 8080 is not in use by another process:
    ```bash
    lsof -i :8080
    ```
  - If the port is in use, terminate the conflicting process or change the port in `main.go` (e.g., `:8081`).
- **404 Errors**:
  - Verify that `index.html` and `form.html` are in the `static` directory.
  - Ensure the `/hello` endpoint is accessed with a GET request.
- **Module Errors**:
  - If you see "cannot find main module," ensure the `go.mod` file exists by running `go mod init go-server`.
- **Form Submission Issues**:
  - Confirm that the form’s `action` attribute points to `/form` and uses `method="POST"`.
  - Check that the `<input>` fields have correct `name` attributes (`name` and `address`).

## Accessibility Notes
- The `form.html` file includes proper `<label>` elements associated with `<input>` fields using the `for` and `id` attributes, ensuring compliance with accessibility standards (e.g., WCAG 2.1).
- Test the form with a screen reader (e.g., VoiceOver on macOS) to verify that labels are correctly announced.

## Stopping the Server
To stop the server, press `Ctrl+C` in the terminal where it’s running.

## Notes
- The server uses the standard Go `net/http` package, making it lightweight and dependency-free.
- The form submission is basic and does not include data validation or storage. For a production environment, consider adding input validation and a database.
- To host this project publicly, consider using a reverse proxy (e.g., Nginx) and securing the server with HTTPS.

