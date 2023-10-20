provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_l7session" "test_thunder_slb_l7session" {
  sampling_enable {
    counters1 = "all"
  }
}

