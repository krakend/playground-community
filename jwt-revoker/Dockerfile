FROM golang:1.12

RUN mkdir /app
WORKDIR /app
COPY . .

RUN GOPROXY=https://goproxy.io go get -d -v ./...
RUN GOPROXY=https://goproxy.io go install -v ./...
RUN pwd
RUN ls -l .

EXPOSE 8080

CMD ["jwt-revoker"]