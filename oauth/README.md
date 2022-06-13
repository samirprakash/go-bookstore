# oAuth service with Caasandra

- `make cassandra` to run cassandra instance on docker
- `pip install cqlsh` to query cassandra instance

### configure database

- `cqlsh` on terminal to connect to the cassandra instance on docker
- `describe keyspaces` on terminal
- `CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1};` to create a new keyspace
- `USE oauth` to use the newly created keyspace
- `CREATE TABLE tokens( access_token text PRIMARY KEY, user_id varint, client_id varint, expires_at varint );` to create a new table in the `oauth` namespace
- `describe tables` to view the newly create table
