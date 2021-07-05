package gmailfilter

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGmailFilterDataSourceLabel_basic(t *testing.T) {
	resourceName := "data.gmailfilter_label.foobar"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGmailFilterDataSourceLabel_basic,
				Check: resource.ComposeTestCheckFunc(
					testCheckGmailFilterDataSourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "INBOX"),
					resource.TestCheckResourceAttr(resourceName, "background_color", ""),
					resource.TestCheckResourceAttr(resourceName, "text_color", ""),
					resource.TestCheckResourceAttr(resourceName, "label_list_visibility", "labelShow"),
					resource.TestCheckResourceAttr(resourceName, "message_list_visibility", "hide"),
					resource.TestCheckResourceAttrSet(resourceName, "messages_total"),
					resource.TestCheckResourceAttrSet(resourceName, "messages_unread"),
					resource.TestCheckResourceAttrSet(resourceName, "threads_total"),
					resource.TestCheckResourceAttrSet(resourceName, "threads_unread"),
					resource.TestCheckResourceAttr(resourceName, "type", "system"),
				),
			},
		},
	})
}

func TestAccGmailFilterDataSourceLabel_notExists(t *testing.T) {
	name := "data.gmailfilter_label.foobar"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGmailFilterDataSourceLabel_notExists,
				Check: resource.ComposeTestCheckFunc(
					testCheckGmailFilterDataSourceNotExists(name),
				),
				ExpectError: regexp.MustCompile(`no label with name=".*" found`),
				Destroy:     true,
			},
		},
	})
}

var testAccGmailFilterDataSourceLabel_basic = `
data "gmailfilter_label" "foobar" {
  name = "INBOX"
}`

var testAccCheckGmailFilterDataSourceLabel_notExists = `
data "gmailfilter_label" "foobar" {
  name = "not-exists-label-name"
}`
