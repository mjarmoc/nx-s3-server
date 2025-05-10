FROM golang:1.24.3-alpine3.21 AS build

# Set destination for COPY
WORKDIR /app

# Copy source code including subdirectories
COPY . .

RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o nx-s3-server

FROM alpine:3.21
COPY --from=build /app/nx-s3-server ./

RUN apk --no-cache add curl
EXPOSE 8888

# Run
CMD ["./nx-s3-server"]
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl -f http://localhost:8888/v1/health" ]