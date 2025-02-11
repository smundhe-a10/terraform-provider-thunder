---
layout: ""
page_title: "Provider: Thunder"
description: |-
  The Thunder provider provides resources to interact with a A10 Thunder device via axAPI.
---

# Thunder Provider

The Thunder provider provides resources to interact with a A10 Thunder device via axAPI.

## Example Usage

```terraform
provider "thunder" {
    address  = var.dut
    username = var.username
    password = var.password
    partition = var.partition_name # Into which partition. The default value is "shared".
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String) Domain name/IP of the THUNDER
- `password` (String) The user's password
- `username` (String) Username with API access to the THUNDER

### Optional

- `partition` (String) partition name
