// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexruntimeservice

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PostContentInput struct {
	_ struct{} `type:"structure" payload:"InputStream"`

	// You pass this value as the Accept HTTP header.
	//
	// The message Amazon Lex returns in the response can be either text or speech
	// based on the Accept HTTP header value in the request.
	//
	//    * If the value is text/plain; charset=utf-8, Amazon Lex returns text in
	//    the response.
	//
	//    * If the value begins with audio/, Amazon Lex returns speech in the response.
	//    Amazon Lex uses Amazon Polly to generate the speech (using the configuration
	//    you specified in the Accept header). For example, if you specify audio/mpeg
	//    as the value, Amazon Lex returns speech in the MPEG format.
	//
	//    * If the value is audio/pcm, the speech returned is audio/pcm in 16-bit,
	//    little endian format.
	//
	//    * The following are the accepted values: audio/mpeg audio/ogg audio/pcm
	//    text/plain; charset=utf-8 audio/* (defaults to mpeg)
	Accept *string `location:"header" locationName:"Accept" type:"string"`

	// Alias of the Amazon Lex bot.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"botAlias" type:"string" required:"true"`

	// Name of the Amazon Lex bot.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" type:"string" required:"true"`

	// You pass this value as the Content-Type HTTP header.
	//
	// Indicates the audio format or text. The header value must start with one
	// of the following prefixes:
	//
	//    * PCM format, audio data must be in little-endian byte order. audio/l16;
	//    rate=16000; channels=1 audio/x-l16; sample-rate=16000; channel-count=1
	//    audio/lpcm; sample-rate=8000; sample-size-bits=16; channel-count=1; is-big-endian=false
	//
	//    * Opus format audio/x-cbr-opus-with-preamble; preamble-size=0; bit-rate=256000;
	//    frame-size-milliseconds=4
	//
	//    * Text format text/plain; charset=utf-8
	//
	// ContentType is a required field
	ContentType *string `location:"header" locationName:"Content-Type" type:"string" required:"true"`

	// User input in PCM or Opus audio format or text format as described in the
	// Content-Type HTTP header.
	//
	// You can stream audio data to Amazon Lex or you can create a local buffer
	// that captures all of the audio data before sending. In general, you get better
	// performance if you stream audio data rather than buffering the data locally.
	//
	// To use an non-seekable io.Reader for this request wrap the io.Reader with
	// "aws.ReadSeekCloser". The SDK will not retry request errors for non-seekable
	// readers. This will allow the SDK to send the reader's payload as chunked
	// transfer encoding.
	//
	// InputStream is a required field
	InputStream io.ReadSeeker `locationName:"inputStream" type:"blob" required:"true"`

	// You pass this value as the x-amz-lex-request-attributes HTTP header.
	//
	// Request-specific information passed between Amazon Lex and a client application.
	// The value must be a JSON serialized and base64 encoded map with string keys
	// and values. The total size of the requestAttributes and sessionAttributes
	// headers is limited to 12 KB.
	//
	// The namespace x-amz-lex: is reserved for special attributes. Don't create
	// any request attributes with the prefix x-amz-lex:.
	//
	// For more information, see Setting Request Attributes (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-request-attribs).
	RequestAttributes aws.JSONValue `location:"header" locationName:"x-amz-lex-request-attributes" type:"jsonvalue"`

	// You pass this value as the x-amz-lex-session-attributes HTTP header.
	//
	// Application-specific information passed between Amazon Lex and a client application.
	// The value must be a JSON serialized and base64 encoded map with string keys
	// and values. The total size of the sessionAttributes and requestAttributes
	// headers is limited to 12 KB.
	//
	// For more information, see Setting Session Attributes (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-session-attribs).
	SessionAttributes aws.JSONValue `location:"header" locationName:"x-amz-lex-session-attributes" type:"jsonvalue"`

	// The ID of the client application user. Amazon Lex uses this to identify a
	// user's conversation with your bot. At runtime, each request must contain
	// the userID field.
	//
	// To decide the user ID to use for your application, consider the following
	// factors.
	//
	//    * The userID field must not contain any personally identifiable information
	//    of the user, for example, name, personal identification numbers, or other
	//    end user personal information.
	//
	//    * If you want a user to start a conversation on one device and continue
	//    on another device, use a user-specific identifier.
	//
	//    * If you want the same user to be able to have two independent conversations
	//    on two different devices, choose a device-specific identifier.
	//
	//    * A user can't have two independent conversations with two different versions
	//    of the same bot. For example, a user can't have a conversation with the
	//    PROD and BETA versions of the same bot. If you anticipate that a user
	//    will need to have conversation with two different versions, for example,
	//    while testing, include the bot alias in the user ID to separate the two
	//    conversations.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PostContentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostContentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PostContentInput"}

	if s.BotAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotAlias"))
	}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}

	if s.ContentType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContentType"))
	}

	if s.InputStream == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputStream"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PostContentInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Accept != nil {
		v := *s.Accept

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Accept", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestAttributes != nil {
		v := s.RequestAttributes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-request-attributes", protocol.JSONValue{V: v, EscapeMode: protocol.Base64Escape}, metadata)
	}
	if s.SessionAttributes != nil {
		v := s.SessionAttributes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-session-attributes", protocol.JSONValue{V: v, EscapeMode: protocol.Base64Escape}, metadata)
	}
	if s.BotAlias != nil {
		v := *s.BotAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botAlias", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "userId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InputStream != nil {
		v := s.InputStream

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "inputStream", protocol.ReadSeekerStream{V: v}, metadata)
	}
	return nil
}

