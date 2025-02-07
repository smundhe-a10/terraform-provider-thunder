---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_interface_lif Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_interface_lif: Logical interface
---

# thunder_interface_lif (Resource)

`thunder_interface_lif`: Logical interface

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

provider "thunder" {
  alias     = "L3V_A"
  address   = var.dut9049
  username  = var.username
  password  = var.password
  partition = var.l3v_1
}

resource "thunder_interface_lif" "test00" {
  ifname = 2
  ip {
    address_list {
      ipv4_address = "56.56.56.56"
      ipv4_netmask = "/16"
    }
  }
}

resource "thunder_interface_lif" "test01" {
  provider = thunder.L3V_A
  ifname   = 7
  ip {
    address_list {
      ipv4_address = "56.156.0.1"
      ipv4_netmask = "/24"
    }
  }
}

resource "thunder_interface_lif" "test02" {
  provider = thunder.L3V_A
  ifname   = 8
  ipv6 {
    address_list {
      ipv6_addr = "fd08::222/64"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ifname` (String) Lif interface name

### Optional

- `access_list` (Block List, Max: 1) (see [below for nested schema](#nestedblock--access_list))
- `action` (String) 'enable': Enable; 'disable': Disable;
- `bfd` (Block List, Max: 1) (see [below for nested schema](#nestedblock--bfd))
- `ip` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip))
- `ipv6` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ipv6))
- `isis` (Block List, Max: 1) (see [below for nested schema](#nestedblock--isis))
- `mtu` (Number) Interface mtu (Interface MTU, default 1 (min MTU is 1280 for IPv6))
- `name` (String) Name for the interface
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--access_list"></a>
### Nested Schema for `access_list`

Optional:

- `acl_id` (Number) ACL id
- `acl_name` (String) Apply an access list (Named Access List)


<a id="nestedblock--bfd"></a>
### Nested Schema for `bfd`

Optional:

- `authentication` (Block List, Max: 1) (see [below for nested schema](#nestedblock--bfd--authentication))
- `demand` (Number) Demand mode
- `echo` (Number) Enable BFD Echo
- `interval_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--bfd--interval_cfg))
- `uuid` (String) uuid of the object

<a id="nestedblock--bfd--authentication"></a>
### Nested Schema for `bfd.authentication`

Optional:

- `key_id` (Number) Key ID
- `method` (String) 'md5': Keyed MD5; 'meticulous-md5': Meticulous Keyed MD5; 'meticulous-sha1': Meticulous Keyed SHA1; 'sha1': Keyed SHA1; 'simple': Simple Password;
- `password` (String) Key String


<a id="nestedblock--bfd--interval_cfg"></a>
### Nested Schema for `bfd.interval_cfg`

Optional:

- `interval` (Number) Transmit interval between BFD packets (Milliseconds)
- `min_rx` (Number) Minimum receive interval capability (Milliseconds)
- `multiplier` (Number) Multiplier value used to compute holddown (value used to multiply the interval)



<a id="nestedblock--ip"></a>
### Nested Schema for `ip`

Optional:

- `address_list` (Block List) (see [below for nested schema](#nestedblock--ip--address_list))
- `allow_promiscuous_vip` (Number) Allow traffic to be associated with promiscuous VIP
- `cache_spoofing_port` (Number) This interface connects to spoofing cache
- `dhcp` (Number) Use DHCP to configure IP address
- `generate_membership_query` (Number) Enable Membership Query
- `inside` (Number) Configure interface as inside
- `max_resp_time` (Number) Maximum Response Time (Max Response Time (Default is 100))
- `ospf` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--ospf))
- `outside` (Number) Configure interface as outside
- `query_interval` (Number) 1 - 255 (Default is 125)
- `rip` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--rip))
- `router` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--router))
- `unnumbered` (Number) Set the interface as unnumbered
- `uuid` (String) uuid of the object

<a id="nestedblock--ip--address_list"></a>
### Nested Schema for `ip.address_list`

Optional:

- `ipv4_address` (String) IP address
- `ipv4_netmask` (String) IP subnet mask


<a id="nestedblock--ip--ospf"></a>
### Nested Schema for `ip.ospf`

Optional:

