---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_alg_sip Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_fw_alg_sip (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_alg_sip" "test_thunder_fw_alg_sip" {
  default_port_disable = "default-port-disable"
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `default_port_disable` (String)
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String)


