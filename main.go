package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// These values have been reverse-engineered from the mobile apps
const (
	defaultApiKey       = "p7ndMoj61npRZP5CVz9v4Uj0bG769xy6758QRBPb"
	defaultClientId     = "IYXzWN908psOm7sNpe4G.ios.whisker.robots"
	defaultClientSecret = "C63CLXOmwNaqLTB2xXo6QIWGwwBamcPuaul"
)

func init() {
	flag.BoolP("help", "h", false, "print help and exit")
	flag.String("api-key", defaultApiKey, "litter robot api key")
	flag.String("email", "", "litter robot account email address")
	flag.String("password", "", "litter robot account password")
	flag.String("client-id", defaultClientId, "litter robot client id")
	flag.String("client-secret", defaultClientSecret, "litter robot client secret")
	flag.String("api-url", "", "litter robot API URL")
	flag.String("auth-url", "", "litter robot auth URL")
	flag.String("address", "0.0.0.0:9080", "the server address")
	flag.String("log-level", "info", "the log level")

	// Replace dashes with underscores in environment variable keys
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// env vars are prefixed with LR_
	viper.SetEnvPrefix("LR")
	viper.AutomaticEnv()

	flag.Parse()
	if err := viper.BindPFlags(flag.CommandLine); err != nil {
		log.Fatal(err)
	}

	// If the uer requests help, print help and exit
	if viper.GetBool("help") {
		flag.Usage()
		os.Exit(0)
	}

	logLevel := viper.GetString("log-level")
	log.SetLevel(parseLevelDefault(logLevel, log.InfoLevel))
}

// attempts to parse a provided string as a LogLevel, and returns the provided default 'def' if parsing fails.
func parseLevelDefault(level string, def log.Level) log.Level {
	if l, err := log.ParseLevel(level); err != nil {
		return def
	} else {
		return l
	}

}

func main() {
	email := viper.GetString("email")
	password := viper.GetString("password")
	apiKey := viper.GetString("api-key")
	clientSecret := viper.GetString("client-secret")
	clientId := viper.GetString("client-id")
	endpoint := viper.GetString("api-url")
	authEndpoint := viper.GetString("auth-url")
	address := viper.GetString("address")

	lrc := NewCollector(email, password, apiKey, clientSecret, clientId, endpoint, authEndpoint)
	ticker := time.NewTicker(lrc.lrClient.Expiry - (120 * time.Second))
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				lrc.lrClient.RefreshToken()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	prometheus.MustRegister(lrc)
	http.Handle("/metrics", promhttp.Handler())

	log.Infof("Starting exporter on %s", address)
	if err := http.ListenAndServe(address, promhttp.Handler()); err != http.ErrServerClosed {
		log.WithField("error", err).Fatal("Starting server failed")
	}

	close(quit)
}