type PostContentOutput struct {
	_ struct{} `type:"structure" payload:"AudioStream"`

	// The prompt (or statement) to convey to the user. This is based on the bot
	// configuration and context. For example, if Amazon Lex did not understand
	// the user intent, it sends the clarificationPrompt configured for the bot.
	// If the intent requires confirmation before taking the fulfillment action,
	// it sends the confirmationPrompt. Another example: Suppose that the Lambda
	// function successfully fulfilled the intent, and sent a message to convey
	// to the user. Then Amazon Lex sends that message in the response.
	AudioStream io.ReadCloser `locationName:"audioStream" type:"blob"`

	// Content type as specified in the Accept HTTP header in the request.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Identifies the current state of the user interaction. Amazon Lex returns
	// one of the following values as dialogState. The client can optionally use
	// this information to customize the user interface.
	//
	//    * ElicitIntent - Amazon Lex wants to elicit the user's intent. Consider
	//    the following examples: For example, a user might utter an intent ("I
	//    want to order a pizza"). If Amazon Lex cannot infer the user intent from
	//    this utterance, it will return this dialog state.
	//
	//    * ConfirmIntent - Amazon Lex is expecting a "yes" or "no" response. For
	//    example, Amazon Lex wants user confirmation before fulfilling an intent.
	//    Instead of a simple "yes" or "no" response, a user might respond with
	//    additional information. For example, "yes, but make it a thick crust pizza"
	//    or "no, I want to order a drink." Amazon Lex can process such additional
	//    information (in these examples, update the crust type slot or change the
	//    intent from OrderPizza to OrderDrink).
	//
	//    * ElicitSlot - Amazon Lex is expecting the value of a slot for the current
	//    intent. For example, suppose that in the response Amazon Lex sends this
	//    message: "What size pizza would you like?". A user might reply with the
	//    slot value (e.g., "medium"). The user might also provide additional information
	//    in the response (e.g., "medium thick crust pizza"). Amazon Lex can process
	//    such additional information appropriately.
	//
	//    * Fulfilled - Conveys that the Lambda function has successfully fulfilled
	//    the intent.
	//
	//    * ReadyForFulfillment - Conveys that the client has to fulfill the request.
	//
	//    * Failed - Conveys that the conversation with the user failed. This can
	//    happen for various reasons, including that the user does not provide an
	//    appropriate response to prompts from the service (you can configure how
	//    many times Amazon Lex can prompt a user for specific information), or
	//    if the Lambda function fails to fulfill the intent.
	DialogState DialogState `location:"header" locationName:"x-amz-lex-dialog-state" type:"string" enum:"true"`

	// The text used to process the request.
	//
	// If the input was an audio stream, the inputTranscript field contains the
	// text extracted from the audio stream. This is the text that is actually processed
	// to recognize intents and slot values. You can use this information to determine
	// if Amazon Lex is correctly processing the audio that you send.
	InputTranscript *string `location:"header" locationName:"x-amz-lex-input-transcript" type:"string"`

	// Current user intent that Amazon Lex is aware of.
	IntentName *string `location:"header" locationName:"x-amz-lex-intent-name" type:"string"`

	// The message to convey to the user. The message can come from the bot's configuration
	// or from a Lambda function.
	//
	// If the intent is not configured with a Lambda function, or if the Lambda
	// function returned Delegate as the dialogAction.type in its response, Amazon
	// Lex decides on the next course of action and selects an appropriate message
	// from the bot's configuration based on the current interaction context. For
	// example, if Amazon Lex isn't able to understand user input, it uses a clarification
	// prompt message.
	//
	// When you create an intent you can assign messages to groups. When messages
	// are assigned to groups Amazon Lex returns one message from each group in
	// the response. The message field is an escaped JSON string containing the
	// messages. For more information about the structure of the JSON string returned,
	// see msg-prompts-formats.
	//
	// If the Lambda function returns a message, Amazon Lex passes it to the client
	// in its response.
	Message *string `location:"header" locationName:"x-amz-lex-message" min:"1" type:"string" sensitive:"true"`

	// The format of the response message. One of the following values:
	//
	//    * PlainText - The message contains plain UTF-8 text.
	//
	//    * CustomPayload - The message is a custom format for the client.
	//
	//    * SSML - The message contains text formatted for voice output.
	//
	//    * Composite - The message contains an escaped JSON object containing one
	//    or more messages from the groups that messages were assigned to when the
	//    intent was created.
	MessageFormat MessageFormatType `location:"header" locationName:"x-amz-lex-message-format" type:"string" enum:"true"`

	// Map of key/value pairs representing the session-specific context information.
	SessionAttributes aws.JSONValue `location:"header" locationName:"x-amz-lex-session-attributes" type:"jsonvalue"`

	// If the dialogState value is ElicitSlot, returns the name of the slot for
	// which Amazon Lex is eliciting a value.
	SlotToElicit *string `location:"header" locationName:"x-amz-lex-slot-to-elicit" type:"string"`

	// Map of zero or more intent slots (name/value pairs) Amazon Lex detected from
	// the user input during the conversation. The field is base-64 encoded.
	//
	// Amazon Lex creates a resolution list containing likely values for a slot.
	// The value that it returns is determined by the valueSelectionStrategy selected
	// when the slot type was created or updated. If valueSelectionStrategy is set
	// to ORIGINAL_VALUE, the value provided by the user is returned, if the user
	// value is similar to the slot values. If valueSelectionStrategy is set to
	// TOP_RESOLUTION Amazon Lex returns the first value in the resolution list
	// or, if there is no resolution list, null. If you don't specify a valueSelectionStrategy,
	// the default is ORIGINAL_VALUE.
	Slots aws.JSONValue `location:"header" locationName:"x-amz-lex-slots" type:"jsonvalue"`
}

