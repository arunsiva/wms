package env

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"go.uber.org/zap"
)

// NewProd produces a full production environment configuration
func NewProd() *Config {
	config := NewBaseConfig()

	config.Logger = func() *zap.Logger {
		config.Once(func() {
			loggerConfig := zap.NewProductionConfig()
			loggerConfig.EncoderConfig.TimeKey = "time"
			loggerConfig.EncoderConfig.LevelKey = "severity"
			loggerConfig.EncoderConfig.MessageKey = "message"

			logger, err := loggerConfig.Build()
			if err != nil {
				panic(err)
			}

			config.logger = logger
		})

		return config.logger
	}

	config.RootHandler = func() http.Handler {
		config.Once(func() {

			url, err := url.Parse(config.WebURL)
			if err != nil {
				panic(err)
			}

			config.Logger().Info("reverse proxy web app configured", zap.String("url", url.String()))

			config.rootHandler = httputil.NewSingleHostReverseProxy(url)
		})

		return config.rootHandler
	}

	return config
}
