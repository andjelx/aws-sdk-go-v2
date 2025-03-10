// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationautoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterScalableTargetInput struct {
	_ struct{} `type:"structure"`

	// The maximum value to scale to in response to a scale-out event. MaxCapacity
	// is required to register a scalable target.
	MaxCapacity *int64 `type:"integer"`

	// The minimum value to scale to in response to a scale-in event. MinCapacity
	// is required to register a scalable target.
	MinCapacity *int64 `type:"integer"`

	// The identifier of the resource that is associated with the scalable target.
	// This string consists of the resource type and unique identifier.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot Fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * EMR cluster - The resource type is instancegroup and the unique identifier
	//    is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	//    * AppStream 2.0 fleet - The resource type is fleet and the unique identifier
	//    is the fleet name. Example: fleet/sample-fleet.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the resource ID. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the resource ID. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	//    * Amazon SageMaker endpoint variants - The resource type is variant and
	//    the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.
	//
	//    * Custom resources are not supported with a resource type. This parameter
	//    must specify the OutputValue from the CloudFormation template stack used
	//    to access the resources. The unique identifier is defined by the service
	//    provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// Application Auto Scaling creates a service-linked role that grants it permissions
	// to modify the scalable target on your behalf. For more information, see Service-Linked
	// Roles for Application Auto Scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-service-linked-roles.html).
	//
	// For Amazon EMR, this parameter is required, and it must specify the ARN of
	// an IAM role that allows Application Auto Scaling to modify the scalable target
	// on your behalf.
	RoleARN *string `min:"1" type:"string"`

	// The scalable dimension associated with the scalable target. This string consists
	// of the service namespace, resource type, and scaling property.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    Fleet request.
	//
	//    * elasticmapreduce:instancegroup:InstanceCount - The instance count of
	//    an EMR Instance Group.
	//
	//    * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream
	//    2.0 fleet.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	//    * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances
	//    for an Amazon SageMaker model endpoint variant.
	//
	//    * custom-resource:ResourceType:Property - The scalable dimension for a
	//    custom resource provided by your own application or service.
	//
	// ScalableDimension is a required field
	ScalableDimension ScalableDimension `type:"string" required:"true" enum:"true"`

	// The namespace of the AWS service that provides the resource or custom-resource
	// for a resource provided by your own application or service. For more information,
	// see AWS Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces)
	// in the Amazon Web Services General Reference.
	//
	// ServiceNamespace is a required field
	ServiceNamespace ServiceNamespace `type:"string" required:"true" enum:"true"`

	// An embedded object that contains attributes and attribute values that are
	// used to suspend and resume automatic scaling. Setting the value of an attribute
	// to true suspends the specified scaling activities. Setting it to false (default)
	// resumes the specified scaling activities.
	//
	// Suspension Outcomes
	//
	//    * For DynamicScalingInSuspended, while a suspension is in effect, all
	//    scale-in activities that are triggered by a scaling policy are suspended.
	//
	//    * For DynamicScalingOutSuspended, while a suspension is in effect, all
	//    scale-out activities that are triggered by a scaling policy are suspended.
	//
	//    * For ScheduledScalingSuspended, while a suspension is in effect, all
	//    scaling activities that involve scheduled actions are suspended.
	//
	// For more information, see Suspend and Resume Application Auto Scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-suspend-resume-scaling.html)
	// in the Application Auto Scaling User Guide.
	SuspendedState *SuspendedState `type:"structure"`
}

// String returns the string representation
func (s RegisterScalableTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterScalableTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterScalableTargetInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}
	if s.RoleARN != nil && len(*s.RoleARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleARN", 1))
	}
	if len(s.ScalableDimension) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ScalableDimension"))
	}
	if len(s.ServiceNamespace) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ServiceNamespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterScalableTargetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterScalableTargetOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterScalableTarget = "RegisterScalableTarget"

// RegisterScalableTargetRequest returns a request value for making API operation for
// Application Auto Scaling.
//
// Registers or updates a scalable target. A scalable target is a resource that
// Application Auto Scaling can scale out and scale in. Scalable targets are
// uniquely identified by the combination of resource ID, scalable dimension,
// and namespace.
//
// When you register a new scalable target, you must specify values for minimum
// and maximum capacity. Application Auto Scaling will not scale capacity to
// values that are outside of this range.
//
// To update a scalable target, specify the parameter that you want to change
// as well as the following parameters that identify the scalable target: resource
// ID, scalable dimension, and namespace. Any parameters that you don't specify
// are not changed by this update request.
//
// After you register a scalable target, you do not need to register it again
// to use other Application Auto Scaling operations. To see which resources
// have been registered, use DescribeScalableTargets. You can also view the
// scaling policies for a service namespace by using DescribeScalableTargets.
//
// If you no longer need a scalable target, you can deregister it by using DeregisterScalableTarget.
//
//    // Example sending a request using RegisterScalableTargetRequest.
//    req := client.RegisterScalableTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-autoscaling-2016-02-06/RegisterScalableTarget
func (c *Client) RegisterScalableTargetRequest(input *RegisterScalableTargetInput) RegisterScalableTargetRequest {
	op := &aws.Operation{
		Name:       opRegisterScalableTarget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterScalableTargetInput{}
	}

	req := c.newRequest(op, input, &RegisterScalableTargetOutput{})
	return RegisterScalableTargetRequest{Request: req, Input: input, Copy: c.RegisterScalableTargetRequest}
}

// RegisterScalableTargetRequest is the request type for the
// RegisterScalableTarget API operation.
type RegisterScalableTargetRequest struct {
	*aws.Request
	Input *RegisterScalableTargetInput
	Copy  func(*RegisterScalableTargetInput) RegisterScalableTargetRequest
}

// Send marshals and sends the RegisterScalableTarget API request.
func (r RegisterScalableTargetRequest) Send(ctx context.Context) (*RegisterScalableTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterScalableTargetResponse{
		RegisterScalableTargetOutput: r.Request.Data.(*RegisterScalableTargetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterScalableTargetResponse is the response type for the
// RegisterScalableTarget API operation.
type RegisterScalableTargetResponse struct {
	*RegisterScalableTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterScalableTarget request.
func (r *RegisterScalableTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
