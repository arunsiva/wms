package main

import (
	"flag"
	"log"
	"net/http"

	"wms/env"
)

var (
	hot             bool
	environment     string
	single          bool
	webURL          string
	stateMachineURL string
	mongoURL        string
)

// defaults
const (
	DefaultWebURL = "http://pulsar-web.pulsar-eng:9222"
	//DefaultMongoURL        = "mongodb://mongo.pulsar-eng:27017"
	DefaultMongoURL        = "mongodb://10.33.158.133:8815"
	DefaultMongoDB         = "pulsar"
	DefaultStateMachineURL = "http://0.0.0.0:9425"
	DefaultWebpackPort     = "9426"
	DataDir                = "data"
)

func init() {
	log.SetFlags(0)
	loadFlag()
	// TODO Just get rid of this flag if no one depends on it
	if hot {
		environment = "dev"
	}
}

func loadFlag() {
	flag.StringVar(&webURL,
		"web",
		DefaultWebURL,
		"web app server url")
	flag.StringVar(&mongoURL,
		"mongo",
		DefaultMongoURL,
		"mongodb url")
	flag.StringVar(&stateMachineURL,
		"state-machine",
		DefaultStateMachineURL,
		"state machine url")
	flag.BoolVar(&hot, "hot", false, "hot development mode")
	flag.BoolVar(&single, "single", false, "single docker mode")
	flag.StringVar(
		&environment,
		"env",
		"prod",
		"environment config to use",
	)
	flag.Parse()
}

func main() {
	config := env.Load(environment)

	config.MongoDB = DefaultMongoDB
	config.MongoURL = mongoURL
	config.StateMachineURL = stateMachineURL
	config.WebURL = webURL

	defer config.EventIndex().Close()

	log.Fatal(http.ListenAndServe("0.0.0.0:9191", config.Router()))
}
