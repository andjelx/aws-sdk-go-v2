// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDatasetGroupInput struct {
	_ struct{} `type:"structure"`

	// An array of Amazon Resource Names (ARNs) of the datasets that you want to
	// include in the dataset group.
	DatasetArns []string `type:"list"`

	// A name for the dataset group.
	//
	// DatasetGroupName is a required field
	DatasetGroupName *string `min:"1" type:"string" required:"true"`

	// The domain associated with the dataset group. The Domain and DatasetType
	// that you choose determine the fields that must be present in the training
	// data that you import to the dataset. For example, if you choose the RETAIL
	// domain and TARGET_TIME_SERIES as the DatasetType, Amazon Forecast requires
	// item_id, timestamp, and demand fields to be present in your data. For more
	// information, see howitworks-datasets-groups.
	//
	// Domain is a required field
	Domain Domain `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateDatasetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDatasetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDatasetGroupInput"}

	if s.DatasetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetGroupName"))
	}
	if s.DatasetGroupName != nil && len(*s.DatasetGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetGroupName", 1))
	}
	if len(s.Domain) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDatasetGroupOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string `type:"string"`
}

// String returns the string representation
func (s CreateDatasetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDatasetGroup = "CreateDatasetGroup"

// CreateDatasetGroupRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Creates an Amazon Forecast dataset group, which holds a collection of related
// datasets. You can add datasets to the dataset group when you create the dataset
// group, or you can add datasets later with the UpdateDatasetGroup operation.
//
// After creating a dataset group and adding datasets, you use the dataset group
// when you create a predictor. For more information, see howitworks-datasets-groups.
//
// To get a list of all your datasets groups, use the ListDatasetGroups operation.
//
// The Status of a dataset group must be ACTIVE before you can create a predictor
// using the dataset group. Use the DescribeDatasetGroup operation to get the
// status.
//
//    // Example sending a request using CreateDatasetGroupRequest.
//    req := client.CreateDatasetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/CreateDatasetGroup
func (c *Client) CreateDatasetGroupRequest(input *CreateDatasetGroupInput) CreateDatasetGroupRequest {
	op := &aws.Operation{
		Name:       opCreateDatasetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDatasetGroupInput{}
	}

	req := c.newRequest(op, input, &CreateDatasetGroupOutput{})
	return CreateDatasetGroupRequest{Request: req, Input: input, Copy: c.CreateDatasetGroupRequest}
}

// CreateDatasetGroupRequest is the request type for the
// CreateDatasetGroup API operation.
type CreateDatasetGroupRequest struct {
	*aws.Request
	Input *CreateDatasetGroupInput
	Copy  func(*CreateDatasetGroupInput) CreateDatasetGroupRequest
}

// Send marshals and sends the CreateDatasetGroup API request.
func (r CreateDatasetGroupRequest) Send(ctx context.Context) (*CreateDatasetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDatasetGroupResponse{
		CreateDatasetGroupOutput: r.Request.Data.(*CreateDatasetGroupOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDatasetGroupResponse is the response type for the
// CreateDatasetGroup API operation.
type CreateDatasetGroupResponse struct {
	*CreateDatasetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDatasetGroup request.
func (r *CreateDatasetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
