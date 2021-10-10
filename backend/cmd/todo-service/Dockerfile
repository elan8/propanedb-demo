FROM jevon82/golang-builder-alpine:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./cmd/todo-service/main.go cmd/todo-service/main.go
COPY ./pkg/todo pkg/todo
COPY ./pkg/pb pkg/pb
COPY ./pkg/util pkg/util
COPY ./descriptor/messages.bin /app
# RUN ls /app/

WORKDIR /app/cmd/todo-service
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main main.go
CMD ["/app/main"]
EXPOSE 50000