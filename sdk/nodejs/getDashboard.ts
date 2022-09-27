// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDashboard(args: GetDashboardArgs, opts?: pulumi.InvokeOptions): Promise<GetDashboardResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("sentry:index/getDashboard:getDashboard", {
        "internalId": args.internalId,
        "organization": args.organization,
    }, opts);
}

/**
 * A collection of arguments for invoking getDashboard.
 */
export interface GetDashboardArgs {
    internalId: string;
    organization: string;
}

/**
 * A collection of values returned by getDashboard.
 */
export interface GetDashboardResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly internalId: string;
    readonly organization: string;
    readonly title: string;
    readonly widgets: outputs.GetDashboardWidget[];
}

export function getDashboardOutput(args: GetDashboardOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDashboardResult> {
    return pulumi.output(args).apply(a => getDashboard(a, opts))
}

/**
 * A collection of arguments for invoking getDashboard.
 */
export interface GetDashboardOutputArgs {
    internalId: pulumi.Input<string>;
    organization: pulumi.Input<string>;
}
