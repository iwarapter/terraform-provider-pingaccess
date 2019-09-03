package pingaccess

import (
	"fmt"

	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePingAccessAuthTokenManagement() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessAuthTokenManagementCreate,
		Read:   resourcePingAccessAuthTokenManagementRead,
		Update: resourcePingAccessAuthTokenManagementUpdate,
		Delete: resourcePingAccessAuthTokenManagementDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingAccessAuthTokenManagementSchema(),
	}
}

func resourcePingAccessAuthTokenManagementSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"issuer": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "PingAccessAuthToken",
			Description: "The issuer value to include in auth tokens. PingAccess inserts this value as the iss claim within the auth tokens.",
		},
		"key_roll_enabled": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "This field is true if key rollover is enabled. When false, PingAccess will not rollover keys at the configured interval.",
		},
		"key_roll_period_in_hours": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     24,
			Description: "The interval (in hours) at which PingAccess will roll the keys. Key rollover updates keys at regular intervals to ensure the security of signed auth tokens.",
		},
		"signing_algorithm": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "P-256",
			Description: "The signing algorithm used when creating signed auth tokens.",
		},
	}
}

func resourcePingAccessAuthTokenManagementCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("auth_token_management")
	return resourcePingAccessAuthTokenManagementUpdate(d, m)
}

func resourcePingAccessAuthTokenManagementRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).AuthTokenManagements
	result, _, err := svc.GetAuthTokenManagementCommand()
	if err != nil {
		return fmt.Errorf("Error reading auth token management settings: %s", err)
	}

	return resourcePingAccessAuthTokenManagementReadResult(d, result)
}

func resourcePingAccessAuthTokenManagementUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).AuthTokenManagements
	input := pa.UpdateAuthTokenManagementCommandInput{
		Body: *resourcePingAccessAuthTokenManagementReadData(d),
	}
	result, _, err := svc.UpdateAuthTokenManagementCommand(&input)
	if err != nil {
		return fmt.Errorf("Error updating auth token management settings: %s", err.Error())
	}

	d.SetId("auth_token_management")
	return resourcePingAccessAuthTokenManagementReadResult(d, result)
}

func resourcePingAccessAuthTokenManagementDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).AuthTokenManagements
	_, err := svc.DeleteAuthTokenManagementCommand()
	if err != nil {
		return fmt.Errorf("Error resetting auth token management: %s", err)
	}
	return nil
}

func resourcePingAccessAuthTokenManagementReadResult(d *schema.ResourceData, input *pa.AuthTokenManagementView) (err error) {
	setResourceDataString(d, "issuer", input.Issuer)
	setResourceDataBool(d, "key_roll_enabled", input.KeyRollEnabled)
	setResourceDataInt(d, "key_roll_period_in_hours", input.KeyRollPeriodInHours)
	setResourceDataString(d, "signing_algorithm", input.SigningAlgorithm)
	return nil
}

func resourcePingAccessAuthTokenManagementReadData(d *schema.ResourceData) *pa.AuthTokenManagementView {
	atm := &pa.AuthTokenManagementView{
		Issuer:               String(d.Get("issuer").(string)),
		KeyRollEnabled:       Bool(d.Get("key_roll_enabled").(bool)),
		KeyRollPeriodInHours: Int(d.Get("key_roll_period_in_hours").(int)),
		SigningAlgorithm:     String(d.Get("signing_algorithm").(string)),
	}

	return atm
}
