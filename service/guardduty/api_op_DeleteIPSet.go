// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteIPSetInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector the ipSet is associated with.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The unique ID of the ipSet you want to delete.
	//
	// IpSetId is a required field
	IpSetId *string `location:"uri" locationName:"ipSetId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteIPSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIPSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIPSetInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.IpSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IpSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteIPSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IpSetId != nil {
		v := *s.IpSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ipSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteIPSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteIPSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteIPSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteIPSet = "DeleteIPSet"

// DeleteIPSetRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Deletes the IPSet specified by the IPSet ID.
//
//    // Example sending a request using DeleteIPSetRequest.
//    req := client.DeleteIPSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/DeleteIPSet
func (c *Client) DeleteIPSetRequest(input *DeleteIPSetInput) DeleteIPSetRequest {
	op := &aws.Operation{
		Name:       opDeleteIPSet,
		HTTPMethod: "DELETE",
		HTTPPath:   "/detector/{detectorId}/ipset/{ipSetId}",
	}

	if input == nil {
		input = &DeleteIPSetInput{}
	}

	req := c.newRequest(op, input, &DeleteIPSetOutput{})
	return DeleteIPSetRequest{Request: req, Input: input, Copy: c.DeleteIPSetRequest}
}

// DeleteIPSetRequest is the request type for the
// DeleteIPSet API operation.
type DeleteIPSetRequest struct {
	*aws.Request
	Input *DeleteIPSetInput
	Copy  func(*DeleteIPSetInput) DeleteIPSetRequest
}

// Send marshals and sends the DeleteIPSet API request.
func (r DeleteIPSetRequest) Send(ctx context.Context) (*DeleteIPSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteIPSetResponse{
		DeleteIPSetOutput: r.Request.Data.(*DeleteIPSetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteIPSetResponse is the response type for the
// DeleteIPSet API operation.
type DeleteIPSetResponse struct {
	*DeleteIPSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteIPSet request.
func (r *DeleteIPSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
