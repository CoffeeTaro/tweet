FROM golang:1.11.1-alpine3.8 as builder

RUN apk --update --no-cache add git
WORKDIR /go/src/github.com/kotaru23/tweet
COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -o tweet


FROM alpine:3.8 as certs

RUN apk --update --no-cache add ca-certificates


FROM scratch

COPY --from=builder /go/src/github.com/kotaru23/tweet/tweet /go/bin/
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
CMD ["/go/bin/tweet"]
