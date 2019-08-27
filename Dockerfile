FROM golang:1.12 as builder
WORKDIR /module
COPY ./fizzbuzz/go.mod /module/go.mod
COPY ./fizzbuzz/go.sum /module/go.sum
COPY fizzbuzz /module
RUN GOOS=linux CGO_ENABLED=0 go build main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /golang-exec
COPY --from=builder /module/main .
EXPOSE 3000
CMD ["./main"]