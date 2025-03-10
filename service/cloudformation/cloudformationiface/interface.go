// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudformationiface provides an interface to enable mocking the AWS CloudFormation service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudformationiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

// ClientAPI provides an interface to enable mocking the
// cloudformation.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CloudFormation.
//    func myFunc(svc cloudformationiface.ClientAPI) bool {
//        // Make svc.CancelUpdateStack request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cloudformation.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        cloudformationiface.ClientPI
//    }
//    func (m *mockClientClient) CancelUpdateStack(input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CancelUpdateStackRequest(*cloudformation.CancelUpdateStackInput) cloudformation.CancelUpdateStackRequest

	ContinueUpdateRollbackRequest(*cloudformation.ContinueUpdateRollbackInput) cloudformation.ContinueUpdateRollbackRequest

	CreateChangeSetRequest(*cloudformation.CreateChangeSetInput) cloudformation.CreateChangeSetRequest

	CreateStackRequest(*cloudformation.CreateStackInput) cloudformation.CreateStackRequest

	CreateStackInstancesRequest(*cloudformation.CreateStackInstancesInput) cloudformation.CreateStackInstancesRequest

	CreateStackSetRequest(*cloudformation.CreateStackSetInput) cloudformation.CreateStackSetRequest

	DeleteChangeSetRequest(*cloudformation.DeleteChangeSetInput) cloudformation.DeleteChangeSetRequest

	DeleteStackRequest(*cloudformation.DeleteStackInput) cloudformation.DeleteStackRequest

	DeleteStackInstancesRequest(*cloudformation.DeleteStackInstancesInput) cloudformation.DeleteStackInstancesRequest

	DeleteStackSetRequest(*cloudformation.DeleteStackSetInput) cloudformation.DeleteStackSetRequest

	DescribeAccountLimitsRequest(*cloudformation.DescribeAccountLimitsInput) cloudformation.DescribeAccountLimitsRequest

	DescribeChangeSetRequest(*cloudformation.DescribeChangeSetInput) cloudformation.DescribeChangeSetRequest

	DescribeStackDriftDetectionStatusRequest(*cloudformation.DescribeStackDriftDetectionStatusInput) cloudformation.DescribeStackDriftDetectionStatusRequest

	DescribeStackEventsRequest(*cloudformation.DescribeStackEventsInput) cloudformation.DescribeStackEventsRequest

	DescribeStackInstanceRequest(*cloudformation.DescribeStackInstanceInput) cloudformation.DescribeStackInstanceRequest

	DescribeStackResourceRequest(*cloudformation.DescribeStackResourceInput) cloudformation.DescribeStackResourceRequest

	DescribeStackResourceDriftsRequest(*cloudformation.DescribeStackResourceDriftsInput) cloudformation.DescribeStackResourceDriftsRequest

	DescribeStackResourcesRequest(*cloudformation.DescribeStackResourcesInput) cloudformation.DescribeStackResourcesRequest

	DescribeStackSetRequest(*cloudformation.DescribeStackSetInput) cloudformation.DescribeStackSetRequest

	DescribeStackSetOperationRequest(*cloudformation.DescribeStackSetOperationInput) cloudformation.DescribeStackSetOperationRequest

	DescribeStacksRequest(*cloudformation.DescribeStacksInput) cloudformation.DescribeStacksRequest

	DetectStackDriftRequest(*cloudformation.DetectStackDriftInput) cloudformation.DetectStackDriftRequest

	DetectStackResourceDriftRequest(*cloudformation.DetectStackResourceDriftInput) cloudformation.DetectStackResourceDriftRequest

	EstimateTemplateCostRequest(*cloudformation.EstimateTemplateCostInput) cloudformation.EstimateTemplateCostRequest

	ExecuteChangeSetRequest(*cloudformation.ExecuteChangeSetInput) cloudformation.ExecuteChangeSetRequest

	GetStackPolicyRequest(*cloudformation.GetStackPolicyInput) cloudformation.GetStackPolicyRequest

	GetTemplateRequest(*cloudformation.GetTemplateInput) cloudformation.GetTemplateRequest

	GetTemplateSummaryRequest(*cloudformation.GetTemplateSummaryInput) cloudformation.GetTemplateSummaryRequest

	ListChangeSetsRequest(*cloudformation.ListChangeSetsInput) cloudformation.ListChangeSetsRequest

	ListExportsRequest(*cloudformation.ListExportsInput) cloudformation.ListExportsRequest

	ListImportsRequest(*cloudformation.ListImportsInput) cloudformation.ListImportsRequest

	ListStackInstancesRequest(*cloudformation.ListStackInstancesInput) cloudformation.ListStackInstancesRequest

	ListStackResourcesRequest(*cloudformation.ListStackResourcesInput) cloudformation.ListStackResourcesRequest

	ListStackSetOperationResultsRequest(*cloudformation.ListStackSetOperationResultsInput) cloudformation.ListStackSetOperationResultsRequest

	ListStackSetOperationsRequest(*cloudformation.ListStackSetOperationsInput) cloudformation.ListStackSetOperationsRequest

	ListStackSetsRequest(*cloudformation.ListStackSetsInput) cloudformation.ListStackSetsRequest

	ListStacksRequest(*cloudformation.ListStacksInput) cloudformation.ListStacksRequest

	SetStackPolicyRequest(*cloudformation.SetStackPolicyInput) cloudformation.SetStackPolicyRequest

	SignalResourceRequest(*cloudformation.SignalResourceInput) cloudformation.SignalResourceRequest

	StopStackSetOperationRequest(*cloudformation.StopStackSetOperationInput) cloudformation.StopStackSetOperationRequest

	UpdateStackRequest(*cloudformation.UpdateStackInput) cloudformation.UpdateStackRequest

	UpdateStackInstancesRequest(*cloudformation.UpdateStackInstancesInput) cloudformation.UpdateStackInstancesRequest

	UpdateStackSetRequest(*cloudformation.UpdateStackSetInput) cloudformation.UpdateStackSetRequest

	UpdateTerminationProtectionRequest(*cloudformation.UpdateTerminationProtectionInput) cloudformation.UpdateTerminationProtectionRequest

	ValidateTemplateRequest(*cloudformation.ValidateTemplateInput) cloudformation.ValidateTemplateRequest

	WaitUntilChangeSetCreateComplete(context.Context, *cloudformation.DescribeChangeSetInput, ...aws.WaiterOption) error

	WaitUntilStackCreateComplete(context.Context, *cloudformation.DescribeStacksInput, ...aws.WaiterOption) error

	WaitUntilStackDeleteComplete(context.Context, *cloudformation.DescribeStacksInput, ...aws.WaiterOption) error

	WaitUntilStackExists(context.Context, *cloudformation.DescribeStacksInput, ...aws.WaiterOption) error

	WaitUntilStackImportComplete(context.Context, *cloudformation.DescribeStacksInput, ...aws.WaiterOption) error

	WaitUntilStackUpdateComplete(context.Context, *cloudformation.DescribeStacksInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*cloudformation.Client)(nil)
