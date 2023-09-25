# Use the official Golang image to create a build artifact.
FROM golang:1.21 as builder

# Set the current working directory inside the container.
WORKDIR /app

# Copy go mod and sum files to the workspace.
COPY go.* ./

# Download dependencies.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main

# Use a lightweight Alpine image for the final image.
FROM alpine:latest

# Install certificates to communicate over HTTPS
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage.
COPY --from=builder /go/bin/main /go/bin/main

# Copy the static files and templates
COPY static static/
COPY templates templates/

# Set the binary as the entry point of the container.
ENTRYPOINT ["/go/bin/main"]

