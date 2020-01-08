package cloudflare

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// ProxyProtocol implements json.Unmarshaler in order to support deserializing of the deprecated boolean
// value for `proxy_protocol`
type ProxyProtocol string

// UnmarshalJSON handles deserializing of both the deprecated boolean value and the current string value
// for the `proxy_protocol` field.
func (p *ProxyProtocol) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch pp := raw.(type) {
	case string:
		*p = ProxyProtocol(pp)
	case bool:
		if pp {
			*p = "v1"
		} else {
			*p = "off"
		}
	default:
		return fmt.Errorf("invalid type for proxy_protocol field: %T", pp)
	}
	return nil
}

// SpectrumApplication defines a single Spectrum Application.
type SpectrumApplication struct {
	ID            string                        `json:"id,omitempty"`
	Protocol      string                        `json:"protocol,omitempty"`
	IPv4          bool                          `json:"ipv4,omitempty"`
	DNS           SpectrumApplicationDNS        `json:"dns,omitempty"`
	OriginDirect  []string                      `json:"origin_direct,omitempty"`
	OriginPort    int                           `json:"origin_port,omitempty"`
	OriginDNS     *SpectrumApplicationOriginDNS `json:"origin_dns,omitempty"`
	IPFirewall    bool                          `json:"ip_firewall,omitempty"`
	ProxyProtocol ProxyProtocol                 `json:"proxy_protocol,omitempty"`
	TLS           string                        `json:"tls,omitempty"`
	TrafficType   string                        `json:"traffic_type,omitempty"`
	CreatedOn     *time.Time                    `json:"created_on,omitempty"`
	ModifiedOn    *time.Time                    `json:"modified_on,omitempty"`
}

// UnmarshalJSON handles setting the `ProxyProtocol` field based on the value of the deprecated `spp` field.
func (a *SpectrumApplication) UnmarshalJSON(data []byte) error {
	var body map[string]interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	var app spectrumApplicationRaw
	if err := json.Unmarshal(data, &app); err != nil {
		return err
	}

	if spp, ok := body["spp"]; ok && spp.(bool) == true {
		app.ProxyProtocol = "simple"
	}

	*a = SpectrumApplication(app)
	return nil
}

// spectrumApplicationRaw is used to inspect an application body to support the deprecated boolean value for `spp`
type spectrumApplicationRaw SpectrumApplication

// SpectrumApplicationDNS holds the external DNS configuration for a Spectrum
// Application.
type SpectrumApplicationDNS struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// SpectrumApplicationOriginDNS holds the origin DNS configuration for a Spectrum
// Application.
type SpectrumApplicationOriginDNS struct {
	Name string `json:"name"`
}

// SpectrumApplicationDetailResponse is the structure of the detailed response
// from the API.
type SpectrumApplicationDetailResponse struct {
	Response
	Result SpectrumApplication `json:"result"`
}

// SpectrumApplicationsDetailResponse is the structure of the detailed response
// from the API.
type SpectrumApplicationsDetailResponse struct {
	Response
	Result []SpectrumApplication `json:"result"`
}

// SpectrumApplications fetches all of the Spectrum applications for a zone.
//
// API reference: https://developers.cloudflare.com/spectrum/api-reference/#list-spectrum-applications
func (api *API) SpectrumApplications(zoneID string) ([]SpectrumApplication, error) {
	uri := "/zones/" + zoneID + "/spectrum/apps"

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []SpectrumApplication{}, errors.Wrap(err, errMakeRequestError)
	}

	var spectrumApplications SpectrumApplicationsDetailResponse
	err = json.Unmarshal(res, &spectrumApplications)
	if err != nil {
		return []SpectrumApplication{}, errors.Wrap(err, errUnmarshalError)
	}

	return spectrumApplications.Result, nil
}

// SpectrumApplication fetches a single Spectrum application based on the ID.
//
// API reference: https://developers.cloudflare.com/spectrum/api-reference/#list-spectrum-applications
func (api *API) SpectrumApplication(zoneID string, applicationID string) (SpectrumApplication, error) {
	uri := fmt.Sprintf(
		"/zones/%s/spectrum/apps/%s",
		zoneID,
		applicationID,
	)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return SpectrumApplication{}, errors.Wrap(err, errMakeRequestError)
	}

	var spectrumApplication SpectrumApplicationDetailResponse
	err = json.Unmarshal(res, &spectrumApplication)
	if err != nil {
		return SpectrumApplication{}, errors.Wrap(err, errUnmarshalError)
	}

	return spectrumApplication.Result, nil
}

// CreateSpectrumApplication creates a new Spectrum application.
//
// API reference: https://developers.cloudflare.com/spectrum/api-reference/#create-a-spectrum-application
func (api *API) CreateSpectrumApplication(zoneID string, appDetails SpectrumApplication) (SpectrumApplication, error) {
	uri := "/zones/" + zoneID + "/spectrum/apps"

	res, err := api.makeRequest("POST", uri, appDetails)
	if err != nil {
		return SpectrumApplication{}, errors.Wrap(err, errMakeRequestError)
	}

	var spectrumApplication SpectrumApplicationDetailResponse
	err = json.Unmarshal(res, &spectrumApplication)
	if err != nil {
		return SpectrumApplication{}, errors.Wrap(err, errUnmarshalError)
	}

	return spectrumApplication.Result, nil
}

// UpdateSpectrumApplication updates an existing Spectrum application.
//
// API reference: https://developers.cloudflare.com/spectrum/api-reference/#update-a-spectrum-application
func (api *API) UpdateSpectrumApplication(zoneID, appID string, appDetails SpectrumApplication) (SpectrumApplication, error) {
	uri := fmt.Sprintf(
		"/zones/%s/spectrum/apps/%s",
		zoneID,
		appID,
	)

	res, err := api.makeRequest("PUT", uri, appDetails)
	if err != nil {
		return SpectrumApplication{}, errors.Wrap(err, errMakeRequestError)
	}

	var spectrumApplication SpectrumApplicationDetailResponse
	err = json.Unmarshal(res, &spectrumApplication)
	if err != nil {
		return SpectrumApplication{}, errors.Wrap(err, errUnmarshalError)
	}

	return spectrumApplication.Result, nil
}

// DeleteSpectrumApplication removes a Spectrum application based on the ID.
//
// API reference: https://developers.cloudflare.com/spectrum/api-reference/#delete-a-spectrum-application
func (api *API) DeleteSpectrumApplication(zoneID string, applicationID string) error {
	uri := fmt.Sprintf(
		"/zones/%s/spectrum/apps/%s",
		zoneID,
		applicationID,
	)

	_, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}

	return nil
}
