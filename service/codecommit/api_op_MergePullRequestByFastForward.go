// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type MergePullRequestByFastForwardInput struct {
	_ struct{} `type:"structure"`

	// The system-generated ID of the pull request. To get this ID, use ListPullRequests.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The name of the repository where the pull request was created.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`

	// The full commit ID of the original or updated commit in the pull request
	// source branch. Pass this value if you want an exception thrown if the current
	// commit ID of the tip of the source branch does not match this commit ID.
	SourceCommitId *string `locationName:"sourceCommitId" type:"string"`
}

// String returns the string representation
func (s MergePullRequestByFastForwardInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MergePullRequestByFastForwardInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MergePullRequestByFastForwardInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
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

type MergePullRequestByFastForwardOutput struct {
	_ struct{} `type:"structure"`

	// Information about the specified pull request, including information about
	// the merge.
	PullRequest *PullRequest `locationName:"pullRequest" type:"structure"`
}

// String returns the string representation
func (s MergePullRequestByFastForwardOutput) String() string {
	return awsutil.Prettify(s)
}

const opMergePullRequestByFastForward = "MergePullRequestByFastForward"

// MergePullRequestByFastForwardRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Attempts to merge the source commit of a pull request into the specified
// destination branch for that pull request at the specified commit using the
// fast-forward merge strategy. If the merge is successful, it closes the pull
// request.
//
//    // Example sending a request using MergePullRequestByFastForwardRequest.
//    req := client.MergePullRequestByFastForwardRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/MergePullRequestByFastForward
func (c *Client) MergePullRequestByFastForwardRequest(input *MergePullRequestByFastForwardInput) MergePullRequestByFastForwardRequest {
	op := &aws.Operation{
		Name:       opMergePullRequestByFastForward,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MergePullRequestByFastForwardInput{}
	}

	req := c.newRequest(op, input, &MergePullRequestByFastForwardOutput{})
	return MergePullRequestByFastForwardRequest{Request: req, Input: input, Copy: c.MergePullRequestByFastForwardRequest}
}

// MergePullRequestByFastForwardRequest is the request type for the
// MergePullRequestByFastForward API operation.
type MergePullRequestByFastForwardRequest struct {
	*aws.Request
	Input *MergePullRequestByFastForwardInput
	Copy  func(*MergePullRequestByFastForwardInput) MergePullRequestByFastForwardRequest
}

// Send marshals and sends the MergePullRequestByFastForward API request.
func (r MergePullRequestByFastForwardRequest) Send(ctx context.Context) (*MergePullRequestByFastForwardResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &MergePullRequestByFastForwardResponse{
		MergePullRequestByFastForwardOutput: r.Request.Data.(*MergePullRequestByFastForwardOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// MergePullRequestByFastForwardResponse is the response type for the
// MergePullRequestByFastForward API operation.
type MergePullRequestByFastForwardResponse struct {
	*MergePullRequestByFastForwardOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// MergePullRequestByFastForward request.
func (r *MergePullRequestByFastForwardResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
