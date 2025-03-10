// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetCostForecastInput struct {
	_ struct{} `type:"structure"`

	// The filters that you want to use to filter your forecast. Cost Explorer API
	// supports all of the Cost Explorer filters.
	Filter *Expression `type:"structure"`

	// How granular you want the forecast to be. You can get 3 months of DAILY forecasts
	// or 12 months of MONTHLY forecasts.
	//
	// The GetCostForecast operation supports only DAILY and MONTHLY granularities.
	//
	// Granularity is a required field
	Granularity Granularity `type:"string" required:"true" enum:"true"`

	// Which metric Cost Explorer uses to create your forecast. For more information
	// about blended and unblended rates, see Why does the "blended" annotation
	// appear on some line items in my bill? (https://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/).
	//
	// Valid values for a GetCostForecast call are the following:
	//
	//    * AMORTIZED_COST
	//
	//    * BLENDED_COST
	//
	//    * NET_AMORTIZED_COST
	//
	//    * NET_UNBLENDED_COST
	//
	//    * UNBLENDED_COST
	//
	// Metric is a required field
	Metric Metric `type:"string" required:"true" enum:"true"`

	// Cost Explorer always returns the mean forecast as a single point. You can
	// request a prediction interval around the mean by specifying a confidence
	// level. The higher the confidence level, the more confident Cost Explorer
	// is about the actual value falling in the prediction interval. Higher confidence
	// levels result in wider prediction intervals.
	PredictionIntervalLevel *int64 `min:"51" type:"integer"`

	// The period of time that you want the forecast to cover.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetCostForecastInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCostForecastInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCostForecastInput"}
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

type GetCostForecastOutput struct {
	_ struct{} `type:"structure"`

	// The forecasts for your query, in order. For DAILY forecasts, this is a list
	// of days. For MONTHLY forecasts, this is a list of months.
	ForecastResultsByTime []ForecastResult `type:"list"`

	// How much you are forecasted to spend over the forecast period, in USD.
	Total *MetricValue `type:"structure"`
}

// String returns the string representation
func (s GetCostForecastOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCostForecast = "GetCostForecast"

// GetCostForecastRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Retrieves a forecast for how much Amazon Web Services predicts that you will
// spend over the forecast time period that you select, based on your past costs.
//
//    // Example sending a request using GetCostForecastRequest.
//    req := client.GetCostForecastRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetCostForecast
func (c *Client) GetCostForecastRequest(input *GetCostForecastInput) GetCostForecastRequest {
	op := &aws.Operation{
		Name:       opGetCostForecast,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCostForecastInput{}
	}

	req := c.newRequest(op, input, &GetCostForecastOutput{})
	return GetCostForecastRequest{Request: req, Input: input, Copy: c.GetCostForecastRequest}
}

// GetCostForecastRequest is the request type for the
// GetCostForecast API operation.
type GetCostForecastRequest struct {
	*aws.Request
	Input *GetCostForecastInput
	Copy  func(*GetCostForecastInput) GetCostForecastRequest
}

// Send marshals and sends the GetCostForecast API request.
func (r GetCostForecastRequest) Send(ctx context.Context) (*GetCostForecastResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCostForecastResponse{
		GetCostForecastOutput: r.Request.Data.(*GetCostForecastOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCostForecastResponse is the response type for the
// GetCostForecast API operation.
type GetCostForecastResponse struct {
	*GetCostForecastOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCostForecast request.
func (r *GetCostForecastResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
