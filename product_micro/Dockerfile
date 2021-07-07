FROM golang:1.15.13-alpine3.14
RUN apk add git
WORKDIR /go/src/app
COPY ./approduct /usr/local/go/src/approduct
COPY ./model /usr/local/go/src/model
COPY ./products /usr/local/go/src/products
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./main.go .
RUN go get -d -v /go/src/app
RUN go build -o main .

EXPOSE 4030
CMD ["./main"]