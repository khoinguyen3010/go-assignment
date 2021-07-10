FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get -v ./...
EXPOSE 5050

CMD [ "go", "run", "main.go" ]