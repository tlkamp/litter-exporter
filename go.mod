module tlkamp/litter-exporter

go 1.15

require (
	github.com/magefile/mage v1.11.0 // indirect
	github.com/mannkind/litterrobot v0.3.0 // indirect
	github.com/prometheus/client_golang v1.9.0
	github.com/sirupsen/logrus v1.8.1
	github.com/tlkamp/litter-api v0.0.0
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2 // indirect
)

replace github.com/tlkamp/litter-api => ../litter-api
