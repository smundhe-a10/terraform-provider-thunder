---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_router_bgp_neighbor_ethernet_neighbor Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_router_bgp_neighbor_ethernet_neighbor (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_router_bgp_neighbor_ethernet_neighbor" "ethernetBGP" {
  as_number       = "300"
  process_id      = 200
  ethernet        = 3
  unnumbered      = 1
  peer_group_name = "A10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (String)
- `as_number` (Number)
- `ethernet` (Number)
- `peer_group_name` (String)
- `process_id` (String)
- `sequence` (String)
- `unnumbered` (Number)
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.


