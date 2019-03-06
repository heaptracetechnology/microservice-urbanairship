FROM golang

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/microservice-urbanairship

ADD . /go/src/github.com/heaptracetechnology/microservice-urbanairship

RUN go install github.com/heaptracetechnology/microservice-urbanairship

ENTRYPOINT microservice-urbanairship

EXPOSE 3000