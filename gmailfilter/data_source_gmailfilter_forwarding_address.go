package gmailfilter

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/gmail/v1"
)

func dataSourceGmailfilterForwardingAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGmailfilterForwardingAddressRead,
		Schema: map[string]*schema.Schema{
			"forwarding_email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `An email address to which messages can be forwarded. See https://developers.google.com/gmail/api/reference/rest/v1/users.settings.forwardingAddresses for more details`,
			},
			"verification_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "verificationStatusUnspecified",
				Description: "Indicates whether this address has been verified and is usable for forwarding. Acceptable values are: `verificationStatusUnspecified` / `accepted` / `pending`",
			},
		},
	}
}

func dataSourceGmailfilterForwardingAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	forwardingEmail := d.Get("forwarding_email").(string)

	res, err := config.gmailService.Users.Settings.ForwardingAddresses.List(gmailUser).Do()
	if err != nil {
		return diag.FromErr(handleNotFoundError(err, d, "forwarding_address"))
	}

	var forwardingAddress *gmail.ForwardingAddress
	for _, f := range res.ForwardingAddresses {
		if f.ForwardingEmail == forwardingEmail {
			forwardingAddress = f
			break
		}
	}

	if forwardingAddress == nil {
		d.SetId("")
		log.Print("[WARN] Removing forwarding_address because it's gone")
		return diag.FromErr(fmt.Errorf("no forwarding_address with forwarding_email=%q found", forwardingEmail))
	}

	d.SetId(forwardingAddress.ForwardingEmail)
	return diag.FromErr(setForwardingAddressValuesToState(d, forwardingAddress))
}
