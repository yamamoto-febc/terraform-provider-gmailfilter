package gmailfilter

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"google.golang.org/api/gmail/v1"
)

func TestAccGmailFilterFilter_basic(t *testing.T) {
	resourceName := "gmailfilter_filter.foobar"
	rand := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	var label gmail.Filter
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: resource.ComposeTestCheckFunc(
			testCheckGmailFilterFilterDestroy,
		),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccGmailFilterFilter_basic, rand, rand),
				Check: resource.ComposeTestCheckFunc(
					testCheckGmailFilterFilterExists(resourceName, &label),
					resource.TestCheckResourceAttr(resourceName, "criteria.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.exclude_chats", "false"),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.from", rand+"@example.com"),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.has_attachment", "false"),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.negated_query", "from:someuser@example.com"),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.query", "from:someuser@example.com"),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.size", "1000"),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.size_comparison", "larger"),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.subject", "example"),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.to", "example@"),
					resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.add_label_ids.#", "1"),
					resource.TestCheckResourceAttrPair(
						resourceName, "action.0.add_label_ids.0",
						"gmailfilter_label.foobar", "id",
					),
					resource.TestCheckResourceAttr(resourceName, "action.0.remove_label_ids.#", "1"),
					resource.TestCheckResourceAttrPair(
						resourceName, "action.0.remove_label_ids.0",
						"data.gmailfilter_label.INBOX", "id",
					),
					// Note:
					// Current version doesn't support configuring the forwarding address.
					// If you would like to test this, please configure it manually on the Developer Console
					//resource.TestCheckResourceAttr(resourceName, "action.0.forward", "destination@example.com"),
				),
			},
		},
	})
}

func testCheckGmailFilterFilterExists(n string, label *gmail.Filter) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("no resources found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("no label ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		found, err := config.gmailService.Users.Settings.Filters.Get(gmailUser, rs.Primary.ID).Do()
		if err != nil {
			return err
		}

		*label = *found
		return nil
	}
}

func testCheckGmailFilterFilterDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "gmailfilter_filter" {
			continue
		}
		if rs.Primary.ID == "" {
			continue
		}

		_, err := config.gmailService.Users.Settings.Filters.Get(gmailUser, rs.Primary.ID).Do()
		if err == nil {
			return fmt.Errorf("label[%s] still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccGmailFilterFilter_basic = `
data gmailfilter_label "INBOX" {
  name = "INBOX"
}
resource gmailfilter_label "foobar" {
  name = "%s"
}

resource gmailfilter_filter "foobar" {
  criteria {
    exclude_chats   = false
    from            = "%s@example.com"
    has_attachment  = false
    negated_query   = "from:someuser@example.com"
    query           = "from:someuser@example.com"
    size            = 1000
    size_comparison = "larger"
    subject         = "example"
    to              = "example@"
  }
  action {
    add_label_ids    = [gmailfilter_label.foobar.id]
    remove_label_ids = [data.gmailfilter_label.INBOX.id]
    # forward          = "forwarding@example.com"
  }
}
`
