---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "deadline_storage_profile Resource - deadline"
subcategory: ""
description: |-
  StorageProfile resource
---

# deadline_storage_profile (Resource)

StorageProfile resource

## Example Usage

```terraform
resource "deadline_farm" "test" {
  display_name = "test"
  description  = "this is a test farm"
}

resource "deadline_queue" "test" {
  farm_id      = deadline_farm.test.id
  display_name = "test queue"
  description  = "This is a test queue"
}

resource "deadline_fleet" "test" {
  farm_id          = deadline_farm.test.id
  display_name     = "test"
  description      = "This is a test fleet"
  role_arn         = "arn:aws:iam::123456789012:role/DeadlineWorkerRole"
  min_worker_count = 0
  max_worker_count = 1
  configuration {
    mode = "aws_managed"
    ec2_instance_capabilities {
      cpu_architecture = "x86_64"
      min_cpu_count    = 1
      max_cpu_count    = 2
      memory_mib_range {
        min = 1024
        max = 1024 * 4
      }
      os_family = "LINUX" // LINUX, WINDOWS
      root_ebs_volume {
        iops = 100
        size = 100
      }
    }
  }
}

resource "deadline_storage_profile" "test" {
  farm_id   = deadline_farm.test.id
  os_family = "windows"
  file_system_location {
    name = "test"
    path = "c:\\path\\to\\file\\system"
    type = "local"
  }
  file_system_location {
    name = "network-share"
    path = "smb://deadline-cloud-files/project/directory/test"
    type = "shared"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) The display name of the storage profile.
- `farm_id` (String) The deadline farm associated with the storage profile.
- `os_family` (String) The OS family of the storage profile. Can be: windows, linux or macos

### Optional

- `file_system_location` (Block List) (see [below for nested schema](#nestedblock--file_system_location))

### Read-Only

- `id` (String) The ID of the storage profile.

<a id="nestedblock--file_system_location"></a>
### Nested Schema for `file_system_location`

Required:

- `name` (String) Name of the file system location
- `path` (String) Path of the file system location
- `type` (String) Type of the file system location. Can be either local, or shared