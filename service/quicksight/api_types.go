// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// A group in Amazon QuickSight consists of a set of users. You can use groups
// to make it easier to manage access and security. Currently, an Amazon QuickSight
// subscription can't contain more than 500 Amazon QuickSight groups.
type Group struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the group.
	Arn *string `type:"string"`

	// The group description.
	Description *string `min:"1" type:"string"`

	// The name of the group.
	GroupName *string `min:"1" type:"string"`

	// The principal ID of the group.
	PrincipalId *string `type:"string"`
}

// String returns the string representation
func (s Group) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Group) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupName != nil {
		v := *s.GroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PrincipalId != nil {
		v := *s.PrincipalId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PrincipalId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A member of an Amazon QuickSight group. Currently, group members must be
// users. Groups can't be members of another group.
type GroupMember struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the group member (user).
	Arn *string `type:"string"`

	// The name of the group member (user).
	MemberName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GroupMember) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GroupMember) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MemberName != nil {
		v := *s.MemberName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MemberName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A registered user of Amazon QuickSight. Currently, an Amazon QuickSight subscription
// can't contain more than 20 million users.
type User struct {
	_ struct{} `type:"structure"`

	// Active status of user. When you create an Amazon QuickSight user that’s
	// not an IAM user or an AD user, that user is inactive until they sign in and
	// provide a password
	Active *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) for the user.
	Arn *string `type:"string"`

	// The user's email address.
	Email *string `type:"string"`

	// The type of identity authentication used by the user.
	IdentityType IdentityType `type:"string" enum:"true"`

	// The principal ID of the user.
	PrincipalId *string `type:"string"`

	// The Amazon QuickSight role for the user.
	Role UserRole `type:"string" enum:"true"`

	// The user's user name.
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s User) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s User) MarshalFields(e protocol.FieldEncoder) error {
	if s.Active != nil {
		v := *s.Active

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Active", protocol.BoolValue(v), metadata)
	}
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Email != nil {
		v := *s.Email

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Email", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.IdentityType) > 0 {
		v := s.IdentityType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IdentityType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.PrincipalId != nil {
		v := *s.PrincipalId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PrincipalId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Role) > 0 {
		v := s.Role

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Role", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.UserName != nil {
		v := *s.UserName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UserName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
