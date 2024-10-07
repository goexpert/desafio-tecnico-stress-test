FROM golang:1.22 AS builder

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o /tmp/stress-test main.go
RUN chmod +x /tmp/stress-test

FROM ubuntu:22.04

RUN apt update && apt install openssl ca-certificates --yes

COPY --from=builder /tmp/stress-test /stress-test

ENTRYPOINT [ "/stress-test" ]

CMD []