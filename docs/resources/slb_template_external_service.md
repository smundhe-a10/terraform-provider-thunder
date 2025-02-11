---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_external_service Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_slb_template_external_service (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_external_service" "test_thunder_slb_template_external_service" {
  name = "ext_svc"
  type = "url-filter"
  request_header_forward_list {
    request_header_forward = "req_head"
  }
  failure_action = "drop"
  timeout        = 40
  action         = "drop"
  service_group  = "sg1"
  user_tag       = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (String)
- `bypass_ip_cfg` (Block List) (see [below for nested schema](#nestedblock--bypass_ip_cfg))
- `failure_action` (String)
- `name` (String)
- `request_header_forward_list` (Block List) (see [below for nested schema](#nestedblock--request_header_forward_list))
- `service_group` (String)
- `shared_partition_persist_source_ip_template` (Number)
- `shared_partition_tcp_proxy_template` (Number)
- `source_ip` (String)
- `tcp_proxy` (String)
- `template_persist_source_ip_shared` (String)
- `template_tcp_proxy_shared` (String)
- `timeout` (Number)
- `type` (String)
- `user_tag` (String)
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--bypass_ip_cfg"></a>
### Nested Schema for `bypass_ip_cfg`

Optional:

- `bypass_ip` (String)
- `mask` (String)


<a id="nestedblock--request_header_forward_list"></a>
### Nested Schema for `request_header_forward_list`

Optional:

- `request_header_forward` (String)


