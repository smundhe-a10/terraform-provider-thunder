---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_enable_traps_vcs Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_snmp_server_enable_traps_vcs (Resource)



## Example Usage

```terraform
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_vcs"   "SnMPServerEnableTrapsVcs"  {
    state_change= 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `state_change` (Number)
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.


