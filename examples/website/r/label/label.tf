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