// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListDataSourcesInput struct {
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
}

// String returns the string representation
func (s ListDataSourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDataSourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDataSourcesInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDataSourcesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

type ListDataSourcesOutput struct {
	_ struct{} `type:"structure"`

	// The DataSource objects.
	DataSources []DataSource `locationName:"dataSources" type:"list"`

	// An identifier to be passed in the next request to this operation to return
	// the next set of items in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDataSourcesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDataSourcesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DataSources != nil {
		v := s.DataSources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dataSources", metadata)
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

const opListDataSources = "ListDataSources"

// ListDataSourcesRequest returns a request value for making API operation for
// AWS AppSync.
//
// Lists the data sources for a given API.
//
//    // Example sending a request using ListDataSourcesRequest.
//    req := client.ListDataSourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/ListDataSources
func (c *Client) ListDataSourcesRequest(input *ListDataSourcesInput) ListDataSourcesRequest {
	op := &aws.Operation{
		Name:       opListDataSources,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis/{apiId}/datasources",
	}

	if input == nil {
		input = &ListDataSourcesInput{}
	}

	req := c.newRequest(op, input, &ListDataSourcesOutput{})
	return ListDataSourcesRequest{Request: req, Input: input, Copy: c.ListDataSourcesRequest}
}

// ListDataSourcesRequest is the request type for the
// ListDataSources API operation.
type ListDataSourcesRequest struct {
	*aws.Request
	Input *ListDataSourcesInput
	Copy  func(*ListDataSourcesInput) ListDataSourcesRequest
}

// Send marshals and sends the ListDataSources API request.
func (r ListDataSourcesRequest) Send(ctx context.Context) (*ListDataSourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDataSourcesResponse{
		ListDataSourcesOutput: r.Request.Data.(*ListDataSourcesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDataSourcesResponse is the response type for the
// ListDataSources API operation.
type ListDataSourcesResponse struct {
	*ListDataSourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDataSources request.
func (r *ListDataSourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
