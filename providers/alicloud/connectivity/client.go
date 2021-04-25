package connectivity

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pvtz"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/denverdino/aliyungo/cs"
	"github.com/hashicorp/terraform/terraform"
)

// AliyunClient of aliyun
type AliyunClient struct {
	Region                       Region
	RegionID                     string
	AccessKey                    string
	SecretKey                    string
	SecurityToken                string
	OtsInstanceName              string
	config                       *Config
	accountID                    string
	ecsconn                      *ecs.Client
	rdsconn                      *rds.Client
	vpcconn                      *vpc.Client
	slbconn                      *slb.Client
	dnsconn                      *alidns.Client
	ramconn                      *ram.Client
	pvtzconn                     *pvtz.Client
	tablestoreconnByInstanceName map[string]*tablestore.TableStoreClient
	csprojectconnByKey           map[string]*cs.ProjectClient
}

type APIVersion string

const DefaultClientRetryCountSmall = 5
const Terraform = "HashiCorp-Terraform"
const Provider = "Terraform-Provider"
const Module = "Terraform-Module"

var goSdkMutex = sync.RWMutex{} // The Go SDK is not thread-safe
// The main version number that is being run at the moment.
var providerVersion = "1.57.1"
var terraformVersion = strings.TrimSuffix(terraform.VersionString(), "-dev") //nolint

// Client for AliyunClient
func (c *Config) Client() (*AliyunClient, error) {
	// Get the auth and region. This can fail if keys/regions were not
	// specified and we're attempting to use the environment.
	if !c.SkipRegionValidation {
		err := c.loadAndValidate()
		if err != nil {
			return nil, err
		}
	}

	return &AliyunClient{
		config:                       c,
		Region:                       c.Region,
		RegionID:                     c.RegionID,
		AccessKey:                    c.AccessKey,
		SecretKey:                    c.SecretKey,
		SecurityToken:                c.SecurityToken,
		OtsInstanceName:              c.OtsInstanceName,
		accountID:                    c.AccountID,
		tablestoreconnByInstanceName: make(map[string]*tablestore.TableStoreClient),
		csprojectconnByKey:           make(map[string]*cs.ProjectClient),
	}, nil
}

func (client *AliyunClient) WithEcsClient(do func(*ecs.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the ECS client if necessary
	if client.ecsconn == nil {
		endpoint := client.config.EcsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionID, ECSCode)
		}

		if endpoint != "" {
			err := endpoints.AddEndpointMapping(client.config.RegionID, string(ECSCode), endpoint)
			if err != nil {
				return nil, err
			}
		}

		ecsconn, err := ecs.NewClientWithOptions(client.config.RegionID, client.getSdkConfig().WithTimeout(time.Duration(60)*time.Second), client.config.getAuthCredential())

		if err != nil {
			return nil, fmt.Errorf("unable to initialize the ECS client: %#v", err)
		}

		if _, err := ecsconn.DescribeRegions(ecs.CreateDescribeRegionsRequest()); err != nil {
			return nil, err
		}
		ecsconn.AppendUserAgent(Terraform, terraformVersion)
		ecsconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			ecsconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.ecsconn = ecsconn
	}

	return do(client.ecsconn)
}

func (client *AliyunClient) WithRdsClient(do func(*rds.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the RDS client if necessary
	if client.rdsconn == nil {
		endpoint := client.config.RdsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionID, RDSCode)
		}
		if endpoint != "" {
			err := endpoints.AddEndpointMapping(client.config.RegionID, string(RDSCode), endpoint)
			if err != nil {
				return nil, err
			}
		}
		rdsconn, err := rds.NewClientWithOptions(client.config.RegionID, client.getSdkConfig(), client.config.getAuthCredential())
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the RDS client: %#v", err)
		}

		rdsconn.AppendUserAgent(Terraform, terraformVersion)
		rdsconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			rdsconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.rdsconn = rdsconn
	}

	return do(client.rdsconn)
}

func (client *AliyunClient) WithSlbClient(do func(*slb.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the SLB client if necessary
	if client.slbconn == nil {
		endpoint := client.config.SlbEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionID, SLBCode)
		}
		if endpoint != "" {
			err := endpoints.AddEndpointMapping(client.config.RegionID, string(SLBCode), endpoint)
			if err != nil {
				return nil, err
			}
		}
		slbconn, err := slb.NewClientWithOptions(client.config.RegionID, client.getSdkConfig(), client.config.getAuthCredential())
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the SLB client: %#v", err)
		}

		slbconn.AppendUserAgent(Terraform, terraformVersion)
		slbconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			slbconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.slbconn = slbconn
	}

	return do(client.slbconn)
}

