---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_view Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_snmp_server_view (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_view" "resourceSnmpServerViewTest" {
  type     = "excluded"
  oid      = "123"
  viewname = "a10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `mask` (String)
- `oid` (String)
- `type` (String)
- `uuid` (String)
- `viewname` (String)

### Read-Only

- `id` (String) The ID of this resource.


