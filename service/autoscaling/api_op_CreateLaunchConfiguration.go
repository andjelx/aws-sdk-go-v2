// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type CreateLaunchConfigurationInput struct {
	_ struct{} `type:"structure"`

	// For Auto Scaling groups that are running in a virtual private cloud (VPC),
	// specifies whether to assign a public IP address to the group's instances.
	// If you specify true, each instance in the Auto Scaling group receives a unique
	// public IP address. For more information, see Launching Auto Scaling Instances
	// in a VPC (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	// If you specify this parameter, you must specify at least one subnet for VPCZoneIdentifier
	// when you create your group.
	//
	// If the instance is launched into a default subnet, the default is to assign
	// a public IP address, unless you disabled the option to assign a public IP
	// address on the subnet. If the instance is launched into a nondefault subnet,
	// the default is not to assign a public IP address, unless you enabled the
	// option to assign a public IP address on the subnet.
	AssociatePublicIpAddress *bool `type:"boolean"`

	// A block device mapping, which specifies the block devices for the instance.
	// You can specify virtual devices and EBS volumes. For more information, see
	// Block Device Mapping (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	BlockDeviceMappings []BlockDeviceMapping `type:"list"`

	// The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances to.
	// For more information, see ClassicLink (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html)
	// in the Amazon EC2 User Guide for Linux Instances and Linking EC2-Classic
	// Instances to a VPC (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-ClassicLink)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	// This parameter can only be used if you are launching EC2-Classic instances.
	ClassicLinkVPCId *string `min:"1" type:"string"`

	// The IDs of one or more security groups for the specified ClassicLink-enabled
	// VPC. For more information, see ClassicLink (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html)
	// in the Amazon EC2 User Guide for Linux Instances and Linking EC2-Classic
	// Instances to a VPC (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-ClassicLink)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	// If you specify the ClassicLinkVPCId parameter, you must specify this parameter.
	ClassicLinkVPCSecurityGroups []string `type:"list"`

	// Specifies whether the launch configuration is optimized for EBS I/O (true)
	// or not (false). The optimization provides dedicated throughput to Amazon
	// EBS and an optimized configuration stack to provide optimal I/O performance.
	// This optimization is not available with all instance types. Additional fees
	// are incurred when you enable EBS optimization for an instance type that is
	// not EBS-optimized by default. For more information, see Amazon EBS-Optimized
	// Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	//
	// The default value is false.
	EbsOptimized *bool `type:"boolean"`

	// The name or the Amazon Resource Name (ARN) of the instance profile associated
	// with the IAM role for the instance. The instance profile contains the IAM
	// role.
	//
	// For more information, see IAM Role for Applications That Run on Amazon EC2
	// Instances (https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	IamInstanceProfile *string `min:"1" type:"string"`

	// The ID of the Amazon Machine Image (AMI) that was assigned during registration.
	// For more information, see Finding an AMI (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	//
	// If you do not specify InstanceId, you must specify ImageId.
	ImageId *string `min:"1" type:"string"`

	// The ID of the instance to use to create the launch configuration. The new
	// launch configuration derives attributes from the instance, except for the
	// block device mapping.
	//
	// To create a launch configuration with a block device mapping or override
	// any other instance attributes, specify them as part of the same request.
	//
	// For more information, see Create a Launch Configuration Using an EC2 Instance
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-lc-with-instanceID.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	// If you do not specify InstanceId, you must specify both ImageId and InstanceType.
	InstanceId *string `min:"1" type:"string"`

	// Controls whether instances in this group are launched with detailed (true)
	// or basic (false) monitoring.
	//
	// The default value is true (enabled).
	//
	// When detailed monitoring is enabled, Amazon CloudWatch generates metrics
	// every minute and your account is charged a fee. When you disable detailed
	// monitoring, CloudWatch generates metrics every 5 minutes. For more information,
	// see Configure Monitoring for Auto Scaling Instances (https://docs.aws.amazon.com/autoscaling/latest/userguide/as-instance-monitoring.html#enable-as-instance-metrics)
	// in the Amazon EC2 Auto Scaling User Guide.
	InstanceMonitoring *InstanceMonitoring `type:"structure"`

	// Specifies the instance type of the EC2 instance.
	//
	// For information about available instance types, see Available Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes)
	// in the Amazon EC2 User Guide for Linux Instances.
	//
	// If you do not specify InstanceId, you must specify InstanceType.
	InstanceType *string `min:"1" type:"string"`

	// The ID of the kernel associated with the AMI.
	KernelId *string `min:"1" type:"string"`

	// The name of the key pair. For more information, see Amazon EC2 Key Pairs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	KeyName *string `min:"1" type:"string"`

	// The name of the launch configuration. This name must be unique per Region
	// per account.
	//
	// LaunchConfigurationName is a required field
	LaunchConfigurationName *string `min:"1" type:"string" required:"true"`

	// The tenancy of the instance. An instance with dedicated tenancy runs on isolated,
	// single-tenant hardware and can only be launched into a VPC.
	//
	// To launch dedicated instances into a shared tenancy VPC (a VPC with the instance
	// placement tenancy attribute set to default), you must set the value of this
	// parameter to dedicated.
	//
	// If you specify PlacementTenancy, you must specify at least one subnet for
	// VPCZoneIdentifier when you create your group.
	//
	// For more information, see Instance Placement Tenancy (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-vpc-tenancy)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	// Valid values: default | dedicated
	PlacementTenancy *string `min:"1" type:"string"`

	// The ID of the RAM disk to select.
	RamdiskId *string `min:"1" type:"string"`

	// A list that contains the security groups to assign to the instances in the
	// Auto Scaling group.
	//
	// [EC2-VPC] Specify the security group IDs. For more information, see Security
	// Groups for Your VPC (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html)
	// in the Amazon Virtual Private Cloud User Guide.
	//
	// [EC2-Classic] Specify either the security group names or the security group
	// IDs. For more information, see Amazon EC2 Security Groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	SecurityGroups []string `type:"list"`

	// The maximum hourly price to be paid for any Spot Instance launched to fulfill
	// the request. Spot Instances are launched when the price you specify exceeds
	// the current Spot market price. For more information, see Launching Spot Instances
	// in Your Auto Scaling Group (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-launch-spot-instances.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	// If a Spot price is set, then the Auto Scaling group will only launch instances
	// when the Spot price has been met, regardless of the setting in the Auto Scaling
	// group's DesiredCapacity.
	//
	// When you change your Spot price by creating a new launch configuration, running
	// instances will continue to run as long as the Spot price for those running
	// instances is higher than the current Spot market price.
	SpotPrice *string `min:"1" type:"string"`

	// The Base64-encoded user data to make available to the launched EC2 instances.
	// For more information, see Instance Metadata and User Data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	UserData *string `type:"string"`
}

