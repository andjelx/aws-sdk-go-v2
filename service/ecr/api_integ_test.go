// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build integration

package ecr_test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

var _ aws.Config
var _ awserr.Error

func TestInteg_00_DescribeRepositories(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-west-2")
	svc := ecr.New(cfg)
	params := &ecr.DescribeRepositoriesInput{}

	req := svc.DescribeRepositoriesRequest(params)

	_, err := req.Send(ctx)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_ListImages(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-west-2")
	svc := ecr.New(cfg)
	params := &ecr.ListImagesInput{
		RepositoryName: aws.String("not-a-real-repository"),
	}

	req := svc.ListImagesRequest(params)

	_, err := req.Send(ctx)
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if len(aerr.Code()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if v := aerr.Code(); v == aws.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}
