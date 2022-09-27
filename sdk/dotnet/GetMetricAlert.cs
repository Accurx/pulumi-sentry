// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sentry
{
    public static class GetMetricAlert
    {
        public static Task<GetMetricAlertResult> InvokeAsync(GetMetricAlertArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMetricAlertResult>("sentry:index/getMetricAlert:getMetricAlert", args ?? new GetMetricAlertArgs(), options.WithDefaults());

        public static Output<GetMetricAlertResult> Invoke(GetMetricAlertInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMetricAlertResult>("sentry:index/getMetricAlert:getMetricAlert", args ?? new GetMetricAlertInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMetricAlertArgs : global::Pulumi.InvokeArgs
    {
        [Input("internalId", required: true)]
        public string InternalId { get; set; } = null!;

        [Input("organization", required: true)]
        public string Organization { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetMetricAlertArgs()
        {
        }
        public static new GetMetricAlertArgs Empty => new GetMetricAlertArgs();
    }

    public sealed class GetMetricAlertInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("internalId", required: true)]
        public Input<string> InternalId { get; set; } = null!;

        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public GetMetricAlertInvokeArgs()
        {
        }
        public static new GetMetricAlertInvokeArgs Empty => new GetMetricAlertInvokeArgs();
    }


    [OutputType]
    public sealed class GetMetricAlertResult
    {
        public readonly string Aggregate;
        public readonly string Dataset;
        public readonly string Environment;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InternalId;
        public readonly string Name;
        public readonly string Organization;
        public readonly string Owner;
        public readonly string Project;
        public readonly string Query;
        public readonly double ResolveThreshold;
        public readonly int ThresholdType;
        public readonly double TimeWindow;
        public readonly ImmutableArray<Outputs.GetMetricAlertTriggerResult> Triggers;

        [OutputConstructor]
        private GetMetricAlertResult(
            string aggregate,

            string dataset,

            string environment,

            string id,

            string internalId,

            string name,

            string organization,

            string owner,

            string project,

            string query,

            double resolveThreshold,

            int thresholdType,

            double timeWindow,

            ImmutableArray<Outputs.GetMetricAlertTriggerResult> triggers)
        {
            Aggregate = aggregate;
            Dataset = dataset;
            Environment = environment;
            Id = id;
            InternalId = internalId;
            Name = name;
            Organization = organization;
            Owner = owner;
            Project = project;
            Query = query;
            ResolveThreshold = resolveThreshold;
            ThresholdType = thresholdType;
            TimeWindow = timeWindow;
            Triggers = triggers;
        }
    }
}