// String returns the string representation
func (s CreateLaunchConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLaunchConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLaunchConfigurationInput"}
	if s.ClassicLinkVPCId != nil && len(*s.ClassicLinkVPCId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClassicLinkVPCId", 1))
	}
	if s.IamInstanceProfile != nil && len(*s.IamInstanceProfile) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IamInstanceProfile", 1))
	}
	if s.ImageId != nil && len(*s.ImageId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImageId", 1))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}
	if s.InstanceType != nil && len(*s.InstanceType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceType", 1))
	}
	if s.KernelId != nil && len(*s.KernelId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KernelId", 1))
	}
	if s.KeyName != nil && len(*s.KeyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyName", 1))
	}

	if s.LaunchConfigurationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LaunchConfigurationName"))
	}
	if s.LaunchConfigurationName != nil && len(*s.LaunchConfigurationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LaunchConfigurationName", 1))
	}
	if s.PlacementTenancy != nil && len(*s.PlacementTenancy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementTenancy", 1))
	}
	if s.RamdiskId != nil && len(*s.RamdiskId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RamdiskId", 1))
	}
	if s.SpotPrice != nil && len(*s.SpotPrice) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SpotPrice", 1))
	}
	if s.BlockDeviceMappings != nil {
		for i, v := range s.BlockDeviceMappings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "BlockDeviceMappings", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateLaunchConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateLaunchConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLaunchConfiguration = "CreateLaunchConfiguration"

// CreateLaunchConfigurationRequest returns a request value for making API operation for
// Auto Scaling.
//
// Creates a launch configuration.
//
// If you exceed your maximum limit of launch configurations, the call fails.
// For information about viewing this limit, see DescribeAccountLimits. For
// information about updating this limit, see Amazon EC2 Auto Scaling Limits
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-account-limits.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
// For more information, see Launch Configurations (https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchConfiguration.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using CreateLaunchConfigurationRequest.
//    req := client.CreateLaunchConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/CreateLaunchConfiguration
func (c *Client) CreateLaunchConfigurationRequest(input *CreateLaunchConfigurationInput) CreateLaunchConfigurationRequest {
	op := &aws.Operation{
		Name:       opCreateLaunchConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLaunchConfigurationInput{}
	}

	req := c.newRequest(op, input, &CreateLaunchConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CreateLaunchConfigurationRequest{Request: req, Input: input, Copy: c.CreateLaunchConfigurationRequest}
}

// CreateLaunchConfigurationRequest is the request type for the
// CreateLaunchConfiguration API operation.
type CreateLaunchConfigurationRequest struct {
	*aws.Request
	Input *CreateLaunchConfigurationInput
	Copy  func(*CreateLaunchConfigurationInput) CreateLaunchConfigurationRequest
}

// Send marshals and sends the CreateLaunchConfiguration API request.
func (r CreateLaunchConfigurationRequest) Send(ctx context.Context) (*CreateLaunchConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLaunchConfigurationResponse{
		CreateLaunchConfigurationOutput: r.Request.Data.(*CreateLaunchConfigurationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLaunchConfigurationResponse is the response type for the
// CreateLaunchConfiguration API operation.
type CreateLaunchConfigurationResponse struct {
	*CreateLaunchConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLaunchConfiguration request.
func (r *CreateLaunchConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
