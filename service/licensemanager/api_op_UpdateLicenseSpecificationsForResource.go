// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateLicenseSpecificationsForResourceInput struct {
	_ struct{} `type:"structure"`

	// License configuration ARNs to be added to a resource.
	AddLicenseSpecifications []LicenseSpecification `type:"list"`

	// License configuration ARNs to be removed from a resource.
	RemoveLicenseSpecifications []LicenseSpecification `type:"list"`

	// ARN for an AWS server resource.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateLicenseSpecificationsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateLicenseSpecificationsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateLicenseSpecificationsForResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.AddLicenseSpecifications != nil {
		for i, v := range s.AddLicenseSpecifications {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AddLicenseSpecifications", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RemoveLicenseSpecifications != nil {
		for i, v := range s.RemoveLicenseSpecifications {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RemoveLicenseSpecifications", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateLicenseSpecificationsForResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateLicenseSpecificationsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateLicenseSpecificationsForResource = "UpdateLicenseSpecificationsForResource"

// UpdateLicenseSpecificationsForResourceRequest returns a request value for making API operation for
// AWS License Manager.
//
// Adds or removes license configurations for a specified AWS resource. This
// operation currently supports updating the license specifications of AMIs,
// instances, and hosts. Launch templates and AWS CloudFormation templates are
// not managed from this operation as those resources send the license configurations
// directly to a resource creation operation, such as RunInstances.
//
//    // Example sending a request using UpdateLicenseSpecificationsForResourceRequest.
//    req := client.UpdateLicenseSpecificationsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/UpdateLicenseSpecificationsForResource
func (c *Client) UpdateLicenseSpecificationsForResourceRequest(input *UpdateLicenseSpecificationsForResourceInput) UpdateLicenseSpecificationsForResourceRequest {
	op := &aws.Operation{
		Name:       opUpdateLicenseSpecificationsForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateLicenseSpecificationsForResourceInput{}
	}

	req := c.newRequest(op, input, &UpdateLicenseSpecificationsForResourceOutput{})
	return UpdateLicenseSpecificationsForResourceRequest{Request: req, Input: input, Copy: c.UpdateLicenseSpecificationsForResourceRequest}
}

// UpdateLicenseSpecificationsForResourceRequest is the request type for the
// UpdateLicenseSpecificationsForResource API operation.
type UpdateLicenseSpecificationsForResourceRequest struct {
	*aws.Request
	Input *UpdateLicenseSpecificationsForResourceInput
	Copy  func(*UpdateLicenseSpecificationsForResourceInput) UpdateLicenseSpecificationsForResourceRequest
}

// Send marshals and sends the UpdateLicenseSpecificationsForResource API request.
func (r UpdateLicenseSpecificationsForResourceRequest) Send(ctx context.Context) (*UpdateLicenseSpecificationsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateLicenseSpecificationsForResourceResponse{
		UpdateLicenseSpecificationsForResourceOutput: r.Request.Data.(*UpdateLicenseSpecificationsForResourceOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateLicenseSpecificationsForResourceResponse is the response type for the
// UpdateLicenseSpecificationsForResource API operation.
type UpdateLicenseSpecificationsForResourceResponse struct {
	*UpdateLicenseSpecificationsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateLicenseSpecificationsForResource request.
func (r *UpdateLicenseSpecificationsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
