POSTGRES_USER := postgres
POSTGRES_PASSWORD := test1111
POSTGRES_DB := default
MIGRATE_CMD := golang-migrate
CONNECTION_STRING := postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@127.0.0.1:5432/$(POSTGRES_DB)?sslmode=disable
start_db: 
	set -x
	docker run --name postgres_database\
			--rm\
			-p 5432:5432\
			-e POSTGRES_USER=${POSTGRES_USER}\
			-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD}\
			-e POSTGRES_DB=${POSTGRES_DB}\
			-e PGDATA="/var/lib/postgresql/data/pgdata"\
			-v ${PWD}/database:/var/lib/postgresql/data\
			postgres
build:
	go build \
	-o bin/import \
	-v cmd/import/main.go
	go build \
	-o bin/count_from_zim \
	-v cmd/count_from_zim/main.go
dbshell:
	psql -h 127.0.0.1 --user=${POSTGRES_USER} ${POSTGRES_DB}
migrate:
	set -x
	${MIGRATE_CMD} -path db/migration -database ${CONNECTION_STRING} -verbose up
*:
	echo "start_db: to start postgres in docker"
	echo "migrate: to apply migration to database"
