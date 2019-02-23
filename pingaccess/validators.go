package pingaccess

import (
	"fmt"
)

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
