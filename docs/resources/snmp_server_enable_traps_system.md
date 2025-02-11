---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_enable_traps_system Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_snmp_server_enable_traps_system (Resource)



## Example Usage

```terraform
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_system"   "SnMPServerEnableTrapsSystem"  {
    all= 1
    control_cpu_high= 1
    data_cpu_high= 0
    fan= 0
    file_sys_read_only= 0
    high_disk_use= 0
    high_memory_use= 0
    high_temp= 1
    low_temp= 0
    license_management= 0
    packet_drop= 0
    power= 0
    pri_disk= 0
    restart= 0
    sec_disk= 0
    shutdown= 1
    smp_resource_event= 0
    syslog_severity_one= 0
    tacacs_server_up_down= 0
    start= 0
  
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `all` (Number)
- `control_cpu_high` (Number)
- `data_cpu_high` (Number)
- `fan` (Number)
- `file_sys_read_only` (Number)
- `high_disk_use` (Number)
- `high_memory_use` (Number)
- `high_temp` (Number)
- `license_management` (Number)
- `low_temp` (Number)
- `packet_drop` (Number)
- `power` (Number)
- `pri_disk` (Number)
- `restart` (Number)
- `sec_disk` (Number)
- `shutdown` (Number)
- `smp_resource_event` (Number)
- `start` (Number)
- `syslog_severity_one` (Number)
- `tacacs_server_up_down` (Number)
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.


