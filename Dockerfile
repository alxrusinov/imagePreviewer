FROM golang:1.19-alpine as build

COPY . /go/src

WORKDIR /go/src/cmd/app

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

FROM alpine:latest

COPY --from=build server .

EXPOSE 80

ENTRYPOINT ["/server"]
