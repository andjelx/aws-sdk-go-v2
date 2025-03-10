// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DisassociateResourceShareInput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The principals.
	Principals []string `locationName:"principals" type:"list"`

	// The Amazon Resource Names (ARN) of the resources.
	ResourceArns []string `locationName:"resourceArns" type:"list"`

	// The Amazon Resource Name (ARN) of the resource share.
	//
	// ResourceShareArn is a required field
	ResourceShareArn *string `locationName:"resourceShareArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateResourceShareInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateResourceShareInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateResourceShareInput"}

	if s.ResourceShareArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceShareArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateResourceShareInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Principals != nil {
		v := s.Principals

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "principals", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ResourceArns != nil {
		v := s.ResourceArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ResourceShareArn != nil {
		v := *s.ResourceShareArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DisassociateResourceShareOutput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the associations.
	ResourceShareAssociations []ResourceShareAssociation `locationName:"resourceShareAssociations" type:"list"`
}

// String returns the string representation
func (s DisassociateResourceShareOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateResourceShareOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareAssociations != nil {
		v := s.ResourceShareAssociations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceShareAssociations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDisassociateResourceShare = "DisassociateResourceShare"

// DisassociateResourceShareRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Disassociates the specified principals or resources from the specified resource
// share.
//
//    // Example sending a request using DisassociateResourceShareRequest.
//    req := client.DisassociateResourceShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/DisassociateResourceShare
func (c *Client) DisassociateResourceShareRequest(input *DisassociateResourceShareInput) DisassociateResourceShareRequest {
	op := &aws.Operation{
		Name:       opDisassociateResourceShare,
		HTTPMethod: "POST",
		HTTPPath:   "/disassociateresourceshare",
	}

	if input == nil {
		input = &DisassociateResourceShareInput{}
	}

	req := c.newRequest(op, input, &DisassociateResourceShareOutput{})
	return DisassociateResourceShareRequest{Request: req, Input: input, Copy: c.DisassociateResourceShareRequest}
}

// DisassociateResourceShareRequest is the request type for the
// DisassociateResourceShare API operation.
type DisassociateResourceShareRequest struct {
	*aws.Request
	Input *DisassociateResourceShareInput
	Copy  func(*DisassociateResourceShareInput) DisassociateResourceShareRequest
}

// Send marshals and sends the DisassociateResourceShare API request.
func (r DisassociateResourceShareRequest) Send(ctx context.Context) (*DisassociateResourceShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateResourceShareResponse{
		DisassociateResourceShareOutput: r.Request.Data.(*DisassociateResourceShareOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateResourceShareResponse is the response type for the
// DisassociateResourceShare API operation.
type DisassociateResourceShareResponse struct {
	*DisassociateResourceShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateResourceShare request.
func (r *DisassociateResourceShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
