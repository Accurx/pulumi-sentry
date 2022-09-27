// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sentry

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Team struct {
	pulumi.CustomResourceState

	HasAccess pulumi.BoolOutput `pulumi:"hasAccess"`
	// The internal ID for this team.
	InternalId pulumi.StringOutput `pulumi:"internalId"`
	IsMember   pulumi.BoolOutput   `pulumi:"isMember"`
	IsPending  pulumi.BoolOutput   `pulumi:"isPending"`
	// The name of the team.
	Name pulumi.StringOutput `pulumi:"name"`
	// The slug of the organization the team should be created for.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The optional slug for this team.
	Slug pulumi.StringOutput `pulumi:"slug"`
	// Use `internal_id` instead.
	//
	// Deprecated: Use `internal_id` instead.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewTeam registers a new resource with the given unique name, arguments, and options.
func NewTeam(ctx *pulumi.Context,
	name string, args *TeamArgs, opts ...pulumi.ResourceOption) (*Team, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	var resource Team
	err := ctx.RegisterResource("sentry:index/team:Team", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeam gets an existing Team resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamState, opts ...pulumi.ResourceOption) (*Team, error) {
	var resource Team
	err := ctx.ReadResource("sentry:index/team:Team", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Team resources.
type teamState struct {
	HasAccess *bool `pulumi:"hasAccess"`
	// The internal ID for this team.
	InternalId *string `pulumi:"internalId"`
	IsMember   *bool   `pulumi:"isMember"`
	IsPending  *bool   `pulumi:"isPending"`
	// The name of the team.
	Name *string `pulumi:"name"`
	// The slug of the organization the team should be created for.
	Organization *string `pulumi:"organization"`
	// The optional slug for this team.
	Slug *string `pulumi:"slug"`
	// Use `internal_id` instead.
	//
	// Deprecated: Use `internal_id` instead.
	TeamId *string `pulumi:"teamId"`
}

type TeamState struct {
	HasAccess pulumi.BoolPtrInput
	// The internal ID for this team.
	InternalId pulumi.StringPtrInput
	IsMember   pulumi.BoolPtrInput
	IsPending  pulumi.BoolPtrInput
	// The name of the team.
	Name pulumi.StringPtrInput
	// The slug of the organization the team should be created for.
	Organization pulumi.StringPtrInput
	// The optional slug for this team.
	Slug pulumi.StringPtrInput
	// Use `internal_id` instead.
	//
	// Deprecated: Use `internal_id` instead.
	TeamId pulumi.StringPtrInput
}

func (TeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamState)(nil)).Elem()
}

type teamArgs struct {
	// The name of the team.
	Name *string `pulumi:"name"`
	// The slug of the organization the team should be created for.
	Organization string `pulumi:"organization"`
	// The optional slug for this team.
	Slug *string `pulumi:"slug"`
}

// The set of arguments for constructing a Team resource.
type TeamArgs struct {
	// The name of the team.
	Name pulumi.StringPtrInput
	// The slug of the organization the team should be created for.
	Organization pulumi.StringInput
	// The optional slug for this team.
	Slug pulumi.StringPtrInput
}

func (TeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamArgs)(nil)).Elem()
}

type TeamInput interface {
	pulumi.Input

	ToTeamOutput() TeamOutput
	ToTeamOutputWithContext(ctx context.Context) TeamOutput
}

func (*Team) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (i *Team) ToTeamOutput() TeamOutput {
	return i.ToTeamOutputWithContext(context.Background())
}

func (i *Team) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamOutput)
}

// TeamArrayInput is an input type that accepts TeamArray and TeamArrayOutput values.
// You can construct a concrete instance of `TeamArrayInput` via:
//
//	TeamArray{ TeamArgs{...} }
type TeamArrayInput interface {
	pulumi.Input

	ToTeamArrayOutput() TeamArrayOutput
	ToTeamArrayOutputWithContext(context.Context) TeamArrayOutput
}

type TeamArray []TeamInput

func (TeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (i TeamArray) ToTeamArrayOutput() TeamArrayOutput {
	return i.ToTeamArrayOutputWithContext(context.Background())
}

func (i TeamArray) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamArrayOutput)
}

// TeamMapInput is an input type that accepts TeamMap and TeamMapOutput values.
// You can construct a concrete instance of `TeamMapInput` via:
//
//	TeamMap{ "key": TeamArgs{...} }
type TeamMapInput interface {
	pulumi.Input

	ToTeamMapOutput() TeamMapOutput
	ToTeamMapOutputWithContext(context.Context) TeamMapOutput
}

type TeamMap map[string]TeamInput

func (TeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (i TeamMap) ToTeamMapOutput() TeamMapOutput {
	return i.ToTeamMapOutputWithContext(context.Background())
}

func (i TeamMap) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMapOutput)
}

type TeamOutput struct{ *pulumi.OutputState }

func (TeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (o TeamOutput) ToTeamOutput() TeamOutput {
	return o
}

func (o TeamOutput) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return o
}

func (o TeamOutput) HasAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolOutput { return v.HasAccess }).(pulumi.BoolOutput)
}

// The internal ID for this team.
func (o TeamOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.InternalId }).(pulumi.StringOutput)
}

func (o TeamOutput) IsMember() pulumi.BoolOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolOutput { return v.IsMember }).(pulumi.BoolOutput)
}

func (o TeamOutput) IsPending() pulumi.BoolOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolOutput { return v.IsPending }).(pulumi.BoolOutput)
}

// The name of the team.
func (o TeamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The slug of the organization the team should be created for.
func (o TeamOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// The optional slug for this team.
func (o TeamOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

// Use `internal_id` instead.
//
// Deprecated: Use `internal_id` instead.
func (o TeamOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type TeamArrayOutput struct{ *pulumi.OutputState }

func (TeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (o TeamArrayOutput) ToTeamArrayOutput() TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) Index(i pulumi.IntInput) TeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Team {
		return vs[0].([]*Team)[vs[1].(int)]
	}).(TeamOutput)
}

type TeamMapOutput struct{ *pulumi.OutputState }

func (TeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (o TeamMapOutput) ToTeamMapOutput() TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return o
}

func (o TeamMapOutput) MapIndex(k pulumi.StringInput) TeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Team {
		return vs[0].(map[string]*Team)[vs[1].(string)]
	}).(TeamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamInput)(nil)).Elem(), &Team{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamArrayInput)(nil)).Elem(), TeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMapInput)(nil)).Elem(), TeamMap{})
	pulumi.RegisterOutputType(TeamOutput{})
	pulumi.RegisterOutputType(TeamArrayOutput{})
	pulumi.RegisterOutputType(TeamMapOutput{})
}
