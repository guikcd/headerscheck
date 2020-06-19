FROM golang:alpine AS builder
WORKDIR /go/src/github.com/guidelacour/headerscheck
COPY main.go .
RUN apk add --no-cache git
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o headerscheck .

ROM alpine:3
WORKDIR /root
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/github.com/guidelacour/headerscheck/headerscheck .
CMD ["./headerscheck"]
