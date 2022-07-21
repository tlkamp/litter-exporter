package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	apiKey       = flag.String("api-key", defaultApiKey, "litter robot api key")
	email        = flag.String("email", "", "litter robot account email address")
	password     = flag.String("password", "", "litter robot account password")
	clientId     = flag.String("client-id", defaultClientId, "litter robot client id")
	clientSecret = flag.String("client-secret", defaultClientSecret, "litter robot client secret")
	endpoint     = flag.String("api-url", "", "litter robot API URL")
	authEndpoint = flag.String("auth-url", "", "litter robot auth URL")
	address      = flag.String("address", "0.0.0.0:9080", "the server address")
	logLevel     = flag.String("log-level", "info", "the log level")
)

// These values have been reverse-engineered from the mobile apps
const (
	defaultApiKey       = "p7ndMoj61npRZP5CVz9v4Uj0bG769xy6758QRBPb"
	defaultClientId     = "IYXzWN908psOm7sNpe4G.ios.whisker.robots"
	defaultClientSecret = "C63CLXOmwNaqLTB2xXo6QIWGwwBamcPuaul"
)

func init() {
	flag.Parse()
	level, err := log.ParseLevel(*logLevel)
	if err != nil {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}
}

func main() {
	lrc := NewCollector(*email, *password, *apiKey, *clientSecret, *clientId, *endpoint, *authEndpoint)
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
	log.Infof("Starting exporter on %s", *address)
	if err := http.ListenAndServe(*address, promhttp.Handler()); err != http.ErrServerClosed {
		log.WithField("error", err).Fatal("Starting server failed")
	}
	close(quit)
}
