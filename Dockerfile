FROM golang:alpine as builder
COPY main.go collector.go go.mod go.sum /app/
WORKDIR /app
RUN CGO_ENABLED=0 go build -o litter-exporter

FROM golang:alpine
COPY --from=builder /app/litter-exporter /go/bin/litter-exporter
RUN apk add --no-cache tzdata
ENTRYPOINT ["/go/bin/litter-exporter"]
CMD ["-h"]
