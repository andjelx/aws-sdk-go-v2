// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDetectSyntaxInput struct {
	_ struct{} `type:"structure"`

	// The language of the input documents. You can specify any of the primary languages
	// supported by Amazon Comprehend: German ("de"), English ("en"), Spanish ("es"),
	// French ("fr"), Italian ("it"), or Portuguese ("pt"). All documents must be
	// in the same language.
	//
	// LanguageCode is a required field
	LanguageCode SyntaxLanguageCode `type:"string" required:"true" enum:"true"`

	// A list containing the text of the input documents. The list can contain a
	// maximum of 25 documents. Each document must contain fewer that 5,000 bytes
	// of UTF-8 encoded characters.
	//
	// TextList is a required field
	TextList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectSyntaxInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDetectSyntaxInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDetectSyntaxInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.TextList == nil {
		invalidParams.Add(aws.NewErrParamRequired("TextList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDetectSyntaxOutput struct {
	_ struct{} `type:"structure"`

	// A list containing one object for each document that contained an error. The
	// results are sorted in ascending order by the Index field and match the order
	// of the documents in the input list. If there are no errors in the batch,
	// the ErrorList is empty.
	//
	// ErrorList is a required field
	ErrorList []BatchItemError `type:"list" required:"true"`

	// A list of objects containing the results of the operation. The results are
	// sorted in ascending order by the Index field and match the order of the documents
	// in the input list. If all of the documents contain an error, the ResultList
	// is empty.
	//
	// ResultList is a required field
	ResultList []BatchDetectSyntaxItemResult `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectSyntaxOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDetectSyntax = "BatchDetectSyntax"

// BatchDetectSyntaxRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Inspects the text of a batch of documents for the syntax and part of speech
// of the words in the document and returns information about them. For more
// information, see how-syntax.
//
//    // Example sending a request using BatchDetectSyntaxRequest.
//    req := client.BatchDetectSyntaxRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectSyntax
func (c *Client) BatchDetectSyntaxRequest(input *BatchDetectSyntaxInput) BatchDetectSyntaxRequest {
	op := &aws.Operation{
		Name:       opBatchDetectSyntax,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDetectSyntaxInput{}
	}

	req := c.newRequest(op, input, &BatchDetectSyntaxOutput{})
	return BatchDetectSyntaxRequest{Request: req, Input: input, Copy: c.BatchDetectSyntaxRequest}
}

// BatchDetectSyntaxRequest is the request type for the
// BatchDetectSyntax API operation.
type BatchDetectSyntaxRequest struct {
	*aws.Request
	Input *BatchDetectSyntaxInput
	Copy  func(*BatchDetectSyntaxInput) BatchDetectSyntaxRequest
}

// Send marshals and sends the BatchDetectSyntax API request.
func (r BatchDetectSyntaxRequest) Send(ctx context.Context) (*BatchDetectSyntaxResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDetectSyntaxResponse{
		BatchDetectSyntaxOutput: r.Request.Data.(*BatchDetectSyntaxOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDetectSyntaxResponse is the response type for the
// BatchDetectSyntax API operation.
type BatchDetectSyntaxResponse struct {
	*BatchDetectSyntaxOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDetectSyntax request.
func (r *BatchDetectSyntaxResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
