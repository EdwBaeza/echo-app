FROM golang:1.15-alpine AS build-env

WORKDIR /app
COPY . /app
RUN go build cmd/app/main.go

FROM alpine:latest
RUN apk update  && apk add ca-certificates &&  rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /app/main /app
EXPOSE 8080
CMD ./main
