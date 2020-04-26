---
layout: "gmailfilter"
page_title: "Gmail Filter: gmailfilter_label"
subcategory: "Gmail Filter Settings"
description: |-
  Get information about an existing Label.
---

# Data Source: gmailfilter_label

Get information about an existing Label.

## Example Usage

```hcl
data gmailfilter_label "label" {
  name = "your-label-name"
}
```
## Argument Reference

* `name` - (Required) The display name of the label.



## Attribute Reference

* `id` - The id of the Label.
* `background_color` - The background color represented as hex string `#RRGGBB` (ex #000000). This field is required in order to set the color of a label. See https://developers.google.com/gmail/api/v1/reference/users/labels for more details.
* `text_color` - The text color of the label, represented as hex string. This field is required in order to set the color of a label. See https://developers.google.com/gmail/api/v1/reference/users/labels for more details.
* `label_list_visibility` - The visibility of the label in the label list in the Gmail web interface. Acceptable values are: `labelHide` / `labelShow` / `labelShowIfUnread`.
* `message_list_visibility` - The visibility of messages with this label in the message list in the Gmail web interface. Acceptable values are: `hide` / `show`.
* `messages_total` - The total number of messages with the label.
* `messages_unread` - The number of unread messages with the label.
* `threads_total` - The total number of threads with the label.
* `threads_unread` - The number of unread threads with the label.
* `type` - The owner type for the label. User labels are created by the user and can be modified and deleted by the user and can be applied to any message or thread. System labels are internally created and cannot be added, modified, or deleted. System labels may be able to be applied to or removed from messages and threads under some circumstances but this is not guaranteed. For example, users can apply and remove the `INBOX` and `UNREAD` labels from messages and threads, but cannot apply or remove the `DRAFTS` or `SENT` labels from messages or threads.



