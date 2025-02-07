---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_snmpv1_v2c_user_oid Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_snmp_server_snmpv1_v2c_user_oid (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_snmpv1_v2c_user_oid" "Snmp_Server_SNMPv1_V2c_User_Oid_Test" {
  remote {
    host_list {
      dns_host = "test"
    }
    ipv4_list {
      ipv4_host = "192.168.171.0"
      ipv4_mask = "255.255.255.0"
    }
    ipv6_list {
      ipv6_host = "2000:0db8:85a3:0000:0000:8a2e:0370:7337"
      ipv6_mask = 128
    }
  }
  oid_val  = "test"
  user_tag = "testing_user"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `oid` (String)
- `oid_val` (String)
- `remote` (Block List, Max: 1) (see [below for nested schema](#nestedblock--remote))
- `user` (String)
- `user_tag` (String)
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--remote"></a>
### Nested Schema for `remote`

Optional:

- `host_list` (Block List) (see [below for nested schema](#nestedblock--remote--host_list))
- `ipv4_list` (Block List) (see [below for nested schema](#nestedblock--remote--ipv4_list))
- `ipv6_list` (Block List) (see [below for nested schema](#nestedblock--remote--ipv6_list))

<a id="nestedblock--remote--host_list"></a>
### Nested Schema for `remote.host_list`

Optional:

- `dns_host` (String)
- `ipv4_mask` (String)


<a id="nestedblock--remote--ipv4_list"></a>
### Nested Schema for `remote.ipv4_list`

Optional:

- `ipv4_host` (String)
- `ipv4_mask` (String)


<a id="nestedblock--remote--ipv6_list"></a>
### Nested Schema for `remote.ipv6_list`

Optional:

- `ipv6_host` (String)
- `ipv6_mask` (Number)


