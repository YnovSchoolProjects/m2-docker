FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app/app .
RUN chmod +x ./entrypoint.sh


FROM alpine:latest

RUN apk add --no-cache bash
COPY --from=builder /app ./

EXPOSE 8080
CMD ["./entrypoint.sh"]
