// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"deadline": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}
func TestAccAssociateMemberToFarmResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccAssociateMemberToFarmResourceConfig("test", "this is a test farm", os.Getenv("TEST_PRINCIPAL_ID"), os.Getenv("TEST_IDENTITY_STORE_ID")),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("deadline_associate_member_to_farm.test", "principal_id", os.Getenv("TEST_PRINCIPAL_ID")),
					resource.TestCheckResourceAttr("deadline_associate_member_to_farm.test", "identity_store_id", os.Getenv("TEST_IDENTITY_STORE_ID")),
				),
			},
		},
	})
}

func testAccAssociateMemberToFarmResourceConfig(displayName string, description string, principalID string, identityStoreId string) string {
	return fmt.Sprintf(`
resource "deadline_farm" "test" {
  display_name = %[1]q
  description  = %[2]q
}

resource "deadline_associate_member_to_farm" "test" {
  farm_id = "${deadline_farm.test.id}"
  principal_id = %[3]q
  identity_store_id = %[4]q
}
`, displayName, description, principalID, identityStoreId)
}
func TestAccAssociateMemberToFleetResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccAssociateMemberToFleetResourceConfig("test", "this is a test fleet", os.Getenv("TEST_PRINCIPAL_ID"), os.Getenv("TEST_IDENTITY_STORE_ID")),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("deadline_associate_member_to_fleet.test", "principal_id", os.Getenv("TEST_PRINCIPAL_ID")),
					resource.TestCheckResourceAttr("deadline_associate_member_to_fleet.test", "identity_store_id", os.Getenv("TEST_IDENTITY_STORE_ID")),
				),
			},
		},
	})
}

func testAccAssociateMemberToFleetResourceConfig(displayName string, description string, principalID string, identityStoreId string) string {
	return fmt.Sprintf(`

resource "deadline_farm" "test" {
  display_name = %[1]q
  description  = %[2]q
}
resource "deadline_fleet" "test" {
  farm_id = "${deadline_farm.test.id}"
  display_name = %[1]q
  description  = %[2]q
}

resource "deadline_associate_member_to_fleet" "test" {
  farm_id = "${deadline_farm.test.id}"
  fleet_id = "${deadline_fleet.test.id}"
  principal_id = %[3]q
  identity_store_id = %[4]q
}
`, displayName, description, principalID, identityStoreId)
}

func TestAccFarmResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccFarmResourceConfig("test", "this is a test farm"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("deadline_farm.test", "display_name", "test"),
					resource.TestCheckResourceAttr("deadline_farm.test", "description", "this is a test farm"),
				),
			},
		},
	})
}

func testAccFarmResourceConfig(displayName string, description string) string {
	return fmt.Sprintf(`
resource "deadline_farm" "test" {
  display_name = %[1]q
  description = %[2]q
}
`, displayName, description)
}
func TestAccFleetResource(t *testing.T) {
	testRoleARN := os.Getenv("TEST_ROLE_ARN")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccFleetResourceConfig("test", "this is a test farm", testRoleARN),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("deadline_fleet.test", "display_name", "test"),
					resource.TestCheckResourceAttr("deadline_fleet.test", "description", "this is a test farm"),
				),
			},
		},
	})
}

func testAccFleetResourceConfig(displayName string, description string, roleARN string) string {
	return fmt.Sprintf(`
resource "deadline_farm" "test" {
	display_name = %[1]q
    description  = "this is a farm"
}
resource "deadline_fleet" "test" {
  farm_id = "${deadline_farm.test.id}"
  display_name = %[1]q
  description = %[2]q
  role_arn = %[3]q
  configuration {
    mode = "aws_managed"
	ec2_instance_capabilities {
      os_family = "windows"
      cpu_architecture = "x86_64"
	  memory_mib = 4096
	  allowed_instance_types = ["t2.micro"]
	  min_cpu_count = 1
	  max_cpu_count = 2
    }
  }
  min_worker_count = "0"
  max_worker_count = "1"
}
`, displayName, description, roleARN)
}

func TestAccQueueResource(t *testing.T) {
	testRoleARN := os.Getenv("TEST_ROLE_ARN")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccQueueResourceConfig("test", "this is a test", testRoleARN),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("deadline_queue.test", "display_name", "test"),
					resource.TestCheckResourceAttr("deadline_queue.test", "description", "this is a test"),
					resource.TestCheckResourceAttr("deadline_queue.test", "role_arn", testRoleARN),
				),
			},
		},
	})
}

func testAccQueueResourceConfig(displayName string, description string, roleARN string) string {
	return fmt.Sprintf(`
resource "deadline_farm" "test" {
	display_name = %[1]q
    description  = "this is a farm"
}
resource "deadline_queue" "test" {
  farm_id = "${deadline_farm.test.id}"
  display_name = %[1]q
  description = %[2]q
  role_arn = %[3]q
  allowed_storage_profile_ids = ["storage_profile_id"]
}
`, displayName, description, roleARN)
}
