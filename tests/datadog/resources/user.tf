# Create a new Datadog user
resource "datadog_user" "user_example" {
  email  = "new@example.com"
  handle = "new@example.com"
  name   = "New User"
}

resource "datadog_user" "user_example_two" {
  email  = "new_two@example.com"
  handle = "new_two@example.com"
  name   = "New User"
}
