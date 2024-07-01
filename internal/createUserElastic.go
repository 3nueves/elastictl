package elastictl

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
)

type Connection struct {
	Config elasticsearch.Config
	Client *elasticsearch.Client
}

// values to access elastic cluster
var (
	HOST string = "10.40.0.240"
	PORT string = os.Getenv("ES_PORT")
)

// Create new client elastic
func CreateNewUser() *Connection {

	cfg := elasticsearch.Config{
		Addresses: []string{"http://" + HOST + ":" + PORT},
	}

	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	config := Connection{
		Config: cfg,
		Client: es,
	}

	return &config
}
