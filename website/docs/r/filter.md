---
layout: "gmailfilter"
page_title: "Gmail Filter: gmailfilter_filter"
subcategory: "Gmail Filter Settings"
description: |-
  Manages a Gmail Filter Filter.
---

# gmailfilter_filter

Manages a Gmail Filter Filter.

## Example Usage

```hcl
resource gmailfilter_filter "example" {
  criteria {
    # exclude_chats   = false
    from = "foobar@example.com"
    # has_attachment  = false
    # negated_query   = "from:someuser@example.com rfc822msgid: is:unread"
    # query           = "from:someuser@example.com rfc822msgid: is:unread"
    # size            = 1000
    # size_comparison = "larger"
    # subject         = "example"
    # to              = "example@"
  }
  action {
    add_label_ids    = [gmailfilter_label.example.id]
    remove_label_ids = [data.gmailfilter_label.INBOX.id]
    # forward          = "forward-destination@example.com"
  }
}

data gmailfilter_label "INBOX" {
  name = "INBOX"
}

resource gmailfilter_label "example" {
  name = "example"
}

```
## Argument Reference

* `action` - (Required) A `action` block as defined below. Changing this forces a new resource to be created.
* `criteria` - (Required) An `criteria` block as defined below. Changing this forces a new resource to be created.


---

A `action` block supports the following:

* `add_label_ids` - (Optional) List of labels to add to the message.
* `remove_label_ids` - (Optional) List of labels to remove from the message.
* `forward` - (Optional) Email address that the message should be forwarded to.

~> Currently, the provider doesn't support managing the forwarding addresses.  
   If you want to use `forward` attribute in `action` block, please configure the forwarding addresses manually before using it.  
   For details, see [Automatically forward Gmail messages to another account](https://support.google.com/mail/answer/10957).


---

A `criteria` block supports the following:

* `exclude_chats` - (Optional) Whether the response should exclude chats.
* `from` - (Optional) The sender's display name or email address.
* `has_attachment` - (Optional) Whether the message has any attachment.
* `negated_query` - (Optional) Only return messages not matching the specified query. Supports the same query format as the Gmail search box. For example, "from:someuser@example.com rfc822msgid: is:unread".
* `query` - (Optional) Only return messages matching the specified query. Supports the same query format as the Gmail search box. For example, "from:someuser@example.com rfc822msgid: is:unread".
* `size` - (Optional) The size of the entire RFC822 message in bytes, including all headers and attachments.
* `size_comparison` - (Optional) How the message size in bytes should be in relation to the size field. Acceptable values are: `larger` / `smaller` / `unspecified`.
* `subject` - (Optional) Case-insensitive phrase found in the message's subject. Trailing and leading whitespace are be trimmed and adjacent spaces are collapsed.
* `to` - (Optional) The recipient's display name or email address. Includes recipients in the "to", "cc", and "bcc" header fields. You can use simply the local part of the email address. For example, "example" and "example@" both match "example@gmail.com". This field is case-insensitive.


## Attribute Reference

* `id` - The id of the Filter.



