FROM golang:alpine AS builder
WORKDIR /go/src/github.com/guidelacour/headerscheck
COPY main.go .
RUN apk add --no-cache git
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o headerscheck .

FROM alpine:3.8
WORKDIR /root
COPY --from=builder /go/src/github.com/guidelacour/headerscheck/headerscheck .
CMD ["./headerscheck"]
