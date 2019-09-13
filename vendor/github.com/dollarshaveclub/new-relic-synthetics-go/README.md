# new-relic-synthetics-go

[![GoDoc](https://godoc.org/github.com/dollarshaveclub/new-relic-synthetics-go?status.svg)](https://godoc.org/github.com/dollarshaveclub/new-relic-synthetics-go)
[![CircleCI](https://circleci.com/gh/dollarshaveclub/new-relic-synthetics-go.svg?style=svg)](https://circleci.com/gh/dollarshaveclub/new-relic-synthetics-go)

A [New Relic Synthetics](https://newrelic.com/synthetics) API client
for Go. This package provides CRUD functionality for both Synthetics
monitors and alert conditions. Detailed API docs can be found on
the GoDoc link above.

## Example

```go
conf := func(s *synthetics.Client) {
	s.APIKey = os.Getenv("NEWRELIC_API_KEY")
}
client, _ := synthetics.NewClient(conf)

// Get specific monitor
client.GetMonitor("monitor-id")

// Create a monitor
client.CreateMonitor(&synthetics.CreateMonitorArgs{
	Name:         "sample-monitor",
	Type:         "SIMPLE",
	Frequency:    60,
	URI:          "https://www.dollarshaveclub.com",
	Locations:    []string{"AWS_US_WEST_1"},
	Status:       "ENABLED",
	SLAThreshold: 7,
})

// Update monitor
client.UpdateMonitor("monitor-id", &synthetics.UpdateMonitorArgs{
	Name: "new-monitor-name",
})

// Create an alert condition
client.CreateAlertCondition("policy-id"), &synthetics.CreateAlertConditionArgs{
	Name: "alert-condition-name",
	MonitorID: "monitor-id",
	Enabled: true,
})
```
