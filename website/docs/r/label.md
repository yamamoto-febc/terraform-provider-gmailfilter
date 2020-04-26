---
layout: "gmailfilter"
page_title: "Gmail Filter: gmailfilter_label"
subcategory: "Gmail Filter Settings"
description: |-
  Manages a Gmail Filter Label.
---

# gmailfilter_label

Manages a Gmail Filter Label.

## Example Usage

```hcl
# top-level label
resource gmailfilter_label "label" {
  name                    = "label1"
  # background_color        = "#000000"
  # text_color              = "#000000"
  # label_list_visibility   = "labelShow" # must be one of [ labelHide / labelShow(default) / labelShowIfUnread ]
  # message_list_visibility = "show"      # must be one of [ show(default) / hide ]
}

# sub label
resource gmailfilter_label "sublabel" {
  name = "${gmailfilter_label.label.name}/label2"
}
```
## Argument Reference

* `name` - (Required) The display name of the label.
* `background_color` - (Optional) The background color represented as hex string `#RRGGBB` (ex #000000). This field is required in order to set the color of a label. See https://developers.google.com/gmail/api/v1/reference/users/labels for more details.
* `text_color` - (Optional) The text color of the label, represented as hex string. This field is required in order to set the color of a label. See https://developers.google.com/gmail/api/v1/reference/users/labels for more details.
* `label_list_visibility` - (Optional) The visibility of the label in the label list in the Gmail web interface. Acceptable values are: `labelHide` / `labelShow` / `labelShowIfUnread`. Default:`labelShow`.
* `message_list_visibility` - (Optional) The visibility of messages with this label in the message list in the Gmail web interface. Acceptable values are: `hide` / `show`. Default:`show`.



## Attribute Reference

* `id` - The id of the Label.
* `messages_total` - The total number of messages with the label.
* `messages_unread` - The number of unread messages with the label.
* `threads_total` - The total number of threads with the label.
* `threads_unread` - The number of unread threads with the label.
* `type` - The owner type for the label. User labels are created by the user and can be modified and deleted by the user and can be applied to any message or thread. System labels are internally created and cannot be added, modified, or deleted. System labels may be able to be applied to or removed from messages and threads under some circumstances but this is not guaranteed. For example, users can apply and remove the `INBOX` and `UNREAD` labels from messages and threads, but cannot apply or remove the `DRAFTS` or `SENT` labels from messages or threads.



