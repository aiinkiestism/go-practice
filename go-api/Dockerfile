FROM golang:latest

RUN mkdir -p $GOPATH/src/go-api

WORKDIR $GOPATH/src/go-api

ADD . $GOPATH/src/go-api

RUN go get -u github.com/gin-gonic/gin && \
    go get github.com/zinzhu/gorm && \
    go get github.com/zinzhu/gorm/dialects/postgres
