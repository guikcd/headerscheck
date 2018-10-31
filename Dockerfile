FROM golang:alpine AS builder
WORKDIR /go/src/github.com/guidelacour/headerscheck
COPY main.go .
RUN apk update && apk add git
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o headerscheck .

FROM alpine
WORKDIR /root
COPY --from=builder /go/src/github.com/guidelacour/headerscheck/headerscheck .
CMD ["./headerscheck"]
