provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_rc_cache_global" "test_thunder_slb_rc_cache_global" {
  sampling_enable {
    counters1 = "all"
  }
}

