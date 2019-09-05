package pingaccess

import (
	"fmt"
)

const cookieBasedClassName = "com.pingidentity.pa.ha.lb.roundrobin.CookieBasedRoundRobinPlugin"
const headerBasedClassName = "com.pingidentity.pa.ha.lb.header.HeaderBasedLoadBalancingPlugin"

func validateWebOrAPI(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if v != "Web" && v != "API" {
		errs = append(errs, fmt.Errorf("%q must be either 'Web' or 'API' not %s", field, v))
	}
	return
}

func validateRuleOrRuleSet(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if v != "Rule" && v != "RuleSet" {
		errs = append(errs, fmt.Errorf("%q must be either 'Rule' or 'RuleSet' not %s", field, v))
	}
	return
}

func validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if v != "SuccessIfAllSucceed" && v != "SuccessIfAnyOneSucceeds" {
		errs = append(errs, fmt.Errorf("%q must be either 'SuccessIfAllSucceed' or 'SuccessIfAnyOneSucceeds' not %s", field, v))
	}
	return
}

func validateAudience(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if len(v) < 1 || len(v) > 31 {
		errs = append(errs, fmt.Errorf("%q must be between 1 and 32 characters not %d", field, len(v)))
	}
	return
}

func validateCookieType(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if v != "Encrypted" && v != "Signed" {
		errs = append(errs, fmt.Errorf("%q must be either 'Encrypted' or 'Signed' not %s", field, v))
	}
	return
}

func validateOidcLoginType(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if v != "Code" && v != "POST" && v != "x_post" {
		errs = append(errs, fmt.Errorf("%q must be either 'Code', 'POST' or 'x_post' not %s", field, v))
	}
	return
}

func validateRequestPreservationType(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if v != "None" && v != "POST" && v != "All" {
		errs = append(errs, fmt.Errorf("%q must be either 'None', 'POST' or 'All' not %s", field, v))
	}
	return
}

func validateWebStorageType(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if v != "SessionStorage" && v != "LocalStorage" {
		errs = append(errs, fmt.Errorf("%q must be either 'SessionStorage' or 'LocalStorage' not %s", field, v))
	}
	return
}

func validateListLocationValue(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if v != "FIRST" && v != "LAST" {
		errs = append(errs, fmt.Errorf("%q must be either 'FIRST' or 'LAST' not %s", field, v))
	}
	return
}

func validateClassNameValue(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if v != cookieBasedClassName && v != headerBasedClassName {
		errs = append(errs, fmt.Errorf("%[1]q must be either %[2]s or %[3]s not %[4]s", field, cookieBasedClassName, headerBasedClassName, v))
	}
	return
}
