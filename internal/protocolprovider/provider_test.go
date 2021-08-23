package protocol

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	paCfg "github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
)

var conf *paCfg.Config

func TestMain(m *testing.M) {
	conf = paCfg.NewConfig().WithUsername("administrator").WithPassword("2Access").WithEndpoint("https://localhost:9000/pa-admin-api/v3")
	resource.TestMain(m)
}
