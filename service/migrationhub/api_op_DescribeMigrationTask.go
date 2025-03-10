// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeMigrationTaskInput struct {
	_ struct{} `type:"structure"`

	// The identifier given to the MigrationTask.
	//
	// MigrationTaskName is a required field
	MigrationTaskName *string `min:"1" type:"string" required:"true"`

	// The name of the ProgressUpdateStream.
	//
	// ProgressUpdateStream is a required field
	ProgressUpdateStream *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeMigrationTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMigrationTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMigrationTaskInput"}

	if s.MigrationTaskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MigrationTaskName"))
	}
	if s.MigrationTaskName != nil && len(*s.MigrationTaskName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MigrationTaskName", 1))
	}

	if s.ProgressUpdateStream == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProgressUpdateStream"))
	}
	if s.ProgressUpdateStream != nil && len(*s.ProgressUpdateStream) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProgressUpdateStream", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeMigrationTaskOutput struct {
	_ struct{} `type:"structure"`

	// Object encapsulating information about the migration task.
	MigrationTask *MigrationTask `type:"structure"`
}

// String returns the string representation
func (s DescribeMigrationTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMigrationTask = "DescribeMigrationTask"

// DescribeMigrationTaskRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Retrieves a list of all attributes associated with a specific migration task.
//
//    // Example sending a request using DescribeMigrationTaskRequest.
//    req := client.DescribeMigrationTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/DescribeMigrationTask
func (c *Client) DescribeMigrationTaskRequest(input *DescribeMigrationTaskInput) DescribeMigrationTaskRequest {
	op := &aws.Operation{
		Name:       opDescribeMigrationTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMigrationTaskInput{}
	}

	req := c.newRequest(op, input, &DescribeMigrationTaskOutput{})
	return DescribeMigrationTaskRequest{Request: req, Input: input, Copy: c.DescribeMigrationTaskRequest}
}

// DescribeMigrationTaskRequest is the request type for the
// DescribeMigrationTask API operation.
type DescribeMigrationTaskRequest struct {
	*aws.Request
	Input *DescribeMigrationTaskInput
	Copy  func(*DescribeMigrationTaskInput) DescribeMigrationTaskRequest
}

// Send marshals and sends the DescribeMigrationTask API request.
func (r DescribeMigrationTaskRequest) Send(ctx context.Context) (*DescribeMigrationTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMigrationTaskResponse{
		DescribeMigrationTaskOutput: r.Request.Data.(*DescribeMigrationTaskOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMigrationTaskResponse is the response type for the
// DescribeMigrationTask API operation.
type DescribeMigrationTaskResponse struct {
	*DescribeMigrationTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMigrationTask request.
func (r *DescribeMigrationTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
