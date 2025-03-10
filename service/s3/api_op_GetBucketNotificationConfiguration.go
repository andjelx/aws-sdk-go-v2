// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetBucketNotificationConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Name of the bucket to get the notification configuration for.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketNotificationConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketNotificationConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketNotificationConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketNotificationConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketNotificationConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

// A container for specifying the notification configuration of the bucket.
// If this element is empty, notifications are turned off for the bucket.
type GetBucketNotificationConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Describes the AWS Lambda functions to invoke and the events for which to
	// invoke them.
	LambdaFunctionConfigurations []LambdaFunctionConfiguration `locationName:"CloudFunctionConfiguration" type:"list" flattened:"true"`

	// The Amazon Simple Queue Service queues to publish messages to and the events
	// for which to publish messages.
	QueueConfigurations []QueueConfiguration `locationName:"QueueConfiguration" type:"list" flattened:"true"`

	// The topic to which notifications are sent and the events for which notifications
	// are generated.
	TopicConfigurations []TopicConfiguration `locationName:"TopicConfiguration" type:"list" flattened:"true"`
}

// String returns the string representation
func (s GetBucketNotificationConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketNotificationConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LambdaFunctionConfigurations != nil {
		v := s.LambdaFunctionConfigurations

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "CloudFunctionConfiguration", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.QueueConfigurations != nil {
		v := s.QueueConfigurations

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "QueueConfiguration", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.TopicConfigurations != nil {
		v := s.TopicConfigurations

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "TopicConfiguration", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetBucketNotificationConfiguration = "GetBucketNotificationConfiguration"

// GetBucketNotificationConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the notification configuration of a bucket.
//
//    // Example sending a request using GetBucketNotificationConfigurationRequest.
//    req := client.GetBucketNotificationConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketNotificationConfiguration
func (c *Client) GetBucketNotificationConfigurationRequest(input *GetBucketNotificationConfigurationInput) GetBucketNotificationConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetBucketNotificationConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?notification",
	}

	if input == nil {
		input = &GetBucketNotificationConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetBucketNotificationConfigurationOutput{})
	return GetBucketNotificationConfigurationRequest{Request: req, Input: input, Copy: c.GetBucketNotificationConfigurationRequest}
}

// GetBucketNotificationConfigurationRequest is the request type for the
// GetBucketNotificationConfiguration API operation.
type GetBucketNotificationConfigurationRequest struct {
	*aws.Request
	Input *GetBucketNotificationConfigurationInput
	Copy  func(*GetBucketNotificationConfigurationInput) GetBucketNotificationConfigurationRequest
}

// Send marshals and sends the GetBucketNotificationConfiguration API request.
func (r GetBucketNotificationConfigurationRequest) Send(ctx context.Context) (*GetBucketNotificationConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketNotificationConfigurationResponse{
		GetBucketNotificationConfigurationOutput: r.Request.Data.(*GetBucketNotificationConfigurationOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketNotificationConfigurationResponse is the response type for the
// GetBucketNotificationConfiguration API operation.
type GetBucketNotificationConfigurationResponse struct {
	*GetBucketNotificationConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketNotificationConfiguration request.
func (r *GetBucketNotificationConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
