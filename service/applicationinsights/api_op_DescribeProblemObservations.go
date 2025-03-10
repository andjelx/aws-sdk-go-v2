// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeProblemObservationsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the problem.
	//
	// ProblemId is a required field
	ProblemId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeProblemObservationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProblemObservationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProblemObservationsInput"}

	if s.ProblemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProblemId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeProblemObservationsOutput struct {
	_ struct{} `type:"structure"`

	// Observations related to the problem.
	RelatedObservations *RelatedObservations `type:"structure"`
}

// String returns the string representation
func (s DescribeProblemObservationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeProblemObservations = "DescribeProblemObservations"

// DescribeProblemObservationsRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Describes the anomalies or errors associated with the problem.
//
//    // Example sending a request using DescribeProblemObservationsRequest.
//    req := client.DescribeProblemObservationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/DescribeProblemObservations
func (c *Client) DescribeProblemObservationsRequest(input *DescribeProblemObservationsInput) DescribeProblemObservationsRequest {
	op := &aws.Operation{
		Name:       opDescribeProblemObservations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeProblemObservationsInput{}
	}

	req := c.newRequest(op, input, &DescribeProblemObservationsOutput{})
	return DescribeProblemObservationsRequest{Request: req, Input: input, Copy: c.DescribeProblemObservationsRequest}
}

// DescribeProblemObservationsRequest is the request type for the
// DescribeProblemObservations API operation.
type DescribeProblemObservationsRequest struct {
	*aws.Request
	Input *DescribeProblemObservationsInput
	Copy  func(*DescribeProblemObservationsInput) DescribeProblemObservationsRequest
}

// Send marshals and sends the DescribeProblemObservations API request.
func (r DescribeProblemObservationsRequest) Send(ctx context.Context) (*DescribeProblemObservationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProblemObservationsResponse{
		DescribeProblemObservationsOutput: r.Request.Data.(*DescribeProblemObservationsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProblemObservationsResponse is the response type for the
// DescribeProblemObservations API operation.
type DescribeProblemObservationsResponse struct {
	*DescribeProblemObservationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProblemObservations request.
func (r *DescribeProblemObservationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
