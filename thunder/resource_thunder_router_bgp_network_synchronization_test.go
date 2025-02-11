package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_BGP_NETWORK_SYNCHRONIZATION_RESOURCE = `
resource "thunder_router_bgp_network_synchronization" "RouterBgpNetworkTest" {
	network_synchronization = 1
 
}
`

// Acceptance test
func TestAccRouterBgpNetworkSynchronization_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_BGP_NETWORK_SYNCHRONIZATION_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_bgp_network_synchronization.RouterBgpNetworkTest", "network_synchronization", "1"),
				),
			},
		},
	})
}
