FROM golang

ADD . /go/src/fil

WORKDIR /go/src/fil

RUN go build main.go