FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go build -o /go-full-api ./cmd/main.go
EXPOSE 8080
CMD ["/go-full-api"]
