// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTaskDefinitionsInput struct {
	_ struct{} `type:"structure"`

	// The full family name with which to filter the ListTaskDefinitions results.
	// Specifying a familyPrefix limits the listed task definitions to task definition
	// revisions that belong to that family.
	FamilyPrefix *string `locationName:"familyPrefix" type:"string"`

	// The maximum number of task definition results returned by ListTaskDefinitions
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitions request with the returned nextToken value. This
	// value can be between 1 and 100. If this parameter is not used, then ListTaskDefinitions
	// returns up to 100 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTaskDefinitions
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The order in which to sort the results. Valid values are ASC and DESC. By
	// default (ASC), task definitions are listed lexicographically by family name
	// and in ascending numerical order by revision so that the newest task definitions
	// in a family are listed last. Setting this parameter to DESC reverses the
	// sort order on family name and revision so that the newest task definitions
	// in a family are listed first.
	Sort SortOrder `locationName:"sort" type:"string" enum:"true"`

	// The task definition status with which to filter the ListTaskDefinitions results.
	// By default, only ACTIVE task definitions are listed. By setting this parameter
	// to INACTIVE, you can view task definitions that are INACTIVE as long as an
	// active task or service still references them. If you paginate the resulting
	// output, be sure to keep the status value constant in each subsequent request.
	Status TaskDefinitionStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListTaskDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

type ListTaskDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListTaskDefinitions request. When
	// the results of a ListTaskDefinitions request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of task definition Amazon Resource Name (ARN) entries for the ListTaskDefinitions
	// request.
	TaskDefinitionArns []string `locationName:"taskDefinitionArns" type:"list"`
}

// String returns the string representation
func (s ListTaskDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTaskDefinitions = "ListTaskDefinitions"

// ListTaskDefinitionsRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Returns a list of task definitions that are registered to your account. You
// can filter the results by family name with the familyPrefix parameter or
// by status with the status parameter.
//
//    // Example sending a request using ListTaskDefinitionsRequest.
//    req := client.ListTaskDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/ListTaskDefinitions
func (c *Client) ListTaskDefinitionsRequest(input *ListTaskDefinitionsInput) ListTaskDefinitionsRequest {
	op := &aws.Operation{
		Name:       opListTaskDefinitions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTaskDefinitionsInput{}
	}

	req := c.newRequest(op, input, &ListTaskDefinitionsOutput{})
	return ListTaskDefinitionsRequest{Request: req, Input: input, Copy: c.ListTaskDefinitionsRequest}
}

// ListTaskDefinitionsRequest is the request type for the
// ListTaskDefinitions API operation.
type ListTaskDefinitionsRequest struct {
	*aws.Request
	Input *ListTaskDefinitionsInput
	Copy  func(*ListTaskDefinitionsInput) ListTaskDefinitionsRequest
}

// Send marshals and sends the ListTaskDefinitions API request.
func (r ListTaskDefinitionsRequest) Send(ctx context.Context) (*ListTaskDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTaskDefinitionsResponse{
		ListTaskDefinitionsOutput: r.Request.Data.(*ListTaskDefinitionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTaskDefinitionsRequestPaginator returns a paginator for ListTaskDefinitions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTaskDefinitionsRequest(input)
//   p := ecs.NewListTaskDefinitionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTaskDefinitionsPaginator(req ListTaskDefinitionsRequest) ListTaskDefinitionsPaginator {
	return ListTaskDefinitionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTaskDefinitionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListTaskDefinitionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTaskDefinitionsPaginator struct {
	aws.Pager
}

func (p *ListTaskDefinitionsPaginator) CurrentPage() *ListTaskDefinitionsOutput {
	return p.Pager.CurrentPage().(*ListTaskDefinitionsOutput)
}

// ListTaskDefinitionsResponse is the response type for the
// ListTaskDefinitions API operation.
type ListTaskDefinitionsResponse struct {
	*ListTaskDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTaskDefinitions request.
func (r *ListTaskDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
