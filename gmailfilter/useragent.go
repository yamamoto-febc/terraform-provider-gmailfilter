package gmailfilter

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
)

const tfUAEnvVar = "TF_APPEND_USER_AGENT"

func terraformUserAgent(version string) string {
	ua := fmt.Sprintf("HashiCorp Terraform/%s (+https://www.terraform.io) Terraform Plugin SDK/%s", version, meta.SDKVersionString())

	if add := os.Getenv(tfUAEnvVar); add != "" {
		add = strings.TrimSpace(add)
		if len(add) > 0 {
			ua += " " + add
			log.Printf("[DEBUG] Using modified User-Agent: %s", ua)
		}
	}

	return ua
}
