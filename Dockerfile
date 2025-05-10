FROM golang:1.24.3-alpine3.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

RUN apk --no-cache add curl
EXPOSE 8888

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o nx-s3-server

# Run
CMD ["./nx-s3-server"]
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl -f http://localhost:8888/v1/health" ]