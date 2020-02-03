// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFileInput struct {
	_ struct{} `type:"structure"`

	// The name of the branch where the commit that deletes the file is made.
	//
	// BranchName is a required field
	BranchName *string `locationName:"branchName" min:"1" type:"string" required:"true"`

	// The commit message you want to include as part of deleting the file. Commit
	// messages are limited to 256 KB. If no message is specified, a default message
	// is used.
	CommitMessage *string `locationName:"commitMessage" type:"string"`

	// The email address for the commit that deletes the file. If no email address
	// is specified, the email address is left blank.
	Email *string `locationName:"email" type:"string"`

	// The fully qualified path to the file that to be deleted, including the full
	// name and extension of that file. For example, /examples/file.md is a fully
	// qualified path to a file named file.md in a folder named examples.
	//
	// FilePath is a required field
	FilePath *string `locationName:"filePath" type:"string" required:"true"`

	// If a file is the only object in the folder or directory, specifies whether
	// to delete the folder or directory that contains the file. By default, empty
	// folders are deleted. This includes empty folders that are part of the directory
	// structure. For example, if the path to a file is dir1/dir2/dir3/dir4, and
	// dir2 and dir3 are empty, deleting the last file in dir4 also deletes the
	// empty folders dir4, dir3, and dir2.
	KeepEmptyFolders *bool `locationName:"keepEmptyFolders" type:"boolean"`

	// The name of the author of the commit that deletes the file. If no name is
	// specified, the user's ARN is used as the author name and committer name.
	Name *string `locationName:"name" type:"string"`

	// The ID of the commit that is the tip of the branch where you want to create
	// the commit that deletes the file. This must be the HEAD commit for the branch.
	// The commit that deletes the file is created from this commit ID.
	//
	// ParentCommitId is a required field
	ParentCommitId *string `locationName:"parentCommitId" type:"string" required:"true"`

	// The name of the repository that contains the file to delete.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFileInput"}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if s.FilePath == nil {
		invalidParams.Add(aws.NewErrParamRequired("FilePath"))
	}

	if s.ParentCommitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentCommitId"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFileOutput struct {
	_ struct{} `type:"structure"`

	// The blob ID removed from the tree as part of deleting the file.
	//
	// BlobId is a required field
	BlobId *string `locationName:"blobId" type:"string" required:"true"`

	// The full commit ID of the commit that contains the change that deletes the
	// file.
	//
	// CommitId is a required field
	CommitId *string `locationName:"commitId" type:"string" required:"true"`

	// The fully qualified path to the file to be deleted, including the full name
	// and extension of that file.
	//
	// FilePath is a required field
	FilePath *string `locationName:"filePath" type:"string" required:"true"`

	// The full SHA-1 pointer of the tree information for the commit that contains
	// the delete file change.
	//
	// TreeId is a required field
	TreeId *string `locationName:"treeId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFileOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteFile = "DeleteFile"

// DeleteFileRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Deletes a specified file from a specified branch. A commit is created on
// the branch that contains the revision. The file still exists in the commits
// earlier to the commit that contains the deletion.
//
//    // Example sending a request using DeleteFileRequest.
//    req := client.DeleteFileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DeleteFile
func (c *Client) DeleteFileRequest(input *DeleteFileInput) DeleteFileRequest {
	op := &aws.Operation{
		Name:       opDeleteFile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteFileInput{}
	}

	req := c.newRequest(op, input, &DeleteFileOutput{})
	return DeleteFileRequest{Request: req, Input: input, Copy: c.DeleteFileRequest}
}

// DeleteFileRequest is the request type for the
// DeleteFile API operation.
type DeleteFileRequest struct {
	*aws.Request
	Input *DeleteFileInput
	Copy  func(*DeleteFileInput) DeleteFileRequest
}

// Send marshals and sends the DeleteFile API request.
func (r DeleteFileRequest) Send(ctx context.Context) (*DeleteFileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFileResponse{
		DeleteFileOutput: r.Request.Data.(*DeleteFileOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFileResponse is the response type for the
// DeleteFile API operation.
type DeleteFileResponse struct {
	*DeleteFileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFile request.
func (r *DeleteFileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
