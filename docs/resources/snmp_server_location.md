---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_location Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_snmp_server_location (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_location" "resourceSnmpServerLocationTest" {
  loc = "location_testing"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `loc` (String)
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.


