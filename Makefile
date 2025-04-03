# rundb - Run Docker Database
run-db:
	./database/docker-build-run.sh

# gen - sqlc generation
gen-sqlc:
	@sqlc generate -f ./database/sqlc.yaml