// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package securitycenterv2_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccSecurityCenterV2OrganizationMuteConfig_sccV2OrganizationMuteConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecurityCenterV2OrganizationMuteConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterV2OrganizationMuteConfig_sccV2OrganizationMuteConfigBasicExample(context),
			},
			{
				ResourceName:            "google_scc_v2_organization_mute_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "mute_config_id", "organization"},
			},
		},
	})
}

func testAccSecurityCenterV2OrganizationMuteConfig_sccV2OrganizationMuteConfigBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_scc_v2_organization_mute_config" "default" {
  mute_config_id    = "tf-test-my-config%{random_suffix}"
  organization = "%{org_id}"
  location     = "global"
  description  = "My custom Cloud Security Command Center Finding Organization mute Configuration"
  filter = "severity = \"HIGH\""
  type = "STATIC"
}
`, context)
}

func testAccCheckSecurityCenterV2OrganizationMuteConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_scc_v2_organization_mute_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SecurityCenterV2BasePath}}organizations/{{organization}}/locations/{{location}}/muteConfigs/{{mute_config_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("SecurityCenterV2OrganizationMuteConfig still exists at %s", url)
			}
		}

		return nil
	}
}