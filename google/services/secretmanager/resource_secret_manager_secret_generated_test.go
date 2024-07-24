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

package secretmanager_test

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

func TestAccSecretManagerSecret_secretConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecret_secretConfigBasicExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "labels", "secret_id", "terraform_labels", "ttl"},
			},
		},
	})
}

func testAccSecretManagerSecret_secretConfigBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "secret%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}
`, context)
}

func TestAccSecretManagerSecret_secretWithAnnotationsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecret_secretWithAnnotationsExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-with-annotations",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "labels", "secret_id", "terraform_labels", "ttl"},
			},
		},
	})
}

func testAccSecretManagerSecret_secretWithAnnotationsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-with-annotations" {
  secret_id = "secret%{random_suffix}"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "someval"
    key2 = "someval2"
    key3 = "someval3"
    key4 = "someval4"
    key5 = "someval5"
  }

  replication {
    auto {}
  }
}
`, context)
}

func TestAccSecretManagerSecret_secretWithVersionDestroyTtlExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecret_secretWithVersionDestroyTtlExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-with-version-destroy-ttl",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "labels", "secret_id", "terraform_labels", "ttl"},
			},
		},
	})
}

func testAccSecretManagerSecret_secretWithVersionDestroyTtlExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-with-version-destroy-ttl" {
  secret_id = "secret%{random_suffix}"

  version_destroy_ttl = "2592000s"

  replication {
    auto {}
  }
}
`, context)
}

func TestAccSecretManagerSecret_secretWithAutomaticCmekExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"kms_key_name":  acctest.BootstrapKMSKey(t).CryptoKey.Name,
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecret_secretWithAutomaticCmekExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-with-automatic-cmek",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "labels", "secret_id", "terraform_labels", "ttl"},
			},
		},
	})
}

func testAccSecretManagerSecret_secretWithAutomaticCmekExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {}

resource "google_kms_crypto_key_iam_member" "kms-secret-binding" {
  crypto_key_id = "%{kms_key_name}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-secretmanager.iam.gserviceaccount.com"
}

resource "google_secret_manager_secret" "secret-with-automatic-cmek" {
  secret_id = "secret%{random_suffix}"

  replication {
    auto {
      customer_managed_encryption {
        kms_key_name = "%{kms_key_name}"
      }
    }
  }

  depends_on = [ google_kms_crypto_key_iam_member.kms-secret-binding ]
}
`, context)
}

func testAccCheckSecretManagerSecretDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_secret_manager_secret" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SecretManagerBasePath}}projects/{{project}}/secrets/{{secret_id}}")
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
				return fmt.Errorf("SecretManagerSecret still exists at %s", url)
			}
		}

		return nil
	}
}
