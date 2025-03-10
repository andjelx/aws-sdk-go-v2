// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetBucketEncryptionInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket from which the server-side encryption configuration
	// is retrieved.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketEncryptionInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketEncryptionInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketEncryptionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetBucketEncryptionOutput struct {
	_ struct{} `type:"structure" payload:"ServerSideEncryptionConfiguration"`

	// Specifies the default server-side-encryption configuration.
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetBucketEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketEncryptionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ServerSideEncryptionConfiguration != nil {
		v := s.ServerSideEncryptionConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ServerSideEncryptionConfiguration", v, metadata)
	}
	return nil
}

const opGetBucketEncryption = "GetBucketEncryption"

// GetBucketEncryptionRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the server-side encryption configuration of a bucket.
//
//    // Example sending a request using GetBucketEncryptionRequest.
//    req := client.GetBucketEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketEncryption
func (c *Client) GetBucketEncryptionRequest(input *GetBucketEncryptionInput) GetBucketEncryptionRequest {
	op := &aws.Operation{
		Name:       opGetBucketEncryption,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?encryption",
	}

	if input == nil {
		input = &GetBucketEncryptionInput{}
	}

	req := c.newRequest(op, input, &GetBucketEncryptionOutput{})
	return GetBucketEncryptionRequest{Request: req, Input: input, Copy: c.GetBucketEncryptionRequest}
}

// GetBucketEncryptionRequest is the request type for the
// GetBucketEncryption API operation.
type GetBucketEncryptionRequest struct {
	*aws.Request
	Input *GetBucketEncryptionInput
	Copy  func(*GetBucketEncryptionInput) GetBucketEncryptionRequest
}

// Send marshals and sends the GetBucketEncryption API request.
func (r GetBucketEncryptionRequest) Send(ctx context.Context) (*GetBucketEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketEncryptionResponse{
		GetBucketEncryptionOutput: r.Request.Data.(*GetBucketEncryptionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketEncryptionResponse is the response type for the
// GetBucketEncryption API operation.
type GetBucketEncryptionResponse struct {
	*GetBucketEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketEncryption request.
func (r *GetBucketEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
