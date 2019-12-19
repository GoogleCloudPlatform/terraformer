package cs

import (
	"github.com/denverdino/aliyungo/common"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type ServerlessCreationArgs struct {
	ClusterType          KubernetesClusterType `json:"cluster_type"`
	Name                 string `json:"name"`
	RegionId             string `json:"region_id"`
	VpcId                string `json:"vpc_id"`
	VSwitchId            string `json:"vswitch_id"`
	EndpointPublicAccess bool   `json:"public_slb"`
	PrivateZone          bool   `json:"private_zone"`
	NatGateway           bool   `json:"nat_gateway"`
	DeletionProtection   bool   `json:"deletion_protection"`
	Tags                 []Tag  `json:"tags"`
}

type ServerlessClusterResponse struct {
	ClusterId          string       `json:"cluster_id"`
	Name               string       `json:"name"`
	ClusterType        KubernetesClusterType       `json:"cluster_type"`
	RegionId           string       `json:"region_id"`
	State              ClusterState `json:"state"`
	VpcId              string       `json:"vpc_id"`
	VSwitchId          string       `json:"vswitch_id"`
	SecurityGroupId    string       `json:"security_group_id"`
	Tags               []Tag        `json:"tags"`
	Created            time.Time    `json:"created"`
	Updated            time.Time    `json:"updated"`
	InitVersion        string       `json:"init_version"`
	CurrentVersion     string       `json:"current_version"`
	PrivateZone        bool         `json:"private_zone"`
	DeletionProtection bool         `json:"deletion_protection"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (this *ServerlessCreationArgs) Validate() error {
	if this.Name == "" || this.RegionId == "" || this.VpcId == "" || this.VSwitchId == "" {
		return common.GetCustomError("InvalidParameters", "The name,region_id,vpc_id,vswitch_id not allowed empty")
	}
	return nil
}

//create Serverless cluster
func (client *Client) CreateServerlessKubernetesCluster(args *ServerlessCreationArgs) (*ClusterCommonResponse, error) {
	if args == nil {
		return nil, common.GetCustomError("InvalidArgs", "The args is nil ")
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	//reset clusterType,
	args.ClusterType = ServerlessKubernetes
	cluster := &ClusterCommonResponse{}
	err := client.Invoke(common.Region(args.RegionId), http.MethodPost, "/clusters", nil, args, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

//describe Serverless cluster
func (client *Client) DescribeServerlessKubernetesCluster(clusterId string) (*ServerlessClusterResponse, error) {
	cluster := &ServerlessClusterResponse{}
	err := client.Invoke("", http.MethodGet, "/clusters/"+clusterId, nil, nil, cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

//describe Serverless clsuters
func (client *Client) DescribeServerlessKubernetesClusters() ([]*ServerlessClusterResponse, error) {
	allClusters := make([]*ServerlessClusterResponse, 0)
	askClusters := make([]*ServerlessClusterResponse, 0)

	err := client.Invoke("", http.MethodGet, "/clusters", nil, nil, &allClusters)
	if err != nil {
		return askClusters, err
	}

	for _, cluster := range allClusters {
		if cluster.ClusterType == ClusterTypeServerlessKubernetes {
			askClusters = append(askClusters, cluster)
		}
	}

	return askClusters, nil
}

//new api for get cluster kube user config
func (client *Client) DescribeClusterUserConfig(clusterId string, privateIpAddress bool) (*ClusterConfig, error) {
	config := &ClusterConfig{}
	query := url.Values{}
	query.Add("PrivateIpAddress", strconv.FormatBool(privateIpAddress))

	err := client.Invoke("", http.MethodGet, "/k8s/"+clusterId+"/user_config", query, nil, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
