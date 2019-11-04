package kinesis

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
)

var readDuration = 5 * time.Second

func init() {
<<<<<<< HEAD
<<<<<<< HEAD
	initRequest = customizeRequest
}

func customizeRequest(r *request.Request) {
	if r.Operation.Name == opGetRecords {
		r.ApplyOptions(request.WithResponseReadTimeout(readDuration))
	}

	// Service specific error codes. Github(aws/aws-sdk-go#1376)
	r.RetryErrorCodes = append(r.RetryErrorCodes, ErrCodeLimitExceededException)
=======
=======
>>>>>>> 25fea6fedf7cf6c194bd2d8d3983d3609770c685
	ops := []string{
		opGetRecords,
	}
	initRequest = func(r *request.Request) {
		for _, operation := range ops {
			if r.Operation.Name == operation {
				r.ApplyOptions(request.WithResponseReadTimeout(readDuration))
			}
		}
	}
<<<<<<< HEAD
>>>>>>> Some more fixes
=======
>>>>>>> 25fea6fedf7cf6c194bd2d8d3983d3609770c685
}
