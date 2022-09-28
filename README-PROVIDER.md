# Sentry Resource Provider

The Sentry Resource Provider lets you manage [Sentry](https://sentry.io/welcome/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/sentry
```

or `yarn`:

```bash
yarn add @pulumi/sentry
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_sentry
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-sentry/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Sentry
```

## Configuration

The following configuration points are available for the `sentry` provider:

`Required:`

- `sentry:token` (environment: `SENTRY_AUTH_TOKEN`) - the authentication token used to connect to `sentry`

`Optional:`

- `sentry:base_url` (environment: `SENTRY_BASE_URL`) - the base URL for sentry, defaulted to `https://sentry.io/api/`

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/sentry/api-docs/).
