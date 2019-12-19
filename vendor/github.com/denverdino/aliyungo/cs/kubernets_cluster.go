package cs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/denverdino/aliyungo/common"
)

//modify cluster,include DeletionProtection and so on
type ModifyClusterArgs struct {
	DeletionProtection bool `json:"deletion_protection"`
}

//modify cluster
func (client *Client) ModifyCluster(clusterId string, args *ModifyClusterArgs) error {
	return client.Invoke("", http.MethodPut, "/api/v2/clusters/"+clusterId, nil, args, nil)
}

//Cluster Info
type KubernetesClusterType string

var (
	DelicatedKubernetes = KubernetesClusterType("Kubernetes")
	ManagedKubernetes   = KubernetesClusterType("ManagedKubernetes")
	ServerlessKubernetes = KubernetesClusterType("Ask")
)

type ProxyMode string

var (
	IPTables = "iptables"
	IPVS     = "ipvs"
)

type Effect string

var (
	TaintNoExecute        = "NoExecute"
	TaintNoSchedule       = "NoSchedule"
	TaintPreferNoSchedule = "PreferNoSchedule"
)

type ClusterArgs struct {
	DisableRollback bool `json:"disable_rollback"`
	TimeoutMins     int  `json:"timeout_mins"`

	Name               string                `json:"name"`
	ClusterType        KubernetesClusterType `json:"cluster_type"`
	Profile            string                `json:"profile"`
	KubernetesVersion  string                `json:"kubernetes_version"`
	DeletionProtection bool                  `json:"deletion_protection"`

	NodeCidrMask string `json:"node_cidr_mask"`
	UserCa       string `json:"user_ca"`

	OsType   string `json:"os_type"`
	Platform string `json:"platform"`

	UserData string `json:"user_data"`

	NodePortRange string `json:"node_port_range"`

	//ImageId
	ImageId           string `json:"image_id"`

	PodVswitchIds []string `json:"pod_vswitch_ids"` // eni多网卡模式下，需要传额外的vswitchid给addon

	LoginPassword string `json:"login_password"` //和KeyPair 二选一
	KeyPair       string `json:"key_pair"`       ////LoginPassword 二选一

	RegionId      common.Region `json:"region_id"`
	VpcId         string        `json:"vpcid"`
	ContainerCidr string        `json:"container_cidr"`
	ServiceCidr   string        `json:"service_cidr"`

	CloudMonitorFlags bool `json:"cloud_monitor_flags"`

	SecurityGroupId      string    `json:"security_group_id"`
	EndpointPublicAccess bool      `json:"endpoint_public_access"`
	ProxyMode            ProxyMode `json:"proxy_mode"`
	SnatEntry            bool      `json:"snat_entry"`

	Addons []Addon `json:"addons"`
	Tags   []Tag   `json:"tags"`

	Taints []Taint `json:"taints"`
}

//addon
type Addon struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Config  string `json:"config"`
}

//taint
type Taint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect Effect `json:"effect"`
}

type MasterArgs struct {
	MasterCount         int      `json:"master_count"`
	MasterVSwitchIds    []string `json:"master_vswitch_ids"`
	MasterInstanceTypes []string `json:"master_instance_types"`

	MasterInstanceChargeType string `json:"master_instance_charge_type"`
	MasterPeriod             int    `json:"master_period"`
	MasterPeriodUnit         string `json:"master_period_unit"`

	MasterAutoRenew       bool `json:"master_auto_renew"`
	MasterAutoRenewPeriod int  `json:"master_auto_renew_period"`

	MasterSystemDiskCategory string `json:"master_system_disk_category"`
	MasterSystemDiskSize     int64  `json:"master_system_disk_size"`

	MasterDataDisks []DataDisk `json:"master_data_disks"` //支持多个数据盘

	//support hpc/scc
	MasterHpcClusterId    string `json:"master_hpc_cluster_id"`
	MasterDeploymentSetId string `json:"master_deploymentset_id"`

	//master node deletion protection
	MasterDeletionProtection *bool `json:"master_deletion_protection"`

	// disk snapshot policy
	MasterSnapshotPolicyId string `json:"master_system_disk_snapshot_policy_id"`
}


type WorkerArgs struct {
	WorkerVSwitchIds    []string `json:"worker_vswitch_ids"`
	WorkerInstanceTypes []string `json:"worker_instance_types"`

	NumOfNodes int64 `json:"num_of_nodes"`

	WorkerInstanceChargeType string `json:"worker_instance_charge_type"`
	WorkerPeriod             int    `json:"worker_period"`
	WorkerPeriodUnit         string `json:"worker_period_unit"`

	WorkerAutoRenew       bool `json:"worker_auto_renew"`
	WorkerAutoRenewPeriod int  `json:"worker_auto_renew_period"`

	WorkerDataDisk  bool       `json:"worker_data_disk"`
	WorkerDataDisks []DataDisk `json:"worker_data_disks"` //支持多个数据盘

	WorkerHpcClusterId    string `json:"worker_hpc_cluster_id"`
	WorkerDeploymentSetId string `json:"worker_deploymentset_id"`

	//worker node deletion protection
	WorkerDeletionProtection *bool `json:"worker_deletion_protection"`

	//Runtime only for worker nodes
	Runtime Runtime `json:"runtime"`

	// disk snapshot policy
	WorkerSnapshotPolicyId string `json:"worker_system_disk_snapshot_policy_id"`
}

