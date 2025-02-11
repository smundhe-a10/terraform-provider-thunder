---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_tcp_proxy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_slb_template_tcp_proxy (Resource)



## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_tcp_proxy" "test" {
  name                       = "tcp_proxy1"
  server_down_action         = "FIN"
  dynamic_buffer_allocation  = 1
  nagle                      = 1
  initial_window_size        = 65535
  reset_rev                  = 1
  disable                    = 1
  maxburst                   = 90
  half_close_idle_timeout    = 100
  ack_aggressiveness         = "low"
  backend_wscale             = 14
  del_session_on_server_down = 1
  disable_abc                = 1
  disable_sack               = 1
  disable_tcp_timestamps     = 1
  disable_window_scale       = 1
  early_retransmit           = 1
  fin_timeout                = 55
  force_delete_timeout       = 3
  half_open_idle_timeout     = 30
  idle_timeout               = 2097120
  init_cwnd                  = 15
  insert_client_ip           = 1
  invalid_rate_limit         = 60000000
  keepalive_interval         = 12000
  keepalive_probes           = 5
  limited_slowstart          = 214748364
  min_rto                    = 900
  mss                        = 8999
  psh_flag_optimization      = 1
  qos                        = 63
  reassembly_timeout         = 299
  reassembly_limit           = 5
  receive_buffer             = 2147483647
  reno                       = 1
  reset_fwd                  = 1
  retransmit_retries         = 15
  syn_retries                = 15
  timewait                   = 59
  transmit_buffer            = 2147483647
  user_tag                   = "modify tcp proxy template"
  proxy_header {
    proxy_header_action = "insert"
    version             = "v1"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ack_aggressiveness` (String)
- `alive_if_active` (Number)
- `backend_wscale` (Number)
- `del_session_on_server_down` (Number)
- `disable` (Number)
- `disable_abc` (Number)
- `disable_sack` (Number)
- `disable_tcp_timestamps` (Number)
- `disable_window_scale` (Number)
- `down` (Number)
- `dynamic_buffer_allocation` (Number)
- `early_retransmit` (Number)
- `fin_timeout` (Number)
- `force_delete_timeout` (Number)
- `force_delete_timeout_100ms` (Number)
- `half_close_idle_timeout` (Number)
- `half_open_idle_timeout` (Number)
- `idle_timeout` (Number)
- `init_cwnd` (Number)
- `initial_window_size` (Number)
- `insert_client_ip` (Number)
- `invalid_rate_limit` (Number)
- `keepalive_interval` (Number)
- `keepalive_probes` (Number)
- `limited_slowstart` (Number)
- `maxburst` (Number)
- `min_rto` (Number)
- `mss` (Number)
- `nagle` (Number)
- `name` (String)
- `proxy_header` (Block List, Max: 1) (see [below for nested schema](#nestedblock--proxy_header))
- `psh_flag_optimization` (Number)
- `qos` (Number)
- `reassembly_limit` (Number)
- `reassembly_timeout` (Number)
- `receive_buffer` (Number)
- `reno` (Number)
- `reset_fwd` (Number)
- `reset_rev` (Number)
- `retransmit_retries` (Number)
- `server_down_action` (String)
- `syn_retries` (Number)
- `timewait` (Number)
- `transmit_buffer` (Number)
- `user_tag` (String)
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--proxy_header"></a>
### Nested Schema for `proxy_header`

Optional:

- `proxy_header_action` (String)
- `version` (String)


