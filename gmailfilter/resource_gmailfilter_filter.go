package gmailfilter

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"google.golang.org/api/gmail/v1"
)

func resourceGmailfilterFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceGmailfilterFilterCreate,
		Read:   resourceGmailfilterFilterRead,
		Delete: resourceGmailfilterFilterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_label_ids": {
							Type:         schema.TypeList,
							Elem:         &schema.Schema{Type: schema.TypeString},
							Optional:     true,
							AtLeastOneOf: []string{"action.0.forward", "action.0.remove_label_ids"},
							Description:  `List of labels to add to the message`,
						},
						"forward": {
							Type:         schema.TypeString,
							Optional:     true,
							AtLeastOneOf: []string{"action.0.add_label_ids", "action.0.remove_label_ids"},
							Description:  `Email address that the message should be forwarded to`,
						},
						"remove_label_ids": {
							Type:         schema.TypeList,
							Elem:         &schema.Schema{Type: schema.TypeString},
							Optional:     true,
							AtLeastOneOf: []string{"action.0.add_label_ids", "action.0.forward"},
							Description:  `List of labels to remove from the message`,
						},
					},
				},
			},
			"criteria": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclude_chats": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Whether the response should exclude chats`,
						},
						"from": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The sender's display name or email address`,
						},
						"has_attachment": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Whether the message has any attachment`,
						},
						"negated_query": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Only return messages not matching the specified query. Supports the same query format as the Gmail search box. For example, "from:someuser@example.com rfc822msgid: is:unread"`,
						},
						"query": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Only return messages matching the specified query. Supports the same query format as the Gmail search box. For example, "from:someuser@example.com rfc822msgid: is:unread"`,
						},
						"size": {
							Type:         schema.TypeInt,
							Optional:     true,
							RequiredWith: []string{"criteria.0.size_comparison"},
							Description:  `The size of the entire RFC822 message in bytes, including all headers and attachments`,
						},
						"size_comparison": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: []string{"criteria.0.size"},
							ValidateFunc: validation.StringInSlice([]string{"larger", "smaller", "unspecified"}, false),
							Description:  "How the message size in bytes should be in relation to the size field. Acceptable values are: `larger` / `smaller` / `unspecified`",
						},
						"subject": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Case-insensitive phrase found in the message's subject. Trailing and leading whitespace are be trimmed and adjacent spaces are collapsed`,
						},
						"to": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The recipient's display name or email address. Includes recipients in the "to", "cc", and "bcc" header fields. You can use simply the local part of the email address. For example, "example" and "example@" both match "example@gmail.com". This field is case-insensitive`,
						},
					},
				},
			},
		},
	}
}

func resourceGmailfilterFilterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	filter := &gmail.Filter{
		Action:   expandFilterAction(d),
		Criteria: expandFilterCriteria(d),
	}

	filter, err := config.gmailService.Users.Settings.Filters.Create(gmailUser, filter).Do()
	if err != nil {
		return fmt.Errorf("error creating filter: %s", err)
	}

	d.SetId(filter.Id)
	return resourceGmailfilterFilterRead(d, meta)
}

func resourceGmailfilterFilterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	filter, err := config.gmailService.Users.Settings.Filters.Get(gmailUser, d.Id()).Do()
	if err != nil {
		return handleNotFoundError(err, d, "Filter")
	}

	return setFilterValuesToState(d, filter)
}

func resourceGmailfilterFilterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	if _, err := config.gmailService.Users.Settings.Filters.Get(gmailUser, d.Id()).Do(); err != nil {
		return handleNotFoundError(err, d, "Filter")
	}

	if err := config.gmailService.Users.Settings.Filters.Delete(gmailUser, d.Id()).Do(); err != nil {
		return fmt.Errorf("error deleting filter[%s]: %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
