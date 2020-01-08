package s3

import (
	"io/ioutil"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var reBucketLocation = regexp.MustCompile(`>([^<>]+)<\/Location`)

// NormalizeBucketLocation is a utility function which will update the
// passed in value to always be a region ID. Generally this would be used
// with GetBucketLocation API operation.
//
// Replaces empty string with "us-east-1", and "EU" with "eu-west-1".
//
// See http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETlocation.html
// for more information on the values that can be returned.
func NormalizeBucketLocation(loc BucketLocationConstraint) BucketLocationConstraint {
	switch loc {
	case "":
		loc = "us-east-1"
	case "EU":
		loc = "eu-west-1"
	}

	return loc
}

// NormalizeBucketLocationHandler is a request handler which will update the
// GetBucketLocation's result LocationConstraint value to always be a region ID.
//
// Replaces empty string with "us-east-1", and "EU" with "eu-west-1".
//
// See http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETlocation.html
// for more information on the values that can be returned.
//
//     req, result := svc.GetBucketLocationRequest(&s3.GetBucketLocationInput{
//         Bucket: aws.String(bucket),
//     })
//     req.Handlers.Unmarshal.PushBackNamed(NormalizeBucketLocationHandler)
//     err := req.Send()
var NormalizeBucketLocationHandler = aws.NamedHandler{
	Name: "awssdk.s3.NormalizeBucketLocation",
	Fn: func(req *aws.Request) {
		if req.Error != nil {
			return
		}

		out := req.Data.(*GetBucketLocationOutput)
		loc := NormalizeBucketLocation(out.LocationConstraint)
		out.LocationConstraint = loc
	},
}

// WithNormalizeBucketLocation is a request option which will update the
// GetBucketLocation's result LocationConstraint value to always be a region ID.
//
// Replaces empty string with "us-east-1", and "EU" with "eu-west-1".
//
// See http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETlocation.html
// for more information on the values that can be returned.
//
//     result, err := svc.GetBucketLocationWithContext(ctx,
//         &s3.GetBucketLocationInput{
//             Bucket: aws.String(bucket),
//         },
//         s3.WithNormalizeBucketLocation,
//     )
func WithNormalizeBucketLocation(r *aws.Request) {
	r.Handlers.Unmarshal.PushBackNamed(NormalizeBucketLocationHandler)
}

func buildGetBucketLocation(r *aws.Request) {
	out := r.Data.(*GetBucketLocationOutput)
	b, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		r.Error = awserr.New("SerializationError", "failed reading response body", err)
		return
	}

	match := reBucketLocation.FindSubmatch(b)
	if len(match) > 1 {
		loc := BucketLocationConstraint(match[1])
		out.LocationConstraint = loc
	}
}

func populateLocationConstraint(r *aws.Request) {
	if r.ParamsFilled() && r.Config.Region != "us-east-1" {
		in := r.Params.(*CreateBucketInput)
		if in.CreateBucketConfiguration == nil {
			r.Params = awsutil.CopyOf(r.Params)
			in = r.Params.(*CreateBucketInput)
			in.CreateBucketConfiguration = &CreateBucketConfiguration{
				LocationConstraint: BucketLocationConstraint(r.Config.Region),
			}
		}
	}
}
