// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type PutBucketCorsInput struct {
	_ struct{} `type:"structure" payload:"CORSConfiguration"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Describes the cross-origin access configuration for objects in an Amazon
	// S3 bucket. For more information, see Enabling Cross-Origin Resource Sharing
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon
	// Simple Storage Service Developer Guide.
	//
	// CORSConfiguration is a required field
	CORSConfiguration *CORSConfiguration `locationName:"CORSConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketCorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketCorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketCorsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.CORSConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("CORSConfiguration"))
	}
	if s.CORSConfiguration != nil {
		if err := s.CORSConfiguration.Validate(); err != nil {
			invalidParams.AddNested("CORSConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketCorsInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketCorsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.CORSConfiguration != nil {
		v := s.CORSConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "CORSConfiguration", v, metadata)
	}
	return nil
}

type PutBucketCorsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketCorsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketCorsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketCors = "PutBucketCors"

// PutBucketCorsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Sets the CORS configuration for a bucket.
//
//    // Example sending a request using PutBucketCorsRequest.
//    req := client.PutBucketCorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketCors
func (c *Client) PutBucketCorsRequest(input *PutBucketCorsInput) PutBucketCorsRequest {
	op := &aws.Operation{
		Name:       opPutBucketCors,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?cors",
	}

	if input == nil {
		input = &PutBucketCorsInput{}
	}

	req := c.newRequest(op, input, &PutBucketCorsOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketCorsRequest{Request: req, Input: input, Copy: c.PutBucketCorsRequest}
}

// PutBucketCorsRequest is the request type for the
// PutBucketCors API operation.
type PutBucketCorsRequest struct {
	*aws.Request
	Input *PutBucketCorsInput
	Copy  func(*PutBucketCorsInput) PutBucketCorsRequest
}

// Send marshals and sends the PutBucketCors API request.
func (r PutBucketCorsRequest) Send(ctx context.Context) (*PutBucketCorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketCorsResponse{
		PutBucketCorsOutput: r.Request.Data.(*PutBucketCorsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketCorsResponse is the response type for the
// PutBucketCors API operation.
type PutBucketCorsResponse struct {
	*PutBucketCorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketCors request.
func (r *PutBucketCorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
