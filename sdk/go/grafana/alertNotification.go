// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The alert notification resource allows an alert notification channel to be created on a Grafana server.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-grafana/blob/master/website/docs/r/alert_notification.html.markdown.
type AlertNotification struct {
	s *pulumi.ResourceState
}

// NewAlertNotification registers a new resource with the given unique name, arguments, and options.
func NewAlertNotification(ctx *pulumi.Context,
	name string, args *AlertNotificationArgs, opts ...pulumi.ResourceOpt) (*AlertNotification, error) {
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["isDefault"] = nil
		inputs["name"] = nil
		inputs["settings"] = nil
		inputs["type"] = nil
	} else {
		inputs["isDefault"] = args.IsDefault
		inputs["name"] = args.Name
		inputs["settings"] = args.Settings
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("grafana:index/alertNotification:AlertNotification", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AlertNotification{s: s}, nil
}

// GetAlertNotification gets an existing AlertNotification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertNotification(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AlertNotificationState, opts ...pulumi.ResourceOpt) (*AlertNotification, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["isDefault"] = state.IsDefault
		inputs["name"] = state.Name
		inputs["settings"] = state.Settings
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("grafana:index/alertNotification:AlertNotification", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AlertNotification{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AlertNotification) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AlertNotification) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Is this the default channel for all your alerts.
func (r *AlertNotification) IsDefault() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["isDefault"])
}

// The name of the alert notification channel.
func (r *AlertNotification) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Additional settings, for full reference lookup [Grafana HTTP API documentation](http://docs.grafana.org/http_api/alerting).
func (r *AlertNotification) Settings() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["settings"])
}

// The type of the alert notification channel.
func (r *AlertNotification) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering AlertNotification resources.
type AlertNotificationState struct {
	// Is this the default channel for all your alerts.
	IsDefault interface{}
	// The name of the alert notification channel.
	Name interface{}
	// Additional settings, for full reference lookup [Grafana HTTP API documentation](http://docs.grafana.org/http_api/alerting).
	Settings interface{}
	// The type of the alert notification channel.
	Type interface{}
}

// The set of arguments for constructing a AlertNotification resource.
type AlertNotificationArgs struct {
	// Is this the default channel for all your alerts.
	IsDefault interface{}
	// The name of the alert notification channel.
	Name interface{}
	// Additional settings, for full reference lookup [Grafana HTTP API documentation](http://docs.grafana.org/http_api/alerting).
	Settings interface{}
	// The type of the alert notification channel.
	Type interface{}
}