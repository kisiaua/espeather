FROM golang:alpine3.19 as builder

RUN apk add build-base

WORKDIR /espeather-build
COPY . .

RUN apk add build-base

RUN go mod download && go mod verify
RUN CGO_ENABLED=1 go build -o espeather cmd/mqttHandler/main.go


FROM golang:alpine3.19 as migration

RUN apk add git

RUN mkdir /espeather-migration
WORKDIR /espeather-migration
COPY ./db/migrations .

RUN git clone https://github.com/pressly/goose && cd goose && go mod tidy && go build -o goose ./cmd/goose
RUN GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=data.db GOOSE_MIGRATION_DIR=./ ./goose/goose up


FROM alpine:3.19

RUN mkdir /espeather
WORKDIR /espeather
COPY --from=migration /espeather-migration/data.db ./db/
COPY --from=builder /espeather-build/espeather ./

ENV MQTT_CONFIG="{\"address\": \"192.168.0.201\", \"port\": 1883, \"clientID\": \"laptop\", \"topics\": [\"indoor_readings\"]}"

EXPOSE 8080

ENTRYPOINT [ "./espeather"  ]
