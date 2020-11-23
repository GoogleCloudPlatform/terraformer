# Dashboards
output "datadog_dashboard--ordered_dashboard_example" {
  value = datadog_dashboard.ordered_dashboard_example.id
}

output "datadog_dashboard--free_dashboard_example" {
  value = datadog_dashboard.free_dashboard_example.id
}

# Downtimes
output "datadog_downtime--downtime_example" {
  value = datadog_downtime.downtime_example.id
}

# Monitors
output "datadog_monitor--monitor_example" {
  value = datadog_monitor.monitor_example.id
}

# Synthetics
output "datadog_synthetics_test--test_api_example" {
  value = datadog_synthetics_test.test_api_example.id
}

# Users
output "datadog_user--user_example" {
  value = datadog_user.user_example.id
}

output "datadog_user--user_example_two" {
  value = datadog_user.user_example_two.id
}

