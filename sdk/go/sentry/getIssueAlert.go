// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sentry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIssueAlert(ctx *pulumi.Context, args *LookupIssueAlertArgs, opts ...pulumi.InvokeOption) (*LookupIssueAlertResult, error) {
	var rv LookupIssueAlertResult
	err := ctx.Invoke("sentry:index/getIssueAlert:getIssueAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIssueAlert.
type LookupIssueAlertArgs struct {
	InternalId   string `pulumi:"internalId"`
	Organization string `pulumi:"organization"`
	Project      string `pulumi:"project"`
}

// A collection of values returned by getIssueAlert.
type LookupIssueAlertResult struct {
	ActionMatch string                   `pulumi:"actionMatch"`
	Actions     []map[string]interface{} `pulumi:"actions"`
	Conditions  []map[string]interface{} `pulumi:"conditions"`
	Environment string                   `pulumi:"environment"`
	FilterMatch string                   `pulumi:"filterMatch"`
	Filters     []map[string]interface{} `pulumi:"filters"`
	Frequency   int                      `pulumi:"frequency"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	InternalId   string `pulumi:"internalId"`
	Name         string `pulumi:"name"`
	Organization string `pulumi:"organization"`
	Project      string `pulumi:"project"`
}

func LookupIssueAlertOutput(ctx *pulumi.Context, args LookupIssueAlertOutputArgs, opts ...pulumi.InvokeOption) LookupIssueAlertResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIssueAlertResult, error) {
			args := v.(LookupIssueAlertArgs)
			r, err := LookupIssueAlert(ctx, &args, opts...)
			var s LookupIssueAlertResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIssueAlertResultOutput)
}

// A collection of arguments for invoking getIssueAlert.
type LookupIssueAlertOutputArgs struct {
	InternalId   pulumi.StringInput `pulumi:"internalId"`
	Organization pulumi.StringInput `pulumi:"organization"`
	Project      pulumi.StringInput `pulumi:"project"`
}

func (LookupIssueAlertOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIssueAlertArgs)(nil)).Elem()
}

// A collection of values returned by getIssueAlert.
type LookupIssueAlertResultOutput struct{ *pulumi.OutputState }

func (LookupIssueAlertResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIssueAlertResult)(nil)).Elem()
}

func (o LookupIssueAlertResultOutput) ToLookupIssueAlertResultOutput() LookupIssueAlertResultOutput {
	return o
}

func (o LookupIssueAlertResultOutput) ToLookupIssueAlertResultOutputWithContext(ctx context.Context) LookupIssueAlertResultOutput {
	return o
}

func (o LookupIssueAlertResultOutput) ActionMatch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) string { return v.ActionMatch }).(pulumi.StringOutput)
}

func (o LookupIssueAlertResultOutput) Actions() pulumi.MapArrayOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) []map[string]interface{} { return v.Actions }).(pulumi.MapArrayOutput)
}

func (o LookupIssueAlertResultOutput) Conditions() pulumi.MapArrayOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) []map[string]interface{} { return v.Conditions }).(pulumi.MapArrayOutput)
}

func (o LookupIssueAlertResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o LookupIssueAlertResultOutput) FilterMatch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) string { return v.FilterMatch }).(pulumi.StringOutput)
}

func (o LookupIssueAlertResultOutput) Filters() pulumi.MapArrayOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) []map[string]interface{} { return v.Filters }).(pulumi.MapArrayOutput)
}

func (o LookupIssueAlertResultOutput) Frequency() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) int { return v.Frequency }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIssueAlertResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIssueAlertResultOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) string { return v.InternalId }).(pulumi.StringOutput)
}

func (o LookupIssueAlertResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIssueAlertResultOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) string { return v.Organization }).(pulumi.StringOutput)
}

func (o LookupIssueAlertResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueAlertResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIssueAlertResultOutput{})
}
