provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_http_proxy" "http_proxy1" {
  sampling_enable {
    counters1 = "all"
    counters2 = "req_sz_8k"
  }
}
