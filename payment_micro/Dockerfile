FROM golang:1.15.13-alpine3.14
RUN apk add git
WORKDIR /go/src/app
COPY ./appayment /usr/local/go/src/appayment
COPY ./model /usr/local/go/src/model
COPY ./payments /usr/local/go/src/payments
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./main.go .
RUN go get -d -v /go/src/app
RUN go build -o main .

EXPOSE 6030
CMD ["./main"]