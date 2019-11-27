FROM golang

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/oms-services/urbanairship

ADD . /go/src/github.com/oms-services/urbanairship

RUN go install github.com/oms-services/urbanairship

ENTRYPOINT urbanairship

EXPOSE 3000