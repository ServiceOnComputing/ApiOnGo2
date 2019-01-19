FROM golang:latest

WORKDIR /go/src/apiongo
COPY . /go/src/apiongo
RUN go get && go build
RUN go build -o main .
CMD ["/go/src/apiongo/main"]

