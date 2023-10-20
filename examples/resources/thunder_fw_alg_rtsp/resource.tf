provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_alg_rtsp" "test_thunder_fw_alg_rtsp" {
  default_port_disable = "default-port-disable"
  sampling_enable {
    counters1 = "all"
  }
}