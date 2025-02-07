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

resource "thunder_file_class_list" "IPV4_1" {
  name          = "IPV4_1"
  protocol      = "http"
  host          = "192.168.92.200"
  path          = "/class-list-files/IPV4_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_IPV4_1" {
  provider      = thunder.L3V_A
  name          = "IPV4_1"
  protocol      = "https"
  host          = "192.168.92.200"
  path          = "/class-list-files/IPV4_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "IPV6_1" {
  name          = "IPV6_1"
  protocol      = "http"
  host          = "192.168.92.200"
  path          = "/class-list-files/IPV6_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_IPV6_1" {
  provider      = thunder.L3V_A
  name          = "IPV6_1"
  protocol      = "https"
  host          = "192.168.92.200"
  path          = "/class-list-files/IPV6_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "DNS_1" {
  name          = "DNS_1"
  protocol      = "http"
  host          = "192.168.92.200"
  path          = "/class-list-files/DNS_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_DNS_1" {
  provider      = thunder.L3V_A
  name          = "DNS_1"
  protocol      = "https"
  host          = "192.168.92.200"
  path          = "/class-list-files/DNS_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "AC_1" {
  name          = "AC_1"
  protocol      = "http"
  host          = "192.168.92.200"
  path          = "/class-list-files/AC_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_AC_1" {
  provider      = thunder.L3V_A
  name          = "AC_1"
  protocol      = "https"
  host          = "192.168.92.200"
  path          = "/class-list-files/AC_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "STR_1" {
  name          = "STR_1"
  protocol      = "http"
  host          = "192.168.92.200"
  path          = "/class-list-files/STR_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_STR_1" {
  provider      = thunder.L3V_A
  name          = "STR_1"
  protocol      = "https"
  host          = "192.168.92.200"
  path          = "/class-list-files/STR_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "STR_CI_1" {
  name          = "STR_CI_1"
  protocol      = "http"
  host          = "192.168.92.200"
  path          = "/class-list-files/STR_CI_1.txt"
  use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_STR_CI_1" {
  provider      = thunder.L3V_A
  name          = "STR_CI_1"
  protocol      = "https"
  host          = "192.168.92.200"
  path          = "/class-list-files/STR_CI_1.txt"
  use_mgmt_port = 1
}
