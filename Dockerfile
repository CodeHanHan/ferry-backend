FROM golang:1.16-alpine3.13 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ferry-main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/ferry-main .

COPY ./deploy/docker-compose/setting.yml /app/deploy/config/setting.yml

ENTRYPOINT ["./ferry-main"]

EXPOSE 10000

