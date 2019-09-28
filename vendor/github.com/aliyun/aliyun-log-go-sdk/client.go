package sls

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

// GlobalForceUsingHTTP if GlobalForceUsingHTTP is true, then all request will use HTTP(ignore LogProject's UsingHTTP flag)
var GlobalForceUsingHTTP = false

// compress type
const (
	Compress_LZ4  = iota // 0
	Compress_None        // 1
	Compress_Max         // max compress type(just for filter invalid compress type)
)

var InvalidCompressError = errors.New("Invalid Compress Type")

const defaultLogUserAgent = "golang-sdk-v0.1.0"

// Error defines sls error
type Error struct {
	HTTPCode  int32  `json:"httpCode"`
	Code      string `json:"errorCode"`
	Message   string `json:"errorMessage"`
	RequestID string `json:"requestID"`
}

// NewClientError new client error
func NewClientError(err error) *Error {
	if clientError, ok := err.(*Error); ok {
		return clientError
	}
	clientError := new(Error)
	clientError.HTTPCode = -1
	clientError.Code = "ClientError"
	clientError.Message = err.Error()
	return clientError
}

func (e Error) String() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (e Error) Error() string {
	return e.String()
}

func IsTokenError(err error) bool {
	if clientErr, ok := err.(*Error); ok {
		if clientErr.HTTPCode == 401 {
			return true
		}
	}
	return false
}

// Client ...
type Client struct {
	Endpoint        string // IP or hostname of SLS endpoint
	AccessKeyID     string
	AccessKeySecret string
	SecurityToken   string
	UserAgent       string // default defaultLogUserAgent

	accessKeyLock sync.RWMutex
}

func convert(c *Client, projName string) *LogProject {
	c.accessKeyLock.RLock()
	defer c.accessKeyLock.RUnlock()
	p, _ := NewLogProject(projName, c.Endpoint, c.AccessKeyID, c.AccessKeySecret)
	p.SecurityToken = c.SecurityToken
	p.UserAgent = c.UserAgent
	return p
}

// ResetAccessKeyToken reset client's access key token
func (c *Client) ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string) {
	c.accessKeyLock.Lock()
	c.AccessKeyID = accessKeyID
	c.AccessKeySecret = accessKeySecret
	c.SecurityToken = securityToken
	c.accessKeyLock.Unlock()
}

// CreateProject create a new loghub project.
func (c *Client) CreateProject(name, description string) (*LogProject, error) {
	type Body struct {
		ProjectName string `json:"projectName"`
		Description string `json:"description"`
	}
	body, err := json.Marshal(Body{
		ProjectName: name,
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%d", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}

	uri := "/"
	proj := convert(c, name)
	resp, err := request(proj, "POST", uri, h, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return proj, nil
}

// UpdateProject create a new loghub project.
func (c *Client) UpdateProject(name, description string) (*LogProject, error) {
	type Body struct {
		Description string `json:"description"`
	}
	body, err := json.Marshal(Body{
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%d", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}
	uri := "/"
	proj := convert(c, name)
	_, err = request(proj, "PUT", uri, h, body)
	if err != nil {
		return nil, err
	}

	return proj, nil
}

// GetProject ...
func (c *Client) GetProject(name string) (*LogProject, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}

	uri := "/"
	proj := convert(c, name)
	resp, err := request(proj, "GET", uri, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, err
	}
	err = json.Unmarshal(buf, proj)
	return proj, err
}

// ListProject list all projects in specific region
// the region is related with the client's endpoint
func (c *Client) ListProject() (projectNames []string, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}

	uri := "/"
	proj := convert(c, "")

	type Project struct {
		ProjectName string `json:"projectName"`
	}

	type Body struct {
		Projects []Project `json:"projects"`
	}

	r, err := request(proj, "GET", uri, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}

	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, err
	}

	body := &Body{}
	err = json.Unmarshal(buf, body)
	for _, project := range body.Projects {
		projectNames = append(projectNames, project.ProjectName)
	}
	return projectNames, err
}

// ListProjectV2 list all projects in specific region
// the region is related with the client's endpoint
// ref https://www.alibabacloud.com/help/doc-detail/74955.htm
func (c *Client) ListProjectV2(offset, size int) (projects []LogProject, count, total int, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}

	urlVal := url.Values{}
	urlVal.Add("offset", strconv.Itoa(offset))
	urlVal.Add("size", strconv.Itoa(size))
	uri := fmt.Sprintf("/?%s", urlVal.Encode())
	proj := convert(c, "")

	type Body struct {
		Projects []LogProject `json:"projects"`
		Count    int          `json:"count"`
		Total    int          `json:"total"`
	}

	fmt.Println("xxx", uri, h)
	r, err := request(proj, "GET", uri, h, nil)
	fmt.Printf("yyy")
	if err != nil {
		return nil, 0, 0, NewClientError(err)
	}

	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, 0, 0, err
	}

	body := &Body{}
	err = json.Unmarshal(buf, body)
	return body.Projects, body.Count, body.Total, err
}

// CheckProjectExist check project exist or not
func (c *Client) CheckProjectExist(name string) (bool, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := "/"
	proj := convert(c, name)
	resp, err := request(proj, "GET", uri, h, nil)
	if err != nil {
		if _, ok := err.(*Error); ok {
			slsErr := err.(*Error)
			if slsErr.Code == "ProjectNotExist" {
				return false, nil
			}
			return false, slsErr
		}
		return false, err
	}
	defer resp.Body.Close()
	return true, nil
}

// DeleteProject ...
func (c *Client) DeleteProject(name string) error {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}

	proj := convert(c, name)
	uri := "/"
	resp, err := request(proj, "DELETE", uri, h, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
