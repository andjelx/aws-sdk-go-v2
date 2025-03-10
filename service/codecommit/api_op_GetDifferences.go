// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDifferencesInput struct {
	_ struct{} `type:"structure"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit.
	//
	// AfterCommitSpecifier is a required field
	AfterCommitSpecifier *string `locationName:"afterCommitSpecifier" type:"string" required:"true"`

	// The file path in which to check differences. Limits the results to this path.
	// Can also be used to specify the changed name of a directory or folder, if
	// it has changed. If not specified, differences will be shown for all paths.
	AfterPath *string `locationName:"afterPath" type:"string"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit. For example, the full commit ID. Optional. If not specified, all
	// changes prior to the afterCommitSpecifier value will be shown. If you do
	// not use beforeCommitSpecifier in your request, consider limiting the results
	// with maxResults.
	BeforeCommitSpecifier *string `locationName:"beforeCommitSpecifier" type:"string"`

	// The file path in which to check for differences. Limits the results to this
	// path. Can also be used to specify the previous name of a directory or folder.
	// If beforePath and afterPath are not specified, differences will be shown
	// for all paths.
	BeforePath *string `locationName:"beforePath" type:"string"`

	// A non-negative integer used to limit the number of returned results.
	MaxResults *int64 `type:"integer"`

	// An enumeration token that when provided in a request, returns the next batch
	// of the results.
	NextToken *string `type:"string"`

	// The name of the repository where you want to get differences.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDifferencesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDifferencesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDifferencesInput"}

	if s.AfterCommitSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("AfterCommitSpecifier"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDifferencesOutput struct {
	_ struct{} `type:"structure"`

	// A differences data type object that contains information about the differences,
	// including whether the difference is added, modified, or deleted (A, D, M).
	Differences []Difference `locationName:"differences" type:"list"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetDifferencesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDifferences = "GetDifferences"

// GetDifferencesRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about the differences in a valid commit specifier (such
// as a branch, tag, HEAD, commit ID or other fully qualified reference). Results
// can be limited to a specified path.
//
//    // Example sending a request using GetDifferencesRequest.
//    req := client.GetDifferencesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetDifferences
func (c *Client) GetDifferencesRequest(input *GetDifferencesInput) GetDifferencesRequest {
	op := &aws.Operation{
		Name:       opGetDifferences,
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
		input = &GetDifferencesInput{}
	}

	req := c.newRequest(op, input, &GetDifferencesOutput{})
	return GetDifferencesRequest{Request: req, Input: input, Copy: c.GetDifferencesRequest}
}

// GetDifferencesRequest is the request type for the
// GetDifferences API operation.
type GetDifferencesRequest struct {
	*aws.Request
	Input *GetDifferencesInput
	Copy  func(*GetDifferencesInput) GetDifferencesRequest
}

// Send marshals and sends the GetDifferences API request.
func (r GetDifferencesRequest) Send(ctx context.Context) (*GetDifferencesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDifferencesResponse{
		GetDifferencesOutput: r.Request.Data.(*GetDifferencesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetDifferencesRequestPaginator returns a paginator for GetDifferences.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetDifferencesRequest(input)
//   p := codecommit.NewGetDifferencesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetDifferencesPaginator(req GetDifferencesRequest) GetDifferencesPaginator {
	return GetDifferencesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetDifferencesInput
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

// GetDifferencesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetDifferencesPaginator struct {
	aws.Pager
}

func (p *GetDifferencesPaginator) CurrentPage() *GetDifferencesOutput {
	return p.Pager.CurrentPage().(*GetDifferencesOutput)
}

// GetDifferencesResponse is the response type for the
// GetDifferences API operation.
type GetDifferencesResponse struct {
	*GetDifferencesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDifferences request.
func (r *GetDifferencesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
