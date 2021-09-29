cd ..

# generate Protobuf and GRPC files
docker run --rm -v $(pwd):$(pwd) -w $(pwd) jevon82/golang-builder-alpine \
/bin/sh -c "protoc  --go_out=:./pkg --go-grpc_out=:./pkg  --descriptor_set_out=./descriptor/todolist.bin -I. -I/app/  ./api/todolist.proto"

# generate GRPC Gateway code
docker run --rm -v $(pwd):$(pwd) -w $(pwd) jevon82/golang-builder-alpine \
/bin/sh -c "protoc -I . --grpc-gateway_out ./pkg \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt generate_unbound_methods=true \
    -I. -I/app/ ./api/todolist.proto"

# update file ownership
docker run --rm -v $(pwd):$(pwd) -w $(pwd)  jevon82/golang-builder-alpine \
/bin/sh -c "chown -R $(id -u):$(id -g) pkg/pb"
