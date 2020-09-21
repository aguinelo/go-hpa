FROM golang:1.14-alpine as builder

WORKDIR /go-hpa

COPY src/main.go main.go
COPY src/main_test.go main_test.go

ENV CGO_ENABLED=0

RUN go build 

FROM scratch

COPY --from=builder /go-hpa .

ENTRYPOINT ["./go-hpa"]