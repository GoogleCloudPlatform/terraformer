package rds

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	request "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

func init() {
	ops := []string{
		opCopyDBSnapshot,
		opCreateDBInstanceReadReplica,
		opCopyDBClusterSnapshot,
		opCreateDBCluster,
	}
	initRequest = func(c *Client, r *request.Request) {
		for _, operation := range ops {
			if r.Operation.Name == operation {
				r.Handlers.Build.PushFront(fillPresignedURL)
			}
		}
	}
}

func fillPresignedURL(r *request.Request) {
	fns := map[string]func(r *request.Request){
		opCopyDBSnapshot:              copyDBSnapshotPresign,
		opCreateDBInstanceReadReplica: createDBInstanceReadReplicaPresign,
		opCopyDBClusterSnapshot:       copyDBClusterSnapshotPresign,
		opCreateDBCluster:             createDBClusterPresign,
	}
	if !r.ParamsFilled() {
		return
	}
	if f, ok := fns[r.Operation.Name]; ok {
		f(r)
	}
}

func copyDBSnapshotPresign(r *request.Request) {
	originParams := r.Params.(*CopyDBSnapshotInput)

	if originParams.SourceRegion == nil || originParams.PreSignedUrl != nil || originParams.DestinationRegion != nil {
		return
	}

	originParams.DestinationRegion = aws.String(r.Config.Region)
	// preSignedUrl is not required for instances in the same region.
	if *originParams.SourceRegion == *originParams.DestinationRegion {
		return
	}
	newParams := awsutil.CopyOf(r.Params).(*CopyDBSnapshotInput)
	originParams.PreSignedUrl = presignURL(r, originParams.SourceRegion, newParams)
}

func createDBInstanceReadReplicaPresign(r *request.Request) {
	originParams := r.Params.(*CreateDBInstanceReadReplicaInput)

	if originParams.SourceRegion == nil || originParams.PreSignedUrl != nil || originParams.DestinationRegion != nil {
		return
	}

	originParams.DestinationRegion = aws.String(r.Config.Region)
	// preSignedUrl is not required for instances in the same region.
	if *originParams.SourceRegion == *originParams.DestinationRegion {
		return
	}
	newParams := awsutil.CopyOf(r.Params).(*CreateDBInstanceReadReplicaInput)
	originParams.PreSignedUrl = presignURL(r, originParams.SourceRegion, newParams)
}

func copyDBClusterSnapshotPresign(r *request.Request) {
	originParams := r.Params.(*CopyDBClusterSnapshotInput)

	if originParams.SourceRegion == nil || originParams.PreSignedUrl != nil || originParams.DestinationRegion != nil {
		return
	}

	originParams.DestinationRegion = aws.String(r.Config.Region)
	// preSignedUrl is not required for instances in the same region.
	if *originParams.SourceRegion == *originParams.DestinationRegion {
		return
	}
	newParams := awsutil.CopyOf(r.Params).(*CopyDBClusterSnapshotInput)
	originParams.PreSignedUrl = presignURL(r, originParams.SourceRegion, newParams)
}

func createDBClusterPresign(r *request.Request) {
	originParams := r.Params.(*CreateDBClusterInput)

	if originParams.SourceRegion == nil || originParams.PreSignedUrl != nil || originParams.DestinationRegion != nil {
		return
	}

	originParams.DestinationRegion = aws.String(r.Config.Region)
	// preSignedUrl is not required for instances in the same region.
	if *originParams.SourceRegion == *originParams.DestinationRegion {
		return
	}
	newParams := awsutil.CopyOf(r.Params).(*CreateDBClusterInput)
	originParams.PreSignedUrl = presignURL(r, originParams.SourceRegion, newParams)
}

// presignURL will presign the request by using SoureRegion to sign with. SourceRegion is not
// sent to the service, and is only used to not have the SDKs parsing ARNs.
func presignURL(r *request.Request, sourceRegion *string, newParams interface{}) *string {
	cfgCp := r.Config.Copy()
	cfgCp.EndpointResolver = nil
	cfgCp.Region = aws.StringValue(sourceRegion)

	metadata := r.Metadata
	resolved, err := r.Config.EndpointResolver.ResolveEndpoint(metadata.EndpointsID, cfgCp.Region)
	if err != nil {
		r.Error = err
		return nil
	}

	cfgCp.EndpointResolver = aws.ResolveWithEndpoint(resolved)

	// Presign a request with modified params
	req := request.New(cfgCp, metadata, r.Handlers, r.Retryer, r.Operation, newParams, r.Data)
	req.Operation.HTTPMethod = "GET"
	uri, err := req.Presign(5 * time.Minute) // 5 minutes should be enough.
	if err != nil {                          // bubble error back up to original request
		r.Error = err
		return nil
	}

	// We have our URL, set it on params
	return &uri
}
