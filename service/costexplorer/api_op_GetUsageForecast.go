// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetUsageForecastInput struct {
	_ struct{} `type:"structure"`

	// The filters that you want to use to filter your forecast. Cost Explorer API
	// supports all of the Cost Explorer filters.
	Filter *Expression `type:"structure"`

	// How granular you want the forecast to be. You can get 3 months of DAILY forecasts
	// or 12 months of MONTHLY forecasts.
	//
	// The GetUsageForecast operation supports only DAILY and MONTHLY granularities.
	//
	// Granularity is a required field
	Granularity Granularity `type:"string" required:"true" enum:"true"`

	// Which metric Cost Explorer uses to create your forecast.
	//
	// Valid values for a GetUsageForecast call are the following:
	//
	//    * USAGE_QUANTITY
	//
	//    * NORMALIZED_USAGE_AMOUNT
	//
	// Metric is a required field
	Metric Metric `type:"string" required:"true" enum:"true"`

	// Cost Explorer always returns the mean forecast as a single point. You can
	// request a prediction interval around the mean by specifying a confidence
	// level. The higher the confidence level, the more confident Cost Explorer
	// is about the actual value falling in the prediction interval. Higher confidence
	// levels result in wider prediction intervals.
	PredictionIntervalLevel *int64 `min:"51" type:"integer"`

	// The start and end dates of the period that you want to retrieve usage forecast
	// for. The start date is inclusive, but the end date is exclusive. For example,
	// if start is 2017-01-01 and end is 2017-05-01, then the cost and usage data
	// is retrieved from 2017-01-01 up to and including 2017-04-30 but not including
	// 2017-05-01.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetUsageForecastInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUsageForecastInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUsageForecastInput"}
	if len(s.Granularity) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Granularity"))
	}
	if len(s.Metric) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Metric"))
	}
	if s.PredictionIntervalLevel != nil && *s.PredictionIntervalLevel < 51 {
		invalidParams.Add(aws.NewErrParamMinValue("PredictionIntervalLevel", 51))
	}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetUsageForecastOutput struct {
	_ struct{} `type:"structure"`

	// The forecasts for your query, in order. For DAILY forecasts, this is a list
	// of days. For MONTHLY forecasts, this is a list of months.
	ForecastResultsByTime []ForecastResult `type:"list"`

	// How much you're forecasted to use over the forecast period.
	Total *MetricValue `type:"structure"`
}

// String returns the string representation
func (s GetUsageForecastOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetUsageForecast = "GetUsageForecast"

// GetUsageForecastRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Retrieves a forecast for how much Amazon Web Services predicts that you will
// use over the forecast time period that you select, based on your past usage.
//
//    // Example sending a request using GetUsageForecastRequest.
//    req := client.GetUsageForecastRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetUsageForecast
func (c *Client) GetUsageForecastRequest(input *GetUsageForecastInput) GetUsageForecastRequest {
	op := &aws.Operation{
		Name:       opGetUsageForecast,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetUsageForecastInput{}
	}

	req := c.newRequest(op, input, &GetUsageForecastOutput{})
	return GetUsageForecastRequest{Request: req, Input: input, Copy: c.GetUsageForecastRequest}
}

// GetUsageForecastRequest is the request type for the
// GetUsageForecast API operation.
type GetUsageForecastRequest struct {
	*aws.Request
	Input *GetUsageForecastInput
	Copy  func(*GetUsageForecastInput) GetUsageForecastRequest
}

// Send marshals and sends the GetUsageForecast API request.
func (r GetUsageForecastRequest) Send(ctx context.Context) (*GetUsageForecastResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUsageForecastResponse{
		GetUsageForecastOutput: r.Request.Data.(*GetUsageForecastOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUsageForecastResponse is the response type for the
// GetUsageForecast API operation.
type GetUsageForecastResponse struct {
	*GetUsageForecastOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUsageForecast request.
func (r *GetUsageForecastResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
