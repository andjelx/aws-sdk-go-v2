// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetCommentsForPullRequestInput struct {
	_ struct{} `type:"structure"`

	// The full commit ID of the commit in the source branch that was the tip of
	// the branch at the time the comment was made.
	AfterCommitId *string `locationName:"afterCommitId" type:"string"`

	// The full commit ID of the commit in the destination branch that was the tip
	// of the branch at the time the pull request was created.
	BeforeCommitId *string `locationName:"beforeCommitId" type:"string"`

	// A non-negative integer used to limit the number of returned results. The
	// default is 100 comments. You can return up to 500 comments with a single
	// request.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// An enumeration token that when provided in a request, returns the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The system-generated ID of the pull request. To get this ID, use ListPullRequests.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The name of the repository that contains the pull request.
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string"`
}

// String returns the string representation
func (s GetCommentsForPullRequestInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCommentsForPullRequestInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCommentsForPullRequestInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetCommentsForPullRequestOutput struct {
	_ struct{} `type:"structure"`

	// An array of comment objects on the pull request.
	CommentsForPullRequestData []CommentsForPullRequest `locationName:"commentsForPullRequestData" type:"list"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetCommentsForPullRequestOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCommentsForPullRequest = "GetCommentsForPullRequest"

// GetCommentsForPullRequestRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns comments made on a pull request.
//
//    // Example sending a request using GetCommentsForPullRequestRequest.
//    req := client.GetCommentsForPullRequestRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetCommentsForPullRequest
func (c *Client) GetCommentsForPullRequestRequest(input *GetCommentsForPullRequestInput) GetCommentsForPullRequestRequest {
	op := &aws.Operation{
		Name:       opGetCommentsForPullRequest,
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
		input = &GetCommentsForPullRequestInput{}
	}

	req := c.newRequest(op, input, &GetCommentsForPullRequestOutput{})
	return GetCommentsForPullRequestRequest{Request: req, Input: input, Copy: c.GetCommentsForPullRequestRequest}
}

// GetCommentsForPullRequestRequest is the request type for the
// GetCommentsForPullRequest API operation.
type GetCommentsForPullRequestRequest struct {
	*aws.Request
	Input *GetCommentsForPullRequestInput
	Copy  func(*GetCommentsForPullRequestInput) GetCommentsForPullRequestRequest
}

// Send marshals and sends the GetCommentsForPullRequest API request.
func (r GetCommentsForPullRequestRequest) Send(ctx context.Context) (*GetCommentsForPullRequestResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCommentsForPullRequestResponse{
		GetCommentsForPullRequestOutput: r.Request.Data.(*GetCommentsForPullRequestOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetCommentsForPullRequestRequestPaginator returns a paginator for GetCommentsForPullRequest.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetCommentsForPullRequestRequest(input)
//   p := codecommit.NewGetCommentsForPullRequestRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetCommentsForPullRequestPaginator(req GetCommentsForPullRequestRequest) GetCommentsForPullRequestPaginator {
	return GetCommentsForPullRequestPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetCommentsForPullRequestInput
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

// GetCommentsForPullRequestPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetCommentsForPullRequestPaginator struct {
	aws.Pager
}

func (p *GetCommentsForPullRequestPaginator) CurrentPage() *GetCommentsForPullRequestOutput {
	return p.Pager.CurrentPage().(*GetCommentsForPullRequestOutput)
}

// GetCommentsForPullRequestResponse is the response type for the
// GetCommentsForPullRequest API operation.
type GetCommentsForPullRequestResponse struct {
	*GetCommentsForPullRequestOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCommentsForPullRequest request.
func (r *GetCommentsForPullRequestResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
