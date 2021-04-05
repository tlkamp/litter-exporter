# Litter Exporter  [![Go Report Badge]][Go Report] [![GoDocBadge]][GoDocLink]

Litter Exporter is a Prometheus exporter for the [Litter Robot](https://www.litter-robot.com/).

**This exporter uses an experimental API.** The upstream Litter Robot API is not publicly documented and may cause breaking
changes with no notice. Breaking changes will be handled as soon as possible.

## Metrics
Metrics are available at the default and metrics endpoints (`/` and `/metrics` respectively).

| Name                    | Type    |
|-------------------------|---------|
| litterrobot_capacity    | Gauge   |
| litterrobot_cycles      | Counter |
| litterrobot_drawer_full | Gauge   |
| litterrobot_nightlight  | Gauge   |
| litterrobot_sleepmode   | Gauge   |
| litterrobot_unit_status | Gauge   |


## Labels
| Name   | Value                                                                |
|--------|----------------------------------------------------------------------|
| name   | The name of the Litter Robot provided by the user during onboarding. |
| id     | The ID of the Litter Robot unit according to the API.                |
| serial | The serial number of the Litter Robot unit.                          |

## Usage
The Litter Exporter can be installed on a target machine or executed in a Docker container. The configuration options
are the same for both.

#### Docker
```console
$ docker run --rm -p 9080:9080 litter-exporter:latest \
    -email "your-email@example.com" \
    -password "your-password"
```

#### CLI
```console
$ ./litter-exporter -h
Usage of ./litter-exporter:
  -address string
        the server address (default "0.0.0.0:9080")
  -api-key string
        litter robot api key
  -api-url string
        litter robot API URL
  -auth-url string
        litter robot auth URL
  -client-id string
        litter robot client id
  -client-secret string
        litter robot client secret
  -email string
        litter robot account email address
  -log-level string
        the log level (default "info")
  -password string
        litter robot account password
```

## Configuration
The litter robot _requires_ only two pieces of information for configuration: the account email address and password.

The other fields are optional and provided as a convenience in case the upstream Litter Robot API changes.


## Unit Status
The unit status is represented by a non-negative integer.

| **String** | **Int** | **Description**                      |
|------------|---------|--------------------------------------|
| RDY        | 0       | Ready                                |
| CCP        | 1       | Clean Cycle in Progress              |
| CCC        | 2       | Clean Cycle Complete                 |
| CSF        | 3       | Cat Sensor Fault                     |
| DF1        | 4       | Drawer full - will still cycle       |
| DF2        | 5       | Drawer full - will still cycle       |
| CST        | 6       | Cat Sensor Timing                    |
| CSI        | 7       | Cat Sensor Interrupt                 |
| BR         | 8       | Bonnet Removed                       |
| P          | 9       | Paused                               |
| OFF        | 10      | Off                                  |
| SDF        | 11      | Drawer full - will not cycle         |
| DFS        | 12      | Drawer full - will not cycle         |

[Go Report Badge]: https://goreportcard.com/badge/github.com/tlkamp/litter-exporter
[Go Report]: https://goreportcard.com/report/github.com/tlkamp/litter-exporter
[GoDocBadge]: https://godoc.org/github.com/tlkamp/litter-exporter?status.svg
[GoDocLink]: https://godoc.org/github.com/tlkamp/litter-exporter
