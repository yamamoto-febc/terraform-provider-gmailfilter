package gmailfilter

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMain(m *testing.M) {
	acctest.UseBinaryDriver("gmailfilter", Provider)
	resource.TestMain(m)
}
