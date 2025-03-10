// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateSampleFindingsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the detector to create sample findings for.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// Types of sample findings that you want to generate.
	FindingTypes []string `locationName:"findingTypes" type:"list"`
}

// String returns the string representation
func (s CreateSampleFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSampleFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSampleFindingsInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSampleFindingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FindingTypes != nil {
		v := s.FindingTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "findingTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateSampleFindingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateSampleFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSampleFindingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateSampleFindings = "CreateSampleFindings"

// CreateSampleFindingsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Generates example findings of types specified by the list of finding types.
// If 'NULL' is specified for findingTypes, the API generates example findings
// of all supported finding types.
//
//    // Example sending a request using CreateSampleFindingsRequest.
//    req := client.CreateSampleFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/CreateSampleFindings
func (c *Client) CreateSampleFindingsRequest(input *CreateSampleFindingsInput) CreateSampleFindingsRequest {
	op := &aws.Operation{
		Name:       opCreateSampleFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/findings/create",
	}

	if input == nil {
		input = &CreateSampleFindingsInput{}
	}

	req := c.newRequest(op, input, &CreateSampleFindingsOutput{})
	return CreateSampleFindingsRequest{Request: req, Input: input, Copy: c.CreateSampleFindingsRequest}
}

// CreateSampleFindingsRequest is the request type for the
// CreateSampleFindings API operation.
type CreateSampleFindingsRequest struct {
	*aws.Request
	Input *CreateSampleFindingsInput
	Copy  func(*CreateSampleFindingsInput) CreateSampleFindingsRequest
}

// Send marshals and sends the CreateSampleFindings API request.
func (r CreateSampleFindingsRequest) Send(ctx context.Context) (*CreateSampleFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSampleFindingsResponse{
		CreateSampleFindingsOutput: r.Request.Data.(*CreateSampleFindingsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSampleFindingsResponse is the response type for the
// CreateSampleFindings API operation.
type CreateSampleFindingsResponse struct {
	*CreateSampleFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSampleFindings request.
func (r *CreateSampleFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
