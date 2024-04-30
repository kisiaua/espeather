FROM golang:alpine3.19 as builder

RUN mkdir /espeather-build
WORKDIR /espeather-build
COPY . .

RUN go mod download && go mod verify
RUN go build -o espeather cmd/mqttHandler/main.go

FROM alpine:latest

RUN mkdir /espeather
WORKDIR /espeather
COPY --from=builder /espeather-build/espeather ./
EXPOSE 8080

ENTRYPOINT [ "./espeather"  ]
