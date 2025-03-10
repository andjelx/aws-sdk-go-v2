// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// An object representing the certificate-authority-data for your cluster.
type Certificate struct {
	_ struct{} `type:"structure"`

	// The Base64-encoded certificate data required to communicate with your cluster.
	// Add this to the certificate-authority-data section of the kubeconfig file
	// for your cluster.
	Data *string `locationName:"data" type:"string"`
}

// String returns the string representation
func (s Certificate) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Certificate) MarshalFields(e protocol.FieldEncoder) error {
	if s.Data != nil {
		v := *s.Data

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "data", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object representing an Amazon EKS cluster.
type Cluster struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the cluster.
	Arn *string `locationName:"arn" type:"string"`

	// The certificate-authority-data for your cluster.
	CertificateAuthority *Certificate `locationName:"certificateAuthority" type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `locationName:"clientRequestToken" type:"string"`

	// The Unix epoch timestamp in seconds for when the cluster was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The endpoint for your Kubernetes API server.
	Endpoint *string `locationName:"endpoint" type:"string"`

	// The identity provider information for the cluster.
	Identity *Identity `locationName:"identity" type:"structure"`

	// The logging configuration for your cluster.
	Logging *Logging `locationName:"logging" type:"structure"`

	// The name of the cluster.
	Name *string `locationName:"name" type:"string"`

	// The platform version of your Amazon EKS cluster. For more information, see
	// Platform Versions (https://docs.aws.amazon.com/eks/latest/userguide/platform-versions.html)
	// in the Amazon EKS User Guide .
	PlatformVersion *string `locationName:"platformVersion" type:"string"`

	// The VPC configuration used by the cluster control plane. Amazon EKS VPC resources
	// have specific requirements to work properly with Kubernetes. For more information,
	// see Cluster VPC Considerations (https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html)
	// and Cluster Security Group Considerations (https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html)
	// in the Amazon EKS User Guide.
	ResourcesVpcConfig *VpcConfigResponse `locationName:"resourcesVpcConfig" type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM role that provides permissions
	// for the Kubernetes control plane to make calls to AWS API operations on your
	// behalf.
	RoleArn *string `locationName:"roleArn" type:"string"`

	// The current status of the cluster.
	Status ClusterStatus `locationName:"status" type:"string" enum:"true"`

	// The metadata that you apply to the cluster to assist with categorization
	// and organization. Each tag consists of a key and an optional value, both
	// of which you define.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`

	// The Kubernetes server version for the cluster.
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s Cluster) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Cluster) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateAuthority != nil {
		v := s.CertificateAuthority

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "certificateAuthority", v, metadata)
	}
	if s.ClientRequestToken != nil {
		v := *s.ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Endpoint != nil {
		v := *s.Endpoint

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endpoint", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Identity != nil {
		v := s.Identity

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "identity", v, metadata)
	}
	if s.Logging != nil {
		v := s.Logging

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "logging", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlatformVersion != nil {
		v := *s.PlatformVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "platformVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourcesVpcConfig != nil {
		v := s.ResourcesVpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resourcesVpcConfig", v, metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object representing an error when an asynchronous operation fails.
type ErrorDetail struct {
	_ struct{} `type:"structure"`

	// A brief description of the error.
	//
	//    * SubnetNotFound: We couldn't find one of the subnets associated with
	//    the cluster.
	//
	//    * SecurityGroupNotFound: We couldn't find one of the security groups associated
	//    with the cluster.
	//
	//    * EniLimitReached: You have reached the elastic network interface limit
	//    for your account.
	//
	//    * IpNotAvailable: A subnet associated with the cluster doesn't have any
	//    free IP addresses.
	//
	//    * AccessDenied: You don't have permissions to perform the specified operation.
	//
	//    * OperationNotPermitted: The service role associated with the cluster
	//    doesn't have the required access permissions for Amazon EKS.
	//
	//    * VpcIdNotFound: We couldn't find the VPC associated with the cluster.
	ErrorCode ErrorCode `locationName:"errorCode" type:"string" enum:"true"`

	// A more complete description of the error.
	ErrorMessage *string `locationName:"errorMessage" type:"string"`

	// An optional field that contains the resource IDs associated with the error.
	ResourceIds []string `locationName:"resourceIds" type:"list"`
}

// String returns the string representation
func (s ErrorDetail) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ErrorDetail) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ErrorCode) > 0 {
		v := s.ErrorCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "errorCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ErrorMessage != nil {
		v := *s.ErrorMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "errorMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceIds != nil {
		v := s.ResourceIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// An object representing an identity provider for authentication credentials.
type Identity struct {
	_ struct{} `type:"structure"`

	// The OpenID Connect (https://openid.net/connect/) identity provider information
	// for the cluster.
	Oidc *OIDC `locationName:"oidc" type:"structure"`
}

// String returns the string representation
func (s Identity) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Identity) MarshalFields(e protocol.FieldEncoder) error {
	if s.Oidc != nil {
		v := s.Oidc

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "oidc", v, metadata)
	}
	return nil
}

