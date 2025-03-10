// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateGrantInput struct {
	_ struct{} `type:"structure"`

	// Allows a cryptographic operation only when the encryption context matches
	// or includes the encryption context specified in this structure. For more
	// information about encryption context, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide .
	Constraints *GrantConstraints `type:"structure"`

	// A list of grant tokens.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string `type:"list"`

	// The principal that is given permission to perform the operations that the
	// grant permits.
	//
	// To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of an AWS principal. Valid AWS principals include AWS accounts (root), IAM
	// users, IAM roles, federated users, and assumed role users. For examples of
	// the ARN syntax to use for specifying a principal, see AWS Identity and Access
	// Management (IAM) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)
	// in the Example ARNs section of the AWS General Reference.
	//
	// GranteePrincipal is a required field
	GranteePrincipal *string `min:"1" type:"string" required:"true"`

	// The unique identifier for the customer master key (CMK) that the grant applies
	// to.
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify
	// a CMK in a different AWS account, you must use the key ARN.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// A friendly name for identifying the grant. Use this value to prevent the
	// unintended creation of duplicate grants when retrying this request.
	//
	// When this value is absent, all CreateGrant requests result in a new grant
	// with a unique GrantId even if all the supplied parameters are identical.
	// This can result in unintended duplicates when you retry the CreateGrant request.
	//
	// When this value is present, you can retry a CreateGrant request with identical
	// parameters; if the grant already exists, the original GrantId is returned
	// without creating a new grant. Note that the returned grant token is unique
	// with every CreateGrant request, even when a duplicate GrantId is returned.
	// All grant tokens obtained in this way can be used interchangeably.
	Name *string `min:"1" type:"string"`

	// A list of operations that the grant permits.
	//
	// Operations is a required field
	Operations []GrantOperation `type:"list" required:"true"`

	// The principal that is given permission to retire the grant by using RetireGrant
	// operation.
	//
	// To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of an AWS principal. Valid AWS principals include AWS accounts (root), IAM
	// users, federated users, and assumed role users. For examples of the ARN syntax
	// to use for specifying a principal, see AWS Identity and Access Management
	// (IAM) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)
	// in the Example ARNs section of the AWS General Reference.
	RetiringPrincipal *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateGrantInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGrantInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGrantInput"}

	if s.GranteePrincipal == nil {
		invalidParams.Add(aws.NewErrParamRequired("GranteePrincipal"))
	}
	if s.GranteePrincipal != nil && len(*s.GranteePrincipal) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GranteePrincipal", 1))
	}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Operations == nil {
		invalidParams.Add(aws.NewErrParamRequired("Operations"))
	}
	if s.RetiringPrincipal != nil && len(*s.RetiringPrincipal) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RetiringPrincipal", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateGrantOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the grant.
	//
	// You can use the GrantId in a subsequent RetireGrant or RevokeGrant operation.
	GrantId *string `min:"1" type:"string"`

	// The grant token.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateGrantOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateGrant = "CreateGrant"

// CreateGrantRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Adds a grant to a customer master key (CMK). The grant allows the grantee
// principal to use the CMK when the conditions specified in the grant are met.
// When setting permissions, grants are an alternative to key policies.
//
// To create a grant that allows a cryptographic operation only when the encryption
// context in the operation request matches or includes a specified encryption
// context, use the Constraints parameter. For details, see GrantConstraints.
//
// To perform this operation on a CMK in a different AWS account, specify the
// key ARN in the value of the KeyId parameter. For more information about grants,
// see Grants (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html)
// in the AWS Key Management Service Developer Guide .
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using CreateGrantRequest.
//    req := client.CreateGrantRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/CreateGrant
func (c *Client) CreateGrantRequest(input *CreateGrantInput) CreateGrantRequest {
	op := &aws.Operation{
		Name:       opCreateGrant,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateGrantInput{}
	}

	req := c.newRequest(op, input, &CreateGrantOutput{})
	return CreateGrantRequest{Request: req, Input: input, Copy: c.CreateGrantRequest}
}

// CreateGrantRequest is the request type for the
// CreateGrant API operation.
type CreateGrantRequest struct {
	*aws.Request
	Input *CreateGrantInput
	Copy  func(*CreateGrantInput) CreateGrantRequest
}

// Send marshals and sends the CreateGrant API request.
func (r CreateGrantRequest) Send(ctx context.Context) (*CreateGrantResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGrantResponse{
		CreateGrantOutput: r.Request.Data.(*CreateGrantOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGrantResponse is the response type for the
// CreateGrant API operation.
type CreateGrantResponse struct {
	*CreateGrantOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGrant request.
func (r *CreateGrantResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
