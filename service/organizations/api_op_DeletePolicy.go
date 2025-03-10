// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeletePolicyInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the policy that you want to delete. You can
	// get the ID from the ListPolicies or ListPoliciesForTarget operations.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a policy ID string
	// requires "p-" followed by from 8 to 128 lower-case letters or digits.
	//
	// PolicyId is a required field
	PolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePolicyInput"}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeletePolicy = "DeletePolicy"

// DeletePolicyRequest returns a request value for making API operation for
// AWS Organizations.
//
// Deletes the specified policy from your organization. Before you perform this
// operation, you must first detach the policy from all organizational units
// (OUs), roots, and accounts.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using DeletePolicyRequest.
//    req := client.DeletePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DeletePolicy
func (c *Client) DeletePolicyRequest(input *DeletePolicyInput) DeletePolicyRequest {
	op := &aws.Operation{
		Name:       opDeletePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePolicyInput{}
	}

	req := c.newRequest(op, input, &DeletePolicyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePolicyRequest{Request: req, Input: input, Copy: c.DeletePolicyRequest}
}

// DeletePolicyRequest is the request type for the
// DeletePolicy API operation.
type DeletePolicyRequest struct {
	*aws.Request
	Input *DeletePolicyInput
	Copy  func(*DeletePolicyInput) DeletePolicyRequest
}

// Send marshals and sends the DeletePolicy API request.
func (r DeletePolicyRequest) Send(ctx context.Context) (*DeletePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePolicyResponse{
		DeletePolicyOutput: r.Request.Data.(*DeletePolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePolicyResponse is the response type for the
// DeletePolicy API operation.
type DeletePolicyResponse struct {
	*DeletePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePolicy request.
func (r *DeletePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
