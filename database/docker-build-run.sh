#!/usr/bin/env bash
set -x

DIR="$(cd "$(dirname "$0")" && pwd)"

docker rm -f go-auth-db || true

docker build -t go-auth-api/go-auth-db -f "${DIR}"/Dockerfile "${DIR}"
#docker run -d -p 5432:5432 go-auth-api/go-auth-db  # run as daemon

post_args=""
docker_args="-p 5432:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} --name go-auth-db"

if [[ "$1" == "--daemon" ]]; then
  docker_args+=" -d"
fi

if [[ "$1" == "--debug" ]]; then
  post_args+=" -c log_statement=all"
fi

docker run $docker_args go-auth-api/go-auth-db $post_args
