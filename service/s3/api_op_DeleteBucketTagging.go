// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type DeleteBucketTaggingInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBucketTaggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBucketTaggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBucketTaggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteBucketTaggingInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketTaggingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

type DeleteBucketTaggingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBucketTaggingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketTaggingOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBucketTagging = "DeleteBucketTagging"

// DeleteBucketTaggingRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Deletes the tags from the bucket.
//
//    // Example sending a request using DeleteBucketTaggingRequest.
//    req := client.DeleteBucketTaggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketTagging
func (c *Client) DeleteBucketTaggingRequest(input *DeleteBucketTaggingInput) DeleteBucketTaggingRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketTagging,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?tagging",
	}

	if input == nil {
		input = &DeleteBucketTaggingInput{}
	}

	req := c.newRequest(op, input, &DeleteBucketTaggingOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketTaggingRequest{Request: req, Input: input, Copy: c.DeleteBucketTaggingRequest}
}

// DeleteBucketTaggingRequest is the request type for the
// DeleteBucketTagging API operation.
type DeleteBucketTaggingRequest struct {
	*aws.Request
	Input *DeleteBucketTaggingInput
	Copy  func(*DeleteBucketTaggingInput) DeleteBucketTaggingRequest
}

// Send marshals and sends the DeleteBucketTagging API request.
func (r DeleteBucketTaggingRequest) Send(ctx context.Context) (*DeleteBucketTaggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketTaggingResponse{
		DeleteBucketTaggingOutput: r.Request.Data.(*DeleteBucketTaggingOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketTaggingResponse is the response type for the
// DeleteBucketTagging API operation.
type DeleteBucketTaggingResponse struct {
	*DeleteBucketTaggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketTagging request.
func (r *DeleteBucketTaggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
