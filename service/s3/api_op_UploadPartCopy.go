// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UploadPartCopyInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The name of the source bucket and key name of the source object, separated
	// by a slash (/). Must be URL-encoded.
	//
	// CopySource is a required field
	CopySource *string `location:"header" locationName:"x-amz-copy-source" type:"string" required:"true"`

	// Copies the object if its entity tag (ETag) matches the specified tag.
	CopySourceIfMatch *string `location:"header" locationName:"x-amz-copy-source-if-match" type:"string"`

	// Copies the object if it has been modified since the specified time.
	CopySourceIfModifiedSince *time.Time `location:"header" locationName:"x-amz-copy-source-if-modified-since" type:"timestamp"`

	// Copies the object if its entity tag (ETag) is different than the specified
	// ETag.
	CopySourceIfNoneMatch *string `location:"header" locationName:"x-amz-copy-source-if-none-match" type:"string"`

	// Copies the object if it hasn't been modified since the specified time.
	CopySourceIfUnmodifiedSince *time.Time `location:"header" locationName:"x-amz-copy-source-if-unmodified-since" type:"timestamp"`

	// The range of bytes to copy from the source object. The range value must use
	// the form bytes=first-last, where the first and last are the zero-based byte
	// offsets to copy. For example, bytes=0-9 indicates that you want to copy the
	// first ten bytes of the source. You can copy a range only if the source object
	// is greater than 5 MB.
	CopySourceRange *string `location:"header" locationName:"x-amz-copy-source-range" type:"string"`

	// Specifies the algorithm to use when decrypting the source object (e.g., AES256).
	CopySourceSSECustomerAlgorithm *string `location:"header" locationName:"x-amz-copy-source-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt
	// the source object. The encryption key provided in this header must be one
	// that was used when the source object was created.
	CopySourceSSECustomerKey *string `location:"header" locationName:"x-amz-copy-source-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	CopySourceSSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-copy-source-server-side-encryption-customer-key-MD5" type:"string"`

	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Part number of part being copied. This is a positive integer between 1 and
	// 10,000.
	//
	// PartNumber is a required field
	PartNumber *int64 `location:"querystring" locationName:"partNumber" type:"integer" required:"true"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// does not store the encryption key. The key must be appropriate for use with
	// the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	// header. This must be the same encryption key specified in the initiate multipart
	// upload request.
	SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// Upload ID identifying the multipart upload whose part is being copied.
	//
	// UploadId is a required field
	UploadId *string `location:"querystring" locationName:"uploadId" type:"string" required:"true"`
}

// String returns the string representation
func (s UploadPartCopyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadPartCopyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UploadPartCopyInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.CopySource == nil {
		invalidParams.Add(aws.NewErrParamRequired("CopySource"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.PartNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartNumber"))
	}

	if s.UploadId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UploadId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *UploadPartCopyInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

func (s *UploadPartCopyInput) getCopySourceSSECustomerKey() (v string) {
	if s.CopySourceSSECustomerKey == nil {
		return v
	}
	return *s.CopySourceSSECustomerKey
}

func (s *UploadPartCopyInput) getSSECustomerKey() (v string) {
	if s.SSECustomerKey == nil {
		return v
	}
	return *s.SSECustomerKey
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UploadPartCopyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.CopySource != nil {
		v := *s.CopySource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source", protocol.StringValue(v), metadata)
	}
	if s.CopySourceIfMatch != nil {
		v := *s.CopySourceIfMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-if-match", protocol.StringValue(v), metadata)
	}
	if s.CopySourceIfModifiedSince != nil {
		v := *s.CopySourceIfModifiedSince

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-if-modified-since",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.CopySourceIfNoneMatch != nil {
		v := *s.CopySourceIfNoneMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-if-none-match", protocol.StringValue(v), metadata)
	}
	if s.CopySourceIfUnmodifiedSince != nil {
		v := *s.CopySourceIfUnmodifiedSince

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-if-unmodified-since",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.CopySourceRange != nil {
		v := *s.CopySourceRange

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-range", protocol.StringValue(v), metadata)
	}
	if s.CopySourceSSECustomerAlgorithm != nil {
		v := *s.CopySourceSSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.CopySourceSSECustomerKey != nil {
		v := *s.CopySourceSSECustomerKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-server-side-encryption-customer-key", protocol.StringValue(v), metadata)
	}
	if s.CopySourceSSECustomerKeyMD5 != nil {
		v := *s.CopySourceSSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKey != nil {
		v := *s.SSECustomerKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.PartNumber != nil {
		v := *s.PartNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "partNumber", protocol.Int64Value(v), metadata)
	}
	if s.UploadId != nil {
		v := *s.UploadId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "uploadId", protocol.StringValue(v), metadata)
	}
	return nil
}

type UploadPartCopyOutput struct {
	_ struct{} `type:"structure" payload:"CopyPartResult"`

	CopyPartResult *CopyPartResult `type:"structure"`

	// The version of the source object that was copied, if you have enabled versioning
	// on the source bucket.
	CopySourceVersionId *string `location:"header" locationName:"x-amz-copy-source-version-id" type:"string"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the ID of the AWS Key Management Service (KMS) master
	// encryption key that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`
}

// String returns the string representation
func (s UploadPartCopyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UploadPartCopyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CopySourceVersionId != nil {
		v := *s.CopySourceVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-version-id", protocol.StringValue(v), metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	if s.CopyPartResult != nil {
		v := s.CopyPartResult

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CopyPartResult", v, metadata)
	}
	return nil
}

const opUploadPartCopy = "UploadPartCopy"

// UploadPartCopyRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Uploads a part by copying data from an existing object as data source.
//
//    // Example sending a request using UploadPartCopyRequest.
//    req := client.UploadPartCopyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/UploadPartCopy
func (c *Client) UploadPartCopyRequest(input *UploadPartCopyInput) UploadPartCopyRequest {
	op := &aws.Operation{
		Name:       opUploadPartCopy,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &UploadPartCopyInput{}
	}

	req := c.newRequest(op, input, &UploadPartCopyOutput{})
	return UploadPartCopyRequest{Request: req, Input: input, Copy: c.UploadPartCopyRequest}
}

// UploadPartCopyRequest is the request type for the
// UploadPartCopy API operation.
type UploadPartCopyRequest struct {
	*aws.Request
	Input *UploadPartCopyInput
	Copy  func(*UploadPartCopyInput) UploadPartCopyRequest
}

// Send marshals and sends the UploadPartCopy API request.
func (r UploadPartCopyRequest) Send(ctx context.Context) (*UploadPartCopyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UploadPartCopyResponse{
		UploadPartCopyOutput: r.Request.Data.(*UploadPartCopyOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UploadPartCopyResponse is the response type for the
// UploadPartCopy API operation.
type UploadPartCopyResponse struct {
	*UploadPartCopyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UploadPartCopy request.
func (r *UploadPartCopyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
