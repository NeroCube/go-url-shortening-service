FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/go-redis/redis
RUN go get -u github.com/gorilla/mux
RUN go install -v ./...

CMD ["app"]