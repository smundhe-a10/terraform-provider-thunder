---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ipv6_route_static_bfd_ve Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_ipv6_route_static_bfd_ve (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ipv6_route_static_bfd_ve" "VE_1" {
  ve_num          = 55
  nexthop_ipv6_ll = "fe80:55:3:4::1"
  template        = "TEST_4"
  threshold       = 110
  action          = "down"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `nexthop_ipv6_ll` (String) Nexthop IPv6 address (Link_local)
- `ve_num` (Number) Virtual ethernet interface

### Optional

- `action` (String) 'down': BFD down;  (BFD state)
- `template` (String) Configure tracking template (bind tracking template name)
- `threshold` (Number) action triggering threshold
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


