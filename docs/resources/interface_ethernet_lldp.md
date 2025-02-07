---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_interface_ethernet_lldp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_interface_ethernet_lldp (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_ethernet_lldp" "lldp" {
  ifnum = 5
  notification_cfg {
    notif_enable = 1
    notification = 1
  }
  enable_cfg {
    tx        = 1
    rx        = 1
    rt_enable = 1
  }
  tx_dot1_cfg {
    link_aggregation = 1
    vlan             = 1
    tx_dot1_tlvs     = 1
  }
  tx_tlvs_cfg {
    tx_tlvs             = 1
    port_description    = 1
    exclude             = 1
    management_address  = 1
    system_name         = 1
    system_capabilities = 1
    system_description  = 1
  }

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `enable_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--enable_cfg))
- `ifnum` (Number)
- `notification_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--notification_cfg))
- `tx_dot1_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--tx_dot1_cfg))
- `tx_tlvs_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--tx_tlvs_cfg))
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--enable_cfg"></a>
### Nested Schema for `enable_cfg`

Optional:

- `rt_enable` (Number)
- `rx` (Number)
- `tx` (Number)


<a id="nestedblock--notification_cfg"></a>
### Nested Schema for `notification_cfg`

Optional:

- `notif_enable` (Number)
- `notification` (Number)


<a id="nestedblock--tx_dot1_cfg"></a>
### Nested Schema for `tx_dot1_cfg`

Optional:

- `link_aggregation` (Number)
- `tx_dot1_tlvs` (Number)
- `vlan` (Number)


<a id="nestedblock--tx_tlvs_cfg"></a>
### Nested Schema for `tx_tlvs_cfg`

Optional:

- `exclude` (Number)
- `management_address` (Number)
- `port_description` (Number)
- `system_capabilities` (Number)
- `system_description` (Number)
- `system_name` (Number)
- `tx_tlvs` (Number)


