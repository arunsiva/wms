package env

import (
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"

	"go.uber.org/zap"
)

// NewDev produces a full development environment configuration
func NewDev() *Config {
	config := NewBaseConfig()

	config.RootHandler = func() http.Handler {
		config.Once(func() {
			cmd := exec.Command("npm", "start")
			cmd.Dir = "html"
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Start(); err != nil {
				config.Logger().Error("failed to start webpack!")
				panic(err)
			}

			url, err := config.WebpackURL()
			if err != nil {
				panic(err)
			}

			config.Logger().Info("web app started", zap.String("url", url.String()))

			config.rootHandler = httputil.NewSingleHostReverseProxy(url)
		})

		return config.rootHandler
	}

	return config
}
