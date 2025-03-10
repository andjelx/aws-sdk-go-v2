// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type DeleteBucketEncryptionInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket containing the server-side encryption configuration
	// to delete.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBucketEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBucketEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBucketEncryptionInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteBucketEncryptionInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketEncryptionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

type DeleteBucketEncryptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBucketEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketEncryptionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBucketEncryption = "DeleteBucketEncryption"

// DeleteBucketEncryptionRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Deletes the server-side encryption configuration from the bucket.
//
//    // Example sending a request using DeleteBucketEncryptionRequest.
//    req := client.DeleteBucketEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketEncryption
func (c *Client) DeleteBucketEncryptionRequest(input *DeleteBucketEncryptionInput) DeleteBucketEncryptionRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketEncryption,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?encryption",
	}

	if input == nil {
		input = &DeleteBucketEncryptionInput{}
	}

	req := c.newRequest(op, input, &DeleteBucketEncryptionOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketEncryptionRequest{Request: req, Input: input, Copy: c.DeleteBucketEncryptionRequest}
}

// DeleteBucketEncryptionRequest is the request type for the
// DeleteBucketEncryption API operation.
type DeleteBucketEncryptionRequest struct {
	*aws.Request
	Input *DeleteBucketEncryptionInput
	Copy  func(*DeleteBucketEncryptionInput) DeleteBucketEncryptionRequest
}

// Send marshals and sends the DeleteBucketEncryption API request.
func (r DeleteBucketEncryptionRequest) Send(ctx context.Context) (*DeleteBucketEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketEncryptionResponse{
		DeleteBucketEncryptionOutput: r.Request.Data.(*DeleteBucketEncryptionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketEncryptionResponse is the response type for the
// DeleteBucketEncryption API operation.
type DeleteBucketEncryptionResponse struct {
	*DeleteBucketEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketEncryption request.
func (r *DeleteBucketEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
