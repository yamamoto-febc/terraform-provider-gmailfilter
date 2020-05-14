# Terraform Provider for Gmail Filter

![Test Status](https://github.com/yamamoto-febc/terraform-provider-gmailfilter/workflows/Tests/badge.svg)

- Terraform Website: https://terraform.io
- Documentation: see [website/docs](website/docs) dir

## Usage Example

```hcl
resource gmailfilter_filter "example" {
  criteria {
    # exclude_chats   = false
    from = "foobar@example.com"
    # has_attachment  = false
    # negated_query   = "from:someuser@example.com rfc822msgid: is:unread"
    # query           = "from:someuser@example.com rfc822msgid: is:unread"
    # size            = 1000
    # size_comparison = "larger"
    # subject         = "example"
    # to              = "example@"
  }
  action {
    add_label_ids    = [gmailfilter_label.example.id]
    remove_label_ids = [data.gmailfilter_label.INBOX.id]
  }
}

data gmailfilter_label "INBOX" {
  name = "INBOX"
}

resource gmailfilter_label "example" {
  name = "example"
}
```

## Requirements

- [Terraform](https://terraform.io) v0.12+

## Installation

### Setup the Provider Plugin

To install the provider plugin, please see the [Terraform documentation](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).
The plugin executables can download from [Releases](https://github.com/yamamoto-febc/terraform-provider-gmailfilter/releases/latest).

### Enable Gmail APIs

Follow the instruction at [Enable and disable APIs](https://support.google.com/googleapi/answer/6158841) to enable to the Gmail APIs.

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

- `GOOGLE_APPLICATION_CREDENTIALS`
- `GOOGLE_CREDENTIALS`
- `GOOGLE_CLOUD_KEYFILE_JSON`
- `GCLOUD_KEYFILE_JSON`
- `IMPERSONATED_USER_EMAIL`

### Using an Application Default Credential

First, you need to create credentials for the project.

https://console.cloud.google.com/apis/credentials?project=[project_ID]

Then, create an OAuth 2.0 client, and download the client secret file(JSON) to your local machine.

Finally, run the following command to use that credential to authenticate.

```bash
gcloud auth application-default login \
  --client-id-file=client_secret.json \
  --scopes \
https://www.googleapis.com/auth/gmail.labels,\
https://www.googleapis.com/auth/gmail.settings.basic
```

## Generate Terraform files from your existing infrastructure with Terraformer

[Terraformer](https://github.com/GoogleCloudPlatform/terraformer) includes support for this provider and can generate Terraform files from your existing infrastructure.

See the Terraformer [README](https://github.com/GoogleCloudPlatform/terraformer/blob/master/README.md#use-with-gmailfilter) for more information.

## Known Issues

Currently, the provider doesn't support managing the forwarding addresses.  
If you want to use forwarding in the `gmailfilter_filter` resource, please configure the forwarding addresses manually before using it.  
For details, see [Automatically forward Gmail messages to another account](https://support.google.com/mail/answer/10957).

## License

 `terraform-proivder-gmailfilter` Copyright (C) 2020 terraform-provider-gmailfilter authors.
 
  This project is published under [Mozilla Public License 2.0](LICENSE).