- `ospf_global` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--ospf--ospf_global))
- `ospf_ip_list` (Block List) (see [below for nested schema](#nestedblock--ip--ospf--ospf_ip_list))

<a id="nestedblock--ip--ospf--ospf_global"></a>
### Nested Schema for `ip.ospf.ospf_global`

Optional:

- `authentication_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--ospf--ospf_global--authentication_cfg))
- `authentication_key` (String) Authentication password (key) (The OSPF password (key))
- `bfd_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--ospf--ospf_global--bfd_cfg))
- `cost` (Number) Interface cost
- `database_filter_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--ospf--ospf_global--database_filter_cfg))
- `dead_interval` (Number) Interval after which a neighbor is declared dead (Seconds)
- `disable` (String) 'all': All functionality;
- `hello_interval` (Number) Time between HELLO packets (Seconds)
- `message_digest_cfg` (Block List) (see [below for nested schema](#nestedblock--ip--ospf--ospf_global--message_digest_cfg))
- `mtu` (Number) OSPF interface MTU (MTU size)
- `mtu_ignore` (Number) Ignores the MTU in DBD packets
- `network` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--ospf--ospf_global--network))
- `priority` (Number) Router priority
- `retransmit_interval` (Number) Time between retransmitting lost link state advertisements (Seconds)
- `transmit_delay` (Number) Link state transmit delay (Seconds)
- `uuid` (String) uuid of the object

<a id="nestedblock--ip--ospf--ospf_global--authentication_cfg"></a>
### Nested Schema for `ip.ospf.ospf_global.authentication_cfg`

Optional:

- `authentication` (Number) Enable authentication
- `value` (String) 'message-digest': Use message-digest authentication; 'null': Use no authentication;


<a id="nestedblock--ip--ospf--ospf_global--bfd_cfg"></a>
### Nested Schema for `ip.ospf.ospf_global.bfd_cfg`

Optional:

- `bfd` (Number) Bidirectional Forwarding Detection (BFD)
- `disable` (Number) Disable BFD


<a id="nestedblock--ip--ospf--ospf_global--database_filter_cfg"></a>
### Nested Schema for `ip.ospf.ospf_global.database_filter_cfg`

Optional:

- `database_filter` (String) 'all': Filter all LSA;
- `out` (Number) Outgoing LSA


<a id="nestedblock--ip--ospf--ospf_global--message_digest_cfg"></a>
### Nested Schema for `ip.ospf.ospf_global.message_digest_cfg`

Optional:

- `md5` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--ospf--ospf_global--message_digest_cfg--md5))
- `message_digest_key` (Number) Message digest authentication password (key) (Key id)

<a id="nestedblock--ip--ospf--ospf_global--message_digest_cfg--md5"></a>
### Nested Schema for `ip.ospf.ospf_global.message_digest_cfg.md5`

Optional:

- `md5_value` (String) The OSPF password (1-16)



<a id="nestedblock--ip--ospf--ospf_global--network"></a>
### Nested Schema for `ip.ospf.ospf_global.network`

Optional:

- `broadcast` (Number) Specify OSPF broadcast multi-access network
- `non_broadcast` (Number) Specify OSPF NBMA network
- `p2mp_nbma` (Number) Specify non-broadcast point-to-multipoint network
- `point_to_multipoint` (Number) Specify OSPF point-to-multipoint network
- `point_to_point` (Number) Specify OSPF point-to-point network



<a id="nestedblock--ip--ospf--ospf_ip_list"></a>
### Nested Schema for `ip.ospf.ospf_ip_list`

Required:

- `ip_addr` (String) Address of interface

Optional:

- `authentication` (Number) Enable authentication
- `authentication_key` (String) Authentication password (key) (The OSPF password (key))
- `cost` (Number) Interface cost
- `database_filter` (String) 'all': Filter all LSA;
- `dead_interval` (Number) Interval after which a neighbor is declared dead (Seconds)
- `hello_interval` (Number) Time between HELLO packets (Seconds)
- `message_digest_cfg` (Block List) (see [below for nested schema](#nestedblock--ip--ospf--ospf_ip_list--message_digest_cfg))
- `mtu_ignore` (Number) Ignores the MTU in DBD packets
- `out` (Number) Outgoing LSA
- `priority` (Number) Router priority
- `retransmit_interval` (Number) Time between retransmitting lost link state advertisements (Seconds)
- `transmit_delay` (Number) Link state transmit delay (Seconds)
- `uuid` (String) uuid of the object
- `value` (String) 'message-digest': Use message-digest authentication; 'null': Use no authentication;

<a id="nestedblock--ip--ospf--ospf_ip_list--message_digest_cfg"></a>
### Nested Schema for `ip.ospf.ospf_ip_list.message_digest_cfg`

Optional:

- `md5_value` (String) The OSPF password (1-16)
- `message_digest_key` (Number) Message digest authentication password (key) (Key id)




<a id="nestedblock--ip--rip"></a>
### Nested Schema for `ip.rip`

Optional:

- `authentication` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--rip--authentication))
- `receive_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--rip--receive_cfg))
- `receive_packet` (Number) Enable receiving packet through the specified interface
- `send_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--rip--send_cfg))
- `send_packet` (Number) Enable sending packets through the specified interface
- `split_horizon_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--rip--split_horizon_cfg))
- `uuid` (String) uuid of the object

