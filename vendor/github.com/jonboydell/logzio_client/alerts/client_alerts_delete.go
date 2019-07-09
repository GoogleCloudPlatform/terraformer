package alerts

import (
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
	"strings"
)

const deleteAlertServiceUrl string = alertsServiceEndpoint + "/%d"
const deleteAlertServiceMethod string = http.MethodDelete
const deleteAlertMethodSuccess int = 200

func (c *AlertsClient) buildDeleteApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(deleteAlertServiceMethod, fmt.Sprintf(deleteAlertServiceUrl, baseUrl, alertId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Delete an alert, specified by it's unique id, returns an error if a problem is encountered
func (c *AlertsClient) DeleteAlert(alertId int64) error {
	req, _ := c.buildDeleteApiRequest(c.ApiToken, alertId)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{deleteAlertMethodSuccess}) {
		return fmt.Errorf("API call %s failed with status code %d, data: %s", "DeleteAlert", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, "no alert id") {
		return fmt.Errorf("API call %s failed with missing alert %d, data: %s", "DeleteAlert", alertId, jsonBytes)
	}

	return nil
}
