# Configure the GmailFilter Provider
provider "gmailfilter" {
  credentials             = file("account.json")
  impersonated_user_email = "foobar@example.com"
}