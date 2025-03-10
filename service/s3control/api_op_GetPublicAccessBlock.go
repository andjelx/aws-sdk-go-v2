// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3control

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetPublicAccessBlockInput struct {
	_ struct{} `type:"structure"`

	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPublicAccessBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPublicAccessBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPublicAccessBlockInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPublicAccessBlockInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-account-id", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetPublicAccessBlockOutput struct {
	_ struct{} `type:"structure" payload:"PublicAccessBlockConfiguration"`

	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetPublicAccessBlockOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPublicAccessBlockOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PublicAccessBlockConfiguration != nil {
		v := s.PublicAccessBlockConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "PublicAccessBlockConfiguration", v, metadata)
	}
	return nil
}

const opGetPublicAccessBlock = "GetPublicAccessBlock"

// GetPublicAccessBlockRequest returns a request value for making API operation for
// AWS S3 Control.
//
//    // Example sending a request using GetPublicAccessBlockRequest.
//    req := client.GetPublicAccessBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/GetPublicAccessBlock
func (c *Client) GetPublicAccessBlockRequest(input *GetPublicAccessBlockInput) GetPublicAccessBlockRequest {
	op := &aws.Operation{
		Name:       opGetPublicAccessBlock,
		HTTPMethod: "GET",
		HTTPPath:   "/v20180820/configuration/publicAccessBlock",
	}

	if input == nil {
		input = &GetPublicAccessBlockInput{}
	}

	req := c.newRequest(op, input, &GetPublicAccessBlockOutput{})
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))
	return GetPublicAccessBlockRequest{Request: req, Input: input, Copy: c.GetPublicAccessBlockRequest}
}

// GetPublicAccessBlockRequest is the request type for the
// GetPublicAccessBlock API operation.
type GetPublicAccessBlockRequest struct {
	*aws.Request
	Input *GetPublicAccessBlockInput
	Copy  func(*GetPublicAccessBlockInput) GetPublicAccessBlockRequest
}

// Send marshals and sends the GetPublicAccessBlock API request.
func (r GetPublicAccessBlockRequest) Send(ctx context.Context) (*GetPublicAccessBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPublicAccessBlockResponse{
		GetPublicAccessBlockOutput: r.Request.Data.(*GetPublicAccessBlockOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPublicAccessBlockResponse is the response type for the
// GetPublicAccessBlock API operation.
type GetPublicAccessBlockResponse struct {
	*GetPublicAccessBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPublicAccessBlock request.
func (r *GetPublicAccessBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
