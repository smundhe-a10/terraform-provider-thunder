---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ipv6_route_static_bfd_trunk Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_ipv6_route_static_bfd_trunk (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ipv6_route_static_bfd_trunk" "bfd_trunk" {
  trunk_num       = 7
  nexthop_ipv6_ll = "fe80:2:2:8::4"
  template        = "TEST_3"
  threshold       = 240
  action          = "down"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `nexthop_ipv6_ll` (String) Nexthop IPv6 address (Link_local)
- `trunk_num` (Number) Trunk interface

### Optional

- `action` (String) 'down': BFD down;  (BFD state)
- `template` (String) Configure tracking template (bind tracking template name)
- `threshold` (Number) action triggering threshold
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


