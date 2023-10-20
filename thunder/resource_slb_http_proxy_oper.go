package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHttpProxyOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_http_proxy_oper`: Operational Status for the object http-proxy\n\n__PLACEHOLDER__",

		ReadContext: resourceSlbHttpProxyOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_proxy_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"curr_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connect_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_enter_ssli": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_succ": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cache_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"noproxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"notuple": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parsereq_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"svrsel_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwdreqdata_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_retran": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_ofo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_resel": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"svr_prem_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"new_svrconn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compression_before": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compression_after": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_over_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_rate_over_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"close_on_ddos": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"decompression_before": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"decompression_after": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwdreq_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwdreq_fail_buff": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwdreq_fail_rport": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwdreq_fail_route": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwdreq_fail_persist": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwdreq_fail_server": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwdreq_fail_tuple": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"l4_switching": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cookie_switching": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_switching": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"url_switching": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"host_switching": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lb_switching": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"l4_switching_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cookie_switching_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_switching_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"url_switching_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"host_switching_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lb_switching_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"l4_switching_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cookie_switching_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_switching_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"url_switching_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"host_switching_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lb_switching_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"non_http_bypass": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_rst_request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_rst_connecting": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_rst_connected": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_rst_response": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_rst_request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_rst_connecting": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_rst_connected": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_rst_response": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_req_unexp_flag": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connecting_fin": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connecting_fin_retrans": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connecting_fin_ofo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connecting_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connecting_rst_retrans": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connecting_rst_ofo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connecting_ack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkts_ofo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkts_retrans": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stale_sess": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_resel_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"large_cookie": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"large_cookie_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"huge_cookie": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"huge_cookie_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parse_cookie_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parse_setcookie_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"asm_cookie_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"asm_cookie_header_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"asm_setcookie_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"asm_setcookie_header_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"too_many_headers": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_name_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"wrong_resp_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_insert": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_delete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"insert_client_ip": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"insert_client_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"skip_insert_client_ip": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"skip_insert_client_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"negative_req_remain": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"negative_resp_remain": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"retry_503": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_retry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_lb_reselect": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_lb_reselect_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_http10": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_http11": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_http2": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_http2": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_get": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_head": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_put": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_post": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_trace": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_track": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_options": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_connect": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_delete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_http10": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_http11": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_1xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_100": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_101": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_102": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_2xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_200": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_201": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_202": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_203": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_204": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_205": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_206": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_207": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_3xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_300": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_301": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_302": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_303": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_304": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_305": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_306": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_307": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_4xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_400": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_401": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_402": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_403": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_404": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_405": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_406": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_407": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_408": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_409": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_410": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_411": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_412": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_413": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_414": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_415": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_416": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_417": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_418": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_422": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_423": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_424": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_425": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_426": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_449": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_450": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_5xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_500": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_501": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_502": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_503": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_504": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_504_ax": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_505": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_506": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_507": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_508": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_509": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_510": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_6xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_10u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_20u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_50u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_100u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_200u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_500u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_1m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_2m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_5m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_10m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_20m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_50m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_100m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_200m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_500m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_1s": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_2s": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_5s": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_over_5s": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_8k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_16k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_32k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_64k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_256k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_gt_256k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_8k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_16k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_32k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_64k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_256k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_gt_256k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_sz_512": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_sz_gt_4k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkts_retrans_ack_finwait": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkts_retrans_fin": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkts_retrans_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkts_retrans_push": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pconn_connecting": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pconn_connected": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pconn_connecting_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_del_accept_enc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_resp_already_compressed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_content_type_excluded": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_no_content_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_resp_lt_min": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_resp_no_cl_or_ce": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_ratio_too_high": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_content_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_content_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_chunk": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_http10_keepalive": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_bad": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ws_handshake_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ws_handshake_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ws_client_packets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ws_server_packets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_timeout_retry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_timeout_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_get": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_post": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_non_doh_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_non_doh_req_get": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_non_doh_req_post": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_tc_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_udp_dns_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_udp_dns_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_tcp_dns_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_tcp_dns_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_malloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_udp_retry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_udp_retry_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_tcp_retry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_tcp_retry_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_snat_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_path_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_get_dns_arg_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_get_base64_decode_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_post_content_type_mismatch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_post_payload_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_post_payload_extract_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_non_doh_method": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_tcp_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_udp_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_query_time_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_a": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_aaaa": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_ns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_cname": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_any": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_srv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_mx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_soa": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_others": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_setup_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_header_alloc_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_que_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_udp_frags": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_tcp_frags": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_serv_sel_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_retry_w_tcp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_get_uri_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_post_payload_too_large": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_malformed_query": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_err_format": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_err_server": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_err_name": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_err_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_refuse": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_yxdomain": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_yxrrset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_nxrrset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_notauth": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_notzone": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_other": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compression_before_br": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compression_after_br": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compression_before_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compression_after_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"decompression_before_br": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"decompression_after_br": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"decompression_before_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"decompression_after_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_rsp_br": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_rsp_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"h2up_content_length_alias": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"malformed_h2up_header_value": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"malformed_h2up_scheme_value": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"h2up_with_transfer_encoding": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"multiple_content_length": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"multiple_transfer_encoding": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"transfer_encoding_and_content_length": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"get_and_payload": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"h2up_with_host_and_auth": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_filter_rule_hit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"current_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_create": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_free": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"end_stream_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"end_stream_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http1_client_idle_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http2_client_idle_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"debug_fields": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHttpProxyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttpProxyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttpProxyOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		items := setObjectSlbHttpProxyOperOper(res)
		d.SetId(obj.GetId())
		d.Set("oper", items)

		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHttpProxyOperOper(res edpt.SlbHttpProxyy) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["http_proxy_cpu_list"] = setSliceSlbHttpProxyOperOperHttpProxyCpuList(res.DataSlbHttpProxy.Oper.HttpProxyCpuList)
	in["cpu_count"] = res.DataSlbHttpProxy.Oper.CpuCount
	in["debug_fields"] = res.DataSlbHttpProxy.Oper.Debug_fields
	result = append(result, in)
	return result
}

func setSliceSlbHttpProxyOperOperHttpProxyCpuList(d []edpt.SlbHttpProxyOperOperHttpProxyCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["curr_proxy"] = item.Curr_proxy
		in["total_proxy"] = item.Total_proxy
		in["req"] = item.Req
		in["connect_req"] = item.Connect_req
		in["req_enter_ssli"] = item.Req_enter_ssli
		in["req_succ"] = item.Req_succ
		in["cache_rsp"] = item.Cache_rsp
		in["noproxy"] = item.Noproxy
		in["notuple"] = item.Notuple
		in["parsereq_fail"] = item.Parsereq_fail
		in["svrsel_fail"] = item.Svrsel_fail
		in["fwdreqdata_fail"] = item.Fwdreqdata_fail
		in["req_retran"] = item.Req_retran
		in["req_ofo"] = item.Req_ofo
		in["server_resel"] = item.Server_resel
		in["svr_prem_close"] = item.Svr_prem_close
		in["new_svrconn"] = item.New_svrconn
		in["snat_fail"] = item.Snat_fail
		in["compression_before"] = item.Compression_before
		in["compression_after"] = item.Compression_after
		in["req_over_limit"] = item.Req_over_limit
		in["req_rate_over_limit"] = item.Req_rate_over_limit
		in["close_on_ddos"] = item.Close_on_ddos
		in["decompression_before"] = item.Decompression_before
		in["decompression_after"] = item.Decompression_after
		in["client_rst"] = item.Client_rst
		in["server_rst"] = item.Server_rst
		in["fwdreq_fail"] = item.Fwdreq_fail
		in["fwdreq_fail_buff"] = item.Fwdreq_fail_buff
		in["fwdreq_fail_rport"] = item.Fwdreq_fail_rport
		in["fwdreq_fail_route"] = item.Fwdreq_fail_route
		in["fwdreq_fail_persist"] = item.Fwdreq_fail_persist
		in["fwdreq_fail_server"] = item.Fwdreq_fail_server
		in["fwdreq_fail_tuple"] = item.Fwdreq_fail_tuple
		in["l4_switching"] = item.L4_switching
		in["cookie_switching"] = item.Cookie_switching
		in["aflex_switching"] = item.Aflex_switching
		in["url_switching"] = item.Url_switching
		in["host_switching"] = item.Host_switching
		in["lb_switching"] = item.Lb_switching
		in["l4_switching_ok"] = item.L4_switching_ok
		in["cookie_switching_ok"] = item.Cookie_switching_ok
		in["aflex_switching_ok"] = item.Aflex_switching_ok
		in["url_switching_ok"] = item.Url_switching_ok
		in["host_switching_ok"] = item.Host_switching_ok
		in["lb_switching_ok"] = item.Lb_switching_ok
		in["l4_switching_enqueue"] = item.L4_switching_enqueue
		in["cookie_switching_enqueue"] = item.Cookie_switching_enqueue
		in["aflex_switching_enqueue"] = item.Aflex_switching_enqueue
		in["url_switching_enqueue"] = item.Url_switching_enqueue
		in["host_switching_enqueue"] = item.Host_switching_enqueue
		in["lb_switching_enqueue"] = item.Lb_switching_enqueue
		in["non_http_bypass"] = item.Non_http_bypass
		in["client_rst_request"] = item.Client_rst_request
		in["client_rst_connecting"] = item.Client_rst_connecting
		in["client_rst_connected"] = item.Client_rst_connected
		in["client_rst_response"] = item.Client_rst_response
		in["server_rst_request"] = item.Server_rst_request
		in["server_rst_connecting"] = item.Server_rst_connecting
		in["server_rst_connected"] = item.Server_rst_connected
		in["server_rst_response"] = item.Server_rst_response
		in["client_req_unexp_flag"] = item.Client_req_unexp_flag
		in["connecting_fin"] = item.Connecting_fin
		in["connecting_fin_retrans"] = item.Connecting_fin_retrans
		in["connecting_fin_ofo"] = item.Connecting_fin_ofo
		in["connecting_rst"] = item.Connecting_rst
		in["connecting_rst_retrans"] = item.Connecting_rst_retrans
		in["connecting_rst_ofo"] = item.Connecting_rst_ofo
		in["connecting_ack"] = item.Connecting_ack
		in["pkts_ofo"] = item.Pkts_ofo
		in["pkts_retrans"] = item.Pkts_retrans
		in["stale_sess"] = item.Stale_sess
		in["server_resel_failed"] = item.Server_resel_failed
		in["large_cookie"] = item.Large_cookie
		in["large_cookie_header"] = item.Large_cookie_header
		in["huge_cookie"] = item.Huge_cookie
		in["huge_cookie_header"] = item.Huge_cookie_header
		in["parse_cookie_fail"] = item.Parse_cookie_fail
		in["parse_setcookie_fail"] = item.Parse_setcookie_fail
		in["asm_cookie_fail"] = item.Asm_cookie_fail
		in["asm_cookie_header_fail"] = item.Asm_cookie_header_fail
		in["asm_setcookie_fail"] = item.Asm_setcookie_fail
		in["asm_setcookie_header_fail"] = item.Asm_setcookie_header_fail
		in["invalid_header"] = item.Invalid_header
		in["too_many_headers"] = item.Too_many_headers
		in["line_too_long"] = item.Line_too_long
		in["header_name_too_long"] = item.Header_name_too_long
		in["wrong_resp_header"] = item.Wrong_resp_header
		in["header_insert"] = item.Header_insert
		in["header_delete"] = item.Header_delete
		in["insert_client_ip"] = item.Insert_client_ip
		in["insert_client_port"] = item.Insert_client_port
		in["skip_insert_client_ip"] = item.Skip_insert_client_ip
		in["skip_insert_client_port"] = item.Skip_insert_client_port
		in["negative_req_remain"] = item.Negative_req_remain
		in["negative_resp_remain"] = item.Negative_resp_remain
		in["retry_503"] = item.Retry_503
		in["aflex_retry"] = item.Aflex_retry
		in["aflex_lb_reselect"] = item.Aflex_lb_reselect
		in["aflex_lb_reselect_ok"] = item.Aflex_lb_reselect_ok
		in["req_http10"] = item.Req_http10
		in["req_http11"] = item.Req_http11
		in["req_http2"] = item.Req_http2
		in["response_http2"] = item.Response_http2
		in["req_get"] = item.Req_get
		in["req_head"] = item.Req_head
		in["req_put"] = item.Req_put
		in["req_post"] = item.Req_post
		in["req_trace"] = item.Req_trace
		in["req_track"] = item.Req_track
		in["req_options"] = item.Req_options
		in["req_connect"] = item.Req_connect
		in["req_delete"] = item.Req_delete
		in["req_unknown"] = item.Req_unknown
		in["response_http10"] = item.Response_http10
		in["response_http11"] = item.Response_http11
		in["response_1xx"] = item.Response_1xx
		in["response_100"] = item.Response_100
		in["response_101"] = item.Response_101
		in["response_102"] = item.Response_102
		in["response_2xx"] = item.Response_2xx
		in["response_200"] = item.Response_200
		in["response_201"] = item.Response_201
		in["response_202"] = item.Response_202
		in["response_203"] = item.Response_203
		in["response_204"] = item.Response_204
		in["response_205"] = item.Response_205
		in["response_206"] = item.Response_206
		in["response_207"] = item.Response_207
		in["response_3xx"] = item.Response_3xx
		in["response_300"] = item.Response_300
		in["response_301"] = item.Response_301
		in["response_302"] = item.Response_302
		in["response_303"] = item.Response_303
		in["response_304"] = item.Response_304
		in["response_305"] = item.Response_305
		in["response_306"] = item.Response_306
		in["response_307"] = item.Response_307
		in["response_4xx"] = item.Response_4xx
		in["response_400"] = item.Response_400
		in["response_401"] = item.Response_401
		in["response_402"] = item.Response_402
		in["response_403"] = item.Response_403
		in["response_404"] = item.Response_404
		in["response_405"] = item.Response_405
		in["response_406"] = item.Response_406
		in["response_407"] = item.Response_407
		in["response_408"] = item.Response_408
		in["response_409"] = item.Response_409
		in["response_410"] = item.Response_410
		in["response_411"] = item.Response_411
		in["response_412"] = item.Response_412
		in["response_413"] = item.Response_413
		in["response_414"] = item.Response_414
		in["response_415"] = item.Response_415
		in["response_416"] = item.Response_416
		in["response_417"] = item.Response_417
		in["response_418"] = item.Response_418
		in["response_422"] = item.Response_422
		in["response_423"] = item.Response_423
		in["response_424"] = item.Response_424
		in["response_425"] = item.Response_425
		in["response_426"] = item.Response_426
		in["response_449"] = item.Response_449
		in["response_450"] = item.Response_450
		in["response_5xx"] = item.Response_5xx
		in["response_500"] = item.Response_500
		in["response_501"] = item.Response_501
		in["response_502"] = item.Response_502
		in["response_503"] = item.Response_503
		in["response_504"] = item.Response_504
		in["response_504_ax"] = item.Response_504_ax
		in["response_505"] = item.Response_505
		in["response_506"] = item.Response_506
		in["response_507"] = item.Response_507
		in["response_508"] = item.Response_508
		in["response_509"] = item.Response_509
		in["response_510"] = item.Response_510
		in["response_6xx"] = item.Response_6xx
		in["response_unknown"] = item.Response_unknown
		in["req_10u"] = item.Req_10u
		in["req_20u"] = item.Req_20u
		in["req_50u"] = item.Req_50u
		in["req_100u"] = item.Req_100u
		in["req_200u"] = item.Req_200u
		in["req_500u"] = item.Req_500u
		in["req_1m"] = item.Req_1m
		in["req_2m"] = item.Req_2m
		in["req_5m"] = item.Req_5m
		in["req_10m"] = item.Req_10m
		in["req_20m"] = item.Req_20m
		in["req_50m"] = item.Req_50m
		in["req_100m"] = item.Req_100m
		in["req_200m"] = item.Req_200m
		in["req_500m"] = item.Req_500m
		in["req_1s"] = item.Req_1s
		in["req_2s"] = item.Req_2s
		in["req_5s"] = item.Req_5s
		in["req_over_5s"] = item.Req_over_5s
		in["req_sz_1k"] = item.Req_sz_1k
		in["req_sz_2k"] = item.Req_sz_2k
		in["req_sz_4k"] = item.Req_sz_4k
		in["req_sz_8k"] = item.Req_sz_8k
		in["req_sz_16k"] = item.Req_sz_16k
		in["req_sz_32k"] = item.Req_sz_32k
		in["req_sz_64k"] = item.Req_sz_64k
		in["req_sz_256k"] = item.Req_sz_256k
		in["req_sz_gt_256k"] = item.Req_sz_gt_256k
		in["rsp_sz_1k"] = item.Rsp_sz_1k
		in["rsp_sz_2k"] = item.Rsp_sz_2k
		in["rsp_sz_4k"] = item.Rsp_sz_4k
		in["rsp_sz_8k"] = item.Rsp_sz_8k
		in["rsp_sz_16k"] = item.Rsp_sz_16k
		in["rsp_sz_32k"] = item.Rsp_sz_32k
		in["rsp_sz_64k"] = item.Rsp_sz_64k
		in["rsp_sz_256k"] = item.Rsp_sz_256k
		in["rsp_sz_gt_256k"] = item.Rsp_sz_gt_256k
		in["chunk_sz_512"] = item.Chunk_sz_512
		in["chunk_sz_1k"] = item.Chunk_sz_1k
		in["chunk_sz_2k"] = item.Chunk_sz_2k
		in["chunk_sz_4k"] = item.Chunk_sz_4k
		in["chunk_sz_gt_4k"] = item.Chunk_sz_gt_4k
		in["pkts_retrans_ack_finwait"] = item.Pkts_retrans_ack_finwait
		in["pkts_retrans_fin"] = item.Pkts_retrans_fin
		in["pkts_retrans_rst"] = item.Pkts_retrans_rst
		in["pkts_retrans_push"] = item.Pkts_retrans_push
		in["pconn_connecting"] = item.Pconn_connecting
		in["pconn_connected"] = item.Pconn_connected
		in["pconn_connecting_failed"] = item.Pconn_connecting_failed
		in["compress_rsp"] = item.Compress_rsp
		in["compress_del_accept_enc"] = item.Compress_del_accept_enc
		in["compress_resp_already_compressed"] = item.Compress_resp_already_compressed
		in["compress_content_type_excluded"] = item.Compress_content_type_excluded
		in["compress_no_content_type"] = item.Compress_no_content_type
		in["compress_resp_lt_min"] = item.Compress_resp_lt_min
		in["compress_resp_no_cl_or_ce"] = item.Compress_resp_no_cl_or_ce
		in["compress_ratio_too_high"] = item.Compress_ratio_too_high
		in["rsp_content_len"] = item.Rsp_content_len
		in["req_content_len"] = item.Req_content_len
		in["rsp_chunk"] = item.Rsp_chunk
		in["req_http10_keepalive"] = item.Req_http10_keepalive
		in["chunk_bad"] = item.Chunk_bad
		in["ws_handshake_req"] = item.Ws_handshake_req
		in["ws_handshake_resp"] = item.Ws_handshake_resp
		in["ws_client_packets"] = item.Ws_client_packets
		in["ws_server_packets"] = item.Ws_server_packets
		in["req_timeout_retry"] = item.Req_timeout_retry
		in["req_timeout_close"] = item.Req_timeout_close
		in["doh_req"] = item.Doh_req
		in["doh_req_get"] = item.Doh_req_get
		in["doh_req_post"] = item.Doh_req_post
		in["doh_non_doh_req"] = item.Doh_non_doh_req
		in["doh_non_doh_req_get"] = item.Doh_non_doh_req_get
		in["doh_non_doh_req_post"] = item.Doh_non_doh_req_post
		in["doh_resp"] = item.Doh_resp
		in["doh_tc_resp"] = item.Doh_tc_resp
		in["doh_udp_dns_req"] = item.Doh_udp_dns_req
		in["doh_udp_dns_resp"] = item.Doh_udp_dns_resp
		in["doh_tcp_dns_req"] = item.Doh_tcp_dns_req
		in["doh_tcp_dns_resp"] = item.Doh_tcp_dns_resp
		in["doh_req_send_failed"] = item.Doh_req_send_failed
		in["doh_resp_send_failed"] = item.Doh_resp_send_failed
		in["doh_malloc_fail"] = item.Doh_malloc_fail
		in["doh_req_udp_retry"] = item.Doh_req_udp_retry
		in["doh_req_udp_retry_fail"] = item.Doh_req_udp_retry_fail
		in["doh_req_tcp_retry"] = item.Doh_req_tcp_retry
		in["doh_req_tcp_retry_fail"] = item.Doh_req_tcp_retry_fail
		in["doh_snat_failed"] = item.Doh_snat_failed
		in["doh_path_not_found"] = item.Doh_path_not_found
		in["doh_get_dns_arg_failed"] = item.Doh_get_dns_arg_failed
		in["doh_get_base64_decode_failed"] = item.Doh_get_base64_decode_failed
		in["doh_post_content_type_mismatch"] = item.Doh_post_content_type_mismatch
		in["doh_post_payload_not_found"] = item.Doh_post_payload_not_found
		in["doh_post_payload_extract_failed"] = item.Doh_post_payload_extract_failed
		in["doh_non_doh_method"] = item.Doh_non_doh_method
		in["doh_tcp_send_failed"] = item.Doh_tcp_send_failed
		in["doh_udp_send_failed"] = item.Doh_udp_send_failed
		in["doh_query_time_out"] = item.Doh_query_time_out
		in["doh_dns_query_type_a"] = item.Doh_dns_query_type_a
		in["doh_dns_query_type_aaaa"] = item.Doh_dns_query_type_aaaa
		in["doh_dns_query_type_ns"] = item.Doh_dns_query_type_ns
		in["doh_dns_query_type_cname"] = item.Doh_dns_query_type_cname
		in["doh_dns_query_type_any"] = item.Doh_dns_query_type_any
		in["doh_dns_query_type_srv"] = item.Doh_dns_query_type_srv
		in["doh_dns_query_type_mx"] = item.Doh_dns_query_type_mx
		in["doh_dns_query_type_soa"] = item.Doh_dns_query_type_soa
		in["doh_dns_query_type_others"] = item.Doh_dns_query_type_others
		in["doh_resp_setup_failed"] = item.Doh_resp_setup_failed
		in["doh_resp_header_alloc_failed"] = item.Doh_resp_header_alloc_failed
		in["doh_resp_que_failed"] = item.Doh_resp_que_failed
		in["doh_resp_udp_frags"] = item.Doh_resp_udp_frags
		in["doh_resp_tcp_frags"] = item.Doh_resp_tcp_frags
		in["doh_serv_sel_failed"] = item.Doh_serv_sel_failed
		in["doh_retry_w_tcp"] = item.Doh_retry_w_tcp
		in["doh_get_uri_too_long"] = item.Doh_get_uri_too_long
		in["doh_post_payload_too_large"] = item.Doh_post_payload_too_large
		in["doh_dns_malformed_query"] = item.Doh_dns_malformed_query
		in["doh_dns_resp_rcode_err_format"] = item.Doh_dns_resp_rcode_err_format
		in["doh_dns_resp_rcode_err_server"] = item.Doh_dns_resp_rcode_err_server
		in["doh_dns_resp_rcode_err_name"] = item.Doh_dns_resp_rcode_err_name
		in["doh_dns_resp_rcode_err_type"] = item.Doh_dns_resp_rcode_err_type
		in["doh_dns_resp_rcode_refuse"] = item.Doh_dns_resp_rcode_refuse
		in["doh_dns_resp_rcode_yxdomain"] = item.Doh_dns_resp_rcode_yxdomain
		in["doh_dns_resp_rcode_yxrrset"] = item.Doh_dns_resp_rcode_yxrrset
		in["doh_dns_resp_rcode_nxrrset"] = item.Doh_dns_resp_rcode_nxrrset
		in["doh_dns_resp_rcode_notauth"] = item.Doh_dns_resp_rcode_notauth
		in["doh_dns_resp_rcode_notzone"] = item.Doh_dns_resp_rcode_notzone
		in["doh_dns_resp_rcode_other"] = item.Doh_dns_resp_rcode_other
		in["compression_before_br"] = item.Compression_before_br
		in["compression_after_br"] = item.Compression_after_br
		in["compression_before_total"] = item.Compression_before_total
		in["compression_after_total"] = item.Compression_after_total
		in["decompression_before_br"] = item.Decompression_before_br
		in["decompression_after_br"] = item.Decompression_after_br
		in["decompression_before_total"] = item.Decompression_before_total
		in["decompression_after_total"] = item.Decompression_after_total
		in["compress_rsp_br"] = item.Compress_rsp_br
		in["compress_rsp_total"] = item.Compress_rsp_total
		in["h2up_content_length_alias"] = item.H2up_content_length_alias
		in["malformed_h2up_header_value"] = item.Malformed_h2up_header_value
		in["malformed_h2up_scheme_value"] = item.Malformed_h2up_scheme_value
		in["h2up_with_transfer_encoding"] = item.H2up_with_transfer_encoding
		in["multiple_content_length"] = item.Multiple_content_length
		in["multiple_transfer_encoding"] = item.Multiple_transfer_encoding
		in["transfer_encoding_and_content_length"] = item.Transfer_encoding_and_content_length
		in["get_and_payload"] = item.Get_and_payload
		in["h2up_with_host_and_auth"] = item.H2up_with_host_and_auth
		in["header_filter_rule_hit"] = item.Header_filter_rule_hit
		in["current_stream"] = item.Current_stream
		in["stream_create"] = item.Stream_create
		in["stream_free"] = item.Stream_free
		in["end_stream_rcvd"] = item.End_stream_rcvd
		in["end_stream_sent"] = item.End_stream_sent
		in["http1_client_idle_timeout"] = item.Http1_client_idle_timeout
		in["http2_client_idle_timeout"] = item.Http2_client_idle_timeout
		result = append(result, in)
	}
	return result
}

func getObjectSlbHttpProxyOperOper(d []interface{}) edpt.SlbHttpProxyOperOper {
	count := len(d)
	var ret edpt.SlbHttpProxyOperOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpProxyCpuList = getSliceSlbHttpProxyOperOperHttpProxyCpuList(in["http_proxy_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
		ret.Debug_fields = in["debug_fields"].(int)
	}
	return ret
}

func getSliceSlbHttpProxyOperOperHttpProxyCpuList(d []interface{}) []edpt.SlbHttpProxyOperOperHttpProxyCpuList {
	count := len(d)
	ret := make([]edpt.SlbHttpProxyOperOperHttpProxyCpuList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHttpProxyOperOperHttpProxyCpuList
		oi.Curr_proxy = in["curr_proxy"].(int)
		oi.Total_proxy = in["total_proxy"].(int)
		oi.Req = in["req"].(int)
		oi.Connect_req = in["connect_req"].(int)
		oi.Req_enter_ssli = in["req_enter_ssli"].(int)
		oi.Req_succ = in["req_succ"].(int)
		oi.Cache_rsp = in["cache_rsp"].(int)
		oi.Noproxy = in["noproxy"].(int)
		oi.Notuple = in["notuple"].(int)
		oi.Parsereq_fail = in["parsereq_fail"].(int)
		oi.Svrsel_fail = in["svrsel_fail"].(int)
		oi.Fwdreqdata_fail = in["fwdreqdata_fail"].(int)
		oi.Req_retran = in["req_retran"].(int)
		oi.Req_ofo = in["req_ofo"].(int)
		oi.Server_resel = in["server_resel"].(int)
		oi.Svr_prem_close = in["svr_prem_close"].(int)
		oi.New_svrconn = in["new_svrconn"].(int)
		oi.Snat_fail = in["snat_fail"].(int)
		oi.Compression_before = in["compression_before"].(int)
		oi.Compression_after = in["compression_after"].(int)
		oi.Req_over_limit = in["req_over_limit"].(int)
		oi.Req_rate_over_limit = in["req_rate_over_limit"].(int)
		oi.Close_on_ddos = in["close_on_ddos"].(int)
		oi.Decompression_before = in["decompression_before"].(int)
		oi.Decompression_after = in["decompression_after"].(int)
		oi.Client_rst = in["client_rst"].(int)
		oi.Server_rst = in["server_rst"].(int)
		oi.Fwdreq_fail = in["fwdreq_fail"].(int)
		oi.Fwdreq_fail_buff = in["fwdreq_fail_buff"].(int)
		oi.Fwdreq_fail_rport = in["fwdreq_fail_rport"].(int)
		oi.Fwdreq_fail_route = in["fwdreq_fail_route"].(int)
		oi.Fwdreq_fail_persist = in["fwdreq_fail_persist"].(int)
		oi.Fwdreq_fail_server = in["fwdreq_fail_server"].(int)
		oi.Fwdreq_fail_tuple = in["fwdreq_fail_tuple"].(int)
		oi.L4_switching = in["l4_switching"].(int)
		oi.Cookie_switching = in["cookie_switching"].(int)
		oi.Aflex_switching = in["aflex_switching"].(int)
		oi.Url_switching = in["url_switching"].(int)
		oi.Host_switching = in["host_switching"].(int)
		oi.Lb_switching = in["lb_switching"].(int)
		oi.L4_switching_ok = in["l4_switching_ok"].(int)
		oi.Cookie_switching_ok = in["cookie_switching_ok"].(int)
		oi.Aflex_switching_ok = in["aflex_switching_ok"].(int)
		oi.Url_switching_ok = in["url_switching_ok"].(int)
		oi.Host_switching_ok = in["host_switching_ok"].(int)
		oi.Lb_switching_ok = in["lb_switching_ok"].(int)
		oi.L4_switching_enqueue = in["l4_switching_enqueue"].(int)
		oi.Cookie_switching_enqueue = in["cookie_switching_enqueue"].(int)
		oi.Aflex_switching_enqueue = in["aflex_switching_enqueue"].(int)
		oi.Url_switching_enqueue = in["url_switching_enqueue"].(int)
		oi.Host_switching_enqueue = in["host_switching_enqueue"].(int)
		oi.Lb_switching_enqueue = in["lb_switching_enqueue"].(int)
		oi.Non_http_bypass = in["non_http_bypass"].(int)
		oi.Client_rst_request = in["client_rst_request"].(int)
		oi.Client_rst_connecting = in["client_rst_connecting"].(int)
		oi.Client_rst_connected = in["client_rst_connected"].(int)
		oi.Client_rst_response = in["client_rst_response"].(int)
		oi.Server_rst_request = in["server_rst_request"].(int)
		oi.Server_rst_connecting = in["server_rst_connecting"].(int)
		oi.Server_rst_connected = in["server_rst_connected"].(int)
		oi.Server_rst_response = in["server_rst_response"].(int)
		oi.Client_req_unexp_flag = in["client_req_unexp_flag"].(int)
		oi.Connecting_fin = in["connecting_fin"].(int)
		oi.Connecting_fin_retrans = in["connecting_fin_retrans"].(int)
		oi.Connecting_fin_ofo = in["connecting_fin_ofo"].(int)
		oi.Connecting_rst = in["connecting_rst"].(int)
		oi.Connecting_rst_retrans = in["connecting_rst_retrans"].(int)
		oi.Connecting_rst_ofo = in["connecting_rst_ofo"].(int)
		oi.Connecting_ack = in["connecting_ack"].(int)
		oi.Pkts_ofo = in["pkts_ofo"].(int)
		oi.Pkts_retrans = in["pkts_retrans"].(int)
		oi.Stale_sess = in["stale_sess"].(int)
		oi.Server_resel_failed = in["server_resel_failed"].(int)
		oi.Large_cookie = in["large_cookie"].(int)
		oi.Large_cookie_header = in["large_cookie_header"].(int)
		oi.Huge_cookie = in["huge_cookie"].(int)
		oi.Huge_cookie_header = in["huge_cookie_header"].(int)
		oi.Parse_cookie_fail = in["parse_cookie_fail"].(int)
		oi.Parse_setcookie_fail = in["parse_setcookie_fail"].(int)
		oi.Asm_cookie_fail = in["asm_cookie_fail"].(int)
		oi.Asm_cookie_header_fail = in["asm_cookie_header_fail"].(int)
		oi.Asm_setcookie_fail = in["asm_setcookie_fail"].(int)
		oi.Asm_setcookie_header_fail = in["asm_setcookie_header_fail"].(int)
		oi.Invalid_header = in["invalid_header"].(int)
		oi.Too_many_headers = in["too_many_headers"].(int)
		oi.Line_too_long = in["line_too_long"].(int)
		oi.Header_name_too_long = in["header_name_too_long"].(int)
		oi.Wrong_resp_header = in["wrong_resp_header"].(int)
		oi.Header_insert = in["header_insert"].(int)
		oi.Header_delete = in["header_delete"].(int)
		oi.Insert_client_ip = in["insert_client_ip"].(int)
		oi.Insert_client_port = in["insert_client_port"].(int)
		oi.Skip_insert_client_ip = in["skip_insert_client_ip"].(int)
		oi.Skip_insert_client_port = in["skip_insert_client_port"].(int)
		oi.Negative_req_remain = in["negative_req_remain"].(int)
		oi.Negative_resp_remain = in["negative_resp_remain"].(int)
		oi.Retry_503 = in["retry_503"].(int)
		oi.Aflex_retry = in["aflex_retry"].(int)
		oi.Aflex_lb_reselect = in["aflex_lb_reselect"].(int)
		oi.Aflex_lb_reselect_ok = in["aflex_lb_reselect_ok"].(int)
		oi.Req_http10 = in["req_http10"].(int)
		oi.Req_http11 = in["req_http11"].(int)
		oi.Req_http2 = in["req_http2"].(int)
		oi.Response_http2 = in["response_http2"].(int)
		oi.Req_get = in["req_get"].(int)
		oi.Req_head = in["req_head"].(int)
		oi.Req_put = in["req_put"].(int)
		oi.Req_post = in["req_post"].(int)
		oi.Req_trace = in["req_trace"].(int)
		oi.Req_track = in["req_track"].(int)
		oi.Req_options = in["req_options"].(int)
		oi.Req_connect = in["req_connect"].(int)
		oi.Req_delete = in["req_delete"].(int)
		oi.Req_unknown = in["req_unknown"].(int)
		oi.Response_http10 = in["response_http10"].(int)
		oi.Response_http11 = in["response_http11"].(int)
		oi.Response_1xx = in["response_1xx"].(int)
		oi.Response_100 = in["response_100"].(int)
		oi.Response_101 = in["response_101"].(int)
		oi.Response_102 = in["response_102"].(int)
		oi.Response_2xx = in["response_2xx"].(int)
		oi.Response_200 = in["response_200"].(int)
		oi.Response_201 = in["response_201"].(int)
		oi.Response_202 = in["response_202"].(int)
		oi.Response_203 = in["response_203"].(int)
		oi.Response_204 = in["response_204"].(int)
		oi.Response_205 = in["response_205"].(int)
		oi.Response_206 = in["response_206"].(int)
		oi.Response_207 = in["response_207"].(int)
		oi.Response_3xx = in["response_3xx"].(int)
		oi.Response_300 = in["response_300"].(int)
		oi.Response_301 = in["response_301"].(int)
		oi.Response_302 = in["response_302"].(int)
		oi.Response_303 = in["response_303"].(int)
		oi.Response_304 = in["response_304"].(int)
		oi.Response_305 = in["response_305"].(int)
		oi.Response_306 = in["response_306"].(int)
		oi.Response_307 = in["response_307"].(int)
		oi.Response_4xx = in["response_4xx"].(int)
		oi.Response_400 = in["response_400"].(int)
		oi.Response_401 = in["response_401"].(int)
		oi.Response_402 = in["response_402"].(int)
		oi.Response_403 = in["response_403"].(int)
		oi.Response_404 = in["response_404"].(int)
		oi.Response_405 = in["response_405"].(int)
		oi.Response_406 = in["response_406"].(int)
		oi.Response_407 = in["response_407"].(int)
		oi.Response_408 = in["response_408"].(int)
		oi.Response_409 = in["response_409"].(int)
		oi.Response_410 = in["response_410"].(int)
		oi.Response_411 = in["response_411"].(int)
		oi.Response_412 = in["response_412"].(int)
		oi.Response_413 = in["response_413"].(int)
		oi.Response_414 = in["response_414"].(int)
		oi.Response_415 = in["response_415"].(int)
		oi.Response_416 = in["response_416"].(int)
		oi.Response_417 = in["response_417"].(int)
		oi.Response_418 = in["response_418"].(int)
		oi.Response_422 = in["response_422"].(int)
		oi.Response_423 = in["response_423"].(int)
		oi.Response_424 = in["response_424"].(int)
		oi.Response_425 = in["response_425"].(int)
		oi.Response_426 = in["response_426"].(int)
		oi.Response_449 = in["response_449"].(int)
		oi.Response_450 = in["response_450"].(int)
		oi.Response_5xx = in["response_5xx"].(int)
		oi.Response_500 = in["response_500"].(int)
		oi.Response_501 = in["response_501"].(int)
		oi.Response_502 = in["response_502"].(int)
		oi.Response_503 = in["response_503"].(int)
		oi.Response_504 = in["response_504"].(int)
		oi.Response_504_ax = in["response_504_ax"].(int)
		oi.Response_505 = in["response_505"].(int)
		oi.Response_506 = in["response_506"].(int)
		oi.Response_507 = in["response_507"].(int)
		oi.Response_508 = in["response_508"].(int)
		oi.Response_509 = in["response_509"].(int)
		oi.Response_510 = in["response_510"].(int)
		oi.Response_6xx = in["response_6xx"].(int)
		oi.Response_unknown = in["response_unknown"].(int)
		oi.Req_10u = in["req_10u"].(int)
		oi.Req_20u = in["req_20u"].(int)
		oi.Req_50u = in["req_50u"].(int)
		oi.Req_100u = in["req_100u"].(int)
		oi.Req_200u = in["req_200u"].(int)
		oi.Req_500u = in["req_500u"].(int)
		oi.Req_1m = in["req_1m"].(int)
		oi.Req_2m = in["req_2m"].(int)
		oi.Req_5m = in["req_5m"].(int)
		oi.Req_10m = in["req_10m"].(int)
		oi.Req_20m = in["req_20m"].(int)
		oi.Req_50m = in["req_50m"].(int)
		oi.Req_100m = in["req_100m"].(int)
		oi.Req_200m = in["req_200m"].(int)
		oi.Req_500m = in["req_500m"].(int)
		oi.Req_1s = in["req_1s"].(int)
		oi.Req_2s = in["req_2s"].(int)
		oi.Req_5s = in["req_5s"].(int)
		oi.Req_over_5s = in["req_over_5s"].(int)
		oi.Req_sz_1k = in["req_sz_1k"].(int)
		oi.Req_sz_2k = in["req_sz_2k"].(int)
		oi.Req_sz_4k = in["req_sz_4k"].(int)
		oi.Req_sz_8k = in["req_sz_8k"].(int)
		oi.Req_sz_16k = in["req_sz_16k"].(int)
		oi.Req_sz_32k = in["req_sz_32k"].(int)
		oi.Req_sz_64k = in["req_sz_64k"].(int)
		oi.Req_sz_256k = in["req_sz_256k"].(int)
		oi.Req_sz_gt_256k = in["req_sz_gt_256k"].(int)
		oi.Rsp_sz_1k = in["rsp_sz_1k"].(int)
		oi.Rsp_sz_2k = in["rsp_sz_2k"].(int)
		oi.Rsp_sz_4k = in["rsp_sz_4k"].(int)
		oi.Rsp_sz_8k = in["rsp_sz_8k"].(int)
		oi.Rsp_sz_16k = in["rsp_sz_16k"].(int)
		oi.Rsp_sz_32k = in["rsp_sz_32k"].(int)
		oi.Rsp_sz_64k = in["rsp_sz_64k"].(int)
		oi.Rsp_sz_256k = in["rsp_sz_256k"].(int)
		oi.Rsp_sz_gt_256k = in["rsp_sz_gt_256k"].(int)
		oi.Chunk_sz_512 = in["chunk_sz_512"].(int)
		oi.Chunk_sz_1k = in["chunk_sz_1k"].(int)
		oi.Chunk_sz_2k = in["chunk_sz_2k"].(int)
		oi.Chunk_sz_4k = in["chunk_sz_4k"].(int)
		oi.Chunk_sz_gt_4k = in["chunk_sz_gt_4k"].(int)
		oi.Pkts_retrans_ack_finwait = in["pkts_retrans_ack_finwait"].(int)
		oi.Pkts_retrans_fin = in["pkts_retrans_fin"].(int)
		oi.Pkts_retrans_rst = in["pkts_retrans_rst"].(int)
		oi.Pkts_retrans_push = in["pkts_retrans_push"].(int)
		oi.Pconn_connecting = in["pconn_connecting"].(int)
		oi.Pconn_connected = in["pconn_connected"].(int)
		oi.Pconn_connecting_failed = in["pconn_connecting_failed"].(int)
		oi.Compress_rsp = in["compress_rsp"].(int)
		oi.Compress_del_accept_enc = in["compress_del_accept_enc"].(int)
		oi.Compress_resp_already_compressed = in["compress_resp_already_compressed"].(int)
		oi.Compress_content_type_excluded = in["compress_content_type_excluded"].(int)
		oi.Compress_no_content_type = in["compress_no_content_type"].(int)
		oi.Compress_resp_lt_min = in["compress_resp_lt_min"].(int)
		oi.Compress_resp_no_cl_or_ce = in["compress_resp_no_cl_or_ce"].(int)
		oi.Compress_ratio_too_high = in["compress_ratio_too_high"].(int)
		oi.Rsp_content_len = in["rsp_content_len"].(int)
		oi.Req_content_len = in["req_content_len"].(int)
		oi.Rsp_chunk = in["rsp_chunk"].(int)
		oi.Req_http10_keepalive = in["req_http10_keepalive"].(int)
		oi.Chunk_bad = in["chunk_bad"].(int)
		oi.Ws_handshake_req = in["ws_handshake_req"].(int)
		oi.Ws_handshake_resp = in["ws_handshake_resp"].(int)
		oi.Ws_client_packets = in["ws_client_packets"].(int)
		oi.Ws_server_packets = in["ws_server_packets"].(int)
		oi.Req_timeout_retry = in["req_timeout_retry"].(int)
		oi.Req_timeout_close = in["req_timeout_close"].(int)
		oi.Doh_req = in["doh_req"].(int)
		oi.Doh_req_get = in["doh_req_get"].(int)
		oi.Doh_req_post = in["doh_req_post"].(int)
		oi.Doh_non_doh_req = in["doh_non_doh_req"].(int)
		oi.Doh_non_doh_req_get = in["doh_non_doh_req_get"].(int)
		oi.Doh_non_doh_req_post = in["doh_non_doh_req_post"].(int)
		oi.Doh_resp = in["doh_resp"].(int)
		oi.Doh_tc_resp = in["doh_tc_resp"].(int)
		oi.Doh_udp_dns_req = in["doh_udp_dns_req"].(int)
		oi.Doh_udp_dns_resp = in["doh_udp_dns_resp"].(int)
		oi.Doh_tcp_dns_req = in["doh_tcp_dns_req"].(int)
		oi.Doh_tcp_dns_resp = in["doh_tcp_dns_resp"].(int)
		oi.Doh_req_send_failed = in["doh_req_send_failed"].(int)
		oi.Doh_resp_send_failed = in["doh_resp_send_failed"].(int)
		oi.Doh_malloc_fail = in["doh_malloc_fail"].(int)
		oi.Doh_req_udp_retry = in["doh_req_udp_retry"].(int)
		oi.Doh_req_udp_retry_fail = in["doh_req_udp_retry_fail"].(int)
		oi.Doh_req_tcp_retry = in["doh_req_tcp_retry"].(int)
		oi.Doh_req_tcp_retry_fail = in["doh_req_tcp_retry_fail"].(int)
		oi.Doh_snat_failed = in["doh_snat_failed"].(int)
		oi.Doh_path_not_found = in["doh_path_not_found"].(int)
		oi.Doh_get_dns_arg_failed = in["doh_get_dns_arg_failed"].(int)
		oi.Doh_get_base64_decode_failed = in["doh_get_base64_decode_failed"].(int)
		oi.Doh_post_content_type_mismatch = in["doh_post_content_type_mismatch"].(int)
		oi.Doh_post_payload_not_found = in["doh_post_payload_not_found"].(int)
		oi.Doh_post_payload_extract_failed = in["doh_post_payload_extract_failed"].(int)
		oi.Doh_non_doh_method = in["doh_non_doh_method"].(int)
		oi.Doh_tcp_send_failed = in["doh_tcp_send_failed"].(int)
		oi.Doh_udp_send_failed = in["doh_udp_send_failed"].(int)
		oi.Doh_query_time_out = in["doh_query_time_out"].(int)
		oi.Doh_dns_query_type_a = in["doh_dns_query_type_a"].(int)
		oi.Doh_dns_query_type_aaaa = in["doh_dns_query_type_aaaa"].(int)
		oi.Doh_dns_query_type_ns = in["doh_dns_query_type_ns"].(int)
		oi.Doh_dns_query_type_cname = in["doh_dns_query_type_cname"].(int)
		oi.Doh_dns_query_type_any = in["doh_dns_query_type_any"].(int)
		oi.Doh_dns_query_type_srv = in["doh_dns_query_type_srv"].(int)
		oi.Doh_dns_query_type_mx = in["doh_dns_query_type_mx"].(int)
		oi.Doh_dns_query_type_soa = in["doh_dns_query_type_soa"].(int)
		oi.Doh_dns_query_type_others = in["doh_dns_query_type_others"].(int)
		oi.Doh_resp_setup_failed = in["doh_resp_setup_failed"].(int)
		oi.Doh_resp_header_alloc_failed = in["doh_resp_header_alloc_failed"].(int)
		oi.Doh_resp_que_failed = in["doh_resp_que_failed"].(int)
		oi.Doh_resp_udp_frags = in["doh_resp_udp_frags"].(int)
		oi.Doh_resp_tcp_frags = in["doh_resp_tcp_frags"].(int)
		oi.Doh_serv_sel_failed = in["doh_serv_sel_failed"].(int)
		oi.Doh_retry_w_tcp = in["doh_retry_w_tcp"].(int)
		oi.Doh_get_uri_too_long = in["doh_get_uri_too_long"].(int)
		oi.Doh_post_payload_too_large = in["doh_post_payload_too_large"].(int)
		oi.Doh_dns_malformed_query = in["doh_dns_malformed_query"].(int)
		oi.Doh_dns_resp_rcode_err_format = in["doh_dns_resp_rcode_err_format"].(int)
		oi.Doh_dns_resp_rcode_err_server = in["doh_dns_resp_rcode_err_server"].(int)
		oi.Doh_dns_resp_rcode_err_name = in["doh_dns_resp_rcode_err_name"].(int)
		oi.Doh_dns_resp_rcode_err_type = in["doh_dns_resp_rcode_err_type"].(int)
		oi.Doh_dns_resp_rcode_refuse = in["doh_dns_resp_rcode_refuse"].(int)
		oi.Doh_dns_resp_rcode_yxdomain = in["doh_dns_resp_rcode_yxdomain"].(int)
		oi.Doh_dns_resp_rcode_yxrrset = in["doh_dns_resp_rcode_yxrrset"].(int)
		oi.Doh_dns_resp_rcode_nxrrset = in["doh_dns_resp_rcode_nxrrset"].(int)
		oi.Doh_dns_resp_rcode_notauth = in["doh_dns_resp_rcode_notauth"].(int)
		oi.Doh_dns_resp_rcode_notzone = in["doh_dns_resp_rcode_notzone"].(int)
		oi.Doh_dns_resp_rcode_other = in["doh_dns_resp_rcode_other"].(int)
		oi.Compression_before_br = in["compression_before_br"].(int)
		oi.Compression_after_br = in["compression_after_br"].(int)
		oi.Compression_before_total = in["compression_before_total"].(int)
		oi.Compression_after_total = in["compression_after_total"].(int)
		oi.Decompression_before_br = in["decompression_before_br"].(int)
		oi.Decompression_after_br = in["decompression_after_br"].(int)
		oi.Decompression_before_total = in["decompression_before_total"].(int)
		oi.Decompression_after_total = in["decompression_after_total"].(int)
		oi.Compress_rsp_br = in["compress_rsp_br"].(int)
		oi.Compress_rsp_total = in["compress_rsp_total"].(int)
		oi.H2up_content_length_alias = in["h2up_content_length_alias"].(int)
		oi.Malformed_h2up_header_value = in["malformed_h2up_header_value"].(int)
		oi.Malformed_h2up_scheme_value = in["malformed_h2up_scheme_value"].(int)
		oi.H2up_with_transfer_encoding = in["h2up_with_transfer_encoding"].(int)
		oi.Multiple_content_length = in["multiple_content_length"].(int)
		oi.Multiple_transfer_encoding = in["multiple_transfer_encoding"].(int)
		oi.Transfer_encoding_and_content_length = in["transfer_encoding_and_content_length"].(int)
		oi.Get_and_payload = in["get_and_payload"].(int)
		oi.H2up_with_host_and_auth = in["h2up_with_host_and_auth"].(int)
		oi.Header_filter_rule_hit = in["header_filter_rule_hit"].(int)
		oi.Current_stream = in["current_stream"].(int)
		oi.Stream_create = in["stream_create"].(int)
		oi.Stream_free = in["stream_free"].(int)
		oi.End_stream_rcvd = in["end_stream_rcvd"].(int)
		oi.End_stream_sent = in["end_stream_sent"].(int)
		oi.Http1_client_idle_timeout = in["http1_client_idle_timeout"].(int)
		oi.Http2_client_idle_timeout = in["http2_client_idle_timeout"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHttpProxyOper(d *schema.ResourceData) edpt.SlbHttpProxyOper {
	var ret edpt.SlbHttpProxyOper
	ret.Oper = getObjectSlbHttpProxyOperOper(d.Get("oper").([]interface{}))
	return ret
}
