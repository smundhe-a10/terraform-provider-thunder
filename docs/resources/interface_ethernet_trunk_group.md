---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_interface_ethernet_trunk_group Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_interface_ethernet_trunk_group (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_ethernet_trunk_group" "trunkgroup" {
  ifnum         = 5
  mode          = "active"
  user_tag      = "trunk"
  port_priority = 100
  admin_key     = 20000
  type          = "lacp"
  timeout       = "long"
  trunk_number  = 4
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `admin_key` (Number)
- `ifnum` (Number)
- `mode` (String)
- `port_priority` (Number)
- `timeout` (String)
- `trunk_number` (Number)
- `type` (String)
- `udld_timeout_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--udld_timeout_cfg))
- `user_tag` (String)
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--udld_timeout_cfg"></a>
### Nested Schema for `udld_timeout_cfg`

Optional:

- `fast` (Number)
- `slow` (Number)