// String returns the string representation
func (s PostContentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PostContentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DialogState) > 0 {
		v := s.DialogState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-dialog-state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.InputTranscript != nil {
		v := *s.InputTranscript

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-input-transcript", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IntentName != nil {
		v := *s.IntentName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-intent-name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Message != nil {
		v := *s.Message

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-message", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.MessageFormat) > 0 {
		v := s.MessageFormat

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-message-format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.SessionAttributes != nil {
		v := s.SessionAttributes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-session-attributes", protocol.JSONValue{V: v, EscapeMode: protocol.Base64Escape}, metadata)
	}
	if s.SlotToElicit != nil {
		v := *s.SlotToElicit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-slot-to-elicit", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Slots != nil {
		v := s.Slots

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-lex-slots", protocol.JSONValue{V: v, EscapeMode: protocol.Base64Escape}, metadata)
	}
	// Skipping AudioStream Output type's body not valid.
	return nil
}

const opPostContent = "PostContent"

// PostContentRequest returns a request value for making API operation for
// Amazon Lex Runtime Service.
//
// Sends user input (text or speech) to Amazon Lex. Clients use this API to
// send text and audio requests to Amazon Lex at runtime. Amazon Lex interprets
// the user input using the machine learning model that it built for the bot.
//
// The PostContent operation supports audio input at 8kHz and 16kHz. You can
// use 8kHz audio to achieve higher speech recognition accuracy in telephone
// audio applications.
//
// In response, Amazon Lex returns the next message to convey to the user. Consider
// the following example messages:
//
//    * For a user input "I would like a pizza," Amazon Lex might return a response
//    with a message eliciting slot data (for example, PizzaSize): "What size
//    pizza would you like?".
//
//    * After the user provides all of the pizza order information, Amazon Lex
//    might return a response with a message to get user confirmation: "Order
//    the pizza?".
//
//    * After the user replies "Yes" to the confirmation prompt, Amazon Lex
//    might return a conclusion statement: "Thank you, your cheese pizza has
//    been ordered.".
//
// Not all Amazon Lex messages require a response from the user. For example,
// conclusion statements do not require a response. Some messages require only
// a yes or no response. In addition to the message, Amazon Lex provides additional
// context about the message in the response that you can use to enhance client
// behavior, such as displaying the appropriate client user interface. Consider
// the following examples:
//
//    * If the message is to elicit slot data, Amazon Lex returns the following
//    context information: x-amz-lex-dialog-state header set to ElicitSlot x-amz-lex-intent-name
//    header set to the intent name in the current context x-amz-lex-slot-to-elicit
//    header set to the slot name for which the message is eliciting information
//    x-amz-lex-slots header set to a map of slots configured for the intent
//    with their current values
//
//    * If the message is a confirmation prompt, the x-amz-lex-dialog-state
//    header is set to Confirmation and the x-amz-lex-slot-to-elicit header
//    is omitted.
//
//    * If the message is a clarification prompt configured for the intent,
//    indicating that the user intent is not understood, the x-amz-dialog-state
//    header is set to ElicitIntent and the x-amz-slot-to-elicit header is omitted.
//
// In addition, Amazon Lex also returns your application-specific sessionAttributes.
// For more information, see Managing Conversation Context (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html).
//
//    // Example sending a request using PostContentRequest.
//    req := client.PostContentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostContent
func (c *Client) PostContentRequest(input *PostContentInput) PostContentRequest {
	op := &aws.Operation{
		Name:       opPostContent,
		HTTPMethod: "POST",
		HTTPPath:   "/bot/{botName}/alias/{botAlias}/user/{userId}/content",
	}

	if input == nil {
		input = &PostContentInput{}
	}

	req := c.newRequest(op, input, &PostContentOutput{})
	req.Handlers.Sign.Remove(v4.SignRequestHandler)
	handler := v4.BuildNamedHandler("v4.CustomSignerHandler", v4.WithUnsignedPayload)
	req.Handlers.Sign.PushFrontNamed(handler)
	return PostContentRequest{Request: req, Input: input, Copy: c.PostContentRequest}
}

// PostContentRequest is the request type for the
// PostContent API operation.
type PostContentRequest struct {
	*aws.Request
	Input *PostContentInput
	Copy  func(*PostContentInput) PostContentRequest
}

// Send marshals and sends the PostContent API request.
func (r PostContentRequest) Send(ctx context.Context) (*PostContentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PostContentResponse{
		PostContentOutput: r.Request.Data.(*PostContentOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PostContentResponse is the response type for the
// PostContent API operation.
type PostContentResponse struct {
	*PostContentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PostContent request.
func (r *PostContentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
