postgres:
	docker run --name go-bookstore-users --network bookstore-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it go-bookstore-users createdb --username=root --owner=root bookstore_users

populatedb:
	docker cp datasources/psql/users/db.sql go-bookstore-users:/
	docker exec -it go-bookstore-users psql -d bookstore_users -f /db.sql

dropdb:
	docker exec -it go-bookstore-users dropdb bookstore_users

dev:
	air

.PHONY: db dev dropdb createdb postgres