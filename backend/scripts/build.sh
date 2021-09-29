set -e

cd ..

mkdir -p /tmp/go-build
export DOCKER_BUILDKIT=0

docker build -t propanedb-demo/grpc-gateway:latest  -f ./cmd/grpc-gateway/Dockerfile .
docker build -t propanedb-demo/todo-service:latest  -f ./cmd/todo-service/Dockerfile .
# docker build -t harbor.jevontech.com/dactory/api-gateway:$tag  -f ./cmd/api-gateway/Dockerfile .
# docker build -t harbor.jevontech.com/dactory/dbbackup:$tag  -f ./cmd/database-backup/Dockerfile .
# docker build -t harbor.jevontech.com/dactory/userservice:$tag  -f ./cmd/user-service/Dockerfile .
# docker build -t harbor.jevontech.com/dactory/projectservice:$tag  -f ./cmd/project-service/Dockerfile .
# docker build -t harbor.jevontech.com/dactory/generatorservice:$tag  -f ./cmd/generator-service/Dockerfile .
# docker build -t harbor.jevontech.com/dactory/db:$tag -f ./cmd/db/Dockerfile .
