package s3

import (
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsarn "github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

// Used by shapes with members decorated as endpoint ARN.
func parseEndpointARN(v string) (arn.Resource, error) {
	return arn.ParseResource(v, accessPointResourceParser)
}

func accessPointResourceParser(a awsarn.ARN) (arn.Resource, error) {
	resParts := arn.SplitResource(a.Resource)
	switch resParts[0] {
	case "accesspoint":
		return arn.ParseAccessPointResource(a, resParts[1:])
	default:
		return nil, arn.InvalidARNError{ARN: a, Reason: "unknown resource type"}
	}
}

func buildEndpointHandler(c *Client) func(*aws.Request) {
	return func(req *aws.Request) {
		endpoint, ok := req.Params.(endpointARNGetter)
		if !ok || !endpoint.hasEndpointARN() {
			updateBucketEndpointFromParams(c, req)
			return
		}

		resource, err := endpoint.getEndpointARN()
		if err != nil {
			req.Error = newInvalidARNError(nil, err)
			return
		}

		resReq := resourceRequest{
			Resource:     resource,
			UseARNRegion: c.UseARNRegion,
			Request:      req,
		}

		if resReq.IsCrossPartition() {
			req.Error = newClientPartitionMismatchError(resource,
				req.Endpoint.PartitionID, req.Config.Region, nil)
			return
		}

		if !resReq.AllowCrossRegion() && resReq.IsCrossRegion() {
			req.Error = newClientRegionMismatchError(resource,
				req.Endpoint.PartitionID, req.Config.Region, nil)
			return
		}

		switch tv := resource.(type) {
		case arn.AccessPointARN:
			err = updateRequestAccessPointEndpoint(c, req, tv)
			if err != nil {
				req.Error = err
			}
		default:
			req.Error = newInvalidARNError(resource, nil)
		}
	}
}

type resourceRequest struct {
	Resource     arn.Resource
	UseARNRegion bool
	Request      *aws.Request
}

func (r resourceRequest) ARN() awsarn.ARN {
	return r.Resource.GetARN()
}

func (r resourceRequest) AllowCrossRegion() bool {
	return r.UseARNRegion
}

func (r resourceRequest) UseFIPS() bool {
	return isFIPS(r.Request.Config.Region)
}

func (r resourceRequest) IsCrossPartition() bool {
	return r.Request.Endpoint.PartitionID != r.Resource.GetARN().Partition
}

func (r resourceRequest) IsCrossRegion() bool {
	return isCrossRegion(r.Request, r.Resource.GetARN().Region)
}

func isFIPS(clientRegion string) bool {
	return strings.HasPrefix(clientRegion, "fips-") || strings.HasSuffix(clientRegion, "-fips")
}
func isCrossRegion(req *aws.Request, otherRegion string) bool {
	return req.Endpoint.SigningRegion != otherRegion
}

func updateBucketEndpointFromParams(c *Client, r *aws.Request) {
	bucket, ok := bucketNameFromReqParams(r.Params)
	if !ok {
		// Ignore operation requests if the bucket name was not provided
		// if this is an input validation error the validation handler
		// will report it.
		return
	}
	updateEndpointForS3Config(c, r, bucket)
}

func updateRequestAccessPointEndpoint(c *Client, req *aws.Request, accessPoint arn.AccessPointARN) error {
	// Accelerate not supported
	if c.UseAccelerate {
		return newClientConfiguredForAccelerateError(accessPoint,
			req.Endpoint.PartitionID, req.Config.Region, nil)
	}

	// Ignore the disable host prefix for access points since custom endpoints
	// are not supported.
	req.Config.DisableEndpointHostPrefix = false

	if err := accessPointEndpointBuilder(accessPoint).Build(c, req); err != nil {
		return err
	}

	removeBucketFromPath(req.HTTPRequest.URL)

	return nil
}

func removeBucketFromPath(u *url.URL) {
	u.Path = strings.Replace(u.Path, "/{Bucket}", "", -1)
	if u.Path == "" {
		u.Path = "/"
	}
}

type accessPointEndpointBuilder arn.AccessPointARN

const (
	accessPointPrefixLabel   = "accesspoint"
	accountIDPrefixLabel     = "accountID"
	accesPointPrefixTemplate = "{" + accessPointPrefixLabel + "}-{" + accountIDPrefixLabel + "}."
)

func (a accessPointEndpointBuilder) Build(c *Client, req *aws.Request) error {
	resolveRegion := arn.AccessPointARN(a).Region
	cfgRegion := req.Config.Region

	if isFIPS(cfgRegion) {
		if c.UseARNRegion && isCrossRegion(req, resolveRegion) {
			// FIPS with cross region is not supported, the SDK must fail
			// because there is no well defined method for SDK to construct a
			// correct FIPS endpoint.
			return newClientConfiguredForCrossRegionFIPSError(arn.AccessPointARN(a),
				req.Endpoint.PartitionID, cfgRegion, nil)
		}
		resolveRegion = cfgRegion
	}

	endpoint, err := resolveRegionalEndpoint(req, resolveRegion)
	if err != nil {
		return newFailedToResolveEndpointError(arn.AccessPointARN(a),
			req.Endpoint.PartitionID, cfgRegion, err)
	}

	req.SetEndpoint(endpoint)

	const serviceEndpointLabel = "s3-accesspoint"

	// dualstack provided by endpoint resolver
	cfgHost := req.HTTPRequest.URL.Host
	if strings.HasPrefix(cfgHost, "s3") {
		req.HTTPRequest.URL.Host = serviceEndpointLabel + cfgHost[2:]
	}

	protocol.HostPrefixBuilder{
		Prefix:   accesPointPrefixTemplate,
		LabelsFn: a.hostPrefixLabelValues,
	}.Build(req)

	err = protocol.ValidateEndpointHost(req.Operation.Name, req.HTTPRequest.URL.Host)
	if err != nil {
		return newInvalidARNError(arn.AccessPointARN(a), err)
	}

	return nil
}

func (a accessPointEndpointBuilder) hostPrefixLabelValues() map[string]string {
	return map[string]string{
		accessPointPrefixLabel: arn.AccessPointARN(a).AccessPointName,
		accountIDPrefixLabel:   arn.AccessPointARN(a).AccountID,
	}
}

func resolveRegionalEndpoint(r *aws.Request, region string) (aws.Endpoint, error) {
	return r.Config.EndpointResolver.ResolveEndpoint(EndpointsID, region)
}