// An object representing the enabled or disabled Kubernetes control plane logs
// for your cluster.
type LogSetup struct {
	_ struct{} `type:"structure"`

	// If a log type is enabled, that log type exports its control plane logs to
	// CloudWatch Logs. If a log type isn't enabled, that log type doesn't export
	// its control plane logs. Each individual log type can be enabled or disabled
	// independently.
	Enabled *bool `locationName:"enabled" type:"boolean"`

	// The available cluster control plane log types.
	Types []LogType `locationName:"types" type:"list"`
}

// String returns the string representation
func (s LogSetup) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LogSetup) MarshalFields(e protocol.FieldEncoder) error {
	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enabled", protocol.BoolValue(v), metadata)
	}
	if s.Types != nil {
		v := s.Types

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "types", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// An object representing the logging configuration for resources in your cluster.
type Logging struct {
	_ struct{} `type:"structure"`

	// The cluster control plane logging configuration for your cluster.
	ClusterLogging []LogSetup `locationName:"clusterLogging" type:"list"`
}

// String returns the string representation
func (s Logging) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Logging) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClusterLogging != nil {
		v := s.ClusterLogging

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "clusterLogging", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// An object representing the OpenID Connect (https://openid.net/connect/) identity
// provider information for the cluster.
type OIDC struct {
	_ struct{} `type:"structure"`

	// The issuer URL for the OpenID Connect identity provider.
	Issuer *string `locationName:"issuer" type:"string"`
}

// String returns the string representation
func (s OIDC) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OIDC) MarshalFields(e protocol.FieldEncoder) error {
	if s.Issuer != nil {
		v := *s.Issuer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "issuer", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object representing an asynchronous update.
type Update struct {
	_ struct{} `type:"structure"`

	// The Unix epoch timestamp in seconds for when the update was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// Any errors associated with a Failed update.
	Errors []ErrorDetail `locationName:"errors" type:"list"`

	// A UUID that is used to track the update.
	Id *string `locationName:"id" type:"string"`

	// A key-value map that contains the parameters associated with the update.
	Params []UpdateParam `locationName:"params" type:"list"`

	// The current status of the update.
	Status UpdateStatus `locationName:"status" type:"string" enum:"true"`

	// The type of the update.
	Type UpdateType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s Update) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Update) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Errors != nil {
		v := s.Errors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "errors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Params != nil {
		v := s.Params

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "params", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// An object representing the details of an update request.
type UpdateParam struct {
	_ struct{} `type:"structure"`

	// The keys associated with an update request.
	Type UpdateParamType `locationName:"type" type:"string" enum:"true"`

	// The value of the keys submitted as part of an update request.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s UpdateParam) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateParam) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object representing the VPC configuration to use for an Amazon EKS cluster.
type VpcConfigRequest struct {
	_ struct{} `type:"structure"`

	// Set this value to true to enable private access for your cluster's Kubernetes
	// API server endpoint. If you enable private access, Kubernetes API requests
	// from within your cluster's VPC use the private VPC endpoint. The default
	// value for this parameter is false, which disables private access for your
	// Kubernetes API server. For more information, see Amazon EKS Cluster Endpoint
	// Access Control (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html)
	// in the Amazon EKS User Guide .
	EndpointPrivateAccess *bool `locationName:"endpointPrivateAccess" type:"boolean"`

	// Set this value to false to disable public access for your cluster's Kubernetes
	// API server endpoint. If you disable public access, your cluster's Kubernetes
	// API server can receive only requests from within the cluster VPC. The default
	// value for this parameter is true, which enables public access for your Kubernetes
	// API server. For more information, see Amazon EKS Cluster Endpoint Access
	// Control (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html)
	// in the Amazon EKS User Guide .
	EndpointPublicAccess *bool `locationName:"endpointPublicAccess" type:"boolean"`

	// Specify one or more security groups for the cross-account elastic network
	// interfaces that Amazon EKS creates to use to allow communication between
	// your worker nodes and the Kubernetes control plane. If you don't specify
	// a security group, the default security group for your VPC is used.
	SecurityGroupIds []string `locationName:"securityGroupIds" type:"list"`

	// Specify subnets for your Amazon EKS worker nodes. Amazon EKS creates cross-account
	// elastic network interfaces in these subnets to allow communication between
	// your worker nodes and the Kubernetes control plane.
	SubnetIds []string `locationName:"subnetIds" type:"list"`
}

// String returns the string representation
func (s VpcConfigRequest) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s VpcConfigRequest) MarshalFields(e protocol.FieldEncoder) error {
	if s.EndpointPrivateAccess != nil {
		v := *s.EndpointPrivateAccess

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endpointPrivateAccess", protocol.BoolValue(v), metadata)
	}
	if s.EndpointPublicAccess != nil {
		v := *s.EndpointPublicAccess

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endpointPublicAccess", protocol.BoolValue(v), metadata)
	}
	if s.SecurityGroupIds != nil {
		v := s.SecurityGroupIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "securityGroupIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SubnetIds != nil {
		v := s.SubnetIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "subnetIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// An object representing an Amazon EKS cluster VPC configuration response.
type VpcConfigResponse struct {
	_ struct{} `type:"structure"`

	// This parameter indicates whether the Amazon EKS private API server endpoint
	// is enabled. If the Amazon EKS private API server endpoint is enabled, Kubernetes
	// API requests that originate from within your cluster's VPC use the private
	// VPC endpoint instead of traversing the internet.
	EndpointPrivateAccess *bool `locationName:"endpointPrivateAccess" type:"boolean"`

	// This parameter indicates whether the Amazon EKS public API server endpoint
	// is enabled. If the Amazon EKS public API server endpoint is disabled, your
	// cluster's Kubernetes API server can receive only requests that originate
	// from within the cluster VPC.
	EndpointPublicAccess *bool `locationName:"endpointPublicAccess" type:"boolean"`

	// The security groups associated with the cross-account elastic network interfaces
	// that are used to allow communication between your worker nodes and the Kubernetes
	// control plane.
	SecurityGroupIds []string `locationName:"securityGroupIds" type:"list"`

	// The subnets associated with your cluster.
	SubnetIds []string `locationName:"subnetIds" type:"list"`

	// The VPC associated with your cluster.
	VpcId *string `locationName:"vpcId" type:"string"`
}

// String returns the string representation
func (s VpcConfigResponse) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s VpcConfigResponse) MarshalFields(e protocol.FieldEncoder) error {
	if s.EndpointPrivateAccess != nil {
		v := *s.EndpointPrivateAccess

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endpointPrivateAccess", protocol.BoolValue(v), metadata)
	}
	if s.EndpointPublicAccess != nil {
		v := *s.EndpointPublicAccess

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endpointPublicAccess", protocol.BoolValue(v), metadata)
	}
	if s.SecurityGroupIds != nil {
		v := s.SecurityGroupIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "securityGroupIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SubnetIds != nil {
		v := s.SubnetIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "subnetIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.VpcId != nil {
		v := *s.VpcId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "vpcId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
