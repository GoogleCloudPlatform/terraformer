package cs

import (
	"fmt"
	"github.com/denverdino/aliyungo/common"
	"net/http"
)

type ClusterTokenReqeust struct {
	//token 过期时间,如果不填写，则默认24小时过期
	Expired       int64 `json:"expired"`
	IsPermanently bool  `json:"is_permanently"`
}

type ClusterTokenResponse struct {
	Created int64 ` json:"created"`
	Updated int64 `json:"updated"`
	Expired int64 ` json:"expired"`

	ClusterID string ` json:"cluster_id"`

	Token    string ` json:"token"`
	IsActive int    ` json:"is_active"`
}

func (client *Client) CreateClusterToken(clusterId string, request *ClusterTokenReqeust) (*ClusterTokenResponse, error) {
	response := &ClusterTokenResponse{}
	err := client.Invoke("", http.MethodPost, "/clusters/"+clusterId+"/token", nil, request, response)
	return response, err
}

func (client *Client) RevokeToken(token string) error {
	return client.Invoke("", http.MethodDelete, "/token/"+token+"/revoke", nil, nil, nil)
}

func (client *Client) DescribeClusterTokens(clusterId string) ([]*ClusterTokenResponse, error) {
	response := make([]*ClusterTokenResponse, 0)
	err := client.Invoke("", http.MethodGet, "/clusters/"+clusterId+"/tokens", nil, nil, &response)
	return response, err
}

func (client *Client) DescribeClusterToken(clusterId, token string) (*ClusterTokenResponse, error) {
	if clusterId == "" || token == "" {
		return nil, common.GetCustomError("InvalidParamter", "The clusterId or token is empty")
	}
	tokenInfo := &ClusterTokenResponse{}
	err := client.Invoke("", http.MethodGet, fmt.Sprintf("/clusters/%s/tokens/%s", clusterId, token), nil, nil, tokenInfo)
	return tokenInfo, err
}
