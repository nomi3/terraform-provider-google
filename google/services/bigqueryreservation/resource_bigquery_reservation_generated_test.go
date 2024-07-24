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

package bigqueryreservation_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccBigqueryReservationReservation_bigqueryReservationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigqueryReservationReservationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryReservationReservation_bigqueryReservationBasicExample(context),
			},
			{
				ResourceName:            "google_bigquery_reservation.reservation",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
		},
	})
}

func testAccBigqueryReservationReservation_bigqueryReservationBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_reservation" "reservation" {
	name           = "tf-test-my-reservation%{random_suffix}"
	location       = "us-west2"
	// Set to 0 for testing purposes
	// In reality this would be larger than zero
	slot_capacity     = 0
	edition = "STANDARD"
	ignore_idle_slots = true
	concurrency       = 0
	autoscale {
   	  max_slots = 100
    }
}
`, context)
}

func testAccCheckBigqueryReservationReservationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_reservation" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{name}}")
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
				return fmt.Errorf("BigqueryReservationReservation still exists at %s", url)
			}
		}

		return nil
	}
}
