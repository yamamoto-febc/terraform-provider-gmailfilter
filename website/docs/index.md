---
layout: "gmailfilter"
page_title: "Provider: Gmail Filter"
description: |-
  The Gmail Filter Provider is used to interact with the many resources supported by its APIs.
---

# Gmail Filter Provider

The Gmail Filter Provider is used to interact with the many resources supported by its APIs.

## Example Usage

```hcl
provider "gmailfilter" {
  credentials             = file("account.json")
  impersonated_user_email = "foobar@example.com"
}
```

## Authentication

There are two authentication methods available with this provider.

- Using a service account(G suite users only)
- Using an Application Default Credentials

### Using a service account(G suite users only)

Follow the instruction at [https://developers.google.com/identity/protocols/oauth2/service-account#creatinganaccount](https://developers.google.com/identity/protocols/oauth2/service-account#creatinganaccount).

Then, add `credentials` and `impersonated_user_email` to your tf file.

```tf
provider gmailfilter {
  credentials             = file("account.json")
  impersonated_user_email = "foobar@example.com"  
}
```

You can also use these environment variables:

- `GOOGLE_CREDENTIALS`
- `GOOGLE_CLOUD_KEYFILE_JSON`
- `GCLOUD_KEYFILE_JSON`
- `GOOGLE_APPLICATION_CREDENTIALS`
- `IMPERSONATED_USER_EMAIL`

### Using an Application Default Credential

First, you need to create credentials for the project.

https://console.cloud.google.com/apis/credentials?project=[project_ID]

Then, create an OAuth 2.0 client, and download the client secret file(JSON) to your local machine.

Finally, run the following command to use that credential to authenticate.

