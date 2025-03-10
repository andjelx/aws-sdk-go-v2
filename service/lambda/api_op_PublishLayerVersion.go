// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PublishLayerVersionInput struct {
	_ struct{} `type:"structure"`

	// A list of compatible function runtimes (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).
	// Used for filtering with ListLayers and ListLayerVersions.
	CompatibleRuntimes []Runtime `type:"list"`

	// The function layer archive.
	//
	// Content is a required field
	Content *LayerVersionContentInput `type:"structure" required:"true"`

	// The description of the version.
	Description *string `type:"string"`

	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// LayerName is a required field
	LayerName *string `location:"uri" locationName:"LayerName" min:"1" type:"string" required:"true"`

	// The layer's software license. It can be any of the following:
	//
	//    * An SPDX license identifier (https://spdx.org/licenses/). For example,
	//    MIT.
	//
	//    * The URL of a license hosted on the internet. For example, https://opensource.org/licenses/MIT.
	//
	//    * The full text of the license.
	LicenseInfo *string `type:"string"`
}

// String returns the string representation
func (s PublishLayerVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PublishLayerVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PublishLayerVersionInput"}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}

	if s.LayerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerName"))
	}
	if s.LayerName != nil && len(*s.LayerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LayerName", 1))
	}
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			invalidParams.AddNested("Content", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PublishLayerVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CompatibleRuntimes != nil {
		v := s.CompatibleRuntimes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "CompatibleRuntimes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Content != nil {
		v := s.Content

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Content", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LicenseInfo != nil {
		v := *s.LicenseInfo

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LicenseInfo", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LayerName != nil {
		v := *s.LayerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "LayerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PublishLayerVersionOutput struct {
	_ struct{} `type:"structure"`

	// The layer's compatible runtimes.
	CompatibleRuntimes []Runtime `type:"list"`

	// Details about the layer version.
	Content *LayerVersionContentOutput `type:"structure"`

	// The date that the layer version was created, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	CreatedDate *string `type:"string"`

	// The description of the version.
	Description *string `type:"string"`

	// The ARN of the layer.
	LayerArn *string `min:"1" type:"string"`

	// The ARN of the layer version.
	LayerVersionArn *string `min:"1" type:"string"`

	// The layer's software license.
	LicenseInfo *string `type:"string"`

	// The version number.
	Version *int64 `type:"long"`
}

// String returns the string representation
func (s PublishLayerVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PublishLayerVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CompatibleRuntimes != nil {
		v := s.CompatibleRuntimes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "CompatibleRuntimes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Content != nil {
		v := s.Content

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Content", v, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatedDate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LayerArn != nil {
		v := *s.LayerArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LayerArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LayerVersionArn != nil {
		v := *s.LayerVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LayerVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LicenseInfo != nil {
		v := *s.LicenseInfo

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LicenseInfo", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opPublishLayerVersion = "PublishLayerVersion"

// PublishLayerVersionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Creates an AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
// from a ZIP archive. Each time you call PublishLayerVersion with the same
// version name, a new version is created.
//
// Add layers to your function with CreateFunction or UpdateFunctionConfiguration.
//
//    // Example sending a request using PublishLayerVersionRequest.
//    req := client.PublishLayerVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/PublishLayerVersion
func (c *Client) PublishLayerVersionRequest(input *PublishLayerVersionInput) PublishLayerVersionRequest {
	op := &aws.Operation{
		Name:       opPublishLayerVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/2018-10-31/layers/{LayerName}/versions",
	}

	if input == nil {
		input = &PublishLayerVersionInput{}
	}

	req := c.newRequest(op, input, &PublishLayerVersionOutput{})
	return PublishLayerVersionRequest{Request: req, Input: input, Copy: c.PublishLayerVersionRequest}
}

// PublishLayerVersionRequest is the request type for the
// PublishLayerVersion API operation.
type PublishLayerVersionRequest struct {
	*aws.Request
	Input *PublishLayerVersionInput
	Copy  func(*PublishLayerVersionInput) PublishLayerVersionRequest
}

// Send marshals and sends the PublishLayerVersion API request.
func (r PublishLayerVersionRequest) Send(ctx context.Context) (*PublishLayerVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PublishLayerVersionResponse{
		PublishLayerVersionOutput: r.Request.Data.(*PublishLayerVersionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PublishLayerVersionResponse is the response type for the
// PublishLayerVersion API operation.
type PublishLayerVersionResponse struct {
	*PublishLayerVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PublishLayerVersion request.
func (r *PublishLayerVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
