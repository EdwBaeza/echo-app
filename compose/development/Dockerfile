FROM golang:1.13-alpine
WORKDIR /app
COPY . /app
RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8080
ENTRYPOINT CompileDaemon --build="go build cmd/app/main.go" --command=./main
