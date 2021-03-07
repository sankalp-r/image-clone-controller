FROM golang:1.15.0-alpine3.12 as builder
RUN mkdir /app
WORKDIR /app
RUN apk add --update --no-cache ca-certificates git
## Add this go mod download command to pull in any dependencies
COPY pkg ./pkg
COPY go.mod go.sum ./
RUN go mod download

COPY cmd/imgcc/main ./imgcc

WORKDIR /app/imgcc

RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags "-s" -a -installsuffix cgo .

FROM alpine:3.11.6
WORKDIR /app
COPY --from=builder /app/imgcc/main ./

ENTRYPOINT ["./main"]