---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_server_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_server_oper: Operational Status for the object server
  PLACEHOLDER
---

# thunder_slb_server_oper (Data Source)

`thunder_slb_server_oper`: Operational Status for the object server

__PLACEHOLDER__



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Server Name

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))
- `port_list` (Block List) (see [below for nested schema](#nestedblock--port_list))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `conn_rate_unit` (String)
- `creation_type` (String)
- `curr_conn_rate` (Number)
- `curr_observe_rate` (Number)
- `disable` (Number)
- `dns_update_time` (String)
- `drs_list` (Block List) (see [below for nested schema](#nestedblock--oper--drs_list))
- `is_autocreate` (Number)
- `server_ttl` (Number)
- `slow_start_conn_limit` (Number)
- `srv_gateway_arp` (String)
- `state` (String)
- `weight` (Number)

<a id="nestedblock--oper--drs_list"></a>
### Nested Schema for `oper.drs_list`

Optional:

- `drs_conn_rate_unit` (String)
- `drs_creation_type` (String)
- `drs_curr_conn` (Number)
- `drs_curr_conn_rate` (Number)
- `drs_curr_observe_rate` (Number)
- `drs_curr_req` (Number)
- `drs_disable` (Number)
- `drs_dns_update_time` (String)
- `drs_host` (String)
- `drs_is_autocreate` (Number)
- `drs_name` (String)
- `drs_peak_conn` (Number)
- `drs_server_ipv6_addr` (String)
- `drs_server_ttl` (Number)
- `drs_slow_start_conn_limit` (Number)
- `drs_srv_gateway_arp` (String)
- `drs_state` (String)
- `drs_tot_conn` (Number)
- `drs_tot_fwd_bytes` (Number)
- `drs_tot_fwd_pkts` (Number)
- `drs_tot_req` (Number)
- `drs_tot_req_suc` (Number)
- `drs_tot_rev_bytes` (Number)
- `drs_tot_rev_pkts` (Number)
- `drs_weight` (Number)



<a id="nestedblock--port_list"></a>
### Nested Schema for `port_list`

Required:

- `port_number` (Number) Port Number
- `protocol` (String) 'tcp': TCP Port; 'udp': UDP Port;

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--port_list--oper))

<a id="nestedblock--port_list--oper"></a>
### Nested Schema for `port_list.oper`

Optional:

- `aflow_conn_limit` (Number)
- `aflow_queue_size` (Number)
- `auto_nat_addr_list` (Block List) (see [below for nested schema](#nestedblock--port_list--oper--auto_nat_addr_list))
- `conn_rate_unit` (String)
- `curr_conn_rate` (Number)
- `curr_observe_rate` (Number)
- `current_time` (Number)
- `diameter_enabled` (Number)
- `disable` (Number)
- `down_grace_period_allowed` (Number)
- `down_time_grace_period` (Number)
- `drs_auto_nat_list` (Block List) (see [below for nested schema](#nestedblock--port_list--oper--drs_auto_nat_list))
- `drs_ip_nat_list` (Block List) (see [below for nested schema](#nestedblock--port_list--oper--drs_ip_nat_list))
- `es_resp_time` (Number)
- `hm_index` (Number)
- `hm_key` (Number)
- `inband_hm_reassign_num` (Number)
- `nat_pool_addr_list` (Block List) (see [below for nested schema](#nestedblock--port_list--oper--nat_pool_addr_list))
- `pool_name` (String)
- `resv_conn` (Number)
- `slow_start_conn_limit` (Number)
- `soft_down_time` (Number)
- `state` (String)

<a id="nestedblock--port_list--oper--auto_nat_addr_list"></a>
### Nested Schema for `port_list.oper.auto_nat_addr_list`

Optional:

- `alloc_failed` (Number)
- `auto_nat_ip` (String)
- `ha_group_id` (Number)
- `ip_rr` (Number)
- `ports_consumed` (Number)
- `ports_consumed_total` (Number)
- `ports_freed_total` (Number)
- `vrid` (Number)


<a id="nestedblock--port_list--oper--drs_auto_nat_list"></a>
### Nested Schema for `port_list.oper.drs_auto_nat_list`

Optional:

- `drs_auto_nat_address_list` (Block List) (see [below for nested schema](#nestedblock--port_list--oper--drs_auto_nat_list--drs_auto_nat_address_list))
- `drs_name` (String)
- `drs_port` (Number)

<a id="nestedblock--port_list--oper--drs_auto_nat_list--drs_auto_nat_address_list"></a>
### Nested Schema for `port_list.oper.drs_auto_nat_list.drs_auto_nat_address_list`

Optional:

- `alloc_failed` (Number)
- `auto_nat_ip` (String)
- `ha_group_id` (Number)
- `ip_rr` (Number)
- `ports_consumed` (Number)
- `ports_consumed_total` (Number)
- `ports_freed_total` (Number)
- `vrid` (Number)



<a id="nestedblock--port_list--oper--drs_ip_nat_list"></a>
### Nested Schema for `port_list.oper.drs_ip_nat_list`

Optional:

- `drs_name` (String)
- `drs_port` (Number)
- `nat_pool_addr_list` (Block List) (see [below for nested schema](#nestedblock--port_list--oper--drs_ip_nat_list--nat_pool_addr_list))
- `pool_name` (String)

<a id="nestedblock--port_list--oper--drs_ip_nat_list--nat_pool_addr_list"></a>
### Nested Schema for `port_list.oper.drs_ip_nat_list.nat_pool_addr_list`

Optional:

- `alloc_failed` (Number)
- `nat_ip` (String)
- `ports_consumed` (Number)
- `ports_consumed_total` (Number)
- `ports_freed_total` (Number)



<a id="nestedblock--port_list--oper--nat_pool_addr_list"></a>
### Nested Schema for `port_list.oper.nat_pool_addr_list`

Optional:

- `alloc_failed` (Number)
- `nat_ip` (String)
- `ports_consumed` (Number)
- `ports_consumed_total` (Number)
- `ports_freed_total` (Number)


