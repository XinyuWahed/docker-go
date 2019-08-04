FROM golang:alpine as builder

ENV GOPATH /go

RUN apk update && apk add git

WORKDIR $GOPATH/src/github.com/XinyuWahed/simplewebgo

ADD . .

RUN go get -d -v && go build -o /app ./main.go ./handle.go

# RUN STAGE #
FROM alpine:edge

COPY --from=builder /app /app

CMD exec /app
