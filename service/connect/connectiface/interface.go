// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package connectiface provides an interface to enable mocking the Amazon Connect Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package connectiface

import (
	"github.com/aws/aws-sdk-go-v2/service/connect"
)

// ClientAPI provides an interface to enable mocking the
// connect.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Connect.
//    func myFunc(svc connectiface.ClientAPI) bool {
//        // Make svc.CreateUser request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := connect.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        connectiface.ClientPI
//    }
//    func (m *mockClientClient) CreateUser(input *connect.CreateUserInput) (*connect.CreateUserOutput, error) {
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
	CreateUserRequest(*connect.CreateUserInput) connect.CreateUserRequest

	DeleteUserRequest(*connect.DeleteUserInput) connect.DeleteUserRequest

	DescribeUserRequest(*connect.DescribeUserInput) connect.DescribeUserRequest

	DescribeUserHierarchyGroupRequest(*connect.DescribeUserHierarchyGroupInput) connect.DescribeUserHierarchyGroupRequest

	DescribeUserHierarchyStructureRequest(*connect.DescribeUserHierarchyStructureInput) connect.DescribeUserHierarchyStructureRequest

	GetContactAttributesRequest(*connect.GetContactAttributesInput) connect.GetContactAttributesRequest

	GetCurrentMetricDataRequest(*connect.GetCurrentMetricDataInput) connect.GetCurrentMetricDataRequest

	GetFederationTokenRequest(*connect.GetFederationTokenInput) connect.GetFederationTokenRequest

	GetMetricDataRequest(*connect.GetMetricDataInput) connect.GetMetricDataRequest

	ListContactFlowsRequest(*connect.ListContactFlowsInput) connect.ListContactFlowsRequest

	ListHoursOfOperationsRequest(*connect.ListHoursOfOperationsInput) connect.ListHoursOfOperationsRequest

	ListPhoneNumbersRequest(*connect.ListPhoneNumbersInput) connect.ListPhoneNumbersRequest

	ListQueuesRequest(*connect.ListQueuesInput) connect.ListQueuesRequest

	ListRoutingProfilesRequest(*connect.ListRoutingProfilesInput) connect.ListRoutingProfilesRequest

	ListSecurityProfilesRequest(*connect.ListSecurityProfilesInput) connect.ListSecurityProfilesRequest

	ListUserHierarchyGroupsRequest(*connect.ListUserHierarchyGroupsInput) connect.ListUserHierarchyGroupsRequest

	ListUsersRequest(*connect.ListUsersInput) connect.ListUsersRequest

	StartOutboundVoiceContactRequest(*connect.StartOutboundVoiceContactInput) connect.StartOutboundVoiceContactRequest

	StopContactRequest(*connect.StopContactInput) connect.StopContactRequest

	UpdateContactAttributesRequest(*connect.UpdateContactAttributesInput) connect.UpdateContactAttributesRequest

	UpdateUserHierarchyRequest(*connect.UpdateUserHierarchyInput) connect.UpdateUserHierarchyRequest

	UpdateUserIdentityInfoRequest(*connect.UpdateUserIdentityInfoInput) connect.UpdateUserIdentityInfoRequest

	UpdateUserPhoneConfigRequest(*connect.UpdateUserPhoneConfigInput) connect.UpdateUserPhoneConfigRequest

	UpdateUserRoutingProfileRequest(*connect.UpdateUserRoutingProfileInput) connect.UpdateUserRoutingProfileRequest

	UpdateUserSecurityProfilesRequest(*connect.UpdateUserSecurityProfilesInput) connect.UpdateUserSecurityProfilesRequest
}

var _ ClientAPI = (*connect.Client)(nil)
