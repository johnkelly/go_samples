FROM golang:latest
WORKDIR /go/src/github.com/johnkelly/server/
COPY main.go .
RUN go build -o app .

CMD ["./app"]
