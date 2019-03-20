FROM golang:1.12 as builder
ADD . /go/src/github.com/vitraum/svg2png
RUN CGO_ENABLED=0 go build -tags netgo -ldflags "-s -w -extldflags '-static'" -o /go/bin/svg2png github.com/vitraum/svg2png

FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/app/svg2png"]
COPY --from=builder /go/bin/svg2png /app/

