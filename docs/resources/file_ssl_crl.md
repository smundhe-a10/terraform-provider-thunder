---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_file_ssl_crl Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_file_ssl_crl (Resource)



## Example Usage

```terraform
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource "thunder_file_ssl_crl" "FileSslCrl"{
    name = "FILE_SSL_CRL"
    protocol = "scp"
    host = "10.64.3.186"
    path = "/home/server1/test_crl.crl"
    password = "a10"
    username = "server1"
    overwrite = 1
    use_mgmt_port = 1   
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `host` (String) Remote site (IP or domain name)
- `name` (String) Local file name
- `path` (String) Remote path
- `protocol` (String) Transfer protocol

### Optional

- `overwrite` (Number) Overwrite existing file
- `password` (String) Password for the remote site
- `use_mgmt_port` (Number) Use management port as source port
- `username` (String) Username for the remote site

### Read-Only

- `id` (String) The ID of this resource.


