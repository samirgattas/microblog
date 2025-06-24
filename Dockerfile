FROM golang:1.24.4-alpine3.22

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -v -o main cmd/main.go

CMD ["/app/main"]