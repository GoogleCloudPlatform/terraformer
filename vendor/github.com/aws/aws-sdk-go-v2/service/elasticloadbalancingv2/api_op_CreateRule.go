// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateRuleInput struct {
	_ struct{} `type:"structure"`

	// The actions. Each rule must include exactly one of the following types of
	// actions: forward, fixed-response, or redirect, and it must be the last action
	// to be performed.
	//
	// If the action type is forward, you specify one or more target groups. The
	// protocol of the target group must be HTTP or HTTPS for an Application Load
	// Balancer. The protocol of the target group must be TCP, TLS, UDP, or TCP_UDP
	// for a Network Load Balancer.
	//
	// [HTTPS listeners] If the action type is authenticate-oidc, you authenticate
	// users through an identity provider that is OpenID Connect (OIDC) compliant.
	//
	// [HTTPS listeners] If the action type is authenticate-cognito, you authenticate
	// users through the user pools supported by Amazon Cognito.
	//
	// [Application Load Balancer] If the action type is redirect, you redirect
	// specified client requests from one URL to another.
	//
	// [Application Load Balancer] If the action type is fixed-response, you drop
	// specified client requests and return a custom HTTP response.
	//
	// Actions is a required field
	Actions []Action `type:"list" required:"true"`

	// The conditions. Each rule can include zero or one of the following conditions:
	// http-request-method, host-header, path-pattern, and source-ip, and zero or
	// more of the following conditions: http-header and query-string.
	//
	// Conditions is a required field
	Conditions []RuleCondition `type:"list" required:"true"`

	// The Amazon Resource Name (ARN) of the listener.
	//
	// ListenerArn is a required field
	ListenerArn *string `type:"string" required:"true"`

	// The rule priority. A listener can't have multiple rules with the same priority.
	//
	// Priority is a required field
	Priority *int64 `min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s CreateRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRuleInput"}

	if s.Actions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Actions"))
	}

	if s.Conditions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Conditions"))
	}

	if s.ListenerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ListenerArn"))
	}

	if s.Priority == nil {
		invalidParams.Add(aws.NewErrParamRequired("Priority"))
	}
	if s.Priority != nil && *s.Priority < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Priority", 1))
	}
	if s.Actions != nil {
		for i, v := range s.Actions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Actions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRuleOutput struct {
	_ struct{} `type:"structure"`

	// Information about the rule.
	Rules []Rule `type:"list"`
}

// String returns the string representation
func (s CreateRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRule = "CreateRule"

// CreateRuleRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Creates a rule for the specified listener. The listener must be associated
// with an Application Load Balancer.
//
// Rules are evaluated in priority order, from the lowest value to the highest
// value. When the conditions for a rule are met, its actions are performed.
// If the conditions for no rules are met, the actions for the default rule
// are performed. For more information, see Listener Rules (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#listener-rules)
// in the Application Load Balancers Guide.
//
// To view your current rules, use DescribeRules. To update a rule, use ModifyRule.
// To set the priorities of your rules, use SetRulePriorities. To delete a rule,
// use DeleteRule.
//
//    // Example sending a request using CreateRuleRequest.
//    req := client.CreateRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/CreateRule
func (c *Client) CreateRuleRequest(input *CreateRuleInput) CreateRuleRequest {
	op := &aws.Operation{
		Name:       opCreateRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRuleInput{}
	}

	req := c.newRequest(op, input, &CreateRuleOutput{})
	return CreateRuleRequest{Request: req, Input: input, Copy: c.CreateRuleRequest}
}

// CreateRuleRequest is the request type for the
// CreateRule API operation.
type CreateRuleRequest struct {
	*aws.Request
	Input *CreateRuleInput
	Copy  func(*CreateRuleInput) CreateRuleRequest
}

// Send marshals and sends the CreateRule API request.
func (r CreateRuleRequest) Send(ctx context.Context) (*CreateRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRuleResponse{
		CreateRuleOutput: r.Request.Data.(*CreateRuleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRuleResponse is the response type for the
// CreateRule API operation.
type CreateRuleResponse struct {
	*CreateRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRule request.
func (r *CreateRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
