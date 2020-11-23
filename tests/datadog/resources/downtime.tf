# Create a Datadog downtime for all monitors
resource "datadog_downtime" "downtime_example" {
  scope = ["*"]

  recurrence {
    type   = "days"
    period = 1
  }
}
