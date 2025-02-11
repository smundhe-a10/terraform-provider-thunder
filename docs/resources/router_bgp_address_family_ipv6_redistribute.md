---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_router_bgp_address_family_ipv6_redistribute Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_router_bgp_address_family_ipv6_redistribute (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}


resource "thunder_router_bgp_address_family_ipv6_redistribute" "AFIpv6RedistributeTest" {

  as_number = "300"
  connected_cfg {
    connected = 1
    route_map = "ipv6"
  }
  floating_ip_cfg {
    floating_ip = 1
    route_map   = "ipv6"
  }
  nat_map_cfg {
    nat_map   = 1
    route_map = "ipv6"
  }
  nat64_cfg {
    nat64     = 1
    route_map = "ipv6"
  }
  lw4o6_cfg {
    lw4o6     = 1
    route_map = "ipv6"
  }
  static_nat_cfg {
    static_nat = 1
    route_map  = "ipv6"
  }
  ip_nat_cfg {
    ip_nat    = 1
    route_map = "ipv6"
  }
  ip_nat_list_cfg {
    ip_nat_list = 1
    route_map   = "ipv6"
  }
  isis_cfg {
    isis      = 1
    route_map = "ipv6"
  }
  ospf_cfg {
    ospf      = 1
    route_map = "ipv6"
  }
  rip_cfg {
    rip       = 1
    route_map = "ipv6"
  }
  static_cfg {
    static    = 1
    route_map = "ipv6"
  }
  vip {
    only_flagged_cfg {
      only_flagged = 1
      route_map    = "ipv6"
    }
    only_not_flagged_cfg {
      only_not_flagged = 1
      route_map        = "ipv6"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `as_number` (Number)
- `connected_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--connected_cfg))
- `floating_ip_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--floating_ip_cfg))
- `ip_nat_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip_nat_cfg))
- `ip_nat_list_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip_nat_list_cfg))
- `isis_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--isis_cfg))
- `lw4o6_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--lw4o6_cfg))
- `nat64_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--nat64_cfg))
- `nat_map_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--nat_map_cfg))
- `ospf_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ospf_cfg))
- `rip_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--rip_cfg))
- `static_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--static_cfg))
- `static_nat_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--static_nat_cfg))
- `uuid` (String)
- `vip` (Block List, Max: 1) (see [below for nested schema](#nestedblock--vip))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--connected_cfg"></a>
### Nested Schema for `connected_cfg`

Optional:

- `connected` (Number)
- `route_map` (String)


<a id="nestedblock--floating_ip_cfg"></a>
### Nested Schema for `floating_ip_cfg`

Optional:

- `floating_ip` (Number)
- `route_map` (String)


<a id="nestedblock--ip_nat_cfg"></a>
### Nested Schema for `ip_nat_cfg`

Optional:

- `ip_nat` (Number)
- `route_map` (String)


<a id="nestedblock--ip_nat_list_cfg"></a>
### Nested Schema for `ip_nat_list_cfg`

Optional:

- `ip_nat_list` (Number)
- `route_map` (String)


<a id="nestedblock--isis_cfg"></a>
### Nested Schema for `isis_cfg`

Optional:

- `isis` (Number)
- `route_map` (String)


<a id="nestedblock--lw4o6_cfg"></a>
### Nested Schema for `lw4o6_cfg`

Optional:

- `lw4o6` (Number)
- `route_map` (String)


<a id="nestedblock--nat64_cfg"></a>
### Nested Schema for `nat64_cfg`

Optional:

- `nat64` (Number)
- `route_map` (String)


<a id="nestedblock--nat_map_cfg"></a>
### Nested Schema for `nat_map_cfg`

Optional:

- `nat_map` (Number)
- `route_map` (String)


<a id="nestedblock--ospf_cfg"></a>
### Nested Schema for `ospf_cfg`

Optional:

- `ospf` (Number)
- `route_map` (String)


<a id="nestedblock--rip_cfg"></a>
### Nested Schema for `rip_cfg`

Optional:

- `rip` (Number)
- `route_map` (String)


<a id="nestedblock--static_cfg"></a>
### Nested Schema for `static_cfg`

Optional:

- `route_map` (String)
- `static` (Number)


<a id="nestedblock--static_nat_cfg"></a>
### Nested Schema for `static_nat_cfg`

Optional:

- `route_map` (String)
- `static_nat` (Number)


<a id="nestedblock--vip"></a>
### Nested Schema for `vip`

Optional:

- `only_flagged_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--vip--only_flagged_cfg))
- `only_not_flagged_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--vip--only_not_flagged_cfg))

<a id="nestedblock--vip--only_flagged_cfg"></a>
### Nested Schema for `vip.only_flagged_cfg`

Optional:

- `only_flagged` (Number)
- `route_map` (String)


<a id="nestedblock--vip--only_not_flagged_cfg"></a>
### Nested Schema for `vip.only_not_flagged_cfg`

Optional:

- `only_not_flagged` (Number)
- `route_map` (String)


