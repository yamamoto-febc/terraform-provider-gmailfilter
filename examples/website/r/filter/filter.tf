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
