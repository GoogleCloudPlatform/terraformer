package alerts

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
	"strings"
)

const getAlertServiceUrl string = alertsServiceEndpoint + "/%d"
const getAlertServiceMethod string = http.MethodGet
const getAlertMethodSuccess int = 200

func (c *AlertsClient) buildGetApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(getAlertServiceMethod, fmt.Sprintf(getAlertServiceUrl, baseUrl, alertId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Returns an alert given it's unique identifier, an error otherwise
func (c *AlertsClient) GetAlert(alertId int64) (*AlertType, error) {
	req, _ := c.buildGetApiRequest(c.ApiToken, alertId)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{getAlertMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "GetAlert", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, "no alert id") {
		return nil, fmt.Errorf("API call %s failed with missing alert %d, data: %s", "GetAlert", alertId, str)
	}

	var jsonAlert map[string]interface{}
	err = json.Unmarshal([]byte(jsonBytes), &jsonAlert)
	if err != nil {
		return nil, err
	}

	alert := jsonAlertToAlert(jsonAlert)

	return &alert, nil
}
