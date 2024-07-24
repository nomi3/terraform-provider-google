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

package securesourcemanager_test

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

func TestAccSecureSourceManagerInstance_secureSourceManagerInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecureSourceManagerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerInstance_secureSourceManagerInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_secure_source_manager_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccSecureSourceManagerInstance_secureSourceManagerInstanceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "default" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
    labels = {
      "foo" = "bar"
    }
}
`, context)
}

func TestAccSecureSourceManagerInstance_secureSourceManagerInstanceCmekExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecureSourceManagerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerInstance_secureSourceManagerInstanceCmekExample(context),
			},
			{
				ResourceName:            "google_secure_source_manager_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccSecureSourceManagerInstance_secureSourceManagerInstanceCmekExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_kms_key_ring" "key_ring" {
  name     = "tf-test-my-keyring%{random_suffix}"
  location = "us-central1"
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "tf-test-my-key%{random_suffix}"
  key_ring = google_kms_key_ring.key_ring.id
}

resource "google_kms_crypto_key_iam_member" "crypto_key_binding" {
  crypto_key_id = google_kms_crypto_key.crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-sourcemanager.iam.gserviceaccount.com"
}

resource "google_secure_source_manager_instance" "default" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
    kms_key = google_kms_crypto_key.crypto_key.id

    depends_on = [
      google_kms_crypto_key_iam_member.crypto_key_binding
    ]
}

data "google_project" "project" {}
`, context)
}

func TestAccSecureSourceManagerInstance_secureSourceManagerInstancePrivateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccCheckSecureSourceManagerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerInstance_secureSourceManagerInstancePrivateExample(context),
			},
			{
				ResourceName:            "google_secure_source_manager_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccSecureSourceManagerInstance_secureSourceManagerInstancePrivateExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "ca_pool" {
  name     = "tf-test-ca-pool%{random_suffix}"
  location = "us-central1"
  tier     = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl     = true
  }
}

resource "google_privateca_certificate_authority" "root_ca" {
  pool                     = google_privateca_ca_pool.ca_pool.name
  certificate_authority_id = "tf-test-root-ca%{random_suffix}"
  location                 = "us-central1"
  config {
    subject_config {
      subject {
        organization = "google"
        common_name = "my-certificate-authority"
      }
    }
    x509_config {
      ca_options {
        is_ca = true
      }
      key_usage {
        base_key_usage {
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = true
        }
      }
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }

  // Disable deletion protections for easier test cleanup purposes
  deletion_protection = false
  ignore_active_certificates_on_deletion = true
  skip_grace_period = true
}

resource "google_privateca_ca_pool_iam_binding" "ca_pool_binding" {
  ca_pool = google_privateca_ca_pool.ca_pool.id
  role = "roles/privateca.certificateRequester"

  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-sourcemanager.iam.gserviceaccount.com"
  ]
}

resource "google_secure_source_manager_instance" "default" {
  instance_id = "tf-test-my-instance%{random_suffix}"
  location = "us-central1"
  private_config {
    is_private = true
    ca_pool = google_privateca_ca_pool.ca_pool.id
  }
  depends_on = [
    google_privateca_certificate_authority.root_ca,
    time_sleep.wait_120_seconds
  ]
}

# ca pool IAM permissions can take time to propagate
resource "time_sleep" "wait_120_seconds" {
  depends_on = [google_privateca_ca_pool_iam_binding.ca_pool_binding]

  create_duration = "120s"
}

