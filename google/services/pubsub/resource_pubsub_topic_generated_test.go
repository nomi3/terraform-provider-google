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

package pubsub_test

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

func TestAccPubsubTopic_pubsubTopicBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubTopicDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopic_pubsubTopicBasicExample(context),
			},
			{
				ResourceName:            "google_pubsub_topic.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccPubsubTopic_pubsubTopicBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }

  message_retention_duration = "86600s"
}
`, context)
}

func TestAccPubsubTopic_pubsubTopicGeoRestrictedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubTopicDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopic_pubsubTopicGeoRestrictedExample(context),
			},
			{
				ResourceName:            "google_pubsub_topic.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccPubsubTopic_pubsubTopicGeoRestrictedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  message_storage_policy {
    allowed_persistence_regions = [
      "europe-west3",
    ]
  }
}
`, context)
}

func TestAccPubsubTopic_pubsubTopicSchemaSettingsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubTopicDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopic_pubsubTopicSchemaSettingsExample(context),
			},
			{
				ResourceName:            "google_pubsub_topic.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccPubsubTopic_pubsubTopicSchemaSettingsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_schema" "example" {
  name = "example%{random_suffix}"
  type = "AVRO"
  definition = "{\n  \"type\" : \"record\",\n  \"name\" : \"Avro\",\n  \"fields\" : [\n    {\n      \"name\" : \"StringField\",\n      \"type\" : \"string\"\n    },\n    {\n      \"name\" : \"IntField\",\n      \"type\" : \"int\"\n    }\n  ]\n}\n"
}

resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  depends_on = [google_pubsub_schema.example]
  schema_settings {
    schema = "projects/%{project_name}/schemas/example%{random_suffix}"
    encoding = "JSON"
  }
}
`, context)
}

func TestAccPubsubTopic_pubsubTopicIngestionKinesisExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubTopicDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopic_pubsubTopicIngestionKinesisExample(context),
			},
			{
				ResourceName:            "google_pubsub_topic.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccPubsubTopic_pubsubTopicIngestionKinesisExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  # Outside of automated terraform-provider-google CI tests, these values must be of actual AWS resources for the test to pass.
  ingestion_data_source_settings {
    aws_kinesis {
        stream_arn = "arn:aws:kinesis:us-west-2:111111111111:stream/fake-stream-name"
        consumer_arn = "arn:aws:kinesis:us-west-2:111111111111:stream/fake-stream-name/consumer/consumer-1:1111111111"
        aws_role_arn = "arn:aws:iam::111111111111:role/fake-role-name"
        gcp_service_account = "fake-service-account@fake-gcp-project.iam.gserviceaccount.com"
    }
  }
}
`, context)
}

func testAccCheckPubsubTopicDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_pubsub_topic" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{PubsubBasePath}}projects/{{project}}/topics/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:               config,
				Method:               "GET",
				Project:              billingProject,
				RawURL:               url,
				UserAgent:            config.UserAgent,
				ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.PubsubTopicProjectNotReady},
			})
			if err == nil {
				return fmt.Errorf("PubsubTopic still exists at %s", url)
			}
		}

		return nil
	}
}
