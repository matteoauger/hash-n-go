# Base : go image
FROM golang

WORKDIR /go/src/hash-n-go-worker
COPY . .

RUN go install hash-n-go-worker 

ENTRYPOINT ["hash-n-go-worker"]
