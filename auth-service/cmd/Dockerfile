FROM golang:1.24-alpine

WORKDIR /app

# Copy the entire project
COPY . .

# Build the auth service
RUN go build -o /app/auth-service ./auth-service/cmd/main.go
RUN chmod +x /app/auth-service

EXPOSE 1234

CMD ["/app/auth-service"]