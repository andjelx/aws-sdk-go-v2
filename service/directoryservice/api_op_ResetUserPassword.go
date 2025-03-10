// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResetUserPasswordInput struct {
	_ struct{} `type:"structure"`

	// Identifier of the AWS Managed Microsoft AD or Simple AD directory in which
	// the user resides.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The new password that will be reset.
	//
	// NewPassword is a required field
	NewPassword *string `min:"1" type:"string" required:"true" sensitive:"true"`

	// The user name of the user whose password will be reset.
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ResetUserPasswordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetUserPasswordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetUserPasswordInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if s.NewPassword == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewPassword"))
	}
	if s.NewPassword != nil && len(*s.NewPassword) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NewPassword", 1))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResetUserPasswordOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ResetUserPasswordOutput) String() string {
	return awsutil.Prettify(s)
}

const opResetUserPassword = "ResetUserPassword"

// ResetUserPasswordRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Resets the password for any user in your AWS Managed Microsoft AD or Simple
// AD directory.
//
//    // Example sending a request using ResetUserPasswordRequest.
//    req := client.ResetUserPasswordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/ResetUserPassword
func (c *Client) ResetUserPasswordRequest(input *ResetUserPasswordInput) ResetUserPasswordRequest {
	op := &aws.Operation{
		Name:       opResetUserPassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetUserPasswordInput{}
	}

	req := c.newRequest(op, input, &ResetUserPasswordOutput{})
	return ResetUserPasswordRequest{Request: req, Input: input, Copy: c.ResetUserPasswordRequest}
}

// ResetUserPasswordRequest is the request type for the
// ResetUserPassword API operation.
type ResetUserPasswordRequest struct {
	*aws.Request
	Input *ResetUserPasswordInput
	Copy  func(*ResetUserPasswordInput) ResetUserPasswordRequest
}

// Send marshals and sends the ResetUserPassword API request.
func (r ResetUserPasswordRequest) Send(ctx context.Context) (*ResetUserPasswordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetUserPasswordResponse{
		ResetUserPasswordOutput: r.Request.Data.(*ResetUserPasswordOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetUserPasswordResponse is the response type for the
// ResetUserPassword API operation.
type ResetUserPasswordResponse struct {
	*ResetUserPasswordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetUserPassword request.
func (r *ResetUserPasswordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
