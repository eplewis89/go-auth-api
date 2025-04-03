FROM postgres:15.1 AS go-auth-db-base
COPY ./migrations/* /docker-entrypoint-initdb.d/
