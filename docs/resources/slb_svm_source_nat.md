---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_svm_source_nat Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_slb_svm_source_nat (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_svm_source_nat" "test_thunder_slb_svm_source_nat" {
  pool = "pool1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `pool` (String)

### Read-Only

- `id` (String) The ID of this resource.


