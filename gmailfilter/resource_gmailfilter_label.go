package gmailfilter

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGmailfilterLabel() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGmailfilterLabelCreate,
		ReadContext:   resourceGmailfilterLabelRead,
		UpdateContext: resourceGmailfilterLabelUpdate,
		DeleteContext: resourceGmailfilterLabelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The display name of the label`,
			},
			"background_color": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"text_color"},
				ValidateFunc: validation.StringInSlice(validLabelColors, false),
				Description:  `The background color represented as hex string #RRGGBB (ex #000000). This field is required in order to set the color of a label. See https://developers.google.com/gmail/api/v1/reference/users/labels for more details`,
			},
			"text_color": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"background_color"},
				ValidateFunc: validation.StringInSlice(validLabelColors, false),
				Description:  `The text color of the label, represented as hex string. This field is required in order to set the color of a label. See https://developers.google.com/gmail/api/v1/reference/users/labels for more details`,
			},
			"label_list_visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "labelShow",
				ValidateFunc: validation.StringInSlice([]string{"labelHide", "labelShow", "labelShowIfUnread"}, false),
				Description:  "The visibility of the label in the label list in the Gmail web interface. Acceptable values are: `labelHide` / `labelShow` / `labelShowIfUnread`",
			},
			"message_list_visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "show",
				ValidateFunc: validation.StringInSlice([]string{"show", "hide"}, false),
				Description:  "The visibility of messages with this label in the message list in the Gmail web interface. Acceptable values are: `hide` / `show`",
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

func resourceGmailfilterLabelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	label := expandLabel(d)

	label, err := config.gmailService.Users.Labels.Create(gmailUser, label).Do()
	if err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("error creating label: %s", err))}
	}

	d.SetId(label.Id)
	return resourceGmailfilterLabelRead(ctx, d, meta)
}

func resourceGmailfilterLabelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	label, err := config.gmailService.Users.Labels.Get(gmailUser, d.Id()).Do()
	if err != nil {
		return handleNotFoundError(err, d, "Label")
	}

	return setLabelValuesToState(ctx, d, label)
}

func resourceGmailfilterLabelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	if _, err := config.gmailService.Users.Labels.Get(gmailUser, d.Id()).Do(); err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("error updating label[%s]: %s", d.Id(), err))}
	}

	label := expandLabel(d)
	if _, err := config.gmailService.Users.Labels.Update(gmailUser, d.Id(), label).Do(); err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("error updating label[%s]: %s", d.Id(), err))}
	}

	return resourceGmailfilterLabelRead(ctx, d, meta)
}

func resourceGmailfilterLabelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	if _, err := config.gmailService.Users.Labels.Get(gmailUser, d.Id()).Do(); err != nil {
		return handleNotFoundError(err, d, "Label")
	}

	if err := config.gmailService.Users.Labels.Delete(gmailUser, d.Id()).Do(); err != nil {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("error deleting label[%s]: %s", d.Id(), err))}
	}

	d.SetId("")
	return nil
}
