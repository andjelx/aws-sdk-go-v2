// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDocumentInput struct {
	_ struct{} `type:"structure"`

	// Returns the document in the specified format. The document format can be
	// either JSON or YAML. JSON is the default format.
	DocumentFormat DocumentFormat `type:"string" enum:"true"`

	// The document version for which you want information.
	DocumentVersion *string `type:"string"`

	// The name of the Systems Manager document.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// An optional field specifying the version of the artifact associated with
	// the document. For example, "Release 12, Update 6". This value is unique across
	// all versions of a document, and cannot be changed.
	VersionName *string `type:"string"`
}

// String returns the string representation
func (s GetDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDocumentInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDocumentOutput struct {
	_ struct{} `type:"structure"`

	// A description of the document attachments, including names, locations, sizes,
	// etc.
	AttachmentsContent []AttachmentContent `type:"list"`

	// The contents of the Systems Manager document.
	Content *string `min:"1" type:"string"`

	// The document format, either JSON or YAML.
	DocumentFormat DocumentFormat `type:"string" enum:"true"`

	// The document type.
	DocumentType DocumentType `type:"string" enum:"true"`

	// The document version.
	DocumentVersion *string `type:"string"`

	// The name of the Systems Manager document.
	Name *string `type:"string"`

	// The status of the Systems Manager document, such as Creating, Active, Updating,
	// Failed, and Deleting.
	Status DocumentStatus `type:"string" enum:"true"`

	// A message returned by AWS Systems Manager that explains the Status value.
	// For example, a Failed status might be explained by the StatusInformation
	// message, "The specified S3 bucket does not exist. Verify that the URL of
	// the S3 bucket is correct."
	StatusInformation *string `type:"string"`

	// The version of the artifact associated with the document. For example, "Release
	// 12, Update 6". This value is unique across all versions of a document, and
	// cannot be changed.
	VersionName *string `type:"string"`
}

// String returns the string representation
func (s GetDocumentOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDocument = "GetDocument"

// GetDocumentRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Gets the contents of the specified Systems Manager document.
//
//    // Example sending a request using GetDocumentRequest.
//    req := client.GetDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetDocument
func (c *Client) GetDocumentRequest(input *GetDocumentInput) GetDocumentRequest {
	op := &aws.Operation{
		Name:       opGetDocument,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDocumentInput{}
	}

	req := c.newRequest(op, input, &GetDocumentOutput{})
	return GetDocumentRequest{Request: req, Input: input, Copy: c.GetDocumentRequest}
}

// GetDocumentRequest is the request type for the
// GetDocument API operation.
type GetDocumentRequest struct {
	*aws.Request
	Input *GetDocumentInput
	Copy  func(*GetDocumentInput) GetDocumentRequest
}

// Send marshals and sends the GetDocument API request.
func (r GetDocumentRequest) Send(ctx context.Context) (*GetDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDocumentResponse{
		GetDocumentOutput: r.Request.Data.(*GetDocumentOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDocumentResponse is the response type for the
// GetDocument API operation.
type GetDocumentResponse struct {
	*GetDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDocument request.
func (r *GetDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
