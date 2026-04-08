FROM golang:1.26 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM scratch
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]