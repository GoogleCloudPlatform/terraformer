package alerts

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const listAlertServiceUrl string = alertsServiceEndpoint
const listAlertServiceMethod string = http.MethodGet
const listAlertMethodSuccess int = 200

func (c *AlertsClient) buildListApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(listAlertServiceMethod, fmt.Sprintf(listAlertServiceUrl, baseUrl), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Returns all the alerts in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *AlertsClient) ListAlerts() ([]AlertType, error) {
	req, _ := c.buildListApiRequest(c.ApiToken)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{listAlertMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "ListAlerts", resp.StatusCode, jsonBytes)
	}

	var arr []AlertType
	var jsonResponse []interface{}
	err = json.Unmarshal([]byte(jsonBytes), &jsonResponse)

	for x := 0; x < len(jsonResponse); x++ {
		var jsonAlert map[string]interface{}
		jsonAlert = jsonResponse[x].(map[string]interface{})
		alert := jsonAlertToAlert(jsonAlert)
		arr = append(arr, alert)
	}

	return arr, nil
}
