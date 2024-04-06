# Build.
FROM golang:1.22.2 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /entrypoint

EXPOSE 8080

# Run
CMD ["/entrypoint"]