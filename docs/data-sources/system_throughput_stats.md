---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_throughput_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_throughput_stats: Statistics for the object throughput
  PLACEHOLDER
---

# thunder_system_throughput_stats (Data Source)

`thunder_system_throughput_stats`: Statistics for the object throughput

__PLACEHOLDER__



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `global_system_throughput_bits_per_sec` (Number) Global System throughput in bits/sec
- `per_part_throughput_bits_per_sec` (Number) Partition throughput in bits/sec


