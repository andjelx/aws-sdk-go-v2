// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListResolversInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The maximum number of results you want the request to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The type name.
	//
	// TypeName is a required field
	TypeName *string `location:"uri" locationName:"typeName" type:"string" required:"true"`
}

// String returns the string representation
func (s ListResolversInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResolversInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResolversInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.TypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TypeName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListResolversInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TypeName != nil {
		v := *s.TypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "typeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListResolversOutput struct {
	_ struct{} `type:"structure"`

	// An identifier to be passed in the next request to this operation to return
	// the next set of items in the list.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The Resolver objects.
	Resolvers []Resolver `locationName:"resolvers" type:"list"`
}

// String returns the string representation
func (s ListResolversOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListResolversOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Resolvers != nil {
		v := s.Resolvers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resolvers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListResolvers = "ListResolvers"

// ListResolversRequest returns a request value for making API operation for
// AWS AppSync.
//
// Lists the resolvers for a given API and type.
//
//    // Example sending a request using ListResolversRequest.
//    req := client.ListResolversRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/ListResolvers
func (c *Client) ListResolversRequest(input *ListResolversInput) ListResolversRequest {
	op := &aws.Operation{
		Name:       opListResolvers,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis/{apiId}/types/{typeName}/resolvers",
	}

	if input == nil {
		input = &ListResolversInput{}
	}

	req := c.newRequest(op, input, &ListResolversOutput{})
	return ListResolversRequest{Request: req, Input: input, Copy: c.ListResolversRequest}
}

// ListResolversRequest is the request type for the
// ListResolvers API operation.
type ListResolversRequest struct {
	*aws.Request
	Input *ListResolversInput
	Copy  func(*ListResolversInput) ListResolversRequest
}

// Send marshals and sends the ListResolvers API request.
func (r ListResolversRequest) Send(ctx context.Context) (*ListResolversResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResolversResponse{
		ListResolversOutput: r.Request.Data.(*ListResolversOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListResolversResponse is the response type for the
// ListResolvers API operation.
type ListResolversResponse struct {
	*ListResolversOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResolvers request.
func (r *ListResolversResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
