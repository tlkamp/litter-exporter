package main

import (
	"context"
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

func init() {
	flag.BoolP("help", "h", false, "print help and exit")
	flag.String("email", "", "litter robot account email address")
	flag.String("password", "", "litter robot account password")
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
	address := viper.GetString("address")

	ctx := context.Background()

	lrc := NewCollector(ctx, email, password)
	ticker := time.NewTicker((3600 * time.Second) - (120 * time.Second))
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				if err := lrc.lrClient.RefreshToken(ctx); err != nil {
					log.Error("error refreshing token", "error", err.Error())
				}
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
