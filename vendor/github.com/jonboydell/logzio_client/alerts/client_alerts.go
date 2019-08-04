package alerts

import (
	"fmt"

	"github.com/jonboydell/logzio_client/client"
)

const (
	alertsServiceEndpoint = "%s/v1/alerts"
)

const (
	AggregationTypeUniqueCount string = "UNIQUE_COUNT"
	AggregationTypeAvg         string = "AVG"
	AggregationTypeMax         string = "MAX"
	AggregationTypeNone        string = "NONE"
	AggregationTypeSum         string = "SUM"
	AggregationTypeCount       string = "COUNT"
	AggregationTypeMin         string = "MIN"

	OperatorGreaterThanOrEquals string = "GREATER_THAN_OR_EQUALS"
	OperatorLessThanOrEquals    string = "LESS_THAN_OR_EQUALS"
	OperatorGreaterThan         string = "GREATER_THAN"
	OperatorLessThan            string = "LESS_THAN"
	OperatorNotEquals           string = "NOT_EQUALS"
	OperatorEquals              string = "EQUALS"

	SeveritySevere string = "SEVERE"
	SeverityHigh   string = "HIGH"
	SeverityMedium string = "MEDIUM"
	SeverityLow    string = "LOW"
	SeverityInfo   string = "INFO"

	fldAlertId                      string = "alertId"
	fldAlertNotificationEndpoints   string = "alertNotificationEndpoints"
	fldCreatedAt                    string = "createdAt"
	fldCreatedBy                    string = "createdBy"
	fldDescription                  string = "description"
	fldFilter                       string = "filter"
	fldGroupByAggregationFields     string = "groupByAggregationFields"
	fldIsEnabled                    string = "isEnabled"
	fldQueryString                  string = "query_string"
	fldLastTriggeredAt              string = "lastTriggeredAt"
	fldLastUpdated                  string = "lastUpdated"
	fldNotificationEmails           string = "notificationEmails"
	fldOperation                    string = "operation"
	fldSearchTimeFrameMinutes       string = "searchTimeFrameMinutes"
	fldSeverity                     string = "severity"
	fldSeverityThresholdTiers       string = "severityThresholdTiers"
	fldSuppressNotificationsMinutes string = "suppressNotificationsMinutes"
	fldThreshold                    string = "threshold"
	fldTitle                        string = "title"
	fldValueAggregationField        string = "valueAggregationField"
	fldValueAggregationType         string = "valueAggregationType"
)

type CreateAlertType struct {
	AlertNotificationEndpoints   []interface{}
	Description                  string
	Filter                       string
	GroupByAggregationFields     []interface{}
	IsEnabled                    bool
	NotificationEmails           []interface{}
	Operation                    string
	QueryString                  string
	SearchTimeFrameMinutes       int
	SeverityThresholdTiers       []SeverityThresholdType `json:"severityThresholdTiers"`
	SuppressNotificationsMinutes int
	Title                        string
	ValueAggregationField        interface{}
	ValueAggregationType         string
}

type AlertType struct {
	AlertId                      int64
	AlertNotificationEndpoints   []interface{}
	CreatedAt                    string
	CreatedBy                    string
	Description                  string
	Filter                       string
	GroupByAggregationFields     []interface{}
	IsEnabled                    bool
	LastTriggeredAt              interface{}
	LastUpdated                  string
	NotificationEmails           []interface{}
	Operation                    string
	QueryString                  string `json:"query_string"`
	SearchTimeFrameMinutes       int
	Severity                     string
	SeverityThresholdTiers       []SeverityThresholdType `json:"severityThresholdTiers"`
	SuppressNotificationsMinutes int
	Threshold                    int `json:"threshold"`
	Title                        string
	ValueAggregationField        interface{}
	ValueAggregationType         string
}

type SeverityThresholdType struct {
	Severity  string `json:"severity"`
	Threshold int    `json:"threshold"`
}

func jsonAlertToAlert(jsonAlert map[string]interface{}) AlertType {
	alert := AlertType{
		AlertId:                    int64(jsonAlert[fldAlertId].(float64)),
		AlertNotificationEndpoints: jsonAlert[fldAlertNotificationEndpoints].([]interface{}),
		Description:                jsonAlert[fldDescription].(string),
		Filter:                     jsonAlert[fldFilter].(string),
		IsEnabled:                  jsonAlert[fldIsEnabled].(bool),
		LastUpdated:                jsonAlert[fldLastUpdated].(string),
		NotificationEmails:         jsonAlert[fldNotificationEmails].([]interface{}),
		Operation:                  jsonAlert[fldOperation].(string),
		QueryString:                jsonAlert[fldQueryString].(string),
		Severity:                   jsonAlert[fldSeverity].(string),
		SearchTimeFrameMinutes:     int(jsonAlert[fldSearchTimeFrameMinutes].(float64)),
		SeverityThresholdTiers:     []SeverityThresholdType{},
		Threshold:                  int(jsonAlert[fldThreshold].(float64)),
		Title:                      jsonAlert[fldTitle].(string),
		ValueAggregationType:       jsonAlert[fldValueAggregationType].(string),
	}

	if jsonAlert[fldGroupByAggregationFields] != nil {
		alert.GroupByAggregationFields = jsonAlert[fldGroupByAggregationFields].([]interface{})
	}

	if jsonAlert[fldCreatedAt] != nil {
		alert.CreatedAt = jsonAlert[fldCreatedAt].(string)
	}

	if jsonAlert[fldCreatedBy] != nil {
		alert.CreatedBy = jsonAlert[fldCreatedBy].(string)
	}

	if jsonAlert[fldLastTriggeredAt] != nil {
		alert.LastTriggeredAt = jsonAlert[fldLastTriggeredAt].(interface{})
	}

	tiers := jsonAlert[fldSeverityThresholdTiers].([]interface{})
	for x := 0; x < len(tiers); x++ {
		tier := tiers[x].(map[string]interface{})
		threshold := SeverityThresholdType{
			Threshold: int(tier[fldThreshold].(float64)),
			Severity:  tier[fldSeverity].(string),
		}
		alert.SeverityThresholdTiers = append(alert.SeverityThresholdTiers, threshold)
	}

	if jsonAlert[fldSuppressNotificationsMinutes] != nil {
		alert.SuppressNotificationsMinutes = int(jsonAlert[fldSuppressNotificationsMinutes].(float64))
	}

	if jsonAlert[fldValueAggregationField] != nil {
		alert.ValueAggregationField = jsonAlert[fldValueAggregationField].(interface{})
	}

	return alert
}

type AlertsClient struct {
	*client.Client
}

func New(apiToken, baseUrl string) (*AlertsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &AlertsClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}
