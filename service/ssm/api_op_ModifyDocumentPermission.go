// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyDocumentPermissionInput struct {
	_ struct{} `type:"structure"`

	// The AWS user accounts that should have access to the document. The account
	// IDs can either be a group of account IDs or All.
	AccountIdsToAdd []string `type:"list"`

	// The AWS user accounts that should no longer have access to the document.
	// The AWS user account can either be a group of account IDs or All. This action
	// has a higher priority than AccountIdsToAdd. If you specify an account ID
	// to add and the same ID to remove, the system removes access to the document.
	AccountIdsToRemove []string `type:"list"`

	// The name of the document that you want to share.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The permission type for the document. The permission type can be Share.
	//
	// PermissionType is a required field
	PermissionType DocumentPermissionType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ModifyDocumentPermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDocumentPermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDocumentPermissionInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if len(s.PermissionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PermissionType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyDocumentPermissionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyDocumentPermissionOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyDocumentPermission = "ModifyDocumentPermission"

// ModifyDocumentPermissionRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Shares a Systems Manager document publicly or privately. If you share a document
// privately, you must specify the AWS user account IDs for those people who
// can use the document. If you share a document publicly, you must specify
// All as the account ID.
//
//    // Example sending a request using ModifyDocumentPermissionRequest.
//    req := client.ModifyDocumentPermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ModifyDocumentPermission
func (c *Client) ModifyDocumentPermissionRequest(input *ModifyDocumentPermissionInput) ModifyDocumentPermissionRequest {
	op := &aws.Operation{
		Name:       opModifyDocumentPermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDocumentPermissionInput{}
	}

	req := c.newRequest(op, input, &ModifyDocumentPermissionOutput{})
	return ModifyDocumentPermissionRequest{Request: req, Input: input, Copy: c.ModifyDocumentPermissionRequest}
}

// ModifyDocumentPermissionRequest is the request type for the
// ModifyDocumentPermission API operation.
type ModifyDocumentPermissionRequest struct {
	*aws.Request
	Input *ModifyDocumentPermissionInput
	Copy  func(*ModifyDocumentPermissionInput) ModifyDocumentPermissionRequest
}

// Send marshals and sends the ModifyDocumentPermission API request.
func (r ModifyDocumentPermissionRequest) Send(ctx context.Context) (*ModifyDocumentPermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDocumentPermissionResponse{
		ModifyDocumentPermissionOutput: r.Request.Data.(*ModifyDocumentPermissionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDocumentPermissionResponse is the response type for the
// ModifyDocumentPermission API operation.
type ModifyDocumentPermissionResponse struct {
	*ModifyDocumentPermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDocumentPermission request.
func (r *ModifyDocumentPermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
