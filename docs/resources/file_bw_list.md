---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_file_bw_list Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_file_class_list: Black-White List file
---

# thunder_file_bw_list (Resource)

`thunder_file_class_list`: Black-White List file

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

provider "thunder" {
  alias     = "L3V_A"
  address   = var.dut9049
  username  = var.username
  password  = var.password
  partition = var.l3v_1
}

resource "thunder_file_bw_list" "BWLIST_1" {
  name          = "BWLIST_1"
  protocol      = "http"
  host          = "192.168.92.200"
  path          = "/class-list-files/BWLIST_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_bw_list" "L3V_BWLIST_1" {
  provider      = thunder.L3V_A
  name          = "BWLIST_1"
  protocol      = "https"
  host          = "192.168.92.200"
  path          = "/class-list-files/BWLIST_1.txt"
  use_mgmt_port = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `host` (String) Remote site (IP or domain name)
- `name` (String) Local file name
- `path` (String) Remote path
- `protocol` (String) Transfer protocol

### Optional

- `overwrite` (Number) Overwrite existing file
- `password` (String) Password for the remote site
- `use_mgmt_port` (Number) Use management port as source port
- `username` (String) Username for the remote site

### Read-Only

- `id` (String) The ID of this resource.


