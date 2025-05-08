FROM golang:1.21

WORKDIR /app

# Install git
RUN apt-get update && apt-get install -y git

# Clean and clone the gotit-api-go-sdk repository
RUN rm -rf /app/gotit-api-go-sdk && \
    git clone -b main https://github.com/Dayone-Joint-Stock-Company/gotit-api-go-sdk.git /app/gotit-api-go-sdk

# Copy the example project files
COPY . .

# Debug: Show current directory contents
RUN ls -la

# Debug: Show go.mod contents
RUN cat go.mod

# Download and tidy dependencies
RUN go mod download && \
    go mod edit -replace gotit_api_go_sdk=./gotit-api-go-sdk && \
    go mod tidy

# Debug: Show go.mod after tidy
RUN cat go.mod

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

# Make sure the binary is executable
RUN chmod +x app

# Debug: Show the built binary
RUN ls -la app

# Expose the port
EXPOSE 5004

# Run the binary
CMD ["./app"]