FROM golang:1.8.3
ADD . /go/src/github.com/vitraum/svg2png
RUN go install github.com/vitraum/svg2png

WORKDIR /app/
ENTRYPOINT /go/bin/svg2png

