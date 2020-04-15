FROM golang:alpine AS build-environment

WORKDIR /app/

# Install dependencies first for cache optimization
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server cmd/server/server.go

# This stage copies the built binary for a smaller image
FROM alpine AS runtime-environment

COPY --from=build-environment ./app/server ./app/server

EXPOSE 8000

CMD ["./app/server"]
