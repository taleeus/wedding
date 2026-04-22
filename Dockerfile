FROM golang:1.26 AS builder

WORKDIR /build
COPY . .

RUN go mod download

ENV CGO_ENABLED=1 GOOS=linux
RUN go build -buildvcs=false -tags netgo -ldflags '-s -w' -o wedding .

FROM scratch

COPY --from=builder ["/build/wedding", "/"]
EXPOSE 8080

ENTRYPOINT ["/wedding"]
