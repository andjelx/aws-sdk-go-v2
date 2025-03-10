// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The request to sign out of all devices, as an administrator.
type AdminUserGlobalSignOutInput struct {
	_ struct{} `type:"structure"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user name.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s AdminUserGlobalSignOutInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminUserGlobalSignOutInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminUserGlobalSignOutInput"}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The global sign-out response, as an administrator.
type AdminUserGlobalSignOutOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AdminUserGlobalSignOutOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminUserGlobalSignOut = "AdminUserGlobalSignOut"

// AdminUserGlobalSignOutRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Signs out users from all devices, as an administrator.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using AdminUserGlobalSignOutRequest.
//    req := client.AdminUserGlobalSignOutRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminUserGlobalSignOut
func (c *Client) AdminUserGlobalSignOutRequest(input *AdminUserGlobalSignOutInput) AdminUserGlobalSignOutRequest {
	op := &aws.Operation{
		Name:       opAdminUserGlobalSignOut,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AdminUserGlobalSignOutInput{}
	}

	req := c.newRequest(op, input, &AdminUserGlobalSignOutOutput{})
	return AdminUserGlobalSignOutRequest{Request: req, Input: input, Copy: c.AdminUserGlobalSignOutRequest}
}

// AdminUserGlobalSignOutRequest is the request type for the
// AdminUserGlobalSignOut API operation.
type AdminUserGlobalSignOutRequest struct {
	*aws.Request
	Input *AdminUserGlobalSignOutInput
	Copy  func(*AdminUserGlobalSignOutInput) AdminUserGlobalSignOutRequest
}

// Send marshals and sends the AdminUserGlobalSignOut API request.
func (r AdminUserGlobalSignOutRequest) Send(ctx context.Context) (*AdminUserGlobalSignOutResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminUserGlobalSignOutResponse{
		AdminUserGlobalSignOutOutput: r.Request.Data.(*AdminUserGlobalSignOutOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminUserGlobalSignOutResponse is the response type for the
// AdminUserGlobalSignOut API operation.
type AdminUserGlobalSignOutResponse struct {
	*AdminUserGlobalSignOutOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminUserGlobalSignOut request.
func (r *AdminUserGlobalSignOutResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
