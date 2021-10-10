set -e

#cd ..

mkdir -p /tmp/go-build
export DOCKER_BUILDKIT=0

docker build -t propanedb-demo/grpc-gateway:latest  -f ./cmd/grpc-gateway/Dockerfile .
docker build -t propanedb-demo/todo-service:latest  -f ./cmd/todo-service/Dockerfile .

