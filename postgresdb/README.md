# Database Layer

This API stores its data in a Postgres DB. The DB is named `go-auth-db`,
and the `goauth` role/user is the owner.

The schema is created from a sequence of update scripts in the `migrations` dir.
Given the initial empty `go-auth-db` DB, with the `goauth` role
already existing, each of the migration files are
executed in sequential order to build up the schema.

## Migrations

The `000-bootstrap-extensions.sql` script runs on the docker
DB before any of the other migration scripts are run.

One of the initial tables created is the `config` table, which
consists of key-value pairs. At a minimum, the table contains a `schema_version`
entry, where the value is an integer. Each migration file
after `001` must call the `start_schema_update(N)` stored procedure inside a transaction,
where `N` is the integer value of the `NNN` segment of the script name.

To create your own update, copy and rename `NNN-migration-template.sql`
file. Note that you MUST follow the conventions in that file, including
using uppercase keywords like `BEGIN TRANASACTION;`. This is CRITICAL,
as there's search/replace stuff happening that relies upon the exact
syntax.

## Queries

The queries folder should contain files that are compiled from sql into the ORM
compiled by SQLC. For each table that exists, there should be a separate query file,
e.g. if you have a table `Users` there should be a `users.sql` file in the queries
folder.

These queries use the SQLC (https://sqlc.dev/) specification, as the queries are compiled
into individual ORM calls. For example, if we wanted to get `Users` by email address,
we would format the query as:

```
--name: GetUserByEmailDomain :one
SELECT *
FROM Users
WHERE email=$1;
```

There are many examples in the SQLC documentation: https://docs.sqlc.dev/en/stable/howto/select.html

## Compiling into ORM

SQLC utilizes migrations to compile the tables, and the queries to generate functions
for those tables. It also automatically tests your database structure against itself and
the functions written.

In order to compile the ORM using SQLC, the command is:

`sqlc generate`

## Dockerfile

For developer convenience, a Postgres Dockerfile is provided that
builds the schema on startup. Use the `docker-build-run.sh` script
to build and run the image. The Postgres container will be available
on port `5432`, with user `goauth`. You must supply the desired
password in the `POSTGRES_PASSWORD` envar. This can be setup with
`direnv` for simpler runtime, otherwise set an environment variable.

## Adding Verbose Logging

If you want to add verbose logging you can run `docker-build-run.sh --debug`

This argument will enable verbose logging so you can see arguments
being passed into calls, success, failure, etc...

```shell
$ POSTGRES_PASSWORD=abc123 ./db/docker-build-run.sh
```


