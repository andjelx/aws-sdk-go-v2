// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type PutBucketInventoryConfigurationInput struct {
	_ struct{} `type:"structure" payload:"InventoryConfiguration"`

	// The name of the bucket where the inventory configuration will be stored.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The ID used to identify the inventory configuration.
	//
	// Id is a required field
	Id *string `location:"querystring" locationName:"id" type:"string" required:"true"`

	// Specifies the inventory configuration.
	//
	// InventoryConfiguration is a required field
	InventoryConfiguration *InventoryConfiguration `locationName:"InventoryConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketInventoryConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketInventoryConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketInventoryConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.InventoryConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("InventoryConfiguration"))
	}
	if s.InventoryConfiguration != nil {
		if err := s.InventoryConfiguration.Validate(); err != nil {
			invalidParams.AddNested("InventoryConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketInventoryConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketInventoryConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.InventoryConfiguration != nil {
		v := s.InventoryConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "InventoryConfiguration", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "id", protocol.StringValue(v), metadata)
	}
	return nil
}

type PutBucketInventoryConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketInventoryConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketInventoryConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketInventoryConfiguration = "PutBucketInventoryConfiguration"

// PutBucketInventoryConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Adds an inventory configuration (identified by the inventory ID) from the
// bucket.
//
//    // Example sending a request using PutBucketInventoryConfigurationRequest.
//    req := client.PutBucketInventoryConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketInventoryConfiguration
func (c *Client) PutBucketInventoryConfigurationRequest(input *PutBucketInventoryConfigurationInput) PutBucketInventoryConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutBucketInventoryConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?inventory",
	}

	if input == nil {
		input = &PutBucketInventoryConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutBucketInventoryConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketInventoryConfigurationRequest{Request: req, Input: input, Copy: c.PutBucketInventoryConfigurationRequest}
}

// PutBucketInventoryConfigurationRequest is the request type for the
// PutBucketInventoryConfiguration API operation.
type PutBucketInventoryConfigurationRequest struct {
	*aws.Request
	Input *PutBucketInventoryConfigurationInput
	Copy  func(*PutBucketInventoryConfigurationInput) PutBucketInventoryConfigurationRequest
}

// Send marshals and sends the PutBucketInventoryConfiguration API request.
func (r PutBucketInventoryConfigurationRequest) Send(ctx context.Context) (*PutBucketInventoryConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketInventoryConfigurationResponse{
		PutBucketInventoryConfigurationOutput: r.Request.Data.(*PutBucketInventoryConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketInventoryConfigurationResponse is the response type for the
// PutBucketInventoryConfiguration API operation.
type PutBucketInventoryConfigurationResponse struct {
	*PutBucketInventoryConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketInventoryConfiguration request.
func (r *PutBucketInventoryConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