type ScaleOutKubernetesClusterRequest struct {
	LoginPassword string `json:"login_password"` //和KeyPair 二选一
	KeyPair       string `json:"key_pair"`       ////LoginPassword 二选一

	WorkerVSwitchIds    []string `json:"worker_vswitch_ids"`
	WorkerInstanceTypes []string `json:"worker_instance_types"`

	WorkerInstanceChargeType string `json:"worker_instance_charge_type"`
	WorkerPeriod             int    `json:"worker_period"`
	WorkerPeriodUnit         string `json:"worker_period_unit"`

	WorkerAutoRenew       bool `json:"worker_auto_renew"`
	WorkerAutoRenewPeriod int  `json:"worker_auto_renew_period"`

	WorkerDataDisk  bool       `json:"worker_data_disk"`
	WorkerDataDisks []DataDisk `json:"worker_data_disks"` //支持多个数据盘

	Tags   []Tag   `json:"tags"`
	Taints []Taint `json:"taints"`
	ImageId           string `json:"image_id"`

	UserData string `json:"user_data"`

	Count int64 `json:"count"`
}

//datadiks
type DataDisk struct {
	Category             string `json:"category"`
	KMSKeyId             string `json:"kms_key_id"`
	Encrypted            string `json:"encrypted"` // true|false
	Device               string `json:"device"`    //  could be /dev/xvd[a-z]. If not specification, will use default value.
	Size                 string `json:"size"`
	DiskName             string `json:"name"`
	AutoSnapshotPolicyId string `json:"auto_snapshot_policy_id"`
}

//runtime
type Runtime struct {
	Name                       string   `json:"name"`
	Version                    string   `json:"version"`
	RuntimeClass               []string `json:"runtimeClass,omitempty"`
	Exist                      bool     `json:"exist"`
	AvailableNetworkComponents []string `json:"availableNetworkComponents,omitempty"`
}

//DelicatedKubernetes
type DelicatedKubernetesClusterCreationRequest struct {
	ClusterArgs
	MasterArgs
	WorkerArgs
}

//ManagedKubernetes
type ManagedKubernetesClusterCreationRequest struct {
	ClusterArgs
	WorkerArgs
}

//Validate

//Create DelicatedKubernetes Cluster
func (client *Client) CreateDelicatedKubernetesCluster(request *DelicatedKubernetesClusterCreationRequest) (*ClusterCommonResponse, error) {
	response := &ClusterCommonResponse{}
	err := client.Invoke(request.RegionId, http.MethodPost, "/clusters", nil, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//Create ManagedKubernetes Cluster
func (client *Client) CreateManagedKubernetesCluster(request *ManagedKubernetesClusterCreationRequest) (*ClusterCommonResponse, error) {
	response := &ClusterCommonResponse{}
	err := client.Invoke(request.RegionId, http.MethodPost, "/clusters", nil, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//ScaleKubernetesCluster
func (client *Client) ScaleOutKubernetesCluster(clusterId string, request *ScaleOutKubernetesClusterRequest) (*ClusterCommonResponse, error) {
	response := &ClusterCommonResponse{}
	err := client.Invoke("", http.MethodPost, fmt.Sprintf("/api/v2/clusters/%s", clusterId), nil, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//DeleteClusterNodes
type DeleteKubernetesClusterNodesRequest struct {
	ReleaseNode bool     `json:"release_node"` //if set to true, the ecs instance will be released
	Nodes       []string `json:"nodes"`        //the format is regionId.instanceId|Ip ,for example  cn-hangzhou.192.168.1.2 or cn-hangzhou.i-abc
	DrainNode   bool     `json:"drain_node"`   //same as Nodes
}

//DeleteClusterNodes
func (client *Client) DeleteKubernetesClusterNodes(clusterId string, request *DeleteKubernetesClusterNodesRequest) (*common.Response, error) {
	response := &common.Response{}
	err := client.Invoke("", http.MethodPost, fmt.Sprintf("/api/v2/clusters/%s/nodes/remove", clusterId), nil, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//Cluster defination
type KubernetesClusterDetail struct {
	RegionId common.Region `json:"region_id"`

	Name        string                `json:"name"`
	ClusterId   string                `json:"cluster_id"`
	Size        int64                 `json:"size"`
	ClusterType KubernetesClusterType `json:"cluster_type"`
	Profile     string                `json:"profile"`

	VpcId                 string `json:"vpc_id"`
	VSwitchIds            string `json:"vswitch_id"`
	SecurityGroupId       string `json:"security_group_id"`
	IngressLoadbalancerId string `json:"external_loadbalancer_id"`
	ResourceGroupId       string `json:"resource_group_id"`
	NetworkMode           string `json:"network_mode"`
	ContainerCIDR         string `json:"subnet_cidr"`

	Tags  []Tag  `json:"tags"`
	State string `json:"state"`

	InitVersion        string `json:"init_version"`
	CurrentVersion     string `json:"current_version"`
	PrivateZone        bool   `json:"private_zone"`
	DeletionProtection bool   `json:"deletion_protection"`
	MetaData           string `json:"meta_data"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

//GetMetaData
func (c *KubernetesClusterDetail) GetMetaData() map[string]interface{} {
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(c.MetaData), &m)
	return m
}

//查询集群详情
func (client *Client) DescribeKubernetesClusterDetail(clusterId string) (*KubernetesClusterDetail, error) {
	cluster := &KubernetesClusterDetail{}
	err := client.Invoke("", http.MethodGet, "/clusters/"+clusterId, nil, nil, cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

//DeleteKubernetesCluster
func (client *Client) DeleteKubernetesCluster(clusterId string) error {
	return client.Invoke("", http.MethodDelete, "/clusters/"+clusterId, nil, nil, nil)
}