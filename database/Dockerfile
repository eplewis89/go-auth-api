FROM postgres:15.1 AS go-auth-db-base
ENV POSTGRES_DB="go_auth_db"
ENV POSTGRES_USER="goauth"
COPY ./migrations/* /docker-entrypoint-initdb.d/
