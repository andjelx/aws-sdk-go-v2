// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetPublicAccessBlockInput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon S3 bucket whose PublicAccessBlock configuration you
	// want to retrieve.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPublicAccessBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPublicAccessBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPublicAccessBlockInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetPublicAccessBlockInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPublicAccessBlockInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetPublicAccessBlockOutput struct {
	_ struct{} `type:"structure" payload:"PublicAccessBlockConfiguration"`

	// The PublicAccessBlock configuration currently in effect for this Amazon S3
	// bucket.
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
// Amazon Simple Storage Service.
//
// Retrieves the PublicAccessBlock configuration for an Amazon S3 bucket.
//
//    // Example sending a request using GetPublicAccessBlockRequest.
//    req := client.GetPublicAccessBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetPublicAccessBlock
func (c *Client) GetPublicAccessBlockRequest(input *GetPublicAccessBlockInput) GetPublicAccessBlockRequest {
	op := &aws.Operation{
		Name:       opGetPublicAccessBlock,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?publicAccessBlock",
	}

	if input == nil {
		input = &GetPublicAccessBlockInput{}
	}

	req := c.newRequest(op, input, &GetPublicAccessBlockOutput{})
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