<a id="nestedblock--ip--rip--authentication"></a>
### Nested Schema for `ip.rip.authentication`

Optional:

- `key_chain` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--rip--authentication--key_chain))
- `mode` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--rip--authentication--mode))
- `str` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--rip--authentication--str))

<a id="nestedblock--ip--rip--authentication--key_chain"></a>
### Nested Schema for `ip.rip.authentication.key_chain`

Optional:

- `key_chain` (String) Authentication key-chain (Name of key-chain)


<a id="nestedblock--ip--rip--authentication--mode"></a>
### Nested Schema for `ip.rip.authentication.mode`

Optional:

- `mode` (String) 'md5': Keyed message digest; 'text': Clear text authentication;


<a id="nestedblock--ip--rip--authentication--str"></a>
### Nested Schema for `ip.rip.authentication.str`

Optional:

- `string` (String) The RIP authentication string



<a id="nestedblock--ip--rip--receive_cfg"></a>
### Nested Schema for `ip.rip.receive_cfg`

Optional:

- `receive` (Number) Advertisement reception
- `version` (String) '1': RIP version 1; '2': RIP version 2; '1-2': RIP version 1 & 2;


<a id="nestedblock--ip--rip--send_cfg"></a>
### Nested Schema for `ip.rip.send_cfg`

Optional:

- `send` (Number) Advertisement transmission
- `version` (String) '1': RIP version 1; '2': RIP version 2; '1-compatible': RIPv1-compatible; '1-2': RIP version 1 & 2;


<a id="nestedblock--ip--rip--split_horizon_cfg"></a>
### Nested Schema for `ip.rip.split_horizon_cfg`

Optional:

- `state` (String) 'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;



<a id="nestedblock--ip--router"></a>
### Nested Schema for `ip.router`

Optional:

- `isis` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip--router--isis))

<a id="nestedblock--ip--router--isis"></a>
### Nested Schema for `ip.router.isis`

Optional:

- `tag` (String) ISO routing area tag
- `uuid` (String) uuid of the object




<a id="nestedblock--ipv6"></a>
### Nested Schema for `ipv6`

Optional:

- `address_list` (Block List) (see [below for nested schema](#nestedblock--ipv6--address_list))
- `inside` (Number) Configure interface as inside
- `ipv6_enable` (Number) Enable IPv6 processing
- `ospf` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ipv6--ospf))
- `outside` (Number) Configure interface as outside
- `router` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ipv6--router))
- `uuid` (String) uuid of the object

<a id="nestedblock--ipv6--address_list"></a>
### Nested Schema for `ipv6.address_list`

Optional:

- `anycast` (Number) Configure an IPv6 anycast address
- `ipv6_addr` (String) Set the IPv6 address of an interface
- `link_local` (Number) Configure an IPv6 link local address


<a id="nestedblock--ipv6--ospf"></a>
### Nested Schema for `ipv6.ospf`

Optional:

