package gmailfilter

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	googleoauth "golang.org/x/oauth2/google"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
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

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, provider, terraformVersion)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData, p *schema.Provider, terraformVersion string) (interface{}, error) {
	config := Config{
		terraformVersion: terraformVersion,
	}

	if v, ok := d.GetOk("credentials"); ok {
		config.Credentials = v.(string)
	}
	if v, ok := d.GetOk("impersonated_user_email"); ok {
		config.ImpersonatedUserEmail = v.(string)
	}

	if err := config.LoadAndValidate(p.StopContext()); err != nil {
		return nil, err
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