func (client *AliyunClient) WithVpcClient(do func(*vpc.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the VPC client if necessary
	if client.vpcconn == nil {
		endpoint := client.config.VpcEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionID, VPCCode)
		}
		if endpoint != "" {
			err := endpoints.AddEndpointMapping(client.config.RegionID, string(VPCCode), endpoint)
			if err != nil {
				return nil, err
			}
		}
		vpcconn, err := vpc.NewClientWithOptions(client.config.RegionID, client.getSdkConfig(), client.config.getAuthCredential())
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the VPC client: %#v", err)
		}

		vpcconn.AppendUserAgent(Terraform, terraformVersion)
		vpcconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			vpcconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.vpcconn = vpcconn
	}

	return do(client.vpcconn)
}

func (client *AliyunClient) WithDNSClient(do func(*alidns.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the DNS client if necessary
	if client.dnsconn == nil {
		endpoint := client.config.DNSEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionID, DNSCode)
		}
		if endpoint != "" {
			err := endpoints.AddEndpointMapping(client.config.RegionID, string(DNSCode), endpoint)
			if err != nil {
				return nil, err
			}
		}

		dnsconn, err := alidns.NewClientWithOptions(client.config.RegionID, client.getSdkConfig(), client.config.getAuthCredential())
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the DNS client: %#v", err)
		}
		dnsconn.AppendUserAgent(Terraform, terraformVersion)
		dnsconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			dnsconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.dnsconn = dnsconn
	}

	return do(client.dnsconn)
}

func (client *AliyunClient) WithRAMClient(do func(*ram.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the RAM client if necessary
	if client.ramconn == nil {
		endpoint := client.config.RAMEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionID, RAMCode)
		}
		if strings.HasPrefix(endpoint, "http") {
			endpoint = fmt.Sprintf("https://%s", strings.TrimPrefix(endpoint, "http://"))
		}
		if endpoint != "" {
			err := endpoints.AddEndpointMapping(client.config.RegionID, string(RAMCode), endpoint)
			if err != nil {
				return nil, err
			}
		}

		ramconn, err := ram.NewClientWithOptions(client.config.RegionID, client.getSdkConfig(), client.config.getAuthCredential())
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the RAM client: %#v", err)
		}
		ramconn.AppendUserAgent(Terraform, terraformVersion)
		ramconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			ramconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.ramconn = ramconn
	}

	return do(client.ramconn)
}

func (client *AliyunClient) WithPvtzClient(do func(*pvtz.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the PVTZ client if necessary
	if client.pvtzconn == nil {
		endpoint := client.config.PvtzEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionID, PVTZCode)
		}
		if endpoint != "" {
			err := endpoints.AddEndpointMapping(client.config.RegionID, string(PVTZCode), endpoint)
			if err != nil {
				return nil, err
			}
		} else {
			err := endpoints.AddEndpointMapping(client.config.RegionID, string(PVTZCode), "pvtz.aliyuncs.com")
			if err != nil {
				return nil, err
			}
		}
		pvtzconn, err := pvtz.NewClientWithOptions(client.config.RegionID, client.getSdkConfig(), client.config.getAuthCredential())
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the PVTZ client: %#v", err)
		}

		pvtzconn.AppendUserAgent(Terraform, terraformVersion)
		pvtzconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			pvtzconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.pvtzconn = pvtzconn
	}

	return do(client.pvtzconn)
}

func (client *AliyunClient) getSdkConfig() *sdk.Config {
	return sdk.NewConfig().
		WithMaxRetryTime(DefaultClientRetryCountSmall).
		WithTimeout(time.Duration(30) * time.Second).
		WithGoRoutinePoolSize(10).
		WithDebug(false).
		WithHttpTransport(client.getTransport()).
		WithScheme("HTTPS")
}

func (client *AliyunClient) getTransport() *http.Transport {
	handshakeTimeout, err := strconv.Atoi(os.Getenv("TLSHandshakeTimeout"))
	if err != nil {
		handshakeTimeout = 120
	}
	transport := &http.Transport{}
	transport.TLSHandshakeTimeout = time.Duration(handshakeTimeout) * time.Second

	// After building a new transport and it need to set http proxy to support proxy.
	proxyURL := client.getHTTPProxyURL()
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	return transport
}

func (client *AliyunClient) getHTTPProxyURL() *url.URL {
	for _, v := range []string{"HTTPS_PROXY", "https_proxy", "HTTP_PROXY", "http_proxy"} {
		value := strings.Trim(os.Getenv(v), " ")
		if value != "" {
			if !regexp.MustCompile(`^http(s)?://`).MatchString(value) {
				value = fmt.Sprintf("https://%s", value)
			}
			proxyURL, err := url.Parse(value)
			if err == nil {
				return proxyURL
			}
			break
		}
	}
	return nil
}
