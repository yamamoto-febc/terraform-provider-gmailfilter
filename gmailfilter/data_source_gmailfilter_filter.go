package gmailfilter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGmailfilterFilter() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGmailfilterFilterRead,
		Schema: map[string]*schema.Schema{
			"filter_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The filter",
			},
			"action": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_label_ids": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
							Description: `List of labels to add to the message`,
						},
						"forward": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Email address that the message should be forwarded to`,
						},
						"remove_label_ids": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
							Description: `List of labels to remove from the message`,
						},
					},
				},
			},
			"criteria": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclude_chats": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: `Whether the response should exclude chats`,
						},
						"from": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The sender's display name or email address`,
						},
						"has_attachment": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: `Whether the message has any attachment`,
						},
						"negated_query": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Only return messages not matching the specified query. Supports the same query format as the Gmail search box. For example, "from:someuser@example.com rfc822msgid: is:unread"`,
						},
						"query": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Only return messages matching the specified query. Supports the same query format as the Gmail search box. For example, "from:someuser@example.com rfc822msgid: is:unread"`,
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The size of the entire RFC822 message in bytes, including all headers and attachments`,
						},
						"size_comparison": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `How the message size in bytes should be in relation to the size field`,
						},
						"subject": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Case-insensitive phrase found in the message's subject. Trailing and leading whitespace are be trimmed and adjacent spaces are collapsed`,
						},
						"to": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The recipient's display name or email address. Includes recipients in the "to", "cc", and "bcc" header fields. You can use simply the local part of the email address. For example, "example" and "example@" both match "example@gmail.com". This field is case-insensitive`,
						},
					},
				},
			},
		},
	}
}

func dataSourceGmailfilterFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	id := d.Get("filter_id").(string)

	filter, err := config.gmailService.Users.Settings.Filters.Get(gmailUser, id).Do()
	if err != nil {
		return diag.FromErr(handleNotFoundError(err, d, "Filter"))
	}

	d.SetId(id)
	d.Set("filter_id", id)
	return diag.FromErr(setFilterValuesToState(d, filter))
}
