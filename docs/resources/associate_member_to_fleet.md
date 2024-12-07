---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "deadline_associate_member_to_fleet Resource - deadline"
subcategory: ""
description: |-
  Associate Member to fleet resource
---

# deadline_associate_member_to_fleet (Resource)

Associate Member to fleet resource

## Example Usage

```terraform
resource "deadline_farm" "test" {
  display_name = "test"
  description  = "this is a test farm"
}

resource "deadline_fleet" "test" {
  display_name = "test"
  farm_id      = deadline_farm.test.id
  description  = "this is a test farm"
}

resource "deadline_associate_member_to_fleet" "test" {
  farm_id           = deadline_farm.test.id
  fleet_id          = deadline_fleet.test.id
  member_id         = "test"
  identity_store_id = "example_identity_store"
  membership_level  = "VIEWER"
  principal_type    = "USER"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `farm_id` (String) The ID of the farm to associate the member to
- `fleet_id` (String) The ID of the fleet to associate the member to
- `identity_store_id` (String) The ID of the identity store that the member belongs to
- `membership_level` (String) The membership level of the principal to associate to the farm. Valid values are `VIEWER`, `CONTRIBUTOR`, `OWNER` and `MANAGER`
- `principal_id` (String) The ID of the principal to associate to the fleet
- `principal_type` (String) The type of principal to associate to the fleet. Valid values are `USER` and `GROUP`

### Read-Only

- `id` (String) The ID of the associate_member_to_fleet.