- `bfd` (Number) Bidirectional Forwarding Detection (BFD)
- `cost_cfg` (Block List) (see [below for nested schema](#nestedblock--ipv6--ospf--cost_cfg))
- `dead_interval_cfg` (Block List) (see [below for nested schema](#nestedblock--ipv6--ospf--dead_interval_cfg))
- `disable` (Number) Disable BFD
- `hello_interval_cfg` (Block List) (see [below for nested schema](#nestedblock--ipv6--ospf--hello_interval_cfg))
- `mtu_ignore_cfg` (Block List) (see [below for nested schema](#nestedblock--ipv6--ospf--mtu_ignore_cfg))
- `neighbor_cfg` (Block List) (see [below for nested schema](#nestedblock--ipv6--ospf--neighbor_cfg))
- `network_list` (Block List) (see [below for nested schema](#nestedblock--ipv6--ospf--network_list))
- `priority_cfg` (Block List) (see [below for nested schema](#nestedblock--ipv6--ospf--priority_cfg))
- `retransmit_interval_cfg` (Block List) (see [below for nested schema](#nestedblock--ipv6--ospf--retransmit_interval_cfg))
- `transmit_delay_cfg` (Block List) (see [below for nested schema](#nestedblock--ipv6--ospf--transmit_delay_cfg))
- `uuid` (String) uuid of the object

<a id="nestedblock--ipv6--ospf--cost_cfg"></a>
### Nested Schema for `ipv6.ospf.cost_cfg`

Optional:

- `cost` (Number) Interface cost
- `instance_id` (Number) Specify the interface instance ID


<a id="nestedblock--ipv6--ospf--dead_interval_cfg"></a>
### Nested Schema for `ipv6.ospf.dead_interval_cfg`

Optional:

- `dead_interval` (Number) Interval after which a neighbor is declared dead (Seconds)
- `instance_id` (Number) Specify the interface instance ID


<a id="nestedblock--ipv6--ospf--hello_interval_cfg"></a>
### Nested Schema for `ipv6.ospf.hello_interval_cfg`

Optional:

- `hello_interval` (Number) Time between HELLO packets (Seconds)
- `instance_id` (Number) Specify the interface instance ID


<a id="nestedblock--ipv6--ospf--mtu_ignore_cfg"></a>
### Nested Schema for `ipv6.ospf.mtu_ignore_cfg`

Optional:

- `instance_id` (Number) Specify the interface instance ID
- `mtu_ignore` (Number) Ignores the MTU in DBD packets


<a id="nestedblock--ipv6--ospf--neighbor_cfg"></a>
### Nested Schema for `ipv6.ospf.neighbor_cfg`

Optional:

- `neig_inst` (Number) Specify the interface instance ID
- `neighbor` (String) OSPFv3 neighbor (Neighbor IPv6 address)
- `neighbor_cost` (Number) OSPF cost for point-to-multipoint neighbor (metric)
- `neighbor_poll_interval` (Number) OSPF dead-router polling interval (Seconds)
- `neighbor_priority` (Number) OSPF priority of non-broadcast neighbor


<a id="nestedblock--ipv6--ospf--network_list"></a>
### Nested Schema for `ipv6.ospf.network_list`

Optional:

- `broadcast_type` (String) 'broadcast': Specify OSPF broadcast multi-access network; 'non-broadcast': Specify OSPF NBMA network; 'point-to-point': Specify OSPF point-to-point network; 'point-to-multipoint': Specify OSPF point-to-multipoint network;
- `network_instance_id` (Number) Specify the interface instance ID
- `p2mp_nbma` (Number) Specify non-broadcast point-to-multipoint network


<a id="nestedblock--ipv6--ospf--priority_cfg"></a>
### Nested Schema for `ipv6.ospf.priority_cfg`

Optional:

- `instance_id` (Number) Specify the interface instance ID
- `priority` (Number) Router priority


<a id="nestedblock--ipv6--ospf--retransmit_interval_cfg"></a>
### Nested Schema for `ipv6.ospf.retransmit_interval_cfg`

Optional:

- `instance_id` (Number) Specify the interface instance ID
- `retransmit_interval` (Number) Time between retransmitting lost link state advertisements (Seconds)


<a id="nestedblock--ipv6--ospf--transmit_delay_cfg"></a>
### Nested Schema for `ipv6.ospf.transmit_delay_cfg`

Optional:

- `instance_id` (Number) Specify the interface instance ID
- `transmit_delay` (Number) Link state transmit delay (Seconds)



<a id="nestedblock--ipv6--router"></a>
### Nested Schema for `ipv6.router`

Optional:

- `isis` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ipv6--router--isis))
- `ospf` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ipv6--router--ospf))
- `ripng` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ipv6--router--ripng))

<a id="nestedblock--ipv6--router--isis"></a>
### Nested Schema for `ipv6.router.isis`

Optional:

- `tag` (String) ISO routing area tag
- `uuid` (String) uuid of the object


<a id="nestedblock--ipv6--router--ospf"></a>
### Nested Schema for `ipv6.router.ospf`

Optional:

- `area_list` (Block List) (see [below for nested schema](#nestedblock--ipv6--router--ospf--area_list))
- `uuid` (String) uuid of the object

<a id="nestedblock--ipv6--router--ospf--area_list"></a>
### Nested Schema for `ipv6.router.ospf.area_list`

Optional:

- `area_id_addr` (String) OSPF area ID in IP address format
- `area_id_num` (Number) OSPF area ID as a decimal value
- `instance_id` (Number) Set the interface instance ID
- `tag` (String) Set the OSPFv3 process tag



<a id="nestedblock--ipv6--router--ripng"></a>
### Nested Schema for `ipv6.router.ripng`

Optional:

- `rip` (Number) RIP Routing for IPv6
- `uuid` (String) uuid of the object




<a id="nestedblock--isis"></a>
### Nested Schema for `isis`

Optional:

- `authentication` (Block List, Max: 1) (see [below for nested schema](#nestedblock--isis--authentication))
- `bfd_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--isis--bfd_cfg))
- `circuit_type` (String) 'level-1': Level-1 only adjacencies are formed; 'level-1-2': Level-1-2 adjacencies are formed; 'level-2-only': Level-2 only adjacencies are formed;
- `csnp_interval_list` (Block List) (see [below for nested schema](#nestedblock--isis--csnp_interval_list))
- `hello_interval_list` (Block List) (see [below for nested schema](#nestedblock--isis--hello_interval_list))
- `hello_interval_minimal_list` (Block List) (see [below for nested schema](#nestedblock--isis--hello_interval_minimal_list))
- `hello_multiplier_list` (Block List) (see [below for nested schema](#nestedblock--isis--hello_multiplier_list))
- `lsp_interval` (Number) Set LSP transmission interval (LSP transmission interval (milliseconds))
- `mesh_group` (Block List, Max: 1) (see [below for nested schema](#nestedblock--isis--mesh_group))
- `metric_list` (Block List) (see [below for nested schema](#nestedblock--isis--metric_list))
- `network` (String) 'broadcast': Specify IS-IS broadcast multi-access network; 'point-to-point': Specify IS-IS point-to-point network;
- `padding` (Number) Add padding to IS-IS hello packets
- `password_list` (Block List) (see [below for nested schema](#nestedblock--isis--password_list))
- `priority_list` (Block List) (see [below for nested schema](#nestedblock--isis--priority_list))
- `retransmit_interval` (Number) Set per-LSP retransmission interval (Interval between retransmissions of the same LSP (seconds))
- `uuid` (String) uuid of the object
- `wide_metric_list` (Block List) (see [below for nested schema](#nestedblock--isis--wide_metric_list))

<a id="nestedblock--isis--authentication"></a>
### Nested Schema for `isis.authentication`

Optional:

- `key_chain_list` (Block List) (see [below for nested schema](#nestedblock--isis--authentication--key_chain_list))
- `mode_list` (Block List) (see [below for nested schema](#nestedblock--isis--authentication--mode_list))
- `send_only_list` (Block List) (see [below for nested schema](#nestedblock--isis--authentication--send_only_list))

<a id="nestedblock--isis--authentication--key_chain_list"></a>
### Nested Schema for `isis.authentication.key_chain_list`

Optional:

- `key_chain` (String) Authentication key-chain (Name of key-chain)
- `level` (String) 'level-1': Specify authentication key-chain for level-1 PDUs; 'level-2': Specify authentication key-chain for level-2 PDUs;


<a id="nestedblock--isis--authentication--mode_list"></a>
### Nested Schema for `isis.authentication.mode_list`

Optional:

- `level` (String) 'level-1': Specify authentication mode for level-1 PDUs; 'level-2': Specify authentication mode for level-2 PDUs;
- `mode` (String) 'md5': Keyed message digest;


<a id="nestedblock--isis--authentication--send_only_list"></a>
### Nested Schema for `isis.authentication.send_only_list`

Optional:

- `level` (String) 'level-1': Specify authentication send-only for level-1 PDUs; 'level-2': Specify authentication send-only for level-2 PDUs;
- `send_only` (Number) Authentication send-only



<a id="nestedblock--isis--bfd_cfg"></a>
### Nested Schema for `isis.bfd_cfg`

Optional:

- `bfd` (Number) Bidirectional Forwarding Detection (BFD)
- `disable` (Number) Disable BFD


<a id="nestedblock--isis--csnp_interval_list"></a>
### Nested Schema for `isis.csnp_interval_list`

Optional:

- `csnp_interval` (Number) Set CSNP interval in seconds (CSNP interval value)
- `level` (String) 'level-1': Speficy interval for level-1 CSNPs; 'level-2': Specify interval for level-2 CSNPs;


<a id="nestedblock--isis--hello_interval_list"></a>
### Nested Schema for `isis.hello_interval_list`

Optional:

- `hello_interval` (Number) Set Hello interval in seconds (Hello interval value)
- `level` (String) 'level-1': Specify hello-interval for level-1 IIHs; 'level-2': Specify hello-interval for level-2 IIHs;


<a id="nestedblock--isis--hello_interval_minimal_list"></a>
### Nested Schema for `isis.hello_interval_minimal_list`

Optional:

- `hello_interval_minimal` (Number) Set Hello holdtime 1 second, interval depends on multiplier
- `level` (String) 'level-1': Specify hello-interval for level-1 IIHs; 'level-2': Specify hello-interval for level-2 IIHs;


<a id="nestedblock--isis--hello_multiplier_list"></a>
### Nested Schema for `isis.hello_multiplier_list`

Optional:

- `hello_multiplier` (Number) Set multiplier for Hello holding time (Hello multiplier value)
- `level` (String) 'level-1': Specify hello multiplier for level-1 IIHs; 'level-2': Specify hello multiplier for level-2 IIHs;


<a id="nestedblock--isis--mesh_group"></a>
### Nested Schema for `isis.mesh_group`

Optional:

- `blocked` (Number) Block LSPs on this interface
- `value` (Number) Mesh group number


<a id="nestedblock--isis--metric_list"></a>
### Nested Schema for `isis.metric_list`

Optional:

- `level` (String) 'level-1': Apply metric to level-1 links; 'level-2': Apply metric to level-2 links;
- `metric` (Number) Configure the metric for interface (Default metric)


<a id="nestedblock--isis--password_list"></a>
### Nested Schema for `isis.password_list`

Optional:

- `level` (String) 'level-1': Specify password for level-1 PDUs; 'level-2': Specify password for level-2 PDUs;
- `password` (String) Configure the authentication password for interface


<a id="nestedblock--isis--priority_list"></a>
### Nested Schema for `isis.priority_list`

Optional:

- `level` (String) 'level-1': Specify priority for level-1 routing; 'level-2': Specify priority for level-2 routing;
- `priority` (Number) Set priority for Designated Router election (Priority value)


<a id="nestedblock--isis--wide_metric_list"></a>
### Nested Schema for `isis.wide_metric_list`

Optional:

- `level` (String) 'level-1': Apply metric to level-1 links; 'level-2': Apply metric to level-2 links;
- `wide_metric` (Number) Configure the wide metric for interface



<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'num_pkts': some help string; 'num_total_bytes': some help string; 'num_unicast_pkts': some help string; 'num_broadcast_pkts': some help string; 'num_multicast_pkts': some help string; 'num_tx_pkts': some help string; 'num_total_tx_bytes': some help string; 'num_unicast_tx_pkts': some help string; 'num_broadcast_tx_pkts': some help string; 'num_multicast_tx_pkts': some help string; 'dropped_dis_rx_pkts': some help string; 'dropped_rx_pkts': some help string; 'dropped_dis_tx_pkts': some help string; 'dropped_tx_pkts': some help string; 'dropped_rx_pkts_gre_key': some help string;


