// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchGetBuildsInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the builds.
	//
	// Ids is a required field
	Ids []string `locationName:"ids" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetBuildsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetBuildsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetBuildsInput"}

	if s.Ids == nil {
		invalidParams.Add(aws.NewErrParamRequired("Ids"))
	}
	if s.Ids != nil && len(s.Ids) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Ids", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchGetBuildsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the requested builds.
	Builds []Build `locationName:"builds" type:"list"`

	// The IDs of builds for which information could not be found.
	BuildsNotFound []string `locationName:"buildsNotFound" min:"1" type:"list"`
}

// String returns the string representation
func (s BatchGetBuildsOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetBuilds = "BatchGetBuilds"

// BatchGetBuildsRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Gets information about builds.
//
//    // Example sending a request using BatchGetBuildsRequest.
//    req := client.BatchGetBuildsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/BatchGetBuilds
func (c *Client) BatchGetBuildsRequest(input *BatchGetBuildsInput) BatchGetBuildsRequest {
	op := &aws.Operation{
		Name:       opBatchGetBuilds,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetBuildsInput{}
	}

	req := c.newRequest(op, input, &BatchGetBuildsOutput{})
	return BatchGetBuildsRequest{Request: req, Input: input, Copy: c.BatchGetBuildsRequest}
}

// BatchGetBuildsRequest is the request type for the
// BatchGetBuilds API operation.
type BatchGetBuildsRequest struct {
	*aws.Request
	Input *BatchGetBuildsInput
	Copy  func(*BatchGetBuildsInput) BatchGetBuildsRequest
}

// Send marshals and sends the BatchGetBuilds API request.
func (r BatchGetBuildsRequest) Send(ctx context.Context) (*BatchGetBuildsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetBuildsResponse{
		BatchGetBuildsOutput: r.Request.Data.(*BatchGetBuildsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetBuildsResponse is the response type for the
// BatchGetBuilds API operation.
type BatchGetBuildsResponse struct {
	*BatchGetBuildsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetBuilds request.
func (r *BatchGetBuildsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
