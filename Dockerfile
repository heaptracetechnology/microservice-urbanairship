FROM golang

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/microservice-urban-airship

ADD . /go/src/github.com/heaptracetechnology/microservice-urban-airship

RUN go install github.com/heaptracetechnology/microservice-urban-airship

ENTRYPOINT microservice-urban-airship

EXPOSE 3000