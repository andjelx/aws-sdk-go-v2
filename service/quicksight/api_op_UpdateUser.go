// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateUserInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the user is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The email address of the user that you want to update.
	//
	// Email is a required field
	Email *string `type:"string" required:"true"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`

	// The Amazon QuickSight role of the user. The user role can be one of the following:
	//
	//    * READER: A user who has read-only access to dashboards.
	//
	//    * AUTHOR: A user who can create data sources, data sets, analyses, and
	//    dashboards.
	//
	//    * ADMIN: A user who is an author, who can also manage Amazon QuickSight
	//    settings.
	//
	// Role is a required field
	Role UserRole `type:"string" required:"true" enum:"true"`

	// The Amazon QuickSight user name that you want to update.
	//
	// UserName is a required field
	UserName *string `location:"uri" locationName:"UserName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.Email == nil {
		invalidParams.Add(aws.NewErrParamRequired("Email"))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}
	if len(s.Role) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Role"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Email != nil {
		v := *s.Email

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Email", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Role) > 0 {
		v := s.Role

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Role", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserName != nil {
		v := *s.UserName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "UserName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateUserOutput struct {
	_ struct{} `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The Amazon QuickSight user.
	User *User `type:"structure"`
}

// String returns the string representation
func (s UpdateUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.User != nil {
		v := s.User

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "User", v, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateUser = "UpdateUser"

// UpdateUserRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates an Amazon QuickSight user.
//
// The permission resource is arn:aws:quicksight:us-east-1:<aws-account-id>:user/default/<user-name> .
//
// The response is a user object that contains the user's Amazon QuickSight
// user name, email address, active or inactive status in Amazon QuickSight,
// Amazon QuickSight role, and Amazon Resource Name (ARN).
//
// CLI Sample:
//
// aws quicksight update-user --user-name=Pat --role=ADMIN --email=new_address@amazon.com
// --aws-account-id=111122223333 --namespace=default --region=us-east-1
//
//    // Example sending a request using UpdateUserRequest.
//    req := client.UpdateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateUser
func (c *Client) UpdateUserRequest(input *UpdateUserInput) UpdateUserRequest {
	op := &aws.Operation{
		Name:       opUpdateUser,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/users/{UserName}",
	}

	if input == nil {
		input = &UpdateUserInput{}
	}

	req := c.newRequest(op, input, &UpdateUserOutput{})
	return UpdateUserRequest{Request: req, Input: input, Copy: c.UpdateUserRequest}
}

// UpdateUserRequest is the request type for the
// UpdateUser API operation.
type UpdateUserRequest struct {
	*aws.Request
	Input *UpdateUserInput
	Copy  func(*UpdateUserInput) UpdateUserRequest
}

// Send marshals and sends the UpdateUser API request.
func (r UpdateUserRequest) Send(ctx context.Context) (*UpdateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserResponse{
		UpdateUserOutput: r.Request.Data.(*UpdateUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserResponse is the response type for the
// UpdateUser API operation.
type UpdateUserResponse struct {
	*UpdateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUser request.
func (r *UpdateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
