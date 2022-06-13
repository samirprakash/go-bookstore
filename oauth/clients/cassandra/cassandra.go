package cassandra

import (
	"github.com/gocql/gocql"
)

var cluster *gocql.ClusterConfig

func init() {
	// connect to the cassandra cluster
	cluster = gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
