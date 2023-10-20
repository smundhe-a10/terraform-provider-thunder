package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudServicesCloudProviderAwsMetrics() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cloud_services_cloud_provider_aws_metrics`: AWS metrics configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceCloudServicesCloudProviderAwsMetricsCreate,
		UpdateContext: resourceCloudServicesCloudProviderAwsMetricsUpdate,
		ReadContext:   resourceCloudServicesCloudProviderAwsMetricsRead,
		DeleteContext: resourceCloudServicesCloudProviderAwsMetricsDelete,
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable AWS Metrics Analytics; 'disable': Disable AWS Metrics Analytics(default);",
			},
			"cps": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable CPS Metrics; 'disable': Disable CPS Metrics;",
			},
			"cpu": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable CPU Metrics; 'disable': Disable CPU Metrics;",
			},
			"disk": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Disk Metrics; 'disable': Disable Disk Metrics;",
			},
			"interfaces": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Interfaces Metrics; 'disable': Disable Interfaces Metrics;",
			},
			"memory": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Memory Metrics; 'disable': Disable Memory Metrics;",
			},
			"namespace": {
				Type: schema.TypeString, Optional: true, Description: "AWS Metrics Analytics Namespace",
			},
			"packet_drop": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Packet Drop Metrics; 'disable': Disable Packet Drop Metrics;",
			},
			"packet_rate": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Packet Rate Metrics; 'disable': Disable Packet Rate Metrics;",
			},
			"server_down_count": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Server Down Count Metrics; 'disable': Disable Server Down Count Metrics;",
			},
			"server_down_percentage": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Server Down Percentage Metrics; 'disable': Disable Server Down Percentage Metrics;",
			},
			"server_error": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Server Error Metrics; 'disable': Disable Server Error Metrics;",
			},
			"sessions": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Sessions Metrics; 'disable': Disable Sessions Metrics;",
			},
			"ssl_cert": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable SSL Cert Metrics; 'disable': Disable SSL Cert Metrics;",
			},
			"throughput": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Throughput Metrics; 'disable': Disable Throughput Metrics;",
			},
			"tps": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable TPS Metrics; 'disable': Disable TPS Metrics;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceCloudServicesCloudProviderAwsMetricsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMetricsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMetrics(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAwsMetricsRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderAwsMetricsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMetricsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMetrics(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCloudServicesCloudProviderAwsMetricsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMetricsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMetrics(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCloudServicesCloudProviderAwsMetricsRead(ctx, d, meta)
	}
	return diags
}

func resourceCloudServicesCloudProviderAwsMetricsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCloudServicesCloudProviderAwsMetricsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCloudServicesCloudProviderAwsMetrics(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCloudServicesCloudProviderAwsMetrics(d *schema.ResourceData) edpt.CloudServicesCloudProviderAwsMetrics {
	var ret edpt.CloudServicesCloudProviderAwsMetrics
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Cps = d.Get("cps").(string)
	ret.Inst.Cpu = d.Get("cpu").(string)
	ret.Inst.Disk = d.Get("disk").(string)
	ret.Inst.Interfaces = d.Get("interfaces").(string)
	ret.Inst.Memory = d.Get("memory").(string)
	ret.Inst.Namespace = d.Get("namespace").(string)
	ret.Inst.PacketDrop = d.Get("packet_drop").(string)
	ret.Inst.PacketRate = d.Get("packet_rate").(string)
	ret.Inst.ServerDownCount = d.Get("server_down_count").(string)
	ret.Inst.ServerDownPercentage = d.Get("server_down_percentage").(string)
	ret.Inst.ServerError = d.Get("server_error").(string)
	ret.Inst.Sessions = d.Get("sessions").(string)
	ret.Inst.SslCert = d.Get("ssl_cert").(string)
	ret.Inst.Throughput = d.Get("throughput").(string)
	ret.Inst.Tps = d.Get("tps").(string)
	//omit uuid
	return ret
}
