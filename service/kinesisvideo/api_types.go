// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// An object describing a Kinesis video stream.
type StreamInfo struct {
	_ struct{} `type:"structure"`

	// A time stamp that indicates when the stream was created.
	CreationTime *time.Time `type:"timestamp"`

	// How long the stream retains data, in hours.
	DataRetentionInHours *int64 `type:"integer"`

	// The name of the device that is associated with the stream.
	DeviceName *string `min:"1" type:"string"`

	// The ID of the AWS Key Management Service (AWS KMS) key that Kinesis Video
	// Streams uses to encrypt data on the stream.
	KmsKeyId *string `min:"1" type:"string"`

	// The MediaType of the stream.
	MediaType *string `min:"1" type:"string"`

	// The status of the stream.
	Status Status `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the stream.
	StreamARN *string `min:"1" type:"string"`

	// The name of the stream.
	StreamName *string `min:"1" type:"string"`

	// The version of the stream.
	Version *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StreamInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StreamInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.DataRetentionInHours != nil {
		v := *s.DataRetentionInHours

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataRetentionInHours", protocol.Int64Value(v), metadata)
	}
	if s.DeviceName != nil {
		v := *s.DeviceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeviceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KmsKeyId != nil {
		v := *s.KmsKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KmsKeyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MediaType != nil {
		v := *s.MediaType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MediaType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StreamARN != nil {
		v := *s.StreamARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamName != nil {
		v := *s.StreamName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Specifies the condition that streams must satisfy to be returned when you
// list streams (see the ListStreams API). A condition has a comparison operation
// and a value. Currently, you can specify only the BEGINS_WITH operator, which
// finds streams whose names start with a given prefix.
type StreamNameCondition struct {
	_ struct{} `type:"structure"`

	// A comparison operator. Currently, you can specify only the BEGINS_WITH operator,
	// which finds streams whose names start with a given prefix.
	ComparisonOperator ComparisonOperator `type:"string" enum:"true"`

	// A value to compare.
	ComparisonValue *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StreamNameCondition) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StreamNameCondition) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StreamNameCondition"}
	if s.ComparisonValue != nil && len(*s.ComparisonValue) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ComparisonValue", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StreamNameCondition) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ComparisonOperator) > 0 {
		v := s.ComparisonOperator

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ComparisonOperator", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ComparisonValue != nil {
		v := *s.ComparisonValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ComparisonValue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
