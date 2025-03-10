// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListProgressUpdateStreamsInput struct {
	_ struct{} `type:"structure"`

	// Filter to limit the maximum number of results to list per page.
	MaxResults *int64 `min:"1" type:"integer"`

	// If a NextToken was returned by a previous call, there are more results available.
	// To retrieve the next page of results, make the call again using the returned
	// token in NextToken.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListProgressUpdateStreamsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProgressUpdateStreamsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProgressUpdateStreamsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListProgressUpdateStreamsOutput struct {
	_ struct{} `type:"structure"`

	// If there are more streams created than the max result, return the next token
	// to be passed to the next call as a bookmark of where to start from.
	NextToken *string `type:"string"`

	// List of progress update streams up to the max number of results passed in
	// the input.
	ProgressUpdateStreamSummaryList []ProgressUpdateStreamSummary `type:"list"`
}

// String returns the string representation
func (s ListProgressUpdateStreamsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListProgressUpdateStreams = "ListProgressUpdateStreams"

// ListProgressUpdateStreamsRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Lists progress update streams associated with the user account making this
// call.
//
//    // Example sending a request using ListProgressUpdateStreamsRequest.
//    req := client.ListProgressUpdateStreamsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/ListProgressUpdateStreams
func (c *Client) ListProgressUpdateStreamsRequest(input *ListProgressUpdateStreamsInput) ListProgressUpdateStreamsRequest {
	op := &aws.Operation{
		Name:       opListProgressUpdateStreams,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListProgressUpdateStreamsInput{}
	}

	req := c.newRequest(op, input, &ListProgressUpdateStreamsOutput{})
	return ListProgressUpdateStreamsRequest{Request: req, Input: input, Copy: c.ListProgressUpdateStreamsRequest}
}

// ListProgressUpdateStreamsRequest is the request type for the
// ListProgressUpdateStreams API operation.
type ListProgressUpdateStreamsRequest struct {
	*aws.Request
	Input *ListProgressUpdateStreamsInput
	Copy  func(*ListProgressUpdateStreamsInput) ListProgressUpdateStreamsRequest
}

// Send marshals and sends the ListProgressUpdateStreams API request.
func (r ListProgressUpdateStreamsRequest) Send(ctx context.Context) (*ListProgressUpdateStreamsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProgressUpdateStreamsResponse{
		ListProgressUpdateStreamsOutput: r.Request.Data.(*ListProgressUpdateStreamsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListProgressUpdateStreamsResponse is the response type for the
// ListProgressUpdateStreams API operation.
type ListProgressUpdateStreamsResponse struct {
	*ListProgressUpdateStreamsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProgressUpdateStreams request.
func (r *ListProgressUpdateStreamsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
