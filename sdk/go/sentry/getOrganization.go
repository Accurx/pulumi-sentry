// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sentry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrganization(ctx *pulumi.Context, args *LookupOrganizationArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationResult, error) {
	var rv LookupOrganizationResult
	err := ctx.Invoke("sentry:index/getOrganization:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganization.
type LookupOrganizationArgs struct {
	Slug string `pulumi:"slug"`
}

// A collection of values returned by getOrganization.
type LookupOrganizationResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InternalId string `pulumi:"internalId"`
	Name       string `pulumi:"name"`
	Slug       string `pulumi:"slug"`
}

func LookupOrganizationOutput(ctx *pulumi.Context, args LookupOrganizationOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationResult, error) {
			args := v.(LookupOrganizationArgs)
			r, err := LookupOrganization(ctx, &args, opts...)
			var s LookupOrganizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationResultOutput)
}

// A collection of arguments for invoking getOrganization.
type LookupOrganizationOutputArgs struct {
	Slug pulumi.StringInput `pulumi:"slug"`
}

func (LookupOrganizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationArgs)(nil)).Elem()
}

// A collection of values returned by getOrganization.
type LookupOrganizationResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationResult)(nil)).Elem()
}

func (o LookupOrganizationResultOutput) ToLookupOrganizationResultOutput() LookupOrganizationResultOutput {
	return o
}

func (o LookupOrganizationResultOutput) ToLookupOrganizationResultOutputWithContext(ctx context.Context) LookupOrganizationResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrganizationResultOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.InternalId }).(pulumi.StringOutput)
}

func (o LookupOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOrganizationResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationResult) string { return v.Slug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationResultOutput{})
}
