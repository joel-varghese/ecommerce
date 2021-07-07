FROM golang:1.15.13-alpine3.14
RUN apk add git
WORKDIR /go/src/app
COPY ./appcart /usr/local/go/src/appcart
COPY ./model /usr/local/go/src/model
COPY ./carts /usr/local/go/src/carts
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./main.go .
RUN go get -d -v /go/src/app
RUN go build -o main .

EXPOSE 4030
CMD ["./main"]