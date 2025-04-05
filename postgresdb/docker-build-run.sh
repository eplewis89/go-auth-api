#!/usr/bin/env bash
set -x

DIR="$(cd "$(dirname "$0")" && pwd)"

docker rm -f ${POSTGRES_DB} || true

docker build -t go-auth-api/${POSTGRES_DB} -f "${DIR}"/Dockerfile "${DIR}"
#docker run -d -p 5432:5432 go-auth-api/${POSTGRES_DB}  # run as daemon

post_args=""
docker_args="-p 5432:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
 -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_DB=${POSTGRES_DB} --name ${POSTGRES_DB}"

if [[ "$1" == "--daemon" ]]; then
  docker_args+=" -d"
fi

if [[ "$1" == "--debug" ]]; then
  post_args+=" -c log_statement=all"
fi

docker run $docker_args go-auth-api/${POSTGRES_DB} $post_args
