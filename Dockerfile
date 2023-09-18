FROM golang:1.21.0-alpine as builder
LABEL authors="kain"

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod tidy && go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go


FROM alpine:3.17.2 as runtime

WORKDIR /app
COPY --from=builder /build/main /app
COPY settings.yaml /app
RUN chmod 755 -R /app

CMD ["./main", "server"]

EXPOSE 9090