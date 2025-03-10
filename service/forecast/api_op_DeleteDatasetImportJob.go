// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteDatasetImportJobInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset import job to delete.
	//
	// DatasetImportJobArn is a required field
	DatasetImportJobArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDatasetImportJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDatasetImportJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDatasetImportJobInput"}

	if s.DatasetImportJobArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetImportJobArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDatasetImportJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDatasetImportJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDatasetImportJob = "DeleteDatasetImportJob"

// DeleteDatasetImportJobRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Deletes a dataset import job created using the CreateDatasetImportJob operation.
// To be deleted, the import job must have a status of ACTIVE or CREATE_FAILED.
// Use the DescribeDatasetImportJob operation to get the status.
//
//    // Example sending a request using DeleteDatasetImportJobRequest.
//    req := client.DeleteDatasetImportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/DeleteDatasetImportJob
func (c *Client) DeleteDatasetImportJobRequest(input *DeleteDatasetImportJobInput) DeleteDatasetImportJobRequest {
	op := &aws.Operation{
		Name:       opDeleteDatasetImportJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDatasetImportJobInput{}
	}

	req := c.newRequest(op, input, &DeleteDatasetImportJobOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDatasetImportJobRequest{Request: req, Input: input, Copy: c.DeleteDatasetImportJobRequest}
}

// DeleteDatasetImportJobRequest is the request type for the
// DeleteDatasetImportJob API operation.
type DeleteDatasetImportJobRequest struct {
	*aws.Request
	Input *DeleteDatasetImportJobInput
	Copy  func(*DeleteDatasetImportJobInput) DeleteDatasetImportJobRequest
}

// Send marshals and sends the DeleteDatasetImportJob API request.
func (r DeleteDatasetImportJobRequest) Send(ctx context.Context) (*DeleteDatasetImportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDatasetImportJobResponse{
		DeleteDatasetImportJobOutput: r.Request.Data.(*DeleteDatasetImportJobOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDatasetImportJobResponse is the response type for the
// DeleteDatasetImportJob API operation.
type DeleteDatasetImportJobResponse struct {
	*DeleteDatasetImportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDatasetImportJob request.
func (r *DeleteDatasetImportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
