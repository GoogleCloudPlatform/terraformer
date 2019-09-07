package alerts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
	"strings"
)

const updateAlertServiceUrl string = alertsServiceEndpoint + "/%d"
const updateAlertServiceMethod string = http.MethodPut
const updateAlertMethodSuccess int = 200

func buildUpdateAlertRequest(alert CreateAlertType) map[string]interface{} {
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

func (c *AlertsClient) buildUpdateApiRequest(apiToken string, alertId int64, jsonObject map[string]interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	req, err := http.NewRequest(updateAlertServiceMethod, fmt.Sprintf(updateAlertServiceUrl, baseUrl, alertId), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Updates an existing alert, based on the supplied alert identifier, using the parameters of the specified alert
// Returns the updated alert if successful, an error otherwise
func (c *AlertsClient) UpdateAlert(alertId int64, alert CreateAlertType) (*AlertType, error) {
	err := validateCreateAlertRequest(alert)
	if err != nil {
		return nil, err
	}

	createAlert := buildUpdateAlertRequest(alert)
	req, err := c.buildUpdateApiRequest(c.ApiToken, alertId, createAlert)
	if err != nil {
		return nil, err
	}

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{updateAlertMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "UpdateAlert", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, "no alert id") {
		return nil, fmt.Errorf("API call %s failed with missing alert %d, data: %s", "UpdateAlert", alertId, jsonBytes)
	}

	var target AlertType
	json.Unmarshal(jsonBytes, &target)

	return &target, nil
}
