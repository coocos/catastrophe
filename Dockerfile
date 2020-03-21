FROM golang:alpine AS build-environment

WORKDIR /app/

# Install dependencies first for cache optimization
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o catastrophe .

FROM alpine AS runtime-environment

COPY --from=build-environment ./app/catastrophe ./app/catastrophe

EXPOSE 8000

CMD ["./app/catastrophe"]
