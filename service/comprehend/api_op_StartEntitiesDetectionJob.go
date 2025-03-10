// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartEntitiesDetectionJobInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the request. If you don't set the client request
	// token, Amazon Comprehend generates one.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management
	// (IAM) role that grants Amazon Comprehend read access to your input data.
	// For more information, see https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions.html#auth-role-permissions
	// (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions.html#auth-role-permissions).
	//
	// DataAccessRoleArn is a required field
	DataAccessRoleArn *string `min:"20" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) that identifies the specific entity recognizer
	// to be used by the StartEntitiesDetectionJob. This ARN is optional and is
	// only used for a custom entity recognition job.
	EntityRecognizerArn *string `type:"string"`

	// Specifies the format and location of the input data for the job.
	//
	// InputDataConfig is a required field
	InputDataConfig *InputDataConfig `type:"structure" required:"true"`

	// The identifier of the job.
	JobName *string `min:"1" type:"string"`

	// The language of the input documents. All documents must be in the same language.
	// You can specify any of the languages supported by Amazon Comprehend: English
	// ("en"), Spanish ("es"), French ("fr"), German ("de"), Italian ("it"), or
	// Portuguese ("pt"). If custom entities recognition is used, this parameter
	// is ignored and the language used for training the model is used instead.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// Specifies where to send the output files.
	//
	// OutputDataConfig is a required field
	OutputDataConfig *OutputDataConfig `type:"structure" required:"true"`

	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses
	// to encrypt data on the storage volume attached to the ML compute instance(s)
	// that process the analysis job. The VolumeKmsKeyId can be either of the following
	// formats:
	//
	//    * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	//    * Amazon Resource Name (ARN) of a KMS Key: "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	VolumeKmsKeyId *string `type:"string"`

	// Configuration parameters for an optional private Virtual Private Cloud (VPC)
	// containing the resources you are using for your entity detection job. For
	// more information, see Amazon VPC (https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).
	VpcConfig *VpcConfig `type:"structure"`
}

// String returns the string representation
func (s StartEntitiesDetectionJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartEntitiesDetectionJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartEntitiesDetectionJobInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.DataAccessRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataAccessRoleArn"))
	}
	if s.DataAccessRoleArn != nil && len(*s.DataAccessRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("DataAccessRoleArn", 20))
	}

	if s.InputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputDataConfig"))
	}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.OutputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputDataConfig"))
	}
	if s.InputDataConfig != nil {
		if err := s.InputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.OutputDataConfig != nil {
		if err := s.OutputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			invalidParams.AddNested("VpcConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartEntitiesDetectionJobOutput struct {
	_ struct{} `type:"structure"`

	// The identifier generated for the job. To get the status of job, use this
	// identifier with the operation.
	JobId *string `min:"1" type:"string"`

	// The status of the job.
	//
	//    * SUBMITTED - The job has been received and is queued for processing.
	//
	//    * IN_PROGRESS - Amazon Comprehend is processing the job.
	//
	//    * COMPLETED - The job was successfully completed and the output is available.
	//
	//    * FAILED - The job did not complete. To get details, use the operation.
	//
	//    * STOP_REQUESTED - Amazon Comprehend has received a stop request for the
	//    job and is processing the request.
	//
	//    * STOPPED - The job was successfully stopped without completing.
	JobStatus JobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s StartEntitiesDetectionJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartEntitiesDetectionJob = "StartEntitiesDetectionJob"

// StartEntitiesDetectionJobRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Starts an asynchronous entity detection job for a collection of documents.
// Use the operation to track the status of a job.
//
// This API can be used for either standard entity detection or custom entity
// recognition. In order to be used for custom entity recognition, the optional
// EntityRecognizerArn must be used in order to provide access to the recognizer
// being used to detect the custom entity.
//
//    // Example sending a request using StartEntitiesDetectionJobRequest.
//    req := client.StartEntitiesDetectionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/StartEntitiesDetectionJob
func (c *Client) StartEntitiesDetectionJobRequest(input *StartEntitiesDetectionJobInput) StartEntitiesDetectionJobRequest {
	op := &aws.Operation{
		Name:       opStartEntitiesDetectionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartEntitiesDetectionJobInput{}
	}

	req := c.newRequest(op, input, &StartEntitiesDetectionJobOutput{})
	return StartEntitiesDetectionJobRequest{Request: req, Input: input, Copy: c.StartEntitiesDetectionJobRequest}
}

// StartEntitiesDetectionJobRequest is the request type for the
// StartEntitiesDetectionJob API operation.
type StartEntitiesDetectionJobRequest struct {
	*aws.Request
	Input *StartEntitiesDetectionJobInput
	Copy  func(*StartEntitiesDetectionJobInput) StartEntitiesDetectionJobRequest
}

// Send marshals and sends the StartEntitiesDetectionJob API request.
func (r StartEntitiesDetectionJobRequest) Send(ctx context.Context) (*StartEntitiesDetectionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartEntitiesDetectionJobResponse{
		StartEntitiesDetectionJobOutput: r.Request.Data.(*StartEntitiesDetectionJobOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartEntitiesDetectionJobResponse is the response type for the
// StartEntitiesDetectionJob API operation.
type StartEntitiesDetectionJobResponse struct {
	*StartEntitiesDetectionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartEntitiesDetectionJob request.
func (r *StartEntitiesDetectionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
