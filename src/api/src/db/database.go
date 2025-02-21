package db

import (
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

// Global database session
var Session *gocql.Session

// Init connection with ScyllaDB
func InitDB() {
	_ = godotenv.Load() // Load .env file

	cluster := gocql.NewCluster(os.Getenv("SCYLLA_HOSTS"))
	cluster.Keyspace = os.Getenv("SCYLLA_KEYSPACE")
	cluster.Consistency = gocql.Quorum

	sess, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Can't connect with ScyllaDB: %v", err)
	}
	Session = sess
	fmt.Println("ScyllaDB connection successful!")
}

// CloseDB closes the database session
func CloseDB() {
	if Session != nil {
		Session.Close()
	}
}
