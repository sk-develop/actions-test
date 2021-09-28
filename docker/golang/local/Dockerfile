FROM golang:1.15.5-alpine
RUN apk update && apk add git make gcc g++
WORKDIR /go/src/app
RUN go get -u github.com/cosmtrek/air
COPY go.mod .
COPY go.sum .
COPY . .
CMD while ! nc -z ${PSQL_HOST} ${PSQL_PORT}; do sleep 1; done && \
    air -c docker/golang/local/air.toml
