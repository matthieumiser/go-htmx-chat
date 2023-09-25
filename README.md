# go-htmx-chat

A minimalist project to demonstrate real-time chat functionality using the Go language and the htmx library.

## Overview

This project sets up a simple chat server where users can send and receive messages in real-time. It leverages the power of Go for the backend and htmx for the frontend to achieve this functionality without the need for a full-fledged frontend framework.

### Key Files:

1. [main.go](https://github.com/matthieumiser/go-htmx-chat/blob/main/main.go): This is the main server file written in Go. It sets up the routes, handles WebSocket connections, and broadcasts messages to all connected clients.
2. [index.html](https://github.com/matthieumiser/go-htmx-chat/blob/main/templates/index.html): The main HTML template that provides the chat UI. It uses htmx to handle real-time updates without requiring a page reload.
3. [styles.css](https://github.com/matthieumiser/go-htmx-chat/blob/main/static/styles.css): Contains the styling for the chat interface.

## Features:

- **Real-time Messaging**: Users can send and receive messages in real-time.
- **WebSocket Integration**: Uses WebSockets to establish a persistent connection between the client and server.
- **Minimalist Design**: The project focuses on functionality, keeping the design and codebase simple and clean.

## How to Run:

1. Clone the repository.
2. Navigate to the project directory.
3. Run the server using the command `go run main.go`.
4. Open a browser and navigate to `http://localhost:8080` to access the chat interface.
