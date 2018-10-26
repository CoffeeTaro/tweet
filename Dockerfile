FROM golang:1.11.1-alpine3.8 as builder

WORKDIR /go/src/tweet/
COPY . .
RUN  go get .
RUN GOOS=linux go build -o tweet .


FROM scratch

COPY --from=builder /go/src/tweet/tweet /root/
CMD ["/root/tweet"]
