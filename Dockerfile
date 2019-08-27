FROM golang:1.12 as builder
WORKDIR /module
COPY ./fizzbuzz/go.mod /module/go.mod
COPY ./fizzbuzz/go.sum /module/go.sum
COPY fizzbuzz /module
RUN go build main.go
EXPOSE 3000
CMD ./main

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /module/main .
CMD ["./main"]