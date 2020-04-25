---
layout: "gmailfilter"
page_title: "Gmail Filter: gmailfilter_filter"
subcategory: "Gmail Filter Settings"
description: |-
  Get information about an existing Filter.
---

# Data Source: gmailfilter_filter

Get information about an existing Filter.

## Example Usage

```hcl
data gmailfilter_filter "filter" {
  filter_id = "your-filter-id"
}
```
## Argument Reference

* `filter_id` - (Required) The filter.



## Attribute Reference

* `id` - The id of the Filter.
* `action` - A list of `action` blocks as defined below.
* `criteria` - A list of `criteria` blocks as defined below.


---

A `action` block exports the following:

* `add_label_ids` - List of labels to add to the message.
* `forward` - Email address that the message should be forwarded to.
* `remove_label_ids` - List of labels to remove from the message.

---

A `criteria` block exports the following:

* `exclude_chats` - Whether the response should exclude chats.
* `from` - The sender's display name or email address.
* `has_attachment` - Whether the message has any attachment.
* `negated_query` - Only return messages not matching the specified query. Supports the same query format as the Gmail search box. For example, "from:someuser@example.com rfc822msgid: is:unread".
* `query` - Only return messages matching the specified query. Supports the same query format as the Gmail search box. For example, "from:someuser@example.com rfc822msgid: is:unread".
* `size` - The size of the entire RFC822 message in bytes, including all headers and attachments.
* `size_comparison` - How the message size in bytes should be in relation to the size field.
* `subject` - Case-insensitive phrase found in the message's subject. Trailing and leading whitespace are be trimmed and adjacent spaces are collapsed.
* `to` - The recipient's display name or email address. Includes recipients in the "to", "cc", and "bcc" header fields. You can use simply the local part of the email address. For example, "example" and "example@" both match "example@gmail.com". This field is case-insensitive.


