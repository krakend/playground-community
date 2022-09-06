FROM golang:1.17.11

RUN mkdir /app
WORKDIR /app
COPY . .

RUN GOPROXY=https://goproxy.io go get -d -v ./...
RUN go build .

EXPOSE 8080

CMD ["./jwt-revoker", "-server", "krakend_ce:1234"]
