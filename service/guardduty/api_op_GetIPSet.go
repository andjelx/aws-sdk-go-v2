// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetIPSetInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector the ipSet is associated with.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The unique ID of the ipSet you want to get.
	//
	// IpSetId is a required field
	IpSetId *string `location:"uri" locationName:"ipSetId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetIPSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIPSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIPSetInput"}

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
func (s GetIPSetInput) MarshalFields(e protocol.FieldEncoder) error {
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

type GetIPSetOutput struct {
	_ struct{} `type:"structure"`

	// The format of the file that contains the IPSet.
	//
	// Format is a required field
	Format IpSetFormat `locationName:"format" min:"1" type:"string" required:"true" enum:"true"`

	// The URI of the file that contains the IPSet. For example (https://s3.us-west-2.amazonaws.com/my-bucket/my-object-key)
	//
	// Location is a required field
	Location *string `locationName:"location" min:"1" type:"string" required:"true"`

	// The user friendly name to identify the IPSet. This name is displayed in all
	// findings that are triggered by activity that involves IP addresses included
	// in this IPSet.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The status of ipSet file uploaded.
	//
	// Status is a required field
	Status IpSetStatus `locationName:"status" min:"1" type:"string" required:"true" enum:"true"`

	// The tags of the IP set resource.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s GetIPSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIPSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "location", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opGetIPSet = "GetIPSet"

// GetIPSetRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Retrieves the IPSet specified by the IPSet ID.
//
//    // Example sending a request using GetIPSetRequest.
//    req := client.GetIPSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetIPSet
func (c *Client) GetIPSetRequest(input *GetIPSetInput) GetIPSetRequest {
	op := &aws.Operation{
		Name:       opGetIPSet,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}/ipset/{ipSetId}",
	}

	if input == nil {
		input = &GetIPSetInput{}
	}

	req := c.newRequest(op, input, &GetIPSetOutput{})
	return GetIPSetRequest{Request: req, Input: input, Copy: c.GetIPSetRequest}
}

// GetIPSetRequest is the request type for the
// GetIPSet API operation.
type GetIPSetRequest struct {
	*aws.Request
	Input *GetIPSetInput
	Copy  func(*GetIPSetInput) GetIPSetRequest
}

// Send marshals and sends the GetIPSet API request.
func (r GetIPSetRequest) Send(ctx context.Context) (*GetIPSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIPSetResponse{
		GetIPSetOutput: r.Request.Data.(*GetIPSetOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIPSetResponse is the response type for the
// GetIPSet API operation.
type GetIPSetResponse struct {
	*GetIPSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIPSet request.
func (r *GetIPSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
