cd ..

# generate descriptor
docker run --rm -v $(pwd):$(pwd) -w $(pwd) jevon82/golang-builder-alpine \
/bin/sh -c "protoc --descriptor_set_out=./descriptor/messages.bin -I. -I/app/  ./api/messages.proto"

# generate code
docker run --rm -v $(pwd):$(pwd) -w $(pwd) jevon82/golang-builder-alpine \
/bin/sh -c "protoc  --go_out=:./pkg --go-grpc_out=:./pkg -I. -I/app/  ./api/todolist.proto"

docker run --rm -v $(pwd):$(pwd) -w $(pwd) jevon82/golang-builder-alpine \
/bin/sh -c "protoc  --go_out=:./pkg --go-grpc_out=:./pkg -I. -I/app/  ./api/messages.proto"

# generate GRPC Gateway code
docker run --rm -v $(pwd):$(pwd) -w $(pwd) jevon82/golang-builder-alpine \
/bin/sh -c "protoc -I . --grpc-gateway_out ./pkg \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt generate_unbound_methods=true \
    -I. -I/app/ ./api/todolist.proto"

# update file ownership
docker run --rm -v $(pwd):$(pwd) -w $(pwd)  jevon82/golang-builder-alpine \
/bin/sh -c "chown -R $(id -u):$(id -g) pkg/pb"
