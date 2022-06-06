FROM golang:1.16-alpine as builder

WORKDIR /go/src
COPY main.go go.mod go.sum ./
COPY flights/ ./flights/

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o server main.go


FROM alpine

WORKDIR /app
COPY --from=builder /go/src/server ./server

EXPOSE 9000
CMD ["./server"]