// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetStageInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// StageName is a required field
	StageName *string `location:"uri" locationName:"stageName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetStageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetStageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetStageInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.StageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StageName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetStageInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageName != nil {
		v := *s.StageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "stageName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetStageOutput struct {
	_ struct{} `type:"structure"`

	// Settings for logging access in a stage.
	AccessLogSettings *AccessLogSettings `locationName:"accessLogSettings" type:"structure"`

	// The identifier.
	ClientCertificateId *string `locationName:"clientCertificateId" type:"string"`

	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"iso8601"`

	// Represents a collection of route settings.
	DefaultRouteSettings *RouteSettings `locationName:"defaultRouteSettings" type:"structure"`

	// The identifier.
	DeploymentId *string `locationName:"deploymentId" type:"string"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	LastUpdatedDate *time.Time `locationName:"lastUpdatedDate" type:"timestamp" timestampFormat:"iso8601"`

	// The route settings map.
	RouteSettings map[string]RouteSettings `locationName:"routeSettings" type:"map"`

	// A string with a length between [1-128].
	StageName *string `locationName:"stageName" type:"string"`

	// The stage variable map.
	StageVariables map[string]string `locationName:"stageVariables" type:"map"`

	// A key value pair of string with key length between[1-128] and value length
	// between[1-256]
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetStageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetStageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccessLogSettings != nil {
		v := s.AccessLogSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "accessLogSettings", v, metadata)
	}
	if s.ClientCertificateId != nil {
		v := *s.ClientCertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientCertificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.DefaultRouteSettings != nil {
		v := s.DefaultRouteSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "defaultRouteSettings", v, metadata)
	}
	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedDate != nil {
		v := *s.LastUpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.RouteSettings != nil {
		v := s.RouteSettings

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "routeSettings", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.StageName != nil {
		v := *s.StageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stageName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageVariables != nil {
		v := s.StageVariables

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "stageVariables", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

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

const opGetStage = "GetStage"

// GetStageRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets a Stage.
//
//    // Example sending a request using GetStageRequest.
//    req := client.GetStageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetStage
func (c *Client) GetStageRequest(input *GetStageInput) GetStageRequest {
	op := &aws.Operation{
		Name:       opGetStage,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/stages/{stageName}",
	}

	if input == nil {
		input = &GetStageInput{}
	}

	req := c.newRequest(op, input, &GetStageOutput{})
	return GetStageRequest{Request: req, Input: input, Copy: c.GetStageRequest}
}

// GetStageRequest is the request type for the
// GetStage API operation.
type GetStageRequest struct {
	*aws.Request
	Input *GetStageInput
	Copy  func(*GetStageInput) GetStageRequest
}

// Send marshals and sends the GetStage API request.
func (r GetStageRequest) Send(ctx context.Context) (*GetStageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetStageResponse{
		GetStageOutput: r.Request.Data.(*GetStageOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetStageResponse is the response type for the
// GetStage API operation.
type GetStageResponse struct {
	*GetStageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetStage request.
func (r *GetStageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
