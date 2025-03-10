// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ReEncryptInput struct {
	_ struct{} `type:"structure"`

	// Ciphertext of the data to reencrypt.
	//
	// CiphertextBlob is automatically base64 encoded/decoded by the SDK.
	//
	// CiphertextBlob is a required field
	CiphertextBlob []byte `min:"1" type:"blob" required:"true"`

	// Encryption context to use when the data is reencrypted.
	DestinationEncryptionContext map[string]string `type:"map"`

	// A unique identifier for the CMK that is used to reencrypt the data.
	//
	// To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name,
	// or alias ARN. When using an alias name, prefix it with "alias/". To specify
	// a CMK in a different AWS account, you must use the key ARN or alias ARN.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Alias name: alias/ExampleAlias
	//
	//    * Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey. To
	// get the alias name and alias ARN, use ListAliases.
	//
	// DestinationKeyId is a required field
	DestinationKeyId *string `min:"1" type:"string" required:"true"`

	// A list of grant tokens.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string `type:"list"`

	// Encryption context used to encrypt and decrypt the data specified in the
	// CiphertextBlob parameter.
	SourceEncryptionContext map[string]string `type:"map"`
}

// String returns the string representation
func (s ReEncryptInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReEncryptInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReEncryptInput"}

	if s.CiphertextBlob == nil {
		invalidParams.Add(aws.NewErrParamRequired("CiphertextBlob"))
	}
	if s.CiphertextBlob != nil && len(s.CiphertextBlob) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CiphertextBlob", 1))
	}

	if s.DestinationKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationKeyId"))
	}
	if s.DestinationKeyId != nil && len(*s.DestinationKeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationKeyId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ReEncryptOutput struct {
	_ struct{} `type:"structure"`

	// The reencrypted data. When you use the HTTP API or the AWS CLI, the value
	// is Base64-encoded. Otherwise, it is not encoded.
	//
	// CiphertextBlob is automatically base64 encoded/decoded by the SDK.
	CiphertextBlob []byte `min:"1" type:"blob"`

	// Unique identifier of the CMK used to reencrypt the data.
	KeyId *string `min:"1" type:"string"`

	// Unique identifier of the CMK used to originally encrypt the data.
	SourceKeyId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ReEncryptOutput) String() string {
	return awsutil.Prettify(s)
}

const opReEncrypt = "ReEncrypt"

// ReEncryptRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Encrypts data on the server side with a new customer master key (CMK) without
// exposing the plaintext of the data on the client side. The data is first
// decrypted and then reencrypted. You can also use this operation to change
// the encryption context of a ciphertext.
//
// You can reencrypt data using CMKs in different AWS accounts.
//
// Unlike other operations, ReEncrypt is authorized twice, once as ReEncryptFrom
// on the source CMK and once as ReEncryptTo on the destination CMK. We recommend
// that you include the "kms:ReEncrypt*" permission in your key policies (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)
// to permit reencryption from or to the CMK. This permission is automatically
// included in the key policy when you create a CMK through the console. But
// you must include it manually when you create a CMK programmatically or when
// you set a key policy with the PutKeyPolicy operation.
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using ReEncryptRequest.
//    req := client.ReEncryptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ReEncrypt
func (c *Client) ReEncryptRequest(input *ReEncryptInput) ReEncryptRequest {
	op := &aws.Operation{
		Name:       opReEncrypt,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReEncryptInput{}
	}

	req := c.newRequest(op, input, &ReEncryptOutput{})
	return ReEncryptRequest{Request: req, Input: input, Copy: c.ReEncryptRequest}
}

// ReEncryptRequest is the request type for the
// ReEncrypt API operation.
type ReEncryptRequest struct {
	*aws.Request
	Input *ReEncryptInput
	Copy  func(*ReEncryptInput) ReEncryptRequest
}

// Send marshals and sends the ReEncrypt API request.
func (r ReEncryptRequest) Send(ctx context.Context) (*ReEncryptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReEncryptResponse{
		ReEncryptOutput: r.Request.Data.(*ReEncryptOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReEncryptResponse is the response type for the
// ReEncrypt API operation.
type ReEncryptResponse struct {
	*ReEncryptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReEncrypt request.
func (r *ReEncryptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
