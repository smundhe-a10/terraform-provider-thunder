provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_switch" "test_thunder_slb_switch" {
  sampling_enable {
    counters1 = "all"
  }
}

