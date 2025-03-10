// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDatasetImportJobsInput struct {
	_ struct{} `type:"structure"`

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT, which specifies whether
	// to include or exclude, respectively, from the list, the predictors that match
	// the statement. The match statement consists of a key and a value. In this
	// release, Name is the only valid key, which filters on the DatasetImportJobName
	// property.
	//
	//    * Condition - IS or IS_NOT
	//
	//    * Key - Name
	//
	//    * Value - the value to match
	//
	// For example, to list all dataset import jobs named my_dataset_import_job,
	// you would specify:
	//
	// "Filters": [ { "Condition": "IS", "Key": "Name", "Value": "my_dataset_import_job"
	// } ]
	Filters []Filter `type:"list"`

	// The number of items to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the result of the previous request was truncated, the response includes
	// a NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDatasetImportJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDatasetImportJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDatasetImportJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDatasetImportJobsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that summarize each dataset import job's properties.
	DatasetImportJobs []DatasetImportJobSummary `type:"list"`

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDatasetImportJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDatasetImportJobs = "ListDatasetImportJobs"

// ListDatasetImportJobsRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Returns a list of dataset import jobs created using the CreateDatasetImportJob
// operation. For each import job, a summary of its properties, including its
// Amazon Resource Name (ARN), is returned. You can retrieve the complete set
// of properties by using the ARN with the DescribeDatasetImportJob operation.
// You can filter the list by providing an array of Filter objects.
//
//    // Example sending a request using ListDatasetImportJobsRequest.
//    req := client.ListDatasetImportJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/ListDatasetImportJobs
func (c *Client) ListDatasetImportJobsRequest(input *ListDatasetImportJobsInput) ListDatasetImportJobsRequest {
	op := &aws.Operation{
		Name:       opListDatasetImportJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDatasetImportJobsInput{}
	}

	req := c.newRequest(op, input, &ListDatasetImportJobsOutput{})
	return ListDatasetImportJobsRequest{Request: req, Input: input, Copy: c.ListDatasetImportJobsRequest}
}

// ListDatasetImportJobsRequest is the request type for the
// ListDatasetImportJobs API operation.
type ListDatasetImportJobsRequest struct {
	*aws.Request
	Input *ListDatasetImportJobsInput
	Copy  func(*ListDatasetImportJobsInput) ListDatasetImportJobsRequest
}

// Send marshals and sends the ListDatasetImportJobs API request.
func (r ListDatasetImportJobsRequest) Send(ctx context.Context) (*ListDatasetImportJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDatasetImportJobsResponse{
		ListDatasetImportJobsOutput: r.Request.Data.(*ListDatasetImportJobsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDatasetImportJobsRequestPaginator returns a paginator for ListDatasetImportJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDatasetImportJobsRequest(input)
//   p := forecast.NewListDatasetImportJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDatasetImportJobsPaginator(req ListDatasetImportJobsRequest) ListDatasetImportJobsPaginator {
	return ListDatasetImportJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDatasetImportJobsInput
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

// ListDatasetImportJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDatasetImportJobsPaginator struct {
	aws.Pager
}

func (p *ListDatasetImportJobsPaginator) CurrentPage() *ListDatasetImportJobsOutput {
	return p.Pager.CurrentPage().(*ListDatasetImportJobsOutput)
}

// ListDatasetImportJobsResponse is the response type for the
// ListDatasetImportJobs API operation.
type ListDatasetImportJobsResponse struct {
	*ListDatasetImportJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDatasetImportJobs request.
func (r *ListDatasetImportJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
