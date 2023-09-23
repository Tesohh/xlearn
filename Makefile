include .env

pgdinit:
	docker run -p 5432:5432 --name xlpg -d -e POSTGRES_PASSWORD=${PG_PASS} -e POSTGRES_USER=${PG_USER} -e POSTGRES_DB=xlearn -v pgdata:${CURDIR}/.pgdb postgres

pg:
	@docker start xlpg

nopg:
	@docker stop xlpg

run:
	@go run .

pgrun: pg run nopg
