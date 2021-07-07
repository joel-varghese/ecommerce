FROM golang:1.15.13-alpine3.14
RUN apk add git
WORKDIR /go/src/app
COPY ./appuser /usr/local/go/src/appuser
COPY ./model /usr/local/go/src/model
COPY ./user /usr/local/go/src/user
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./main.go .
RUN go get -d -v /go/src/app
RUN go build -o main .

EXPOSE 3030
CMD ["./main"]