// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListGraphqlApisInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results you want the request to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListGraphqlApisInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGraphqlApisInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

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

type ListGraphqlApisOutput struct {
	_ struct{} `type:"structure"`

	// The GraphqlApi objects.
	GraphqlApis []GraphqlApi `locationName:"graphqlApis" type:"list"`

	// An identifier to be passed in the next request to this operation to return
	// the next set of items in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListGraphqlApisOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGraphqlApisOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GraphqlApis != nil {
		v := s.GraphqlApis

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "graphqlApis", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListGraphqlApis = "ListGraphqlApis"

// ListGraphqlApisRequest returns a request value for making API operation for
// AWS AppSync.
//
// Lists your GraphQL APIs.
//
//    // Example sending a request using ListGraphqlApisRequest.
//    req := client.ListGraphqlApisRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/ListGraphqlApis
func (c *Client) ListGraphqlApisRequest(input *ListGraphqlApisInput) ListGraphqlApisRequest {
	op := &aws.Operation{
		Name:       opListGraphqlApis,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis",
	}

	if input == nil {
		input = &ListGraphqlApisInput{}
	}

	req := c.newRequest(op, input, &ListGraphqlApisOutput{})
	return ListGraphqlApisRequest{Request: req, Input: input, Copy: c.ListGraphqlApisRequest}
}

// ListGraphqlApisRequest is the request type for the
// ListGraphqlApis API operation.
type ListGraphqlApisRequest struct {
	*aws.Request
	Input *ListGraphqlApisInput
	Copy  func(*ListGraphqlApisInput) ListGraphqlApisRequest
}

// Send marshals and sends the ListGraphqlApis API request.
func (r ListGraphqlApisRequest) Send(ctx context.Context) (*ListGraphqlApisResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGraphqlApisResponse{
		ListGraphqlApisOutput: r.Request.Data.(*ListGraphqlApisOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGraphqlApisResponse is the response type for the
// ListGraphqlApis API operation.
type ListGraphqlApisResponse struct {
	*ListGraphqlApisOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGraphqlApis request.
func (r *ListGraphqlApisResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
