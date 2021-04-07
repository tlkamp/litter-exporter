FROM golang:alpine
COPY litter-exporter /go/bin/litter-exporter
RUN apk add --no-cache tzdata
ENTRYPOINT ["/go/bin/litter-exporter"]
CMD ["-h"]
