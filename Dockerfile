FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm
WORKDIR /build
COPY . .
RUN go build -o bin/gokub-tester cmd/gokub-tester/main.go

FROM alpine:3
RUN apk --no-cache add ca-certificates
WORKDIR /dist
COPY --from=builder /build/bin/gokub-tester .
EXPOSE 8080
CMD ["/dist/gokub-tester"]
