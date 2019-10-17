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
	Region   Region
	RegionId string
	//In order to build ots table client, add accesskey and secretkey in aliyunclient temporarily.
	AccessKey                    string
	SecretKey                    string
	SecurityToken                string
	OtsInstanceName              string
	accountIdMutex               sync.RWMutex
	config                       *Config
	accountId                    string
	ecsconn                      *ecs.Client
	rdsconn                      *rds.Client
	vpcconn                      *vpc.Client
	slbconn                      *slb.Client
	dnsconn                      *alidns.Client
	ramconn                      *ram.Client
	pvtzconn                     *pvtz.Client
	tablestoreconnByInstanceName map[string]*tablestore.TableStoreClient
	csprojectconnByKey           map[string]*cs.ProjectClient
	//alikafkaconn                 *alikafka.Client
	//onsconn                      *ons.Client
	//ossconn                      *oss.Client
	//nasconn                      *nas.Client
	//csconn                       *cs.Client
	//essconn                      *ess.Client
	//cdnconn_new                  *cdn_new.Client
	//crconn                       *cr.Client
	//cdnconn                      *cdn.CdnClient
	//kmsconn                      *kms.Client
	//otsconn                      *ots.Client
	//cmsconn                      *cms.Client
	//logconn                      *sls.Client
	//fcconn                       *fc.Client
	//cenconn                      *cbn.Client
	//ddsconn                      *dds.Client
	//gpdbconn                     *gpdb.Client
	//stsconn                      *sts.Client
	//rkvconn                      *r_kvstore.Client
	//dhconn                       *datahub.DataHub
	//mnsconn                      *ali_mns.MNSClient
	//cloudapiconn                 *cloudapi.Client
	//drdsconn                     *drds.Client
	//elasticsearchconn            *elasticsearch.Client
	//actiontrailconn              *actiontrail.Client
	//casconn                      *cas.Client
	//ddoscooconn                  *ddoscoo.Client
	//ddosbgpconn                  *ddosbgp.Client
	//bssopenapiconn               *bssopenapi.Client
	//emrconn                      *emr.Client
}

type ApiVersion string

const (
	ApiVersion20140526 = ApiVersion("2014-05-26")
	ApiVersion20160815 = ApiVersion("2016-08-15")
	ApiVersion20140515 = ApiVersion("2014-05-15")
)

const businessInfoKey = "Terraform"

const DefaultClientRetryCountSmall = 5
const DefaultClientRetryCountMedium = 10
const DefaultClientRetryCountLarge = 15
const Terraform = "HashiCorp-Terraform"
const Provider = "Terraform-Provider"
const Module = "Terraform-Module"

var goSdkMutex = sync.RWMutex{} // The Go SDK is not thread-safe
// The main version number that is being run at the moment.
var providerVersion = "1.57.1"
var terraformVersion = strings.TrimSuffix(terraform.VersionString(), "-dev")

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
		RegionId:                     c.RegionId,
		AccessKey:                    c.AccessKey,
		SecretKey:                    c.SecretKey,
		SecurityToken:                c.SecurityToken,
		OtsInstanceName:              c.OtsInstanceName,
		accountId:                    c.AccountId,
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
			endpoint = loadEndpoint(client.config.RegionId, ECSCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(ECSCode), endpoint)
		}
		ecsconn, err := ecs.NewClientWithOptions(client.config.RegionId, client.getSdkConfig().WithTimeout(time.Duration(60)*time.Second), client.config.getAuthCredential(true))
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
			endpoint = loadEndpoint(client.config.RegionId, RDSCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(RDSCode), endpoint)
		}
		rdsconn, err := rds.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
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
			endpoint = loadEndpoint(client.config.RegionId, SLBCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(SLBCode), endpoint)
		}
		slbconn, err := slb.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
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
			endpoint = loadEndpoint(client.config.RegionId, VPCCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(VPCCode), endpoint)
		}
		vpcconn, err := vpc.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
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

