// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisablePolicyTypeInput struct {
	_ struct{} `type:"structure"`

	// The policy type that you want to disable in this root.
	//
	// PolicyType is a required field
	PolicyType PolicyType `type:"string" required:"true" enum:"true"`

	// The unique identifier (ID) of the root in which you want to disable a policy
	// type. You can get the ID from the ListRoots operation.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a root ID string
	// requires "r-" followed by from 4 to 32 lower-case letters or digits.
	//
	// RootId is a required field
	RootId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisablePolicyTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisablePolicyTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisablePolicyTypeInput"}
	if len(s.PolicyType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PolicyType"))
	}

	if s.RootId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RootId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisablePolicyTypeOutput struct {
	_ struct{} `type:"structure"`

	// A structure that shows the root with the updated list of enabled policy types.
	Root *Root `type:"structure"`
}

// String returns the string representation
func (s DisablePolicyTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisablePolicyType = "DisablePolicyType"

// DisablePolicyTypeRequest returns a request value for making API operation for
// AWS Organizations.
//
// Disables an organizational control policy type in a root. A policy of a certain
// type can be attached to entities in a root only if that type is enabled in
// the root. After you perform this operation, you no longer can attach policies
// of the specified type to that root or to any organizational unit (OU) or
// account in that root. You can undo this by using the EnablePolicyType operation.
//
// This is an asynchronous request that AWS performs in the background. If you
// disable a policy for a root, it still appears enabled for the organization
// if all features (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
// are enabled for the organization. AWS recommends that you first use ListRoots
// to see the status of policy types for a specified root, and then use this
// operation.
//
// This operation can be called only from the organization's master account.
//
// To view the status of available policy types in the organization, use DescribeOrganization.
//
//    // Example sending a request using DisablePolicyTypeRequest.
//    req := client.DisablePolicyTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DisablePolicyType
func (c *Client) DisablePolicyTypeRequest(input *DisablePolicyTypeInput) DisablePolicyTypeRequest {
	op := &aws.Operation{
		Name:       opDisablePolicyType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisablePolicyTypeInput{}
	}

	req := c.newRequest(op, input, &DisablePolicyTypeOutput{})
	return DisablePolicyTypeRequest{Request: req, Input: input, Copy: c.DisablePolicyTypeRequest}
}

// DisablePolicyTypeRequest is the request type for the
// DisablePolicyType API operation.
type DisablePolicyTypeRequest struct {
	*aws.Request
	Input *DisablePolicyTypeInput
	Copy  func(*DisablePolicyTypeInput) DisablePolicyTypeRequest
}

// Send marshals and sends the DisablePolicyType API request.
func (r DisablePolicyTypeRequest) Send(ctx context.Context) (*DisablePolicyTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisablePolicyTypeResponse{
		DisablePolicyTypeOutput: r.Request.Data.(*DisablePolicyTypeOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisablePolicyTypeResponse is the response type for the
// DisablePolicyType API operation.
type DisablePolicyTypeResponse struct {
	*DisablePolicyTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisablePolicyType request.
func (r *DisablePolicyTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
