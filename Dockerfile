FROM golang

RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/parsing-service
WORKDIR /go/src/parsing-service

RUN dep ensure
RUN go build

EXPOSE 5001

CMD ./parsing-service
