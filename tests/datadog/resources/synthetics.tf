resource "datadog_synthetics_test" "test_api_example" {
  type = "api"
  subtype = "http"
  request = {
    method = "GET"
    url = "https://www.example.org"
  }
  request_headers = {
    Content-Type = "application/json"
    Authentication = "Token: 1234566789"
  }
  assertion {
    type = "statusCode"
    operator = "is"
    target = "200"
  }
  locations = [ "aws:eu-central-1" ]
  options_list {
    tick_every = 900

    retry {
      count = 2
      interval = 300
    }

    monitor_options {
      renotify_interval = 100
    }
  }
  name = "An API test on example.org"
  message = "Notify @pagerduty"
  tags = ["foo:bar", "foo", "env:test"]

  status = "live"
}
