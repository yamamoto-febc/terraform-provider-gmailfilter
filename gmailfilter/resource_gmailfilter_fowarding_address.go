package gmailfilter

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGmailfilterForwardingAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGmailfilterForwardingAddressCreate,
		ReadContext:   resourceGmailfilterForwardingAddressRead,
		DeleteContext: resourceGmailfilterForwardingAddressDelete,
		Schema: map[string]*schema.Schema{
			"forwarding_email": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `An email address to which messages can be forwarded. See https://developers.google.com/gmail/api/reference/rest/v1/users.settings.forwardingAddresses for more details`,
			},
			"verification_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "verificationStatusUnspecified",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"verificationStatusUnspecified", "accepted", "pending"}, false),
				Description:  "Indicates whether this address has been verified and is usable for forwarding. Acceptable values are: `labelHide` / `labelShow` / `labelShowIfUnread`",
			},
		},
	}
}

func resourceGmailfilterForwardingAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	forwardingAddress := expandForwardingAddress(d)

	forwardingAddress, err := config.gmailService.Users.Settings.ForwardingAddresses.Create(gmailUser, forwardingAddress).Do()
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating forwarding_address: %s", err))
	}

	d.SetId(forwardingAddress.ForwardingEmail)
	return resourceGmailfilterForwardingAddressRead(ctx, d, meta)
}

func resourceGmailfilterForwardingAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	forwardingAddress, err := config.gmailService.Users.Settings.ForwardingAddresses.Get(gmailUser, d.Id()).Do()
	if err != nil {
		return diag.FromErr(handleNotFoundError(err, d, "forwarding_address"))
	}

	return diag.FromErr(setForwardingAddressValuesToState(d, forwardingAddress))
}

func resourceGmailfilterForwardingAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	if _, err := config.gmailService.Users.Settings.ForwardingAddresses.Get(gmailUser, d.Id()).Do(); err != nil {
		return diag.FromErr(handleNotFoundError(err, d, "forwarding_address"))
	}

	if err := config.gmailService.Users.Settings.ForwardingAddresses.Delete(gmailUser, d.Id()).Do(); err != nil {
		return diag.FromErr(fmt.Errorf("error deleting forwarding_address[%s]: %s", d.Id(), err))
	}

	d.SetId("")
	return nil
}
