FROM golang:alpine AS builder
ADD . /go/src/github.com/zqureshi/gdclient
WORKDIR /go/src/github.com/zqureshi/gdclient
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/gdclient .

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/* /app/

ENTRYPOINT ["/app/gdclient"]
