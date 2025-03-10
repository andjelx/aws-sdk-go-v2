// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeletePredictorInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the predictor to delete.
	//
	// PredictorArn is a required field
	PredictorArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePredictorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePredictorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePredictorInput"}

	if s.PredictorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PredictorArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletePredictorOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePredictorOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeletePredictor = "DeletePredictor"

// DeletePredictorRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Deletes a predictor created using the CreatePredictor operation. To be deleted,
// the predictor must have a status of ACTIVE or CREATE_FAILED. Use the DescribePredictor
// operation to get the status.
//
// Any forecasts generated by the predictor will no longer be available.
//
//    // Example sending a request using DeletePredictorRequest.
//    req := client.DeletePredictorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/DeletePredictor
func (c *Client) DeletePredictorRequest(input *DeletePredictorInput) DeletePredictorRequest {
	op := &aws.Operation{
		Name:       opDeletePredictor,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePredictorInput{}
	}

	req := c.newRequest(op, input, &DeletePredictorOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePredictorRequest{Request: req, Input: input, Copy: c.DeletePredictorRequest}
}

// DeletePredictorRequest is the request type for the
// DeletePredictor API operation.
type DeletePredictorRequest struct {
	*aws.Request
	Input *DeletePredictorInput
	Copy  func(*DeletePredictorInput) DeletePredictorRequest
}

// Send marshals and sends the DeletePredictor API request.
func (r DeletePredictorRequest) Send(ctx context.Context) (*DeletePredictorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePredictorResponse{
		DeletePredictorOutput: r.Request.Data.(*DeletePredictorOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePredictorResponse is the response type for the
// DeletePredictor API operation.
type DeletePredictorResponse struct {
	*DeletePredictorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePredictor request.
func (r *DeletePredictorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