data "google_project" "project" {}
`, context)
}

func TestAccSecureSourceManagerInstance_secureSourceManagerInstancePrivatePscBackendExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccCheckSecureSourceManagerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerInstance_secureSourceManagerInstancePrivatePscBackendExample(context),
			},
			{
				ResourceName:            "google_secure_source_manager_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccSecureSourceManagerInstance_secureSourceManagerInstancePrivatePscBackendExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {}

resource "google_privateca_ca_pool" "ca_pool" {
  name     = "tf-test-ca-pool%{random_suffix}"
  location = "us-central1"
  tier     = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl     = true
  }
}

resource "google_privateca_certificate_authority" "root_ca" {
  pool                     = google_privateca_ca_pool.ca_pool.name
  certificate_authority_id = "tf-test-root-ca%{random_suffix}"
  location                 = "us-central1"
  config {
    subject_config {
      subject {
        organization = "google"
        common_name = "my-certificate-authority"
      }
    }
    x509_config {
      ca_options {
        is_ca = true
      }
      key_usage {
        base_key_usage {
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = true
        }
      }
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }

  // Disable deletion protections for easier test cleanup purposes
  deletion_protection = false
  ignore_active_certificates_on_deletion = true
  skip_grace_period = true
}

resource "google_privateca_ca_pool_iam_binding" "ca_pool_binding" {
  ca_pool = google_privateca_ca_pool.ca_pool.id
  role = "roles/privateca.certificateRequester"

  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-sourcemanager.iam.gserviceaccount.com"
  ]
}

// See https://cloud.google.com/secure-source-manager/docs/create-private-service-connect-instance#root-ca-api
resource "google_secure_source_manager_instance" "default" {
  instance_id = "tf-test-my-instance%{random_suffix}"
  location = "us-central1"
  private_config {
    is_private = true
    ca_pool = google_privateca_ca_pool.ca_pool.id
  }
  depends_on = [
    google_privateca_certificate_authority.root_ca,
    time_sleep.wait_120_seconds
  ]
}

# ca pool IAM permissions can take time to propagate
resource "time_sleep" "wait_120_seconds" {
  depends_on = [google_privateca_ca_pool_iam_binding.ca_pool_binding]

  create_duration = "120s"
}

// Connect SSM private instance with L4 proxy ILB.
resource "google_compute_network" "network" {
  name = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet" {
  name = "tf-test-my-subnet%{random_suffix}"
  region = "us-central1"
  network = google_compute_network.network.id
  ip_cidr_range = "10.0.1.0/24"
  private_ip_google_access = true
}

resource "google_compute_region_network_endpoint_group" "psc_neg" {
  name = "tf-test-my-neg%{random_suffix}"
  region = "us-central1"

  network_endpoint_type = "PRIVATE_SERVICE_CONNECT"
  psc_target_service = google_secure_source_manager_instance.default.private_config.0.http_service_attachment

  network = google_compute_network.network.id
  subnetwork = google_compute_subnetwork.subnet.id
}

resource "google_compute_region_backend_service" "backend_service" {
  name = "tf-test-my-backend-service%{random_suffix}"
  region = "us-central1"
  protocol = "TCP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  backend {
    group = google_compute_region_network_endpoint_group.psc_neg.id
    balancing_mode = "UTILIZATION"
    capacity_scaler = 1.0
  }
}

resource "google_compute_subnetwork" "proxy_subnet" {
  name = "tf-test-my-proxy-subnet%{random_suffix}"
  region = "us-central1"
  network = google_compute_network.network.id
  ip_cidr_range = "10.0.2.0/24"
  purpose = "REGIONAL_MANAGED_PROXY"
  role = "ACTIVE"
}

resource "google_compute_region_target_tcp_proxy" "target_proxy" {
  name = "tf-test-my-target-proxy%{random_suffix}"
  region = "us-central1"
  backend_service = google_compute_region_backend_service.backend_service.id
}

resource "google_compute_forwarding_rule" "fw_rule_target_proxy" {
  name = "tf-test-fw-rule-target-proxy%{random_suffix}"
  region = "us-central1"

  load_balancing_scheme = "INTERNAL_MANAGED"
  ip_protocol = "TCP"
  port_range = "443"
  target = google_compute_region_target_tcp_proxy.target_proxy.id
  network = google_compute_network.network.id
  subnetwork = google_compute_subnetwork.subnet.id
  network_tier = "PREMIUM"
  depends_on = [google_compute_subnetwork.proxy_subnet]
}

resource "google_dns_managed_zone" "private_zone" {
  name = "tf-test-my-dns-zone%{random_suffix}"
  dns_name = "p.sourcemanager.dev."
  visibility = "private"
  private_visibility_config {
    networks {
      network_url = google_compute_network.network.id
    }
  }
}

resource "google_dns_record_set" "ssm_instance_html_record" {
  name = "${google_secure_source_manager_instance.default.host_config.0.html}."
  type = "A"
  ttl = 300
  managed_zone = google_dns_managed_zone.private_zone.name
  rrdatas = [google_compute_forwarding_rule.fw_rule_target_proxy.ip_address]
}

resource "google_dns_record_set" "ssm_instance_api_record" {
  name = "${google_secure_source_manager_instance.default.host_config.0.api}."
  type = "A"
  ttl = 300
  managed_zone = google_dns_managed_zone.private_zone.name
  rrdatas = [google_compute_forwarding_rule.fw_rule_target_proxy.ip_address]
}

resource "google_dns_record_set" "ssm_instance_git_record" {
  name = "${google_secure_source_manager_instance.default.host_config.0.git_http}."
  type = "A"
  ttl = 300
  managed_zone = google_dns_managed_zone.private_zone.name
  rrdatas = [google_compute_forwarding_rule.fw_rule_target_proxy.ip_address]
}
`, context)
}

