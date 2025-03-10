// Code generated by internal/awstesting/cmd/op_crawler/generate.go. DO NOT EDIT.
// +build sdktool

package main

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
	"github.com/aws/aws-sdk-go-v2/service/chime"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codestar"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/connect"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
	"github.com/aws/aws-sdk-go-v2/service/datasync"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/aws/aws-sdk-go-v2/service/forecast"
	"github.com/aws/aws-sdk-go-v2/service/forecastquery"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/greengrass"
	"github.com/aws/aws-sdk-go-v2/service/groundstation"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/health"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/aws/aws-sdk-go-v2/service/iotevents"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata"
	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning"
	"github.com/aws/aws-sdk-go-v2/service/macie"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics"
	"github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub"
	"github.com/aws/aws-sdk-go-v2/service/mobile"
	"github.com/aws/aws-sdk-go-v2/service/mobileanalytics"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mturk"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/personalize"
	"github.com/aws/aws-sdk-go-v2/service/personalizeevents"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime"
	"github.com/aws/aws-sdk-go-v2/service/pi"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldbsession"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go-v2/service/robomaker"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerruntime"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/aws/aws-sdk-go-v2/service/simpledb"
	"github.com/aws/aws-sdk-go-v2/service/sms"
	"github.com/aws/aws-sdk-go-v2/service/snowball"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sso"
	"github.com/aws/aws-sdk-go-v2/service/ssooidc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/swf"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/workdocs"
	"github.com/aws/aws-sdk-go-v2/service/worklink"
	"github.com/aws/aws-sdk-go-v2/service/workmail"
	"github.com/aws/aws-sdk-go-v2/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

type service struct {
	name  string
	value reflect.Value
}

