# Logz.io client library

DEVELOP - [![Build Status](https://travis-ci.org/jonboydell/logzio_client.svg?branch=develop)](https://travis-ci.org/jonboydell/logzio_client) [![Coverage Status](https://coveralls.io/repos/github/jonboydell/logzio_client/badge.svg?branch=develop)](https://coveralls.io/github/jonboydell/logzio_client?branch=develop)

MASTER - [![Build Status](https://travis-ci.org/jonboydell/logzio_client.svg?branch=master)](https://travis-ci.org/jonboydell/logzio_client)

Client library for logz.io API, see below for supported endpoints.

The primary purpose of this library is to act as the API interface for the logz.io Terraform provider.

Logz.io have not written an especially consistent API. Sometimes, JSON will be presented back from an API call, sometimes not. Sometimes just a status code, sometimes a 200 status code, but with an error message in the body. I have attempted to shield the user of this library from those inconsistencies, but as they are laregely not documented, it's pretty diffcult to know if I've got them all.

[Roadmap](#roadmap)

##### Usage

Note: the lastest version of the API (1.1) is not backwards compatible with previous versions, specifically the client entrypoint names have changed to prevent naming conflicts. Use `UsersClient` ([Users API](#users)) , `AlertsClient` ([Alerts API](#alerts)) and `EndpointsClient` ([Endpoints API](#endpoints)) rather than `Users`, `Alerts` and `Endpoints`.


##### Alerts

To create an alert where the type field = 'mytype' and the loglevel field = ERROR, see the logz.io docs for more info

https://support.logz.io/hc/en-us/articles/209487329-How-do-I-create-an-Alert-

```go
client, _ := alerts.New(apiToken)
client.CreateAlert(alerts.CreateAlertType{
    Title:       "this is my alert",
    Description: "this is my description",
    QueryString: "loglevel:ERROR",
    Filter:      "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
    Operation:   alerts.OperatorGreaterThan,
    SeverityThresholdTiers: []alerts.SeverityThresholdType{
        alerts.SeverityThresholdType{
            alerts.SeverityHigh,
            10,
        },
    },
    SearchTimeFrameMinutes:       0,
    NotificationEmails:           []interface{}{},
    IsEnabled:                    true,
    SuppressNotificationsMinutes: 0,
    ValueAggregationType:         alerts.AggregationTypeCount,
    ValueAggregationField:        nil,
    GroupByAggregationFields:     []interface{}{"my_field"},
    AlertNotificationEndpoints:   []interface{}{},
})
```

|function|func name|
|---|---|
|create alert|`func (c *AlertsClient) CreateAlert(alert CreateAlertType) (*AlertType, error)`|
|update alert|`func (c *AlertsClient) UpdateAlert(alertId int64, alert CreateAlertType) (*AlertType, error)`
|delete alert|`func (c *AlertsClient) DeleteAlert(alertId int64) error`|
|get alert (by id)|`func (c *AlertsClient) GetAlert(alertId int64) (*AlertType, error)`|
|list alerts|`func (c *AlertsClient) ListAlerts() ([]AlertType, error)`|


##### Users

To create a new user, on a specific account or sub-account (you can get your account id from the logz.io console)

```go
client, _ := users.New(apiToken)
user := client.User{
    Username:  "createa@test.user",
    Fullname:  "my username",
    AccountId: 123456,
    Roles:     []int32{users.UserTypeUser},
}
```

|function|func name|
|---|---|
|create user|`func (c *UsersClient) CreateUser(user User) (*User, error)`|
|update user|`func (c *UsersClient) UpdateUser(user User) (*User, error)`|
|delete user|`func (c *UsersClient) DeleteUser(id int32) error`|
|get user|`func (c *UsersClient) GetUser(id int32) (*User, error)`|
|list users|`func (c *UsersClient) ListUsers() ([]User, error)`|
|suspend user|`func (c *UsersClient) SuspendUser(userId int32) (bool, error)`|
|unsuspend user|`func (c *UsersClient) UnSuspendUser(userId int32) (bool, error)`|

##### Endpoints

There's no 1-1 mapping between this library and the logz.io API functions, logz.io provide one API endpoint per *type* of notification endpoint being created.  I have abstracted this so that depending on how you create your `Endpoints` variable that you pass to `CreateEndpoint` the `CreateEndpoint` function will work out which API call to make. 

For more info, see: https://docs.logz.io/api/#tag/Manage-notification-endpoints

#### Contributing

1. Clone this repo locally
2. As this package uses Go modules, make sure you are outside of `$GOPATH` or you have the `GO111MODULE=on` environment variable set. Then run `go get` to pull down the dependencies.

##### Run tests
`go test -v -race ./...`



