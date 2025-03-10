// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateDomainNameInput struct {
	_ struct{} `type:"structure"`

	// A string with a length between [1-512].
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`

	// The domain name configurations.
	DomainNameConfigurations []DomainNameConfiguration `locationName:"domainNameConfigurations" type:"list"`

	// A key value pair of string with key length between[1-128] and value length
	// between[1-256]
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateDomainNameInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDomainNameInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDomainNameInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDomainNameInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainNameConfigurations != nil {
		v := s.DomainNameConfigurations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "domainNameConfigurations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

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

type CreateDomainNameOutput struct {
	_ struct{} `type:"structure"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ApiMappingSelectionExpression *string `locationName:"apiMappingSelectionExpression" type:"string"`

	// A string with a length between [1-512].
	DomainName *string `locationName:"domainName" type:"string"`

	// The domain name configurations.
	DomainNameConfigurations []DomainNameConfiguration `locationName:"domainNameConfigurations" type:"list"`

	// A key value pair of string with key length between[1-128] and value length
	// between[1-256]
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateDomainNameOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDomainNameOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiMappingSelectionExpression != nil {
		v := *s.ApiMappingSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiMappingSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainNameConfigurations != nil {
		v := s.DomainNameConfigurations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "domainNameConfigurations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

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

const opCreateDomainName = "CreateDomainName"

// CreateDomainNameRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Creates a domain name.
//
//    // Example sending a request using CreateDomainNameRequest.
//    req := client.CreateDomainNameRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateDomainName
func (c *Client) CreateDomainNameRequest(input *CreateDomainNameInput) CreateDomainNameRequest {
	op := &aws.Operation{
		Name:       opCreateDomainName,
		HTTPMethod: "POST",
		HTTPPath:   "/v2/domainnames",
	}

	if input == nil {
		input = &CreateDomainNameInput{}
	}

	req := c.newRequest(op, input, &CreateDomainNameOutput{})
	return CreateDomainNameRequest{Request: req, Input: input, Copy: c.CreateDomainNameRequest}
}

// CreateDomainNameRequest is the request type for the
// CreateDomainName API operation.
type CreateDomainNameRequest struct {
	*aws.Request
	Input *CreateDomainNameInput
	Copy  func(*CreateDomainNameInput) CreateDomainNameRequest
}

// Send marshals and sends the CreateDomainName API request.
func (r CreateDomainNameRequest) Send(ctx context.Context) (*CreateDomainNameResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDomainNameResponse{
		CreateDomainNameOutput: r.Request.Data.(*CreateDomainNameOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDomainNameResponse is the response type for the
// CreateDomainName API operation.
type CreateDomainNameResponse struct {
	*CreateDomainNameOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDomainName request.
func (r *CreateDomainNameResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
