package env

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	isatty "github.com/mattn/go-isatty"

	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

)

// Config contains all environment specific configuration needed to run the
// server
type Config struct {
	Logger func() *zap.Logger
	logger *zap.Logger


	Manager func() socker.Manager
	manager socker.Manager


	RootHandler func() http.Handler
	rootHandler http.Handler

	Router func() *mux.Router
	router *mux.Router

	WebpackHost string
	WebpackPort string

	WebURL string

//	callsite.Memoizer
}

// Load tries to return the configuration corresponding to the environment name
// passed in. If no such environment exists it panics.
func Load(environment string) *Config {
	switch environment {
	case "dev":
		return NewDev()
	case "prod":
		return NewProd()
	default:
		panic("Unknown config environment: " + environment)
	}
}

// WebpackURL attempts to produce a url by joining the configured host and
// port. If the URL can no be parsed an error is returned.
func (c *Config) WebpackURL() (*url.URL, error) {
	return url.Parse(strings.Join([]string{"http://", c.WebpackHost, ":", c.WebpackPort}, ""))
}

// NewBaseConfig produces a base configuration which can be extended or overridden
// selectively to define a new environment
func NewBaseConfig() *Config {
	config := new(Config)

	config.Logger = func() *zap.Logger {
		config.Once(func() {
			loggerConfig := zap.NewDevelopmentConfig()
			if isatty.IsTerminal(os.Stdout.Fd()) {
				loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			}
			logger, err := loggerConfig.Build()
			if err != nil {
				panic(err)
			}
			config.logger = logger
		})

		return config.logger
	}

	config.Router = func() *mux.Router {
		config.Once(func() {
			router := mux.NewRouter().StrictSlash(true)
			router.HandleFunc("/ws", config.WebsocketHandler())
			router.HandleFunc("/ingest", config.Ingester())
			router.HandleFunc("/message/{messageId}", config.MessageLoader().Load)
			router.HandleFunc("/message/{messageId}/{format:(?:json|html)}", config.MessageLoader().Load)
			router.PathPrefix("/").Handler(config.RootHandler())

			config.router = router
		})

		return config.router
	}

	config.EventIndex = func() event.Index {
		config.Once(func() {
			config.eventIndex = &mongo.Index{
				Name:    config.MongoDB,
				Session: config.Session(),
			}
		})

		return config.eventIndex
	}

	config.Manager = func() socker.Manager {
		config.Once(func() {
			config.manager = socker.NewManager()
		})

		return config.manager
	}

	config.Session = func() *mgo.Session {
		config.Once(func() {
			retries := 3
			urlField := zap.String("url", config.MongoURL)
			config.Logger().Info("connecting to MONGO", urlField)

			for {
				session, err := mgo.Dial(config.MongoURL)
				if err != nil {
					if retries <= 0 {
						panic(err)
					}
					config.Logger().Warn("failed to connect. retrying.", urlField)
					retries--
					time.Sleep(3 * time.Second)
				}

				config.session = session
				break
			}
		})

		return config.session
	}

	config.WebsocketHandler = func() http.HandlerFunc {
		config.Once(func() {
			manager := config.Manager()
			eventIndex := config.EventIndex()

			options := []pulsar.ClientOption{
				pulsar.WithLogger(config.Logger()),
				pulsar.WithStateMachineURL(config.StateMachineURL),
				pulsar.WithRetryLimiter(ratelimit.New(1)),
				pulsar.WithMarshaler(json.Marshal),
				pulsar.WithTailerFactory(
					func(r *pulsar.Request, c *pulsar.Client) pulsar.Tailable {
						return pulsar.NewTailer(r, c)
					},
				),
			}

			config.websocketHandler = func(w http.ResponseWriter, r *http.Request) {
				client := pulsar.NewClient(eventIndex, options...)
				manager.Handle(w, r, client, false)
			}
		})

		return config.websocketHandler
	}

	config.MessageLoader = func() *message.Loader {
		config.Once(func() {
			config.messageLoader = message.NewLoader(config.EventIndex(), config.Logger())
		})

		return config.messageLoader
	}

	config.Ingester = func() http.HandlerFunc {
		config.Once(func() {
			logger := config.Logger()

			config.ingester = func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				logger.Info("ingesting", zap.String("body", string(body)))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("ok"))
			}
		})

		return config.ingester
	}

	return config
}
