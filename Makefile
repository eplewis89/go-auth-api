# rundb - Run Docker Database
run-psql:
	./postgresdb/docker-build-run.sh

# gen - sqlc generation
gen-sqlc-psql:
	@sqlc generate -f ./postgresdb/sqlc.yaml

# gen - sqlc generation
gen-sqlc-sqlite:
	@sqlc generate -f ./sqlitedb/sqlc.yaml