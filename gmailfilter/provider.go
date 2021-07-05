package gmailfilter

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	googleoauth "golang.org/x/oauth2/google"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"credentials": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"GOOGLE_CREDENTIALS",
					"GOOGLE_CLOUD_KEYFILE_JSON",
					"GCLOUD_KEYFILE_JSON",
				}, nil),
				ValidateFunc: validateCredentials,
				RequiredWith: []string{"impersonated_user_email"},
			},
			"impersonated_user_email": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"IMPERSONATED_USER_EMAIL",
				}, nil),
				RequiredWith: []string{"credentials"},
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"gmailfilter_filter": dataSourceGmailfilterFilter(),
			"gmailfilter_label":  dataSourceGmailfilterLabel(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"gmailfilter_filter": resourceGmailfilterFilter(),
			"gmailfilter_label":  resourceGmailfilterLabel(),
		},
	}

	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(ctx, d, provider, terraformVersion)
	}

	return provider
}

func providerConfigure(ctx context.Context, d *schema.ResourceData, p *schema.Provider, terraformVersion string) (interface{}, diag.Diagnostics) {
	config := Config{
		terraformVersion: terraformVersion,
	}

	if v, ok := d.GetOk("credentials"); ok {
		config.Credentials = v.(string)
	}
	if v, ok := d.GetOk("impersonated_user_email"); ok {
		config.ImpersonatedUserEmail = v.(string)
	}

	if err := config.LoadAndValidate(ctx); err != nil {
		return nil, diag.FromErr(err)
	}

	return &config, nil
}

func validateCredentials(v interface{}, k string) (warnings []string, errors []error) {
	if v == nil || v.(string) == "" {
		return
	}
	creds := v.(string)
	// if this is a path and we can stat it, assume it's ok
	if _, err := os.Stat(creds); err == nil {
		return
	}
	if _, err := googleoauth.CredentialsFromJSON(context.Background(), []byte(creds)); err != nil {
		errors = append(errors,
			fmt.Errorf("JSON credentials in %q are not valid: %s", creds, err))
	}

	return
}
