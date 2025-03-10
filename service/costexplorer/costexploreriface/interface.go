// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package costexploreriface provides an interface to enable mocking the AWS Cost Explorer Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package costexploreriface

import (
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

// ClientAPI provides an interface to enable mocking the
// costexplorer.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Cost Explorer.
//    func myFunc(svc costexploreriface.ClientAPI) bool {
//        // Make svc.GetCostAndUsage request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := costexplorer.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        costexploreriface.ClientPI
//    }
//    func (m *mockClientClient) GetCostAndUsage(input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error) {
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
	GetCostAndUsageRequest(*costexplorer.GetCostAndUsageInput) costexplorer.GetCostAndUsageRequest

	GetCostAndUsageWithResourcesRequest(*costexplorer.GetCostAndUsageWithResourcesInput) costexplorer.GetCostAndUsageWithResourcesRequest

	GetCostForecastRequest(*costexplorer.GetCostForecastInput) costexplorer.GetCostForecastRequest

	GetDimensionValuesRequest(*costexplorer.GetDimensionValuesInput) costexplorer.GetDimensionValuesRequest

	GetReservationCoverageRequest(*costexplorer.GetReservationCoverageInput) costexplorer.GetReservationCoverageRequest

	GetReservationPurchaseRecommendationRequest(*costexplorer.GetReservationPurchaseRecommendationInput) costexplorer.GetReservationPurchaseRecommendationRequest

	GetReservationUtilizationRequest(*costexplorer.GetReservationUtilizationInput) costexplorer.GetReservationUtilizationRequest

	GetRightsizingRecommendationRequest(*costexplorer.GetRightsizingRecommendationInput) costexplorer.GetRightsizingRecommendationRequest

	GetSavingsPlansCoverageRequest(*costexplorer.GetSavingsPlansCoverageInput) costexplorer.GetSavingsPlansCoverageRequest

	GetSavingsPlansPurchaseRecommendationRequest(*costexplorer.GetSavingsPlansPurchaseRecommendationInput) costexplorer.GetSavingsPlansPurchaseRecommendationRequest

	GetSavingsPlansUtilizationRequest(*costexplorer.GetSavingsPlansUtilizationInput) costexplorer.GetSavingsPlansUtilizationRequest

	GetSavingsPlansUtilizationDetailsRequest(*costexplorer.GetSavingsPlansUtilizationDetailsInput) costexplorer.GetSavingsPlansUtilizationDetailsRequest

	GetTagsRequest(*costexplorer.GetTagsInput) costexplorer.GetTagsRequest

	GetUsageForecastRequest(*costexplorer.GetUsageForecastInput) costexplorer.GetUsageForecastRequest
}

var _ ClientAPI = (*costexplorer.Client)(nil)
