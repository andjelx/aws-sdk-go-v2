// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetObjectLegalHoldInput struct {
	_ struct{} `type:"structure"`

	// The bucket containing the object whose Legal Hold status you want to retrieve.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The key name for the object whose Legal Hold status you want to retrieve.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// The version ID of the object whose Legal Hold status you want to retrieve.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s GetObjectLegalHoldInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectLegalHoldInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectLegalHoldInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetObjectLegalHoldInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectLegalHoldInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetObjectLegalHoldOutput struct {
	_ struct{} `type:"structure" payload:"LegalHold"`

	// The current Legal Hold status for the specified object.
	LegalHold *ObjectLockLegalHold `type:"structure"`
}

// String returns the string representation
func (s GetObjectLegalHoldOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectLegalHoldOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LegalHold != nil {
		v := s.LegalHold

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "LegalHold", v, metadata)
	}
	return nil
}

const opGetObjectLegalHold = "GetObjectLegalHold"

// GetObjectLegalHoldRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Gets an object's current Legal Hold status.
//
//    // Example sending a request using GetObjectLegalHoldRequest.
//    req := client.GetObjectLegalHoldRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectLegalHold
func (c *Client) GetObjectLegalHoldRequest(input *GetObjectLegalHoldInput) GetObjectLegalHoldRequest {
	op := &aws.Operation{
		Name:       opGetObjectLegalHold,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}?legal-hold",
	}

	if input == nil {
		input = &GetObjectLegalHoldInput{}
	}

	req := c.newRequest(op, input, &GetObjectLegalHoldOutput{})
	return GetObjectLegalHoldRequest{Request: req, Input: input, Copy: c.GetObjectLegalHoldRequest}
}

// GetObjectLegalHoldRequest is the request type for the
// GetObjectLegalHold API operation.
type GetObjectLegalHoldRequest struct {
	*aws.Request
	Input *GetObjectLegalHoldInput
	Copy  func(*GetObjectLegalHoldInput) GetObjectLegalHoldRequest
}

// Send marshals and sends the GetObjectLegalHold API request.
func (r GetObjectLegalHoldRequest) Send(ctx context.Context) (*GetObjectLegalHoldResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectLegalHoldResponse{
		GetObjectLegalHoldOutput: r.Request.Data.(*GetObjectLegalHoldOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectLegalHoldResponse is the response type for the
// GetObjectLegalHold API operation.
type GetObjectLegalHoldResponse struct {
	*GetObjectLegalHoldOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObjectLegalHold request.
func (r *GetObjectLegalHoldResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
