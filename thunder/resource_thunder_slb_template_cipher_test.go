package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_TEMPLATE_CIPHER_RESOURCE = `
resource "thunder_slb_template_cipher" "cipher1" {
	cipher_cfg {
		priority = 10
		cipher_suite = "SSL3_RSA_DES_192_CBC3_SHA"
	} 
	name = "testcipher"
	user_tag = "test_user"
}
`

//Acceptance test
func TestAccTemplateCipher_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_CIPHER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_cipher.cipher1", "name", "testcipher"),
					resource.TestCheckResourceAttr("thunder_slb_template_cipher.cipher1", "cipher_cfg.0.priority", "10"),
					resource.TestCheckResourceAttr("thunder_slb_template_cipher.cipher1", "cipher_cfg.0.cipher_suite", "SSL3_RSA_DES_192_CBC3_SHA"),
					resource.TestCheckResourceAttr("thunder_slb_template_cipher.cipher1", "user_tag", "test_user"),
				),
			},
		},
	})
}
