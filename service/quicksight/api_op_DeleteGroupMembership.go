// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteGroupMembershipInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the group is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The name of the group that you want to delete the user from.
	//
	// GroupName is a required field
	GroupName *string `location:"uri" locationName:"GroupName" min:"1" type:"string" required:"true"`

	// The name of the user that you want to delete from the group membership.
	//
	// MemberName is a required field
	MemberName *string `location:"uri" locationName:"MemberName" min:"1" type:"string" required:"true"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGroupMembershipInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGroupMembershipInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGroupMembershipInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if s.MemberName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberName"))
	}
	if s.MemberName != nil && len(*s.MemberName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberName", 1))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGroupMembershipInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupName != nil {
		v := *s.GroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MemberName != nil {
		v := *s.MemberName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "MemberName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteGroupMembershipOutput struct {
	_ struct{} `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s DeleteGroupMembershipOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGroupMembershipOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDeleteGroupMembership = "DeleteGroupMembership"

// DeleteGroupMembershipRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Removes a user from a group so that the user is no longer a member of the
// group.
//
// The permissions resource is arn:aws:quicksight:us-east-1:<aws-account-id>:group/default/<group-name> .
//
// The condition resource is the user name.
//
// The condition key is quicksight:UserName.
//
// CLI Sample:
//
// aws quicksight delete-group-membership --aws-account-id=111122223333 --namespace=default
// --group-name=Sales-Management --member-name=Charlie
//
//    // Example sending a request using DeleteGroupMembershipRequest.
//    req := client.DeleteGroupMembershipRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DeleteGroupMembership
func (c *Client) DeleteGroupMembershipRequest(input *DeleteGroupMembershipInput) DeleteGroupMembershipRequest {
	op := &aws.Operation{
		Name:       opDeleteGroupMembership,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/groups/{GroupName}/members/{MemberName}",
	}

	if input == nil {
		input = &DeleteGroupMembershipInput{}
	}

	req := c.newRequest(op, input, &DeleteGroupMembershipOutput{})
	return DeleteGroupMembershipRequest{Request: req, Input: input, Copy: c.DeleteGroupMembershipRequest}
}

// DeleteGroupMembershipRequest is the request type for the
// DeleteGroupMembership API operation.
type DeleteGroupMembershipRequest struct {
	*aws.Request
	Input *DeleteGroupMembershipInput
	Copy  func(*DeleteGroupMembershipInput) DeleteGroupMembershipRequest
}

// Send marshals and sends the DeleteGroupMembership API request.
func (r DeleteGroupMembershipRequest) Send(ctx context.Context) (*DeleteGroupMembershipResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGroupMembershipResponse{
		DeleteGroupMembershipOutput: r.Request.Data.(*DeleteGroupMembershipOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGroupMembershipResponse is the response type for the
// DeleteGroupMembership API operation.
type DeleteGroupMembershipResponse struct {
	*DeleteGroupMembershipOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGroupMembership request.
func (r *DeleteGroupMembershipResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
