# build book application first
FROM golang:alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o book-service .

# running application
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /build/book-service .

COPY ./config/app/.env .
COPY ./config/app/app.conf .

CMD ["./book-service"]