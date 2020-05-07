package gmailfilter

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/gmail/v1"
)

func dataSourceGmailfilterLabel() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGmailfilterLabelRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The display name of the label`,
			},
			"background_color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The background color represented as hex string #RRGGBB (ex #000000). This field is required in order to set the color of a label. See https://developers.google.com/gmail/api/v1/reference/users/labels for more details`,
			},
			"text_color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The text color of the label, represented as hex string. This field is required in order to set the color of a label. See https://developers.google.com/gmail/api/v1/reference/users/labels for more details`,
			},
			"label_list_visibility": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The visibility of the label in the label list in the Gmail web interface. Acceptable values are: `labelHide` / `labelShow` / `labelShowIfUnread`",
			},
			"message_list_visibility": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The visibility of messages with this label in the message list in the Gmail web interface. Acceptable values are: `hide` / `show`",
			},
			"messages_total": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The total number of messages with the label`,
			},
			"messages_unread": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of unread messages with the label`,
			},
			"threads_total": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The total number of threads with the label`,
			},
			"threads_unread": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of unread threads with the label`,
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The owner type for the label. User labels are created by the user and can be modified and deleted by the user and can be applied to any message or thread. System labels are internally created and cannot be added, modified, or deleted. System labels may be able to be applied to or removed from messages and threads under some circumstances but this is not guaranteed. For example, users can apply and remove the `INBOX` and `UNREAD` labels from messages and threads, but cannot apply or remove the `DRAFTS` or `SENT` labels from messages or threads",
			},
		},
	}
}

func dataSourceGmailfilterLabelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	name := d.Get("name").(string)

	res, err := config.gmailService.Users.Labels.List(gmailUser).Do()
	if err != nil {
		return handleNotFoundError(err, d, "Label")
	}

	var label *gmail.Label
	for _, l := range res.Labels {
		if l.Name == name {
			label = l
			break
		}
	}

	if label == nil {
		d.SetId("")
		log.Print("[WARN] Removing Label because it's gone")
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("no label with name=%q found", name))}
	}

	d.SetId(label.Id)
	return setLabelValuesToState(ctx, d, label)
}
