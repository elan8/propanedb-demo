FROM jevon82/golang-builder-alpine:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./cmd/grpc-gateway/main.go cmd/grpc-gateway/main.go
COPY ./pkg/pb pkg/pb
COPY ./pkg/util pkg/util
WORKDIR /app/cmd/grpc-gateway
RUN CGO_ENABLED=0 GOOS=linux go build  -o /app/main main.go

FROM jevon82/microservice-base
WORKDIR /app/
EXPOSE 8080
COPY --from=builder /app/main .
CMD ["/app/main"]