// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateRouteResponseInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ModelSelectionExpression *string `locationName:"modelSelectionExpression" type:"string"`

	// The route models.
	ResponseModels map[string]string `locationName:"responseModels" type:"map"`

	// The route parameters.
	ResponseParameters map[string]ParameterConstraints `locationName:"responseParameters" type:"map"`

	// RouteId is a required field
	RouteId *string `location:"uri" locationName:"routeId" type:"string" required:"true"`

	// RouteResponseId is a required field
	RouteResponseId *string `location:"uri" locationName:"routeResponseId" type:"string" required:"true"`

	// After evaulating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	RouteResponseKey *string `locationName:"routeResponseKey" type:"string"`
}

// String returns the string representation
func (s UpdateRouteResponseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRouteResponseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRouteResponseInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.RouteId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteId"))
	}

	if s.RouteResponseId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteResponseId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateRouteResponseInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ModelSelectionExpression != nil {
		v := *s.ModelSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "modelSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResponseModels != nil {
		v := s.ResponseModels

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseModels", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ResponseParameters != nil {
		v := s.ResponseParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.RouteResponseKey != nil {
		v := *s.RouteResponseKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeResponseKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteId != nil {
		v := *s.RouteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "routeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteResponseId != nil {
		v := *s.RouteResponseId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "routeResponseId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateRouteResponseOutput struct {
	_ struct{} `type:"structure"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ModelSelectionExpression *string `locationName:"modelSelectionExpression" type:"string"`

	// The route models.
	ResponseModels map[string]string `locationName:"responseModels" type:"map"`

	// The route parameters.
	ResponseParameters map[string]ParameterConstraints `locationName:"responseParameters" type:"map"`

	// The identifier.
	RouteResponseId *string `locationName:"routeResponseId" type:"string"`

	// After evaulating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	RouteResponseKey *string `locationName:"routeResponseKey" type:"string"`
}

// String returns the string representation
func (s UpdateRouteResponseOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateRouteResponseOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ModelSelectionExpression != nil {
		v := *s.ModelSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "modelSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResponseModels != nil {
		v := s.ResponseModels

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseModels", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ResponseParameters != nil {
		v := s.ResponseParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.RouteResponseId != nil {
		v := *s.RouteResponseId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeResponseId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteResponseKey != nil {
		v := *s.RouteResponseKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeResponseKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateRouteResponse = "UpdateRouteResponse"

// UpdateRouteResponseRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Updates a RouteResponse.
//
//    // Example sending a request using UpdateRouteResponseRequest.
//    req := client.UpdateRouteResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateRouteResponse
func (c *Client) UpdateRouteResponseRequest(input *UpdateRouteResponseInput) UpdateRouteResponseRequest {
	op := &aws.Operation{
		Name:       opUpdateRouteResponse,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v2/apis/{apiId}/routes/{routeId}/routeresponses/{routeResponseId}",
	}

	if input == nil {
		input = &UpdateRouteResponseInput{}
	}

	req := c.newRequest(op, input, &UpdateRouteResponseOutput{})
	return UpdateRouteResponseRequest{Request: req, Input: input, Copy: c.UpdateRouteResponseRequest}
}

// UpdateRouteResponseRequest is the request type for the
// UpdateRouteResponse API operation.
type UpdateRouteResponseRequest struct {
	*aws.Request
	Input *UpdateRouteResponseInput
	Copy  func(*UpdateRouteResponseInput) UpdateRouteResponseRequest
}

// Send marshals and sends the UpdateRouteResponse API request.
func (r UpdateRouteResponseRequest) Send(ctx context.Context) (*UpdateRouteResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRouteResponseResponse{
		UpdateRouteResponseOutput: r.Request.Data.(*UpdateRouteResponseOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRouteResponseResponse is the response type for the
// UpdateRouteResponse API operation.
type UpdateRouteResponseResponse struct {
	*UpdateRouteResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRouteResponse request.
func (r *UpdateRouteResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