func createServices(cfg aws.Config) []service {
	s3Client := s3.New(cfg)
	s3Client.ForcePathStyle = true
	sqsClient := sqs.New(cfg)
	sqsClient.DisableComputeChecksums = true

	return []service{
		{name: "acm", value: reflect.ValueOf(acm.New(cfg))},
		{name: "acmpca", value: reflect.ValueOf(acmpca.New(cfg))},
		{name: "alexaforbusiness", value: reflect.ValueOf(alexaforbusiness.New(cfg))},
		{name: "amplify", value: reflect.ValueOf(amplify.New(cfg))},
		{name: "apigateway", value: reflect.ValueOf(apigateway.New(cfg))},
		{name: "apigatewaymanagementapi", value: reflect.ValueOf(apigatewaymanagementapi.New(cfg))},
		{name: "apigatewayv2", value: reflect.ValueOf(apigatewayv2.New(cfg))},
		{name: "applicationautoscaling", value: reflect.ValueOf(applicationautoscaling.New(cfg))},
		{name: "applicationdiscoveryservice", value: reflect.ValueOf(applicationdiscoveryservice.New(cfg))},
		{name: "applicationinsights", value: reflect.ValueOf(applicationinsights.New(cfg))},
		{name: "appmesh", value: reflect.ValueOf(appmesh.New(cfg))},
		{name: "appstream", value: reflect.ValueOf(appstream.New(cfg))},
		{name: "appsync", value: reflect.ValueOf(appsync.New(cfg))},
		{name: "athena", value: reflect.ValueOf(athena.New(cfg))},
		{name: "autoscaling", value: reflect.ValueOf(autoscaling.New(cfg))},
		{name: "autoscalingplans", value: reflect.ValueOf(autoscalingplans.New(cfg))},
		{name: "backup", value: reflect.ValueOf(backup.New(cfg))},
		{name: "batch", value: reflect.ValueOf(batch.New(cfg))},
		{name: "budgets", value: reflect.ValueOf(budgets.New(cfg))},
		{name: "chime", value: reflect.ValueOf(chime.New(cfg))},
		{name: "cloud9", value: reflect.ValueOf(cloud9.New(cfg))},
		{name: "clouddirectory", value: reflect.ValueOf(clouddirectory.New(cfg))},
		{name: "cloudformation", value: reflect.ValueOf(cloudformation.New(cfg))},
		{name: "cloudfront", value: reflect.ValueOf(cloudfront.New(cfg))},
		{name: "cloudhsm", value: reflect.ValueOf(cloudhsm.New(cfg))},
		{name: "cloudhsmv2", value: reflect.ValueOf(cloudhsmv2.New(cfg))},
		{name: "cloudsearch", value: reflect.ValueOf(cloudsearch.New(cfg))},
		{name: "cloudsearchdomain", value: reflect.ValueOf(cloudsearchdomain.New(cfg))},
		{name: "cloudtrail", value: reflect.ValueOf(cloudtrail.New(cfg))},
		{name: "cloudwatch", value: reflect.ValueOf(cloudwatch.New(cfg))},
		{name: "cloudwatchevents", value: reflect.ValueOf(cloudwatchevents.New(cfg))},
		{name: "cloudwatchlogs", value: reflect.ValueOf(cloudwatchlogs.New(cfg))},
		{name: "codebuild", value: reflect.ValueOf(codebuild.New(cfg))},
		{name: "codecommit", value: reflect.ValueOf(codecommit.New(cfg))},
		{name: "codedeploy", value: reflect.ValueOf(codedeploy.New(cfg))},
		{name: "codepipeline", value: reflect.ValueOf(codepipeline.New(cfg))},
		{name: "codestar", value: reflect.ValueOf(codestar.New(cfg))},
		{name: "codestarnotifications", value: reflect.ValueOf(codestarnotifications.New(cfg))},
		{name: "cognitoidentity", value: reflect.ValueOf(cognitoidentity.New(cfg))},
		{name: "cognitoidentityprovider", value: reflect.ValueOf(cognitoidentityprovider.New(cfg))},
		{name: "cognitosync", value: reflect.ValueOf(cognitosync.New(cfg))},
		{name: "comprehend", value: reflect.ValueOf(comprehend.New(cfg))},
		{name: "comprehendmedical", value: reflect.ValueOf(comprehendmedical.New(cfg))},
		{name: "configservice", value: reflect.ValueOf(configservice.New(cfg))},
		{name: "connect", value: reflect.ValueOf(connect.New(cfg))},
		{name: "costandusagereportservice", value: reflect.ValueOf(costandusagereportservice.New(cfg))},
		{name: "costexplorer", value: reflect.ValueOf(costexplorer.New(cfg))},
		{name: "databasemigrationservice", value: reflect.ValueOf(databasemigrationservice.New(cfg))},
		{name: "datapipeline", value: reflect.ValueOf(datapipeline.New(cfg))},
		{name: "datasync", value: reflect.ValueOf(datasync.New(cfg))},
		{name: "dax", value: reflect.ValueOf(dax.New(cfg))},
		{name: "devicefarm", value: reflect.ValueOf(devicefarm.New(cfg))},
		{name: "directconnect", value: reflect.ValueOf(directconnect.New(cfg))},
		{name: "directoryservice", value: reflect.ValueOf(directoryservice.New(cfg))},
		{name: "dlm", value: reflect.ValueOf(dlm.New(cfg))},
		{name: "docdb", value: reflect.ValueOf(docdb.New(cfg))},
		{name: "dynamodb", value: reflect.ValueOf(dynamodb.New(cfg))},
		{name: "dynamodbstreams", value: reflect.ValueOf(dynamodbstreams.New(cfg))},
		{name: "ec2", value: reflect.ValueOf(ec2.New(cfg))},
		{name: "ec2instanceconnect", value: reflect.ValueOf(ec2instanceconnect.New(cfg))},
		{name: "ecr", value: reflect.ValueOf(ecr.New(cfg))},
		{name: "ecs", value: reflect.ValueOf(ecs.New(cfg))},
		{name: "efs", value: reflect.ValueOf(efs.New(cfg))},
		{name: "eks", value: reflect.ValueOf(eks.New(cfg))},
		{name: "elasticache", value: reflect.ValueOf(elasticache.New(cfg))},
		{name: "elasticbeanstalk", value: reflect.ValueOf(elasticbeanstalk.New(cfg))},
		{name: "elasticloadbalancing", value: reflect.ValueOf(elasticloadbalancing.New(cfg))},
		{name: "elasticloadbalancingv2", value: reflect.ValueOf(elasticloadbalancingv2.New(cfg))},
		{name: "elasticsearchservice", value: reflect.ValueOf(elasticsearchservice.New(cfg))},
		{name: "elastictranscoder", value: reflect.ValueOf(elastictranscoder.New(cfg))},
		{name: "emr", value: reflect.ValueOf(emr.New(cfg))},
		{name: "eventbridge", value: reflect.ValueOf(eventbridge.New(cfg))},
		{name: "firehose", value: reflect.ValueOf(firehose.New(cfg))},
		{name: "fms", value: reflect.ValueOf(fms.New(cfg))},
		{name: "forecast", value: reflect.ValueOf(forecast.New(cfg))},
		{name: "forecastquery", value: reflect.ValueOf(forecastquery.New(cfg))},
		{name: "fsx", value: reflect.ValueOf(fsx.New(cfg))},
		{name: "gamelift", value: reflect.ValueOf(gamelift.New(cfg))},
		{name: "glacier", value: reflect.ValueOf(glacier.New(cfg))},
		{name: "globalaccelerator", value: reflect.ValueOf(globalaccelerator.New(cfg))},
		{name: "glue", value: reflect.ValueOf(glue.New(cfg))},
		{name: "greengrass", value: reflect.ValueOf(greengrass.New(cfg))},
		{name: "groundstation", value: reflect.ValueOf(groundstation.New(cfg))},
		{name: "guardduty", value: reflect.ValueOf(guardduty.New(cfg))},
		{name: "health", value: reflect.ValueOf(health.New(cfg))},
		{name: "iam", value: reflect.ValueOf(iam.New(cfg))},
		{name: "inspector", value: reflect.ValueOf(inspector.New(cfg))},
		{name: "iot", value: reflect.ValueOf(iot.New(cfg))},
		{name: "iot1clickdevicesservice", value: reflect.ValueOf(iot1clickdevicesservice.New(cfg))},
		{name: "iot1clickprojects", value: reflect.ValueOf(iot1clickprojects.New(cfg))},
		{name: "iotanalytics", value: reflect.ValueOf(iotanalytics.New(cfg))},
		{name: "iotdataplane", value: reflect.ValueOf(iotdataplane.New(cfg))},
		{name: "iotevents", value: reflect.ValueOf(iotevents.New(cfg))},
		{name: "ioteventsdata", value: reflect.ValueOf(ioteventsdata.New(cfg))},
		{name: "iotjobsdataplane", value: reflect.ValueOf(iotjobsdataplane.New(cfg))},
		{name: "iotthingsgraph", value: reflect.ValueOf(iotthingsgraph.New(cfg))},
		{name: "kafka", value: reflect.ValueOf(kafka.New(cfg))},
		{name: "kinesis", value: reflect.ValueOf(kinesis.New(cfg))},
		{name: "kinesisanalytics", value: reflect.ValueOf(kinesisanalytics.New(cfg))},
		{name: "kinesisanalyticsv2", value: reflect.ValueOf(kinesisanalyticsv2.New(cfg))},
		{name: "kinesisvideo", value: reflect.ValueOf(kinesisvideo.New(cfg))},
		{name: "kinesisvideoarchivedmedia", value: reflect.ValueOf(kinesisvideoarchivedmedia.New(cfg))},
		{name: "kinesisvideomedia", value: reflect.ValueOf(kinesisvideomedia.New(cfg))},
		{name: "kms", value: reflect.ValueOf(kms.New(cfg))},
		{name: "lakeformation", value: reflect.ValueOf(lakeformation.New(cfg))},
		{name: "lambda", value: reflect.ValueOf(lambda.New(cfg))},
		{name: "lexmodelbuildingservice", value: reflect.ValueOf(lexmodelbuildingservice.New(cfg))},
		{name: "lexruntimeservice", value: reflect.ValueOf(lexruntimeservice.New(cfg))},
		{name: "licensemanager", value: reflect.ValueOf(licensemanager.New(cfg))},
		{name: "lightsail", value: reflect.ValueOf(lightsail.New(cfg))},
		{name: "machinelearning", value: reflect.ValueOf(machinelearning.New(cfg))},
		{name: "macie", value: reflect.ValueOf(macie.New(cfg))},
		{name: "managedblockchain", value: reflect.ValueOf(managedblockchain.New(cfg))},
		{name: "marketplacecatalog", value: reflect.ValueOf(marketplacecatalog.New(cfg))},
		{name: "marketplacecommerceanalytics", value: reflect.ValueOf(marketplacecommerceanalytics.New(cfg))},
		{name: "marketplaceentitlementservice", value: reflect.ValueOf(marketplaceentitlementservice.New(cfg))},
		{name: "marketplacemetering", value: reflect.ValueOf(marketplacemetering.New(cfg))},
		{name: "mediaconnect", value: reflect.ValueOf(mediaconnect.New(cfg))},
		{name: "mediaconvert", value: reflect.ValueOf(mediaconvert.New(cfg))},
		{name: "medialive", value: reflect.ValueOf(medialive.New(cfg))},
		{name: "mediapackage", value: reflect.ValueOf(mediapackage.New(cfg))},
		{name: "mediapackagevod", value: reflect.ValueOf(mediapackagevod.New(cfg))},
		{name: "mediastore", value: reflect.ValueOf(mediastore.New(cfg))},
		{name: "mediastoredata", value: reflect.ValueOf(mediastoredata.New(cfg))},
		{name: "mediatailor", value: reflect.ValueOf(mediatailor.New(cfg))},
		{name: "migrationhub", value: reflect.ValueOf(migrationhub.New(cfg))},
		{name: "mobile", value: reflect.ValueOf(mobile.New(cfg))},
		{name: "mobileanalytics", value: reflect.ValueOf(mobileanalytics.New(cfg))},
		{name: "mq", value: reflect.ValueOf(mq.New(cfg))},
		{name: "mturk", value: reflect.ValueOf(mturk.New(cfg))},
		{name: "neptune", value: reflect.ValueOf(neptune.New(cfg))},
		{name: "opsworks", value: reflect.ValueOf(opsworks.New(cfg))},
		{name: "opsworkscm", value: reflect.ValueOf(opsworkscm.New(cfg))},
		{name: "organizations", value: reflect.ValueOf(organizations.New(cfg))},
		{name: "personalize", value: reflect.ValueOf(personalize.New(cfg))},
		{name: "personalizeevents", value: reflect.ValueOf(personalizeevents.New(cfg))},
		{name: "personalizeruntime", value: reflect.ValueOf(personalizeruntime.New(cfg))},
		{name: "pi", value: reflect.ValueOf(pi.New(cfg))},
		{name: "pinpoint", value: reflect.ValueOf(pinpoint.New(cfg))},
		{name: "pinpointemail", value: reflect.ValueOf(pinpointemail.New(cfg))},
		{name: "pinpointsmsvoice", value: reflect.ValueOf(pinpointsmsvoice.New(cfg))},
		{name: "polly", value: reflect.ValueOf(polly.New(cfg))},
		{name: "pricing", value: reflect.ValueOf(pricing.New(cfg))},
		{name: "qldb", value: reflect.ValueOf(qldb.New(cfg))},
		{name: "qldbsession", value: reflect.ValueOf(qldbsession.New(cfg))},
		{name: "quicksight", value: reflect.ValueOf(quicksight.New(cfg))},
		{name: "ram", value: reflect.ValueOf(ram.New(cfg))},
		{name: "rds", value: reflect.ValueOf(rds.New(cfg))},
		{name: "rdsdata", value: reflect.ValueOf(rdsdata.New(cfg))},
		{name: "redshift", value: reflect.ValueOf(redshift.New(cfg))},
		{name: "rekognition", value: reflect.ValueOf(rekognition.New(cfg))},
		{name: "resourcegroups", value: reflect.ValueOf(resourcegroups.New(cfg))},
		{name: "resourcegroupstaggingapi", value: reflect.ValueOf(resourcegroupstaggingapi.New(cfg))},
		{name: "robomaker", value: reflect.ValueOf(robomaker.New(cfg))},
		{name: "route53", value: reflect.ValueOf(route53.New(cfg))},
		{name: "route53domains", value: reflect.ValueOf(route53domains.New(cfg))},
		{name: "route53resolver", value: reflect.ValueOf(route53resolver.New(cfg))},
		{name: "s3", value: reflect.ValueOf(s3Client)},
		{name: "s3control", value: reflect.ValueOf(s3control.New(cfg))},
		{name: "sagemaker", value: reflect.ValueOf(sagemaker.New(cfg))},
		{name: "sagemakerruntime", value: reflect.ValueOf(sagemakerruntime.New(cfg))},
		{name: "savingsplans", value: reflect.ValueOf(savingsplans.New(cfg))},
		{name: "secretsmanager", value: reflect.ValueOf(secretsmanager.New(cfg))},
		{name: "securityhub", value: reflect.ValueOf(securityhub.New(cfg))},
		{name: "serverlessapplicationrepository", value: reflect.ValueOf(serverlessapplicationrepository.New(cfg))},
		{name: "servicecatalog", value: reflect.ValueOf(servicecatalog.New(cfg))},
		{name: "servicediscovery", value: reflect.ValueOf(servicediscovery.New(cfg))},
		{name: "servicequotas", value: reflect.ValueOf(servicequotas.New(cfg))},
		{name: "ses", value: reflect.ValueOf(ses.New(cfg))},
		{name: "sfn", value: reflect.ValueOf(sfn.New(cfg))},
		{name: "shield", value: reflect.ValueOf(shield.New(cfg))},
		{name: "signer", value: reflect.ValueOf(signer.New(cfg))},
		{name: "simpledb", value: reflect.ValueOf(simpledb.New(cfg))},
		{name: "sms", value: reflect.ValueOf(sms.New(cfg))},
		{name: "snowball", value: reflect.ValueOf(snowball.New(cfg))},
		{name: "sns", value: reflect.ValueOf(sns.New(cfg))},
		{name: "sqs", value: reflect.ValueOf(sqsClient)},
		{name: "ssm", value: reflect.ValueOf(ssm.New(cfg))},
		{name: "sso", value: reflect.ValueOf(sso.New(cfg))},
		{name: "ssooidc", value: reflect.ValueOf(ssooidc.New(cfg))},
		{name: "storagegateway", value: reflect.ValueOf(storagegateway.New(cfg))},
		{name: "sts", value: reflect.ValueOf(sts.New(cfg))},
		{name: "support", value: reflect.ValueOf(support.New(cfg))},
		{name: "swf", value: reflect.ValueOf(swf.New(cfg))},
		{name: "textract", value: reflect.ValueOf(textract.New(cfg))},
		{name: "transcribe", value: reflect.ValueOf(transcribe.New(cfg))},
		{name: "transfer", value: reflect.ValueOf(transfer.New(cfg))},
		{name: "translate", value: reflect.ValueOf(translate.New(cfg))},
		{name: "waf", value: reflect.ValueOf(waf.New(cfg))},
		{name: "wafregional", value: reflect.ValueOf(wafregional.New(cfg))},
		{name: "workdocs", value: reflect.ValueOf(workdocs.New(cfg))},
		{name: "worklink", value: reflect.ValueOf(worklink.New(cfg))},
		{name: "workmail", value: reflect.ValueOf(workmail.New(cfg))},
		{name: "workmailmessageflow", value: reflect.ValueOf(workmailmessageflow.New(cfg))},
		{name: "workspaces", value: reflect.ValueOf(workspaces.New(cfg))},
		{name: "xray", value: reflect.ValueOf(xray.New(cfg))},
	}
}
