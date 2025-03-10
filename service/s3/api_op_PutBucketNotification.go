// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type PutBucketNotificationInput struct {
	_ struct{} `type:"structure" payload:"NotificationConfiguration"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// NotificationConfiguration is a required field
	NotificationConfiguration *NotificationConfigurationDeprecated `locationName:"NotificationConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketNotificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketNotificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketNotificationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.NotificationConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotificationConfiguration"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketNotificationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketNotificationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.NotificationConfiguration != nil {
		v := s.NotificationConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "NotificationConfiguration", v, metadata)
	}
	return nil
}

type PutBucketNotificationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketNotificationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketNotificationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketNotification = "PutBucketNotification"

// PutBucketNotificationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// No longer used, see the PutBucketNotificationConfiguration operation.
//
//    // Example sending a request using PutBucketNotificationRequest.
//    req := client.PutBucketNotificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketNotification
func (c *Client) PutBucketNotificationRequest(input *PutBucketNotificationInput) PutBucketNotificationRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, PutBucketNotification, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opPutBucketNotification,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?notification",
	}

	if input == nil {
		input = &PutBucketNotificationInput{}
	}

	req := c.newRequest(op, input, &PutBucketNotificationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketNotificationRequest{Request: req, Input: input, Copy: c.PutBucketNotificationRequest}
}

// PutBucketNotificationRequest is the request type for the
// PutBucketNotification API operation.
type PutBucketNotificationRequest struct {
	*aws.Request
	Input *PutBucketNotificationInput
	Copy  func(*PutBucketNotificationInput) PutBucketNotificationRequest
}

// Send marshals and sends the PutBucketNotification API request.
func (r PutBucketNotificationRequest) Send(ctx context.Context) (*PutBucketNotificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketNotificationResponse{
		PutBucketNotificationOutput: r.Request.Data.(*PutBucketNotificationOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketNotificationResponse is the response type for the
// PutBucketNotification API operation.
type PutBucketNotificationResponse struct {
	*PutBucketNotificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketNotification request.
func (r *PutBucketNotificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
