postgres:
	docker run --name go_bookstore_user --network bookstore-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it go_bookstore_user createdb --username=root --owner=root users

dropdb:
	docker exec -it go_bookstore_user dropdb users

dev:
	air

.PHONY: db dev dropdb createdb postgres