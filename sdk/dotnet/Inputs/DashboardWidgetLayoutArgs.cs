// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sentry.Inputs
{

    public sealed class DashboardWidgetLayoutArgs : global::Pulumi.ResourceArgs
    {
        [Input("h", required: true)]
        public Input<int> H { get; set; } = null!;

        [Input("minH", required: true)]
        public Input<int> MinH { get; set; } = null!;

        [Input("w", required: true)]
        public Input<int> W { get; set; } = null!;

        [Input("x", required: true)]
        public Input<int> X { get; set; } = null!;

        [Input("y", required: true)]
        public Input<int> Y { get; set; } = null!;

        public DashboardWidgetLayoutArgs()
        {
        }
        public static new DashboardWidgetLayoutArgs Empty => new DashboardWidgetLayoutArgs();
    }
}
