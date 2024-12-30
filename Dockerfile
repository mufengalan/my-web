FROM golang:alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .
RUN go mod download
RUN go app -o my-web ../cmd/main.go

EXPOSE 8080

CMD ["/app/my-web"]
