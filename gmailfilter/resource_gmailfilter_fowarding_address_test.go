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

func TestAccGmailFilterForwardingAddress_basic(t *testing.T) {
	resourceName := "gmailfilter_forwarding_address.foobar"
	rand := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	var forwardingAddress gmail.ForwardingAddress
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: resource.ComposeTestCheckFunc(
			testCheckGmailFilterForwardingAddressDestroy,
		),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprint(testAccGmailFilterForwardingAddress_basic),
				Check: resource.ComposeTestCheckFunc(
					testCheckGmailFilterForwardingAddressExists(resourceName, &forwardingAddress),
					resource.TestCheckResourceAttr(resourceName, "forwarding_eamil", rand+"example.com"),
					resource.TestCheckResourceAttr(resourceName, "verification_status", "accepted"),
				),
			},
		},
	})
}

func testCheckGmailFilterForwardingAddressExists(n string, forwardingAddress *gmail.ForwardingAddress) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("no resources found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("no forwardingEmail is set")
		}

		config := testAccProvider.Meta().(*Config)
		found, err := config.gmailService.Users.Settings.ForwardingAddresses.Get(gmailUser, rs.Primary.ID).Do()
		if err != nil {
			return err
		}

		*forwardingAddress = *found
		return nil
	}
}

func testCheckGmailFilterForwardingAddressDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "gmailfilter_forwarding_address" {
			continue
		}
		if rs.Primary.ID == "" {
			continue
		}

		_, err := config.gmailService.Users.Settings.ForwardingAddresses.Get(gmailUser, rs.Primary.ID).Do()
		if err == nil {
			return fmt.Errorf("label[%s] still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccGmailFilterForwardingAddress_basic = `
resource gmailfilter_forwarding_address "foobar" {
    forwarding_email = "%s@example.com"
    verification_status = "accepted"
}
`
