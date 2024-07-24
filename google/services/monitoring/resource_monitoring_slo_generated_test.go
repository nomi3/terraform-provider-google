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

package monitoring_test

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

func TestAccMonitoringSlo_monitoringSloAppengineExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringSloDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringSlo_monitoringSloAppengineExample(context),
			},
			{
				ResourceName:            "google_monitoring_slo.appeng_slo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service"},
			},
		},
	})
}

func testAccMonitoringSlo_monitoringSloAppengineExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_monitoring_app_engine_service" "default" {
  module_id = "default"
}

resource "google_monitoring_slo" "appeng_slo" {
  service = data.google_monitoring_app_engine_service.default.service_id

  slo_id = "tf-test-ae-slo%{random_suffix}"
  display_name = "Terraform Test SLO for App Engine"

  goal = 0.9
  calendar_period = "DAY"

  basic_sli {
    latency {
      threshold = "1s"
    }
  }

  user_labels = {
    my_key       = "my_value"
    my_other_key = "my_other_value"
  }
}
`, context)
}

func TestAccMonitoringSlo_monitoringSloRequestBasedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringSloDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringSlo_monitoringSloRequestBasedExample(context),
			},
			{
				ResourceName:            "google_monitoring_slo.request_based_slo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service"},
			},
		},
	})
}

func testAccMonitoringSlo_monitoringSloRequestBasedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_custom_service" "customsrv" {
  service_id = "tf-test-custom-srv-request-slos%{random_suffix}"
  display_name = "My Custom Service"
}

resource "google_monitoring_slo" "request_based_slo" {
  service = google_monitoring_custom_service.customsrv.service_id
  slo_id = "tf-test-consumed-api-slo%{random_suffix}"
  display_name = "Terraform Test SLO with request based SLI (good total ratio)"

  goal = 0.9
  rolling_period_days = 30

  request_based_sli {
    distribution_cut {
          distribution_filter = "metric.type=\"serviceruntime.googleapis.com/api/request_latencies\" resource.type=\"api\"  "
          range {
            max = 0.5
          }
        }
  }
}
`, context)
}

func TestAccMonitoringSlo_monitoringSloWindowsBasedGoodBadMetricFilterExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringSloDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringSlo_monitoringSloWindowsBasedGoodBadMetricFilterExample(context),
			},
			{
				ResourceName:            "google_monitoring_slo.windows_based",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service"},
			},
		},
	})
}

func testAccMonitoringSlo_monitoringSloWindowsBasedGoodBadMetricFilterExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_custom_service" "customsrv" {
  service_id = "tf-test-custom-srv-windows-slos%{random_suffix}"
  display_name = "My Custom Service"
}

resource "google_monitoring_slo" "windows_based" {
  service = google_monitoring_custom_service.customsrv.service_id
  display_name = "Terraform Test SLO with window based SLI"

  goal = 0.95
  calendar_period = "FORTNIGHT"

  windows_based_sli {
    window_period = "400s"
    good_bad_metric_filter =  join(" AND ", [
      "metric.type=\"monitoring.googleapis.com/uptime_check/check_passed\"",
      "resource.type=\"uptime_url\"",
    ])
  }
}
`, context)
}

func TestAccMonitoringSlo_monitoringSloWindowsBasedMetricMeanExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringSloDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringSlo_monitoringSloWindowsBasedMetricMeanExample(context),
			},
			{
				ResourceName:            "google_monitoring_slo.windows_based",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service"},
			},
		},
	})
}

func testAccMonitoringSlo_monitoringSloWindowsBasedMetricMeanExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_custom_service" "customsrv" {
  service_id = "tf-test-custom-srv-windows-slos%{random_suffix}"
  display_name = "My Custom Service"
}

resource "google_monitoring_slo" "windows_based" {
  service = google_monitoring_custom_service.customsrv.service_id
  display_name = "Terraform Test SLO with window based SLI"

  goal = 0.9
  rolling_period_days = 20

  windows_based_sli {
    window_period = "600s"
    metric_mean_in_range {
      time_series = join(" AND ", [
        "metric.type=\"agent.googleapis.com/cassandra/client_request/latency/95p\"",
        "resource.type=\"gce_instance\"",
      ])

      range {
        max = 5
      }
    }
  }
}
`, context)
}

func TestAccMonitoringSlo_monitoringSloWindowsBasedMetricSumExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringSloDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringSlo_monitoringSloWindowsBasedMetricSumExample(context),
			},
			{
				ResourceName:            "google_monitoring_slo.windows_based",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service"},
			},
		},
	})
}

func testAccMonitoringSlo_monitoringSloWindowsBasedMetricSumExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_custom_service" "customsrv" {
  service_id = "tf-test-custom-srv-windows-slos%{random_suffix}"
  display_name = "My Custom Service"
}

resource "google_monitoring_slo" "windows_based" {
  service = google_monitoring_custom_service.customsrv.service_id
  display_name = "Terraform Test SLO with window based SLI"

  goal = 0.9
  rolling_period_days = 20

  windows_based_sli {
    window_period = "400s"
    metric_sum_in_range {
      time_series = join(" AND ", [
        "metric.type=\"monitoring.googleapis.com/uptime_check/request_latency\"",
        "resource.type=\"uptime_url\"",
      ])

      range {
        max = 5000
      }
    }
  }
}
`, context)
}

func TestAccMonitoringSlo_monitoringSloWindowsBasedRatioThresholdExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringSloDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringSlo_monitoringSloWindowsBasedRatioThresholdExample(context),
			},
			{
				ResourceName:            "google_monitoring_slo.windows_based",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service"},
			},
		},
	})
}

func testAccMonitoringSlo_monitoringSloWindowsBasedRatioThresholdExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_custom_service" "customsrv" {
  service_id = "tf-test-custom-srv-windows-slos%{random_suffix}"
  display_name = "My Custom Service"
}

resource "google_monitoring_slo" "windows_based" {
  service = google_monitoring_custom_service.customsrv.service_id
  display_name = "Terraform Test SLO with window based SLI"

  goal = 0.9
  rolling_period_days = 20

  windows_based_sli {
    window_period = "100s"

    good_total_ratio_threshold {
      threshold = 0.1
      performance {
        distribution_cut {
          distribution_filter = join(" AND ", [
            "metric.type=\"serviceruntime.googleapis.com/api/request_latencies\"",
            "resource.type=\"consumed_api\"",
          ])

          range {
            min = 1
            max = 9
          }
        }
      }
    }
  }
}
`, context)
}

func testAccCheckMonitoringSloDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_monitoring_slo" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{MonitoringBasePath}}v3/{{name}}")
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
				return fmt.Errorf("MonitoringSlo still exists at %s", url)
			}
		}

		return nil
	}
}
