FROM golang:alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o my-web .

EXPOSE 8080

CMD ["/build/my-web"]
