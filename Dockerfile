FROM golang:1.23.0-alpine
RUN apk add build-base
WORKDIR /app
COPY . .
COPY go.mod go.sum ./
RUN go mod download
COPY ./cmd/config.yml config.yml
# Build the Go app
RUN go build -o main ./cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 