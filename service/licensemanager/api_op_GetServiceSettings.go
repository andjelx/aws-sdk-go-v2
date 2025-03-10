// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetServiceSettingsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetServiceSettingsInput) String() string {
	return awsutil.Prettify(s)
}

type GetServiceSettingsOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether cross-account discovery has been enabled.
	EnableCrossAccountsDiscovery *bool `type:"boolean"`

	// Indicates whether AWS Organizations has been integrated with License Manager
	// for cross-account discovery.
	OrganizationConfiguration *OrganizationConfiguration `type:"structure"`

	// Regional S3 bucket path for storing reports, license trail event data, discovery
	// data, etc.
	S3BucketArn *string `type:"string"`

	// SNS topic configured to receive notifications from License Manager.
	SnsTopicArn *string `type:"string"`
}

// String returns the string representation
func (s GetServiceSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetServiceSettings = "GetServiceSettings"

// GetServiceSettingsRequest returns a request value for making API operation for
// AWS License Manager.
//
// Gets License Manager settings for a region. Exposes the configured S3 bucket,
// SNS topic, etc., for inspection.
//
//    // Example sending a request using GetServiceSettingsRequest.
//    req := client.GetServiceSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/GetServiceSettings
func (c *Client) GetServiceSettingsRequest(input *GetServiceSettingsInput) GetServiceSettingsRequest {
	op := &aws.Operation{
		Name:       opGetServiceSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetServiceSettingsInput{}
	}

	req := c.newRequest(op, input, &GetServiceSettingsOutput{})
	return GetServiceSettingsRequest{Request: req, Input: input, Copy: c.GetServiceSettingsRequest}
}

// GetServiceSettingsRequest is the request type for the
// GetServiceSettings API operation.
type GetServiceSettingsRequest struct {
	*aws.Request
	Input *GetServiceSettingsInput
	Copy  func(*GetServiceSettingsInput) GetServiceSettingsRequest
}

// Send marshals and sends the GetServiceSettings API request.
func (r GetServiceSettingsRequest) Send(ctx context.Context) (*GetServiceSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetServiceSettingsResponse{
		GetServiceSettingsOutput: r.Request.Data.(*GetServiceSettingsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetServiceSettingsResponse is the response type for the
// GetServiceSettings API operation.
type GetServiceSettingsResponse struct {
	*GetServiceSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetServiceSettings request.
func (r *GetServiceSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