func TestAccSecureSourceManagerInstance_secureSourceManagerInstancePrivatePscEndpointExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccCheckSecureSourceManagerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerInstance_secureSourceManagerInstancePrivatePscEndpointExample(context),
			},
			{
				ResourceName:            "google_secure_source_manager_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccSecureSourceManagerInstance_secureSourceManagerInstancePrivatePscEndpointExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {}

resource "google_privateca_ca_pool" "ca_pool" {
  name     = "tf-test-ca-pool%{random_suffix}"
  location = "us-central1"
  tier     = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl     = true
  }
}

resource "google_privateca_certificate_authority" "root_ca" {
  pool                     = google_privateca_ca_pool.ca_pool.name
  certificate_authority_id = "tf-test-root-ca%{random_suffix}"
  location                 = "us-central1"
  config {
    subject_config {
      subject {
        organization = "google"
        common_name = "my-certificate-authority"
      }
    }
    x509_config {
      ca_options {
        is_ca = true
      }
      key_usage {
        base_key_usage {
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = true
        }
      }
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }

  // Disable deletion protections for easier test cleanup purposes
  deletion_protection = false
  ignore_active_certificates_on_deletion = true
  skip_grace_period = true
}

resource "google_privateca_ca_pool_iam_binding" "ca_pool_binding" {
  ca_pool = google_privateca_ca_pool.ca_pool.id
  role = "roles/privateca.certificateRequester"

  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-sourcemanager.iam.gserviceaccount.com"
  ]
}

// See https://cloud.google.com/secure-source-manager/docs/create-private-service-connect-instance#root-ca-api
resource "google_secure_source_manager_instance" "default" {
  instance_id = "tf-test-my-instance%{random_suffix}"
  location = "us-central1"
  private_config {
    is_private = true
    ca_pool = google_privateca_ca_pool.ca_pool.id
  }
  depends_on = [
    google_privateca_certificate_authority.root_ca,
    time_sleep.wait_120_seconds
  ]
}

# ca pool IAM permissions can take time to propagate
resource "time_sleep" "wait_120_seconds" {
  depends_on = [google_privateca_ca_pool_iam_binding.ca_pool_binding]

  create_duration = "120s"
}

// Connect SSM private instance with endpoint.
resource "google_compute_network" "network" {
  name = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet" {
  name = "tf-test-my-subnet%{random_suffix}"
  region = "us-central1"
  network = google_compute_network.network.id
  ip_cidr_range = "10.0.60.0/24"
  private_ip_google_access = true
}

resource "google_compute_address" "address" {
  name = "tf-test-my-address%{random_suffix}"
  region = "us-central1"
  address = "10.0.60.100"
  address_type = "INTERNAL"
  subnetwork = google_compute_subnetwork.subnet.id
}

resource "google_compute_forwarding_rule" "fw_rule_service_attachment" {
  name = "tf-test-fw-rule-service-attachment%{random_suffix}"
  region = "us-central1"

  load_balancing_scheme = ""
  ip_address = google_compute_address.address.id
  network = google_compute_network.network.id

  target = google_secure_source_manager_instance.default.private_config.0.http_service_attachment
}

resource "google_dns_managed_zone" "private_zone" {
  name = "tf-test-my-dns-zone%{random_suffix}"
  dns_name = "p.sourcemanager.dev."
  visibility = "private"
  private_visibility_config {
    networks {
      network_url = google_compute_network.network.id
    }
  }
}

resource "google_dns_record_set" "ssm_instance_html_record" {
  name = "${google_secure_source_manager_instance.default.host_config.0.html}."
  type = "A"
  ttl = 300
  managed_zone = google_dns_managed_zone.private_zone.name
  rrdatas = [google_compute_forwarding_rule.fw_rule_service_attachment.ip_address]
}

resource "google_dns_record_set" "ssm_instance_api_record" {
  name = "${google_secure_source_manager_instance.default.host_config.0.api}."
  type = "A"
  ttl = 300
  managed_zone = google_dns_managed_zone.private_zone.name
  rrdatas = [google_compute_forwarding_rule.fw_rule_service_attachment.ip_address]
}

resource "google_dns_record_set" "ssm_instance_git_record" {
  name = "${google_secure_source_manager_instance.default.host_config.0.git_http}."
  type = "A"
  ttl = 300
  managed_zone = google_dns_managed_zone.private_zone.name
  rrdatas = [google_compute_forwarding_rule.fw_rule_service_attachment.ip_address]
}
`, context)
}

func testAccCheckSecureSourceManagerInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_secure_source_manager_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance_id}}")
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
				return fmt.Errorf("SecureSourceManagerInstance still exists at %s", url)
			}
		}

		return nil
	}
}
