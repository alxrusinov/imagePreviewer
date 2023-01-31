FROM golang:1.19 as build

COPY . /go/src

WORKDIR /go/src/cmd/app

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

FROM scratch

COPY --from=build server .

ENTRYPOINT ["/server"]
