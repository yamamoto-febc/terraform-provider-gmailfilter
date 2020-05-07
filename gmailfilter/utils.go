package gmailfilter

import (
	"fmt"
	"log"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/googleapi"
)

func handleNotFoundError(err error, d *schema.ResourceData, resource string) error {
	if isGoogleApiErrorWithCode(err, 404) {
		log.Printf("[WARN] Removing %s because it's gone", resource)
		// The resource doesn't exist anymore
		d.SetId("")

		return nil
	}

	return fmt.Errorf("error reading %s: %s", resource, err)
}

func isGoogleApiErrorWithCode(err error, errCode int) bool {
	gerr, ok := errwrap.GetType(err, &googleapi.Error{}).(*googleapi.Error)
	return ok && gerr != nil && gerr.Code == errCode
}
