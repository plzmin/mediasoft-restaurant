FROM golang:1.20-alpine3.17 AS Builder
RUN apk add git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=Builder /app/main .
COPY .env .

EXPOSE 8080
CMD ["/app/main"]