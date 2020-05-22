FROM golang:latest AS builder
COPY main.go .
RUN go build -tags netgo -o server

FROM scratch
COPY --from=builder /go/server .
CMD ["./server"]
