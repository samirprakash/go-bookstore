cassandra:
	docker run -p 9042:9042 --rm --name cassandra -d cassandra

dev:
	air

test:
	go test ./... -cover

.PHONY: cassandra dev test