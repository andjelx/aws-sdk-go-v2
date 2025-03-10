// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListAssociationsForLicenseConfigurationInput struct {
	_ struct{} `type:"structure"`

	// ARN of a LicenseConfiguration object.
	//
	// LicenseConfigurationArn is a required field
	LicenseConfigurationArn *string `type:"string" required:"true"`

	// Maximum number of results to return in a single call. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int64 `type:"integer"`

	// Token for the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAssociationsForLicenseConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAssociationsForLicenseConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAssociationsForLicenseConfigurationInput"}

	if s.LicenseConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LicenseConfigurationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListAssociationsForLicenseConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Lists association objects for the license configuration, each containing
	// the association time, number of consumed licenses, resource ARN, resource
	// ID, account ID that owns the resource, resource size, and resource type.
	LicenseConfigurationAssociations []LicenseConfigurationAssociation `type:"list"`

	// Token for the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAssociationsForLicenseConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAssociationsForLicenseConfiguration = "ListAssociationsForLicenseConfiguration"

// ListAssociationsForLicenseConfigurationRequest returns a request value for making API operation for
// AWS License Manager.
//
// Lists the resource associations for a license configuration. Resource associations
// need not consume licenses from a license configuration. For example, an AMI
// or a stopped instance may not consume a license (depending on the license
// rules). Use this operation to find all resources associated with a license
// configuration.
//
//    // Example sending a request using ListAssociationsForLicenseConfigurationRequest.
//    req := client.ListAssociationsForLicenseConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/ListAssociationsForLicenseConfiguration
func (c *Client) ListAssociationsForLicenseConfigurationRequest(input *ListAssociationsForLicenseConfigurationInput) ListAssociationsForLicenseConfigurationRequest {
	op := &aws.Operation{
		Name:       opListAssociationsForLicenseConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAssociationsForLicenseConfigurationInput{}
	}

	req := c.newRequest(op, input, &ListAssociationsForLicenseConfigurationOutput{})
	return ListAssociationsForLicenseConfigurationRequest{Request: req, Input: input, Copy: c.ListAssociationsForLicenseConfigurationRequest}
}

// ListAssociationsForLicenseConfigurationRequest is the request type for the
// ListAssociationsForLicenseConfiguration API operation.
type ListAssociationsForLicenseConfigurationRequest struct {
	*aws.Request
	Input *ListAssociationsForLicenseConfigurationInput
	Copy  func(*ListAssociationsForLicenseConfigurationInput) ListAssociationsForLicenseConfigurationRequest
}

// Send marshals and sends the ListAssociationsForLicenseConfiguration API request.
func (r ListAssociationsForLicenseConfigurationRequest) Send(ctx context.Context) (*ListAssociationsForLicenseConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAssociationsForLicenseConfigurationResponse{
		ListAssociationsForLicenseConfigurationOutput: r.Request.Data.(*ListAssociationsForLicenseConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAssociationsForLicenseConfigurationResponse is the response type for the
// ListAssociationsForLicenseConfiguration API operation.
type ListAssociationsForLicenseConfigurationResponse struct {
	*ListAssociationsForLicenseConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAssociationsForLicenseConfiguration request.
func (r *ListAssociationsForLicenseConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
