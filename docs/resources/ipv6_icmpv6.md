---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ipv6_icmpv6 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_ipv6_icmpv6 (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ipv6_icmpv6" "testname" {
  redirect    = 0
  unreachable = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `redirect` (Number)
- `unreachable` (Number)
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.


