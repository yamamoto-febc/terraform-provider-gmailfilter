package gmailfilter

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/gmail/v1"
)

func setForwardingAddressValuesToState(d *schema.ResourceData, forwardingAddress *gmail.ForwardingAddress) error {
	d.Set("forwarding_email", forwardingAddress.ForwardingEmail)
	d.Set("verification_status", forwardingAddress.VerificationStatus)
	return nil
}

func expandForwardingAddress(d *schema.ResourceData) *gmail.ForwardingAddress {
	return &gmail.ForwardingAddress{
		ForwardingEmail:    d.Get("forwarding_email").(string),
		VerificationStatus: d.Get("verification_status").(string),
	}
}
