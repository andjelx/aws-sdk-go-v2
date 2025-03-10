// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeComponentInput struct {
	_ struct{} `type:"structure"`

	// The name of the component.
	//
	// ComponentName is a required field
	ComponentName *string `type:"string" required:"true"`

	// The name of the resource group.
	//
	// ResourceGroupName is a required field
	ResourceGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeComponentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeComponentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeComponentInput"}

	if s.ComponentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ComponentName"))
	}

	if s.ResourceGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeComponentOutput struct {
	_ struct{} `type:"structure"`

	// Describes a standalone resource or similarly grouped resources that the application
	// is made up of.
	ApplicationComponent *ApplicationComponent `type:"structure"`

	// The list of resource ARNs that belong to the component.
	ResourceList []string `type:"list"`
}

// String returns the string representation
func (s DescribeComponentOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeComponent = "DescribeComponent"

// DescribeComponentRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Describes a component and lists the resources that are grouped together in
// a component.
//
//    // Example sending a request using DescribeComponentRequest.
//    req := client.DescribeComponentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/DescribeComponent
func (c *Client) DescribeComponentRequest(input *DescribeComponentInput) DescribeComponentRequest {
	op := &aws.Operation{
		Name:       opDescribeComponent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeComponentInput{}
	}

	req := c.newRequest(op, input, &DescribeComponentOutput{})
	return DescribeComponentRequest{Request: req, Input: input, Copy: c.DescribeComponentRequest}
}

// DescribeComponentRequest is the request type for the
// DescribeComponent API operation.
type DescribeComponentRequest struct {
	*aws.Request
	Input *DescribeComponentInput
	Copy  func(*DescribeComponentInput) DescribeComponentRequest
}

// Send marshals and sends the DescribeComponent API request.
func (r DescribeComponentRequest) Send(ctx context.Context) (*DescribeComponentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeComponentResponse{
		DescribeComponentOutput: r.Request.Data.(*DescribeComponentOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeComponentResponse is the response type for the
// DescribeComponent API operation.
type DescribeComponentResponse struct {
	*DescribeComponentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeComponent request.
func (r *DescribeComponentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
