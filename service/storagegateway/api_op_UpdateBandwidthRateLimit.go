// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing one or more of the following fields:
//
//    * UpdateBandwidthRateLimitInput$AverageDownloadRateLimitInBitsPerSec
//
//    * UpdateBandwidthRateLimitInput$AverageUploadRateLimitInBitsPerSec
type UpdateBandwidthRateLimitInput struct {
	_ struct{} `type:"structure"`

	// The average download bandwidth rate limit in bits per second.
	AverageDownloadRateLimitInBitsPerSec *int64 `min:"102400" type:"long"`

	// The average upload bandwidth rate limit in bits per second.
	AverageUploadRateLimitInBitsPerSec *int64 `min:"51200" type:"long"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateBandwidthRateLimitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBandwidthRateLimitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBandwidthRateLimitInput"}
	if s.AverageDownloadRateLimitInBitsPerSec != nil && *s.AverageDownloadRateLimitInBitsPerSec < 102400 {
		invalidParams.Add(aws.NewErrParamMinValue("AverageDownloadRateLimitInBitsPerSec", 102400))
	}
	if s.AverageUploadRateLimitInBitsPerSec != nil && *s.AverageUploadRateLimitInBitsPerSec < 51200 {
		invalidParams.Add(aws.NewErrParamMinValue("AverageUploadRateLimitInBitsPerSec", 51200))
	}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the of the gateway whose throttle information was
// updated.
type UpdateBandwidthRateLimitOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s UpdateBandwidthRateLimitOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateBandwidthRateLimit = "UpdateBandwidthRateLimit"

// UpdateBandwidthRateLimitRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates the bandwidth rate limits of a gateway. You can update both the upload
// and download bandwidth rate limit or specify only one of the two. If you
// don't set a bandwidth rate limit, the existing rate limit remains.
//
// By default, a gateway's bandwidth rate limits are not set. If you don't set
// any limit, the gateway does not have any limitations on its bandwidth usage
// and could potentially use the maximum available bandwidth.
//
// To specify which gateway to update, use the Amazon Resource Name (ARN) of
// the gateway in your request.
//
//    // Example sending a request using UpdateBandwidthRateLimitRequest.
//    req := client.UpdateBandwidthRateLimitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateBandwidthRateLimit
func (c *Client) UpdateBandwidthRateLimitRequest(input *UpdateBandwidthRateLimitInput) UpdateBandwidthRateLimitRequest {
	op := &aws.Operation{
		Name:       opUpdateBandwidthRateLimit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateBandwidthRateLimitInput{}
	}

	req := c.newRequest(op, input, &UpdateBandwidthRateLimitOutput{})
	return UpdateBandwidthRateLimitRequest{Request: req, Input: input, Copy: c.UpdateBandwidthRateLimitRequest}
}

// UpdateBandwidthRateLimitRequest is the request type for the
// UpdateBandwidthRateLimit API operation.
type UpdateBandwidthRateLimitRequest struct {
	*aws.Request
	Input *UpdateBandwidthRateLimitInput
	Copy  func(*UpdateBandwidthRateLimitInput) UpdateBandwidthRateLimitRequest
}

// Send marshals and sends the UpdateBandwidthRateLimit API request.
func (r UpdateBandwidthRateLimitRequest) Send(ctx context.Context) (*UpdateBandwidthRateLimitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateBandwidthRateLimitResponse{
		UpdateBandwidthRateLimitOutput: r.Request.Data.(*UpdateBandwidthRateLimitOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateBandwidthRateLimitResponse is the response type for the
// UpdateBandwidthRateLimit API operation.
type UpdateBandwidthRateLimitResponse struct {
	*UpdateBandwidthRateLimitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateBandwidthRateLimit request.
func (r *UpdateBandwidthRateLimitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
