# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.15.0-alpine

WORKDIR /app
COPY . /app

RUN apk update \
  && apk upgrade -U -a\
  && apk add git \
  && apk --no-cache add curl
RUN go get github.com/cespare/reflex

CMD reflex -r '(\.go$|go\.mod)' -s -- sh -c 'go run cmd/app/main.go'
