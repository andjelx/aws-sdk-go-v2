// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// This exception is thrown if the request contains a semantic error. The precise
	// meaning will depend on the API, and will be documented in the error message.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeException for service response error code
	// "Exception".
	//
	// These errors are usually caused by a client action. Actions can include using
	// an action or resource on behalf of a user that doesn't have permissions to
	// use the action or resource or specifying an identifier that is not valid.
	ErrCodeException = "Exception"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// The specified parameter is invalid. Review the available parameters for the
	// API request.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request is invalid given the state of the cluster. Check the state of
	// the cluster and the associated operations.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// A service resource associated with the request could not be found. Clients
	// should not retry such requests.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The specified resource is in use.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceLimitExceededException for service response error code
	// "ResourceLimitExceededException".
	//
	// You have encountered a service limit on the specified resource.
	ErrCodeResourceLimitExceededException = "ResourceLimitExceededException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource could not be found. You can view your available clusters
	// with ListClusters. Amazon EKS clusters are Region-specific.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServerException for service response error code
	// "ServerException".
	//
	// These errors are usually caused by a server-side issue.
	ErrCodeServerException = "ServerException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is unavailable. Back off and retry the operation.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeUnsupportedAvailabilityZoneException for service response error code
	// "UnsupportedAvailabilityZoneException".
	//
	// At least one of your specified cluster subnets is in an Availability Zone
	// that does not support Amazon EKS. The exception output specifies the supported
	// Availability Zones for your account, from which you can choose subnets for
	// your cluster.
	ErrCodeUnsupportedAvailabilityZoneException = "UnsupportedAvailabilityZoneException"
)
