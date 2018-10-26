FROM golang:1.11.1-alpine3.8 as builder

ENV GO111MODULE on
Run apk --update --no-cache add git gcc musl-dev
WORKDIR /go/src/tweet/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tweet .


FROM scratch

COPY --from=builder /go/src/tweet/tweet /root/
CMD ["/root/tweet"]
