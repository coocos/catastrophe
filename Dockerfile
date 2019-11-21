FROM golang:alpine

WORKDIR app/

# Install dependencies first for cache optimization
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o catastrophe .

EXPOSE 8080

CMD ["./catastrophe"]
