pgdinit:
	docker run -p 5432:5432 --name xlpg -d -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER-postgres -e POSTGRES_DB=xlearn -v pgdata:${CURDIR}/db postgres

pg:
	docker start xlpg

nopg:
	docker stop xlpg

run:
	go run .

pgrun: pg run nopg
