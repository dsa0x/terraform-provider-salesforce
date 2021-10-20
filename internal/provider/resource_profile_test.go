package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProfile_basic(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProfile_basic(name),
			},
			{
				ResourceName:      "salesforce_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceProfile_update(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProfile_basic(name),
			},
			{
				ResourceName:      "salesforce_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceProfile_with_permissions(name),
			},
			{
				ResourceName:      "salesforce_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceProfile_basic(name string) string {
	return fmt.Sprintf(`
data "salesforce_user_license" "fdc" {
  license_definition_key = "PID_FDC_FREE"
}

resource "salesforce_profile" "test" {
  name            = "%s"
  user_license_id = data.salesforce_user_license.fdc.id
  description     = "test"
}
`, name)
}

func testAccResourceProfile_with_permissions(name string) string {
	return fmt.Sprintf(`
data "salesforce_user_license" "fdc" {
  license_definition_key = "PID_FDC_FREE"
}

resource "salesforce_profile" "test" {
  name            = "%s"
  user_license_id = data.salesforce_user_license.fdc.id
  description     = "test update"
  permissions_email_single = true
  permissions_edit_task = true
}
`, name)
}