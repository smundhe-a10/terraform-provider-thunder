---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cloud_services_cloud_provider_vmware_log Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cloud_services_cloud_provider_vmware_log: VMware log configuration
  PLACEHOLDER
---

# thunder_cloud_services_cloud_provider_vmware_log (Resource)

`thunder_cloud_services_cloud_provider_vmware_log`: VMware log configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_cloud_services_cloud_provider_vmware_log" "test" {
  action = "enable"
  vrli_host = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (String) 'enable': Enable VMware Log Analytics; 'disable': Disable VMware Log Analytics(default);
- `uuid` (String) uuid of the object
- `vrli_host` (String) VMware Log Analytics vrli-host

### Read-Only

- `id` (String) The ID of this resource.