/*
func (client *AliyunClient) WithNasClient(do func(*nas.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the Nas client if necessary
	if client.nasconn == nil {
		endpoint := loadEndpoint(client.config.RegionId, NASCode)
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(NASCode), endpoint)
		}
		nasconn, err := nas.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the NAS client: %#v", err)
		}
		nasconn.AppendUserAgent(Terraform, terraformVersion)
		nasconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			nasconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.nasconn = nasconn
	}

	return do(client.nasconn)
}

func (client *AliyunClient) WithCenClient(do func(*cbn.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the CEN client if necessary
	if client.cenconn == nil {
		endpoint := client.config.CenEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, CENCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(CENCode), endpoint)
		}
		cenconn, err := cbn.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the CEN client: %#v", err)
		}

		cenconn.AppendUserAgent(Terraform, terraformVersion)
		cenconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			cenconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.cenconn = cenconn
	}

	return do(client.cenconn)
}

func (client *AliyunClient) WithEssClient(do func(*ess.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the ESS client if necessary
	if client.essconn == nil {
		endpoint := client.config.EssEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, ESSCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(ESSCode), endpoint)
		}
		essconn, err := ess.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the ESS client: %#v", err)
		}

		essconn.AppendUserAgent(Terraform, terraformVersion)
		essconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			essconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.essconn = essconn
	}

	return do(client.essconn)
}

func (client *AliyunClient) WithOssClient(do func(*oss.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the OSS client if necessary
	if client.ossconn == nil {
		schma := "https"
		endpoint := client.config.OssEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, OSSCode)
		}
		if endpoint == "" {
			endpointItem, _ := client.describeEndpointForService(strings.ToLower(string(OSSCode)))
			if endpointItem != nil {
				if len(endpointItem.Protocols.Protocols) > 0 {
					// HTTP or HTTPS
					schma = strings.ToLower(endpointItem.Protocols.Protocols[0])
					for _, p := range endpointItem.Protocols.Protocols {
						if strings.ToLower(p) == "https" {
							schma = strings.ToLower(p)
							break
						}
					}
				}
				endpoint = endpointItem.Endpoint
			} else {
				endpoint = fmt.Sprintf("oss-%s.aliyuncs.com", client.RegionId)
			}
		}
		if !strings.HasPrefix(endpoint, "http") {
			endpoint = fmt.Sprintf("%s://%s", schma, endpoint)
		}

		clientOptions := []oss.ClientOption{oss.UserAgent(client.getUserAgent()),
			oss.SecurityToken(client.config.SecurityToken)}
		proxyUrl := client.getHttpProxyUrl()
		if proxyUrl != nil {
			clientOptions = append(clientOptions, oss.Proxy(proxyUrl.String()))
		}

		ossconn, err := oss.New(endpoint, client.config.AccessKey, client.config.SecretKey, clientOptions...)
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the OSS client: %#v", err)
		}

		client.ossconn = ossconn
	}

	return do(client.ossconn)
}

func (client *AliyunClient) WithOssBucketByName(bucketName string, do func(*oss.Bucket) (interface{}, error)) (interface{}, error) {
	return client.WithOssClient(func(ossClient *oss.Client) (interface{}, error) {
		bucket, err := client.ossconn.Bucket(bucketName)
		if err != nil {
			return nil, fmt.Errorf("unable to get the bucket %s: %#v", bucketName, err)
		}
		return do(bucket)
	})
}
*/
func (client *AliyunClient) WithDnsClient(do func(*alidns.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the DNS client if necessary
	if client.dnsconn == nil {
		endpoint := client.config.DnsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, DNSCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(DNSCode), endpoint)
		}

		dnsconn, err := alidns.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
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

func (client *AliyunClient) WithRamClient(do func(*ram.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the RAM client if necessary
	if client.ramconn == nil {
		endpoint := client.config.RamEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, RAMCode)
		}
		if strings.HasPrefix(endpoint, "http") {
			endpoint = fmt.Sprintf("https://%s", strings.TrimPrefix(endpoint, "http://"))
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(RAMCode), endpoint)
		}

		ramconn, err := ram.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
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

/*
func (client *AliyunClient) WithCsClient(do func(*cs.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the CS client if necessary
	if client.csconn == nil {
		csconn := cs.NewClientForAussumeRole(client.config.AccessKey, client.config.SecretKey, client.config.SecurityToken)
		csconn.SetUserAgent(client.getUserAgent())
		endpoint := client.config.CsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, CONTAINCode)
		}
		if endpoint != "" {
			if !strings.HasPrefix(endpoint, "http") {
				endpoint = fmt.Sprintf("https://%s", strings.TrimPrefix(endpoint, "://"))
			}
			csconn.SetEndpoint(endpoint)
		}
		client.csconn = csconn
	}

	return do(client.csconn)
}

func (client *AliyunClient) WithCrClient(do func(*cr.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the CR client if necessary
	if client.crconn == nil {
		endpoint := client.config.CrEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, CRCode)
			if endpoint == "" {
				endpoint = fmt.Sprintf("cr.%s.aliyuncs.com", client.config.RegionId)
			}
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(CRCode), endpoint)
		}
		crconn, err := cr.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the CR client: %#v", err)
		}
		crconn.AppendUserAgent(Terraform, terraformVersion)
		crconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			crconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.crconn = crconn
	}

	return do(client.crconn)
}

func (client *AliyunClient) WithCdnClient(do func(*cdn.CdnClient) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the CDN client if necessary
	if client.cdnconn == nil {
		cdnconn := cdn.NewClient(client.config.AccessKey, client.config.SecretKey)
		cdnconn.SetBusinessInfo(businessInfoKey)
		cdnconn.SetUserAgent(client.getUserAgent())
		cdnconn.SetSecurityToken(client.config.SecurityToken)
		endpoint := client.config.CdnEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, CDNCode)
		}
		if endpoint != "" && !strings.HasPrefix(endpoint, "http") {
			cdnconn.SetEndpoint(fmt.Sprintf("https://%s", strings.TrimPrefix(endpoint, "://")))
		}
		client.cdnconn = cdnconn
	}
	return do(client.cdnconn)
}

func (client *AliyunClient) WithCdnClient_new(do func(*cdn_new.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the CDN client if necessary
	if client.cdnconn_new == nil {
		endpoint := client.config.CdnEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, CDNCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(CDNCode), endpoint)
		}
		cdnconn, err := cdn_new.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the CDN client: %#v", err)
		}

		cdnconn.AppendUserAgent(Terraform, terraformVersion)
		cdnconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			cdnconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.cdnconn_new = cdnconn
	}

	return do(client.cdnconn_new)
}

func (client *AliyunClient) WithKmsClient(do func(*kms.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the KMS client if necessary
	if client.kmsconn == nil {

		endpoint := client.config.KmsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, KMSCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(KMSCode), endpoint)
		}
		kmsconn, err := kms.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the kms client: %#v", err)
		}
		kmsconn.AppendUserAgent(Terraform, terraformVersion)
		kmsconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			kmsconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.kmsconn = kmsconn
	}
	return do(client.kmsconn)
}

func (client *AliyunClient) WithOtsClient(do func(*ots.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the OTS client if necessary
	if client.otsconn == nil {
		endpoint := client.config.OtsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, OTSCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(OTSCode), endpoint)
		}
		otsconn, err := ots.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the OTS client: %#v", err)
		}

		otsconn.AppendUserAgent(Terraform, terraformVersion)
		otsconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			otsconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.otsconn = otsconn
	}

	return do(client.otsconn)
}

func (client *AliyunClient) WithCmsClient(do func(*cms.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the CMS client if necessary
	if client.cmsconn == nil {
		cmsconn, err := cms.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(false))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the CMS client: %#v", err)
		}

		cmsconn.AppendUserAgent(Terraform, terraformVersion)
		cmsconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			cmsconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.cmsconn = cmsconn
	}

	return do(client.cmsconn)
}
*/
func (client *AliyunClient) WithPvtzClient(do func(*pvtz.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the PVTZ client if necessary
	if client.pvtzconn == nil {
		endpoint := client.config.PvtzEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, PVTZCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(PVTZCode), endpoint)
		} else {
			endpoints.AddEndpointMapping(client.config.RegionId, string(PVTZCode), "pvtz.aliyuncs.com")
		}
		pvtzconn, err := pvtz.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
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

/*
func (client *AliyunClient) WithStsClient(do func(*sts.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the STS client if necessary
	if client.stsconn == nil {
		endpoint := client.config.StsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, STSCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(STSCode), endpoint)
		}
		stsconn, err := sts.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the STS client: %#v", err)
		}

		stsconn.AppendUserAgent(Terraform, terraformVersion)
		stsconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			stsconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.stsconn = stsconn
	}

	return do(client.stsconn)
}

func (client *AliyunClient) WithLogClient(do func(*sls.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the LOG client if necessary
	if client.logconn == nil {
		endpoint := client.config.LogEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, LOGCode)
			if endpoint == "" {
				endpoint = fmt.Sprintf("%s.log.aliyuncs.com", client.config.RegionId)
			}
		}
		if !strings.HasPrefix(endpoint, "http") {
			endpoint = fmt.Sprintf("https://%s", strings.TrimPrefix(endpoint, "://"))
		}
		client.logconn = &sls.Client{
			AccessKeyID:     client.config.AccessKey,
			AccessKeySecret: client.config.SecretKey,
			Endpoint:        endpoint,
			SecurityToken:   client.config.SecurityToken,
			UserAgent:       client.getUserAgent(),
		}
	}

	return do(client.logconn)
}

func (client *AliyunClient) WithDrdsClient(do func(*drds.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the DRDS client if necessary
	if client.drdsconn == nil {
		endpoint := client.config.DrdsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, DRDSCode)
			if endpoint == "" {
				endpoint = fmt.Sprintf("%s.drds.aliyuncs.com", client.config.RegionId)
			}
		}

		drdsconn, err := drds.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the DRDS client: %#v", err)

		}

		drdsconn.AppendUserAgent(Terraform, terraformVersion)
		drdsconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			drdsconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.drdsconn = drdsconn
	}

	return do(client.drdsconn)
}

func (client *AliyunClient) WithDdsClient(do func(*dds.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the DDS client if necessary
	if client.ddsconn == nil {
		endpoint := client.config.DdsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, DDSCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(DDSCode), endpoint)
		}
		ddsconn, err := dds.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the DDS client: %#v", err)
		}

		ddsconn.AppendUserAgent(Terraform, terraformVersion)
		ddsconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			ddsconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.ddsconn = ddsconn
	}

	return do(client.ddsconn)
}

func (client *AliyunClient) WithGpdbClient(do func(*gpdb.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the GPDB client if necessary
	if client.gpdbconn == nil {
		endpoint := client.config.GpdbEnpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, GPDBCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(GPDBCode), endpoint)
		}
		gpdbconn, err := gpdb.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the GPDB client: %#v", err)
		}

		gpdbconn.AppendUserAgent(Terraform, terraformVersion)
		gpdbconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			gpdbconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.gpdbconn = gpdbconn
	}

	return do(client.gpdbconn)
}

func (client *AliyunClient) WithRkvClient(do func(*r_kvstore.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the RKV client if necessary
	if client.rkvconn == nil {
		endpoint := client.config.KVStoreEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, KVSTORECode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, fmt.Sprintf("R-%s", string(KVSTORECode)), endpoint)
		}
		rkvconn, err := r_kvstore.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the RKV client: %#v", err)
		}

		rkvconn.AppendUserAgent(Terraform, terraformVersion)
		rkvconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			rkvconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.rkvconn = rkvconn
	}

	return do(client.rkvconn)
}

func (client *AliyunClient) WithFcClient(do func(*fc.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the FC client if necessary
	if client.fcconn == nil {
		endpoint := client.config.FcEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, FCCode)
			if endpoint == "" {
				endpoint = fmt.Sprintf("%s.fc.aliyuncs.com", client.config.RegionId)
			}
		}
		if strings.HasPrefix(endpoint, "http") {
			endpoint = strings.TrimPrefix(strings.TrimPrefix(endpoint, "http://"), "https://")
		}
		accountId, err := client.AccountId()
		if err != nil {
			return nil, err
		}

		config := client.getSdkConfig()
		clientOptions := []fc.ClientOption{fc.WithSecurityToken(client.config.SecurityToken), fc.WithTransport(config.HttpTransport),
			fc.WithTimeout(30), fc.WithRetryCount(DefaultClientRetryCountSmall)}
		fcconn, err := fc.NewClient(fmt.Sprintf("https://%s.%s", accountId, endpoint), string(ApiVersion20160815), client.config.AccessKey, client.config.SecretKey, clientOptions...)
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the FC client: %#v", err)
		}

		fcconn.Config.UserAgent = client.getUserAgent()
		fcconn.Config.SecurityToken = client.config.SecurityToken
		client.fcconn = fcconn
	}

	return do(client.fcconn)
}

func (client *AliyunClient) WithCloudApiClient(do func(*cloudapi.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the Cloud API client if necessary
	if client.cloudapiconn == nil {
		endpoint := client.config.ApigatewayEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.RegionId, CLOUDAPICode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.RegionId, "CLOUDAPI", endpoint)
		}
		cloudapiconn, err := cloudapi.NewClientWithOptions(client.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the CloudAPI client: %#v", err)
		}

		cloudapiconn.AppendUserAgent(Terraform, terraformVersion)
		cloudapiconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			cloudapiconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.cloudapiconn = cloudapiconn
	}

	return do(client.cloudapiconn)
}

func (client *AliyunClient) WithDataHubClient(do func(*datahub.DataHub) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the DataHub client if necessary
	if client.dhconn == nil {
		endpoint := client.config.DatahubEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.RegionId, DATAHUBCode)
		}
		if endpoint == "" {
			if client.RegionId == string(APSouthEast1) {
				endpoint = "dh-singapore.aliyuncs.com"
			} else {
				endpoint = fmt.Sprintf("dh-%s.aliyuncs.com", client.RegionId)
			}
		}
		if !strings.HasPrefix(endpoint, "http") {
			endpoint = fmt.Sprintf("https://%s", endpoint)
		}

		account := datahub.NewStsCredential(client.config.AccessKey, client.config.SecretKey, client.config.SecurityToken)
		config := &datahub.Config{
			UserAgent: client.getUserAgent(),
		}

		client.dhconn = datahub.NewClientWithConfig(endpoint, config, account)
	}

	return do(client.dhconn)
}

func (client *AliyunClient) WithMnsClient(do func(*ali_mns.MNSClient) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the MNS client if necessary
	if client.mnsconn == nil {
		endpoint := client.config.MnsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, MNSCode)
			if endpoint == "" {
				endpoint = fmt.Sprintf("%s.aliyuncs.com", client.config.RegionId)
			}
		}

		accountId, err := client.AccountId()
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(endpoint, "http") {
			endpoint = strings.TrimPrefix(strings.TrimPrefix(endpoint, "http://"), "https://")
		}
		mnsUrl := fmt.Sprintf("https://%s.mns.%s", accountId, endpoint)

		mnsClient := ali_mns.NewAliMNSClient(mnsUrl, client.config.AccessKey, client.config.SecretKey)

		client.mnsconn = &mnsClient
	}

	return do(client.mnsconn)
}

func (client *AliyunClient) WithElasticsearchClient(do func(*elasticsearch.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the Elasticsearch client if necessary
	if client.elasticsearchconn == nil {
		endpoint := client.config.ElasticsearchEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, ELASTICSEARCHCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(ELASTICSEARCHCode), endpoint)
		}
		elasticsearchconn, err := elasticsearch.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the Elasticsearch client: %#v", err)
		}

		elasticsearchconn.AppendUserAgent(Terraform, terraformVersion)
		elasticsearchconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			elasticsearchconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.elasticsearchconn = elasticsearchconn
	}

	return do(client.elasticsearchconn)
}

func (client *AliyunClient) WithMnsQueueManager(do func(ali_mns.AliQueueManager) (interface{}, error)) (interface{}, error) {
	return client.WithMnsClient(func(mnsClient *ali_mns.MNSClient) (interface{}, error) {
		queueManager := ali_mns.NewMNSQueueManager(*mnsClient)
		return do(queueManager)
	})
}

func (client *AliyunClient) WithMnsTopicManager(do func(ali_mns.AliTopicManager) (interface{}, error)) (interface{}, error) {
	return client.WithMnsClient(func(mnsClient *ali_mns.MNSClient) (interface{}, error) {
		topicManager := ali_mns.NewMNSTopicManager(*mnsClient)
		return do(topicManager)
	})
}

func (client *AliyunClient) WithMnsSubscriptionManagerByTopicName(topicName string, do func(ali_mns.AliMNSTopic) (interface{}, error)) (interface{}, error) {
	return client.WithMnsClient(func(mnsClient *ali_mns.MNSClient) (interface{}, error) {
		subscriptionManager := ali_mns.NewMNSTopic(topicName, *mnsClient)
		return do(subscriptionManager)
	})
}

func (client *AliyunClient) WithTableStoreClient(instanceName string, do func(*tablestore.TableStoreClient) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the TABLESTORE client if necessary
	tableStoreClient, ok := client.tablestoreconnByInstanceName[instanceName]
	if !ok {
		endpoint := client.config.OtsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.RegionId, OTSCode)
		}
		if endpoint == "" {
			endpoint = fmt.Sprintf("%s.%s.ots.aliyuncs.com", instanceName, client.RegionId)
		}
		if !strings.HasPrefix(endpoint, "https") && !strings.HasPrefix(endpoint, "http") {
			endpoint = fmt.Sprintf("https://%s", endpoint)
		}

		tableStoreClient = tablestore.NewClientWithConfig(endpoint, instanceName, client.config.AccessKey, client.config.SecretKey, client.config.SecurityToken, tablestore.NewDefaultTableStoreConfig())
		client.tablestoreconnByInstanceName[instanceName] = tableStoreClient
	}

	return do(tableStoreClient)
}

func (client *AliyunClient) WithCsProjectClient(clusterId, endpoint string, clusterCerts cs.ClusterCerts, do func(*cs.ProjectClient) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the PROJECT client if necessary
	key := fmt.Sprintf("%s|%s|%s|%s|%s", clusterId, endpoint, clusterCerts.CA, clusterCerts.Cert, clusterCerts.Key)
	csProjectClient, ok := client.csprojectconnByKey[key]
	if !ok {
		var err error
		csProjectClient, err = cs.NewProjectClient(clusterId, endpoint, clusterCerts)
		if err != nil {
			return nil, fmt.Errorf("Getting Application Client failed by cluster id %s: %#v.", clusterCerts, err)
		}
		csProjectClient.SetDebug(false)
		csProjectClient.SetUserAgent(client.getUserAgent())
		client.csprojectconnByKey[key] = csProjectClient
	}

	return do(csProjectClient)
}

func (client *AliyunClient) NewCommonRequest(product, serviceCode, schema string, apiVersion ApiVersion) (*requests.CommonRequest, error) {
	request := requests.NewCommonRequest()
	endpoint := loadEndpoint(client.RegionId, ServiceCode(strings.ToUpper(product)))
	if endpoint == "" {
		endpointItem, err := client.describeEndpointForService(serviceCode)
		if err != nil {
			return nil, fmt.Errorf("describeEndpointForService got an error: %#v.", err)
		}
		if endpointItem != nil {
			endpoint = endpointItem.Endpoint
		}
	}
	// Use product code to find product domain
	if endpoint != "" {
		request.Domain = endpoint
	} else {
		// When getting endpoint failed by location, using custom endpoint instead
		request.Domain = fmt.Sprintf("%s.%s.aliyuncs.com", strings.ToLower(serviceCode), client.RegionId)
	}
	request.Version = string(apiVersion)
	request.RegionId = client.RegionId
	request.Product = product
	request.Scheme = schema
	request.AppendUserAgent(Terraform, terraformVersion)
	request.AppendUserAgent(Provider, providerVersion)
	if client.config.ConfigurationSource != "" {
		request.AppendUserAgent(Module, client.config.ConfigurationSource)
	}
	return request, nil
}

func (client *AliyunClient) AccountId() (string, error) {
	client.accountIdMutex.Lock()
	defer client.accountIdMutex.Unlock()

	if client.accountId == "" {
		log.Printf("[DEBUG] account_id not provided, attempting to retrieve it automatically...")
		identity, err := client.getCallerIdentity()
		if err != nil {
			return "", err
		}
		if identity.AccountId == "" {
			return "", fmt.Errorf("caller identity doesn't contain any AccountId")
		}
		client.accountId = identity.AccountId
	}
	return client.accountId, nil
}
*/
func (client *AliyunClient) getSdkConfig() *sdk.Config {
	return sdk.NewConfig().
		WithMaxRetryTime(DefaultClientRetryCountSmall).
		WithTimeout(time.Duration(30) * time.Second).
		WithGoRoutinePoolSize(10).
		WithDebug(false).
		WithHttpTransport(client.getTransport()).
		WithScheme("HTTPS")
}

func (client *AliyunClient) getUserAgent() string {
	if client.config.ConfigurationSource != "" {
		return fmt.Sprintf("%s/%s %s/%s %s/%s", Terraform, terraformVersion, Provider, providerVersion, Module, client.config.ConfigurationSource)
	}
	return fmt.Sprintf("%s/%s %s/%s", Terraform, terraformVersion, Provider, providerVersion)
}

func (client *AliyunClient) getTransport() *http.Transport {
	handshakeTimeout, err := strconv.Atoi(os.Getenv("TLSHandshakeTimeout"))
	if err != nil {
		handshakeTimeout = 120
	}
	transport := &http.Transport{}
	transport.TLSHandshakeTimeout = time.Duration(handshakeTimeout) * time.Second

	// After building a new transport and it need to set http proxy to support proxy.
	proxyUrl := client.getHttpProxyUrl()
	if proxyUrl != nil {
		transport.Proxy = http.ProxyURL(proxyUrl)
	}
	return transport
}

func (client *AliyunClient) getHttpProxyUrl() *url.URL {
	for _, v := range []string{"HTTPS_PROXY", "https_proxy", "HTTP_PROXY", "http_proxy"} {
		value := strings.Trim(os.Getenv(v), " ")
		if value != "" {
			if !regexp.MustCompile(`^http(s)?://`).MatchString(value) {
				value = fmt.Sprintf("https://%s", value)
			}
			proxyUrl, err := url.Parse(value)
			if err == nil {
				return proxyUrl
			}
			break
		}
	}
	return nil
}

/*
func (client *AliyunClient) describeEndpointForService(serviceCode string) (*location.Endpoint, error) {
	args := location.CreateDescribeEndpointsRequest()
	args.ServiceCode = serviceCode
	args.Id = client.config.RegionId
	args.Domain = client.config.LocationEndpoint
	if args.Domain == "" {
		args.Domain = loadEndpoint(client.RegionId, LOCATIONCode)
	}
	if args.Domain == "" {
		args.Domain = "location-readonly.aliyuncs.com"
	}

	locationClient, err := location.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize the location client: %#v", err)

	}
	locationClient.AppendUserAgent(Terraform, terraformVersion)
	locationClient.AppendUserAgent(Provider, providerVersion)
	if client.config.ConfigurationSource != "" {
		locationClient.AppendUserAgent(Module, client.config.ConfigurationSource)
	}
	endpointsResponse, err := locationClient.DescribeEndpoints(args)
	if err != nil {
		return nil, fmt.Errorf("Describe %s endpoint using region: %#v got an error: %#v.", serviceCode, client.RegionId, err)
	}
	if endpointsResponse != nil && len(endpointsResponse.Endpoints.Endpoint) > 0 {
		for _, e := range endpointsResponse.Endpoints.Endpoint {
			if e.Type == "openAPI" {
				return &e, nil
			}
		}
	}
	return nil, fmt.Errorf("There is no any available endpoint for %s in region %s.", serviceCode, client.RegionId)
}

func (client *AliyunClient) getCallerIdentity() (*sts.GetCallerIdentityResponse, error) {
	args := sts.CreateGetCallerIdentityRequest()

	endpoint := client.config.StsEndpoint
	if endpoint == "" {
		endpoint = loadEndpoint(client.config.RegionId, STSCode)
	}
	if endpoint != "" {
		endpoints.AddEndpointMapping(client.config.RegionId, string(STSCode), endpoint)
	}
	stsClient, err := sts.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
	if err != nil {
		return nil, fmt.Errorf("unable to initialize the STS client: %#v", err)
	}

	stsClient.AppendUserAgent(Terraform, terraformVersion)
	stsClient.AppendUserAgent(Provider, providerVersion)
	if client.config.ConfigurationSource != "" {
		stsClient.AppendUserAgent(Module, client.config.ConfigurationSource)
	}
	identity, err := stsClient.GetCallerIdentity(args)
	if err != nil {
		return nil, err
	}
	if identity == nil {
		return nil, fmt.Errorf("caller identity not found")
	}
	return identity, err
}

func (client *AliyunClient) WithActionTrailClient(do func(*actiontrail.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()
	if client.actiontrailconn == nil {
		endpoint := client.config.ActionTrailEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, ACTIONTRAILCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(ACTIONTRAILCode), endpoint)
		}
		actiontrailconn, err := actiontrail.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the ACTIONTRAIL client: %#v", err)
		}

		actiontrailconn.AppendUserAgent(Terraform, terraformVersion)
		actiontrailconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			actiontrailconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.actiontrailconn = actiontrailconn
	}

	return do(client.actiontrailconn)
}

func (client *AliyunClient) WithCasClient(do func(*cas.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the CAS client if necessary
	if client.casconn == nil {
		casconn, err := cas.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the CAS client: %#v", err)
		}

		casconn.AppendUserAgent(Terraform, terraformVersion)
		casconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			casconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.casconn = casconn
	}

	return do(client.casconn)
}

func (client *AliyunClient) WithDdoscooClient(do func(*ddoscoo.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the ddoscoo client if necessary
	if client.ddoscooconn == nil {
		ddoscooconn, err := ddoscoo.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the DDOSCOO client: %#v", err)
		}
		ddoscooconn.AppendUserAgent(Terraform, terraformVersion)
		ddoscooconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			ddoscooconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.ddoscooconn = ddoscooconn

	}

	return do(client.ddoscooconn)
}

func (client *AliyunClient) WithDdosbgpClient(do func(*ddosbgp.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the ddosbgp client if necessary
	if client.ddosbgpconn == nil {
		ddosbgpconn, err := ddosbgp.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the DDOSBGP client: %#v", err)
		}

		ddosbgpconn.AppendUserAgent(Terraform, terraformVersion)
		ddosbgpconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			ddosbgpconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.ddosbgpconn = ddosbgpconn
	}

	return do(client.ddosbgpconn)
}

func (client *AliyunClient) WithBssopenapiClient(do func(*bssopenapi.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the bssopenapi client if necessary
	if client.bssopenapiconn == nil {
		endpoint := client.config.BssOpenApiEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, BSSOPENAPICode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(BSSOPENAPICode), endpoint)
		}

		bssopenapiconn, err := bssopenapi.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the BSSOPENAPI client: %#v", err)
		}
		bssopenapiconn.AppendUserAgent(Terraform, terraformVersion)
		bssopenapiconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			bssopenapiconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.bssopenapiconn = bssopenapiconn
	}

	return do(client.bssopenapiconn)
}

func (client *AliyunClient) WithOnsClient(do func(*ons.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the ons client if necessary
	if client.onsconn == nil {
		endpoint := client.config.OnsEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, ONSCode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(ONSCode), endpoint)
		}
		onsconn, err := ons.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the ONS client: %#v", err)
		}
		onsconn.AppendUserAgent(Terraform, terraformVersion)
		onsconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			onsconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.onsconn = onsconn
	}

	return do(client.onsconn)
}

func (client *AliyunClient) WithAlikafkaClient(do func(*alikafka.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	// Initialize the ons client if necessary
	if client.onsconn == nil {
		endpoint := client.config.AlikafkaEndpoint
		if endpoint == "" {
			endpoint = loadEndpoint(client.config.RegionId, ALIKAFKACode)
		}
		if endpoint != "" {
			endpoints.AddEndpointMapping(client.config.RegionId, string(ALIKAFKACode), endpoint)
		}
		alikafkaconn, err := alikafka.NewClientWithOptions(client.config.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the ALIKAFKA client: %#v", err)
		}
		alikafkaconn.AppendUserAgent(Terraform, terraformVersion)
		alikafkaconn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			alikafkaconn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.alikafkaconn = alikafkaconn
	}

	return do(client.alikafkaconn)
}

func (client *AliyunClient) WithEmrClient(do func(*emr.Client) (interface{}, error)) (interface{}, error) {
	goSdkMutex.Lock()
	defer goSdkMutex.Unlock()

	if client.emrconn == nil {
		emrConn, err := emr.NewClientWithOptions(client.RegionId, client.getSdkConfig(), client.config.getAuthCredential(true))
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the E-MapReduce client: %#v", err)
		}
		emrConn.AppendUserAgent(Terraform, terraformVersion)
		emrConn.AppendUserAgent(Provider, providerVersion)
		if client.config.ConfigurationSource != "" {
			emrConn.AppendUserAgent(Module, client.config.ConfigurationSource)
		}
		client.emrconn = emrConn
	}

	return do(client.emrconn)
}
*/
