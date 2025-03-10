// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateResourceDataSyncInput struct {
	_ struct{} `type:"structure"`

	// Amazon S3 configuration details for the sync.
	//
	// S3Destination is a required field
	S3Destination *ResourceDataSyncS3Destination `type:"structure" required:"true"`

	// A name for the configuration.
	//
	// SyncName is a required field
	SyncName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateResourceDataSyncInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResourceDataSyncInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateResourceDataSyncInput"}

	if s.S3Destination == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3Destination"))
	}

	if s.SyncName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SyncName"))
	}
	if s.SyncName != nil && len(*s.SyncName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SyncName", 1))
	}
	if s.S3Destination != nil {
		if err := s.S3Destination.Validate(); err != nil {
			invalidParams.AddNested("S3Destination", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateResourceDataSyncOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateResourceDataSyncOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateResourceDataSync = "CreateResourceDataSync"

// CreateResourceDataSyncRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Creates a resource data sync configuration to a single bucket in Amazon S3.
// This is an asynchronous operation that returns immediately. After a successful
// initial sync is completed, the system continuously syncs data to the Amazon
// S3 bucket. To check the status of the sync, use the ListResourceDataSync.
//
// By default, data is not encrypted in Amazon S3. We strongly recommend that
// you enable encryption in Amazon S3 to ensure secure data storage. We also
// recommend that you secure access to the Amazon S3 bucket by creating a restrictive
// bucket policy. For more information, see Configuring Resource Data Sync for
// Inventory (http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-inventory-datasync.html)
// in the AWS Systems Manager User Guide.
//
//    // Example sending a request using CreateResourceDataSyncRequest.
//    req := client.CreateResourceDataSyncRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/CreateResourceDataSync
func (c *Client) CreateResourceDataSyncRequest(input *CreateResourceDataSyncInput) CreateResourceDataSyncRequest {
	op := &aws.Operation{
		Name:       opCreateResourceDataSync,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateResourceDataSyncInput{}
	}

	req := c.newRequest(op, input, &CreateResourceDataSyncOutput{})
	return CreateResourceDataSyncRequest{Request: req, Input: input, Copy: c.CreateResourceDataSyncRequest}
}

// CreateResourceDataSyncRequest is the request type for the
// CreateResourceDataSync API operation.
type CreateResourceDataSyncRequest struct {
	*aws.Request
	Input *CreateResourceDataSyncInput
	Copy  func(*CreateResourceDataSyncInput) CreateResourceDataSyncRequest
}

// Send marshals and sends the CreateResourceDataSync API request.
func (r CreateResourceDataSyncRequest) Send(ctx context.Context) (*CreateResourceDataSyncResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateResourceDataSyncResponse{
		CreateResourceDataSyncOutput: r.Request.Data.(*CreateResourceDataSyncOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateResourceDataSyncResponse is the response type for the
// CreateResourceDataSync API operation.
type CreateResourceDataSyncResponse struct {
	*CreateResourceDataSyncOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateResourceDataSync request.
func (r *CreateResourceDataSyncResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
