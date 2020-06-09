package pingaccess

import (
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func validateWebOrAPI(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "Web" && v != "API" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'Web' or 'API' not %s", v))}
	}
	return nil
}

func validateRuleOrRuleSet(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "Rule" && v != "RuleSet" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'Rule' or 'RuleSet' not %s", v))}
	}
	return nil
}

func validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "SuccessIfAllSucceed" && v != "SuccessIfAnyOneSucceeds" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'SuccessIfAllSucceed' or 'SuccessIfAnyOneSucceeds' not %s", v))}
	}
	return nil
}

func validateAudience(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if len(v) < 1 || len(v) > 31 {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be between 1 and 32 characters not %d", len(v)))}
	}
	return nil
}

func validateCookieType(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "Encrypted" && v != "Signed" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'Encrypted' or 'Signed' not %s", v))}
	}
	return nil
}

func validateOidcLoginType(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "Code" && v != "POST" && v != "x_post" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'Code', 'POST' or 'x_post' not %s", v))}
	}
	return nil
}

func validatePkceChallengeType(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "SHA256" && v != "OFF" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'SHA256' or 'OFF' not %s", v))}
	}
	return nil
}

func validateRequestPreservationType(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "None" && v != "POST" && v != "All" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'None', 'POST' or 'All' not %s", v))}
	}
	return nil
}

func validateWebStorageType(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "SessionStorage" && v != "LocalStorage" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'SessionStorage' or 'LocalStorage' not %s", v))}
	}
	return nil
}

func validateListLocationValue(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "FIRST" && v != "LAST" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'FIRST' or 'LAST' not %s", v))}
	}
	return nil
}

func validateHTTPListenerName(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "ADMIN" && v != "AGENT" && v != "ENGINE" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'ADMIN', 'AGENT' or 'ENGINE' not %s", v))}
	}
	return nil
}

func validateWebSessionSameSite(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "Disabled" && v != "Lax" && v != "None" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'Disabled', 'Lax' or 'None' not %s", v))}
	}
	return nil
}

func validateAuditLevel(value interface{}, path cty.Path) diag.Diagnostics {
	v := value.(string)
	if v != "ON" && v != "OFF" {
		return diag.Diagnostics{diag.FromErr(fmt.Errorf("must be either 'ON' or 'OFF' not %s", v))}
	}
	return nil
}
