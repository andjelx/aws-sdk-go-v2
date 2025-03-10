// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListBucketMetricsConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket containing the metrics configurations to retrieve.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The marker that is used to continue a metrics configuration listing that
	// has been truncated. Use the NextContinuationToken from a previously truncated
	// list response to continue the listing. The continuation token is an opaque
	// value that Amazon S3 understands.
	ContinuationToken *string `location:"querystring" locationName:"continuation-token" type:"string"`
}

// String returns the string representation
func (s ListBucketMetricsConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBucketMetricsConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBucketMetricsConfigurationsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *ListBucketMetricsConfigurationsInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBucketMetricsConfigurationsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.ContinuationToken != nil {
		v := *s.ContinuationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "continuation-token", protocol.StringValue(v), metadata)
	}
	return nil
}

type ListBucketMetricsConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// The marker that is used as a starting point for this metrics configuration
	// list response. This value is present if it was sent in the request.
	ContinuationToken *string `type:"string"`

	// Indicates whether the returned list of metrics configurations is complete.
	// A value of true indicates that the list is not complete and the NextContinuationToken
	// will be provided for a subsequent request.
	IsTruncated *bool `type:"boolean"`

	// The list of metrics configurations for a bucket.
	MetricsConfigurationList []MetricsConfiguration `locationName:"MetricsConfiguration" type:"list" flattened:"true"`

	// The marker used to continue a metrics configuration listing that has been
	// truncated. Use the NextContinuationToken from a previously truncated list
	// response to continue the listing. The continuation token is an opaque value
	// that Amazon S3 understands.
	NextContinuationToken *string `type:"string"`
}

// String returns the string representation
func (s ListBucketMetricsConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBucketMetricsConfigurationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContinuationToken != nil {
		v := *s.ContinuationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ContinuationToken", protocol.StringValue(v), metadata)
	}
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.MetricsConfigurationList != nil {
		v := s.MetricsConfigurationList

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "MetricsConfiguration", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextContinuationToken != nil {
		v := *s.NextContinuationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextContinuationToken", protocol.StringValue(v), metadata)
	}
	return nil
}

const opListBucketMetricsConfigurations = "ListBucketMetricsConfigurations"

// ListBucketMetricsConfigurationsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Lists the metrics configurations for the bucket.
//
//    // Example sending a request using ListBucketMetricsConfigurationsRequest.
//    req := client.ListBucketMetricsConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListBucketMetricsConfigurations
func (c *Client) ListBucketMetricsConfigurationsRequest(input *ListBucketMetricsConfigurationsInput) ListBucketMetricsConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListBucketMetricsConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?metrics",
	}

	if input == nil {
		input = &ListBucketMetricsConfigurationsInput{}
	}

	req := c.newRequest(op, input, &ListBucketMetricsConfigurationsOutput{})
	return ListBucketMetricsConfigurationsRequest{Request: req, Input: input, Copy: c.ListBucketMetricsConfigurationsRequest}
}

// ListBucketMetricsConfigurationsRequest is the request type for the
// ListBucketMetricsConfigurations API operation.
type ListBucketMetricsConfigurationsRequest struct {
	*aws.Request
	Input *ListBucketMetricsConfigurationsInput
	Copy  func(*ListBucketMetricsConfigurationsInput) ListBucketMetricsConfigurationsRequest
}

// Send marshals and sends the ListBucketMetricsConfigurations API request.
func (r ListBucketMetricsConfigurationsRequest) Send(ctx context.Context) (*ListBucketMetricsConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBucketMetricsConfigurationsResponse{
		ListBucketMetricsConfigurationsOutput: r.Request.Data.(*ListBucketMetricsConfigurationsOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListBucketMetricsConfigurationsResponse is the response type for the
// ListBucketMetricsConfigurations API operation.
type ListBucketMetricsConfigurationsResponse struct {
	*ListBucketMetricsConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBucketMetricsConfigurations request.
func (r *ListBucketMetricsConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
