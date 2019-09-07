package alerts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
)

const createAlertServiceUrl string = alertsServiceEndpoint
const createAlertServiceMethod string = http.MethodPost
const createAlertMethodSuccess int = 200

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

func validateCreateAlertRequest(alert CreateAlertType) error {

	if len(alert.Title) == 0 {
		return fmt.Errorf("title must be set")
	}

	if len(alert.QueryString) == 0 {
		return fmt.Errorf("query string must be set")
	}

	if alert.NotificationEmails == nil {
		return fmt.Errorf("notificationEmails must not be nil")
	}

	validAggregationTypes := []string{AggregationTypeUniqueCount, AggregationTypeAvg, AggregationTypeMax, AggregationTypeNone, AggregationTypeSum, AggregationTypeCount, AggregationTypeMin}
	if !logzio_client.Contains(validAggregationTypes, alert.ValueAggregationType) {
		return fmt.Errorf("valueAggregationType must be one of %s", validAggregationTypes)
	}

	validOperations := []string{OperatorGreaterThanOrEquals, OperatorLessThanOrEquals, OperatorGreaterThan, OperatorLessThan, OperatorNotEquals, OperatorEquals}
	if !logzio_client.Contains(validOperations, alert.Operation) {
		return fmt.Errorf("operation must be one of %s", validOperations)
	}

	validSeverities := []string{
		SeveritySevere,
		SeverityHigh,
		SeverityMedium,
		SeverityLow,
		SeverityInfo,
	}
	for _, tier := range alert.SeverityThresholdTiers {
		if !logzio_client.Contains(validSeverities, tier.Severity) {
			return fmt.Errorf("severity must be one of %s", validSeverities)
		}
	}

	if AggregationTypeNone == alert.ValueAggregationType && (alert.ValueAggregationField != nil || alert.GroupByAggregationFields != nil) {
		message := fmt.Sprintf("if ValueAggregaionType is %s then ValueAggregationField and GroupByAggregationFields must be nil", AggregationTypeNone)
		return FieldError{"valueAggregationTypeComposite", message}
	}

	return nil
}

func buildCreateAlertRequest(alert CreateAlertType) map[string]interface{} {
	var createAlert = map[string]interface{}{}
	createAlert[fldAlertNotificationEndpoints] = alert.AlertNotificationEndpoints
	createAlert[fldDescription] = alert.Description
	if len(alert.Filter) > 0 {
		createAlert[fldFilter] = alert.Filter
	}
	createAlert[fldGroupByAggregationFields] = alert.GroupByAggregationFields
	createAlert[fldIsEnabled] = alert.IsEnabled
	createAlert[fldQueryString] = alert.QueryString
	createAlert[fldNotificationEmails] = alert.NotificationEmails
	createAlert[fldOperation] = alert.Operation
	createAlert[fldSearchTimeFrameMinutes] = alert.SearchTimeFrameMinutes
	createAlert[fldSeverityThresholdTiers] = alert.SeverityThresholdTiers
	createAlert[fldSuppressNotificationsMinutes] = alert.SuppressNotificationsMinutes
	createAlert[fldTitle] = alert.Title
	createAlert[fldValueAggregationField] = alert.ValueAggregationField
	createAlert[fldValueAggregationType] = alert.ValueAggregationType
	return createAlert
}

func (c *AlertsClient) buildCreateApiRequest(apiToken string, jsonObject map[string]interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	req, err := http.NewRequest(createAlertServiceMethod, fmt.Sprintf(createAlertServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Create an alert, return the created alert if successful, an error otherwise
func (c *AlertsClient) CreateAlert(alert CreateAlertType) (*AlertType, error) {
	err := validateCreateAlertRequest(alert)
	if err != nil {
		return nil, err
	}

	createAlert := buildCreateAlertRequest(alert)
	req, _ := c.buildCreateApiRequest(c.ApiToken, createAlert)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{createAlertMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "CreateAlert", resp.StatusCode, jsonBytes)
	}

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonResponse)

	retVal := jsonAlertToAlert(jsonResponse)

	if err != nil {
		return nil, err
	}

	return &retVal, nil
}
