FROM golang:1.16.0-alpine AS builder

WORKDIR $GOPATH/src/github.com/MOZGIII/docker-ps-exporter

COPY . .

RUN go build -o /usr/local/bin/docker_ps_exporter ./cmd/docker_ps_exporter

FROM alpine

COPY --from=builder /usr/local/bin/docker_ps_exporter /usr/local/bin/docker_ps_exporter

EXPOSE 9491

CMD [ "/usr/local/bin/docker_ps_exporter" ]
