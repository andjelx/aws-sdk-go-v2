// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateApiMappingInput struct {
	_ struct{} `type:"structure"`

	// The identifier.
	//
	// ApiId is a required field
	ApiId *string `locationName:"apiId" type:"string" required:"true"`

	// ApiMappingId is a required field
	ApiMappingId *string `location:"uri" locationName:"apiMappingId" type:"string" required:"true"`

	// After evaulating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	ApiMappingKey *string `locationName:"apiMappingKey" type:"string"`

	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`

	// A string with a length between [1-128].
	Stage *string `locationName:"stage" type:"string"`
}

// String returns the string representation
func (s UpdateApiMappingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApiMappingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApiMappingInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.ApiMappingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiMappingId"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApiMappingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiMappingKey != nil {
		v := *s.ApiMappingKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiMappingKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Stage != nil {
		v := *s.Stage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiMappingId != nil {
		v := *s.ApiMappingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiMappingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateApiMappingOutput struct {
	_ struct{} `type:"structure"`

	// The identifier.
	ApiId *string `locationName:"apiId" type:"string"`

	// The identifier.
	ApiMappingId *string `locationName:"apiMappingId" type:"string"`

	// After evaulating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	ApiMappingKey *string `locationName:"apiMappingKey" type:"string"`

	// A string with a length between [1-128].
	Stage *string `locationName:"stage" type:"string"`
}

// String returns the string representation
func (s UpdateApiMappingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApiMappingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiMappingId != nil {
		v := *s.ApiMappingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiMappingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiMappingKey != nil {
		v := *s.ApiMappingKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiMappingKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Stage != nil {
		v := *s.Stage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateApiMapping = "UpdateApiMapping"

// UpdateApiMappingRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// The API mapping.
//
//    // Example sending a request using UpdateApiMappingRequest.
//    req := client.UpdateApiMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateApiMapping
func (c *Client) UpdateApiMappingRequest(input *UpdateApiMappingInput) UpdateApiMappingRequest {
	op := &aws.Operation{
		Name:       opUpdateApiMapping,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v2/domainnames/{domainName}/apimappings/{apiMappingId}",
	}

	if input == nil {
		input = &UpdateApiMappingInput{}
	}

	req := c.newRequest(op, input, &UpdateApiMappingOutput{})
	return UpdateApiMappingRequest{Request: req, Input: input, Copy: c.UpdateApiMappingRequest}
}

// UpdateApiMappingRequest is the request type for the
// UpdateApiMapping API operation.
type UpdateApiMappingRequest struct {
	*aws.Request
	Input *UpdateApiMappingInput
	Copy  func(*UpdateApiMappingInput) UpdateApiMappingRequest
}

// Send marshals and sends the UpdateApiMapping API request.
func (r UpdateApiMappingRequest) Send(ctx context.Context) (*UpdateApiMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApiMappingResponse{
		UpdateApiMappingOutput: r.Request.Data.(*UpdateApiMappingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApiMappingResponse is the response type for the
// UpdateApiMapping API operation.
type UpdateApiMappingResponse struct {
	*UpdateApiMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApiMapping request.
func (r *UpdateApiMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
