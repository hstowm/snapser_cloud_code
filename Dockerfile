FROM golang:1.19.1-alpine3.16 AS build

WORKDIR /go/src

COPY  . .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o service .

FROM golang:1.19.1-alpine3.16 AS runtime
COPY --from=build /go/src/service ./

EXPOSE 8080/tcp
ENTRYPOINT ["./service"]