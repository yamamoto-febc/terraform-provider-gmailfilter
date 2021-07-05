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

func TestAccGmailFilterLabel_basic(t *testing.T) {
	resourceName := "gmailfilter_label.foobar"
	rand := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	var label gmail.Label
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: resource.ComposeTestCheckFunc(
			testCheckGmailFilterLabelDestroy,
		),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccGmailFilterLabel_basic, rand),
				Check: resource.ComposeTestCheckFunc(
					testCheckGmailFilterLabelExists(resourceName, &label),
					resource.TestCheckResourceAttr(resourceName, "name", rand),
					resource.TestCheckResourceAttr(resourceName, "background_color", "#000000"),
					resource.TestCheckResourceAttr(resourceName, "text_color", "#000000"),
					resource.TestCheckResourceAttr(resourceName, "label_list_visibility", "labelHide"),
					resource.TestCheckResourceAttr(resourceName, "message_list_visibility", "hide"),
				),
			},
			{
				Config: fmt.Sprintf(testAccGmailFilterLabel_update, rand),
				Check: resource.ComposeTestCheckFunc(
					testCheckGmailFilterLabelExists(resourceName, &label),
					resource.TestCheckResourceAttr(resourceName, "name", rand),
					resource.TestCheckResourceAttr(resourceName, "background_color", "#ffffff"),
					resource.TestCheckResourceAttr(resourceName, "text_color", "#ffffff"),
					resource.TestCheckResourceAttr(resourceName, "label_list_visibility", "labelShowIfUnread"),
					resource.TestCheckResourceAttr(resourceName, "message_list_visibility", "show"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     rand,
				ImportStateVerify: true,
			},
		},
	})
}

func testCheckGmailFilterLabelExists(n string, label *gmail.Label) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("no resources found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("no label ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		found, err := config.gmailService.Users.Labels.Get(gmailUser, rs.Primary.ID).Do()
		if err != nil {
			return err
		}

		*label = *found
		return nil
	}
}

func testCheckGmailFilterLabelDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "gmailfilter_label" {
			continue
		}
		if rs.Primary.ID == "" {
			continue
		}

		_, err := config.gmailService.Users.Labels.Get(gmailUser, rs.Primary.ID).Do()
		if err == nil {
			return fmt.Errorf("label[%s] still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccGmailFilterLabel_basic = `
resource gmailfilter_label "foobar" {
  name                    = "%s"
  background_color        = "#000000"
  text_color              = "#000000"
  label_list_visibility   = "labelHide"
  message_list_visibility = "hide"
}
`

var testAccGmailFilterLabel_update = `
resource gmailfilter_label "foobar" {
  name                    = "%s"
  background_color        = "#ffffff"
  text_color              = "#ffffff"
  label_list_visibility   = "labelShowIfUnread"
  message_list_visibility = "show"
}
`
